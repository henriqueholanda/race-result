package repository

import (
	"bufio"
	"github.com/henriqueholanda/race-result/domain/models"
	"os"
	"regexp"
)

type RaceResultRepository struct {
	raceResult models.RaceStatistics
}

func NewRaceResult() *RaceResultRepository {
	return &RaceResultRepository{make(models.RaceStatistics, 0)}
}

func (r *RaceResultRepository) GetList() models.RaceStatistics {
	file, _ := os.Open(os.Getenv("GOPATH") + "/src/github.com/henriqueholanda/race-result/race_log.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		if r.isFileHeader(line) {
			line++
			continue
		}

		r.raceResult = append(r.raceResult, r.parseLine(scanner.Text()))
	}

	return r.raceResult
}

func (r *RaceResultRepository) isFileHeader(line int) bool {
	return line == 0
}

func (r *RaceResultRepository) parseLine(line string) models.RacePilotStatistics {
	pilotStatistic := models.RacePilotStatistics{}
	rp := regexp.MustCompile("[a-zA-Z.0-9:,]+")
	parsedFileLine := rp.FindAllString(line, -1)

	pilotStatistic.Number = parsedFileLine[1]
	pilotStatistic.Name = parsedFileLine[2]
	pilotStatistic.Lap = parsedFileLine[3]
	pilotStatistic.LapTime = parsedFileLine[4]
	pilotStatistic.LapSpeedAverage = parsedFileLine[5]

	return pilotStatistic
}
