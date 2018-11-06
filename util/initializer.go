package util

import (
	"strconv"
	"time"
)

const (
	_START_YEAR     = 2008
	_END_YEAR       = 2018
	_JANUARY_INDEX  = 1
	_DECEMBER_INDEX = 12
)

func InitializeMonths() (months map[string]int64) {
	months = make(map[string]int64)
	monthsTimestamps := GetMonthsTimestamps()
	index := 0
	for year := _START_YEAR; year <= _END_YEAR; year++ {
		for month := _JANUARY_INDEX; month <= _DECEMBER_INDEX; month++ {
			key := time.Month(month).String() + " " + strconv.Itoa(year)
			months[key] = monthsTimestamps[index]
			index++
		}
	}
	return months
}

func InitializeYears() (years map[string]int64) {
	years = make(map[string]int64)
	yearsTimestamps := GetYearsTimestamp()
	index := 0
	for year := _START_YEAR; year <= _END_YEAR; year++ {
		key := strconv.Itoa(year)
		years[key] = yearsTimestamps[index]
		index++
	}
	return years
}

func GetYearsTimestamp() (timestamps []int64) {
	for year := _START_YEAR; year <= _END_YEAR; year++ {
		unixTime := time.Date(year, _JANUARY_INDEX, 1, 0, 0, 0, 0, time.UTC).Unix()
		timestamps = append(timestamps, unixTime)
	}
	return timestamps
}

// Return all months in timestamp style.
func GetMonthsTimestamps() (timestamps []int64) {
	for year := _START_YEAR; year <= _END_YEAR; year++ {
		for month := _JANUARY_INDEX; month <= _DECEMBER_INDEX; month++ {
			unixTime := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC).Unix()
			timestamps = append(timestamps, unixTime)
		}
	}
	return timestamps
}

// First year - 2008, Last year - 2018
func GetYearsNames() []string {
	return []string{
		"2008", "2009", "2010",
		"2011", "2012", "2013",
		"2014", "2015", "2016",
		"2017", "2018",
	}
}

// Return all month names
func GetMonthsNames() []string {
	return []string{
		"January", "February", "March",
		"April", "May", "June",
		"Jule", "August", "September",
		"October", "November",
	}
}
