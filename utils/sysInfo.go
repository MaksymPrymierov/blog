package utils

import (
	"fmt"
	"os"
)

func GetUptimeServer() (int, int, int) {
	var day, hour, minute, data int

	file, _ := os.Open("/proc/uptime")
	defer file.Close()

	fmt.Fscanf(file, "%d", &data)

	day = ((data / 60) / 60) / 24
	hour = ((data / 60) / 60) - (day * 60)
	minute = (data / 60) - (hour * 60)

	return day, hour, minute
}
