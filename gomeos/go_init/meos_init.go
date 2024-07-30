package go_init

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib
#include "meos.h"
#include "meos_catalog.h"
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
