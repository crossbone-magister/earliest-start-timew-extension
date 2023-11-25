package main

import (
	"fmt"
	"os"

	"github.com/crossbone-magister/earliest-start/logic"
	"github.com/crossbone-magister/earliest-start/output"
	"github.com/crossbone-magister/timewlib"
)

func main() {
	parsed, err := timewlib.Parse(os.Stdin)
	if err == nil {
		intervals, err := timewlib.Process(parsed.Intervals)
		if err == nil {
			var earliest_start = logic.FindEarliestStart(intervals)
			var latest_end = logic.FindLatestEnd(intervals)
			println(output.FormatEarliestStart(earliest_start))
			println(output.FormatLatestEnd(latest_end))
		} else {
			printErrorAndExit(err)
		}
	} else {
		printErrorAndExit(err)
	}

}

func printErrorAndExit(err error) {
	fmt.Printf("Error while parsing: %s\n", err)
	os.Exit(1)
}
