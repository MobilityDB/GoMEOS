package gomeos

import (
	"fmt"
)

func ExampleTPointAsGeoJson() {
	tf_seq := TGeomPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &TGeomPointSeq{})
	res := TPointAsGeoJson(tf_seq, 1, 15, "EPSG:4326")
	fmt.Println(res)
	// Output:
	// {"type":"MultiPoint","crs":{"type":"name","properties":{"name":"EPSG:4326"}},"bbox":[1.000000000000000,1.000000000000000,2.000000000000000,2.000000000000000],"coordinates":[[1,1],[2,2]]}
}

func ExampleTPointValueSet() {
	tf_seq := TGeomPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &TGeomPointSeq{})
	geoms, _ := TPointValueSet(tf_seq)
	fmt.Println(GeoAsText(geoms[0], 10))
	fmt.Println(GeoAsText(geoms[1], 10))
	// Output:
	// POINT(1 1)
	// POINT(2 2)
}

func ExampleTPointSTBoxes() {
	tf_seq := TGeomPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &TGeomPointSeq{})
	stboxes, _ := TPointSTBoxes(tf_seq, 100)
	for i := 0; i < len(stboxes); i++ {
		fmt.Println(STBoxOut(stboxes[i], 10))
	}
	// Output:
	// STBOX XT(((1,1),(1,1)),[2022-10-01 00:00:00+00, 2022-10-01 00:00:00+00])
	// STBOX XT(((2,2),(2,2)),[2022-10-02 00:00:00+00, 2022-10-02 00:00:00+00])
}

func ExampleTPointToSTBox() {
	tg := TGeomPointIn("{POINT(1 1)@2022-10-01 00:00:00+00, POINT(2 2)@2022-10-02 00:00:00+00}", &TGeomPointSeq{})
	stbox := TPointToSTBox(tg)
	fmt.Println(STBoxOut(stbox, 10))
	// Output:
	// STBOX XT(((1,1),(2,2)),[2022-10-01 00:00:00+00, 2022-10-02 00:00:00+00])
}
