package time_test

import (
	"fmt"
	"testing"
	"time"

	gomeos "github.com/MobilityDB/GoMEOS"
	"github.com/alecthomas/assert/v2"
)

func TestUtil(t *testing.T) {
	// Example usage:
	timestampTz := gomeos.TimestamptzIn("2020-06-03 07:21:38.834+00")
	// timestampTz := int64(1620403200000000) // Example TimestampTz value
	fmt.Println("Original TimestampTz:", timestampTz)

	// Convert to time.Time
	datetime := gomeos.TimestamptzToDatetime(timestampTz)
	fmt.Println("Converted to time.Time:", datetime)

	// Convert back to TimestampTz
	timestampTzBack := gomeos.DatetimeToTimestamptz(datetime)
	fmt.Println("Converted back to TimestampTz:", timestampTzBack)
}

func TestDatetimeToTimestamptz(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	time := time.Now()
	pg_tstz := gomeos.DatetimeToTimestamptz(time)
	assert.Equal(t, gomeos.TimestamptzOut(pg_tstz), "")
}
