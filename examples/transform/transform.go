package main

//MEOS example: meos/examples/ais_transform.c
import (
	"fmt"

	gomeos "github.com/MobilityDB/GoMEOS"
)

func main() {
	gomeos.MeosInitialize("UTC")
	trip := "SRID=4326;[POINT(6.369747 55.209627)@2023-01-08 18:21:44+01,POINT(6.368953 55.210777)@2023-01-08 18:22:55+01,POINT(6.368603 55.211165)@2023-01-08 18:23:15+01,POINT(6.367535 55.212192)@2023-01-08 18:23:44+01,POINT(6.36672 55.21301)@2023-01-08 18:24:05+01,POINT(6.355373 55.22781)@2023-01-08 18:30:15+01,POINT(6.35513 55.228128)@2023-01-08 18:30:25+01,POINT(6.352437 55.23207)@2023-01-08 18:32:54+01,POINT(6.352097 55.232548)@2023-01-08 18:33:15+01,POINT(6.351793 55.23299)@2023-01-08 18:33:34+01]"
	seq := gomeos.TGeomPointIn(trip, &gomeos.TGeomPointSeq{})
	output := gomeos.TPointAsEWKT(seq, 6)
	fmt.Printf("----------------------------\n")
	fmt.Printf(" Original trip in SRID 4326\n")
	fmt.Printf("----------------------------\n%s\n", output)

	transformed_seq := gomeos.TPointTransform(seq, &gomeos.TGeomPointSeq{}, 25832)
	transformed_output := gomeos.TPointAsEWKT(transformed_seq, 6)
	fmt.Printf("----------------------------\n")
	fmt.Printf(" Original trip in SRID 4326\n")
	fmt.Printf("----------------------------\n%s\n", transformed_output)
}
