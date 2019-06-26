package main

import (
	"github.com/henriqueholanda/race-result/application"
	"github.com/henriqueholanda/race-result/domain"
	"github.com/henriqueholanda/race-result/domain/pilot_statistics"
	"github.com/henriqueholanda/race-result/infrastructure/repository"
)

func main() {
	raceClassifierBuilder := domain.NewRaceResultBuilder(
		pilot_statistics.NewRaceSpeedAverage(),
		pilot_statistics.NewPilotBestLap(),
		pilot_statistics.NewRaceDuration(),
	)

	raceResultHandler := domain.NewRaceResultHandler(
		repository.NewRaceResult(),
		domain.NewRaceClassifier(),
		raceClassifierBuilder,
	)

	application.NewApplication(raceResultHandler).Start()
}
