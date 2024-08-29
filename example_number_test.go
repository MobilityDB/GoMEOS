package gomeos

import (
	"fmt"
)

func ExampleNewFloatSet() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	fmt.Println(g_fs.FloatSetOut(5))
	// Output:
	//{1.23, 2.9, 3.6}
}

func ExampleFloatSet_ToIntSet() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	g_is := g_fs.ToIntSet()
	fmt.Println(g_is.IntSetOut())
	// Output: {1, 2, 3}
}

func ExampleFloatSet_ToSpanSet() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	g_fss := g_fs.ToSpanSet()
	fmt.Println(g_fss.FloatSpanSetOut(5))
	// Output: {[1.23, 1.23], [2.9, 2.9], [3.6, 3.6]}
}

func ExampleFloatSet_NumElements() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	fmt.Println(g_fs.NumElements())
	// Output: 3
}

func ExampleFloatSet_Elements() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	fmt.Println(g_fs.Elements())
	// Output: [1.23 2.9 3.6]
}

func ExampleFloatSet_StartElement() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	fmt.Println(g_fs.StartElement())
	// Output: 1.23
}

func ExampleFloatSet_EndElement() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	fmt.Println(g_fs.EndElement())
	// Output: 3.6
}

func ExampleFloatSet_ElementN() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	fmt.Println(g_fs.ElementN(0))
	fmt.Println(g_fs.ElementN(1))
	fmt.Println(g_fs.ElementN(2))
	// Output:
	// 1.23
	// 2.9
	// 3.6
}

func ExampleFloatSet_Shift() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	shift := g_fs.Shift(1.0)
	fmt.Println(shift.FloatSetOut(5))
	// Output: {2.23, 3.9, 4.6}
}

func ExampleFloatSet_Scale() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	scale := g_fs.Scale(2.0)
	fmt.Println(scale.FloatSetOut(5))
	// Output: {1.23, 2.63928, 3.23}
}

func ExampleFloatSet_ShiftScale() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	ss := g_fs.ShiftScale(1.0, 2.0)
	fmt.Println(ss.FloatSetOut(5))
	// Output: {2.23, 3.63928, 4.23}
}

func ExampleFloatSet_Contains() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	con, _ := g_fs.Contains(2.0)
	fmt.Println(con)

	g_fs2 := NewFloatSet("{2.9, 4.70, 5.97, 1.23, 3.6}")
	conSet, _ := g_fs2.Contains(g_fs)
	fmt.Println(conSet)
	// Output:
	// false
	// false
}

func ExampleFloatSet_IsLeft() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	left, _ := g_fs.IsLeft(4.0)
	fmt.Println(left)
	// Output: true
}

func ExampleFloatSet_IsOverOrLeft() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	overleft, _ := g_fs.IsOverOrLeft(3.0)
	fmt.Println(overleft)
	// Output: false
}

func ExampleFloatSet_IsRight() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	right, _ := g_fs.IsRight(0.0)
	fmt.Println(right)
	// Output: true
}

func ExampleFloatSet_IsOverOrRight() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	overright, _ := g_fs.IsOverOrRight(3.0)
	fmt.Println(overright)
	// Output: false
}

func ExampleFloatSet_Intersection() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	g_fs2 := NewFloatSet("{2.9, 4.70, 5.97, 1.23, 3.6}")
	inter, _ := g_fs.Intersection(&g_fs2)
	fmt.Println(inter.FloatSetOut(5))
	// Output: {1.23, 2.9, 3.6}
}

func ExampleFloatSet_Mul() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	g_fs2 := NewFloatSet("{2.9, 4.70, 5.97, 1.23, 3.6}")
	mul, _ := g_fs.Mul(&g_fs2)
	fmt.Println(mul.FloatSetOut(5))
	// Output: {1.23, 2.9, 3.6}
}

func ExampleFloatSet_Minus() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	g_fs2 := NewFloatSet("{2.9, 4.70, 5.97, 1.23, 3.6}")
	inter, _ := g_fs2.Intersection(&g_fs)
	fmt.Println(inter.FloatSetOut(5))
	// Output: {1.23, 2.9, 3.6}
}

func ExampleFloatSet_Sub() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	sub, _ := g_fs.Sub(2.0)
	fmt.Println(sub.FloatSetOut(5))
	// Output: {1.23, 2.9, 3.6}
}

func ExampleFloatSet_Distance() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	dist, _ := g_fs.Distance(5.0)
	fmt.Println(dist)
	// Output: 1.4
}

func ExampleFloatSet_Union() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	g_fs2 := NewFloatSet("{2.9, 4.70, 5.97, 1.23, 3.6}")
	union, _ := g_fs.Union(&g_fs2)
	fmt.Println(union.FloatSetOut(10))
	// Output: {1.23, 2.9, 3.6, 4.7, 5.97}
}

func ExampleFloatSet_Add() {
	g_fs := NewFloatSet("{1.23, 2.9, 3.6}")
	added, _ := g_fs.Add(4.0)
	fmt.Println(added.FloatSetOut(10))
	// Output: {1.23, 2.9, 3.6, 4}
}

func ExampleNewFloatSpan() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	fmt.Println(g_fs.FloatSpanOut(10))
	// Output: [1.1, 3.3]
}

func ExampleFloatSpan_ToIntSpan() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	g_is := g_fs.ToIntSpan()
	fmt.Println(g_is.IntSpanOut())
	// Output: [1, 4)
}
func ExampleFloatSpan_Lower() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	fmt.Println(g_fs.Lower())
	// Output: 1.1
}
func ExampleFloatSpan_Upper() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	fmt.Println(g_fs.Upper())
	// Output: 3.3
}
func ExampleFloatSpan_Width() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	fmt.Println(g_fs.Width())
	// Output: 2.2
}
func ExampleFloatSpan_ShiftScale() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	shiftedScaled := g_fs.ShiftScale(1.1, 2.2)
	fmt.Println(shiftedScaled.FloatSpanOut(5))
	// Output: [2.2, 4.4]
}
func ExampleFloatSpan_Scale() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	scaled := g_fs.Scale(2.2)
	fmt.Println(scaled.FloatSpanOut(5))
	// Output: [1.1, 3.3]
}
func ExampleFloatSpan_IsAdjacent() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	is_adj, err := g_fs.IsAdjacent(3.4)
	fmt.Println(is_adj, err)
	// Output: false <nil>
}
func ExampleFloatSpan_Contains() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	contains, err := g_fs.Contains(2.2)
	fmt.Println(contains, err)
	// Output: true <nil>
}
func ExampleFloatSpan_IsSame() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	is_same, err := g_fs.IsSame(2.2)
	fmt.Println(is_same, err)
	// Output: false <nil>
}
func ExampleFloatSpan_IsLeft() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	is_left, err := g_fs.IsLeft(4.4)
	fmt.Println(is_left, err)
	// Output: true <nil>
}
func ExampleFloatSpan_IsOverOrLeft() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	is_over_or_left, err := g_fs.IsOverOrLeft(3.3)
	fmt.Println(is_over_or_left, err)
	// Output: true <nil>
}
func ExampleFloatSpan_IsRight() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	is_right, err := g_fs.IsRight(0.0)
	fmt.Println(is_right, err)
	// Output: true <nil>
}
func ExampleFloatSpan_IsOverOrRight() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	is_over_or_right, err := g_fs.IsOverOrRight(2.2)
	fmt.Println(is_over_or_right, err)
	// Output: false <nil>
}
func ExampleFloatSpan_Distance() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	g_fs_2 := NewFloatSpan("[2.2, 4.4]")
	distance, err := g_fs.Distance(g_fs_2)
	fmt.Println(distance, err)
	distance, err = g_fs.Distance(20)
	fmt.Println(distance, err)
	// Output:
	//0 <nil>
	//4.625393849793197e+18 <nil>
}
func ExampleFloatSpan_Intersection() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	g_fs_2 := NewFloatSpan("[2.2, 4.4]")
	res, err := g_fs.Intersection(g_fs_2)
	fmt.Println(res.FloatSpanOut(5), err)
	// Output: [2.2, 3.3] <nil>
}
func ExampleFloatSpan_Minus() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	res, err := g_fs.Minus(2.2)
	fmt.Println(res.FloatSpanSetOut(5), err)
	// Output: {[1.1, 2.2), (2.2, 3.3]} <nil>
}
func ExampleFloatSpan_Union() {
	g_fs := NewFloatSpan("[1.1, 3.3]")
	res, err := g_fs.Union(2.2)
	fmt.Println(res.FloatSpanSetOut(5), err)
	// Output: {[1.1, 3.3]} <nil>
}
func ExampleFloatSpanSet_FloatSpanSetOut() {
	g_fss := NewFloatSpanSet("{[1.1,3.3]}")
	fmt.Println(g_fss.FloatSpanSetOut(10))
	// Output: {[1.1, 3.3]}
}
func ExampleFloatSpanSet_ToSpan() {
	g_fss := NewFloatSpanSet("{[1.1,3.3]}")
	g_fs := g_fss.ToSpan()
	fmt.Println(g_fs.FloatSpanOut(10))
	// Output: [1.1, 3.3]
}
func ExampleFloatSpanSet_ToIntSpanSet() {
	g_fss := NewFloatSpanSet("{[1.1,4.4), [9.9,100.0]}")
	ciss := g_fss.ToIntSpanSet()
	fmt.Println(ciss.IntSpanSetOut())
	// Output: {[1, 4), [9, 101)}
}
func ExampleFloatSpanSet_NumSpans() {
	g_fss := NewFloatSpanSet("{[1.1,4.4), [9.9,100.0]}")
	fmt.Println(g_fss.NumSpans())
	// Output: 2
}
func ExampleFloatSpanSet_Width() {
	g_fss := NewFloatSpanSet("{[1.1,4.4), [9.9,100.0]}")
	fmt.Println(g_fss.Width(true))
	fmt.Println(g_fss.Width(false))
	// Output: 98.9
	// 93.39999999999999
}

func createIntSet() IntSet {
	return NewIntSet("{1, 2, 3}")
}

func createIntSet2() IntSet {
	return NewIntSet("{3, 4, 5}")
}

func ExampleNewIntSet() {
	g_is := createIntSet()
	fmt.Println(g_is.IntSetOut())
	// Output:
	// {1, 2, 3}
}

func ExampleIntSet_ToFloatSet() {
	g_is := createIntSet()
	g_fs := g_is.ToFloatSet()
	fmt.Println(g_fs.FloatSetOut(5))
	// Output:
	// {1, 2, 3}
}

func ExampleIntSet_ToSpanSet() {
	g_is := createIntSet()
	g_iss := g_is.ToSpanSet()
	fmt.Println(g_iss.IntSpanSetOut())
	// Output:
	// {[1, 4)}
}

func ExampleIntSet_NumElements() {
	g_is := createIntSet()
	fmt.Println(g_is.NumElements())
	// Output:
	// 3
}

func ExampleIntSet_Elements() {
	g_is := createIntSet()
	fmt.Println(g_is.Elements())
	// Output:
	// [1 2 3]
}

func ExampleIntSet_StartElement() {
	g_is := createIntSet()
	fmt.Println(g_is.StartElement())
	// Output:
	// 1
}

func ExampleIntSet_EndElement() {
	g_is := createIntSet()
	fmt.Println(g_is.EndElement())
	// Output:
	// 3
}

func ExampleIntSet_ElementN() {
	g_is := createIntSet()
	fmt.Println(g_is.ElementN(0))
	fmt.Println(g_is.ElementN(1))
	fmt.Println(g_is.ElementN(2))
	// Output:
	// 1
	// 2
	// 3
}

func ExampleIntSet_Shift() {
	g_is := createIntSet()
	g_is_shift := g_is.Shift(1)
	fmt.Println(g_is_shift.IntSetOut())
	// Output:
	// {2, 3, 4}
}

func ExampleIntSet_Scale() {
	g_is := createIntSet()
	g_is_scale := g_is.Scale(2)
	fmt.Println(g_is_scale.IntSetOut())
	// Output:
	// {1, 2, 4}
}

func ExampleIntSet_ShiftScale() {
	g_is := createIntSet()
	g_is_ss := g_is.ShiftScale(1, 2)
	fmt.Println(g_is_ss.IntSetOut())
	// Output:
	// {2, 3, 5}
}

func ExampleIntSet_Contains() {
	g_is := createIntSet()
	con, _ := g_is.Contains(2)
	fmt.Println(con)
	g_is2 := createIntSet2()
	conSet, _ := g_is2.Contains(g_is)
	fmt.Println(conSet)
	// Output:
	// true
	// false
}

func ExampleIntSet_IsLeft() {
	g_is := createIntSet()
	left, _ := g_is.IsLeft(4)
	fmt.Println(left)
	// Output:
	// true
}

func ExampleIntSet_IsOverOrLeft() {
	g_is := createIntSet()
	overleft, _ := g_is.IsOverOrLeft(3)
	fmt.Println(overleft)
	// Output:
	// true
}

func ExampleIntSet_IsRight() {
	g_is := createIntSet()
	right, _ := g_is.IsRight(0)
	fmt.Println(right)
	// Output:
	// true
}

func ExampleIntSet_IsOverOrRight() {
	g_is := createIntSet()
	overright, _ := g_is.IsOverOrRight(3)
	fmt.Println(overright)
	// Output:
	// false
}

func ExampleIntSet_Intersection() {
	g_is := createIntSet()
	g_is2 := createIntSet2()
	inter, _ := g_is.Intersection(&g_is2)
	fmt.Println(inter.IntSetOut())
	// Output:
	// {3}
}

func ExampleIntSet_Minus() {
	g_is := createIntSet()
	g_is2 := createIntSet2()
	mins, _ := g_is.Minus(&g_is2)
	fmt.Println(mins.IntSetOut())
	// Output:
	// {1, 2}
}

func ExampleIntSet_Distance() {
	g_is := createIntSet()
	dist, _ := g_is.Distance(5)
	fmt.Println(dist)
	// Output:
	// 2
}

func ExampleIntSet_Sub() {
	g_is := createIntSet()
	sub, _ := g_is.Sub(2)
	fmt.Println(sub.IntSetOut())
	// Output:
	// {1, 3}
}

func ExampleIntSet_Mul() {
	g_is := createIntSet()
	g_is2 := createIntSet2()
	mul, _ := g_is.Mul(&g_is2)
	fmt.Println(mul.IntSetOut())
	// Output:
	// {3}
}

func ExampleIntSet_Union() {
	g_is := createIntSet()
	g_is2 := createIntSet2()
	union, _ := g_is.Union(&g_is2)
	fmt.Println(union.IntSetOut())
	// Output:
	// {1, 2, 3, 4, 5}
}

func ExampleIntSet_Add() {
	g_is := createIntSet()
	added, _ := g_is.Add(4)
	fmt.Println(added.IntSetOut())
	// Output:
	// {1, 2, 3, 4}
}

func createIntSpan() *IntSpan {
	g_is_in := "[4,5]"
	g_is := NewIntSpan(g_is_in)
	return g_is
}

func createIntSpan2() *IntSpan {
	g_is_in := "[1,3]"
	g_is := NewIntSpan(g_is_in)
	return g_is
}

func ExampleIntSpan_IntSpanOut() {
	g_is := createIntSpan()
	fmt.Println(g_is.IntSpanOut())
	// Output:
	// [4, 6)
}

func ExampleIntSpan_ToSpanSet() {
	g_is := createIntSpan()
	g_iss := g_is.ToSpanSet()
	fmt.Println(g_iss.IntSpanSetOut())
	// Output:
	// {[4, 6)}
}

func ExampleIntSpan_ToFloatSpan() {
	g_is := createIntSpan()
	g_fs := g_is.ToFloatSpan()
	fmt.Println(g_fs.FloatSpanOut(10))
	// Output:
	// [4, 5]
}

func ExampleIntSpan_IsAdjacent() {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.IsAdjacent(g_is2)
	fmt.Println(res)
	fmt.Println(err)
	res, err = g_is.IsAdjacent(5)
	fmt.Println(res)
	fmt.Println(err)
	// Output:
	// true
	// <nil>
	// false
	// <nil>
}

func ExampleIntSpan_Contains() {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.Contains(g_is2)
	fmt.Println(res)
	fmt.Println(err)
	res, err = g_is.Contains(200)
	fmt.Println(res)
	fmt.Println(err)
	// Output:
	// false
	// <nil>
	// false
	// <nil>
}

func ExampleIntSpan_IsSame() {
	g_is := createIntSpan()
	res, err := g_is.IsSame(200)
	fmt.Println(res)
	fmt.Println(err)
	g_is2 := createIntSpan()
	res, err = g_is.IsSame(g_is2)
	fmt.Println(res)
	fmt.Println(err)
	// Output:
	// false
	// <nil>
	// true
	// <nil>
}

func ExampleIntSpan_IsLeft() {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.IsLeft(g_is2)
	fmt.Println(res)
	fmt.Println(err)
	res, err = g_is.IsLeft(6)
	fmt.Println(res)
	fmt.Println(err)
	// Output:
	// false
	// <nil>
	// true
	// <nil>
}

func ExampleIntSpan_IsOverOrLeft() {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.IsOverOrLeft(g_is2)
	fmt.Println(res)
	fmt.Println(err)
	res, err = g_is.IsOverOrLeft(4)
	fmt.Println(res)
	fmt.Println(err)
	// Output:
	// false
	// <nil>
	// false
	// <nil>
}

func ExampleIntSpan_IsRight() {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.IsRight(g_is2)
	fmt.Println(res)
	fmt.Println(err)
	res, err = g_is.IsRight(6)
	fmt.Println(res)
	fmt.Println(err)
	// Output:
	//true
	//<nil>
	//false
	//<nil>

}

func ExampleIntSpan_IsOverOrRight() {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.IsOverOrRight(g_is2)
	fmt.Println(res)
	fmt.Println(err)
	res, err = g_is.IsOverOrRight(6)
	fmt.Println(res)
	fmt.Println(err)
	// Output:
	//true
	//<nil>
	//false
	//<nil>
}

func ExampleIntSpan_Union() {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.Union(g_is2)
	fmt.Println(res.IntSpanSetOut())
	fmt.Println(err)
	// Output:
	//{[1, 6)}
	//<nil>

}

func ExampleIntSpan_Minus() {
	g_is := createIntSpan()
	g_is2 := createIntSpan2()
	res, err := g_is.Minus(g_is2)
	fmt.Println(res.IntSpanSetOut())
	fmt.Println(err)
	// Output:
	//{[4, 6)}
	//<nil>
}

func ExampleIntSpan_Distance() {
	g_is := createIntSpan()
	res, err := g_is.Distance(200)
	fmt.Println(res)
	fmt.Println(err)
	// Output:
	//195
	//<nil>
}

func ExampleIntSpan_Add() {
	g_is := createIntSpan()
	res, err := g_is.Add(200)
	fmt.Println(res.IntSpanSetOut())
	fmt.Println(err)
	// Output:
	//{[4, 6), [200, 201)}
	//<nil>
}
