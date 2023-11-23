package logic

import "github.com/crossbone-magister/timewlib"

func FindEarliestStart(intervals []timewlib.Interval) timewlib.Interval {
	var earliest_start = intervals[0]
	for _, interval := range intervals {
		if interval.StartHour() < earliest_start.StartHour() || (earliest_start.StartHour() == interval.StartHour() && interval.StartMinute() < earliest_start.StartMinute()) {
			earliest_start = interval
		}
	}
	return earliest_start
}

func FindLatestStop(intervals []timewlib.Interval) timewlib.Interval {
	var latest_stop = intervals[0]
	for _, interval := range intervals {
		if interval.EndHour() > latest_stop.EndHour() || (latest_stop.EndHour() == interval.EndHour() && interval.EndMinute() > latest_stop.EndMinute()) {
			latest_stop = interval
		}
	}
	return latest_stop
}
