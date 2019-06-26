package application

import (
	"github.com/henriqueholanda/race-result/domain"
	"log"
)

type Application struct {
	raceResultHandler *domain.RaceResultHandler
}

func NewApplication(rrh *domain.RaceResultHandler) *Application {
	return &Application{rrh}
}

func (a *Application) Start() {
	log.Print("Starting application")
	a.raceResultHandler.GenerateResult()
	log.Print("Finishing application")
}
