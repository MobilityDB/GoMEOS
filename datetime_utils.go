package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import (
	"time"
	"unsafe"

	"github.com/leekchan/timeutil"
)

func DateToDateADT(t time.Time) C.DateADT {
	c_date_in := C.CString(t.Format("2006-01-02"))
	defer C.free(unsafe.Pointer(c_date_in))
	c_date := C.pg_date_in(c_date_in)
	return c_date
}

func DateADTToDate(d C.DateADT) time.Time {
	dateStr := C.GoString(C.pg_date_out(d))
	layout := "2006-01-02"
	parsedDate, _ := time.Parse(layout, dateStr)
	return parsedDate
}

func IntervalToTimeDelta(interval C.Interval) timeutil.Timedelta {
	microsecond := int(interval.time)
	day := int(interval.day)
	month := int(interval.month)
	dr := timeutil.Timedelta{
		Microseconds: time.Duration(microsecond),
		Days:         time.Duration(day) + time.Duration(month*30),
	}
	return dr
}

func TimeDeltaToInterval(td timeutil.Timedelta) C.Interval {
	var interval C.Interval
	interval.time = C.TimeOffset(td.Microseconds)
	interval.day = C.int(td.Days)
	interval.month = C.int(0)
	return interval
}

func TimestamptzIn(timeStr string) C.TimestampTz {
	CtimeStr := C.CString(timeStr)
	defer C.free(unsafe.Pointer(CtimeStr))
	tstz := C.pg_timestamptz_in(CtimeStr, C.int(-1))
	return tstz
}

func TimestamptzOut(ts C.TimestampTz) string {
	timeStr := C.GoString(C.pg_timestamptz_out(ts))
	return timeStr
}

// Constants for the transformation
const (
	postgresEpoch             = 946684800000000 // Number of microseconds from Go epoch to PostgreSQL epoch (2000-01-01)
	microsecondsToNanoseconds = 1000            // Conversion factor from microseconds to nanoseconds
)

// Transform TimestampTz to time.Time
func TimestamptzToDatetime(tstz C.TimestampTz) time.Time {
	// Convert PostgreSQL timestamp to nanoseconds since Go epoch
	nanoseconds := (tstz + postgresEpoch) * microsecondsToNanoseconds
	return time.Unix(0, int64(nanoseconds)).UTC() // UTC time
}

// Transform time.Time to TimestampTz
func DatetimeToTimestamptz(t time.Time) C.TimestampTz {
	// Convert Go time to nanoseconds since Go epoch
	nanoseconds := t.UnixNano()
	// Subtract the PostgreSQL epoch and convert to microseconds
	return C.int64((nanoseconds / microsecondsToNanoseconds) - postgresEpoch)
}
