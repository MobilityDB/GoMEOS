package gomeos

// The casts between the time types of Go and those of MEOS convert through the
// generated wrappers of the MEOS conversions, and refuse what the other side
// cannot state: a fraction of a microsecond, an interval with months for a
// time.Duration, the clock 24:00:00 for a time.Time. Each example reads its
// values as #ExampleTemporalOut_size reads its string, through the generated
// wrappers of the functions package. The counts are those of the MEOS
// bindings specification, doc/specs/meos-bindings-algebra.md of MobilityDB.

import (
	"fmt"
	"time"

	"github.com/MobilityDB/GoMEOS/functions"
)

func ExampleDatetimeToTimestamptz() {
	ts, err := functions.DatetimeToTimestamptz(time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC))
	fmt.Println(err)
	us, err := functions.TimestamptzToUnixUs(ts)
	fmt.Println(us, err)
	back, err := functions.TimestamptzToDatetime(ts)
	fmt.Println(back.UTC().Format(time.RFC3339Nano), err)

	_, err = functions.DatetimeToTimestamptz(time.Date(2001, 1, 1, 0, 0, 0, 1, time.UTC))
	fmt.Println(err != nil)
	// Output:
	// <nil>
	// 978307200000000 <nil>
	// 2001-01-01T00:00:00Z <nil>
	// true
}

func ExampleDateToDateADT() {
	d, err := functions.DateToDateADT(time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC))
	fmt.Println(err)
	text, err := functions.DateOut(d)
	fmt.Println(text, err)
	days, err := functions.DateToUnixDays(d)
	fmt.Println(days, err)
	back, err := functions.DateADTToDate(d)
	fmt.Println(back.Format(time.DateOnly), err)
	// Output:
	// <nil>
	// 2025-03-01 <nil>
	// 20148 <nil>
	// 2025-03-01 <nil>
}

func ExampleIntervalToTimedelta() {
	interv, err := functions.IntervalIn("2 days 03:00:00.5", -1)
	fmt.Println(err)
	d, err := functions.IntervalToTimedelta(interv)
	fmt.Println(d, err)

	months, err := functions.IntervalIn("1 mon 2 days", -1)
	fmt.Println(err)
	_, err = functions.IntervalToTimedelta(months)
	fmt.Println(err != nil)
	// Output:
	// <nil>
	// 51h0m0.5s <nil>
	// <nil>
	// true
}

func ExampleTimeADTToTime() {
	t, err := functions.TimeIn("13:45:30.5", -1)
	fmt.Println(err)
	clock, err := functions.TimeADTToTime(t)
	fmt.Println(clock.Format("2006-01-02 15:04:05.999999"), err)
	back, err := functions.TimeToTimeADT(clock)
	fmt.Println(back == t, err)

	end, err := functions.TimeIn("24:00:00", -1)
	fmt.Println(err)
	_, err = functions.TimeADTToTime(end)
	fmt.Println(err != nil)
	// Output:
	// <nil>
	// 0000-01-01 13:45:30.5 <nil>
	// true <nil>
	// <nil>
	// true
}

func ExampleTimetzADTToTimetz() {
	t, err := functions.TimetzIn("13:45:30+02", -1)
	fmt.Println(err)
	clock, err := functions.TimetzADTToTimetz(t)
	fmt.Println(clock.Format("15:04:05-07:00"), err)
	back, err := functions.TimetzToTimetzADT(clock)
	fmt.Println(err)
	text, err := functions.TimetzOut(back)
	fmt.Println(text, err)
	// Output:
	// <nil>
	// 13:45:30+02:00 <nil>
	// <nil>
	// 13:45:30+02 <nil>
}
