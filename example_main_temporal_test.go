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

func ExampleTemporalAsWKB() {
	tb1 := NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02,FALSE@2022-10-03}")
	wkb, _ := TemporalAsWKB(tb1)
	res := TemporalFromWKB(wkb)
	fmt.Println(res.String())
	// Output:
	// {f@2022-10-01 00:00:00+00, f@2022-10-02 00:00:00+00, f@2022-10-03 00:00:00+00}
}

func ExampleTemporalAsHexWKB() {
	tb1 := NewTBoolSeq("{FALSE@2022-10-01, FALSE@2022-10-02,FALSE@2022-10-03}")
	wkb, _ := TemporalAsHexWKB(tb1)
	fmt.Println(wkb)
	// Output:
	// 011A000603000000030000E06E8FEC8C020000004046AD008D02000000A01DCB148D0200
}
