package gomeos

import (
	"fmt"
	"time"
)

func ExampleTBoolFromBaseTemporal() {
	tb_seq := TBoolIn("{FALSE@2022-10-01, FALSE@2022-10-02,TRUE@2022-10-03}", &TBoolSeq{})
	res := TBoolFromBaseTemporal(true, tb_seq, &TBoolSeq{})
	fmt.Println(TBoolOut(res))
	// Output:
	// {t@2022-10-01 00:00:00+00, t@2022-10-02 00:00:00+00, t@2022-10-03 00:00:00+00}
}
func ExampleTBoolSeqSetFromBaseTstzspanset() {
	tstzspanset := NewTsTzSpanSet("{[2023-01-01 08:00:00+00, 2023-01-02 08:00:00+00), (2023-01-03 10:00:00+00, 2023-01-04 12:00:00+00]}")
	tss := TBoolSeqSetFromBaseTstzspanset(true, tstzspanset)
	fmt.Println(TBoolOut(tss))
	// Output:
	// {[t@2023-01-01 08:00:00+00, t@2023-01-02 08:00:00+00), (t@2023-01-03 10:00:00+00, t@2023-01-04 12:00:00+00]}
}

func ExampleTBoolOut() {
	g_is := NewTBoolInst("TRUE@2022-10-01")
	fmt.Println(TBoolOut(g_is))
	// Output:
	// t@2022-10-01 00:00:00+00
}

func ExampleTBoolValueAtTimestamp() {
	g_is := NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02,TRUE@2022-10-03}")
	ts, _ := time.Parse("2006-01-02", "2022-10-01")
	res := TBoolValueAtTimestamp(g_is, ts)
	fmt.Println(res)
	// Output:
	// false
}

func ExampleEverEqTBoolBool() {
	tb1 := NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02,FALSE@2022-10-03}")
	tb2 := true
	res := EverEqTBoolBool(tb1, tb2)
	fmt.Println(res)
	// Output:
	// false

}

func ExampleTEqTBoolBool() {
	tb1 := NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02,FALSE@2022-10-03}")
	tb2 := false
	res := TEqTBoolBool(tb1, tb2)
	fmt.Println(res.String())
	// Output:
	// {t@2022-10-01 00:00:00+00, t@2022-10-02 00:00:00+00, t@2022-10-03 00:00:00+00}

}

func ExampleTNEqTBoolBool() {
	tb1 := NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02,FALSE@2022-10-03}")
	tb2 := false
	res := TNEqTBoolBool(tb1, tb2)
	fmt.Println(res.String())
	// Output:
	// {f@2022-10-01 00:00:00+00, f@2022-10-02 00:00:00+00, f@2022-10-03 00:00:00+00}

}

func ExampleTAndTBoolBool() {
	tb1 := NewTBoolSeq("{TRUE@2022-10-01, TRUE@2022-10-02,TRUE@2022-10-03}")
	tb2 := true
	res := TAndTBoolBool(tb1, tb2)
	fmt.Println(res.String())
	// Output:
	// {t@2022-10-01 00:00:00+00, t@2022-10-02 00:00:00+00, t@2022-10-03 00:00:00+00}

}

func ExampleTNotTBool() {
	tb1 := NewTBoolSeq("{TRUE@2022-10-01, TRUE@2022-10-02,TRUE@2022-10-03}")
	res := TNotTBool(tb1, &TBoolSeq{})
	fmt.Println(res.String())
	// Output:
	// {f@2022-10-01 00:00:00+00, f@2022-10-02 00:00:00+00, f@2022-10-03 00:00:00+00}

}

func ExampleTBoolWhenTrue() {
	tb1 := NewTBoolSeq("{TRUE@2022-10-01, FALSE@2022-10-02,TRUE@2022-10-03}")
	res := TBoolWhenTrue(tb1)
	fmt.Println(res.TsTzSpanSetOut())
	// Output:
	// {[2022-10-01 00:00:00+00, 2022-10-01 00:00:00+00], [2022-10-03 00:00:00+00, 2022-10-03 00:00:00+00]}

}
