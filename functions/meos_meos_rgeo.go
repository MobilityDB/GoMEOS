package functions

/*
#include <stddef.h>
#include "meos.h"
#include "meos_catalog.h"
#include "meos_geo.h"
#include "meos_internal.h"
#include "meos_internal_geo.h"
#include "meos_npoint.h"
#include "meos_cbuffer.h"
#include "meos_pose.h"
#include "meos_rgeo.h"
#include "meos_h3.h"
#include "meos_quadbin.h"
#include "meos_json.h"
#include "meos_pointcloud.h"

// cgo reads C.union_* as a union TYPE; alias the union operators so C.g<name> resolves them as functions.
#define gunion_bigint_set union_bigint_set
#define gunion_bigint_span union_bigint_span
#define gunion_bigint_spanset union_bigint_spanset
#define gunion_cbuffer_set union_cbuffer_set
#define gunion_date_set union_date_set
#define gunion_date_span union_date_span
#define gunion_date_spanset union_date_spanset
#define gunion_float_set union_float_set
#define gunion_float_span union_float_span
#define gunion_float_spanset union_float_spanset
#define gunion_geo_set union_geo_set
#define gunion_h3index_set union_h3index_set
#define gunion_int_set union_int_set
#define gunion_int_span union_int_span
#define gunion_int_spanset union_int_spanset
#define gunion_jsonb_set union_jsonb_set
#define gunion_npoint_set union_npoint_set
#define gunion_pcpatch_set union_pcpatch_set
#define gunion_pcpoint_set union_pcpoint_set
#define gunion_pose_set union_pose_set
#define gunion_posechain_set union_posechain_set
#define gunion_quadbin_set union_quadbin_set
#define gunion_s2cell_set union_s2cell_set
#define gunion_set_bigint union_set_bigint
#define gunion_set_cbuffer union_set_cbuffer
#define gunion_set_date union_set_date
#define gunion_set_float union_set_float
#define gunion_set_geo union_set_geo
#define gunion_set_h3index union_set_h3index
#define gunion_set_int union_set_int
#define gunion_set_jsonb union_set_jsonb
#define gunion_set_npoint union_set_npoint
#define gunion_set_pcpatch union_set_pcpatch
#define gunion_set_pcpoint union_set_pcpoint
#define gunion_set_pose union_set_pose
#define gunion_set_posechain union_set_posechain
#define gunion_set_quadbin union_set_quadbin
#define gunion_set_s2cell union_set_s2cell
#define gunion_set_set union_set_set
#define gunion_set_text union_set_text
#define gunion_set_timestamptz union_set_timestamptz
#define gunion_set_value union_set_value
#define gunion_span_bigint union_span_bigint
#define gunion_span_date union_span_date
#define gunion_span_float union_span_float
#define gunion_span_int union_span_int
#define gunion_span_span union_span_span
#define gunion_span_spanset union_span_spanset
#define gunion_span_timestamptz union_span_timestamptz
#define gunion_span_value union_span_value
#define gunion_spanset_bigint union_spanset_bigint
#define gunion_spanset_date union_spanset_date
#define gunion_spanset_float union_spanset_float
#define gunion_spanset_int union_spanset_int
#define gunion_spanset_span union_spanset_span
#define gunion_spanset_spanset union_spanset_spanset
#define gunion_spanset_timestamptz union_spanset_timestamptz
#define gunion_spanset_value union_spanset_value
#define gunion_stbox_stbox union_stbox_stbox
#define gunion_tbox_tbox union_tbox_tbox
#define gunion_text_set union_text_set
#define gunion_timestamptz_set union_timestamptz_set
#define gunion_timestamptz_span union_timestamptz_span
#define gunion_timestamptz_spanset union_timestamptz_spanset
#define gunion_tpcbox_tpcbox union_tpcbox_tpcbox
#define gunion_value_set union_value_set
#define gunion_value_span union_value_span
#define gunion_value_spanset union_value_spanset
*/
import "C"
import (
	"unsafe"
)

var _ = unsafe.Pointer(nil)

// TrgeometryIn wraps MEOS C function trgeometry_in.
func TrgeometryIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.trgeometry_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryFromMFJSON wraps MEOS C function trgeometry_from_mfjson.
func TrgeometryFromMFJSON(mfjson string) (_r0 *Temporal, _err error) {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	C.meos_errno_reset()
	_cret := C.trgeometry_from_mfjson(_c_mfjson)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryOut wraps MEOS C function trgeometry_out.
func TrgeometryOut(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_out(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TrgeometryAsText wraps MEOS C function trgeometry_as_text.
func TrgeometryAsText(temp *Temporal, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_as_text(temp._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TrgeometryAsEWKT wraps MEOS C function trgeometry_as_ewkt.
func TrgeometryAsEWKT(temp *Temporal, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_as_ewkt(temp._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TrgeometryinstMake wraps MEOS C function trgeometryinst_make.
func TrgeometryinstMake(geom *Geom, pose *Pose, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometryinst_make(geom._inner, pose._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TrgeometryseqMake wraps MEOS C function trgeometryseq_make.
func TrgeometryseqMake(geom *Geom, instants unsafe.Pointer, count int, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometryseq_make(geom._inner, (**C.TInstant)(unsafe.Pointer(instants)), C.int(count), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TrgeometryseqsetMake wraps MEOS C function trgeometryseqset_make.
func TrgeometryseqsetMake(geom *Geom, sequences unsafe.Pointer, count int, normalize bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometryseqset_make(geom._inner, (**C.TSequence)(unsafe.Pointer(sequences)), C.int(count), C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TrgeometryseqsetMakeGaps wraps MEOS C function trgeometryseqset_make_gaps.
func TrgeometryseqsetMakeGaps(geom *Geom, instants unsafe.Pointer, count int, interp Interpolation, maxt *Interval, maxdist float64) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometryseqset_make_gaps(geom._inner, (**C.TInstant)(unsafe.Pointer(instants)), C.int(count), C.interpType(interp), maxt._inner, C.double(maxdist))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// GeometryTposeToTrgeometry wraps MEOS C function geometry_tpose_to_trgeometry.
func GeometryTposeToTrgeometry(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.geometry_tpose_to_trgeometry(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryToTpose wraps MEOS C function trgeometry_to_tpose.
func TrgeometryToTpose(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_to_tpose(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryToTgeompoint wraps MEOS C function trgeometry_to_tgeompoint.
func TrgeometryToTgeompoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_to_tgeompoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryToTgeometry wraps MEOS C function trgeometry_to_tgeometry.
func TrgeometryToTgeometry(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_to_tgeometry(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryEndInstant wraps MEOS C function trgeometry_end_instant.
func TrgeometryEndInstant(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_end_instant(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TrgeometryEndSequence wraps MEOS C function trgeometry_end_sequence.
func TrgeometryEndSequence(temp *Temporal) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_end_sequence(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TrgeometryEndValue wraps MEOS C function trgeometry_end_value.
func TrgeometryEndValue(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TrgeometryGeom wraps MEOS C function trgeometry_geom.
func TrgeometryGeom(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_geom(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TrgeometryInstantN wraps MEOS C function trgeometry_instant_n.
func TrgeometryInstantN(temp *Temporal, n int) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_instant_n(temp._inner, C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TrgeometryInstants wraps MEOS C function trgeometry_instants.
func TrgeometryInstants(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_instants(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TrgeometryPoints wraps MEOS C function trgeometry_points.
func TrgeometryPoints(temp *Temporal) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_points(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TrgeometryYaw wraps MEOS C function trgeometry_yaw.
func TrgeometryYaw(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_yaw(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryPitch wraps MEOS C function trgeometry_pitch.
func TrgeometryPitch(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_pitch(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryRoll wraps MEOS C function trgeometry_roll.
func TrgeometryRoll(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_roll(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometrySegments wraps MEOS C function trgeometry_segments.
func TrgeometrySegments(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_segments(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TrgeometrySequenceN wraps MEOS C function trgeometry_sequence_n.
func TrgeometrySequenceN(temp *Temporal, i int) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_sequence_n(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TrgeometrySequences wraps MEOS C function trgeometry_sequences.
func TrgeometrySequences(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_sequences(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TrgeometryStartInstant wraps MEOS C function trgeometry_start_instant.
func TrgeometryStartInstant(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_start_instant(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TrgeometryStartSequence wraps MEOS C function trgeometry_start_sequence.
func TrgeometryStartSequence(temp *Temporal) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_start_sequence(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TrgeometryStartValue wraps MEOS C function trgeometry_start_value.
func TrgeometryStartValue(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TrgeometryValue wraps MEOS C function trgeometry_value.
func TrgeometryValue(temp *Temporal) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// TrgeometryValues wraps MEOS C function trgeometry_values.
func TrgeometryValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TrgeometryUnnest wraps MEOS C function trgeometry_unnest.
func TrgeometryUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_unnest(temp._inner, (***C.GSERIALIZED)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TrgeometryValueN wraps MEOS C function trgeometry_value_n.
func TrgeometryValueN(temp *Temporal, n int) (_r0 bool, _r1 *Geom, _err error) {
	var _out_result *C.GSERIALIZED
	C.meos_errno_reset()
	_cret := C.trgeometry_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Geom{_inner: _out_result}, nil
}


// TrgeometryTraversedArea wraps MEOS C function trgeometry_traversed_area.
func TrgeometryTraversedArea(temp *Temporal, unary_union bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_traversed_area(temp._inner, C.bool(unary_union))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TrgeometryCentroid wraps MEOS C function trgeometry_centroid.
func TrgeometryCentroid(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_centroid(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryConvexHull wraps MEOS C function trgeometry_convex_hull.
func TrgeometryConvexHull(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_convex_hull(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TrgeometryBodyPointTrajectory wraps MEOS C function trgeometry_body_point_trajectory.
func TrgeometryBodyPointTrajectory(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_body_point_trajectory(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometrySpaceBoxes wraps MEOS C function trgeometry_space_boxes.
func TrgeometrySpaceBoxes(temp *Temporal, xsize float64, ysize float64, zsize float64, sorigin *Geom, bitmatrix bool, border_inc bool, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_space_boxes(temp._inner, C.double(xsize), C.double(ysize), C.double(zsize), sorigin._inner, C.bool(bitmatrix), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TrgeometryTimeBoxes wraps MEOS C function trgeometry_time_boxes.
func TrgeometryTimeBoxes(temp *Temporal, duration *Interval, torigin int64, bitmatrix bool, border_inc bool, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_time_boxes(temp._inner, duration._inner, C.TimestampTz(torigin), C.bool(bitmatrix), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TrgeometrySpaceTimeBoxes wraps MEOS C function trgeometry_space_time_boxes.
func TrgeometrySpaceTimeBoxes(temp *Temporal, xsize float64, ysize float64, zsize float64, duration *Interval, sorigin *Geom, torigin int64, bitmatrix bool, border_inc bool, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_space_time_boxes(temp._inner, C.double(xsize), C.double(ysize), C.double(zsize), duration._inner, sorigin._inner, C.TimestampTz(torigin), C.bool(bitmatrix), C.bool(border_inc), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TrgeometryStboxes wraps MEOS C function trgeometry_stboxes.
func TrgeometryStboxes(temp *Temporal, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_stboxes(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TrgeometrySplitNStboxes wraps MEOS C function trgeometry_split_n_stboxes.
func TrgeometrySplitNStboxes(temp *Temporal, box_count int, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_split_n_stboxes(temp._inner, C.int(box_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TrgeometrySplitEachNStboxes wraps MEOS C function trgeometry_split_each_n_stboxes.
func TrgeometrySplitEachNStboxes(temp *Temporal, elem_count int, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_split_each_n_stboxes(temp._inner, C.int(elem_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TrgeometryLength wraps MEOS C function trgeometry_length.
func TrgeometryLength(temp *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_length(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TrgeometryCumulativeLength wraps MEOS C function trgeometry_cumulative_length.
func TrgeometryCumulativeLength(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_cumulative_length(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryAngularSpeed wraps MEOS C function trgeometry_angular_speed.
func TrgeometryAngularSpeed(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_angular_speed(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometrySpeed wraps MEOS C function trgeometry_speed.
func TrgeometrySpeed(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_speed(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryTwcentroid wraps MEOS C function trgeometry_twcentroid.
func TrgeometryTwcentroid(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_twcentroid(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TrgeometryAppendTinstant wraps MEOS C function trgeometry_append_tinstant.
func TrgeometryAppendTinstant(temp *Temporal, inst *TInstant, interp Interpolation, maxdist float64, maxt *Interval, expand bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_append_tinstant(temp._inner, inst._inner, C.interpType(interp), C.double(maxdist), maxt._inner, C.bool(expand))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryAppendTsequence wraps MEOS C function trgeometry_append_tsequence.
func TrgeometryAppendTsequence(temp *Temporal, seq *TSequence, expand bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_append_tsequence(temp._inner, seq._inner, C.bool(expand))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryDeleteTimestamptz wraps MEOS C function trgeometry_delete_timestamptz.
func TrgeometryDeleteTimestamptz(temp *Temporal, t int64, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_delete_timestamptz(temp._inner, C.TimestampTz(t), C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryDeleteTstzset wraps MEOS C function trgeometry_delete_tstzset.
func TrgeometryDeleteTstzset(temp *Temporal, s *Set, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_delete_tstzset(temp._inner, s._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryDeleteTstzspan wraps MEOS C function trgeometry_delete_tstzspan.
func TrgeometryDeleteTstzspan(temp *Temporal, s *Span, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_delete_tstzspan(temp._inner, s._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryDeleteTstzspanset wraps MEOS C function trgeometry_delete_tstzspanset.
func TrgeometryDeleteTstzspanset(temp *Temporal, ss *SpanSet, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_delete_tstzspanset(temp._inner, ss._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryMerge wraps MEOS C function trgeometry_merge.
func TrgeometryMerge(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_merge(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryMergeArray wraps MEOS C function trgeometry_merge_array.
func TrgeometryMergeArray(temparr unsafe.Pointer, count int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_merge_array((**C.Temporal)(unsafe.Pointer(temparr)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryRound wraps MEOS C function trgeometry_round.
func TrgeometryRound(temp *Temporal, maxdd int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_round(temp._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryRotate wraps MEOS C function trgeometry_rotate.
func TrgeometryRotate(temp *Temporal, angle float64, x0 float64, y0 float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_rotate(temp._inner, C.double(angle), C.double(x0), C.double(y0))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryRotateX wraps MEOS C function trgeometry_rotate_x.
func TrgeometryRotateX(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_rotate_x(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryRotateY wraps MEOS C function trgeometry_rotate_y.
func TrgeometryRotateY(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_rotate_y(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryRotateZ wraps MEOS C function trgeometry_rotate_z.
func TrgeometryRotateZ(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_rotate_z(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometrySetInterp wraps MEOS C function trgeometry_set_interp.
func TrgeometrySetInterp(temp *Temporal, interp Interpolation) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_set_interp(temp._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryTranslate wraps MEOS C function trgeometry_translate.
func TrgeometryTranslate(temp *Temporal, deltax float64, deltay float64, deltaz float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_translate(temp._inner, C.double(deltax), C.double(deltay), C.double(deltaz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryAsTinstant wraps MEOS C function trgeometry_as_tinstant.
func TrgeometryAsTinstant(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_as_tinstant(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TrgeometryAsTsequence wraps MEOS C function trgeometry_as_tsequence.
func TrgeometryAsTsequence(temp *Temporal, interp_str string) (_r0 *TSequence, _err error) {
	_c_interp_str := C.CString(interp_str)
	defer C.free(unsafe.Pointer(_c_interp_str))
	C.meos_errno_reset()
	_cret := C.trgeometry_as_tsequence(temp._inner, _c_interp_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TrgeometryAsTsequenceset wraps MEOS C function trgeometry_as_tsequenceset.
func TrgeometryAsTsequenceset(temp *Temporal, interp_str string) (_r0 *TSequenceSet, _err error) {
	_c_interp_str := C.CString(interp_str)
	defer C.free(unsafe.Pointer(_c_interp_str))
	C.meos_errno_reset()
	_cret := C.trgeometry_as_tsequenceset(temp._inner, _c_interp_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TrgeometryAfterTimestamptz wraps MEOS C function trgeometry_after_timestamptz.
func TrgeometryAfterTimestamptz(temp *Temporal, t int64, strict bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_after_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryBeforeTimestamptz wraps MEOS C function trgeometry_before_timestamptz.
func TrgeometryBeforeTimestamptz(temp *Temporal, t int64, strict bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_before_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryRestrictValues wraps MEOS C function trgeometry_restrict_values.
func TrgeometryRestrictValues(temp *Temporal, s *Set, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_restrict_values(temp._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryRestrictTimestamptz wraps MEOS C function trgeometry_restrict_timestamptz.
func TrgeometryRestrictTimestamptz(temp *Temporal, t int64, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_restrict_timestamptz(temp._inner, C.TimestampTz(t), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryRestrictTstzset wraps MEOS C function trgeometry_restrict_tstzset.
func TrgeometryRestrictTstzset(temp *Temporal, s *Set, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_restrict_tstzset(temp._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryRestrictTstzspan wraps MEOS C function trgeometry_restrict_tstzspan.
func TrgeometryRestrictTstzspan(temp *Temporal, s *Span, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_restrict_tstzspan(temp._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryRestrictTstzspanset wraps MEOS C function trgeometry_restrict_tstzspanset.
func TrgeometryRestrictTstzspanset(temp *Temporal, ss *SpanSet, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_restrict_tstzspanset(temp._inner, ss._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryAtGeom wraps MEOS C function trgeometry_at_geom.
func TrgeometryAtGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_at_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryMinusGeom wraps MEOS C function trgeometry_minus_geom.
func TrgeometryMinusGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_minus_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryAtSTBOX wraps MEOS C function trgeometry_at_stbox.
func TrgeometryAtSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_at_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryMinusSTBOX wraps MEOS C function trgeometry_minus_stbox.
func TrgeometryMinusSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_minus_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryAtValue wraps MEOS C function trgeometry_at_value.
func TrgeometryAtValue(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_at_value(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryMinusValue wraps MEOS C function trgeometry_minus_value.
func TrgeometryMinusValue(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_minus_value(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryAtValues wraps MEOS C function trgeometry_at_values.
func TrgeometryAtValues(temp *Temporal, s *Set) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_at_values(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryMinusValues wraps MEOS C function trgeometry_minus_values.
func TrgeometryMinusValues(temp *Temporal, s *Set) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_minus_values(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryAtTimestamptz wraps MEOS C function trgeometry_at_timestamptz.
func TrgeometryAtTimestamptz(temp *Temporal, t int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_at_timestamptz(temp._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryMinusTimestamptz wraps MEOS C function trgeometry_minus_timestamptz.
func TrgeometryMinusTimestamptz(temp *Temporal, t int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_minus_timestamptz(temp._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryAtTstzset wraps MEOS C function trgeometry_at_tstzset.
func TrgeometryAtTstzset(temp *Temporal, s *Set) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_at_tstzset(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryMinusTstzset wraps MEOS C function trgeometry_minus_tstzset.
func TrgeometryMinusTstzset(temp *Temporal, s *Set) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_minus_tstzset(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryAtTstzspan wraps MEOS C function trgeometry_at_tstzspan.
func TrgeometryAtTstzspan(temp *Temporal, s *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_at_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryMinusTstzspan wraps MEOS C function trgeometry_minus_tstzspan.
func TrgeometryMinusTstzspan(temp *Temporal, s *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_minus_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryAtTstzspanset wraps MEOS C function trgeometry_at_tstzspanset.
func TrgeometryAtTstzspanset(temp *Temporal, ss *SpanSet) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_at_tstzspanset(temp._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TrgeometryMinusTstzspanset wraps MEOS C function trgeometry_minus_tstzspanset.
func TrgeometryMinusTstzspanset(temp *Temporal, ss *SpanSet) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.trgeometry_minus_tstzspanset(temp._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTrgeometryGeo wraps MEOS C function tdistance_trgeometry_geo.
func TdistanceTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTrgeometryTpoint wraps MEOS C function tdistance_trgeometry_tpoint.
func TdistanceTrgeometryTpoint(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_trgeometry_tpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTrgeometryTrgeometry wraps MEOS C function tdistance_trgeometry_trgeometry.
func TdistanceTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// NadSTBOXTrgeometry wraps MEOS C function nad_stbox_trgeometry.
func NadSTBOXTrgeometry(box *STBox, temp *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_stbox_trgeometry(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTrgeometryGeo wraps MEOS C function nad_trgeometry_geo.
func NadTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTrgeometrySTBOX wraps MEOS C function nad_trgeometry_stbox.
func NadTrgeometrySTBOX(temp *Temporal, box *STBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_trgeometry_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTrgeometryTpoint wraps MEOS C function nad_trgeometry_tpoint.
func NadTrgeometryTpoint(temp1 *Temporal, temp2 *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_trgeometry_tpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTrgeometryTrgeometry wraps MEOS C function nad_trgeometry_trgeometry.
func NadTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NaiTrgeometryGeo wraps MEOS C function nai_trgeometry_geo.
func NaiTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// NaiTrgeometryTpoint wraps MEOS C function nai_trgeometry_tpoint.
func NaiTrgeometryTpoint(temp1 *Temporal, temp2 *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_trgeometry_tpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// NaiTrgeometryTrgeometry wraps MEOS C function nai_trgeometry_trgeometry.
func NaiTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// ShortestlineTrgeometryGeo wraps MEOS C function shortestline_trgeometry_geo.
func ShortestlineTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// ShortestlineTrgeometryTpoint wraps MEOS C function shortestline_trgeometry_tpoint.
func ShortestlineTrgeometryTpoint(temp1 *Temporal, temp2 *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_trgeometry_tpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// ShortestlineTrgeometryTrgeometry wraps MEOS C function shortestline_trgeometry_trgeometry.
func ShortestlineTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TdwithinGeoTrgeometry wraps MEOS C function tdwithin_geo_trgeometry.
func TdwithinGeoTrgeometry(gs *Geom, temp *Temporal, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_geo_trgeometry(gs._inner, temp._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdwithinTrgeometryGeo wraps MEOS C function tdwithin_trgeometry_geo.
func TdwithinTrgeometryGeo(temp *Temporal, gs *Geom, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_trgeometry_geo(temp._inner, gs._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdwithinTrgeometryTrgeometry wraps MEOS C function tdwithin_trgeometry_trgeometry.
func TdwithinTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdwithin_trgeometry_trgeometry(temp1._inner, temp2._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AlwaysEqGeoTrgeometry wraps MEOS C function always_eq_geo_trgeometry.
func AlwaysEqGeoTrgeometry(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_geo_trgeometry(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTrgeometryGeo wraps MEOS C function always_eq_trgeometry_geo.
func AlwaysEqTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTrgeometryTrgeometry wraps MEOS C function always_eq_trgeometry_trgeometry.
func AlwaysEqTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeGeoTrgeometry wraps MEOS C function always_ne_geo_trgeometry.
func AlwaysNeGeoTrgeometry(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_geo_trgeometry(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTrgeometryGeo wraps MEOS C function always_ne_trgeometry_geo.
func AlwaysNeTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTrgeometryTrgeometry wraps MEOS C function always_ne_trgeometry_trgeometry.
func AlwaysNeTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqGeoTrgeometry wraps MEOS C function ever_eq_geo_trgeometry.
func EverEqGeoTrgeometry(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_geo_trgeometry(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTrgeometryGeo wraps MEOS C function ever_eq_trgeometry_geo.
func EverEqTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTrgeometryTrgeometry wraps MEOS C function ever_eq_trgeometry_trgeometry.
func EverEqTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeGeoTrgeometry wraps MEOS C function ever_ne_geo_trgeometry.
func EverNeGeoTrgeometry(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_geo_trgeometry(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTrgeometryGeo wraps MEOS C function ever_ne_trgeometry_geo.
func EverNeTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTrgeometryTrgeometry wraps MEOS C function ever_ne_trgeometry_trgeometry.
func EverNeTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqGeoTrgeometry wraps MEOS C function teq_geo_trgeometry.
func TeqGeoTrgeometry(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_geo_trgeometry(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTrgeometryGeo wraps MEOS C function teq_trgeometry_geo.
func TeqTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneGeoTrgeometry wraps MEOS C function tne_geo_trgeometry.
func TneGeoTrgeometry(gs *Geom, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_geo_trgeometry(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTrgeometryGeo wraps MEOS C function tne_trgeometry_geo.
func TneTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// EcontainsGeoTrgeometry wraps MEOS C function econtains_geo_trgeometry.
func EcontainsGeoTrgeometry(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.econtains_geo_trgeometry(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcontainsGeoTrgeometry wraps MEOS C function acontains_geo_trgeometry.
func AcontainsGeoTrgeometry(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acontains_geo_trgeometry(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcoversGeoTrgeometry wraps MEOS C function ecovers_geo_trgeometry.
func EcoversGeoTrgeometry(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ecovers_geo_trgeometry(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcoversGeoTrgeometry wraps MEOS C function acovers_geo_trgeometry.
func AcoversGeoTrgeometry(gs *Geom, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acovers_geo_trgeometry(gs._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EcoversTrgeometryGeo wraps MEOS C function ecovers_trgeometry_geo.
func EcoversTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ecovers_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AcoversTrgeometryGeo wraps MEOS C function acovers_trgeometry_geo.
func AcoversTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.acovers_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdisjointTrgeometryGeo wraps MEOS C function edisjoint_trgeometry_geo.
func EdisjointTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edisjoint_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdisjointTrgeometryGeo wraps MEOS C function adisjoint_trgeometry_geo.
func AdisjointTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adisjoint_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EintersectsTrgeometryGeo wraps MEOS C function eintersects_trgeometry_geo.
func EintersectsTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AintersectsTrgeometryGeo wraps MEOS C function aintersects_trgeometry_geo.
func AintersectsTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.aintersects_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EtouchesTrgeometryGeo wraps MEOS C function etouches_trgeometry_geo.
func EtouchesTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.etouches_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AtouchesTrgeometryGeo wraps MEOS C function atouches_trgeometry_geo.
func AtouchesTrgeometryGeo(temp *Temporal, gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.atouches_trgeometry_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdwithinTrgeometryGeo wraps MEOS C function edwithin_trgeometry_geo.
func EdwithinTrgeometryGeo(temp *Temporal, gs *Geom, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edwithin_trgeometry_geo(temp._inner, gs._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdwithinTrgeometryGeo wraps MEOS C function adwithin_trgeometry_geo.
func AdwithinTrgeometryGeo(temp *Temporal, gs *Geom, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adwithin_trgeometry_geo(temp._inner, gs._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdisjointTrgeometryTrgeometry wraps MEOS C function edisjoint_trgeometry_trgeometry.
func EdisjointTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edisjoint_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdisjointTrgeometryTrgeometry wraps MEOS C function adisjoint_trgeometry_trgeometry.
func AdisjointTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adisjoint_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EintersectsTrgeometryTrgeometry wraps MEOS C function eintersects_trgeometry_trgeometry.
func EintersectsTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AintersectsTrgeometryTrgeometry wraps MEOS C function aintersects_trgeometry_trgeometry.
func AintersectsTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.aintersects_trgeometry_trgeometry(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EdwithinTrgeometryTrgeometry wraps MEOS C function edwithin_trgeometry_trgeometry.
func EdwithinTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.edwithin_trgeometry_trgeometry(temp1._inner, temp2._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AdwithinTrgeometryTrgeometry wraps MEOS C function adwithin_trgeometry_trgeometry.
func AdwithinTrgeometryTrgeometry(temp1 *Temporal, temp2 *Temporal, dist float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.adwithin_trgeometry_trgeometry(temp1._inner, temp2._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}

