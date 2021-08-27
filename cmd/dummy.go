package cmd

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/der-eismann/telemetry/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func (app *App) DummyPacket(ctx context.Context, cmd *cobra.Command, args []string) {

	type SessionPacket struct {
		Header  util.PacketHeader
		Session util.PacketSessionData
	}

	type LapDataPacket struct {
		Header  util.PacketHeader
		LapData util.PacketLapData
	}

	type ParticipantDataPacket struct {
		Header           util.PacketHeader
		ParticipantsData util.PacketParticipantsData
	}

	type ClassificationDataPacket struct {
		Header                  util.PacketHeader
		FinalClassificationData util.PacketFinalClassificationData
	}

	header := util.PacketHeader{
		PacketFormat:            2020,
		GameMajorVersion:        1,
		GameMinorVersion:        10,
		PacketVersion:           1,
		PacketID:                0,
		SessionUID:              123536456,
		SessionTime:             1.23436,
		FrameIdentifier:         1231251325,
		PlayerCarIndex:          5,
		SecondaryPlayerCarIndex: 6,
	}

	sessionHeader := header
	sessionHeader.PacketID = SESSION

	lapdataHeader := header
	lapdataHeader.PacketID = LAPDATA

	participantsdataHeader := header
	participantsdataHeader.PacketID = PARTICIPANTS

	// classificationDataHeader := header
	// classificationDataHeader.PacketID = CLASSIFICATION

	// classification := util.PacketFinalClassificationData{}
	// classification.NumCars = 20
	// for i := uint8(0); i < 21; i++ {
	// 	classification.ClassificationData[i].Position = i + 1
	// 	classification.ClassificationData[i].NumLaps = 32
	// 	classification.ClassificationData[i].GridPosition = 20 - i
	// 	classification.ClassificationData[i].Points = i
	// 	classification.ClassificationData[i].NumPitStops = uint8(rand.Intn(5))
	// 	classification.ClassificationData[i].ResultStatus = 3
	// 	classification.ClassificationData[i].BestLapTime = rand.Float32() + 1
	// 	classification.ClassificationData[i].TotalRaceTime = rand.Float64() + float64(i)
	// 	classification.ClassificationData[i].PenaltiesTime = (i + 1) * 2
	// 	classification.ClassificationData[i].NumPenalties = i
	// 	classification.ClassificationData[i].NumTyreStints = 3
	// 	classification.ClassificationData[i].TyreStintsActual = [8]uint8{7, 8, 16}
	// 	classification.ClassificationData[i].TyreStintsVisual = [8]uint8{7, 8, 16}
	// }

	// classificationPacket := ClassificationDataPacket{
	// 	Header:                  classificationDataHeader,
	// 	FinalClassificationData: classification,
	// }

	packet := SessionPacket{
		Header: sessionHeader,
		Session: util.PacketSessionData{
			Weather:                   1,
			TrackTemperature:          34,
			AirTemperature:            20,
			TotalLaps:                 16,
			TrackLength:               5234,
			SessionType:               1,
			TrackID:                   4,
			Formula:                   0,
			SessionTimeLeft:           4500,
			SessionDuration:           9000,
			PitSpeedLimit:             60,
			GamePaused:                0,
			IsSpectating:              0,
			SpectatorCarIndex:         0,
			SLIProNativeSupport:       0,
			NumMarshalZones:           16,
			MarshalZones:              [21]util.MarshalZone{},
			SafetyCarStatus:           1,
			NetworkGame:               1,
			NumWeatherForecastSamples: 4,
			WeatherForecastSamples: [56]util.WeatherForecastSample{
				{
					Weather:     0,
					TimeOffset:  0,
					SessionType: 1,
				},
				{
					Weather:     1,
					TimeOffset:  5,
					SessionType: 2,
				},
				{
					Weather:     2,
					TimeOffset:  10,
					SessionType: 5,
				},
				{
					Weather:     4,
					TimeOffset:  15,
					SessionType: 10,
				},
			},
		},
	}
	participantsData := util.PacketParticipantsData{}
	participantsData.NumActiveCars = 21
	lapdata := util.PacketLapData{}
	for i := uint8(0); i < participantsData.NumActiveCars; i++ {
		lapdata.LapData[i].LapDistance = rand.Float32()
		lapdata.LapData[i].TotalDistance = rand.Float32()
		lapdata.LapData[i].SafetyCarDelta = rand.Float32()
		lapdata.LapData[i].CarPosition = i + 1
		lapdata.LapData[i].CurrentLapNum = uint8(rand.Intn(30))
		lapdata.LapData[i].PitStatus = uint8(rand.Intn(2))
		lapdata.LapData[i].Sector = uint8(rand.Intn(2))
		lapdata.LapData[i].CurrentLapInvalid = uint8(rand.Intn(1))
		lapdata.LapData[i].Penalties = uint8(rand.Intn(3) * 3)
		lapdata.LapData[i].GridPosition = i + 1
		lapdata.LapData[i].DriverStatus = uint8(rand.Intn(4))
		lapdata.LapData[i].ResultStatus = uint8(rand.Intn(6))

	}

	lapdataPacket := LapDataPacket{
		Header:  lapdataHeader,
		LapData: lapdata,
	}

	for i := uint8(0); i < participantsData.NumActiveCars; i++ {
		participantsData.Participants[i].AiControlled = 1
		participantsData.Participants[i].DriverID = i + 19
		participantsData.Participants[i].TeamID = uint8(rand.Intn(10))
		participantsData.Participants[i].RaceNumber = uint8(rand.Intn(100))
		participantsData.Participants[i].Nationality = uint8(rand.Intn(60))
		participantsData.Participants[i].Name = [48]byte{}
		participantsData.Participants[i].YourTelemetry = 1
	}

	participantsDataPacket := ParticipantDataPacket{
		Header:           participantsdataHeader,
		ParticipantsData: participantsData,
	}

	buf := new(bytes.Buffer)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:20777")
	if err != nil {
		logrus.Error(err)
	}
	pc, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		logrus.Error(err)
	}

	for {
		select {
		case <-ticker.C:
			packet.Session.SessionTimeLeft = packet.Session.SessionTimeLeft - 1
			err := binary.Write(buf, binary.LittleEndian, packet)
			if err != nil {
				fmt.Println("binary.Write failed:", err)
			}

			n, err := pc.Write(buf.Bytes()[:bufferSizes[SESSION]])
			if err != nil {
				logrus.Error(err)
			}
			buf.Reset()
			logrus.Printf("%d bytes sent", n)

			err = binary.Write(buf, binary.LittleEndian, lapdataPacket)
			if err != nil {
				fmt.Println("binary.Write failed:", err)
			}

			n, err = pc.Write(buf.Bytes()[:bufferSizes[LAPDATA]])
			if err != nil {
				logrus.Error(err)
			}
			buf.Reset()
			logrus.Printf("%d bytes sent", n)

			err = binary.Write(buf, binary.LittleEndian, participantsDataPacket)
			if err != nil {
				fmt.Println("binary.Write failed:", err)
			}

			n, err = pc.Write(buf.Bytes()[:bufferSizes[PARTICIPANTS]])
			if err != nil {
				logrus.Error(err)
			}
			buf.Reset()
			logrus.Printf("%d bytes sent", n)

			// err = binary.Write(buf, binary.LittleEndian, classificationPacket)
			// if err != nil {
			// 	fmt.Println("binary.Write failed:", err)
			// }

			// n, err = pc.Write(buf.Bytes()[:bufferSizes[CLASSIFICATION]])
			// if err != nil {
			// 	logrus.Error(err)
			// }
			// buf.Reset()
			// logrus.Printf("%d bytes sent", n)
		}
	}

}
