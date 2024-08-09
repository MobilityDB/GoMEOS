package temporal_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
)

func createGeom() gomeos.Geom {
	return gomeos.NewGeom("SRID=99;POINT(0 0)", -1)
}

func TestCreateGeom(t *testing.T) {
	geom := createGeom()
	assert.Equal(t, geom.GeomOut(), "")
}
