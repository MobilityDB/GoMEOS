package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
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

func TimestamptzToDatetime(ts C.TimestampTz) time.Time {
	timeStr := C.GoString(C.pg_timestamptz_out(ts))
	layout := "2006-01-02 15:04:05-07"
	parsedDate, _ := time.Parse(layout, timeStr)
	return parsedDate
}

func DatetimeToTimestamptz(t time.Time) C.TimestampTz {
	timeStr := t.Format("2006-01-02 15:04:05")
	CtimeStr := C.CString(timeStr)
	defer C.free(unsafe.Pointer(CtimeStr))
	tstz := C.pg_timestamptz_in(CtimeStr, C.int(-1))
	return tstz
}

func TimestamptzOut(ts C.TimestampTz) string {
	timeStr := C.GoString(C.pg_timestamptz_out(ts))
	return timeStr
}
