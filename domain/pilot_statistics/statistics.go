package pilot_statistics

import (
	"github.com/henriqueholanda/race-result/domain/models"
)

type StatisticsInterface interface {
	Generate(raceResult models.RaceStatistics, pilotStatistics models.RacePilotStatistics) models.RacePilotStatistics
}
