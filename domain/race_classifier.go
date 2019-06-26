package domain

import (
	"github.com/henriqueholanda/race-result/domain/models"
	"github.com/henriqueholanda/race-result/infrastructure/helper"
	"sort"
	"strconv"
)

type RaceClassifier struct {
	NumberOfLaps int
	punishment   map[string]int
}

func NewRaceClassifier() *RaceClassifier {
	return &RaceClassifier{
		4,
		map[string]int{
			"1": 1000000000,
			"2": 100000000,
			"3": 10000000,
			"4": 0,
		},
	}
}

func (rc *RaceClassifier) Classify(resultByPilot map[string]models.RaceStatistics) []string {
	classification := []map[int]map[string]string{}

	for pilotNumber, data := range resultByPilot {
		classification = append(classification, rc.generateClassification(pilotNumber, data))
	}

	return rc.orderClassification(classification)
}

func (rc *RaceClassifier) generateClassification(pilotNumber string, data models.RaceStatistics) map[int]map[string]string {
	return map[int]map[string]string{
		rc.getPilotRaceTime(data): {
			"pilot_number": pilotNumber, "laps": strconv.Itoa(len(data)),
		},
	}
}

func (rc *RaceClassifier) getPilotRaceTime(data models.RaceStatistics) int {
	var totalTimeInt int

	for _, lap := range data {
		totalTimeInt += helper.ParseTimeToIntValue(lap.LapTime)
	}

	return totalTimeInt
}

func (rc *RaceClassifier) orderClassification(classification []map[int]map[string]string) []string {
	var orderedLapsTime []int
	var finalClassification []string
	timeByPilot := make(map[int]string)

	for _, data := range classification {
		for timeRace, pilotData := range data {
			timeRaceWithPunish := timeRace + rc.punishment[pilotData["laps"]]
			orderedLapsTime = append(orderedLapsTime, timeRaceWithPunish)
			timeByPilot[timeRaceWithPunish] = pilotData["pilot_number"]
		}
	}

	sort.Slice(orderedLapsTime, func(i, j int) bool { return orderedLapsTime[i] < orderedLapsTime[j] })

	for _, raceTime := range orderedLapsTime {
		finalClassification = append(finalClassification, timeByPilot[raceTime])
	}

	return finalClassification
}
