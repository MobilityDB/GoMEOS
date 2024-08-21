package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

// TPointOut Return a temporal geometry/geography point from its Well-Known Text (WKT) representation
func TPointOut[TP TPoint](tp TP, maxdd int) string {
	c_point := C.tpoint_out(tp.Inner(), C.int(maxdd))
	defer C.free(unsafe.Pointer(c_point))
	point_out := C.GoString(c_point)
	return point_out
}

// TPointAsText Return the Well-Known Text (WKT) representation of a temporal point
func TPointAsText[TP TPoint](tp TP, maxdd int) string {
	c_text := C.tpoint_as_text(tp.Inner(), C.int(maxdd))
	defer C.free(unsafe.Pointer(c_text))
	text_out := C.GoString(c_text)
	return text_out
}

// TPointAsEWKT Return the Extended Well-Known Text (EWKT) representation of a temporal point
func TPointAsEWKT[T TPoint](temp T, maxdd int) string {
	return C.GoString(C.tpoint_as_ewkt(temp.Inner(), C.int(maxdd)))
}

// TPointGeoAsEWKT Return the Extended Well-Known Text (EWKT) representation of the geometry of a temporal point
func TPointGeoAsEWKT[T TPoint](temp T, maxdd int) string {
	c_inst := C.cast_temporal_to_tinstant(temp.Inner())
	c_value := c_inst.value
	c_geo := (*C.GSERIALIZED)(unsafe.Pointer(&c_value))
	return C.GoString(C.geo_as_ewkt(c_geo, C.int(maxdd)))
}

// TPointAsGeoJson Return the GeoJson representation of the geometry of a temporal point
func TPointAsGeoJson[TP TPoint](temp TP, option int, precision int, srs string) string {
	traj := TPointTrajectory(temp)
	return GeoAsGeojson(traj, option, precision, srs)
}

// TPointToSTBox Return a temporal point converted to a spatiotemporal box
func TPointToSTBox[TP TPoint](tp TP) *STBox {
	st_box := C.tpoint_to_stbox(tp.Inner())
	return &STBox{_inner: st_box}
}

// TPointStartValue Return the start value of a temporal point
func TPointStartValue[TP TPoint](tp TP) *Geom {
	cValue := C.tpoint_start_value(tp.Inner())
	return &Geom{_inner: cValue}
}

// TPointEndValue Return the end value of a temporal point
func TPointEndValue[TP TPoint](tp TP) *Geom {
	cValue := C.tpoint_end_value(tp.Inner())
	return &Geom{_inner: cValue}
}

// TPointValueAtTimestamp Return the value of a temporal point at a timestamptz
func TPointValueAtTimestamp[TP TPoint](tp TP, ts time.Time) *Geom {
	tpointinst, _ := TemporalToGeomPointInst(TemporalAtTimestamptz(tp, ts))
	return TPointStartValue(tpointinst)
}

// TPointValueSet Return the array of base values of a temporal geometry point
func TPointValueSet[TP TPoint](tp TP) ([]*Geom, error) {
	var count C.int

	// Call the C function
	cValues := C.tpoint_values(tp.Inner(), &count)
	if cValues == nil {
		return nil, fmt.Errorf("failed to retrieve tpoint values")
	}

	// Convert the C array of pointers to a Go slice
	length := int(count)
	values := (*[1 << 28]*C.GSERIALIZED)(unsafe.Pointer(cValues))[:length:length]
	// Convert the C bool values to Go bool values
	goValues := make([]*Geom, length)
	for i := 0; i < length; i++ {
		goValues[i] = &Geom{_inner: values[i]}
	}
	return goValues, nil
}

// TPointLength Returns the length of the trajectory.
func TPointLength[TP TPoint](tp TP) float64 {
	return float64(C.tpoint_length(tp.Inner()))
}

// CumulativeLength Returns the cumulative length of the trajectory.
func CumulativeLength[TP TPoint](tp TP) Temporal {
	res := C.tpoint_cumulative_length(tp.Inner())
	return CreateTemporal(res)
}

// TPointSpeed Returns the speed of the temporal point.
func TPointSpeed[TP TPoint](tp TP) Temporal {
	res := C.tpoint_speed(tp.Inner())
	return CreateTemporal(res)
}

// TPointGetX Returns the x coordinate of the temporal point.
func TPointGetX[TP TPoint](tp TP) Temporal {
	res := C.tpoint_get_x(tp.Inner())
	return CreateTemporal(res)
}

// TPointGetY Returns the y coordinate of the temporal point.
func TPointGetY[TP TPoint](tp TP) Temporal {
	res := C.tpoint_get_y(tp.Inner())
	return CreateTemporal(res)
}

// TPointGetZ Returns the z coordinate of the temporal point.
func TPointGetZ[TP TPoint](tp TP) Temporal {
	res := C.tpoint_get_z(tp.Inner())
	return CreateTemporal(res)
}

// TPointSTBoxes Return an array of spatiotemporal boxes from the segments of a temporal point
func TPointSTBoxes[TP TPoint](tp TP, max_count int) ([]*STBox, error) {
	var count C.int

	// Call the C function
	cValues := C.tpoint_stboxes(tp.Inner(), C.int(max_count), &count)
	if cValues == nil {
		return nil, fmt.Errorf("failed to retrieve bool values")
	}

	// Convert the C bool array to a Go slice
	length := int(count)
	values := unsafe.Slice((*C.STBox)(cValues), length)
	// Convert the C bool values to Go bool values
	goValues := make([]*STBox, length)
	for i := 0; i < length; i++ {
		goValues[i] = &STBox{_inner: &values[i]}
	}
	return goValues, nil
}

func TPointTransform[T TPoint](temp T, output T, srid_to int) T {
	c_temp := C.tpoint_transform(temp.Inner(), C.int(srid_to))
	output.Init(c_temp)
	return output
}

func TPointTrajectory[TP TPoint](tp TP) *Geom {
	trajectory := C.tpoint_trajectory(tp.Inner())
	return &Geom{_inner: trajectory}
}

//func TpointAtGeomTime[T Temporal](temp T, new_temp T, geom *Geom) T {
//	c_temp := C.tpoint_at_geom_time(temp.Inner(), geom._inner, nil, nil)
//	new_temp.Init(c_temp)
//	return new_temp
//}
//
//func TpointMinusGeomTime[T Temporal](temp T, new_temp T, geom *Geom) T {
//	c_temp := C.tpoint_minus_geom_time(temp.Inner(), geom._inner, nil, nil)
//	new_temp.Init(c_temp)
//	return new_temp
//}
