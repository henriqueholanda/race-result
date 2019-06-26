package pilot_statistics

import (
	"github.com/henriqueholanda/race-result/domain/models"
	"github.com/henriqueholanda/race-result/infrastructure/helper"
	"math"
)

type RaceDuration struct {
}

func NewRaceDuration() *RaceDuration {
	return &RaceDuration{}
}

func (rca *RaceDuration) Generate(raceResult models.RaceStatistics, pilotStatistics models.RacePilotStatistics) models.RacePilotStatistics {
	var raceDuration int
	var milliseconds int

	for _, data := range raceResult {
		minutes, seconds, millisecondsInt := helper.SplitLapTime(data.LapTime)
		raceDuration += seconds + 60*minutes
		milliseconds += millisecondsInt
	}

	raceDuration += int(math.Round(float64(milliseconds) / float64(1000)))

	pilotStatistics.RaceDuration = helper.FormatBySeconds(raceDuration)

	return pilotStatistics
}
