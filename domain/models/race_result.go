package models

type RaceResult struct {
	Classification         []string
	PilotStatistics        map[string]RacePilotStatistics
	PilotNumberRaceBestLap string
	PilotNameRaceBestLap   string
	RaceBestLapTime        string
}

type RaceStatistics []RacePilotStatistics
