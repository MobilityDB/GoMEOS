package functions

/*
#include <stddef.h>
#include "meos.h"
#include "meos_catalog.h"
#include "meos_cbuffer.h"
#include "meos_cellindex.h"
#include "meos_geo.h"
#include "meos_h3.h"
#include "meos_internal.h"
#include "meos_internal_geo.h"
#include "meos_json.h"
#include "meos_npoint.h"
#include "meos_pointcloud.h"
#include "meos_pose.h"
#include "meos_quadbin.h"
#include "meos_raster.h"
#include "meos_rgeo.h"
#include "meos_s2cell.h"
#include "pg_basetypes.h"
#include "pg_bool.h"
#include "pg_date.h"
#include "pg_float.h"
#include "pg_int.h"
#include "pg_interval.h"
#include "pg_text.h"
#include "pg_time.h"
#include "pg_timestamp.h"

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

// MeosProjGetContext wraps MEOS C function meos_proj_get_context.
func MeosProjGetContext() (_r0 *PJContext, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_proj_get_context()
	if _err = meosError(); _err != nil {
		return
	}
	return &PJContext{_inner: _cret}, nil
}


// PointRound wraps MEOS C function point_round.
func PointRound(gs *Geom, maxdd int) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.point_round(gs._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// STBOXSet wraps MEOS C function stbox_set.
func STBOXSet(hasx bool, hasz bool, geodetic bool, srid int32, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, s *Span) (_r0 *STBox, _err error) {
	var _out_result C.STBox
	C.meos_errno_reset()
	C.stbox_set(C.bool(hasx), C.bool(hasz), C.bool(geodetic), C.int32(srid), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), s._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: &_out_result}, nil
}


// GboxSetSTBOX wraps MEOS C function gbox_set_stbox.
func GboxSetSTBOX(box *GBox, srid int32) (_r0 *STBox, _err error) {
	var _out_result C.STBox
	C.meos_errno_reset()
	C.gbox_set_stbox(box._inner, C.int32_t(srid), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: &_out_result}, nil
}


// GeoSetSTBOX wraps MEOS C function geo_set_stbox.
func GeoSetSTBOX(gs *Geom) (_r0 bool, _r1 *STBox, _err error) {
	var _out_result C.STBox
	C.meos_errno_reset()
	_cret := C.geo_set_stbox(gs._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &STBox{_inner: &_out_result}, nil
}


// SpatialsetSetSTBOX wraps MEOS C function spatialset_set_stbox.
func SpatialsetSetSTBOX(set *Set) (_r0 *STBox, _err error) {
	var _out_result C.STBox
	C.meos_errno_reset()
	C.spatialset_set_stbox(set._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: &_out_result}, nil
}


// STBOXSetBox3d wraps MEOS C function stbox_set_box3d.
func STBOXSetBox3d(box *STBox) (_r0 *Box3D, _err error) {
	var _out_result C.BOX3D
	C.meos_errno_reset()
	C.stbox_set_box3d(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &Box3D{_inner: &_out_result}, nil
}


// STBOXSetGbox wraps MEOS C function stbox_set_gbox.
func STBOXSetGbox(box *STBox) (_r0 *GBox, _err error) {
	var _out_result C.GBOX
	C.meos_errno_reset()
	C.stbox_set_gbox(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &GBox{_inner: &_out_result}, nil
}


// TstzsetSetSTBOX wraps MEOS C function tstzset_set_stbox.
func TstzsetSetSTBOX(s *Set) (_r0 *STBox, _err error) {
	var _out_result C.STBox
	C.meos_errno_reset()
	C.tstzset_set_stbox(s._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: &_out_result}, nil
}


// TimestamptzSetSTBOX wraps MEOS C function timestamptz_set_stbox.
func TimestamptzSetSTBOX(t int64) (_r0 *STBox, _err error) {
	var _out_result C.STBox
	C.meos_errno_reset()
	C.timestamptz_set_stbox(C.TimestampTz(t), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: &_out_result}, nil
}


// TstzspanSetSTBOX wraps MEOS C function tstzspan_set_stbox.
func TstzspanSetSTBOX(s *Span) (_r0 *STBox, _err error) {
	var _out_result C.STBox
	C.meos_errno_reset()
	C.tstzspan_set_stbox(s._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: &_out_result}, nil
}


// TstzspansetSetSTBOX wraps MEOS C function tstzspanset_set_stbox.
func TstzspansetSetSTBOX(s *SpanSet) (_r0 *STBox, _err error) {
	var _out_result C.STBox
	C.meos_errno_reset()
	C.tstzspanset_set_stbox(s._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: &_out_result}, nil
}


// STBOXExpand wraps MEOS C function stbox_expand.
func STBOXExpand(box1 *STBox, box2 *STBox) (_err error) {
	C.meos_errno_reset()
	C.stbox_expand(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// STBOXExpandSpaceSet wraps MEOS C function stbox_expand_space_set.
func STBOXExpandSpaceSet(box *STBox, d float64) (_r0 bool, _r1 *STBox, _err error) {
	var _out_result C.STBox
	C.meos_errno_reset()
	_cret := C.stbox_expand_space_set(box._inner, C.double(d), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &STBox{_inner: &_out_result}, nil
}


// STBOXContains wraps MEOS C function stbox_contains.
func STBOXContains(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_contains(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXContained wraps MEOS C function stbox_contained.
func STBOXContained(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_contained(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXOverlaps wraps MEOS C function stbox_overlaps.
func STBOXOverlaps(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_overlaps(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXSame wraps MEOS C function stbox_same.
func STBOXSame(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_same(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXAdjacent wraps MEOS C function stbox_adjacent.
func STBOXAdjacent(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_adjacent(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXLeft wraps MEOS C function stbox_left.
func STBOXLeft(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_left(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXRight wraps MEOS C function stbox_right.
func STBOXRight(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_right(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXOverleft wraps MEOS C function stbox_overleft.
func STBOXOverleft(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_overleft(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXOverright wraps MEOS C function stbox_overright.
func STBOXOverright(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_overright(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXBelow wraps MEOS C function stbox_below.
func STBOXBelow(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_below(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXAbove wraps MEOS C function stbox_above.
func STBOXAbove(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_above(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXOverbelow wraps MEOS C function stbox_overbelow.
func STBOXOverbelow(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_overbelow(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXOverabove wraps MEOS C function stbox_overabove.
func STBOXOverabove(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_overabove(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXFront wraps MEOS C function stbox_front.
func STBOXFront(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_front(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXBack wraps MEOS C function stbox_back.
func STBOXBack(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_back(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXOverfront wraps MEOS C function stbox_overfront.
func STBOXOverfront(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_overfront(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXOverback wraps MEOS C function stbox_overback.
func STBOXOverback(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_overback(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXBefore wraps MEOS C function stbox_before.
func STBOXBefore(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_before(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXAfter wraps MEOS C function stbox_after.
func STBOXAfter(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_after(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXOverbefore wraps MEOS C function stbox_overbefore.
func STBOXOverbefore(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_overbefore(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// STBOXOverafter wraps MEOS C function stbox_overafter.
func STBOXOverafter(box1 *STBox, box2 *STBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_overafter(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// InterSTBOXSTBOX wraps MEOS C function inter_stbox_stbox.
func InterSTBOXSTBOX(box1 *STBox, box2 *STBox) (_r0 bool, _r1 *STBox, _err error) {
	var _out_result C.STBox
	C.meos_errno_reset()
	_cret := C.inter_stbox_stbox(box1._inner, box2._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &STBox{_inner: &_out_result}, nil
}


// STBOXGeo wraps MEOS C function stbox_geo.
func STBOXGeo(box *STBox) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_geo(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// STBOXSortkeyIntl wraps MEOS C function stbox_sortkey_intl.
func STBOXSortkeyIntl(box *STBox, size unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_sortkey_intl(box._inner, (*C.size_t)(unsafe.Pointer(size)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TgeogpointinstIn wraps MEOS C function tgeogpointinst_in.
func TgeogpointinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeogpointinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TgeogpointseqIn wraps MEOS C function tgeogpointseq_in.
func TgeogpointseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeogpointseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TgeogpointseqsetIn wraps MEOS C function tgeogpointseqset_in.
func TgeogpointseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeogpointseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TgeompointinstIn wraps MEOS C function tgeompointinst_in.
func TgeompointinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeompointinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TgeompointseqIn wraps MEOS C function tgeompointseq_in.
func TgeompointseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeompointseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TgeompointseqsetIn wraps MEOS C function tgeompointseqset_in.
func TgeompointseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeompointseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TgeographyinstIn wraps MEOS C function tgeographyinst_in.
func TgeographyinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeographyinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TgeographyseqIn wraps MEOS C function tgeographyseq_in.
func TgeographyseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeographyseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TgeographyseqsetIn wraps MEOS C function tgeographyseqset_in.
func TgeographyseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeographyseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TgeometryinstIn wraps MEOS C function tgeometryinst_in.
func TgeometryinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeometryinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TgeometryseqIn wraps MEOS C function tgeometryseq_in.
func TgeometryseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeometryseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TgeometryseqsetIn wraps MEOS C function tgeometryseqset_in.
func TgeometryseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tgeometryseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TspatialSetSTBOX wraps MEOS C function tspatial_set_stbox.
func TspatialSetSTBOX(temp *Temporal) (_r0 *STBox, _err error) {
	var _out_result C.STBox
	C.meos_errno_reset()
	C.tspatial_set_stbox(temp._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: &_out_result}, nil
}


// TgeoinstSetSTBOX wraps MEOS C function tgeoinst_set_stbox.
func TgeoinstSetSTBOX(inst *TInstant, box *STBox) (_err error) {
	C.meos_errno_reset()
	C.tgeoinst_set_stbox(inst._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TspatialseqSetSTBOX wraps MEOS C function tspatialseq_set_stbox.
func TspatialseqSetSTBOX(seq *TSequence, box *STBox) (_err error) {
	C.meos_errno_reset()
	C.tspatialseq_set_stbox(seq._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TspatialseqsetSetSTBOX wraps MEOS C function tspatialseqset_set_stbox.
func TspatialseqsetSetSTBOX(ss *TSequenceSet, box *STBox) (_err error) {
	C.meos_errno_reset()
	C.tspatialseqset_set_stbox(ss._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TgeoRestrictElevation wraps MEOS C function tgeo_restrict_elevation.
func TgeoRestrictElevation(temp *Temporal, s *Span, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_restrict_elevation(temp._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoRestrictElevationIntl wraps MEOS C function tgeo_restrict_elevation_intl.
func TgeoRestrictElevationIntl(temp *Temporal, s *Span, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_restrict_elevation_intl(temp._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoRestrictGeom wraps MEOS C function tgeo_restrict_geom.
func TgeoRestrictGeom(temp *Temporal, gs *Geom, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_restrict_geom(temp._inner, gs._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoRestrictSTBOX wraps MEOS C function tgeo_restrict_stbox.
func TgeoRestrictSTBOX(temp *Temporal, box *STBox, border_inc bool, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_restrict_stbox(temp._inner, box._inner, C.bool(border_inc), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoinstRestrictGeom wraps MEOS C function tgeoinst_restrict_geom.
func TgeoinstRestrictGeom(inst *TInstant, gs *Geom, atfunc bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoinst_restrict_geom(inst._inner, gs._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TgeoinstRestrictSTBOX wraps MEOS C function tgeoinst_restrict_stbox.
func TgeoinstRestrictSTBOX(inst *TInstant, box *STBox, border_inc bool, atfunc bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoinst_restrict_stbox(inst._inner, box._inner, C.bool(border_inc), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TgeoseqRestrictGeom wraps MEOS C function tgeoseq_restrict_geom.
func TgeoseqRestrictGeom(seq *TSequence, gs *Geom, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoseq_restrict_geom(seq._inner, gs._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoseqRestrictSTBOX wraps MEOS C function tgeoseq_restrict_stbox.
func TgeoseqRestrictSTBOX(seq *TSequence, box *STBox, border_inc bool, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoseq_restrict_stbox(seq._inner, box._inner, C.bool(border_inc), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoseqsetRestrictGeom wraps MEOS C function tgeoseqset_restrict_geom.
func TgeoseqsetRestrictGeom(ss *TSequenceSet, gs *Geom, atfunc bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoseqset_restrict_geom(ss._inner, gs._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TgeoseqsetRestrictSTBOX wraps MEOS C function tgeoseqset_restrict_stbox.
func TgeoseqsetRestrictSTBOX(ss *TSequenceSet, box *STBox, border_inc bool, atfunc bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoseqset_restrict_stbox(ss._inner, box._inner, C.bool(border_inc), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// GeoEdgeCtxMake wraps MEOS C function geo_edge_ctx_make.
func GeoEdgeCtxMake(gs *Geom) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_edge_ctx_make(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// GeoIsPlanarLinear wraps MEOS C function geo_is_planar_linear.
func GeoIsPlanarLinear(gs *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_is_planar_linear(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoClipSubject wraps MEOS C function geo_clip_subject.
func GeoClipSubject(gs *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_clip_subject(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoIsPlanarAreal wraps MEOS C function geo_is_planar_areal.
func GeoIsPlanarAreal(gs *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_is_planar_areal(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoEveryPartBoundsArea wraps MEOS C function geo_every_part_bounds_area.
func GeoEveryPartBoundsArea(gs *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_every_part_bounds_area(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoIsPointSet wraps MEOS C function geo_is_point_set.
func GeoIsPointSet(gs *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_is_point_set(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoMeosCoverage wraps MEOS C function geo_meos_coverage.
func GeoMeosCoverage(gs *Geom) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_meos_coverage(gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// GeoPointsCovered wraps MEOS C function geo_points_covered.
func GeoPointsCovered(pts *Geom, gs *Geom, covered bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_points_covered(pts._inner, gs._inner, C.bool(covered))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoClipLinearGeom wraps MEOS C function geo_clip_linear_geom.
func GeoClipLinearGeom(line *Geom, gs *Geom, inside bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_clip_linear_geom(line._inner, gs._inner, C.bool(inside))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// GeoEdgeCtxFree wraps MEOS C function geo_edge_ctx_free.
func GeoEdgeCtxFree(ctx unsafe.Pointer) (_err error) {
	C.meos_errno_reset()
	C.geo_edge_ctx_free(unsafe.Pointer(ctx))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// GeoIntersects2d wraps MEOS C function geo_intersects2d.
func GeoIntersects2d(gs1 *Geom, gs2 *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_intersects2d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoIntersects2dCtx wraps MEOS C function geo_intersects2d_ctx.
func GeoIntersects2dCtx(gs *Geom, ctx unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_intersects2d_ctx(gs._inner, unsafe.Pointer(ctx))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// GeoCovers2d wraps MEOS C function geo_covers2d.
func GeoCovers2d(gs1 *Geom, gs2 *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_covers2d(gs1._inner, gs2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpointLinearInterGeom wraps MEOS C function tpoint_linear_inter_geom.
func TpointLinearInterGeom(temp *Temporal, gs *Geom, clip bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_linear_inter_geom(temp._inner, gs._inner, C.bool(clip))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointLinearInterGeomCtx wraps MEOS C function tpoint_linear_inter_geom_ctx.
func TpointLinearInterGeomCtx(temp *Temporal, ctx unsafe.Pointer, clip bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_linear_inter_geom_ctx(temp._inner, unsafe.Pointer(ctx), C.bool(clip))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointLinearDwithinGeom wraps MEOS C function tpoint_linear_dwithin_geom.
func TpointLinearDwithinGeom(temp *Temporal, gs *Geom, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_linear_dwithin_geom(temp._inner, gs._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointLinearDwithinGeomCtx wraps MEOS C function tpoint_linear_dwithin_geom_ctx.
func TpointLinearDwithinGeomCtx(temp *Temporal, ctx unsafe.Pointer, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_linear_dwithin_geom_ctx(temp._inner, unsafe.Pointer(ctx), C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointLinearDistanceGeom wraps MEOS C function tpoint_linear_distance_geom.
func TpointLinearDistanceGeom(temp *Temporal, gs *Geom) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_linear_distance_geom(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpointLinearRestrictGeom wraps MEOS C function tpoint_linear_restrict_geom.
func TpointLinearRestrictGeom(temp *Temporal, gs *Geom, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_linear_restrict_geom(temp._inner, gs._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// GeomMeosCoverage wraps MEOS C function geom_meos_coverage.
func GeomMeosCoverage(geom unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.geom_meos_coverage((*C.LWGEOM)(unsafe.Pointer(geom)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// STBOXNad wraps MEOS C function stbox_nad.
func STBOXNad(box1 *STBox, box2 *STBox, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.stbox_nad(box1._inner, box2._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TspatialinstSRID wraps MEOS C function tspatialinst_srid.
func TspatialinstSRID(inst *TInstant) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tspatialinst_srid(inst._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TpointseqAzimuth wraps MEOS C function tpointseq_azimuth.
func TpointseqAzimuth(seq *TSequence) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseq_azimuth(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TpointseqCumulativeLength wraps MEOS C function tpointseq_cumulative_length.
func TpointseqCumulativeLength(seq *TSequence, prevlength float64, spheroid bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseq_cumulative_length(seq._inner, C.double(prevlength), C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TpointseqIsSimple wraps MEOS C function tpointseq_is_simple.
func TpointseqIsSimple(seq *TSequence) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseq_is_simple(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpointseqLength wraps MEOS C function tpointseq_length.
func TpointseqLength(seq *TSequence, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseq_length(seq._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TpointseqLinearTrajectory wraps MEOS C function tpointseq_linear_trajectory.
func TpointseqLinearTrajectory(seq *TSequence, unary_union bool) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseq_linear_trajectory(seq._inner, C.bool(unary_union))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TgeoseqStboxes wraps MEOS C function tgeoseq_stboxes.
func TgeoseqStboxes(seq *TSequence, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoseq_stboxes(seq._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TgeoseqSplitNStboxes wraps MEOS C function tgeoseq_split_n_stboxes.
func TgeoseqSplitNStboxes(seq *TSequence, max_count int, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoseq_split_n_stboxes(seq._inner, C.int(max_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TpointseqsetAzimuth wraps MEOS C function tpointseqset_azimuth.
func TpointseqsetAzimuth(ss *TSequenceSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseqset_azimuth(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TpointseqsetCumulativeLength wraps MEOS C function tpointseqset_cumulative_length.
func TpointseqsetCumulativeLength(ss *TSequenceSet, spheroid bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseqset_cumulative_length(ss._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TpointseqsetIsSimple wraps MEOS C function tpointseqset_is_simple.
func TpointseqsetIsSimple(ss *TSequenceSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseqset_is_simple(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpointseqsetLength wraps MEOS C function tpointseqset_length.
func TpointseqsetLength(ss *TSequenceSet, spheroid bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseqset_length(ss._inner, C.bool(spheroid))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TgeoseqsetStboxes wraps MEOS C function tgeoseqset_stboxes.
func TgeoseqsetStboxes(ss *TSequenceSet, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoseqset_stboxes(ss._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TgeoseqsetSplitNStboxes wraps MEOS C function tgeoseqset_split_n_stboxes.
func TgeoseqsetSplitNStboxes(ss *TSequenceSet, max_count int, count unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeoseqset_split_n_stboxes(ss._inner, C.int(max_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TpointGetCoord wraps MEOS C function tpoint_get_coord.
func TpointGetCoord(temp *Temporal, coord int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpoint_get_coord(temp._inner, C.int(coord))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeominstTgeoginst wraps MEOS C function tgeominst_tgeoginst.
func TgeominstTgeoginst(inst *TInstant, oper bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeominst_tgeoginst(inst._inner, C.bool(oper))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TgeomseqTgeogseq wraps MEOS C function tgeomseq_tgeogseq.
func TgeomseqTgeogseq(seq *TSequence, oper bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeomseq_tgeogseq(seq._inner, C.bool(oper))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TgeomseqsetTgeogseqset wraps MEOS C function tgeomseqset_tgeogseqset.
func TgeomseqsetTgeogseqset(ss *TSequenceSet, oper bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeomseqset_tgeogseqset(ss._inner, C.bool(oper))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TgeomTgeog wraps MEOS C function tgeom_tgeog.
func TgeomTgeog(temp *Temporal, oper bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeom_tgeog(temp._inner, C.bool(oper))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeoTpoint wraps MEOS C function tgeo_tpoint.
func TgeoTpoint(temp *Temporal, oper bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeo_tpoint(temp._inner, C.bool(oper))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TspatialinstSetSRID wraps MEOS C function tspatialinst_set_srid.
func TspatialinstSetSRID(inst *TInstant, srid int32) (_err error) {
	C.meos_errno_reset()
	C.tspatialinst_set_srid(inst._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TpointseqMakeSimple wraps MEOS C function tpointseq_make_simple.
func TpointseqMakeSimple(seq *TSequence, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseq_make_simple(seq._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TspatialseqSetSRID wraps MEOS C function tspatialseq_set_srid.
func TspatialseqSetSRID(seq *TSequence, srid int32) (_err error) {
	C.meos_errno_reset()
	C.tspatialseq_set_srid(seq._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TpointseqsetMakeSimple wraps MEOS C function tpointseqset_make_simple.
func TpointseqsetMakeSimple(ss *TSequenceSet, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseqset_make_simple(ss._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TspatialseqsetSetSRID wraps MEOS C function tspatialseqset_set_srid.
func TspatialseqsetSetSRID(ss *TSequenceSet, srid int32) (_err error) {
	C.meos_errno_reset()
	C.tspatialseqset_set_srid(ss._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TpointseqTwcentroid wraps MEOS C function tpointseq_twcentroid.
func TpointseqTwcentroid(seq *TSequence) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseq_twcentroid(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// TpointseqsetTwcentroid wraps MEOS C function tpointseqset_twcentroid.
func TpointseqsetTwcentroid(ss *TSequenceSet) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.tpointseqset_twcentroid(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}

