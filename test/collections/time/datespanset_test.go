package time_test

import (
	"testing"
	"time"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createDateSpanSet() *gomeos.DateSpanSet {
	return gomeos.NewDateSpanSet("{[2019-09-08, 2019-09-10], [2019-09-11, 2019-09-12]}")
}

func TestNewDateSpanSet(t *testing.T) {
	gomeos.MeosInitialize()
	g_dss := createDateSpanSet()
	assert.Equal(t, g_dss.DateSpanSetOut(), "{[2019-09-08, 2019-09-13)}")
	gomeos.MeosFinalize()
}

func TestDSSUnion(t *testing.T) {
	gomeos.MeosInitialize()
	g_iss := createDateSpanSet()
	g_is_2 := createDateSet2()
	// Dateset
	res, _ := g_iss.Union(g_is_2)
	assert.Equal(t, res.Output(), "Nil")
	// Date
	d := time.Date(2019, 12, 15, 0, 0, 0, 0, time.UTC)
	res, _ = g_iss.Union(d)
	assert.Equal(t, res.Output(), "{[2019-09-08, 2019-09-13), [2019-12-15, 2019-12-16)}")
	// DateSpan
	ds := createDateSpan()
	res, _ = g_iss.Union(ds)
	assert.Equal(t, res.Output(), "{[2019-09-08, 2019-09-13)}")
	// DateSpanSet
	dss := createDateSpanSet()
	res, _ = g_iss.Union(dss)
	assert.Equal(t, res.Output(), "{[2019-09-08, 2019-09-13)}")
	gomeos.MeosFinalize()
}
