package pilot_statistics_test

import (
	"github.com/henriqueholanda/race-result/domain/models"
	"github.com/henriqueholanda/race-result/domain/pilot_statistics"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestShouldReturnPilotBestLap(t *testing.T) {
	raceStatistics := []models.RacePilotStatistics{}

	pilot := models.RacePilotStatistics{Number: "001", Name: "Senna", LapSpeedAverage: "44.444"}
	pilot.LapTime = "01:01.333"
	pilot.LapSpeedAverage = "235,0"
	raceStatistics = append(raceStatistics, pilot)
	pilot.LapTime = "01:01.000"
	pilot.LapSpeedAverage = "240,0"
	raceStatistics = append(raceStatistics, pilot)
	pilot.LapTime = "01:03.900"
	pilot.LapSpeedAverage = "245,0"
	raceStatistics = append(raceStatistics, pilot)
	pilot.LapTime = "01:00.200"
	pilot.LapSpeedAverage = "242,0"
	raceStatistics = append(raceStatistics, pilot)

	pilotBestLap := pilot_statistics.NewPilotBestLap()

	pilotStatistics := models.RacePilotStatistics{
		LapAmount: len(raceStatistics),
	}

	assert.Equal(t, "01:00.200", pilotBestLap.Generate(raceStatistics, pilotStatistics).BestLap)
}
