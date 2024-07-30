package time_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos/collections/times"
	"github.com/MobilityDB/GoMEOS/gomeos/go_init"
	"github.com/alecthomas/assert/v2"
)

func createTsTzSet() *times.TsTzSet {
	return times.NewTsTzSet("{2023-01-01 08:09:21+00, 2023-02-01 23:45:52+00, 2023-03-01 23:43:46+00}")
}

func TestNewTsTzSet(t *testing.T) {
	go_init.MeosInitialize()
	g_dss := createTsTzSet()
	assert.Equal(t, g_dss.TsTzSetOut(), "{\"2023-01-01 10:09:21+02\", \"2023-02-02 01:45:52+02\", \"2023-03-02 01:43:46+02\"}")
	go_init.MeosFinalize()
}
