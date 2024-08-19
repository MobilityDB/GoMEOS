package time_test

import (
	"testing"

	gomeos "github.com/MobilityDB/GoMEOS"
	"github.com/alecthomas/assert/v2"
)

func createDateSpan() *gomeos.DateSpan {
	return gomeos.NewDateSpan("(2019-09-08, 2019-09-10)")
}

// Test function
func TestNewDateSpan(t *testing.T) {
	// gomeos.MeosInitialize()
	g_is := gomeos.NewDateSpan("(2019-09-08, 2019-09-10)")
	if g_is == nil {
		t.Fatal("NewDateSpan returned null")
	}
	// defer g_is.Free() // Ensure memory is freed

	output := g_is.DateSpanOut()
	t.Logf("DateSpanOut returned: %s", output)
	assert.Equal(t, output, "[2019-09-09, 2019-09-10)")
	// gomeos.MeosFinalize()
}
