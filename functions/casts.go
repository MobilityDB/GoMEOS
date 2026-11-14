package functions

// Casts between the time types of MEOS and those of Go. Each converts through
// the generated wrapper of a MEOS conversion and builds the Go value with the
// constructor of its type, as the casts of JMEOS and PyMEOS-CFFI do
// (doc/specs/meos-bindings-algebra.md of MobilityDB, section 5). A value MEOS
// cannot state answers the error of MEOS, a value Go cannot state the error of
// the cast. tools/codegen.py copies this file, tools/templates/casts.go.tmpl,
// verbatim into functions/casts.go.

/*
#include "meos.h"
*/
import "C"
import (
	"fmt"
	"math"
	"time"
	"unsafe"
)

const (
	nanosPerMicro = int64(time.Microsecond)
	microsPerSec  = int64(time.Second / time.Microsecond)
	secondsPerDay = int64(24 * 60 * 60)
	microsPerDay  = secondsPerDay * microsPerSec
)

// wholeMicros returns the microseconds of a count of nanoseconds, which MEOS
// states in whole microseconds.
func wholeMicros(nanos int64, value any) (int64, error) {
	if nanos%nanosPerMicro != 0 {
		return 0, fmt.Errorf("a MEOS value holds whole microseconds: %v", value)
	}
	return nanos / nanosPerMicro, nil
}

// unixMicros returns the microseconds since the Unix epoch of the wall clock of
// t read in UTC.
func unixMicros(t time.Time) (int64, error) {
	sec := t.Unix()
	if sec > math.MaxInt64/microsPerSec-1 || sec < math.MinInt64/microsPerSec+1 {
		return 0, fmt.Errorf("a MEOS timestamp holds a 64-bit count of microseconds: %v", t)
	}
	us, err := wholeMicros(int64(t.Nanosecond()), t)
	if err != nil {
		return 0, err
	}
	return sec*microsPerSec + us, nil
}

// DatetimeToTimestamptz casts t through unix_us_to_timestamptz.
func DatetimeToTimestamptz(t time.Time) (int64, error) {
	us, err := unixMicros(t)
	if err != nil {
		return 0, err
	}
	return UnixUsToTimestamptz(us)
}

// TimestamptzToDatetime casts a timestamp with time zone through
// timestamptz_to_unix_us, in UTC.
func TimestamptzToDatetime(ts int64) (time.Time, error) {
	us, err := TimestamptzToUnixUs(ts)
	if err != nil {
		return time.Time{}, err
	}
	return time.UnixMicro(us).UTC(), nil
}

// DatetimeToTimestamp casts the wall clock of t through unix_us_to_timestamp.
func DatetimeToTimestamp(t time.Time) (int64, error) {
	us, err := unixMicros(time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(),
		t.Second(), t.Nanosecond(), time.UTC))
	if err != nil {
		return 0, err
	}
	return UnixUsToTimestamp(us)
}

// TimestampToDatetime casts a timestamp through timestamp_to_unix_us, its wall
// clock in UTC.
func TimestampToDatetime(ts int64) (time.Time, error) {
	us, err := TimestampToUnixUs(ts)
	if err != nil {
		return time.Time{}, err
	}
	return time.UnixMicro(us).UTC(), nil
}

// DateToDateADT casts the date of t through unix_days_to_date.
func DateToDateADT(t time.Time) (int32, error) {
	days := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).Unix() / secondsPerDay
	if days < math.MinInt32 || days > math.MaxInt32 {
		return 0, fmt.Errorf("a MEOS date holds a 32-bit count of days: %v", t)
	}
	return UnixDaysToDate(int32(days))
}

// DateADTToDate casts a date through date_to_unix_days, at midnight UTC.
func DateADTToDate(d int32) (time.Time, error) {
	days, err := DateToUnixDays(d)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(days)*secondsPerDay, 0).UTC(), nil
}

// TimedeltaToInterval casts d through interval_from_parts, its whole days as
// days and the rest as microseconds.
func TimedeltaToInterval(d time.Duration) (*Interval, error) {
	day := 24 * time.Hour
	us, err := wholeMicros(int64(d%day), d)
	if err != nil {
		return nil, err
	}
	return IntervalFromParts(0, int32(d/day), us)
}

// IntervalToTimedelta casts an interval through interval_to_parts; an interval
// with months, or beyond the range of a time.Duration, is refused.
func IntervalToTimedelta(interv *Interval) (time.Duration, error) {
	var months, days C.int32
	var us C.int64
	if _, err := IntervalToParts(interv, unsafe.Pointer(&months), unsafe.Pointer(&days),
		unsafe.Pointer(&us)); err != nil {
		return 0, err
	}
	if months != 0 {
		return 0, fmt.Errorf("a time.Duration cannot state an interval of %d months", int(months))
	}
	total := int64(days)*microsPerDay + int64(us)
	if total > math.MaxInt64/nanosPerMicro || total < math.MinInt64/nanosPerMicro {
		return 0, fmt.Errorf("a time.Duration cannot state an interval of %d days", int(days))
	}
	return time.Duration(total) * time.Microsecond, nil
}

// clockMicros returns the microseconds since midnight of the clock of t.
func clockMicros(t time.Time) (int64, error) {
	ns := int64(t.Hour())*int64(time.Hour) + int64(t.Minute())*int64(time.Minute) +
		int64(t.Second())*int64(time.Second) + int64(t.Nanosecond())
	return wholeMicros(ns, t)
}

// clockTime returns the time.Time at us microseconds after the start of the
// first day of year 0 in loc, as time.Parse answers a clock; 24:00:00 is refused.
func clockTime(us int64, loc *time.Location) (time.Time, error) {
	if us >= microsPerDay {
		return time.Time{}, fmt.Errorf("a time.Time cannot state the clock 24:00:00")
	}
	return time.Date(0, time.January, 1, 0, 0, 0, 0, loc).Add(time.Duration(us) * time.Microsecond), nil
}

// TimeToTimeADT casts the clock of t through time_from_parts.
func TimeToTimeADT(t time.Time) (int64, error) {
	us, err := clockMicros(t)
	if err != nil {
		return 0, err
	}
	return TimeFromParts(us)
}

// TimeADTToTime casts a time through time_to_parts.
func TimeADTToTime(t int64) (time.Time, error) {
	var us C.int64
	if _, err := TimeToParts(t, unsafe.Pointer(&us)); err != nil {
		return time.Time{}, err
	}
	return clockTime(int64(us), time.UTC)
}

// TimetzToTimetzADT casts the clock of t and the offset of its zone through
// timetz_from_parts.
func TimetzToTimetzADT(t time.Time) (unsafe.Pointer, error) {
	us, err := clockMicros(t)
	if err != nil {
		return nil, err
	}
	_, offset := t.Zone()
	return TimetzFromParts(us, int32(offset))
}

// TimetzADTToTimetz casts a time with time zone through timetz_to_parts, its
// offset as a fixed zone.
func TimetzADTToTimetz(timetz unsafe.Pointer) (time.Time, error) {
	var us C.int64
	var offset C.int32
	if _, err := TimetzToParts(timetz, unsafe.Pointer(&us), unsafe.Pointer(&offset)); err != nil {
		return time.Time{}, err
	}
	return clockTime(int64(us), time.FixedZone("", int(offset)))
}
