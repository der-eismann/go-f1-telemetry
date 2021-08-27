package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/der-eismann/telemetry/pkg/util"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	timePeriod     = 5 * time.Second
	maxMessageSize = 1500
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (app *App) generateJSON(hub *Hub) {
	for {
		select {
		case messageType := <-app.c:
			switch messageType {
			case SESSION:
				sessionInfo, err := app.createSessionInfo()
				if err != nil {
					logrus.Error(err)
					return
				}
				hub.broadcast <- sessionInfo
			case LAPDATA:
				rankingInfo, err := app.createRankingInfo()
				if err != nil {
					logrus.Error(err)
					return
				}
				hub.broadcast <- rankingInfo
			case CARSTATUS:
				carStatusInfo, err := app.createStatusInfo()
				if err != nil {
					logrus.Error(err)
					return
				}
				hub.broadcast <- carStatusInfo
			case MOTION:
				positionInfo, err := app.createPositionInfo()
				if err != nil {
					logrus.Error(err)
					return
				}
				hub.broadcast <- positionInfo
			case CLASSIFICATION:
				if app.Data.PacketSessionData.SessionType == 10 {
					raceResult, err := app.createRaceResult()
					if err != nil {
						logrus.Error(err)
						return
					}
					hub.broadcast <- raceResult
				}
			}
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write([]byte("----"))
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (app *App) serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		EnableCompression: true,
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			logrus.Println(err)
		}
		return
	}

	client := &Client{hub: hub, conn: ws, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

func (app *App) createSessionInfo() ([]byte, error) {
	session := Session{
		MessageType:      "SessionInfo",
		Weather:          app.Data.PacketSessionData.Weather,
		TrackTemperature: app.Data.PacketSessionData.TrackTemperature,
		AirTemperature:   app.Data.PacketSessionData.AirTemperature,
		Position:         app.Data.PacketLapData.LapData[app.Data.PlayerArrayID].CarPosition,
		TotalCars:        app.Data.NumActiveCars,
		LapsDone:         app.Data.PacketLapData.LapData[app.Data.PlayerArrayID].CurrentLapNum,
		TotalLaps:        app.Data.PacketSessionData.TotalLaps,
		TrackLength:      app.Data.PacketSessionData.TrackLength,
		SessionType:      util.SessionType[app.Data.PacketSessionData.SessionType],
		Track:            util.TrackName[app.Data.PacketSessionData.TrackID][0],
		Country:          util.TrackName[app.Data.PacketSessionData.TrackID][1],
		SessionTimeLeft:  app.Data.PacketSessionData.SessionTimeLeft,
		SessionDuration:  app.Data.PacketSessionData.SessionDuration,
	}
	forecasts := []WeatherForecast{}

	for i := uint8(0); i < app.Data.PacketSessionData.NumWeatherForecastSamples; i++ {
		forecast := WeatherForecast{
			SessionType: util.SessionType[app.Data.PacketSessionData.WeatherForecastSamples[i].SessionType],
			TimeOffset:  app.Data.PacketSessionData.WeatherForecastSamples[i].TimeOffset,
			Weather:     app.Data.PacketSessionData.WeatherForecastSamples[i].Weather,
		}
		forecasts = append(forecasts, forecast)
	}
	session.WeatherForecasts = forecasts

	json, err := json.Marshal(session)
	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

func (app *App) createRankingInfo() ([]byte, error) {
	rankingData := [22]RankingData{}
	activeCars := uint8(0)
	for i := uint8(0); i < 22; i++ {
		rankingData[i].IsPlayer = false
		if i == app.Data.PlayerArrayID {
			rankingData[i].IsPlayer = true
		}
		rankingData[i].LapData = app.Data.PacketLapData.LapData[i]
		rankingData[i].AI = app.Data.PacketParticipantsData.Participants[i].AiControlled
		rankingData[i].Team = util.Teams[app.Data.PacketParticipantsData.Participants[i].TeamID][0]
		rankingData[i].TeamColor = util.Teams[app.Data.PacketParticipantsData.Participants[i].TeamID][1]
		rankingData[i].Tyres = util.VisualTyreCompound[app.Data.PacketCarStatusData.CarStatusData[i].VisualTyreCompound]
		rankingData[i].TyreAge = app.Data.PacketCarStatusData.CarStatusData[i].TyresAgeLaps
		rankingData[i].Nation = util.Nationality[app.Data.PacketParticipantsData.Participants[i].Nationality][1]
		rankingData[i].DrivenDistance = "0.0 km"
		if app.Data.PacketLapData.LapData[i].TotalDistance > 0 {
			rankingData[i].DrivenDistance = fmt.Sprintf("%.1f km", app.Data.PacketLapData.LapData[i].TotalDistance/1000)
		}
		if rankingData[i].AI == 1 {
			rankingData[i].Name = util.Drivers[app.Data.PacketParticipantsData.Participants[i].DriverID][0] + " " + util.Drivers[app.Data.PacketParticipantsData.Participants[i].DriverID][1]
		} else if app.Data.PacketSessionData.NetworkGame == 0 {
			upperCaseName := bytes.Trim(app.Data.PacketParticipantsData.Participants[i].Name[:], "\x00")
			if len(upperCaseName) > 0 {
				rankingData[i].Name = string(upperCaseName[0]) + strings.ToLower(string(upperCaseName[1:]))
			} else {
				rankingData[i].Name = ""
			}
		} else {
			rankingData[i].Name = "Player " + strconv.Itoa(int(app.Data.PacketParticipantsData.Participants[i].DriverID))
		}
		if rankingData[i].Tyres == "" && app.Data.PacketCarStatusData.CarStatusData[i].ActualTyreCompound != 0 {
			logrus.Errorf("Unknown tyre - Actual %d - Visual %d", app.Data.PacketCarStatusData.CarStatusData[i].ActualTyreCompound, app.Data.PacketCarStatusData.CarStatusData[i].VisualTyreCompound)
		}
		if rankingData[i].Tyres == "" {
			rankingData[i].Tyres = "dry"
		}
		if app.Data.PacketParticipantsData.Participants[i].RaceNumber != 0 {
			activeCars++
		}
		if app.Data.BestLapTyres[i].BestLapTimeInMS != app.Data.PacketSessionHistoryData[i].LapHistoryData[app.Data.PacketSessionHistoryData[i].BestLapTimeLapNum].LapTimeInMS {
			if app.Data.PacketCarStatusData.CarStatusData[i].VisualTyreCompound == 0 {
				continue
			}
			app.Data.BestLapTyres[i].BestLapTimeInMS = app.Data.PacketSessionHistoryData[i].LapHistoryData[app.Data.PacketSessionHistoryData[i].BestLapTimeLapNum].LapTimeInMS
			app.Data.BestLapTyres[i].Tyres = util.VisualTyreCompound[app.Data.PacketCarStatusData.CarStatusData[i].VisualTyreCompound]
			rankingData[i].TyresBestLap = util.VisualTyreCompound[app.Data.PacketCarStatusData.CarStatusData[i].VisualTyreCompound]
		} else {
			rankingData[i].TyresBestLap = app.Data.BestLapTyres[i].Tyres
		}
	}
	app.Data.NumActiveCars = activeCars

	sort.SliceStable(rankingData[:app.Data.NumActiveCars], func(i, j int) bool {
		return rankingData[i].LapData.CarPosition < rankingData[j].LapData.CarPosition
	})
	ranking := Ranking{
		MessageType: "RankingInfo",
		SessionType: util.SessionType[app.Data.PacketSessionData.SessionType],
		NumCars:     app.Data.NumActiveCars,
		RankingData: rankingData,
	}

	json, err := json.Marshal(ranking)
	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

func (app *App) createStatusInfo() ([]byte, error) {
	carStatus := CarStatus{
		MessageType:             "CarStatusInfo",
		BrakesTemperature:       app.Data.PacketCarTelemetryData.CarTelemetryData[app.Data.PlayerArrayID].BrakesTemperature,
		TyresSurfaceTemperature: app.Data.PacketCarTelemetryData.CarTelemetryData[app.Data.PlayerArrayID].TyresSurfaceTemperature,
		TyresInnerTemperature:   app.Data.PacketCarTelemetryData.CarTelemetryData[app.Data.PlayerArrayID].TyresInnerTemperature,
		EngineTemperature:       app.Data.PacketCarTelemetryData.CarTelemetryData[app.Data.PlayerArrayID].EngineTemperature,
		FuelRemainingLaps:       app.Data.PacketCarStatusData.CarStatusData[app.Data.PlayerArrayID].FuelRemainingLaps,
		TyreVisualCompound:      app.Data.PacketCarStatusData.CarStatusData[app.Data.PlayerArrayID].VisualTyreCompound,
		ErsStoreEnergy:          app.Data.PacketCarStatusData.CarStatusData[app.Data.PlayerArrayID].ERSStoreEnergy,
		ErsDeployMode:           app.Data.PacketCarStatusData.CarStatusData[app.Data.PlayerArrayID].ERSDeployMode,
		TyresWear:               app.Data.PacketCarDamageData.CarDamageData[app.Data.PlayerArrayID].TyresWear,
		FrontLeftWingDamage:     app.Data.PacketCarDamageData.CarDamageData[app.Data.PlayerArrayID].FrontLeftWingDamage,
		FrontRightWingDamage:    app.Data.PacketCarDamageData.CarDamageData[app.Data.PlayerArrayID].FrontRightWingDamage,
		RearWingDamage:          app.Data.PacketCarDamageData.CarDamageData[app.Data.PlayerArrayID].RearWingDamage,
		DiffuserDamage:          app.Data.PacketCarDamageData.CarDamageData[app.Data.PlayerArrayID].DiffuserDamage,
		SidepodDamage:           app.Data.PacketCarDamageData.CarDamageData[app.Data.PlayerArrayID].SidepodDamage,
	}

	json, err := json.Marshal(carStatus)
	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

func (app *App) createRaceResult() ([]byte, error) {
	resultData := [22]RaceResultData{}
	for i := uint8(0); i < app.Data.PacketFinalClassificationData.NumCars; i++ {
		resultData[i] = RaceResultData{
			Position:      app.Data.PacketFinalClassificationData.ClassificationData[i].Position,
			GridPosition:  app.Data.PacketFinalClassificationData.ClassificationData[i].GridPosition,
			NumLaps:       app.Data.PacketFinalClassificationData.ClassificationData[i].NumLaps,
			Team:          util.Teams[app.Data.PacketParticipantsData.Participants[i].TeamID][0],
			TeamColor:     util.Teams[app.Data.PacketParticipantsData.Participants[i].TeamID][1],
			Nation:        util.Nationality[app.Data.PacketParticipantsData.Participants[i].Nationality][1],
			Points:        app.Data.PacketFinalClassificationData.ClassificationData[i].Points,
			NumPitStops:   app.Data.PacketFinalClassificationData.ClassificationData[i].NumPitStops,
			ResultStatus:  app.Data.PacketFinalClassificationData.ClassificationData[i].ResultStatus,
			BestLapTime:   float32(app.Data.PacketFinalClassificationData.ClassificationData[i].BestLapTimeInMS) / 1000,
			TotalRaceTime: app.Data.PacketFinalClassificationData.ClassificationData[i].TotalRaceTime,
			PenaltiesTime: app.Data.PacketFinalClassificationData.ClassificationData[i].PenaltiesTime,
			NumTyreStints: app.Data.PacketFinalClassificationData.ClassificationData[i].NumTyreStints,
		}
		if app.Data.PacketParticipantsData.Participants[i].AiControlled == 1 {
			resultData[i].Name = util.Drivers[app.Data.PacketParticipantsData.Participants[i].DriverID][0] + " " + util.Drivers[app.Data.PacketParticipantsData.Participants[i].DriverID][1]
		} else if app.Data.PacketSessionData.NetworkGame == 0 {
			upperCaseName := bytes.Trim(app.Data.PacketParticipantsData.Participants[i].Name[:], "\x00")
			if len(upperCaseName) > 0 {
				resultData[i].Name = string(upperCaseName[0]) + strings.ToLower(string(upperCaseName[1:]))
			} else {
				resultData[i].Name = ""
			}
		} else {
			resultData[i].Name = "Player " + strconv.Itoa(int(app.Data.PacketParticipantsData.Participants[i].DriverID))
		}
		for j := uint8(0); j < app.Data.PacketFinalClassificationData.ClassificationData[i].NumTyreStints; j++ {
			resultData[i].Tyres[j] = util.VisualTyreCompound[app.Data.PacketFinalClassificationData.ClassificationData[i].TyreStintsVisual[j]]
		}
	}

	sort.SliceStable(resultData[:app.Data.PacketFinalClassificationData.NumCars], func(i, j int) bool {
		return resultData[i].Position < resultData[j].Position
	})

	raceResult := RaceResult{
		MessageType:    "RaceResult",
		NumCars:        app.Data.PacketFinalClassificationData.NumCars,
		RaceResultData: resultData,
	}

	json, err := json.Marshal(raceResult)
	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

func (app *App) createPositionInfo() ([]byte, error) {
	positions := [22]Position{}
	trackID := app.Data.PacketSessionData.TrackID

	for i := uint8(0); i < 22; i++ {
		x, z := rotateCoordinates(
			app.Data.PacketMotionData.CarMotionData[i].WorldPositionX,
			app.Data.PacketMotionData.CarMotionData[i].WorldPositionZ,
			util.BorderCoordinates[trackID].Rotation,
		)

		positions[i].WorldPositionX = int(x) - util.BorderCoordinates[trackID].MinRotated.X
		positions[i].WorldPositionZ = int(z) - util.BorderCoordinates[trackID].MinRotated.Z
		positions[i].TeamColor = util.Teams[app.Data.PacketParticipantsData.Participants[i].TeamID][1]
		if app.Data.PacketParticipantsData.Participants[i].AiControlled == 0 {
			if len(string(bytes.Trim(app.Data.PacketParticipantsData.Participants[i].Name[:], "\x00"))) > 2 {
				positions[i].Name = strings.ToUpper(string(bytes.Trim(app.Data.PacketParticipantsData.Participants[i].Name[:], "\x00"))[0:3])
			}
		} else {
			if len(util.Drivers[app.Data.PacketParticipantsData.Participants[i].DriverID][1]) > 2 {
				runeName := []rune(strings.ToUpper(util.Drivers[app.Data.PacketParticipantsData.Participants[i].DriverID][1]))
				positions[i].Name = string(runeName[0:3])
			}
		}
	}

	positionData := PositionData{
		MessageType: "PositionInfo",
		NumCars:     app.Data.NumActiveCars,
		Positions:   positions,
		TrackID:     trackID,
		Width:       -util.BorderCoordinates[trackID].MinRotated.X + util.BorderCoordinates[trackID].MaxRotated.X,
		Height:      -util.BorderCoordinates[trackID].MinRotated.Z + util.BorderCoordinates[trackID].MaxRotated.Z,
		PlayerID:    app.Data.PlayerArrayID,
	}

	json, err := json.Marshal(positionData)
	if err != nil {
		return []byte{}, err
	}

	return json, nil
}

func rotateCoordinates(x, z, deg float32) (float64, float64) {
	rotation := float64(deg) * (math.Pi / 180)
	xRot := (float64(x) * math.Cos(rotation)) - (float64(z) * math.Sin(rotation))
	zRot := (float64(x) * math.Sin(rotation)) + (float64(z) * math.Cos(rotation))
	return xRot, zRot
}
