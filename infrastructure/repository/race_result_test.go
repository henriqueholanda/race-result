package repository_test

import (
	"github.com/henriqueholanda/race-result/domain/models"
	"github.com/henriqueholanda/race-result/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnRaceStatistics(t *testing.T) {
	raceResultRepository := repository.NewRaceResult()
	raceStatistics := raceResultRepository.GetList()

	assert.NotZero(t, len(raceStatistics))
	assert.IsType(t, models.RaceStatistics{}, raceStatistics)
}
