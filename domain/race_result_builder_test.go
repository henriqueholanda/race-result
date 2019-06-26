package domain_test

import (
	"github.com/henriqueholanda/race-result/domain"
	"github.com/henriqueholanda/race-result/domain/models"
	"github.com/henriqueholanda/race-result/domain/pilot_statistics"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestShouldReturnRacePilotStatistics(t *testing.T) {
	racePilotStatistics := []models.RacePilotStatistics{}

	pilot := models.RacePilotStatistics{Number: "001", Name: "Senna", LapSpeedAverage: "44.444"}
	pilot.LapTime = "01:01.333"
	pilot.LapSpeedAverage = "235,0"
	racePilotStatistics = append(racePilotStatistics, pilot)
	pilot.LapTime = "01:01.000"
	pilot.LapSpeedAverage = "240,0"
	racePilotStatistics = append(racePilotStatistics, pilot)
	pilot.LapTime = "01:03.900"
	pilot.LapSpeedAverage = "245,0"
	racePilotStatistics = append(racePilotStatistics, pilot)
	pilot.LapTime = "01:00.200"
	pilot.LapSpeedAverage = "242,0"
	racePilotStatistics = append(racePilotStatistics, pilot)

	pilotResultBuilder := domain.NewRaceResultBuilder(
		pilot_statistics.NewRaceSpeedAverage(),
		pilot_statistics.NewPilotBestLap(),
		pilot_statistics.NewRaceDuration(),
	)
	pilotStatistic := pilotResultBuilder.Build(racePilotStatistics)

	assert.Equal(t, "01:00.200", pilotStatistic.BestLap)
	assert.Equal(t, "04:06", pilotStatistic.RaceDuration)
	assert.Equal(t, 240.5, pilotStatistic.SpeedRaceAverage)
}
