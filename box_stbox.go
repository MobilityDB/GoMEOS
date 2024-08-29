package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"

*/
import "C"

type STBox struct {
	_inner *C.STBox
}

func STBoxOut(stbox *STBox, max_decimals int) string {
	return C.GoString(C.stbox_out(stbox._inner, C.int(max_decimals)))
}
