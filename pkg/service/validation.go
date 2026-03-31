package service

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidW(repeat []string) (bool, error) {
	if len(repeat) < 2 {
		return false, fmt.Errorf("parameter 'w' requires days")
	}

	vals := strings.Split(repeat[1], ",")
	if len(vals) < 1 || len(vals) > 7 {
		return false, fmt.Errorf("len of sequence after 'w' must be between 1 and 7 (including borders)")
	}
	for _, v := range vals {
		n, err := strconv.Atoi(v)
		if err != nil {
			return false, fmt.Errorf("after 'w' must be numbers")
		}
		if n < 1 || n > 7 {
			return false, fmt.Errorf("numbers after 'w' must be between 1 and 7 (including borders)")
		}
	}

	return true, nil
}

func isValidM(repeat []string) (bool, error) {
	if len(repeat) < 2 {
		return false, fmt.Errorf("parameter 'm' requires days")
	}

	days := strings.Split(repeat[1], ",")
	if len(days) < 1 || len(days) > 31 {
		return false, fmt.Errorf("len of days sequence after 'm' must be between 1 and 31 (including borders)")
	}

	for _, v := range days {
		n, err := strconv.Atoi(v)
		if err != nil {
			return false, fmt.Errorf("after 'm' must be numbers")
		}
		if n < -2 || n > 31 || n == 0 {
			return false, fmt.Errorf("day numbers after 'm' must be between 1 and 31 (including borders)")
		}
	}

	if len(repeat) > 2 {
		months := strings.Split(repeat[2], ",")
		if len(months) > 12 {
			return false, fmt.Errorf("len of months sequense after 'm' must be less or equal than 12")
		}

		for _, v := range months {
			n, err := strconv.Atoi(v)
			if err != nil {
				return false, fmt.Errorf("after 'm' must be numbers")
			}
			if n < 1 || n > 12 {
				return false, fmt.Errorf("month numbers after 'm' must be between 1 and 12 (including borders)")
			}
		}
	}

	return true, nil
}

func isValid(repeat []string) (bool, error) {
	if !(repeat[0] == "d" || repeat[0] == "y" || repeat[0] == "w" || repeat[0] == "m") {
		return false, fmt.Errorf("first argument of repeat must be 'd'|'y'|'w'|'m'")
	}

	switch repeat[0] {
	case "d":
		if len(repeat) < 2 {
			return false, fmt.Errorf("parameter 'd' requires day")
		}
		n, err := strconv.Atoi(repeat[1])
		if err != nil {
			return false, fmt.Errorf("after 'd' must be number")
		}
		if n < 1 || n > 400 {
			return false, fmt.Errorf("'d': number of days must be between 1 and 400 (including borders)")
		}

	case "y":
		if len(repeat) > 1 {
			return false, fmt.Errorf("must be only one argument - 'y'")
		}

	case "w":
		return isValidW(repeat)

	case "m":
		return isValidM(repeat)

	default:
		return false, fmt.Errorf("unknown type of rule")
	}

	return true, nil
}
