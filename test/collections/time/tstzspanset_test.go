package time_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos/collections/times"
	"github.com/MobilityDB/GoMEOS/gomeos/go_init"
	"github.com/alecthomas/assert/v2"
)

func createTsTzSpanSet() *times.TsTzSpanSet {
	return times.NewTsTzSpanSet("{[2023-01-01 00:00:00+00, 2023-01-31 23:59:59+00], [2023-02-01 00:00:00+00, 2023-02-28 23:59:59+00], [2023-03-01 00:00:00+00, 2023-03-31 23:59:59+00]}")
}

func TestNewTsTzSpanSet(t *testing.T) {
	go_init.MeosInitialize()
	g_dss := createTsTzSpanSet()
	assert.Equal(t, g_dss.TsTzSpanSetOut(), "{[2023-01-01 02:00:00+02, 2023-02-01 01:59:59+02], [2023-02-01 02:00:00+02, 2023-03-01 01:59:59+02], [2023-03-01 02:00:00+02, 2023-04-01 01:59:59+02]}")
	go_init.MeosFinalize()
}
