package temporal_test

import (
	"testing"

	gomeos "github.com/MobilityDB/GoMEOS"
	"github.com/alecthomas/assert/v2"
)

func TestTSequenceInc(t *testing.T) {
	gomeos.MeosInitialize("UTC")
	tt := gomeos.TTextIn("{hello@2022-10-01 00:00:00+00, world@2022-10-02 00:00:00+00}", &gomeos.TTextSeq{})
	assert.Equal(t, true, gomeos.TSequenceLowerInclude(tt))
}
