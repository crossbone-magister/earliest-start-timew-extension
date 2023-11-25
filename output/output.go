package output

import (
	"fmt"

	"github.com/crossbone-magister/timewlib"
)

func FormatEarliestStart(earliest_start timewlib.Interval) string {
	year, month, day := earliest_start.StartDate()
	return fmt.Sprintf("Earliest Start: %02d:%02d on %d-%02d-%02d", earliest_start.StartHour(), earliest_start.StartMinute(), year, month, day)
}

func FormatLatestEnd(latest_end timewlib.Interval) string {
	year, month, day := latest_end.EndDate()
	return fmt.Sprintf("Latest End: %02d:%02d on %d-%02d-%02d", latest_end.EndHour(), latest_end.EndMinute(), year, month, day)
}
