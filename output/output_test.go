package output

import (
	"fmt"
	"testing"
	"time"

	"github.com/crossbone-magister/timewlib"
)

func TestFormatEarliestStart(t *testing.T) {
	t.Run("Hour and minute < 10", func(t *testing.T) {
		interval := timewlib.NewInterval(8, 3, 9, 4)
		year, month, day := time.Now().Date()
		expected := fmt.Sprintf("Earliest Start: 08:03 on %d-%02d-%02d", year, month, day)
		actual := FormatEarliestStart(*interval)
		if expected != actual {
			t.Errorf("TestFormatEarliestStart failed: [%s] != [%s]", expected, actual)
		}
	})
	t.Run("Hour and minute >= 10", func(t *testing.T) {
		interval := timewlib.NewInterval(10, 11, 12, 13)
		year, month, day := time.Now().Date()
		expected := fmt.Sprintf("Earliest Start: 10:11 on %d-%02d-%02d", year, month, day)
		actual := FormatEarliestStart(*interval)
		if expected != actual {
			t.Errorf("TestFormatEarliestStart failed: [%s] != [%s]", expected, actual)
		}
	})
}

func TestFormatLatestEnd(t *testing.T) {
	t.Run("Hour and minute < 10", func(t *testing.T) {
		interval := timewlib.NewInterval(8, 3, 9, 4)
		year, month, day := time.Now().Date()
		expected := fmt.Sprintf("Latest End: 09:04 on %d-%02d-%02d", year, month, day)
		actual := FormatLatestEnd(*interval)
		if expected != actual {
			t.Errorf("TestFormatLatestEnd failed: [%s] != [%s]", expected, actual)
		}
	})
	t.Run("Hour and minute >= 10", func(t *testing.T) {
		interval := timewlib.NewInterval(10, 11, 12, 13)
		year, month, day := time.Now().Date()
		expected := fmt.Sprintf("Latest End: 12:13 on %d-%02d-%02d", year, month, day)
		actual := FormatLatestEnd(*interval)
		if expected != actual {
			t.Errorf("TestFormatLatestEnd failed: [%s] != [%s]", expected, actual)
		}
	})
}
