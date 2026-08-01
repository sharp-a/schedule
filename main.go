package main

import "schedule/internal/beknazar"

func main() {
	beknazar.PrintSchedule(
		[]string{"0102", "1100", "0010", "2001"},
		"###",
		"!!!",
	)
}