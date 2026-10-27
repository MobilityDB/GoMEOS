package gomeos

// A string argument MEOS accepts as NULL is a *string in the generated
// package, and nil passes NULL. The MF-JSON of a temporal point names no
// coordinate reference system for a nil srs and names the one given otherwise.
// The value reaches the generated package as #ExampleTemporalFromPointer has it.

import (
	"fmt"
	"strings"

	"github.com/MobilityDB/GoMEOS/functions"
)

func ExampleTemporalAsMFJSON_nilSrs() {
	temp, err := functions.TgeompointIn("[Point(1 1)@2001-01-01, Point(2 2)@2001-01-02]")
	if err != nil {
		fmt.Println(err)
		return
	}
	none, err := functions.TemporalAsMFJSON(temp, 0, 0, 6, nil)
	fmt.Println(strings.Contains(none, `"crs"`), err)

	srs := "EPSG:4326"
	named, err := functions.TemporalAsMFJSON(temp, 0, 0, 6, &srs)
	fmt.Println(strings.Contains(named, `"EPSG:4326"`), err)
	// Output:
	// false <nil>
	// true <nil>
}
