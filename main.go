package main

import (
	"fmt"

	"github.com/MobilityDB/GoMEOS/gomeos/collections/time"
)

func createDateSet() *time.DateSet {
	return time.NewDateSet("{2019-09-08, 2019-09-10, 2019-09-11}")
}

func main() {
	g_is := createDateSet()
	fmt.Println(g_is.DateSetOut())
}
