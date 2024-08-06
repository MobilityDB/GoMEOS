package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

func MeosInitialize(tz string) {
	// TODO: initialize timezone and etc, implement the real meos initialize
	C.meos_initialize(C.CString(tz), nil)
}

func MeosFinalize() {
	C.meos_finalize()
}
