package helper_test

import (
	"github.com/henriqueholanda/race-result/infrastructure/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnFormattedDataPassing128Seconds(t *testing.T) {
	expected := "02:28"
	formatted := helper.FormatBySeconds(148)
	assert.Equal(t, expected, formatted)
}

func TestShouldReturnFormattedDataPassing50Seconds(t *testing.T) {
	expected := "00:40"
	formatted := helper.FormatBySeconds(40)
	assert.Equal(t, expected, formatted)
}

func TestShouldReturnFormattedDataPassing130Seconds(t *testing.T) {
	expected := "02:30"
	formatted := helper.FormatBySeconds(150)
	assert.Equal(t, expected, formatted)
}

func TestShouldReturnFormattedDataPassing60Seconds(t *testing.T) {
	expected := "01:00"
	formatted := helper.FormatBySeconds(60)
	assert.Equal(t, expected, formatted)
}

func TestShouldReturnFormattedDateDiffInSeconds(t *testing.T) {
	expected := "00:40"
	formatted := helper.CalculateDifferenceDates("02:00", "02:40")
	assert.Equal(t, expected, formatted)
}

func TestShouldReturnFormattedDateDiffOneMinute(t *testing.T) {
	expected := "01:00"
	formatted := helper.CalculateDifferenceDates("03:00", "04:00")
	assert.Equal(t, expected, formatted)
}

func TestShouldReturnIntValueOfTime(t *testing.T) {
	expected := 100
	formatted := helper.ParseTimeToIntValue("01:00")
	assert.Equal(t, expected, formatted)
}

func TestShouldReturnFormattedLapSpeed(t *testing.T) {
	expected := 45.5
	formatted := helper.FormatLapSpeed("45,5")
	assert.Equal(t, expected, formatted)
}

func TestShouldReturnSplittedLapTime(t *testing.T) {
	expectedMin := 2
	expectedSec := 0
	expectedMil := 34

	minutes, seconds, miliseconds := helper.SplitLapTime("02:00:34")
	assert.Equal(t, expectedMin, minutes)
	assert.Equal(t, expectedSec, seconds)
	assert.Equal(t, expectedMil, miliseconds)
}
