package time_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos/collections/times"
	"github.com/MobilityDB/GoMEOS/gomeos/go_init"
	"github.com/alecthomas/assert/v2"
)

func createDateSpanSet() *times.DateSpanSet {
	return times.NewDateSpanSet("{[2019-09-08, 2019-09-10], [2019-09-11, 2019-09-12]}")
}

func TestNewDateSpanSet(t *testing.T) {
	go_init.MeosInitialize()
	g_dss := createDateSpanSet()
	assert.Equal(t, g_dss.DateSpanSetOut(), "{[2019-09-08, 2019-09-13)}")
	go_init.MeosFinalize()
}
