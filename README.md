# schedule
Team project for Pre-Piscine program

Schedule Table is a Go program that displays a schedule as a colored ASCII table in the terminal.
This program gives you visualisation of your schedule in terminal

For running this code - use this command "go run ./cmd/schedule/main.go"

If you wanna alter your schedule, please proceed to directory called cmd and then schedule. Then, go to main.go file and adjust arguments in the "beknazar.PrintSchedule()" function:

1) First argument: must be a slice with single size of every row (0 - free, no color; 1 - busy, red; 2 - important, yellow)

2) Second argument: with which symbols do we need to label busy cells

3) Third argument: with which symbols do we need to label important cells

Example:

Input
printSchedule(
    []string{"0102", "1100"},
    "###",
    "!!!",
)
Output

At runtime, ### is displayed in red and !!! is displayed in yellow.

+---------+---------+---------+---------+
|         |   ###   |         |   !!!   |
+---------+---------+---------+---------+
|   ###   |   ###   |         |         |
+---------+---------+---------+---------+


Team
Anuar - sharp-a
Baqytgul - bolathanbakythan-cmyk
Beknazar - dynamicmines
