package exporter

import (
	"encoding/csv"
	"fmt"
	"github.com/henriqueholanda/race-result/domain/models"
	"github.com/henriqueholanda/race-result/infrastructure/helper"
	"log"
	"os"
	"strconv"
)

type RaceResultExporter struct {
}

func NewRaceResultExporter() *RaceResultExporter {
	return &RaceResultExporter{}
}

func (r *RaceResultExporter) Export(raceResult models.RaceResult) {
	fileCsv, _ := os.OpenFile("race_result.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer fileCsv.Close()

	position := 1

	csvWriter := csv.NewWriter(fileCsv)
	strWrite := []string{"POSITION", "NUMBER", "PILOT", "LAPS", "TIME", "VELOCITY AVERAGE", "BEST LAP", "DIFFERENCE FOR WINNER"}
	csvWriter.Write(strWrite)

	differenceTimeToWinner := "00:00"
	timeRaceWinner := ""

	pilotStatistics := raceResult.PilotStatistics
	for _, pilotNumber := range raceResult.Classification {

		if position == 1 {
			timeRaceWinner = pilotStatistics[pilotNumber].RaceDuration
		}

		if position > 1 {
			differenceTimeToWinner = helper.CalculateDifferenceDates(
				timeRaceWinner,
				pilotStatistics[pilotNumber].RaceDuration,
			)
		}

		value := []string{
			strconv.Itoa(position),
			pilotStatistics[pilotNumber].Number,
			pilotStatistics[pilotNumber].Name,
			strconv.Itoa(pilotStatistics[pilotNumber].LapAmount),
			pilotStatistics[pilotNumber].RaceDuration,
			fmt.Sprintf("%f", pilotStatistics[pilotNumber].SpeedRaceAverage),
			pilotStatistics[pilotNumber].BestLap,
			differenceTimeToWinner,
		}

		csvWriter.Write(value)
		csvWriter.Flush()
		position++
	}

	csvWriter.Write([]string{"BEST LAP ON RACE"})
	csvWriter.Flush()
	csvWriter.Write([]string{"PILOT", "NUMBER", "TIME"})
	csvWriter.Flush()

	bestLap := []string{
		raceResult.PilotNameRaceBestLap,
		raceResult.PilotNumberRaceBestLap,
		raceResult.RaceBestLapTime,
	}
	csvWriter.Write(bestLap)
	csvWriter.Flush()

	log.Print("CSV created")
}
