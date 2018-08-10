package models

import (
	"time"
)

type CurrentTime struct {
	Year  int
	Month time.Month
	Day   int
	Hour  int
	Min   int
	Sec   int
}

func NewTime(Year int, Month time.Month, Day, Hour, Min, Sec int) *CurrentTime {
	return &CurrentTime{Year, Month, Day, Hour, Min, Sec}
}

func GetCurrentTime() CurrentTime {
	cTime := time.Now()

	var currentTime CurrentTime

	currentTime.Year, currentTime.Month, currentTime.Day = cTime.Date()
	currentTime.Hour, currentTime.Min, currentTime.Sec = cTime.Clock()

	return currentTime
}
