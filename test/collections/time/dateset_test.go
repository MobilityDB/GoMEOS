package time_test

import (
	"testing"
	"time"

	gomeos "github.com/MobilityDB/GoMEOS"
	"github.com/alecthomas/assert/v2"
	"github.com/leekchan/timeutil"
)

func createDateSet() *gomeos.DateSet {
	return gomeos.NewDateSet("{2019-09-08, 2019-09-10, 2019-12-13}")
}

func createDateSet2() *gomeos.DateSet {
	return gomeos.NewDateSet("{2019-09-08, 2019-09-10, 2019-12-13, 2019-12-18}")
}

func TestNewDateSet(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	g_is := createDateSet()
	assert.Equal(t, g_is.DateSetOut(), "{2019-09-08, 2019-09-10, 2019-12-13}")
	// // gomeos.MeosFinalize()
}

func TestDSDuration(t *testing.T) {
	// // gomeos.MeosInitialize()
	g_is := createDateSet()
	assert.Equal(t, g_is.Duration(), timeutil.Timedelta{Microseconds: 0, Days: 97})
	// // gomeos.MeosFinalize()
}

func TestDSStartElement(t *testing.T) {
	// // gomeos.MeosInitialize()
	g_is := createDateSet()
	assert.Equal(t, g_is.StartElement().Format("2006-01-02"), "2019-09-08")
	// // gomeos.MeosFinalize()
}

func TestDSEndElement(t *testing.T) {
	// // gomeos.MeosInitialize()
	g_is := createDateSet()
	assert.Equal(t, g_is.EndElement().Format("2006-01-02"), "2019-12-13")
	// // gomeos.MeosFinalize()
}

func TestDSElementN(t *testing.T) {
	// // gomeos.MeosInitialize()
	g_is := createDateSet()
	assert.Equal(t, g_is.ElementN(1).Format("2006-01-02"), "2019-09-10")
	// // gomeos.MeosFinalize()
}

func TestDSElements(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	dates := g_is.Elements()
	assert.Equal(t, dates[0].Format("2006-01-02"), "2019-09-08")
	assert.Equal(t, dates[1].Format("2006-01-02"), "2019-09-10")
	assert.Equal(t, dates[2].Format("2006-01-02"), "2019-12-13")
	// gomeos.MeosFinalize()
}

func TestDSShiftScale(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	// shift int pass
	g_ss, _ := g_is.ShiftScale(5, nil)
	dates := g_ss.Elements()
	assert.Equal(t, dates[0].Format("2006-01-02"), "2019-09-13")
	assert.Equal(t, dates[1].Format("2006-01-02"), "2019-09-15")
	assert.Equal(t, dates[2].Format("2006-01-02"), "2019-12-18")

	// shift timedelta pass
	g_ss2, _ := g_is.ShiftScale(timeutil.Timedelta{Days: 5}, nil)
	dates = g_ss2.Elements()
	assert.Equal(t, dates[0].Format("2006-01-02"), "2019-09-13")
	assert.Equal(t, dates[1].Format("2006-01-02"), "2019-09-15")
	assert.Equal(t, dates[2].Format("2006-01-02"), "2019-12-18")

	// shift int pass
	g_ss3, _ := g_is.ShiftScale(nil, 200)
	dates = g_ss3.Elements()
	assert.Equal(t, dates[0].Format("2006-01-02"), "2019-09-08")
	assert.Equal(t, dates[1].Format("2006-01-02"), "2019-09-12")
	assert.Equal(t, dates[2].Format("2006-01-02"), "2020-03-27")

	// shift timedelta pass
	g_ss4, _ := g_is.ShiftScale(nil, timeutil.Timedelta{Days: 200})
	dates = g_ss4.Elements()
	assert.Equal(t, dates[0].Format("2006-01-02"), "2019-09-08")
	assert.Equal(t, dates[1].Format("2006-01-02"), "2019-09-12")
	assert.Equal(t, dates[2].Format("2006-01-02"), "2020-03-27")

	// gomeos.MeosFinalize()
}

func TestDSShiftAndScale(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	// shift int pass
	g_ss, _ := g_is.Shift(5)
	dates := g_ss.Elements()
	assert.Equal(t, dates[0].Format("2006-01-02"), "2019-09-13")
	assert.Equal(t, dates[1].Format("2006-01-02"), "2019-09-15")
	assert.Equal(t, dates[2].Format("2006-01-02"), "2019-12-18")

	// shift timedelta pass
	g_ss2, _ := g_is.Shift(timeutil.Timedelta{Days: 5})
	dates = g_ss2.Elements()
	assert.Equal(t, dates[0].Format("2006-01-02"), "2019-09-13")
	assert.Equal(t, dates[1].Format("2006-01-02"), "2019-09-15")
	assert.Equal(t, dates[2].Format("2006-01-02"), "2019-12-18")

	// shift int pass
	g_ss3, _ := g_is.Scale(200)
	dates = g_ss3.Elements()
	assert.Equal(t, dates[0].Format("2006-01-02"), "2019-09-08")
	assert.Equal(t, dates[1].Format("2006-01-02"), "2019-09-12")
	assert.Equal(t, dates[2].Format("2006-01-02"), "2020-03-27")

	// shift timedelta pass
	g_ss4, _ := g_is.Scale(timeutil.Timedelta{Days: 200})
	dates = g_ss4.Elements()
	assert.Equal(t, dates[0].Format("2006-01-02"), "2019-09-08")
	assert.Equal(t, dates[1].Format("2006-01-02"), "2019-09-12")
	assert.Equal(t, dates[2].Format("2006-01-02"), "2020-03-27")

	// gomeos.MeosFinalize()
}

func TestDSContains(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	g_is_2 := createDateSet2()
	res, _ := g_is.Contains(g_is_2)
	assert.False(t, res)
	res, _ = g_is_2.Contains(g_is)
	assert.True(t, res)
	s1 := time.Date(2009, 11, 17, 0, 0, 0, 0, time.UTC)
	s2 := time.Date(2019, 9, 10, 0, 0, 0, 0, time.UTC)
	res, _ = g_is.Contains(s1)
	assert.False(t, res)
	res, _ = g_is.Contains(s2)
	assert.True(t, res)
	// gomeos.MeosFinalize()
}

func TestDSOverlaps(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	g_is_2 := createDateSet2()
	// Dateset
	res, _ := g_is.Overlaps(g_is_2)
	assert.True(t, res)
	// Date
	d := time.Date(2019, 9, 12, 0, 0, 0, 0, time.UTC)
	res, _ = g_is.Overlaps(d)
	assert.False(t, res)
	// DateSpan
	ds := createDateSpan()
	res, _ = g_is.Overlaps(ds)
	assert.False(t, res)
	// DateSpanSet
	dss := createDateSpanSet()
	res, _ = g_is.Overlaps(dss)
	assert.False(t, res)
	// gomeos.MeosFinalize()
}

func TestDSIsLeft(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	g_is_2 := createDateSet2()
	// Dateset
	res, _ := g_is.IsLeft(g_is_2)
	assert.True(t, res)
	// Date
	d := time.Date(2019, 9, 12, 0, 0, 0, 0, time.UTC)
	res, _ = g_is.IsLeft(d)
	assert.False(t, res)
	// DateSpan
	ds := createDateSpan()
	res, _ = g_is.IsLeft(ds)
	assert.False(t, res)
	// DateSpanSet
	dss := createDateSpanSet()
	res, _ = g_is.IsLeft(dss)
	assert.False(t, res)
	// gomeos.MeosFinalize()
}

func TestDSIsOverOrLeft(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	g_is_2 := createDateSet2()
	// Dateset
	res, _ := g_is.IsOverOrLeft(g_is_2)
	assert.True(t, res)
	// Date
	d := time.Date(2019, 9, 12, 0, 0, 0, 0, time.UTC)
	res, _ = g_is.IsOverOrLeft(d)
	assert.False(t, res)
	// DateSpan
	ds := createDateSpan()
	res, _ = g_is.IsOverOrLeft(ds)
	assert.False(t, res)
	// DateSpanSet
	dss := createDateSpanSet()
	res, _ = g_is.IsOverOrLeft(dss)
	assert.False(t, res)
	// gomeos.MeosFinalize()
}

func TestDSIsRight(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	g_is_2 := createDateSet2()
	// Dateset
	res, _ := g_is.IsRight(g_is_2)
	assert.False(t, res)
	// Date
	d := time.Date(2019, 9, 12, 0, 0, 0, 0, time.UTC)
	res, _ = g_is.IsRight(d)
	assert.False(t, res)
	// DateSpan
	ds := createDateSpan()
	res, _ = g_is.IsRight(ds)
	assert.False(t, res)
	// DateSpanSet
	dss := createDateSpanSet()
	res, _ = g_is.IsRight(dss)
	assert.False(t, res)
	// gomeos.MeosFinalize()
}

func TestDSIsOverOrRight(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	g_is_2 := createDateSet2()
	// Dateset
	res, _ := g_is.IsOverOrRight(g_is_2)
	assert.True(t, res)
	// Date
	d := time.Date(2019, 9, 12, 0, 0, 0, 0, time.UTC)
	res, _ = g_is.IsOverOrRight(d)
	assert.False(t, res)
	// DateSpan
	ds := createDateSpan()
	res, _ = g_is.IsOverOrRight(ds)
	assert.False(t, res)
	// DateSpanSet
	dss := createDateSpanSet()
	res, _ = g_is.IsOverOrRight(dss)
	assert.True(t, res)
	// gomeos.MeosFinalize()
}

func TestDSDistance(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	g_is_2 := createDateSet2()
	// Dateset
	res, _ := g_is.Distance(g_is_2)
	assert.Equal(t, res, timeutil.Timedelta{})
	// Date
	d := time.Date(2019, 12, 15, 0, 0, 0, 0, time.UTC)
	res, _ = g_is.Distance(d)
	assert.Equal(t, res, timeutil.Timedelta{Days: 2})
	// DateSpan
	ds := createDateSpan()
	res, _ = g_is.Distance(ds)
	assert.Equal(t, res, timeutil.Timedelta{})
	// DateSpanSet
	dss := createDateSpanSet()
	res, _ = g_is.Distance(dss)
	assert.Equal(t, res, timeutil.Timedelta{})
	// gomeos.MeosFinalize()
}

// func TestDSIntersection(t *testing.T) {
// 	gomeos.MeosInitialize("UTC")
// 	g_is := createDateSet()
// 	g_is_2 := createDateSet2()
// 	// Dateset
// 	res, _ := g_is.Intersection(g_is_2)
// 	assert.Equal(t, res.Output(), "{2019-09-08, 2019-09-10, 2019-12-13}")
// 	// Date
// 	d := time.Date(2019, 12, 15, 0, 0, 0, 0, time.UTC)
// 	res, _ = g_is.Intersection(d)
// 	assert.Equal(t, res.Output(), "Nil")
// 	// DateSpan
// 	ds := createDateSpan()
// 	res, _ = g_is.Intersection(ds)
// 	assert.Equal(t, res.Output(), "Nil")
// 	// DateSpanSet
// 	dss := createDateSpanSet()
// 	res, _ = g_is.Intersection(dss)
// 	assert.Equal(t, res.Output(), "{[2019-09-08, 2019-09-09), [2019-09-10, 2019-09-11)}")
// 	// gomeos.MeosFinalize()
// }

// func TestDSIntersectionInterface(t *testing.T) {
// 	gomeos.MeosInitialize("UTC")
// 	g_is := createDateSet()
// 	g_is_2 := createDateSet2()
// 	// Dateset
// 	res, _ := g_is.Intersection(g_is_2)
// 	if result, ok := res.(type); !ok {
// 		fmt.Errorf("not converted to concrete type!")
// 		assert.Equal(t, result.DateSetOut(), "{2019-09-08, 2019-09-10, 2019-12-13}")
// 	}
// // Date
// d := time.Date(2019, 12, 15, 0, 0, 0, 0, time.UTC)
// res, _ = g_is.Intersection(d)
// assert.Equal(t, res.Output(), "Nil")
// // DateSpan
// ds := createDateSpan()
// res, _ = g_is.Intersection(ds)
// assert.Equal(t, res.Output(), "Nil")
// // DateSpanSet
// dss := createDateSpanSet()
// res, _ = g_is.Intersection(dss)
// assert.Equal(t, res.Output(), "{[2019-09-08, 2019-09-09), [2019-09-10, 2019-09-11)}")
// // gomeos.MeosFinalize()
// }

func TestDSMinus(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	g_is_2 := createDateSet2()
	// Dateset
	res, _ := g_is.Minus(g_is_2)
	assert.Equal(t, res.Output(), "Nil")
	// Date
	d := time.Date(2019, 12, 15, 0, 0, 0, 0, time.UTC)
	res, _ = g_is.Minus(d)
	assert.Equal(t, res.Output(), "{2019-09-08, 2019-09-10, 2019-12-13}")
	// DateSpan
	ds := createDateSpan()
	res, _ = g_is.Minus(ds)
	assert.Equal(t, res.Output(), "{[2019-09-08, 2019-09-09), [2019-09-10, 2019-09-11), [2019-12-13, 2019-12-14)}")
	// DateSpanSet
	dss := createDateSpanSet()
	res, _ = g_is.Minus(dss)
	assert.Equal(t, res.Output(), "{[2019-12-13, 2019-12-14)}")
	// gomeos.MeosFinalize()
}

func TestDSUnion(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := createDateSet()
	g_is_2 := createDateSet2()
	// Dateset
	res, _ := g_is.Union(g_is_2)
	assert.Equal(t, res.Output(), "{2019-09-08, 2019-09-10, 2019-12-13, 2019-12-18}")
	// Date
	d := time.Date(2019, 12, 15, 0, 0, 0, 0, time.UTC)
	res, _ = g_is.Union(d)
	assert.Equal(t, res.Output(), "{2019-09-08, 2019-09-10, 2019-12-13, 2019-12-15}")
	// DateSpan
	ds := createDateSpan()
	res, _ = g_is.Union(ds)
	assert.Equal(t, res.Output(), "{[2019-09-08, 2019-09-11), [2019-12-13, 2019-12-14)}")
	// DateSpanSet
	dss := createDateSpanSet()
	res, _ = g_is.Union(dss)
	assert.Equal(t, res.Output(), "{[2019-09-08, 2019-09-13), [2019-12-13, 2019-12-14)}")
	// gomeos.MeosFinalize()
}
