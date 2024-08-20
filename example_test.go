package gomeos

import (
	"fmt"
)

//func ExampleTBoolFromBaseTemporal() {
//	tb_seq := TBoolIn("{FALSE@2022-10-01, FALSE@2022-10-02,TRUE@2022-10-03}", &TBoolSeq{})
//	res := TBoolFromBaseTemporal(true, tb_seq, &TBoolSeq{})
//	fmt.Println(TBoolOut(res))
//	// Output:
//	// {t@2022-10-01 00:00:00+00, t@2022-10-02 00:00:00+00, t@2022-10-03 00:00:00+00}
//}

func ExampleTPointToSTBox() {
	tg := TGeomPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &TGeomPointSeq{})
	stbox := TPointToSTBox(tg)
	fmt.Println(STBoxOut(stbox, 10))
	// Output:
	// STBOX XT(((1,1),(2,2)),[2022-10-01 00:00:00+00, 2022-10-02 00:00:00+00])
}
