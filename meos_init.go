package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"

func MeosInitialize(tz string) {
	C.meos_initialize(C.CString(tz), nil)
}

func MeosFinalize() {
	C.meos_finalize()
}
