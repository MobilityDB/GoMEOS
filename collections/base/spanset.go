package base

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

type SpanSet struct {
	_inner *C.SpanSet
}
