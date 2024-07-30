package time_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos/go_init"
	"github.com/MobilityDB/GoMEOS/gomeos/main_t"
	"github.com/alecthomas/assert/v2"
)

func createTFloat() *main_t.TFloat {
	return main_t.NewTFloat("[1@2020-03-01, 10@2020-03-10]")
}

func TestNewTFloat(t *testing.T) {
	go_init.MeosInitialize()
	g_is := createTFloat()
	assert.Equal(t, g_is.TFloatOut(5), "[1@2020-03-01 00:00:00+01, 10@2020-03-10 00:00:00+01]")
	go_init.MeosFinalize()
}
