package time_test

import (
	"testing"

	"github.com/MobilityDB/GoMEOS/gomeos"
	"github.com/alecthomas/assert/v2"
	"github.com/leekchan/timeutil"
)

func createTsTzSet() *gomeos.TsTzSet {
	return gomeos.NewTsTzSet("{2023-01-01 08:09:21+00, 2023-02-01 23:45:52+00, 2023-03-01 08:09:21+00}")
}

func TestNewTsTzSet(t *testing.T) {
	gomeos.MeosInitialize()
	g_dss := createTsTzSet()
	assert.Equal(t, g_dss.TsTzSetOut(), "{\"2023-01-01 10:09:21+02\", \"2023-02-02 01:45:52+02\", \"2023-03-01 10:09:21+02\"}")
	gomeos.MeosFinalize()
}

func TestTTSDuration(t *testing.T) {
	gomeos.MeosInitialize()
	g_dss := createTsTzSet()
	assert.Equal(t, g_dss.Duration(), timeutil.Timedelta{Days: 59})
	gomeos.MeosFinalize()
}

func TestTSStartElement(t *testing.T) {
	gomeos.MeosInitialize()
	g_is := createTsTzSet()
	assert.Equal(t, g_is.StartElement().Format("2006-01-02 15:04:05-07"), "2023-01-01 10:09:21+02")
	gomeos.MeosFinalize()
}

func TestTSEndElement(t *testing.T) {
	gomeos.MeosInitialize()
	g_is := createTsTzSet()
	assert.Equal(t, g_is.EndElement().Format("2006-01-02 15:04:05-07"), "2023-03-01 10:09:21+02")
	gomeos.MeosFinalize()
}

func TestTSElementN(t *testing.T) {
	gomeos.MeosInitialize()
	g_is := createTsTzSet()
	assert.Equal(t, g_is.ElementN(1).Format("2006-01-02 15:04:05-07"), "2023-02-02 01:45:52+02")
	gomeos.MeosFinalize()
}

func TestTSSElements(t *testing.T) {
	gomeos.MeosInitialize()
	g_is := createTsTzSet()
	dates := g_is.Elements()
	assert.Equal(t, dates[0].Format("2006-01-02 15:04:05-07"), "2023-01-01 10:09:21+02")
	assert.Equal(t, dates[1].Format("2006-01-02 15:04:05-07"), "2023-02-02 01:45:52+02")
	assert.Equal(t, dates[2].Format("2006-01-02 15:04:05-07"), "2023-03-01 10:09:21+02")
	gomeos.MeosFinalize()
}
