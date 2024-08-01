package time_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createTsTzSpan() *gomeos.TsTzSpan {
	return gomeos.NewTsTzSpan("[2023-01-01 00:00:00+00, 2023-03-01 00:00:00+00]")
}

func TestNewTsTzSpan(t *testing.T) {
	gomeos.MeosInitialize()
	g_dss := createTsTzSpan()
	assert.Equal(t, g_dss.TsTzSpanOut(), "[2023-01-01 02:00:00+02, 2023-03-01 02:00:00+02]")
	gomeos.MeosFinalize()
}
