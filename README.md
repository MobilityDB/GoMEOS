# GoMEOS

[MEOS (Mobility Engine, Open Source)](https://www.libmeos.org/) is a C library which enables the manipulation of
temporal and spatio-temporal data based on [MobilityDB](https://mobilitydb.com/)'s data types and functions.

GoMEOS is a Go library that wraps the MEOS C library using [CGO](https://pkg.go.dev/cmd/cgo), providing a set of Go functions that allows to use MEOS functionality by directly accessing C structs and C functions.

GoMEOS exposes the functionality of MEOS and is meant to be used directly by the user.

# Usage

## Installation

You'll need to install the MEOS library by compling MobilityDB.
```shell
git clone git@github.com:MobilityDB/MobilityDB.git
cd MobilityDB
mkdir build
cd build
cmake -DMEOS=on ..
make
sudo make install
```
GoMEOS requires [CGO](https://pkg.go.dev/cmd/cgo) (CGO_ENABLED=1) in order to be built.

Then install GoMEOS by `go get` command.
```shell
go get github.com/MobilityDB/GoMEOS
```

# Quick Start
> **IMPORTANT** Before using any GoMEOS function, always call `MeosInitialize()`. Otherwise, the library will crash with a `Segmentation Fault` error. You should also always call `MeosFinalize()` at the end of your code.

```go
// main.go
package main

import (
	"fmt"

	gomeos "github.com/MobilityDB/GoMEOS"
)

func main() {
	gomeos.MeosInitialize("UTC")

	/* Input temporal points in WKT format */
	inst_wkt := "POINT(1 1)@2000-01-01"
	seq_disc_wkt := "{POINT(1 1)@2000-01-01, POINT(2 2)@2000-01-02}"
	seq_linear_wkt := "[POINT(1 1)@2000-01-01, POINT(2 2)@2000-01-02]"
	seq_step_wkt := "Interp=Step;[POINT(1 1)@2000-01-01, POINT(2 2)@2000-01-02]"
	ss_linear_wkt := "{[POINT(1 1)@2000-01-01, POINT(2 2)@2000-01-02],[POINT(3 3)@2000-01-03, POINT(3 3)@2000-01-04]}"
	ss_step_wkt := "Interp=Step;{[POINT(1 1)@2000-01-01, POINT(2 2)@2000-01-02],[POINT(3 3)@2000-01-03, POINT(3 3)@2000-01-04]}"

	/* Read WKT into temporal point object */
	inst := gomeos.NewTGeomPointInst(inst_wkt)
	seq_disc := gomeos.NewTGeomPointSeq(seq_disc_wkt)
	seq_linear := gomeos.NewTGeomPointSeq(seq_linear_wkt)
	seq_step := gomeos.NewTGeomPointSeq(seq_step_wkt)
	ss_linear := gomeos.NewTGeomPointSeqSet(ss_linear_wkt)
	ss_step := gomeos.NewTGeomPointSeqSet(ss_step_wkt)

	fmt.Println("--------------------Temporal Instant--------------------")
	fmt.Println(inst.TPointOut(5))
	fmt.Println("--------------------Temporal Sequence with Discrete Interpolation--------------------")
	fmt.Println(seq_disc.TPointOut(5))
	fmt.Println("--------------------Temporal Sequence with Linear Interpolation--------------------")
	fmt.Println(seq_linear.TPointOut(5))
	fmt.Println("--------------------Temporal Sequence with Step Interpolation--------------------")
	fmt.Println(seq_step.TPointOut(5))
	fmt.Println("--------------------Temporal Sequence Set with Linear Interpolation--------------------")
	fmt.Println(ss_linear.TPointOut(5))
	fmt.Println("--------------------Temporal Sequence Set with Step Interpolation--------------------")
	fmt.Println(ss_step.TPointOut(5))
	gomeos.MeosFinalize()
}

```
To run this file:

```shell
go build main.go
./main
```

```
output:
--------------------Temporal Instant--------------------
POINT(1 1)@2000-01-01 00:00:00+00
--------------------Temporal Sequence with Discrete Interpolation--------------------
{POINT(1 1)@2000-01-01 00:00:00+00, POINT(2 2)@2000-01-02 00:00:00+00}
--------------------Temporal Sequence with Linear Interpolation--------------------
[POINT(1 1)@2000-01-01 00:00:00+00, POINT(2 2)@2000-01-02 00:00:00+00]
--------------------Temporal Sequence with Step Interpolation--------------------
Interp=Step;[POINT(1 1)@2000-01-01 00:00:00+00, POINT(2 2)@2000-01-02 00:00:00+00]
--------------------Temporal Sequence Set with Linear Interpolation--------------------
{[POINT(1 1)@2000-01-01 00:00:00+00, POINT(2 2)@2000-01-02 00:00:00+00], [POINT(3 3)@2000-01-03 00:00:00+00, POINT(3 3)@2000-01-04 00:00:00+00]}
--------------------Temporal Sequence Set with Step Interpolation--------------------
Interp=Step;{[POINT(1 1)@2000-01-01 00:00:00+00, POINT(2 2)@2000-01-02 00:00:00+00], [POINT(3 3)@2000-01-03 00:00:00+00, POINT(3 3)@2000-01-04 00:00:00+00]}
```

For more examples, you can see the [examples](https://github.com/MobilityDB/GoMEOS/tree/main/examples).

```shell
// Run ais.go example
git clone git@github.com:MobilityDB/GoMEOS.git
cd GoMEOS
go run examples/ais/assemble/ais.go
```

# Documentation
Visit [pkg.go.dev](https://pkg.go.dev/github.com/MobilityDB/GoMEOS) for a more complete and detailed documentation.