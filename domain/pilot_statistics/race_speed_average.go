package pilot_statistics

import (
	"github.com/henriqueholanda/race-result/domain/models"
	"github.com/henriqueholanda/race-result/infrastructure/helper"
)

type RaceSpeedAverage struct {
}

func NewRaceSpeedAverage() *RaceSpeedAverage {
	return &RaceSpeedAverage{}
}

func (rca *RaceSpeedAverage) Generate(raceResult models.RaceStatistics, pilotStatistics models.RacePilotStatistics) models.RacePilotStatistics {
	var speedAverage float64

	for _, data := range raceResult {
		speedAverage += helper.FormatLapSpeed(data.LapSpeedAverage)
	}

	pilotStatistics.SpeedRaceAverage = speedAverage / float64(len(raceResult))

	return pilotStatistics
}
