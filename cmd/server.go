package cmd

import (
	"bytes"
	"context"
	"encoding/binary"
	"net"
	"net/http"
	"time"

	"github.com/der-eismann/telemetry/pkg/util"
	"github.com/getlantern/systray"
	"github.com/gobuffalo/packr/v2"
	"github.com/rebuy-de/rebuy-go-sdk/v3/pkg/cmdutil"
	"github.com/sirupsen/logrus"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

// Assignment of names to Packet IDs
const (
	MOTION         = 0
	SESSION        = 1
	LAPDATA        = 2
	EVENT          = 3
	PARTICIPANTS   = 4
	CARSETUPS      = 5
	CARTELEMETRY   = 6
	CARSTATUS      = 7
	CLASSIFICATION = 8
	LOBBYINFO      = 9
	HEADER         = 255
)

// Sizes of the different UDP packets
var bufferSizes = map[int]int{
	HEADER:         24,
	MOTION:         1464,
	SESSION:        251,
	LAPDATA:        1190,
	EVENT:          32,
	PARTICIPANTS:   1213,
	CARSETUPS:      1102,
	CARTELEMETRY:   1307,
	CARSTATUS:      1344,
	CLASSIFICATION: 839,
	LOBBYINFO:      1169,
}

type App struct {
	Data Telemetry
	c    chan int
}

// Telemetry is a struct that contains all available telemetry data
type Telemetry struct {
	PlayerArrayID                 uint8
	SessionID                     uint64
	BestLapTyres                  [22]BestLapTyres
	NumActiveCars                 uint8
	PacketMotionData              util.PacketMotionData
	PacketSessionData             util.PacketSessionData
	PacketLapData                 util.PacketLapData
	PacketParticipantsData        util.PacketParticipantsData
	PacketCarSetupData            util.PacketCarSetupData
	PacketCarTelemetryData        util.PacketCarTelemetryData
	PacketCarStatusData           util.PacketCarStatusData
	PacketFinalClassificationData util.PacketFinalClassificationData
	PacketLobbyInfoData           util.PacketLobbyInfoData
}

type BestLapTyres struct {
	BestLapTime float32
	Tyres       string
}

func systrayOnReady() {
	systray.SetTemplateIcon(util.Icon, util.Icon)
	systray.SetTitle("F1 Telemetry")
	systray.SetTooltip("F1 Telemetry")
	mOpenUI := systray.AddMenuItem("Open UI", "Open Website in browser")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			select {
			case <-mQuitOrig.ClickedCh:
				cmdutil.Exit(0)
			case <-mOpenUI.ClickedCh:
				open.Run("http://localhost:8080")
			}
		}
	}()
}

func (app *App) Listen(ctx context.Context, cmd *cobra.Command, args []string) {
	go systray.Run(systrayOnReady, nil)

	app.c = make(chan int)
	lc := net.ListenConfig{}
	pc, err := lc.ListenPacket(ctx, "udp", ":20777")
	if err != nil {
		logrus.Fatal(err)
	}
	defer pc.Close()
	logrus.Printf("server listening on %s\n", pc.LocalAddr().String())

	hub := newHub()
	go hub.run()
	go app.generateJSON(hub)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		app.serveWs(hub, w, r)
	})

	box := packr.New("static", "../static")
	http.Handle("/", http.FileServer(box))

	go http.ListenAndServe("127.0.0.1:8080", nil)
	//open.Run("http://localhost:8080")

	for {
		buf := make([]byte, 2048)
		_, _, err := pc.ReadFrom(buf)
		if err != nil {
			logrus.Error(err)
		}
		start := time.Now()

		r := bytes.NewReader(buf[:bufferSizes[HEADER]])
		packetHeader := util.PacketHeader{}
		if err := binary.Read(r, binary.LittleEndian, &packetHeader); err != nil {
			logrus.Errorln("binary.Read failed:", err)
		}
		app.Data.PlayerArrayID = packetHeader.PlayerCarIndex
		if app.Data.SessionID != packetHeader.SessionUID {
			app.Data.BestLapTyres = [22]BestLapTyres{}
			app.Data.SessionID = packetHeader.SessionUID
		}

		switch packetHeader.PacketID {
		case MOTION:
			r := bytes.NewReader(buf[bufferSizes[HEADER]:bufferSizes[MOTION]])
			if err := binary.Read(r, binary.LittleEndian, &app.Data.PacketMotionData); err != nil {
				logrus.Errorln("MOTION binary.Read failed:", err)
			}
			app.c <- MOTION
			elapsed := time.Since(start)
			logrus.Debugf("MOTION read took %s", elapsed)
		case SESSION:
			r := bytes.NewReader(buf[bufferSizes[HEADER]:bufferSizes[SESSION]])
			if err := binary.Read(r, binary.LittleEndian, &app.Data.PacketSessionData); err != nil {
				logrus.Errorln("SESSION binary.Read failed:", err)
			}
			app.c <- SESSION
			elapsed := time.Since(start)
			logrus.Debugf("SESSION read took %s", elapsed)
		case LAPDATA:
			r := bytes.NewReader(buf[bufferSizes[HEADER]:bufferSizes[LAPDATA]])
			if err := binary.Read(r, binary.LittleEndian, &app.Data.PacketLapData); err != nil {
				logrus.Errorln("LAPDATA binary.Read failed:", err)
			}
			app.c <- LAPDATA
			elapsed := time.Since(start)
			logrus.Debugf("LAPDATA read took %s", elapsed)
		case PARTICIPANTS:
			r := bytes.NewReader(buf[bufferSizes[HEADER]:bufferSizes[PARTICIPANTS]])
			if err := binary.Read(r, binary.LittleEndian, &app.Data.PacketParticipantsData); err != nil {
				logrus.Errorln("PARTICIPANTS binary.Read failed:", err)
			}
			app.c <- PARTICIPANTS
			elapsed := time.Since(start)
			logrus.Debugf("PARTICIPANTS read took %s", elapsed)
		case CARSETUPS:
			r := bytes.NewReader(buf[bufferSizes[HEADER]:bufferSizes[CARSETUPS]])
			if err := binary.Read(r, binary.LittleEndian, &app.Data.PacketCarSetupData); err != nil {
				logrus.Errorln("CARSETUPS binary.Read failed:", err)
			}
			app.c <- CARSETUPS
			elapsed := time.Since(start)
			logrus.Debugf("CARSETUPS read took %s", elapsed)
		case CARTELEMETRY:
			r := bytes.NewReader(buf[bufferSizes[HEADER]:bufferSizes[CARTELEMETRY]])
			if err := binary.Read(r, binary.LittleEndian, &app.Data.PacketCarTelemetryData); err != nil {
				logrus.Errorln("CARTELEMETRY binary.Read failed:", err)
			}
			app.c <- CARTELEMETRY
			elapsed := time.Since(start)
			logrus.Debugf("CARTELEMETRY read took %s", elapsed)
		case CARSTATUS:
			r := bytes.NewReader(buf[bufferSizes[HEADER]:bufferSizes[CARSTATUS]])
			if err := binary.Read(r, binary.LittleEndian, &app.Data.PacketCarStatusData); err != nil {
				logrus.Errorln("CARSTATUS binary.Read failed:", err)
			}
			app.c <- CARSTATUS
			elapsed := time.Since(start)
			logrus.Debugf("CARSTATUS read took %s", elapsed)
		case CLASSIFICATION:
			r := bytes.NewReader(buf[bufferSizes[HEADER]:bufferSizes[CLASSIFICATION]])
			if err := binary.Read(r, binary.LittleEndian, &app.Data.PacketFinalClassificationData); err != nil {
				logrus.Errorln("CLASSIFICATION binary.Read failed:", err)
			}
			app.c <- CLASSIFICATION
			elapsed := time.Since(start)
			logrus.Debugf("CLASSIFICATION read took %s", elapsed)
		case LOBBYINFO:
			r := bytes.NewReader(buf[bufferSizes[HEADER]:bufferSizes[LOBBYINFO]])
			if err := binary.Read(r, binary.LittleEndian, &app.Data.PacketLobbyInfoData); err != nil {
				logrus.Errorln("LOBBYINFO binary.Read failed:", err)
			}
			app.c <- LOBBYINFO
			elapsed := time.Since(start)
			logrus.Debugf("LOBBYINFO read took %s", elapsed)
		}
	}
}
