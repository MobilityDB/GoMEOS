package time_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos/collections/times"
	"github.com/MobilityDB/GoMEOS/gomeos/go_init"
	"github.com/alecthomas/assert/v2"
)

func createTsTzSpan() *times.TsTzSpan {
	return times.NewTsTzSpan("[2023-01-01 00:00:00+00, 2023-03-01 00:00:00+00]")
}

func TestNewTsTzSpan(t *testing.T) {
	go_init.MeosInitialize()
	g_dss := createTsTzSpan()
	assert.Equal(t, g_dss.TsTzSpanOut(), "[2023-01-01 02:00:00+02, 2023-03-01 02:00:00+02]")
	go_init.MeosFinalize()
}
