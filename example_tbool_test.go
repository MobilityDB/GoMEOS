package gomeos

import (
	"fmt"
	"testing"
)

func ExampleTBoolOut() {
	MeosInitialize("UTC")
	g_is := NewTBoolInst("TRUE@2022-10-01")
	fmt.Println(TBoolOut(g_is))
	// Output:
	// t@2022-10-01 00:00:00+00
}

func TestTPointToSTBox(t *testing.T) {
	MeosInitialize("UTC")
	tg := TGeomPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &TGeomPointSeq{})
	stbox := TPointToSTBox(tg)
	fmt.Println(STBoxOut(stbox, 10))
	// Output:
	// STBOX XT(((1,1),(2,2)),[2022-10-01 00:00:00+00, 2022-10-02 00:00:00+00])
}
