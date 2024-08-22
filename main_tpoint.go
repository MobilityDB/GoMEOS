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

// TPointIsSimple Return true if a temporal point does not self-intersect
func TPointIsSimple[TP TPoint](tp TP) bool {
	return bool(C.tpoint_is_simple(tp.Inner()))
}

// BearingTPointPoint Return the temporal bearing between a temporal point and a point
func BearingTPointPoint[TP TPoint](tp TP, g *Geom, invert bool) Temporal {
	c_temp := C.bearing_tpoint_point(tp.Inner(), g._inner, C.bool(invert))
	return CreateTemporal(c_temp)
}

// BearingTPointTPoint Return the temporal bearing between two temporal points
func BearingTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) Temporal {
	c_temp := C.bearing_tpoint_tpoint(tp1.Inner(), tp2.Inner())
	return CreateTemporal(c_temp)
}

// TPointAzimuth Return the temporal azimuth of a temporal geometry point
func TPointAzimuth[TP TPoint](tp TP) Temporal {
	res := C.tpoint_azimuth(tp.Inner())
	return CreateTemporal(res)
}

// TPointAngularDifference Return the temporal angular difference of a temporal geometry point
func TPointAngularDifference[TP TPoint](tp TP) Temporal {
	res := C.tpoint_angular_difference(tp.Inner())
	return CreateTemporal(res)
}

// TPointTwcentroid Return the time-weighed centroid of a temporal geometry point
func TPointTwcentroid[TP TPoint](tp TP) *Geom {
	res := C.tpoint_twcentroid(tp.Inner())
	return &Geom{_inner: res}
}

// TPointSrid Return the SRID of a temporal point
func TPointSrid[TP TPoint](tp TP) int {
	return int(C.tpoint_srid(tp.Inner()))
}

// TPointSetSrid Return a temporal point with the coordinates set to an SRID
func TPointSetSrid[TP TPoint](tp TP, srid int, output TP) TP {
	c_temp := C.tpoint_set_srid(tp.Inner(), C.int(srid))
	output.Init(c_temp)
	return output
}

// TPointRound Return a temporal point with the precision of the coordinates set to a number of decimal
func TPointRound[TP TPoint](tp TP, max_decimals int) Temporal {
	res := C.tpoint_round(tp.Inner(), C.int(max_decimals))
	return CreateTemporal(res)
}

// TPointExpandSpace Return the bounding box of a temporal point expanded on the spatial dimension
func TPointExpandSpace[TP TPoint](tp TP, other float64) *STBox {
	return &STBox{
		_inner: C.tpoint_expand_space(tp.Inner(), C.double(other)),
	}
}

// TPointAtValue Return a temporal point restricted to a point
func TPointAtValue[TP TPoint](tp TP, value *Geom) Temporal {
	c_temp := C.tpoint_at_value(tp.Inner(), value._inner)
	return CreateTemporal(c_temp)
}

// TPointAtGeomTime Return a temporal point restricted to a geometry
func TpointAtGeomTime[T Temporal](temp T, new_temp T, geom *Geom) T {
	c_temp := C.tpoint_at_geom_time(temp.Inner(), geom._inner, nil, nil)
	new_temp.Init(c_temp)
	return new_temp
}

// TPointAtStbox Return a temporal point restricted to a spatiotemporal box
func TPointAtStbox[TP TPoint](tp TP, stbox *STBox, border_inc bool) Temporal {
	res := C.tpoint_at_stbox(tp.Inner(), stbox._inner, C.bool(border_inc))
	return CreateTemporal(res)
}

// TPointMinusValue Return a temporal point minus a point
func TPointMinusValue[TP TPoint](tp TP, value *Geom) Temporal {
	c_temp := C.tpoint_minus_value(tp.Inner(), value._inner)
	return CreateTemporal(c_temp)
}

// TPointMinusGeomTime Return a temporal point minus a geometry
func TpointMinusGeomTime[T Temporal](temp T, new_temp T, geom *Geom) T {
	c_temp := C.tpoint_minus_geom_time(temp.Inner(), geom._inner, nil, nil)
	new_temp.Init(c_temp)
	return new_temp
}

// TPointMinusStbox Return a temporal point minus a spatiotemporal box
func TPointMinusStbox[TP TPoint](tp TP, stbox *STBox, border_inc bool) Temporal {
	res := C.tpoint_minus_stbox(tp.Inner(), stbox._inner, C.bool(border_inc))
	return CreateTemporal(res)
}

// LeftTPointTPoint Return true if a temporal point is to the left of a spatiotemporal box
func LeftTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.left_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// LeftTPointSTBox Return true if the first temporal point is to the left of the second one
func LeftTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.left_tpoint_stbox(tp.Inner(), stbox._inner))
}

// OverleftTPointSTBox returns true if a temporal point is overleft of a spatiotemporal box.
func OverleftTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.overleft_tpoint_stbox(tp.Inner(), stbox._inner))
}

// OverleftTPointTPoint returns true if the first temporal point is overleft of the second temporal point.
func OverleftTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.overleft_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// RightTPointSTBox returns true if a temporal point is to the right of a spatiotemporal box.
func RightTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.right_tpoint_stbox(tp.Inner(), stbox._inner))
}

// RightTPointTPoint returns true if the first temporal point is to the right of the second temporal point.
func RightTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.right_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// OverrightTPointSTBox returns true if a temporal point is overright of a spatiotemporal box.
func OverrightTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.overright_tpoint_stbox(tp.Inner(), stbox._inner))
}

// OverrightTPointTPoint returns true if the first temporal point is overright of the second temporal point.
func OverrightTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.overright_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// BelowTPointSTBox returns true if a temporal point is below a spatiotemporal box.
func BelowTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.below_tpoint_stbox(tp.Inner(), stbox._inner))
}

// BelowTPointTPoint returns true if the first temporal point is below the second temporal point.
func BelowTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.below_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// OverbelowTPointSTBox returns true if a temporal point is overbelow a spatiotemporal box.
func OverbelowTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.overbelow_tpoint_stbox(tp.Inner(), stbox._inner))
}

// OverbelowTPointTPoint returns true if the first temporal point is overbelow the second temporal point.
func OverbelowTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.overbelow_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// AboveTPointSTBox returns true if a temporal point is above a spatiotemporal box.
func AboveTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.above_tpoint_stbox(tp.Inner(), stbox._inner))
}

// AboveTPointTPoint returns true if the first temporal point is above the second temporal point.
func AboveTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.above_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// OveraboveTPointSTBox returns true if a temporal point is overabove a spatiotemporal box.
func OveraboveTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.overabove_tpoint_stbox(tp.Inner(), stbox._inner))
}

// OveraboveTPointTPoint returns true if the first temporal point is overabove the second temporal point.
func OveraboveTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.overabove_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// FrontTPointSTBox returns true if a temporal point is in front of a spatiotemporal box.
func FrontTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.front_tpoint_stbox(tp.Inner(), stbox._inner))
}

// FrontTPointTPoint returns true if the first temporal point is in front of the second temporal point.
func FrontTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.front_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// OverfrontTPointSTBox returns true if a temporal point is overfront of a spatiotemporal box.
func OverfrontTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.overfront_tpoint_stbox(tp.Inner(), stbox._inner))
}

// OverfrontTPointTPoint returns true if the first temporal point is overfront of the second temporal point.
func OverfrontTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.overfront_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// BackTPointSTBox returns true if a temporal point is behind a spatiotemporal box.
func BackTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.back_tpoint_stbox(tp.Inner(), stbox._inner))
}

// BackTPointTPoint returns true if the first temporal point is behind the second temporal point.
func BackTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.back_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// OverbackTPointSTBox returns true if a temporal point is overback of a spatiotemporal box.
func OverbackTPointSTBox[TP TPoint](tp TP, stbox *STBox) bool {
	return bool(C.overback_tpoint_stbox(tp.Inner(), stbox._inner))
}

// OverbackTPointTPoint returns true if the first temporal point is overback of the second temporal point.
func OverbackTPointTPoint[TP1 TPoint, TP2 TPoint](tp1 TP1, tp2 TP2) bool {
	return bool(C.overback_tpoint_tpoint(tp1.Inner(), tp2.Inner()))
}

// AlwaysContainsGeoTPoint returns true if the geometry contains the temporal point.
func AlwaysContainsGeoTPoint(gs *Geom, temp Temporal) bool {
	return int(C.acontains_geo_tpoint(gs._inner, temp.Inner())) > 0
}

// AlwaysDisjointTPointGeo returns true if the temporal point is disjoint from the geometry.
func AlwaysDisjointTPointGeo(temp Temporal, gs *Geom) bool {
	return int(C.adisjoint_tpoint_geo(temp.Inner(), gs._inner)) > 0
}

// AlwaysDisjointTPointTPoint returns true if the two temporal points are disjoint.
func AlwaysDisjointTPointTPoint(temp1, temp2 Temporal) bool {
	return int(C.adisjoint_tpoint_tpoint(temp1.Inner(), temp2.Inner())) > 0
}

// AlwaysDWithinTPointGeo returns true if the temporal point is within the specified distance of the geometry.
func AlwaysDWithinTPointGeo(temp Temporal, gs *Geom, dist float64) bool {
	return int(C.adwithin_tpoint_geo(temp.Inner(), gs._inner, C.double(dist))) > 0
}

// AlwaysDWithinTPointTPoint returns true if the two temporal points are within the specified distance.
func AlwaysDWithinTPointTPoint(temp1, temp2 Temporal, dist float64) bool {
	return int(C.adwithin_tpoint_tpoint(temp1.Inner(), temp2.Inner(), C.double(dist))) > 0
}

// AlwaysIntersectsTPointGeo returns true if the temporal point intersects the geometry.
func AlwaysIntersectsTPointGeo(temp Temporal, gs *Geom) bool {
	return int(C.aintersects_tpoint_geo(temp.Inner(), gs._inner)) > 0
}

// AlwaysIntersectsTPointTPoint returns true if the two temporal points intersect.
func AlwaysIntersectsTPointTPoint(temp1, temp2 Temporal) bool {
	return int(C.aintersects_tpoint_tpoint(temp1.Inner(), temp2.Inner())) > 0
}

// AlwaysTouchesTPointGeo returns true if the two temporal points touch.
func AlwaysTouchesTPointGeo(temp Temporal, gs *Geom) bool {
	return int(C.atouches_tpoint_geo(temp.Inner(), gs._inner)) > 0
}

func EverContainsGeoTPoint(gs *Geom, temp Temporal) bool {
	return int(C.econtains_geo_tpoint(gs._inner, temp.Inner())) > 0
}

func EverDisjointTPointGeo(temp Temporal, gs *Geom) bool {
	return int(C.edisjoint_tpoint_geo(temp.Inner(), gs._inner)) > 0
}

func EverDisjointTPointTPoint(temp1, temp2 Temporal) bool {
	return int(C.edisjoint_tpoint_tpoint(temp1.Inner(), temp2.Inner())) > 0
}

func EverDWithinTPointGeo(temp Temporal, gs *Geom, dist float64) bool {
	return int(C.edwithin_tpoint_geo(temp.Inner(), gs._inner, C.double(dist))) > 0
}

func EverDWithinTPointTPoint(temp1, temp2 Temporal, dist float64) bool {
	return int(C.edwithin_tpoint_tpoint(temp1.Inner(), temp2.Inner(), C.double(dist))) > 0
}

func EverIntersectsTPointGeo(temp Temporal, gs *Geom) bool {
	return int(C.eintersects_tpoint_geo(temp.Inner(), gs._inner)) > 0
}

func EverIntersectsTPointTPoint(temp1, temp2 Temporal) bool {
	return int(C.eintersects_tpoint_tpoint(temp1.Inner(), temp2.Inner())) > 0
}

func EverTouchesTPointGeo(temp Temporal, gs *Geom) bool {
	return int(C.etouches_tpoint_geo(temp.Inner(), gs._inner)) > 0
}

func AlwaysEqTPointPoint(temp Temporal, gs *Geom) bool {
	return int(C.always_eq_tpoint_point(temp.Inner(), gs._inner)) > 0
}

func AlwaysNeTPointPoint(temp Temporal, gs *Geom) bool {
	return int(C.always_ne_tpoint_point(temp.Inner(), gs._inner)) > 0
}

func EverEqTPointPoint(temp Temporal, gs *Geom) bool {
	return int(C.ever_eq_tpoint_point(temp.Inner(), gs._inner)) > 0
}

func EverNeTPointPoint(temp Temporal, gs *Geom) bool {
	return int(C.ever_ne_tpoint_point(temp.Inner(), gs._inner)) > 0
}

func TContainsGeoTPoint[TP TPoint](gs *Geom, temp TP, restr, atvalue bool) Temporal {
	res := C.tcontains_geo_tpoint(gs._inner, temp.Inner(), C.bool(restr), C.bool(atvalue))
	return CreateTemporal(res)
}

func TDisjointTPointGeo[TP TPoint](temp TP, gs *Geom, restr, atvalue bool) Temporal {
	res := C.tdisjoint_tpoint_geo(temp.Inner(), gs._inner, C.bool(restr), C.bool(atvalue))
	return CreateTemporal(res)
}

func TDisjointTPointTPoint[TP1 TPoint, TP2 TPoint](temp1 TP1, temp2 TP2, restr, atvalue bool) Temporal {
	res := C.tdisjoint_tpoint_tpoint(temp1.Inner(), temp2.Inner(), C.bool(restr), C.bool(atvalue))
	return CreateTemporal(res)
}

func TDWithinTPointGeo[TP TPoint](temp TP, gs *Geom, dist float64, restr, atvalue bool) Temporal {
	res := C.tdwithin_tpoint_geo(temp.Inner(), gs._inner, C.double(dist), C.bool(restr), C.bool(atvalue))
	return CreateTemporal(res)
}

func TDWithinTPointTPoint[TP1 TPoint, TP2 TPoint](temp1 TP1, temp2 TP2, dist float64, restr, atvalue bool) Temporal {
	res := C.tdwithin_tpoint_tpoint(temp1.Inner(), temp2.Inner(), C.double(dist), C.bool(restr), C.bool(atvalue))
	return CreateTemporal(res)
}

func TIntersectsTPointGeo[TP TPoint](temp TP, gs *Geom, restr, atvalue bool) Temporal {
	res := C.tintersects_tpoint_geo(temp.Inner(), gs._inner, C.bool(restr), C.bool(atvalue))
	return CreateTemporal(res)
}

func TIntersectsTPointTPoint[TP1 TPoint, TP2 TPoint](temp1 TP1, temp2 TP2, restr, atvalue bool) Temporal {
	res := C.tintersects_tpoint_tpoint(temp1.Inner(), temp2.Inner(), C.bool(restr), C.bool(atvalue))
	return CreateTemporal(res)
}

func TTouchesTPointGeo[TP TPoint](temp TP, gs *Geom, restr, atvalue bool) Temporal {
	res := C.ttouches_tpoint_geo(temp.Inner(), gs._inner, C.bool(restr), C.bool(atvalue))
	return CreateTemporal(res)
}

func DistanceTPointPoint[TP TPoint](temp TP, gs *Geom) Temporal {
	res := C.distance_tpoint_point(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}

func DistanceTPointTPoint[TP1 TPoint, TP2 TPoint](temp1 TP1, temp2 TP2) Temporal {
	res := C.distance_tpoint_tpoint(temp1.Inner(), temp2.Inner())
	return CreateTemporal(res)
}

func TEqTPointPoint[TP TPoint](temp TP, gs *Geom) Temporal {
	res := C.teq_tpoint_point(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}

func TNeTPointPoint[TP TPoint](temp TP, gs *Geom) Temporal {
	res := C.tne_tpoint_point(temp.Inner(), gs._inner)
	return CreateTemporal(res)
}

// TPointTransform Return a temporal point transformed to another SRID
func TPointTransform[T TPoint](temp T, output T, srid_to int) T {
	c_temp := C.tpoint_transform(temp.Inner(), C.int(srid_to))
	output.Init(c_temp)
	return output
}

// TPointTrajectory Return the trajectory of a temporal point
func TPointTrajectory[TP TPoint](tp TP) *Geom {
	trajectory := C.tpoint_trajectory(tp.Inner())
	return &Geom{_inner: trajectory}
}
