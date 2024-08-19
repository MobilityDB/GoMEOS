# GoMEOS

[MEOS (Mobility Engine, Open Source)](https://www.libmeos.org/) is a C library which enables the manipulation of
temporal and spatio-temporal data based on [MobilityDB](https://mobilitydb.com/)'s data types and functions.

GoMEOS is a Go library that wraps the MEOS C library using [CGO](https://pkg.go.dev/cmd/cgo), providing a set of Go functions that allows to use MEOS functionality by directly accessing C structs and C functions.

GoMEOS exposes the functionality of MEOS and is meant to be used directly by the user.

# Usage

## Installation
```shell
go get github.com/MobilityDB/GoMEOS
```
You also need to install the MEOS library by compling MobilityDB.
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

## Path Configuration
To give CGO the access to MEOS library, flags(which is paths to both the dynamic library and header files) need to be configured in `gomeos.go`.

In MacOS, the flag setting can be:
```shell
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib
```

In Linux, the flag setting can be:
```shell
#cgo CFLAGS: -I/usr/local/include/
#cgo LDFLAGS: -L/usr/local/lib -lmeos -Wl,-rpath,/usr/local/lib
```
Note that the directory may be different after compling and installing on different systems. So it is highly recommended to check the directory for both dynamic library and header files.

# Quick Start
> **IMPORTANT** Before using any GoMEOS function, always call `MeosInitialize()`. Otherwise, the library will crash with a `Segmentation Fault` error. You should also always call `MeosFinalize()` at the end of your code.

```go
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

# Documentation
Visit [pkg.go.dev](https://pkg.go.dev/github.com/MobilityDB/GoMEOS) for a more complete and detailed documentation.