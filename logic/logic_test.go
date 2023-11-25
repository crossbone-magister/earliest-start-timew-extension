package logic_test

import (
	"testing"

	"github.com/crossbone-magister/earliest-start/logic"
	"github.com/crossbone-magister/timewlib"
)

func TestEarliestStart(t *testing.T) {
	var earliest_expected = timewlib.NewInterval(9, 0, 10, 0)
	var intervals = []timewlib.Interval{
		*timewlib.NewInterval(10, 00, 11, 00),
		*earliest_expected,
		*timewlib.NewInterval(11, 0, 12, 0),
	}
	var earliest_actual = logic.FindEarliestStart(intervals)
	if earliest_actual.StartHour() != earliest_expected.StartHour() && earliest_actual.StartMinute() != earliest_expected.StartHour() {
		t.Errorf("Expected earliest start to be %s, found instead %s", earliest_expected.String(), earliest_actual.String())
	}
}

func TestLatestEnd(t *testing.T) {
	var latest_expected = timewlib.NewInterval(11, 0, 12, 0)
	var intervals = []timewlib.Interval{
		*timewlib.NewInterval(10, 00, 11, 00),
		*latest_expected,
		*timewlib.NewInterval(9, 0, 10, 0),
	}
	var latest_actual = logic.FindLatestEnd(intervals)
	if latest_actual.EndHour() != latest_expected.EndHour() && latest_actual.EndMinute() != latest_expected.EndHour() {
		t.Errorf("Expected latest end to be %s, found instead %s", latest_expected.String(), latest_actual.String())
	}
}
