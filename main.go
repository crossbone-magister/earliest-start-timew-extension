package main

import (
	"fmt"
	"os"

	"github.com/crossbone-magister/earliest-start/logic"
	"github.com/crossbone-magister/timewlib"
)

func main() {
	parsed, err := timewlib.Parse(os.Stdin)
	if err == nil {
		intervals, err := timewlib.Process(parsed.Intervals)
		if err == nil {
			var earliest_start = logic.FindEarliestStart(intervals)
			var latest_stop = logic.FindLatestStop(intervals)
			//TODO Improvement on timewlib is required to get start date and end date
			println("Earliest Start")
			println("--------------")
			println(fmt.Sprintf("%d:%d on %s", earliest_start.StartHour(), earliest_start.StartMinute(), earliest_start.String()))
			println("Latest Stop")
			println("-----------")
			println(fmt.Sprintf("%d:%d on %s", latest_stop.StartHour(), latest_stop.StartMinute(), latest_stop.String()))
		} else {
			printErrorAndExit(err)
		}
	} else {
		printErrorAndExit(err)
	}

}

func printErrorAndExit(err error) {
	println("Error while parsing: ", err)
	os.Exit(1)
}
