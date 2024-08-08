package time_test

import (
	"testing"
	"time"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func TestDatetimeToTimestamptz(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	time := time.Now()
	pg_tstz := gomeos.DatetimeToTimestamptz(time)
	assert.Equal(t, gomeos.TimestamptzOut(pg_tstz), "")
}
