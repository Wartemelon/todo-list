package service

import (
	"time"

	"github.com/Wartemelon/TODO-list/pkg/db"
)

func CheckDate(task *db.Task) error {
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(dateFormat)
	}

	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return err
	}

	nowDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	taskDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if afterNow(nowDate, taskDate) {
		if len(task.Repeat) == 0 {
			task.Date = now.Format(dateFormat)
		} else {
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		}
	}

	return nil
}
