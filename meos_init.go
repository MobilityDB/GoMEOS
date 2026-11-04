package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import "unsafe"

func MeosInitialize(tz string) {
	C.meos_initialize()
	if tz != "" {
		c_tz := C.CString(tz)
		defer C.free(unsafe.Pointer(c_tz))
		C.meos_initialize_timezone(c_tz)
	}
}

func MeosFinalize() {
	C.meos_finalize()
}
