package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"

*/
import "C"

type TBox struct {
	_inner *C.TBox
}

func TBoxOut(tbox *TBox, max_decimals int) string {
	return C.GoString(C.tbox_out(tbox._inner, C.int(max_decimals)))
}
