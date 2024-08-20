package gomeos

import (
	"fmt"
	"testing"
)

//import (
//	"fmt"
//	"testing"
//)
//
//func TestTBoolFromBaseTemporal(t *testing.T) {
//	tb_seq := TBoolIn("{FALSE@2022-10-01, FALSE@2022-10-02,TRUE@2022-10-03}", &TBoolSeq{})
//	res := TBoolFromBaseTemporal(true, tb_seq, &TBoolSeq{})
//	fmt.Println(TBoolOut(res))
//	// Output:
//	// {t@2022-10-01 00:00:00+00, t@2022-10-02 00:00:00+00, t@2022-10-03 00:00:00+00}
//}
//
//func TestTBoolSeqSetFromBaseTstzspanset(t *testing.T) {
//	tstzspanset := NewTsTzSpanSet("{[2023-01-01 08:00:00+00, 2023-01-02 08:00:00+00), (2023-01-03 10:00:00+00, 2023-01-04 12:00:00+00]}")
//	tss := TBoolSeqSetFromBaseTstzspanset(true, tstzspanset)
//	fmt.Println(TBoolOut(tss))
//	// Output:
//	// {[t@2023-01-01 08:00:00+00, t@2023-01-02 08:00:00+00), (t@2023-01-03 10:00:00+00, t@2023-01-04 12:00:00+00]}
//}
//
//func TestTBoolOut(t *testing.T) {
//	g_is := NewTBoolInst("TRUE@2022-10-01")
//	fmt.Println(TBoolOut(g_is))
//	// Output:
//	// t@2022-10-01 00:00:00+00
//}
//
//func TestTPointToSTBox(t *testing.T) {
//	tg := TGeomPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &TGeomPointSeq{})
//	stbox := TPointToSTBox(tg)
//	fmt.Println(STBoxOut(stbox, 10))
//	// Output:
//	// STBOX XT(((1,1),(2,2)),[2022-10-01 00:00:00+00, 2022-10-02 00:00:00+00])
//}

func TestTNumberToSpan(t *testing.T) {
	tf_seq := TFloatIn("{1.2@2022-10-01, 2.3@2022-10-02,3.4@2022-10-03}", &TFloatSeq{})
	fs := TNumberToSpan(tf_seq, &FloatSpan{})
	fmt.Println(fs.FloatSpanOut(10))
	// Output
	// [1.2, 3.4]
}
