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

// PoseAsEWKB wraps MEOS C function pose_as_ewkb.
func PoseAsEWKB(pose *Pose, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_ewkb(pose._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PoseAsEWKT wraps MEOS C function pose_as_ewkt.
func PoseAsEWKT(pose *Pose, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_ewkt(pose._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PoseAsHexewkb wraps MEOS C function pose_as_hexewkb.
func PoseAsHexewkb(pose *Pose, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_hexewkb(pose._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PoseAsHexwkb wraps MEOS C function pose_as_hexwkb.
func PoseAsHexwkb(pose *Pose, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_hexwkb(pose._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PoseAsText wraps MEOS C function pose_as_text.
func PoseAsText(pose *Pose, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_text(pose._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PoseAsWKB wraps MEOS C function pose_as_wkb.
func PoseAsWKB(pose *Pose, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_wkb(pose._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PoseFromWKB wraps MEOS C function pose_from_wkb.
func PoseFromWKB(wkb unsafe.Pointer, size uint) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseFromHexwkb wraps MEOS C function pose_from_hexwkb.
func PoseFromHexwkb(hexwkb string) (_r0 *Pose, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.pose_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseFromText wraps MEOS C function pose_from_text.
func PoseFromText(str string) (_r0 *Pose, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.pose_from_text(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseIn wraps MEOS C function pose_in.
func PoseIn(str string) (_r0 *Pose, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.pose_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseOut wraps MEOS C function pose_out.
func PoseOut(pose *Pose, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_out(pose._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PoseFromGeopose wraps MEOS C function pose_from_geopose.
func PoseFromGeopose(json string) (_r0 *Pose, _err error) {
	_c_json := C.CString(json)
	defer C.free(unsafe.Pointer(_c_json))
	C.meos_errno_reset()
	_cret := C.pose_from_geopose(_c_json)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseAsGeopose wraps MEOS C function pose_as_geopose.
func PoseAsGeopose(pose *Pose, conformance int, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_as_geopose(pose._inner, C.int(conformance), C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TposeFromGeopose wraps MEOS C function tpose_from_geopose.
func TposeFromGeopose(json string) (_r0 *Temporal, _err error) {
	_c_json := C.CString(json)
	defer C.free(unsafe.Pointer(_c_json))
	C.meos_errno_reset()
	_cret := C.tpose_from_geopose(_c_json)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeAsGeopose wraps MEOS C function tpose_as_geopose.
func TposeAsGeopose(temp *Temporal, conformance int, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_as_geopose(temp._inner, C.int(conformance), C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TposeAsGeoposeStreamHeader wraps MEOS C function tpose_as_geopose_stream_header.
func TposeAsGeoposeStreamHeader(temp *Temporal, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_as_geopose_stream_header(temp._inner, C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TposeAsGeoposeStreamElement wraps MEOS C function tpose_as_geopose_stream_element.
func TposeAsGeoposeStreamElement(temp *Temporal, inst *TInstant, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_as_geopose_stream_element(temp._inner, inst._inner, C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TposeAsGeoposeStream wraps MEOS C function tpose_as_geopose_stream.
func TposeAsGeoposeStream(temp *Temporal, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_as_geopose_stream(temp._inner, C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GeoposeFrames wraps MEOS C function geopose_frames.
func GeoposeFrames(count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.geopose_frames((*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// GeoposeFrame wraps MEOS C function geopose_frame.
func GeoposeFrame(frame_id int32) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.geopose_frame(C.int32_t(frame_id))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// GeoposeFrameName wraps MEOS C function geopose_frame_name.
func GeoposeFrameName(frame_id int32) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.geopose_frame_name(C.int32_t(frame_id))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// GeoposeFrameSRID wraps MEOS C function geopose_frame_srid.
func GeoposeFrameSRID(frame_id int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.geopose_frame_srid(C.int32_t(frame_id))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// GeoposeFrameIsGeographic wraps MEOS C function geopose_frame_is_geographic.
func GeoposeFrameIsGeographic(frame_id int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geopose_frame_is_geographic(C.int32_t(frame_id))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseApplyGeo wraps MEOS C function pose_apply_geo.
func PoseApplyGeo(pose *Pose, body *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_apply_geo(pose._inner, body._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TposeApplyGeo wraps MEOS C function tpose_apply_geo.
func TposeApplyGeo(temp *Temporal, body *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_apply_geo(temp._inner, body._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeComposePose wraps MEOS C function tpose_compose_pose.
func TposeComposePose(body *Temporal, frame *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_compose_pose(body._inner, frame._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// PoseComposeTpose wraps MEOS C function pose_compose_tpose.
func PoseComposeTpose(body *Pose, frame *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_compose_tpose(body._inner, frame._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeComposeTpose wraps MEOS C function tpose_compose_tpose.
func TposeComposeTpose(body *Temporal, frame *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_compose_tpose(body._inner, frame._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeInverse wraps MEOS C function tpose_inverse.
func TposeInverse(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_inverse(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeRotate wraps MEOS C function tpose_rotate.
func TposeRotate(temp *Temporal, angle float64, x0 float64, y0 float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_rotate(temp._inner, C.double(angle), C.double(x0), C.double(y0))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeRotateX wraps MEOS C function tpose_rotate_x.
func TposeRotateX(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_rotate_x(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeRotateY wraps MEOS C function tpose_rotate_y.
func TposeRotateY(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_rotate_y(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeRotateZ wraps MEOS C function tpose_rotate_z.
func TposeRotateZ(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_rotate_z(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeTranslate wraps MEOS C function tpose_translate.
func TposeTranslate(temp *Temporal, deltax float64, deltay float64, deltaz float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_translate(temp._inner, C.double(deltax), C.double(deltay), C.double(deltaz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// PoseCopy wraps MEOS C function pose_copy.
func PoseCopy(pose *Pose) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_copy(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseMake2d wraps MEOS C function pose_make_2d.
func PoseMake2d(x float64, y float64, theta float64, geodetic bool, srid int32) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_make_2d(C.double(x), C.double(y), C.double(theta), C.bool(geodetic), C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseMake3d wraps MEOS C function pose_make_3d.
func PoseMake3d(x float64, y float64, z float64, W float64, X float64, Y float64, Z float64, geodetic bool, srid int32) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_make_3d(C.double(x), C.double(y), C.double(z), C.double(W), C.double(X), C.double(Y), C.double(Z), C.bool(geodetic), C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseMakePoint2d wraps MEOS C function pose_make_point2d.
func PoseMakePoint2d(gs *Geom, theta float64) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_make_point2d(gs._inner, C.double(theta))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseMakePoint3d wraps MEOS C function pose_make_point3d.
func PoseMakePoint3d(gs *Geom, W float64, X float64, Y float64, Z float64) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_make_point3d(gs._inner, C.double(W), C.double(X), C.double(Y), C.double(Z))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseMakePoint3dYpr wraps MEOS C function pose_make_point3d_ypr.
func PoseMakePoint3dYpr(gs *Geom, yaw float64, pitch float64, roll float64) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_make_point3d_ypr(gs._inner, C.double(yaw), C.double(pitch), C.double(roll))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseToPoint wraps MEOS C function pose_to_point.
func PoseToPoint(pose *Pose) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_to_point(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// PoseToSTBOX wraps MEOS C function pose_to_stbox.
func PoseToSTBOX(pose *Pose) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_to_stbox(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// PoseHash wraps MEOS C function pose_hash.
func PoseHash(pose *Pose) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_hash(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// PoseHashExtended wraps MEOS C function pose_hash_extended.
func PoseHashExtended(pose *Pose, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_hash_extended(pose._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// PoseMemSize wraps MEOS C function pose_mem_size.
func PoseMemSize(pose *Pose) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_mem_size(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// PoseQuaternion wraps MEOS C function pose_quaternion.
func PoseQuaternion(pose *Pose, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_quaternion(pose._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PoseYpr wraps MEOS C function pose_ypr.
func PoseYpr(pose *Pose, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_ypr(pose._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PoseYaw wraps MEOS C function pose_yaw.
func PoseYaw(pose *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_yaw(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// PosePitch wraps MEOS C function pose_pitch.
func PosePitch(pose *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_pitch(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// PoseRoll wraps MEOS C function pose_roll.
func PoseRoll(pose *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_roll(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// PoseAngularDistance wraps MEOS C function pose_angular_distance.
func PoseAngularDistance(pose1 *Pose, pose2 *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_angular_distance(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// PoseCompose wraps MEOS C function pose_compose.
func PoseCompose(body *Pose, frame *Pose) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_compose(body._inner, frame._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseInverse wraps MEOS C function pose_inverse.
func PoseInverse(pose *Pose) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_inverse(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseNormalize wraps MEOS C function pose_normalize.
func PoseNormalize(pose *Pose) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_normalize(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseRound wraps MEOS C function pose_round.
func PoseRound(pose *Pose, maxdd int) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_round(pose._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PosearrRound wraps MEOS C function posearr_round.
func PosearrRound(posearr unsafe.Pointer, count int, maxdd int) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posearr_round((**C.Pose)(unsafe.Pointer(posearr)), C.int(count), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PoseSetSRID wraps MEOS C function pose_set_srid.
func PoseSetSRID(pose *Pose, srid int32) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_set_srid(pose._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseSRID wraps MEOS C function pose_srid.
func PoseSRID(pose *Pose) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_srid(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// PoseTransform wraps MEOS C function pose_transform.
func PoseTransform(pose *Pose, srid int32) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_transform(pose._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseTransformPipeline wraps MEOS C function pose_transform_pipeline.
func PoseTransformPipeline(pose *Pose, pipelinestr string, srid int32, is_forward bool) (_r0 *Pose, _err error) {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	C.meos_errno_reset()
	_cret := C.pose_transform_pipeline(pose._inner, _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PoseTstzspanToSTBOX wraps MEOS C function pose_tstzspan_to_stbox.
func PoseTstzspanToSTBOX(pose *Pose, s *Span) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_tstzspan_to_stbox(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// PoseTimestamptzToSTBOX wraps MEOS C function pose_timestamptz_to_stbox.
func PoseTimestamptzToSTBOX(pose *Pose, t int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_timestamptz_to_stbox(pose._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// DistancePoseGeo wraps MEOS C function distance_pose_geo.
func DistancePoseGeo(pose *Pose, gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_pose_geo(pose._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistancePosePose wraps MEOS C function distance_pose_pose.
func DistancePosePose(pose1 *Pose, pose2 *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_pose_pose(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistancePoseSTBOX wraps MEOS C function distance_pose_stbox.
func DistancePoseSTBOX(pose *Pose, box *STBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_pose_stbox(pose._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// PoseCmp wraps MEOS C function pose_cmp.
func PoseCmp(pose1 *Pose, pose2 *Pose) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_cmp(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// PoseEq wraps MEOS C function pose_eq.
func PoseEq(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_eq(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseGe wraps MEOS C function pose_ge.
func PoseGe(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_ge(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseGt wraps MEOS C function pose_gt.
func PoseGt(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_gt(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseLe wraps MEOS C function pose_le.
func PoseLe(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_le(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseLt wraps MEOS C function pose_lt.
func PoseLt(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_lt(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PoseNe wraps MEOS C function pose_ne.
func PoseNe(pose1 *Pose, pose2 *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_ne(pose1._inner, pose2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PosesetIn wraps MEOS C function poseset_in.
func PosesetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.poseset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PosesetOut wraps MEOS C function poseset_out.
func PosesetOut(s *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.poseset_out(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PosesetMake wraps MEOS C function poseset_make.
func PosesetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.poseset_make((**C.Pose)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PoseToSet wraps MEOS C function pose_to_set.
func PoseToSet(pose *Pose) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_to_set(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PosesetEndValue wraps MEOS C function poseset_end_value.
func PosesetEndValue(s *Set) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.poseset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PosesetStartValue wraps MEOS C function poseset_start_value.
func PosesetStartValue(s *Set) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.poseset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PosesetValueN wraps MEOS C function poseset_value_n.
func PosesetValueN(s *Set, n int) (_r0 bool, _r1 *Pose, _err error) {
	var _out_result *C.Pose
	C.meos_errno_reset()
	_cret := C.poseset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pose{_inner: _out_result}, nil
}


// PosesetValues wraps MEOS C function poseset_values.
func PosesetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.poseset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ContainedPoseSet wraps MEOS C function contained_pose_set.
func ContainedPoseSet(pose *Pose, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_pose_set(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetPose wraps MEOS C function contains_set_pose.
func ContainsSetPose(s *Set, pose *Pose) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_pose(s._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionPoseSet wraps MEOS C function intersection_pose_set.
func IntersectionPoseSet(pose *Pose, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_pose_set(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetPose wraps MEOS C function intersection_set_pose.
func IntersectionSetPose(s *Set, pose *Pose) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_pose(s._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusPoseSet wraps MEOS C function minus_pose_set.
func MinusPoseSet(pose *Pose, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_pose_set(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetPose wraps MEOS C function minus_set_pose.
func MinusSetPose(s *Set, pose *Pose) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_pose(s._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PoseUnionTransfn wraps MEOS C function pose_union_transfn.
func PoseUnionTransfn(state *Set, pose *Pose) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_union_transfn(state._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionPoseSet wraps MEOS C function union_pose_set.
func UnionPoseSet(pose *Pose, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_pose_set(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetPose wraps MEOS C function union_set_pose.
func UnionSetPose(s *Set, pose *Pose) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_pose(s._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TposeFromMFJSON wraps MEOS C function tpose_from_mfjson.
func TposeFromMFJSON(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpose_from_mfjson(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeIn wraps MEOS C function tpose_in.
func TposeIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpose_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeinstMake wraps MEOS C function tposeinst_make.
func TposeinstMake(pose *Pose, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tposeinst_make(pose._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TposeFromBaseTemp wraps MEOS C function tpose_from_base_temp.
func TposeFromBaseTemp(pose *Pose, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_from_base_temp(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeseqFromBaseTstzset wraps MEOS C function tposeseq_from_base_tstzset.
func TposeseqFromBaseTstzset(pose *Pose, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tposeseq_from_base_tstzset(pose._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TposeseqFromBaseTstzspan wraps MEOS C function tposeseq_from_base_tstzspan.
func TposeseqFromBaseTstzspan(pose *Pose, s *Span, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tposeseq_from_base_tstzspan(pose._inner, s._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TposeseqsetFromBaseTstzspanset wraps MEOS C function tposeseqset_from_base_tstzspanset.
func TposeseqsetFromBaseTstzspanset(pose *Pose, ss *SpanSet, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tposeseqset_from_base_tstzspanset(pose._inner, ss._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TposeMake wraps MEOS C function tpose_make.
func TposeMake(tpoint *Temporal, ttheta *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_make(tpoint._inner, ttheta._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeToTpoint wraps MEOS C function tpose_to_tpoint.
func TposeToTpoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_to_tpoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeEndValue wraps MEOS C function tpose_end_value.
func TposeEndValue(temp *Temporal) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// TposePoints wraps MEOS C function tpose_points.
func TposePoints(temp *Temporal) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_points(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TposeYaw wraps MEOS C function tpose_yaw.
func TposeYaw(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_yaw(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposePitch wraps MEOS C function tpose_pitch.
func TposePitch(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_pitch(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeRoll wraps MEOS C function tpose_roll.
func TposeRoll(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_roll(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeSpeed wraps MEOS C function tpose_speed.
func TposeSpeed(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_speed(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeAngularSpeed wraps MEOS C function tpose_angular_speed.
func TposeAngularSpeed(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_angular_speed(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeStartValue wraps MEOS C function tpose_start_value.
func TposeStartValue(temp *Temporal) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// TposeTrajectory wraps MEOS C function tpose_trajectory.
func TposeTrajectory(temp *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_trajectory(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TposeValue wraps MEOS C function tpose_value.
func TposeValue(temp *Temporal) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// TposeValueAtTimestamptz wraps MEOS C function tpose_value_at_timestamptz.
func TposeValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 *Pose, _err error) {
	var _out_result *C.Pose
	C.meos_errno_reset()
	_cret := C.tpose_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pose{_inner: _out_result}, nil
}


// TposeValueN wraps MEOS C function tpose_value_n.
func TposeValueN(temp *Temporal, n int) (_r0 bool, _r1 *Pose, _err error) {
	var _out_result *C.Pose
	C.meos_errno_reset()
	_cret := C.tpose_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pose{_inner: _out_result}, nil
}


// TposeValues wraps MEOS C function tpose_values.
func TposeValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TposeUnnest wraps MEOS C function tpose_unnest.
func TposeUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_unnest(temp._inner, (***C.Pose)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TposeAtElevation wraps MEOS C function tpose_at_elevation.
func TposeAtElevation(temp *Temporal, s *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_at_elevation(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeAtGeom wraps MEOS C function tpose_at_geom.
func TposeAtGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_at_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeAtSTBOX wraps MEOS C function tpose_at_stbox.
func TposeAtSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_at_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeAtPose wraps MEOS C function tpose_at_pose.
func TposeAtPose(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_at_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeMinusElevation wraps MEOS C function tpose_minus_elevation.
func TposeMinusElevation(temp *Temporal, s *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_minus_elevation(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeMinusGeom wraps MEOS C function tpose_minus_geom.
func TposeMinusGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_minus_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeMinusPose wraps MEOS C function tpose_minus_pose.
func TposeMinusPose(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_minus_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposeMinusSTBOX wraps MEOS C function tpose_minus_stbox.
func TposeMinusSTBOX(temp *Temporal, box *STBox, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpose_minus_stbox(temp._inner, box._inner, C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTposePose wraps MEOS C function tdistance_tpose_pose.
func TdistanceTposePose(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTposeGeo wraps MEOS C function tdistance_tpose_geo.
func TdistanceTposeGeo(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tpose_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTposeTpose wraps MEOS C function tdistance_tpose_tpose.
func TdistanceTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// NadTposeGeo wraps MEOS C function nad_tpose_geo.
func NadTposeGeo(temp *Temporal, gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tpose_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTposePose wraps MEOS C function nad_tpose_pose.
func NadTposePose(temp *Temporal, pose *Pose) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTposeSTBOX wraps MEOS C function nad_tpose_stbox.
func NadTposeSTBOX(temp *Temporal, box *STBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tpose_stbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTposeTpose wraps MEOS C function nad_tpose_tpose.
func NadTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NaiTposeGeo wraps MEOS C function nai_tpose_geo.
func NaiTposeGeo(temp *Temporal, gs *Geom) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tpose_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// NaiTposePose wraps MEOS C function nai_tpose_pose.
func NaiTposePose(temp *Temporal, pose *Pose) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// NaiTposeTpose wraps MEOS C function nai_tpose_tpose.
func NaiTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.nai_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// ShortestlineTposeGeo wraps MEOS C function shortestline_tpose_geo.
func ShortestlineTposeGeo(temp *Temporal, gs *Geom) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tpose_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// ShortestlineTposePose wraps MEOS C function shortestline_tpose_pose.
func ShortestlineTposePose(temp *Temporal, pose *Pose) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// ShortestlineTposeTpose wraps MEOS C function shortestline_tpose_tpose.
func ShortestlineTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.shortestline_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// AlwaysEqPoseTpose wraps MEOS C function always_eq_pose_tpose.
func AlwaysEqPoseTpose(pose *Pose, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTposePose wraps MEOS C function always_eq_tpose_pose.
func AlwaysEqTposePose(temp *Temporal, pose *Pose) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTposeTpose wraps MEOS C function always_eq_tpose_tpose.
func AlwaysEqTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNePoseTpose wraps MEOS C function always_ne_pose_tpose.
func AlwaysNePoseTpose(pose *Pose, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTposePose wraps MEOS C function always_ne_tpose_pose.
func AlwaysNeTposePose(temp *Temporal, pose *Pose) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTposeTpose wraps MEOS C function always_ne_tpose_tpose.
func AlwaysNeTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqPoseTpose wraps MEOS C function ever_eq_pose_tpose.
func EverEqPoseTpose(pose *Pose, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTposePose wraps MEOS C function ever_eq_tpose_pose.
func EverEqTposePose(temp *Temporal, pose *Pose) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTposeTpose wraps MEOS C function ever_eq_tpose_tpose.
func EverEqTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNePoseTpose wraps MEOS C function ever_ne_pose_tpose.
func EverNePoseTpose(pose *Pose, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTposePose wraps MEOS C function ever_ne_tpose_pose.
func EverNeTposePose(temp *Temporal, pose *Pose) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTposeTpose wraps MEOS C function ever_ne_tpose_tpose.
func EverNeTposeTpose(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tpose_tpose(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqPoseTpose wraps MEOS C function teq_pose_tpose.
func TeqPoseTpose(pose *Pose, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTposePose wraps MEOS C function teq_tpose_pose.
func TeqTposePose(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnePoseTpose wraps MEOS C function tne_pose_tpose.
func TnePoseTpose(pose *Pose, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_pose_tpose(pose._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTposePose wraps MEOS C function tne_tpose_pose.
func TneTposePose(temp *Temporal, pose *Pose) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tpose_pose(temp._inner, pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// PosechainIn wraps MEOS C function posechain_in.
func PosechainIn(str string) (_r0 unsafe.Pointer, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.posechain_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainOut wraps MEOS C function posechain_out.
func PosechainOut(pc unsafe.Pointer, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_out((*C.PoseChain)(unsafe.Pointer(pc)), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PosechainAsText wraps MEOS C function posechain_as_text.
func PosechainAsText(pc unsafe.Pointer, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_as_text((*C.PoseChain)(unsafe.Pointer(pc)), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PosechainAsEWKB wraps MEOS C function posechain_as_ewkb.
func PosechainAsEWKB(pc unsafe.Pointer, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_as_ewkb((*C.PoseChain)(unsafe.Pointer(pc)), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainAsEWKT wraps MEOS C function posechain_as_ewkt.
func PosechainAsEWKT(pc unsafe.Pointer, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_as_ewkt((*C.PoseChain)(unsafe.Pointer(pc)), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PosechainAsWKB wraps MEOS C function posechain_as_wkb.
func PosechainAsWKB(pc unsafe.Pointer, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_as_wkb((*C.PoseChain)(unsafe.Pointer(pc)), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainAsHexewkb wraps MEOS C function posechain_as_hexewkb.
func PosechainAsHexewkb(pc unsafe.Pointer, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_as_hexewkb((*C.PoseChain)(unsafe.Pointer(pc)), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PosechainAsHexwkb wraps MEOS C function posechain_as_hexwkb.
func PosechainAsHexwkb(pc unsafe.Pointer, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_as_hexwkb((*C.PoseChain)(unsafe.Pointer(pc)), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PosechainFromWKB wraps MEOS C function posechain_from_wkb.
func PosechainFromWKB(wkb unsafe.Pointer, size uint) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainFromHexwkb wraps MEOS C function posechain_from_hexwkb.
func PosechainFromHexwkb(hexwkb string) (_r0 unsafe.Pointer, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.posechain_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainMake wraps MEOS C function posechain_make.
func PosechainMake(poses unsafe.Pointer, count int) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_make((**C.Pose)(unsafe.Pointer(poses)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainCopy wraps MEOS C function posechain_copy.
func PosechainCopy(pc unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_copy((*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainAppend wraps MEOS C function posechain_append.
func PosechainAppend(pc unsafe.Pointer, pose *Pose) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_append((*C.PoseChain)(unsafe.Pointer(pc)), pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PoseToPosechain wraps MEOS C function pose_to_posechain.
func PoseToPosechain(pose *Pose) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.pose_to_posechain(pose._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainToPose wraps MEOS C function posechain_to_pose.
func PosechainToPose(pc unsafe.Pointer) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_to_pose((*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PosechainPrefixPose wraps MEOS C function posechain_prefix_pose.
func PosechainPrefixPose(pc unsafe.Pointer, n int) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_prefix_pose((*C.PoseChain)(unsafe.Pointer(pc)), C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PosechainToPoint wraps MEOS C function posechain_to_point.
func PosechainToPoint(pc unsafe.Pointer) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_to_point((*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// PosechainToSTBOX wraps MEOS C function posechain_to_stbox.
func PosechainToSTBOX(pc unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_to_stbox((*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// PosechainTimestamptzToSTBOX wraps MEOS C function posechain_timestamptz_to_stbox.
func PosechainTimestamptzToSTBOX(pc unsafe.Pointer, t int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_timestamptz_to_stbox((*C.PoseChain)(unsafe.Pointer(pc)), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// PosechainTstzspanToSTBOX wraps MEOS C function posechain_tstzspan_to_stbox.
func PosechainTstzspanToSTBOX(pc unsafe.Pointer, s *Span) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_tstzspan_to_stbox((*C.PoseChain)(unsafe.Pointer(pc)), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// PosechainNumPoses wraps MEOS C function posechain_num_poses.
func PosechainNumPoses(pc unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_num_poses((*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// PosechainStartPose wraps MEOS C function posechain_start_pose.
func PosechainStartPose(pc unsafe.Pointer) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_start_pose((*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PosechainEndPose wraps MEOS C function posechain_end_pose.
func PosechainEndPose(pc unsafe.Pointer) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_end_pose((*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PosechainPoseN wraps MEOS C function posechain_pose_n.
func PosechainPoseN(pc unsafe.Pointer, n int) (_r0 *Pose, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_pose_n((*C.PoseChain)(unsafe.Pointer(pc)), C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pose{_inner: _cret}, nil
}


// PosechainPoses wraps MEOS C function posechain_poses.
func PosechainPoses(pc unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_poses((*C.PoseChain)(unsafe.Pointer(pc)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainHash wraps MEOS C function posechain_hash.
func PosechainHash(pc unsafe.Pointer) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_hash((*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// PosechainHashExtended wraps MEOS C function posechain_hash_extended.
func PosechainHashExtended(pc unsafe.Pointer, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_hash_extended((*C.PoseChain)(unsafe.Pointer(pc)), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// PosechainMemSize wraps MEOS C function posechain_mem_size.
func PosechainMemSize(pc unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_mem_size((*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// PosechainRound wraps MEOS C function posechain_round.
func PosechainRound(pc unsafe.Pointer, maxdd int) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_round((*C.PoseChain)(unsafe.Pointer(pc)), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainSRID wraps MEOS C function posechain_srid.
func PosechainSRID(pc unsafe.Pointer) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_srid((*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// PosechainSetSRID wraps MEOS C function posechain_set_srid.
func PosechainSetSRID(pc unsafe.Pointer, srid int32) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_set_srid((*C.PoseChain)(unsafe.Pointer(pc)), C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainTransform wraps MEOS C function posechain_transform.
func PosechainTransform(pc unsafe.Pointer, srid_to int32) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_transform((*C.PoseChain)(unsafe.Pointer(pc)), C.int32_t(srid_to))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainTransformPipeline wraps MEOS C function posechain_transform_pipeline.
func PosechainTransformPipeline(pc unsafe.Pointer, pipelinestr string, srid int32, is_forward bool) (_r0 unsafe.Pointer, _err error) {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	C.meos_errno_reset()
	_cret := C.posechain_transform_pipeline((*C.PoseChain)(unsafe.Pointer(pc)), _c_pipelinestr, C.int32_t(srid), C.bool(is_forward))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TposechainFromGeopose wraps MEOS C function tposechain_from_geopose.
func TposechainFromGeopose(json string) (_r0 *Temporal, _err error) {
	_c_json := C.CString(json)
	defer C.free(unsafe.Pointer(_c_json))
	C.meos_errno_reset()
	_cret := C.tposechain_from_geopose(_c_json)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposechainAsGeopose wraps MEOS C function tposechain_as_geopose.
func TposechainAsGeopose(temp *Temporal, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_as_geopose(temp._inner, C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TposechainarrAsGeopose wraps MEOS C function tposechainarr_as_geopose.
func TposechainarrAsGeopose(temparr unsafe.Pointer, count int, precision int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechainarr_as_geopose((**C.Temporal)(unsafe.Pointer(temparr)), C.int(count), C.int(precision))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PosechainEq wraps MEOS C function posechain_eq.
func PosechainEq(pc1 unsafe.Pointer, pc2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_eq((*C.PoseChain)(unsafe.Pointer(pc1)), (*C.PoseChain)(unsafe.Pointer(pc2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PosechainNe wraps MEOS C function posechain_ne.
func PosechainNe(pc1 unsafe.Pointer, pc2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_ne((*C.PoseChain)(unsafe.Pointer(pc1)), (*C.PoseChain)(unsafe.Pointer(pc2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PosechainCmp wraps MEOS C function posechain_cmp.
func PosechainCmp(pc1 unsafe.Pointer, pc2 unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_cmp((*C.PoseChain)(unsafe.Pointer(pc1)), (*C.PoseChain)(unsafe.Pointer(pc2)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// PosechainLt wraps MEOS C function posechain_lt.
func PosechainLt(pc1 unsafe.Pointer, pc2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_lt((*C.PoseChain)(unsafe.Pointer(pc1)), (*C.PoseChain)(unsafe.Pointer(pc2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PosechainLe wraps MEOS C function posechain_le.
func PosechainLe(pc1 unsafe.Pointer, pc2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_le((*C.PoseChain)(unsafe.Pointer(pc1)), (*C.PoseChain)(unsafe.Pointer(pc2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PosechainGt wraps MEOS C function posechain_gt.
func PosechainGt(pc1 unsafe.Pointer, pc2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_gt((*C.PoseChain)(unsafe.Pointer(pc1)), (*C.PoseChain)(unsafe.Pointer(pc2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PosechainGe wraps MEOS C function posechain_ge.
func PosechainGe(pc1 unsafe.Pointer, pc2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_ge((*C.PoseChain)(unsafe.Pointer(pc1)), (*C.PoseChain)(unsafe.Pointer(pc2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PosechainsetIn wraps MEOS C function posechainset_in.
func PosechainsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.posechainset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PosechainsetOut wraps MEOS C function posechainset_out.
func PosechainsetOut(s *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.posechainset_out(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PosechainsetMake wraps MEOS C function posechainset_make.
func PosechainsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.posechainset_make((**C.PoseChain)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PosechainToSet wraps MEOS C function posechain_to_set.
func PosechainToSet(pc unsafe.Pointer) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_to_set((*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PosechainsetEndValue wraps MEOS C function posechainset_end_value.
func PosechainsetEndValue(s *Set) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechainset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PosechainsetStartValue wraps MEOS C function posechainset_start_value.
func PosechainsetStartValue(s *Set) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechainset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TODO posechainset_value_n: unhandled OUTPUT_SCALAR shape PoseChain **
// func PosechainsetValueN(...) { /* not yet handled by codegen */ }


// PosechainsetValues wraps MEOS C function posechainset_values.
func PosechainsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.posechainset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ContainedPosechainSet wraps MEOS C function contained_posechain_set.
func ContainedPosechainSet(pc unsafe.Pointer, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_posechain_set((*C.PoseChain)(unsafe.Pointer(pc)), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetPosechain wraps MEOS C function contains_set_posechain.
func ContainsSetPosechain(s *Set, pc unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_posechain(s._inner, (*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionPosechainSet wraps MEOS C function intersection_posechain_set.
func IntersectionPosechainSet(pc unsafe.Pointer, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_posechain_set((*C.PoseChain)(unsafe.Pointer(pc)), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetPosechain wraps MEOS C function intersection_set_posechain.
func IntersectionSetPosechain(s *Set, pc unsafe.Pointer) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_posechain(s._inner, (*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusPosechainSet wraps MEOS C function minus_posechain_set.
func MinusPosechainSet(pc unsafe.Pointer, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_posechain_set((*C.PoseChain)(unsafe.Pointer(pc)), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetPosechain wraps MEOS C function minus_set_posechain.
func MinusSetPosechain(s *Set, pc unsafe.Pointer) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_posechain(s._inner, (*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PosechainUnionTransfn wraps MEOS C function posechain_union_transfn.
func PosechainUnionTransfn(state *Set, pc unsafe.Pointer) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.posechain_union_transfn(state._inner, (*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionPosechainSet wraps MEOS C function union_posechain_set.
func UnionPosechainSet(pc unsafe.Pointer, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_posechain_set((*C.PoseChain)(unsafe.Pointer(pc)), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetPosechain wraps MEOS C function union_set_posechain.
func UnionSetPosechain(s *Set, pc unsafe.Pointer) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_posechain(s._inner, (*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TposechainIn wraps MEOS C function tposechain_in.
func TposechainIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tposechain_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposechainFromMFJSON wraps MEOS C function tposechain_from_mfjson.
func TposechainFromMFJSON(mfjson string) (_r0 *Temporal, _err error) {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	C.meos_errno_reset()
	_cret := C.tposechain_from_mfjson(_c_mfjson)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposechaininstMake wraps MEOS C function tposechaininst_make.
func TposechaininstMake(pc unsafe.Pointer, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechaininst_make((*C.PoseChain)(unsafe.Pointer(pc)), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TposechainFromBaseTemp wraps MEOS C function tposechain_from_base_temp.
func TposechainFromBaseTemp(pc unsafe.Pointer, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_from_base_temp((*C.PoseChain)(unsafe.Pointer(pc)), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposechainseqFromBaseTstzset wraps MEOS C function tposechainseq_from_base_tstzset.
func TposechainseqFromBaseTstzset(pc unsafe.Pointer, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechainseq_from_base_tstzset((*C.PoseChain)(unsafe.Pointer(pc)), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TposechainseqFromBaseTstzspan wraps MEOS C function tposechainseq_from_base_tstzspan.
func TposechainseqFromBaseTstzspan(pc unsafe.Pointer, s *Span, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechainseq_from_base_tstzspan((*C.PoseChain)(unsafe.Pointer(pc)), s._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TposechainseqsetFromBaseTstzspanset wraps MEOS C function tposechainseqset_from_base_tstzspanset.
func TposechainseqsetFromBaseTstzspanset(pc unsafe.Pointer, ss *SpanSet, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechainseqset_from_base_tstzspanset((*C.PoseChain)(unsafe.Pointer(pc)), ss._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TposechainToTpose wraps MEOS C function tposechain_to_tpose.
func TposechainToTpose(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_to_tpose(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposechainEndValue wraps MEOS C function tposechain_end_value.
func TposechainEndValue(temp *Temporal) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TposechainNumPoses wraps MEOS C function tposechain_num_poses.
func TposechainNumPoses(temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_num_poses(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TposechainStartValue wraps MEOS C function tposechain_start_value.
func TposechainStartValue(temp *Temporal) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TposechainValue wraps MEOS C function tposechain_value.
func TposechainValue(temp *Temporal) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TODO tposechain_value_n: unhandled OUTPUT_SCALAR shape PoseChain **
// func TposechainValueN(...) { /* not yet handled by codegen */ }


// TposechainValues wraps MEOS C function tposechain_values.
func TposechainValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TposechainUnnest wraps MEOS C function tposechain_unnest.
func TposechainUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_unnest(temp._inner, (***C.PoseChain)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TposechainRotate wraps MEOS C function tposechain_rotate.
func TposechainRotate(temp *Temporal, angle float64, x0 float64, y0 float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_rotate(temp._inner, C.double(angle), C.double(x0), C.double(y0))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposechainRotateX wraps MEOS C function tposechain_rotate_x.
func TposechainRotateX(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_rotate_x(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposechainRotateY wraps MEOS C function tposechain_rotate_y.
func TposechainRotateY(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_rotate_y(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposechainRotateZ wraps MEOS C function tposechain_rotate_z.
func TposechainRotateZ(temp *Temporal, angle float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_rotate_z(temp._inner, C.double(angle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposechainTranslate wraps MEOS C function tposechain_translate.
func TposechainTranslate(temp *Temporal, deltax float64, deltay float64, deltaz float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_translate(temp._inner, C.double(deltax), C.double(deltay), C.double(deltaz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposechainAtValue wraps MEOS C function tposechain_at_value.
func TposechainAtValue(temp *Temporal, pc unsafe.Pointer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_at_value(temp._inner, (*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TposechainMinusValue wraps MEOS C function tposechain_minus_value.
func TposechainMinusValue(temp *Temporal, pc unsafe.Pointer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tposechain_minus_value(temp._inner, (*C.PoseChain)(unsafe.Pointer(pc)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AlwaysEqPosechainTposechain wraps MEOS C function always_eq_posechain_tposechain.
func AlwaysEqPosechainTposechain(posechain unsafe.Pointer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_posechain_tposechain((*C.PoseChain)(unsafe.Pointer(posechain)), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTposechainPosechain wraps MEOS C function always_eq_tposechain_posechain.
func AlwaysEqTposechainPosechain(temp *Temporal, posechain unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tposechain_posechain(temp._inner, (*C.PoseChain)(unsafe.Pointer(posechain)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTposechainTposechain wraps MEOS C function always_eq_tposechain_tposechain.
func AlwaysEqTposechainTposechain(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tposechain_tposechain(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNePosechainTposechain wraps MEOS C function always_ne_posechain_tposechain.
func AlwaysNePosechainTposechain(posechain unsafe.Pointer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_posechain_tposechain((*C.PoseChain)(unsafe.Pointer(posechain)), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTposechainPosechain wraps MEOS C function always_ne_tposechain_posechain.
func AlwaysNeTposechainPosechain(temp *Temporal, posechain unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tposechain_posechain(temp._inner, (*C.PoseChain)(unsafe.Pointer(posechain)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTposechainTposechain wraps MEOS C function always_ne_tposechain_tposechain.
func AlwaysNeTposechainTposechain(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tposechain_tposechain(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqPosechainTposechain wraps MEOS C function ever_eq_posechain_tposechain.
func EverEqPosechainTposechain(posechain unsafe.Pointer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_posechain_tposechain((*C.PoseChain)(unsafe.Pointer(posechain)), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTposechainPosechain wraps MEOS C function ever_eq_tposechain_posechain.
func EverEqTposechainPosechain(temp *Temporal, posechain unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tposechain_posechain(temp._inner, (*C.PoseChain)(unsafe.Pointer(posechain)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTposechainTposechain wraps MEOS C function ever_eq_tposechain_tposechain.
func EverEqTposechainTposechain(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tposechain_tposechain(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNePosechainTposechain wraps MEOS C function ever_ne_posechain_tposechain.
func EverNePosechainTposechain(posechain unsafe.Pointer, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_posechain_tposechain((*C.PoseChain)(unsafe.Pointer(posechain)), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTposechainPosechain wraps MEOS C function ever_ne_tposechain_posechain.
func EverNeTposechainPosechain(temp *Temporal, posechain unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tposechain_posechain(temp._inner, (*C.PoseChain)(unsafe.Pointer(posechain)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTposechainTposechain wraps MEOS C function ever_ne_tposechain_tposechain.
func EverNeTposechainTposechain(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tposechain_tposechain(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqPosechainTposechain wraps MEOS C function teq_posechain_tposechain.
func TeqPosechainTposechain(posechain unsafe.Pointer, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_posechain_tposechain((*C.PoseChain)(unsafe.Pointer(posechain)), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTposechainPosechain wraps MEOS C function teq_tposechain_posechain.
func TeqTposechainPosechain(temp *Temporal, posechain unsafe.Pointer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tposechain_posechain(temp._inner, (*C.PoseChain)(unsafe.Pointer(posechain)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnePosechainTposechain wraps MEOS C function tne_posechain_tposechain.
func TnePosechainTposechain(posechain unsafe.Pointer, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_posechain_tposechain((*C.PoseChain)(unsafe.Pointer(posechain)), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTposechainPosechain wraps MEOS C function tne_tposechain_posechain.
func TneTposechainPosechain(temp *Temporal, posechain unsafe.Pointer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tposechain_posechain(temp._inner, (*C.PoseChain)(unsafe.Pointer(posechain)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}

