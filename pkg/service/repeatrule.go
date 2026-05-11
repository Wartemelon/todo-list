package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

func nextDateW(now, date time.Time, repeat []string) (string, error) {
	vals := strings.Split(repeat[1], ",")
	weekdays := make([]time.Weekday, len(vals))
	for i, v := range vals {
		num, _ := strconv.Atoi(v)
		weekdays[i] = time.Weekday(num % 7)
	}

	if !afterNow(date, now) {
		date = now
	}

	for {
		date = date.AddDate(0, 0, 1)
		for _, v := range weekdays {
			if date.Weekday() == v {
				return date.Format(dateFormat), nil
			}
		}
	}
}

func nextDateM(now, date time.Time, repeat []string) (string, error) {
	var day [32]bool
	var lastDay, beforeLastDay bool

	days := strings.Split(repeat[1], ",")
	for _, v := range days {
		num, _ := strconv.Atoi(v)
		if num > 0 {
			day[num] = true
		} else if num == -1 {
			lastDay = true
		} else if num == -2 {
			beforeLastDay = true
		}
	}

	var month [13]bool
	haveMonth := false
	if len(repeat) > 2 {
		haveMonth = true
		months := strings.Split(repeat[2], ",")
		for _, v := range months {
			num, _ := strconv.Atoi(v)
			month[num] = true
		}
	}

	if !afterNow(date, now) {
		date = now
	}

	for {
		date = date.AddDate(0, 0, 1)
		if day[date.Day()] ||
			(date.AddDate(0, 0, 1).Day() == 1 && lastDay) ||
			(date.AddDate(0, 0, 2).Day() == 1 && beforeLastDay) {
			if !haveMonth || month[date.Month()] {
				return date.Format(dateFormat), nil
			}
		}
	}

}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("repeat rule can't be empty")
	}

	components := strings.Fields(repeat)
	ok, err := isValid(components)
	if !ok {
		return "", err
	}

	date, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", err
	}

	switch components[0] {
	case "d":
		n, _ := strconv.Atoi(components[1])

		for {
			date = date.AddDate(0, 0, n)
			if afterNow(date, now) {
				break
			}
		}

		return date.Format(dateFormat), nil

	case "y":
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}

		return date.Format(dateFormat), nil

	case "w":
		return nextDateW(now, date, components)

	case "m":
		return nextDateM(now, date, components)

	default:
		return "", nil // Заглушка, потому что в isValid уже проверяется, что правило лежит в заданном наборе букв
	}
}
