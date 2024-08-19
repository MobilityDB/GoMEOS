package temporal_test

import (
	"github.com/MobilityDB/GoMEOS/gomeos"
)

func createTInt() *gomeos.TIntInst {
	return gomeos.NewTIntInst("1@2000-01-01")
}
