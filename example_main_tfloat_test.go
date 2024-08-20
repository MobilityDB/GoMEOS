package gomeos

import (
	"fmt"
	"time"
)

func ExampleTFloatFromBaseTemporal() {
	tb_seq := TBoolIn("{FALSE@2022-10-01, FALSE@2022-10-02,TRUE@2022-10-03}", &TBoolSeq{})
	res := TFloatFromBaseTemporal(2, tb_seq, &TFloatSeq{})
	fmt.Println(TFloatOut(res, 10))
	// Output:
	// {2@2022-10-01 00:00:00+00, 2@2022-10-02 00:00:00+00, 2@2022-10-03 00:00:00+00}
}

func ExampleTFloatToTInt() {
	tf_seq := TFloatIn("{1.2@2022-10-01, 2.3@2022-10-02,3.4@2022-10-03}", &TFloatSeq{})
	res, _ := TFloatToTInt(tf_seq, &TIntSeq{})
	fmt.Println(TIntOut(res))
	// Output:
	// {1@2022-10-01 00:00:00+00, 2@2022-10-02 00:00:00+00, 3@2022-10-03 00:00:00+00}
}

func ExampleTFloatValues() {
	tf_seq := TFloatIn("{1.2@2022-10-01, 2.3@2022-10-02,3.4@2022-10-03}", &TFloatSeq{})
	res, _ := TFloatValues(tf_seq)
	fmt.Println(res)
	// Output:
	// [1.2 2.3 3.4]
}

func ExampleTFloatValueAtTimestamp() {
	tf_seq := TFloatIn("{1.2@2022-10-01, 2.3@2022-10-02,3.4@2022-10-03}", &TFloatSeq{})
	ts, _ := time.Parse("2006-01-02", "2022-10-01")
	res := TFloatValueAtTimestamp(tf_seq, ts)
	fmt.Println(res)
	// Output:
	// 1.2
}
