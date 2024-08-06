package time_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createTsTzSpanSet() *gomeos.TsTzSpanSet {
	return gomeos.NewTsTzSpanSet("{[2023-01-01 00:00:00+00, 2023-01-31 23:59:59+00], [2023-02-01 00:00:00+00, 2023-02-28 23:59:59+00], [2023-03-01 00:00:00+00, 2023-03-31 23:59:59+00]}")
}

func TestNewTsTzSpanSet(t *testing.T) {
	// gomeos.MeosInitialize()
	g_dss := createTsTzSpanSet()
	assert.Equal(t, g_dss.TsTzSpanSetOut(), "{[2023-01-01 02:00:00+02, 2023-02-01 01:59:59+02], [2023-02-01 02:00:00+02, 2023-03-01 01:59:59+02], [2023-03-01 02:00:00+02, 2023-04-01 01:59:59+02]}")
	// gomeos.MeosFinalize()
}
