package cmd

import "github.com/der-eismann/telemetry/pkg/util"

type Ranking struct {
	MessageType string
	SessionType string
	NumCars     uint8
	RankingData [22]RankingData
}

type RankingData struct {
	LapData        util.LapData
	IsPlayer       bool
	AI             uint8
	Team           string
	TeamColor      string
	Tyres          string
	TyresBestLap   string
	TyreAge        uint8
	Name           string
	Nation         string
	DrivenDistance string
}

type Session struct {
	MessageType      string
	Weather          uint8
	TrackTemperature int8
	AirTemperature   int8
	Position         uint8
	TotalCars        uint8
	LapsDone         uint8
	TotalLaps        uint8
	TrackLength      uint16
	SessionType      string
	Track            string
	Country          string
	SessionTimeLeft  uint16
	SessionDuration  uint16
	WeatherForecasts []WeatherForecast
}

type WeatherForecast struct {
	SessionType string
	TimeOffset  uint8
	Weather     uint8
}

type CarStatus struct {
	MessageType             string
	BrakesTemperature       [4]uint16
	TyresSurfaceTemperature [4]uint8
	TyresInnerTemperature   [4]uint8
	EngineTemperature       uint16
	TyresPressure           [4]float32
	FuelRemainingLaps       float32
	TyresWear               [4]uint8
	TyreVisualCompound      uint8
	TyresDamage             [4]uint8
	FrontLeftWingDamage     uint8
	FuelMix                 uint8
	FrontRightWingDamage    uint8
	RearWingDamage          uint8
	EngineDamage            uint8
	GearBoxDamage           uint8
	ErsStoreEnergy          float32
	ErsDeployMode           uint8
}

// RaceResult contains all information about the end of a race
// similar to PacketFinalClassificationData, but stripped some fields
type RaceResult struct {
	MessageType    string
	NumCars        uint8
	RaceResultData [22]RaceResultData
}

type RaceResultData struct {
	Position      uint8
	GridPosition  uint8
	NumLaps       uint8
	Points        uint8
	Team          string
	TeamColor     string
	Name          string
	Nation        string
	NumPitStops   uint8
	ResultStatus  uint8
	BestLapTime   float32
	TotalRaceTime float64
	PenaltiesTime uint8
	NumTyreStints uint8
	Tyres         [8]string
}

type PositionData struct {
	MessageType string
	NumCars     uint8
	Width       int
	Height      int
	TrackID     int8
	PlayerID    uint8
	Positions   [22]Position
}

type Position struct {
	WorldPositionX int
	WorldPositionZ int
	TeamColor      string
	Name           string
}
