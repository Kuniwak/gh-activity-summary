package daterange

import "time"

func NewDateRange(since, until time.Time) []time.Time {
	dates := []time.Time{}

	since = time.Date(since.Year(), since.Month(), 1, 0, 0, 0, 0, time.Local)
	until = time.Date(until.Year(), until.Month(), 1, 0, 0, 0, 0, time.Local)

	for d := since; d.Before(until); d = d.AddDate(0, 1, 0) {
		dates = append(dates, d)
	}

	return dates
}
