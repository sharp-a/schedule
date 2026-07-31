package main
import "strings"
const (
    red = "\033[31m"
    yellow = "\033[33m"
    reset = "\033[0m"
)
func formatCell(cell byte, busy string, important string) string {
	text := ""
	switch cell {
    case '0':
	    text = ""
    case '1':
	    text = busy
    case '2':
	    text = important
	}

	left := (9 - len(text)) / 2
    right := 9 - len(text) - left

	cellText := strings.Repeat(" ", left) +
	    text +
	    strings.Repeat(" ", right)
	switch cell {
    case '1':
        return red + cellText + reset
    case '2':
	    return yellow + cellText + reset
    default:
	return cellText
}
}






