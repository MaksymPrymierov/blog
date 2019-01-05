package utils

import (
	"fmt"
	"os"
)

func GetSecondUptime() int {
	var second int

	file, _ := os.Open("/proc/uptime")
	defer file.Close()

	fmt.Fscanf(file, "%d", &second)

	return second
}

func GetUptimeServer() (int, int, int) {
	var day, hour, minute, second int
	second = GetSecondUptime()

	day = ((second / 60) / 60) / 24
	hour = ((second / 60) / 60)
	minute = (second / 60) - (hour * 60)
	hour -= day * 24

	return day, hour, minute
}
