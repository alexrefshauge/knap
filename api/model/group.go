package model

import "time"

func GroupByWeekday(presses []time.Time) [][]time.Time {
	groups := make([][]time.Time, 7)
	pressIndex := 0
	for i := 1; i <= 7; i++ {
		groupIndex := i - 1
		groups[groupIndex] = make([]time.Time, 0)
		for ; pressIndex < len(presses) && int(presses[pressIndex].Weekday()) == i%7; pressIndex++ {

			groups[groupIndex] = append(groups[groupIndex], presses[pressIndex])
		}
	}
	return groups
}
