package number_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos/collections/number"
	"github.com/alecthomas/assert/v2"
)

func createIntSet() *number.IntSet {
	g_is_in := "{1,2,3}"
	g_is := number.NewIntSet(g_is_in)
	return g_is
}

func TestISInOutput(t *testing.T) {
	g_is := createIntSet()
	assert.Equal(t, g_is.IntSetOut(), "{1, 2, 3}")
}

func TestISToSpanSet(t *testing.T) {
	g_is := createIntSet()
	assert.Equal(t, g_is.ToSpanSet().IntSpanSetOut(), "{[1, 4)}")
}
