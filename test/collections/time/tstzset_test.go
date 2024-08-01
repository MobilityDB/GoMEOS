package time_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createTsTzSet() *gomeos.TsTzSet {
	return gomeos.NewTsTzSet("{2023-01-01 08:09:21+00, 2023-02-01 23:45:52+00, 2023-03-01 23:43:46+00}")
}

func TestNewTsTzSet(t *testing.T) {
	gomeos.MeosInitialize()
	g_dss := createTsTzSet()
	assert.Equal(t, g_dss.TsTzSetOut(), "{\"2023-01-01 10:09:21+02\", \"2023-02-02 01:45:52+02\", \"2023-03-02 01:43:46+02\"}")
	gomeos.MeosFinalize()
}
