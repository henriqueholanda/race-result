package domain_test

import (
	"github.com/henriqueholanda/race-result/domain"
	"github.com/henriqueholanda/race-result/domain/models"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestShouldReturnSortedClassificationWhenAllPilotsFinishedTheRace(t *testing.T) {
	groupedResult := map[string]models.RaceStatistics{}

	groupedResult["001"] = append(groupedResult["001"], models.RacePilotStatistics{Number: "001", LapTime: "1:02.000"})
	groupedResult["001"] = append(groupedResult["001"], models.RacePilotStatistics{Number: "001", LapTime: "1:01.700"})
	groupedResult["001"] = append(groupedResult["001"], models.RacePilotStatistics{Number: "001", LapTime: "1:06.000"})

	groupedResult["002"] = append(groupedResult["002"], models.RacePilotStatistics{Number: "002", LapTime: "1:02.010"})
	groupedResult["002"] = append(groupedResult["002"], models.RacePilotStatistics{Number: "002", LapTime: "1:01.040"})
	groupedResult["002"] = append(groupedResult["002"], models.RacePilotStatistics{Number: "002", LapTime: "1:05.200"})

	groupedResult["003"] = append(groupedResult["003"], models.RacePilotStatistics{Number: "003", LapTime: "1:22.010"})
	groupedResult["003"] = append(groupedResult["003"], models.RacePilotStatistics{Number: "003", LapTime: "1:21.040"})
	groupedResult["003"] = append(groupedResult["003"], models.RacePilotStatistics{Number: "003", LapTime: "1:25.200"})

	racerClassifier := domain.NewRaceClassifier()
	classificationRace := racerClassifier.Classify(groupedResult)

	assert.Equal(t, []string{"002", "001", "003"}, classificationRace)
}

func TestShouldReturnSortedClassificationWhenOnePilotNotFinishedTheRace(t *testing.T) {
	groupedResult := map[string]models.RaceStatistics{}

	groupedResult["001"] = append(groupedResult["001"], models.RacePilotStatistics{Number: "001", LapTime: "1:02.000"})
	groupedResult["001"] = append(groupedResult["001"], models.RacePilotStatistics{Number: "001", LapTime: "1:01.700"})
	groupedResult["001"] = append(groupedResult["001"], models.RacePilotStatistics{Number: "001", LapTime: "1:06.000"})

	groupedResult["002"] = append(groupedResult["002"], models.RacePilotStatistics{Number: "002", LapTime: "1:02.010"})
	groupedResult["002"] = append(groupedResult["002"], models.RacePilotStatistics{Number: "002", LapTime: "1:01.040"})
	groupedResult["002"] = append(groupedResult["002"], models.RacePilotStatistics{Number: "002", LapTime: "1:05.200"})

	groupedResult["003"] = append(groupedResult["003"], models.RacePilotStatistics{Number: "003", LapTime: "1:22.010"})
	groupedResult["003"] = append(groupedResult["003"], models.RacePilotStatistics{Number: "003", LapTime: "1:21.040"})

	racerClassifier := domain.NewRaceClassifier()
	classificationRace := racerClassifier.Classify(groupedResult)

	assert.Equal(t, []string{"002", "001", "003"}, classificationRace)
}

func TestShouldReturnSortedClassificationWhenOnlyOnePilotFinishedTheRace(t *testing.T) {
	groupedResult := map[string]models.RaceStatistics{}

	groupedResult["001"] = append(groupedResult["001"], models.RacePilotStatistics{Number: "001", LapTime: "1:02.000"})
	groupedResult["001"] = append(groupedResult["001"], models.RacePilotStatistics{Number: "001", LapTime: "1:01.700"})
	groupedResult["001"] = append(groupedResult["001"], models.RacePilotStatistics{Number: "001", LapTime: "1:06.000"})

	groupedResult["002"] = append(groupedResult["002"], models.RacePilotStatistics{Number: "002", LapTime: "1:02.010"})
	groupedResult["002"] = append(groupedResult["002"], models.RacePilotStatistics{Number: "002", LapTime: "1:01.040"})

	groupedResult["003"] = append(groupedResult["003"], models.RacePilotStatistics{Number: "003", LapTime: "1:22.010"})
	groupedResult["003"] = append(groupedResult["003"], models.RacePilotStatistics{Number: "003", LapTime: "1:21.040"})

	racerClassifier := domain.NewRaceClassifier()
	classificationRace := racerClassifier.Classify(groupedResult)

	assert.Equal(t, []string{"001", "002", "003"}, classificationRace)
}
