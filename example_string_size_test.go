package gomeos

// A string MEOS answers comes with its size, written through a size_t *, so the
// generated wrapper takes no such argument and answers the string alone, as
// #ExampleTemporalAsMFJSON_nilSrs has it. The size counts bytes, so a value holding
// a character UTF-8 writes in two bytes reads back whole, and a function whose only
// parameter is that size takes no argument at all.

import (
	"fmt"

	"github.com/MobilityDB/GoMEOS/functions"
)

func ExampleTemporalOut_size() {
	temp, err := functions.TtextIn(`{"Zürich"@2001-01-01, "Liège"@2001-01-02}`)
	if err != nil {
		fmt.Println(err)
		return
	}
	text, err := functions.TemporalOut(temp, 0)
	fmt.Println(text, err)

	version, err := functions.MeosVersion()
	fmt.Println(version != "", err)
	// Output:
	// {"Zürich"@2001-01-01 00:00:00+00, "Liège"@2001-01-02 00:00:00+00} <nil>
	// true <nil>
}
