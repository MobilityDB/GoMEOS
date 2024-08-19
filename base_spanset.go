package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"

*/
import "C"

func SpansetUnionTransfn[SS SpanSet](state SS, ss SS, output SS) SS {
	c_ss := C.spanset_union_transfn(state.Inner(), ss.Inner())
	output.Init(c_ss)
	return output
}

func SpansetUnionFinalfn[SS SpanSet](state SS, output SS) SS {
	c_ss := C.spanset_union_finalfn(state.Inner())
	output.Init(c_ss)
	return output
}
