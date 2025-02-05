package pomodoro

import (
	"fmt"
	"time"
)

// DailySummary function returns daily pomodoro and breaks durations
func DailySummary(day time.Time, config *IntervalConfig) ([]time.Duration, error) {
	dPomo, err := config.repo.CategorySummary(day, CategoryPomodoro)
	if err != nil {
		return nil, err
	}

	dBreaks, err := config.repo.CategorySummary(day, "%Breaks")
	if err != nil {
		return nil, err
	}

	return []time.Duration{dPomo, dBreaks}, nil
}

type LineSeries struct {
	Name   string
	Labels map[int]string
	Values []float64
}

// RangeSummary function returns pomodoro and breaks series of n days range summary
func RangeSummary(start time.Time, n int, config *IntervalConfig) ([]LineSeries, error) {
	pomoSeries := LineSeries{
		Name:   "Pomodoro",
		Labels: make(map[int]string),
		Values: make([]float64, n),
	}

	breaksSeries := LineSeries{
		Name:   "Break",
		Labels: make(map[int]string),
		Values: make([]float64, n),
	}

	for i := 0; i < n; i++ {
		day := start.AddDate(0, 0, -i)
		ds, err := DailySummary(day, config)
		if err != nil {
			return nil, err
		}

		label := fmt.Sprintf("%02d/%s", day.Day(), day.Format("Jan"))

		pomoSeries.Labels[i] = label
		pomoSeries.Values[i] = ds[0].Seconds()

		breaksSeries.Labels[i] = label
		breaksSeries.Values[i] = ds[1].Seconds()
	}

	return []LineSeries{
		pomoSeries,
		breaksSeries,
	}, nil
}
