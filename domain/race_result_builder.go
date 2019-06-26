package domain

import (
	"github.com/henriqueholanda/race-result/domain/models"
	"github.com/henriqueholanda/race-result/domain/pilot_statistics"
)

type RaceResultBuilder struct {
	statistics []pilot_statistics.StatisticsInterface
}

func NewRaceResultBuilder(statistics ...pilot_statistics.StatisticsInterface) *RaceResultBuilder {
	return &RaceResultBuilder{
		statistics: statistics,
	}
}

func (rrb *RaceResultBuilder) Build(raceStatistics models.RaceStatistics) models.RacePilotStatistics {
	pilotStatistics := models.RacePilotStatistics{
		LapAmount: len(raceStatistics),
	}

	for _, statistic := range rrb.statistics {
		pilotStatistics = statistic.Generate(raceStatistics, pilotStatistics)
	}

	return pilotStatistics
}
