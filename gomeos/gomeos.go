package gomeos

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#define gunion_set_set union_set_set
#define gunion_set_float union_set_float
*/
import "C"
