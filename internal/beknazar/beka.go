package beknazar

import "fmt"
func PrintSchedule (string, busy string, important string) {
	columns := len(rows[0])
    border := "+" + strings.Repeat("---------+", columns)

    fmt.Println(border)

    for _, row := range rows {
        fmt.Print("|")

        for i := 0; i < len(row); i++ {
            fmt.Print(formatCell(row[i], busy, important))
            fmt.Print("|")
        }
}
}