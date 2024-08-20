package gomeos

import (
	"fmt"
)

func ExampleTNumberToSpan() {
	tf_seq := TFloatIn("{1.2@2022-10-01, 2.3@2022-10-02,3.4@2022-10-03}", &TFloatSeq{})
	fs := TNumberToSpan(tf_seq, &FloatSpan{})
	fmt.Println(fs.FloatSpanOut(10))
	// Output:
	// [1.2, 3.4]
}

func ExampleTNumberValueSpans() {
	tf_seq := TFloatIn("{1.2@2022-10-01, 2.3@2022-10-02,3.4@2022-10-03}", &TFloatSeq{})
	fs := TNumberValueSpans(tf_seq, &FloatSpanSet{})
	fmt.Println(fs.FloatSpanSetOut(10))
	// Output:
	// {[1.2, 1.2], [2.3, 2.3], [3.4, 3.4]}
}
