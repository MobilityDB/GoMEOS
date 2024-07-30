package time_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos/collections/times"
	"github.com/MobilityDB/GoMEOS/gomeos/go_init"
	"github.com/alecthomas/assert/v2"
)

func createDateSpan() *times.DateSpan {
	return times.NewDateSpan("(2019-09-08, 2019-09-10)")
}

// Test function
func TestNewDateSpan(t *testing.T) {
	go_init.MeosInitialize()
	g_is := times.NewDateSpan("(2019-09-08, 2019-09-10)")
	if g_is == nil {
		t.Fatal("NewDateSpan returned null")
	}
	// defer g_is.Free() // Ensure memory is freed

	output := g_is.DateSpanOut()
	t.Logf("DateSpanOut returned: %s", output)
	assert.Equal(t, output, "[2019-09-09, 2019-09-10)")
	go_init.MeosFinalize()
}
