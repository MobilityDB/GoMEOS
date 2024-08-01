package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

func MeosInitialize() {
	// TODO: initialize timezone and etc, implement the real meos initialize
	C.meos_initialize(C.CString("UTC-2"), nil)
}

func MeosFinalize() {
	C.meos_finalize()
}
