package pilot_statistics

import (
	"github.com/henriqueholanda/race-result/domain/models"
	"github.com/henriqueholanda/race-result/infrastructure/helper"
	"sort"
)

type PilotBestLap struct {
}

func NewPilotBestLap() *PilotBestLap {
	return &PilotBestLap{}
}

func (rbl *PilotBestLap) Generate(raceResult models.RaceStatistics, pilotStatistics models.RacePilotStatistics) models.RacePilotStatistics {
	lapsTime := make(map[int]string)

	for _, data := range raceResult {
		pilotStatistics.Name = data.Name
		pilotStatistics.Number = data.Number
		lapsTime[helper.ParseTimeToIntValue(data.LapTime)] = data.LapTime
	}

	pilotStatistics.BestLap = rbl.getPilotBestLap(lapsTime)

	return pilotStatistics
}

func (rbl *PilotBestLap) getPilotBestLap(lapsTime map[int]string) string {
	lapTime := []int{}

	for time := range lapsTime {
		lapTime = append(lapTime, time)
	}
	sort.Ints(lapTime)

	return lapsTime[lapTime[0]]
}
