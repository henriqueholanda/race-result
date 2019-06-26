package domain

import (
	"github.com/henriqueholanda/race-result/domain/models"
	"github.com/henriqueholanda/race-result/infrastructure/exporter"
	"github.com/henriqueholanda/race-result/infrastructure/helper"
	"github.com/henriqueholanda/race-result/infrastructure/repository"
	"log"
	"sort"
)

type RaceResultHandler struct {
	repository        *repository.RaceResultRepository
	raceClassifier    *RaceClassifier
	raceResultBuilder *RaceResultBuilder
	exporter          *exporter.RaceResultExporter
}

func NewRaceResultHandler(repository *repository.RaceResultRepository, raceClassifier *RaceClassifier, raceResultBuilder *RaceResultBuilder, exporter *exporter.RaceResultExporter) *RaceResultHandler {
	return &RaceResultHandler{
		repository:        repository,
		raceClassifier:    raceClassifier,
		raceResultBuilder: raceResultBuilder,
		exporter:          exporter,
	}
}

func (rh *RaceResultHandler) GenerateResult() {
	log.Print("Generating race result")
	var raceResult = models.RaceResult{}

	raceResultGrouped := rh.groupRaceResultByPilot(
		rh.repository.GetList(),
	)

	pilotStatisticsChan := make(chan map[string]models.RacePilotStatistics)
	raceClassificationChan := make(chan []string)

	go rh.getPilotStatistics(raceResultGrouped, pilotStatisticsChan)
	go rh.getClassification(raceResultGrouped, raceClassificationChan)

	raceResult.PilotStatistics = <-pilotStatisticsChan
	raceResult.Classification = <-raceClassificationChan

	raceResult = rh.getRaceStatistics(raceResult)

	rh.exporter.Export(raceResult)
}

func (rh *RaceResultHandler) groupRaceResultByPilot(raceResult models.RaceStatistics) map[string]models.RaceStatistics {
	groupedResult := map[string]models.RaceStatistics{}

	for _, data := range raceResult {
		groupedResult[data.Number] = append(groupedResult[data.Number], data)
	}

	return groupedResult
}

func (rh *RaceResultHandler) getRaceStatistics(raceResult models.RaceResult) models.RaceResult {
	pilot := make(map[int]map[string]string, 0)
	sortLap := []int{}

	for _, data := range raceResult.PilotStatistics {
		timeInt := helper.ParseTimeToIntValue(data.BestLap)
		pilot[timeInt] = map[string]string{
			"pilot_name":   data.Name,
			"pilot_number": data.Number,
			"lap_time":     data.BestLap,
		}
		sortLap = append(sortLap, timeInt)
	}

	sort.Ints(sortLap)
	bestLap := sortLap[0]

	raceResult.PilotNameRaceBestLap = pilot[bestLap]["pilot_name"]
	raceResult.PilotNumberRaceBestLap = pilot[bestLap]["pilot_number"]
	raceResult.RaceBestLapTime = pilot[bestLap]["lap_time"]

	return raceResult
}

func (rh *RaceResultHandler) getPilotStatistics(raceResultGrouped map[string]models.RaceStatistics, pilotStatisticsChan chan map[string]models.RacePilotStatistics) {
	pilotStatistics := make(map[string]models.RacePilotStatistics)
	for pilotNumber, laps := range raceResultGrouped {
		pilotStatistics[pilotNumber] = rh.raceResultBuilder.Build(laps)
	}

	pilotStatisticsChan <- pilotStatistics
}

func (rh *RaceResultHandler) getClassification(raceResultGrouped map[string]models.RaceStatistics, raceClassificationChan chan []string) {
	raceClassificationChan <- rh.raceClassifier.Classify(raceResultGrouped)
}
