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
