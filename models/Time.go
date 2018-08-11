package models

import (
	"time"
)

/* Structure for current time */
type CurrentTime struct {
	Year  int
	Month time.Month
	Day   int
	Hour  int
	Min   int
	Sec   int
}

/* Init */
func NewTime(Year int, Month time.Month, Day, Hour, Min, Sec int) *CurrentTime {
	return &CurrentTime{Year, Month, Day, Hour, Min, Sec}
}

/* Function return current time */
func GetCurrentTime() CurrentTime {
	/* Get current time */
	cTime := time.Now()

	/* Init Current time */
	var currentTime CurrentTime
	currentTime.Year, currentTime.Month, currentTime.Day = cTime.Date()
	currentTime.Hour, currentTime.Min, currentTime.Sec = cTime.Clock()

	return currentTime
}
