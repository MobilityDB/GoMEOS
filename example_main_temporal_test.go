package gomeos

import "fmt"

func ExampleEverEqTemporalTemporal() {
	MeosInitialize("UTC")
	tb1 := NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02,FALSE@2022-10-03}")
	tb2 := NewTBoolInst("TRUE@2022-10-01")
	res := EverEqTemporalTemporal(tb1, tb2)
	fmt.Println(res)
	// Output:
	// false
}

func ExampleTEqTemporalTemporal() {
	MeosInitialize("UTC")
	tb1 := NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02,FALSE@2022-10-03}")
	tb2 := NewTBoolInst("TRUE@2022-10-01")
	res := TEqTemporalTemporal(tb1, tb2)
	fmt.Println(res)
	// Output:
	// f@2022-10-01 00:00:00+00
}

func ExampleTNEqTemporalTemporal() {
	MeosInitialize("UTC")
	tb1 := NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02,FALSE@2022-10-03}")
	tb2 := NewTBoolInst("TRUE@2022-10-01")
	res := TNEqTemporalTemporal(tb1, tb2)
	fmt.Println(res)
	// Output:
	// t@2022-10-01 00:00:00+00
}

//func ExampleTNumberToSpan() {
//	MeosInitialize("UTC")
//	tf_seq := TFloatIn("{1.2@2022-10-01, 2.3@2022-10-02,3.4@2022-10-03}", &TFloatSeq{})
//	fs := TNumberToSpan(tf_seq, &FloatSpan{})
//	fmt.Println(fs.FloatSpanOut(10))
//	// Output
//	// 1
//}
