package beknazar

import (
	"fmt"
	"strings"

	"schedule/internal/anuar"
	"schedule/internal/baqytgul"
)

func PrintSchedule(rows []string, busy string, important string) {
	errorMessage := anuar.ValidateSchedule(rows)

	if errorMessage != "" {
		fmt.Println(errorMessage)
		return
	}

	columns := len(rows[0])
	border := "+" + strings.Repeat("---------+", columns)

	fmt.Println(border)

	for _, row := range rows {
		fmt.Print("|")

		for i := 0; i < len(row); i++ {
			cell := baqytgul.FormatCell(row[i], busy, important)

			fmt.Print(cell)
			fmt.Print("|")
		}

		fmt.Println()
		fmt.Println(border)
	}
}