package gomeos_test

import (
	"fmt"
	"testing"

	gomeos "github.com/MobilityDB/GoMEOS"
)

func ExampleTBoolOut() {
	gomeos.MeosInitialize("UTC")
	g_is := gomeos.NewTBoolInst("TRUE@2022-10-01")
	fmt.Println(gomeos.TBoolOut(g_is))
	// Output:
	// t@2022-10-01 00:00:00+00
}

func TestTPointToSTBox(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	tg := gomeos.TGeomPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &gomeos.TGeomPointSeq{})
	stbox := gomeos.TPointToSTBox(tg)
	fmt.Println(gomeos.STBoxOut(stbox, 10))
	// Output:
	// STBOX XT(((1,1),(2,2)),[2022-10-01 00:00:00+00, 2022-10-02 00:00:00+00])
}
