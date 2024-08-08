package main

import (
	"fmt"

	"github.com/MobilityDB/GoMEOS/gomeos"
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

}
