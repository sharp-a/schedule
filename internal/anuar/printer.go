package anuar

import "strconv"

func validateSchedule(rows []string) string {

	switch {
	case len(rows) == 0 || rows[0] == "":
		return "error: empty schedule"
	}

	for index, i := range rows {
		if index == len(rows)-1 {

		} else {
			if len(i) == len(rows[index+1]) {
				continue
			} else {
				return "error: rows have different lengths"
			}
		}
	}

	for _, i := range rows {
		num, error := strconv.Atoi(i)

		if error != nil || num < 0 {
			return "error: invalid symbol (only 0/1/2 allowed)"
		}

		for num > 0 {
			if num % 10 == 0 || num % 10 == 1 || num % 10 == 2 {
				num /= 10
			} else {
				return "error: invalid symbol (only 0/1/2 allowed)"
			}
		}
	}

	return ""


}



// error: empty schedule 
// error: rows have different lengths
// error: invalid symbol (only 0/1/2 allowed)