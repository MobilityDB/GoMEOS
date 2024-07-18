package main

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"

	"github.com/MobilityDB/GoMEOS/collections/number"
)

func main() {
	C.meos_initialize(nil, nil)
	g_iss_in := "{[1,3], [8,10]}"
	g_iss := number.NewIntSpanSet(g_iss_in)
	g_iss_out := g_iss.IntSpanSetOut()
	w := g_iss.Width()
	sp := g_iss.SpanN(2)
	fmt.Printf("WKT: %s\nMFJSON: %s\nWidth: %d\n", g_iss_in, g_iss_out, w)
	fmt.Println(sp.IntSpanOut())
	fmt.Println(g_iss.NumSpans())
	spans := g_iss.Spans()
	for i, s := range spans {
		fmt.Println(i)
		fmt.Println(s.IntSpanOut())
	}
	// g_is_in := "(2, 5]"
	// g_is := number.NewIntSpan(g_is_in)
	// g_is_out := g_is.IntSpanOut()
	// fmt.Printf("WKT: %s\nMFJSON: %s\n", g_is_in, g_is_out)
}
