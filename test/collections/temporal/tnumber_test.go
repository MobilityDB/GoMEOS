package temporal_test

import (
	gomeos "github.com/MobilityDB/GoMEOS"
)

func createTInt() *gomeos.TIntInst {
	return gomeos.NewTIntInst("1@2000-01-01")
}
