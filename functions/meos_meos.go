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
#define gunion_int_set union_int_set
#define gunion_int_span union_int_span
#define gunion_int_spanset union_int_spanset
#define gunion_jsonb_set union_jsonb_set
#define gunion_npoint_set union_npoint_set
#define gunion_pcpatch_set union_pcpatch_set
#define gunion_pcpoint_set union_pcpoint_set
#define gunion_pose_set union_pose_set
#define gunion_posechain_set union_posechain_set
#define gunion_set_bigint union_set_bigint
#define gunion_set_cbuffer union_set_cbuffer
#define gunion_set_date union_set_date
#define gunion_set_float union_set_float
#define gunion_set_geo union_set_geo
#define gunion_set_int union_set_int
#define gunion_set_jsonb union_set_jsonb
#define gunion_set_npoint union_set_npoint
#define gunion_set_pcpatch union_set_pcpatch
#define gunion_set_pcpoint union_set_pcpoint
#define gunion_set_pose union_set_pose
#define gunion_set_posechain union_set_posechain
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

// BoolIn wraps MEOS C function bool_in.
func BoolIn(str string) (_r0 bool, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.bool_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BoolOut wraps MEOS C function bool_out.
func BoolOut(b bool) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.bool_out(C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// Float8Out wraps MEOS C function float8_out.
func Float8Out(num float64, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.float8_out(C.double(num), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// DateIn wraps MEOS C function date_in.
func DateIn(str string) (_r0 int32, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.date_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DateOut wraps MEOS C function date_out.
func DateOut(date int32) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.date_out(C.DateADT(date))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// IntervalCmp wraps MEOS C function interval_cmp.
func IntervalCmp(interv1 *Interval, interv2 *Interval) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.interval_cmp(interv1._inner, interv2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// IntervalIn wraps MEOS C function interval_in.
func IntervalIn(str string, typmod int32) (_r0 *Interval, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.interval_in(_c_str, C.int32(typmod))
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// IntervalOut wraps MEOS C function interval_out.
func IntervalOut(interv *Interval) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.interval_out(interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TimeIn wraps MEOS C function time_in.
func TimeIn(str string, typmod int32) (_r0 int64, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.time_in(_c_str, C.int32(typmod))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimeOut wraps MEOS C function time_out.
func TimeOut(time int64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.time_out(C.TimeADT(time))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TimestampIn wraps MEOS C function timestamp_in.
func TimestampIn(str string, typmod int32) (_r0 int64, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.timestamp_in(_c_str, C.int32(typmod))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestampOut wraps MEOS C function timestamp_out.
func TimestampOut(ts int64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamp_out(C.Timestamp(ts))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TimestamptzIn wraps MEOS C function timestamptz_in.
func TimestamptzIn(str string, typmod int32) (_r0 int64, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.timestamptz_in(_c_str, C.int32(typmod))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TimestamptzOut wraps MEOS C function timestamptz_out.
func TimestamptzOut(tstz int64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_out(C.TimestampTz(tstz))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// CstringToText wraps MEOS C function cstring_to_text.
func CstringToText(str string) (_r0 string, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.cstring_to_text(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextToCstring wraps MEOS C function text_to_cstring.
func TextToCstring(txt string) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_to_cstring(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TextIn wraps MEOS C function text_in.
func TextIn(str string) (_r0 string, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.text_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextOut wraps MEOS C function text_out.
func TextOut(txt string) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_out(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TODO text_cmp: unsupported param Oid
// func TextCmp(...) { /* not yet handled by codegen */ }


// TextCopy wraps MEOS C function text_copy.
func TextCopy(txt string) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_copy(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextInitcap wraps MEOS C function text_initcap.
func TextInitcap(txt string) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_initcap(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextLower wraps MEOS C function text_lower.
func TextLower(txt string) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_lower(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextUpper wraps MEOS C function text_upper.
func TextUpper(txt string) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_upper(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextcatTextText wraps MEOS C function textcat_text_text.
func TextcatTextText(txt1 string, txt2 string) (_r0 string, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.textcat_text_text(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// MeosArrayCreate wraps MEOS C function meos_array_create.
func MeosArrayCreate(elem_size int) (_r0 *MeosArray, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_array_create(C.int(elem_size))
	if _err = meosError(); _err != nil {
		return
	}
	return &MeosArray{_inner: _cret}, nil
}


// MeosArrayAdd wraps MEOS C function meos_array_add.
func MeosArrayAdd(array *MeosArray, value unsafe.Pointer) (_err error) {
	C.meos_errno_reset()
	C.meos_array_add(array._inner, unsafe.Pointer(value))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosArrayGet wraps MEOS C function meos_array_get.
func MeosArrayGet(array *MeosArray, n int) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_array_get(array._inner, C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// MeosArrayCount wraps MEOS C function meos_array_count.
func MeosArrayCount(array *MeosArray) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_array_count(array._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// MeosArrayReset wraps MEOS C function meos_array_reset.
func MeosArrayReset(array *MeosArray) (_err error) {
	C.meos_errno_reset()
	C.meos_array_reset(array._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosArrayResetFree wraps MEOS C function meos_array_reset_free.
func MeosArrayResetFree(array *MeosArray) (_err error) {
	C.meos_errno_reset()
	C.meos_array_reset_free(array._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosArrayDestroy wraps MEOS C function meos_array_destroy.
func MeosArrayDestroy(array *MeosArray) (_err error) {
	C.meos_errno_reset()
	C.meos_array_destroy(array._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosArrayDestroyFree wraps MEOS C function meos_array_destroy_free.
func MeosArrayDestroyFree(array *MeosArray) (_err error) {
	C.meos_errno_reset()
	C.meos_array_destroy_free(array._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// RtreeCreateIntspan wraps MEOS C function rtree_create_intspan.
func RtreeCreateIntspan() (_r0 *RTree, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_create_intspan()
	if _err = meosError(); _err != nil {
		return
	}
	return &RTree{_inner: _cret}, nil
}


// RtreeCreateBigintspan wraps MEOS C function rtree_create_bigintspan.
func RtreeCreateBigintspan() (_r0 *RTree, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_create_bigintspan()
	if _err = meosError(); _err != nil {
		return
	}
	return &RTree{_inner: _cret}, nil
}


// RtreeCreateFloatspan wraps MEOS C function rtree_create_floatspan.
func RtreeCreateFloatspan() (_r0 *RTree, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_create_floatspan()
	if _err = meosError(); _err != nil {
		return
	}
	return &RTree{_inner: _cret}, nil
}


// RtreeCreateDatespan wraps MEOS C function rtree_create_datespan.
func RtreeCreateDatespan() (_r0 *RTree, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_create_datespan()
	if _err = meosError(); _err != nil {
		return
	}
	return &RTree{_inner: _cret}, nil
}


// RtreeCreateTstzspan wraps MEOS C function rtree_create_tstzspan.
func RtreeCreateTstzspan() (_r0 *RTree, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_create_tstzspan()
	if _err = meosError(); _err != nil {
		return
	}
	return &RTree{_inner: _cret}, nil
}


// RtreeCreateTBOX wraps MEOS C function rtree_create_tbox.
func RtreeCreateTBOX() (_r0 *RTree, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_create_tbox()
	if _err = meosError(); _err != nil {
		return
	}
	return &RTree{_inner: _cret}, nil
}


// RtreeCreateSTBOX wraps MEOS C function rtree_create_stbox.
func RtreeCreateSTBOX() (_r0 *RTree, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_create_stbox()
	if _err = meosError(); _err != nil {
		return
	}
	return &RTree{_inner: _cret}, nil
}


// RtreeCreateTpcbox wraps MEOS C function rtree_create_tpcbox.
func RtreeCreateTpcbox() (_r0 *RTree, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_create_tpcbox()
	if _err = meosError(); _err != nil {
		return
	}
	return &RTree{_inner: _cret}, nil
}


// RtreeFree wraps MEOS C function rtree_free.
func RtreeFree(rtree *RTree) (_err error) {
	C.meos_errno_reset()
	C.rtree_free(rtree._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// RtreeNumEntries wraps MEOS C function rtree_num_entries.
func RtreeNumEntries(rtree *RTree) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_num_entries(rtree._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RtreeMemSize wraps MEOS C function rtree_mem_size.
func RtreeMemSize(rtree *RTree) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_mem_size(rtree._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// RtreeHeight wraps MEOS C function rtree_height.
func RtreeHeight(rtree *RTree) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_height(rtree._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RtreeInsert wraps MEOS C function rtree_insert.
func RtreeInsert(rtree *RTree, box unsafe.Pointer, id int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_insert(rtree._inner, unsafe.Pointer(box), C.int64_t(id))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RtreeLoad wraps MEOS C function rtree_load.
func RtreeLoad(rtree *RTree, boxes unsafe.Pointer, ids unsafe.Pointer, count int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_load(rtree._inner, unsafe.Pointer(boxes), (*C.int64_t)(unsafe.Pointer(ids)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RtreeInsertTemporal wraps MEOS C function rtree_insert_temporal.
func RtreeInsertTemporal(rtree *RTree, temp *Temporal, id int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_insert_temporal(rtree._inner, temp._inner, C.int64_t(id))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RtreeInsertTemporalSplit wraps MEOS C function rtree_insert_temporal_split.
func RtreeInsertTemporalSplit(rtree *RTree, temp *Temporal, id int64, maxboxes int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_insert_temporal_split(rtree._inner, temp._inner, C.int64_t(id), C.int(maxboxes))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RtreeSearch wraps MEOS C function rtree_search.
func RtreeSearch(rtree *RTree, op IndexSearchOp, query unsafe.Pointer) (_r0 int, _r1 *MeosArray, _err error) {
	var _out_result C.MeosArray
	C.meos_errno_reset()
	_cret := C.rtree_search(rtree._inner, C.IndexSearchOp(op), unsafe.Pointer(query), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), &MeosArray{_inner: &_out_result}, nil
}


// RtreeJoin wraps MEOS C function rtree_join.
func RtreeJoin(rtree1 *RTree, rtree2 *RTree, op IndexSearchOp) (_r0 int, _r1 *MeosArray, _err error) {
	var _out_result C.MeosArray
	C.meos_errno_reset()
	_cret := C.rtree_join(rtree1._inner, rtree2._inner, C.IndexSearchOp(op), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), &MeosArray{_inner: &_out_result}, nil
}


// RtreeSearchTemporal wraps MEOS C function rtree_search_temporal.
func RtreeSearchTemporal(rtree *RTree, op IndexSearchOp, temp *Temporal) (_r0 int, _r1 *MeosArray, _err error) {
	var _out_result C.MeosArray
	C.meos_errno_reset()
	_cret := C.rtree_search_temporal(rtree._inner, C.IndexSearchOp(op), temp._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), &MeosArray{_inner: &_out_result}, nil
}


// RtreeSearchTemporalDedup wraps MEOS C function rtree_search_temporal_dedup.
func RtreeSearchTemporalDedup(rtree *RTree, op IndexSearchOp, temp *Temporal, maxboxes int) (_r0 int, _r1 *MeosArray, _err error) {
	var _out_result C.MeosArray
	C.meos_errno_reset()
	_cret := C.rtree_search_temporal_dedup(rtree._inner, C.IndexSearchOp(op), temp._inner, C.int(maxboxes), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), &MeosArray{_inner: &_out_result}, nil
}


// RtreeNnCursorOpen wraps MEOS C function rtree_nn_cursor_open.
func RtreeNnCursorOpen(rtree *RTree, query unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_nn_cursor_open(rtree._inner, unsafe.Pointer(query))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RtreeNnCursorNext wraps MEOS C function rtree_nn_cursor_next.
func RtreeNnCursorNext(cursor unsafe.Pointer, id_out unsafe.Pointer, dist_out unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_nn_cursor_next((*C.RTreeNNCursor)(unsafe.Pointer(cursor)), (*C.int64_t)(unsafe.Pointer(id_out)), (*C.double)(unsafe.Pointer(dist_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RtreeNnCursorClose wraps MEOS C function rtree_nn_cursor_close.
func RtreeNnCursorClose(cursor unsafe.Pointer) (_err error) {
	C.meos_errno_reset()
	C.rtree_nn_cursor_close((*C.RTreeNNCursor)(unsafe.Pointer(cursor)))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// SptreeCreateIntspan wraps MEOS C function sptree_create_intspan.
func SptreeCreateIntspan(kind SPTreeKind) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_create_intspan(C.SPTreeKind(kind))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SptreeCreateBigintspan wraps MEOS C function sptree_create_bigintspan.
func SptreeCreateBigintspan(kind SPTreeKind) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_create_bigintspan(C.SPTreeKind(kind))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SptreeCreateFloatspan wraps MEOS C function sptree_create_floatspan.
func SptreeCreateFloatspan(kind SPTreeKind) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_create_floatspan(C.SPTreeKind(kind))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SptreeCreateDatespan wraps MEOS C function sptree_create_datespan.
func SptreeCreateDatespan(kind SPTreeKind) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_create_datespan(C.SPTreeKind(kind))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SptreeCreateTstzspan wraps MEOS C function sptree_create_tstzspan.
func SptreeCreateTstzspan(kind SPTreeKind) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_create_tstzspan(C.SPTreeKind(kind))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SptreeCreateTBOX wraps MEOS C function sptree_create_tbox.
func SptreeCreateTBOX(kind SPTreeKind) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_create_tbox(C.SPTreeKind(kind))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SptreeCreateSTBOX wraps MEOS C function sptree_create_stbox.
func SptreeCreateSTBOX(kind SPTreeKind) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_create_stbox(C.SPTreeKind(kind))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SptreeCreateTpcbox wraps MEOS C function sptree_create_tpcbox.
func SptreeCreateTpcbox(kind SPTreeKind) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_create_tpcbox(C.SPTreeKind(kind))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SptreeFree wraps MEOS C function sptree_free.
func SptreeFree(sptree unsafe.Pointer) (_err error) {
	C.meos_errno_reset()
	C.sptree_free((*C.SPTree)(unsafe.Pointer(sptree)))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// SptreeNumEntries wraps MEOS C function sptree_num_entries.
func SptreeNumEntries(sptree unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_num_entries((*C.SPTree)(unsafe.Pointer(sptree)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// SptreeMemSize wraps MEOS C function sptree_mem_size.
func SptreeMemSize(sptree unsafe.Pointer) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_mem_size((*C.SPTree)(unsafe.Pointer(sptree)))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// SptreeHeight wraps MEOS C function sptree_height.
func SptreeHeight(sptree unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_height((*C.SPTree)(unsafe.Pointer(sptree)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// SptreeInsert wraps MEOS C function sptree_insert.
func SptreeInsert(sptree unsafe.Pointer, box unsafe.Pointer, id int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_insert((*C.SPTree)(unsafe.Pointer(sptree)), unsafe.Pointer(box), C.int64_t(id))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SptreeLoad wraps MEOS C function sptree_load.
func SptreeLoad(sptree unsafe.Pointer, boxes unsafe.Pointer, ids unsafe.Pointer, count int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_load((*C.SPTree)(unsafe.Pointer(sptree)), unsafe.Pointer(boxes), (*C.int64_t)(unsafe.Pointer(ids)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SptreeInsertTemporal wraps MEOS C function sptree_insert_temporal.
func SptreeInsertTemporal(sptree unsafe.Pointer, temp *Temporal, id int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_insert_temporal((*C.SPTree)(unsafe.Pointer(sptree)), temp._inner, C.int64_t(id))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SptreeInsertTemporalSplit wraps MEOS C function sptree_insert_temporal_split.
func SptreeInsertTemporalSplit(sptree unsafe.Pointer, temp *Temporal, id int64, maxboxes int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_insert_temporal_split((*C.SPTree)(unsafe.Pointer(sptree)), temp._inner, C.int64_t(id), C.int(maxboxes))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SptreeSearch wraps MEOS C function sptree_search.
func SptreeSearch(sptree unsafe.Pointer, op IndexSearchOp, query unsafe.Pointer) (_r0 int, _r1 *MeosArray, _err error) {
	var _out_result C.MeosArray
	C.meos_errno_reset()
	_cret := C.sptree_search((*C.SPTree)(unsafe.Pointer(sptree)), C.IndexSearchOp(op), unsafe.Pointer(query), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), &MeosArray{_inner: &_out_result}, nil
}


// SptreeJoin wraps MEOS C function sptree_join.
func SptreeJoin(sptree1 unsafe.Pointer, sptree2 unsafe.Pointer, op IndexSearchOp) (_r0 int, _r1 *MeosArray, _err error) {
	var _out_result C.MeosArray
	C.meos_errno_reset()
	_cret := C.sptree_join((*C.SPTree)(unsafe.Pointer(sptree1)), (*C.SPTree)(unsafe.Pointer(sptree2)), C.IndexSearchOp(op), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), &MeosArray{_inner: &_out_result}, nil
}


// SptreeSearchTemporal wraps MEOS C function sptree_search_temporal.
func SptreeSearchTemporal(sptree unsafe.Pointer, op IndexSearchOp, temp *Temporal) (_r0 int, _r1 *MeosArray, _err error) {
	var _out_result C.MeosArray
	C.meos_errno_reset()
	_cret := C.sptree_search_temporal((*C.SPTree)(unsafe.Pointer(sptree)), C.IndexSearchOp(op), temp._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), &MeosArray{_inner: &_out_result}, nil
}


// SptreeSearchTemporalDedup wraps MEOS C function sptree_search_temporal_dedup.
func SptreeSearchTemporalDedup(sptree unsafe.Pointer, op IndexSearchOp, temp *Temporal, maxboxes int) (_r0 int, _r1 *MeosArray, _err error) {
	var _out_result C.MeosArray
	C.meos_errno_reset()
	_cret := C.sptree_search_temporal_dedup((*C.SPTree)(unsafe.Pointer(sptree)), C.IndexSearchOp(op), temp._inner, C.int(maxboxes), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), &MeosArray{_inner: &_out_result}, nil
}


// SptreeNnCursorOpen wraps MEOS C function sptree_nn_cursor_open.
func SptreeNnCursorOpen(sptree unsafe.Pointer, query unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_nn_cursor_open((*C.SPTree)(unsafe.Pointer(sptree)), unsafe.Pointer(query))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SptreeNnCursorNext wraps MEOS C function sptree_nn_cursor_next.
func SptreeNnCursorNext(cursor unsafe.Pointer, id_out unsafe.Pointer, dist_out unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.sptree_nn_cursor_next((*C.SPNNCursor)(unsafe.Pointer(cursor)), (*C.int64_t)(unsafe.Pointer(id_out)), (*C.double)(unsafe.Pointer(dist_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SptreeNnCursorClose wraps MEOS C function sptree_nn_cursor_close.
func SptreeNnCursorClose(cursor unsafe.Pointer) (_err error) {
	C.meos_errno_reset()
	C.sptree_nn_cursor_close((*C.SPNNCursor)(unsafe.Pointer(cursor)))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TODO meos_initialize_allocator: unsupported param meos_malloc_fn
// func MeosInitializeAllocator(...) { /* not yet handled by codegen */ }


// MeosInitializeNoexitErrorHandler wraps MEOS C function meos_initialize_noexit_error_handler.
func MeosInitializeNoexitErrorHandler() (_err error) {
	C.meos_errno_reset()
	C.meos_initialize_noexit_error_handler()
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosInitializeTimezone wraps MEOS C function meos_initialize_timezone.
func MeosInitializeTimezone(name string) (_err error) {
	_c_name := C.CString(name)
	defer C.free(unsafe.Pointer(_c_name))
	C.meos_errno_reset()
	C.meos_initialize_timezone(_c_name)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosInitializeCollation wraps MEOS C function meos_initialize_collation.
func MeosInitializeCollation() (_err error) {
	C.meos_errno_reset()
	C.meos_initialize_collation()
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosFinalizeTimezone wraps MEOS C function meos_finalize_timezone.
func MeosFinalizeTimezone() (_err error) {
	C.meos_errno_reset()
	C.meos_finalize_timezone()
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosFinalizeCollation wraps MEOS C function meos_finalize_collation.
func MeosFinalizeCollation() (_err error) {
	C.meos_errno_reset()
	C.meos_finalize_collation()
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosFinalizeProjsrs wraps MEOS C function meos_finalize_projsrs.
func MeosFinalizeProjsrs() (_err error) {
	C.meos_errno_reset()
	C.meos_finalize_projsrs()
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosFinalizeWays wraps MEOS C function meos_finalize_ways.
func MeosFinalizeWays() (_err error) {
	C.meos_errno_reset()
	C.meos_finalize_ways()
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosInitializePointcloud wraps MEOS C function meos_initialize_pointcloud.
func MeosInitializePointcloud() (_err error) {
	C.meos_errno_reset()
	C.meos_initialize_pointcloud()
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosSetDatestyle wraps MEOS C function meos_set_datestyle.
func MeosSetDatestyle(newval string, extra unsafe.Pointer) (_r0 bool, _err error) {
	_c_newval := C.CString(newval)
	defer C.free(unsafe.Pointer(_c_newval))
	C.meos_errno_reset()
	_cret := C.meos_set_datestyle(_c_newval, unsafe.Pointer(extra))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// MeosSetIntervalstyle wraps MEOS C function meos_set_intervalstyle.
func MeosSetIntervalstyle(newval string, extra int) (_r0 bool, _err error) {
	_c_newval := C.CString(newval)
	defer C.free(unsafe.Pointer(_c_newval))
	C.meos_errno_reset()
	_cret := C.meos_set_intervalstyle(_c_newval, C.int(extra))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// MeosGetDatestyle wraps MEOS C function meos_get_datestyle.
func MeosGetDatestyle() (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_get_datestyle()
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// MeosGetIntervalstyle wraps MEOS C function meos_get_intervalstyle.
func MeosGetIntervalstyle() (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_get_intervalstyle()
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// MeosSetSpatialRefSysCsv wraps MEOS C function meos_set_spatial_ref_sys_csv.
func MeosSetSpatialRefSysCsv(path string) (_err error) {
	_c_path := C.CString(path)
	defer C.free(unsafe.Pointer(_c_path))
	C.meos_errno_reset()
	C.meos_set_spatial_ref_sys_csv(_c_path)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosSetWaysCsv wraps MEOS C function meos_set_ways_csv.
func MeosSetWaysCsv(path string) (_err error) {
	_c_path := C.CString(path)
	defer C.free(unsafe.Pointer(_c_path))
	C.meos_errno_reset()
	C.meos_set_ways_csv(_c_path)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosInitialize wraps MEOS C function meos_initialize.
func MeosInitialize() (_err error) {
	C.meos_errno_reset()
	C.meos_initialize()
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosFinalize wraps MEOS C function meos_finalize.
func MeosFinalize() (_err error) {
	C.meos_errno_reset()
	C.meos_finalize()
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosVersion wraps MEOS C function meos_version.
func MeosVersion() (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_version()
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// MeosFullVersion wraps MEOS C function meos_full_version.
func MeosFullVersion() (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_full_version()
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// MobilitydbVersion wraps MEOS C function mobilitydb_version.
func MobilitydbVersion() (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.mobilitydb_version()
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// MobilitydbFullVersion wraps MEOS C function mobilitydb_full_version.
func MobilitydbFullVersion() (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.mobilitydb_full_version()
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// BigintsetIn wraps MEOS C function bigintset_in.
func BigintsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.bigintset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// BigintsetOut wraps MEOS C function bigintset_out.
func BigintsetOut(set *Set) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintset_out(set._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// BigintspanExpand wraps MEOS C function bigintspan_expand.
func BigintspanExpand(s *Span, value int64) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspan_expand(s._inner, C.int64_t(value))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// BigintspanIn wraps MEOS C function bigintspan_in.
func BigintspanIn(str string) (_r0 *Span, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.bigintspan_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// BigintspanOut wraps MEOS C function bigintspan_out.
func BigintspanOut(s *Span) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspan_out(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// BigintspansetIn wraps MEOS C function bigintspanset_in.
func BigintspansetIn(str string) (_r0 *SpanSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.bigintspanset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// BigintspansetOut wraps MEOS C function bigintspanset_out.
func BigintspansetOut(ss *SpanSet) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspanset_out(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// DatesetIn wraps MEOS C function dateset_in.
func DatesetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.dateset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// DatesetOut wraps MEOS C function dateset_out.
func DatesetOut(s *Set) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.dateset_out(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// DatespanIn wraps MEOS C function datespan_in.
func DatespanIn(str string) (_r0 *Span, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.datespan_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// DatespanOut wraps MEOS C function datespan_out.
func DatespanOut(s *Span) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.datespan_out(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// DatespansetIn wraps MEOS C function datespanset_in.
func DatespansetIn(str string) (_r0 *SpanSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.datespanset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// DatespansetOut wraps MEOS C function datespanset_out.
func DatespansetOut(ss *SpanSet) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.datespanset_out(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// FloatsetIn wraps MEOS C function floatset_in.
func FloatsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.floatset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// FloatsetOut wraps MEOS C function floatset_out.
func FloatsetOut(set *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.floatset_out(set._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// FloatspanExpand wraps MEOS C function floatspan_expand.
func FloatspanExpand(s *Span, value float64) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_expand(s._inner, C.double(value))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatspanIn wraps MEOS C function floatspan_in.
func FloatspanIn(str string) (_r0 *Span, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.floatspan_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatspanOut wraps MEOS C function floatspan_out.
func FloatspanOut(s *Span, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_out(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// FloatspansetIn wraps MEOS C function floatspanset_in.
func FloatspansetIn(str string) (_r0 *SpanSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.floatspanset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// FloatspansetOut wraps MEOS C function floatspanset_out.
func FloatspansetOut(ss *SpanSet, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_out(ss._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// IntsetIn wraps MEOS C function intset_in.
func IntsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.intset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntsetOut wraps MEOS C function intset_out.
func IntsetOut(set *Set) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.intset_out(set._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// IntspanExpand wraps MEOS C function intspan_expand.
func IntspanExpand(s *Span, value int32) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intspan_expand(s._inner, C.int32(value))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntspanIn wraps MEOS C function intspan_in.
func IntspanIn(str string) (_r0 *Span, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.intspan_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntspanOut wraps MEOS C function intspan_out.
func IntspanOut(s *Span) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.intspan_out(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// IntspansetIn wraps MEOS C function intspanset_in.
func IntspansetIn(str string) (_r0 *SpanSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.intspanset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntspansetOut wraps MEOS C function intspanset_out.
func IntspansetOut(ss *SpanSet) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.intspanset_out(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// SetAsHexwkb wraps MEOS C function set_as_hexwkb.
func SetAsHexwkb(s *Set, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.set_as_hexwkb(s._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// SetAsWKB wraps MEOS C function set_as_wkb.
func SetAsWKB(s *Set, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.set_as_wkb(s._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SetFromHexwkb wraps MEOS C function set_from_hexwkb.
func SetFromHexwkb(hexwkb string) (_r0 *Set, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.set_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// SetFromWKB wraps MEOS C function set_from_wkb.
func SetFromWKB(wkb unsafe.Pointer, size uint) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.set_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// SpanAsHexwkb wraps MEOS C function span_as_hexwkb.
func SpanAsHexwkb(s *Span, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.span_as_hexwkb(s._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// SpanAsWKB wraps MEOS C function span_as_wkb.
func SpanAsWKB(s *Span, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.span_as_wkb(s._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SpanFromHexwkb wraps MEOS C function span_from_hexwkb.
func SpanFromHexwkb(hexwkb string) (_r0 *Span, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.span_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpanFromWKB wraps MEOS C function span_from_wkb.
func SpanFromWKB(wkb unsafe.Pointer, size uint) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.span_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpansetAsHexwkb wraps MEOS C function spanset_as_hexwkb.
func SpansetAsHexwkb(ss *SpanSet, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_as_hexwkb(ss._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// SpansetAsWKB wraps MEOS C function spanset_as_wkb.
func SpansetAsWKB(ss *SpanSet, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_as_wkb(ss._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SpansetFromHexwkb wraps MEOS C function spanset_from_hexwkb.
func SpansetFromHexwkb(hexwkb string) (_r0 *SpanSet, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.spanset_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// SpansetFromWKB wraps MEOS C function spanset_from_wkb.
func SpansetFromWKB(wkb unsafe.Pointer, size uint) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TextsetIn wraps MEOS C function textset_in.
func TextsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.textset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TextsetOut wraps MEOS C function textset_out.
func TextsetOut(set *Set) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.textset_out(set._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TstzsetIn wraps MEOS C function tstzset_in.
func TstzsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tstzset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TstzsetOut wraps MEOS C function tstzset_out.
func TstzsetOut(set *Set) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzset_out(set._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TstzspanIn wraps MEOS C function tstzspan_in.
func TstzspanIn(str string) (_r0 *Span, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tstzspan_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TstzspanOut wraps MEOS C function tstzspan_out.
func TstzspanOut(s *Span) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_out(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TstzspansetIn wraps MEOS C function tstzspanset_in.
func TstzspansetIn(str string) (_r0 *SpanSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tstzspanset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TstzspansetOut wraps MEOS C function tstzspanset_out.
func TstzspansetOut(ss *SpanSet) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_out(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// BigintsetMake wraps MEOS C function bigintset_make.
func BigintsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintset_make((*C.int64_t)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// BigintspanMake wraps MEOS C function bigintspan_make.
func BigintspanMake(lower int64, upper int64, lower_inc bool, upper_inc bool) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspan_make(C.int64_t(lower), C.int64_t(upper), C.bool(lower_inc), C.bool(upper_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// DatesetMake wraps MEOS C function dateset_make.
func DatesetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.dateset_make((*C.DateADT)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// DatespanMake wraps MEOS C function datespan_make.
func DatespanMake(lower int32, upper int32, lower_inc bool, upper_inc bool) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.datespan_make(C.DateADT(lower), C.DateADT(upper), C.bool(lower_inc), C.bool(upper_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatsetMake wraps MEOS C function floatset_make.
func FloatsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.floatset_make((*C.double)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// FloatspanMake wraps MEOS C function floatspan_make.
func FloatspanMake(lower float64, upper float64, lower_inc bool, upper_inc bool) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_make(C.double(lower), C.double(upper), C.bool(lower_inc), C.bool(upper_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntsetMake wraps MEOS C function intset_make.
func IntsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intset_make((*C.int)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntspanMake wraps MEOS C function intspan_make.
func IntspanMake(lower int, upper int, lower_inc bool, upper_inc bool) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intspan_make(C.int(lower), C.int(upper), C.bool(lower_inc), C.bool(upper_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SetCopy wraps MEOS C function set_copy.
func SetCopy(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.set_copy(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// SpanCopy wraps MEOS C function span_copy.
func SpanCopy(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.span_copy(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpansetCopy wraps MEOS C function spanset_copy.
func SpansetCopy(ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_copy(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// SpansetMake wraps MEOS C function spanset_make.
func SpansetMake(spans *Span, count int) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_make(spans._inner, C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TextsetMake wraps MEOS C function textset_make.
func TextsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.textset_make((**C.text)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TstzsetMake wraps MEOS C function tstzset_make.
func TstzsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzset_make((*C.TimestampTz)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TstzspanMake wraps MEOS C function tstzspan_make.
func TstzspanMake(lower int64, upper int64, lower_inc bool, upper_inc bool) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_make(C.TimestampTz(lower), C.TimestampTz(upper), C.bool(lower_inc), C.bool(upper_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// BigintToSet wraps MEOS C function bigint_to_set.
func BigintToSet(i int64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_to_set(C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// BigintToSpan wraps MEOS C function bigint_to_span.
func BigintToSpan(i int) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_to_span(C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// BigintToSpanset wraps MEOS C function bigint_to_spanset.
func BigintToSpanset(i int) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_to_spanset(C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// DateToSet wraps MEOS C function date_to_set.
func DateToSet(d int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.date_to_set(C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// DateToSpan wraps MEOS C function date_to_span.
func DateToSpan(d int32) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.date_to_span(C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// DateToSpanset wraps MEOS C function date_to_spanset.
func DateToSpanset(d int32) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.date_to_spanset(C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// DatesetToTstzset wraps MEOS C function dateset_to_tstzset.
func DatesetToTstzset(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.dateset_to_tstzset(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// DatespanToTstzspan wraps MEOS C function datespan_to_tstzspan.
func DatespanToTstzspan(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.datespan_to_tstzspan(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// DatespansetToTstzspanset wraps MEOS C function datespanset_to_tstzspanset.
func DatespansetToTstzspanset(ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.datespanset_to_tstzspanset(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// FloatToSet wraps MEOS C function float_to_set.
func FloatToSet(d float64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.float_to_set(C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// FloatToSpan wraps MEOS C function float_to_span.
func FloatToSpan(d float64) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.float_to_span(C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatToSpanset wraps MEOS C function float_to_spanset.
func FloatToSpanset(d float64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.float_to_spanset(C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// FloatsetToIntset wraps MEOS C function floatset_to_intset.
func FloatsetToIntset(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.floatset_to_intset(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// FloatspanToIntspan wraps MEOS C function floatspan_to_intspan.
func FloatspanToIntspan(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_to_intspan(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatspanToBigintspan wraps MEOS C function floatspan_to_bigintspan.
func FloatspanToBigintspan(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_to_bigintspan(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatspansetToIntspanset wraps MEOS C function floatspanset_to_intspanset.
func FloatspansetToIntspanset(ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_to_intspanset(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntToSet wraps MEOS C function int_to_set.
func IntToSet(i int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.int_to_set(C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntToSpan wraps MEOS C function int_to_span.
func IntToSpan(i int) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.int_to_span(C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntToSpanset wraps MEOS C function int_to_spanset.
func IntToSpanset(i int) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.int_to_spanset(C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntsetToFloatset wraps MEOS C function intset_to_floatset.
func IntsetToFloatset(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intset_to_floatset(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntspanToFloatspan wraps MEOS C function intspan_to_floatspan.
func IntspanToFloatspan(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intspan_to_floatspan(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntspanToBigintspan wraps MEOS C function intspan_to_bigintspan.
func IntspanToBigintspan(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intspan_to_bigintspan(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// BigintspanToIntspan wraps MEOS C function bigintspan_to_intspan.
func BigintspanToIntspan(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspan_to_intspan(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// BigintspanToFloatspan wraps MEOS C function bigintspan_to_floatspan.
func BigintspanToFloatspan(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspan_to_floatspan(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntspansetToFloatspanset wraps MEOS C function intspanset_to_floatspanset.
func IntspansetToFloatspanset(ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.intspanset_to_floatspanset(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// SetToSpan wraps MEOS C function set_to_span.
func SetToSpan(s *Set) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.set_to_span(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SetToSpanset wraps MEOS C function set_to_spanset.
func SetToSpanset(s *Set) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.set_to_spanset(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// SpanToSpanset wraps MEOS C function span_to_spanset.
func SpanToSpanset(s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.span_to_spanset(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TextToSet wraps MEOS C function text_to_set.
func TextToSet(txt string) (_r0 *Set, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_to_set(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TimestamptzToSet wraps MEOS C function timestamptz_to_set.
func TimestamptzToSet(t int64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_to_set(C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TimestamptzToSpan wraps MEOS C function timestamptz_to_span.
func TimestamptzToSpan(t int64) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_to_span(C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TimestamptzToSpanset wraps MEOS C function timestamptz_to_spanset.
func TimestamptzToSpanset(t int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_to_spanset(C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TstzsetToDateset wraps MEOS C function tstzset_to_dateset.
func TstzsetToDateset(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzset_to_dateset(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TstzspanToDatespan wraps MEOS C function tstzspan_to_datespan.
func TstzspanToDatespan(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_to_datespan(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TstzspansetToDatespanset wraps MEOS C function tstzspanset_to_datespanset.
func TstzspansetToDatespanset(ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_to_datespanset(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// BigintsetEndValue wraps MEOS C function bigintset_end_value.
func BigintsetEndValue(s *Set) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// BigintsetStartValue wraps MEOS C function bigintset_start_value.
func BigintsetStartValue(s *Set) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// BigintsetValueN wraps MEOS C function bigintset_value_n.
func BigintsetValueN(s *Set, n int) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.int64_t
	C.meos_errno_reset()
	_cret := C.bigintset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// BigintsetValues wraps MEOS C function bigintset_values.
func BigintsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// BigintspanLower wraps MEOS C function bigintspan_lower.
func BigintspanLower(s *Span) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspan_lower(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// BigintspanUpper wraps MEOS C function bigintspan_upper.
func BigintspanUpper(s *Span) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspan_upper(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// BigintspanWidth wraps MEOS C function bigintspan_width.
func BigintspanWidth(s *Span) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspan_width(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// BigintspansetLower wraps MEOS C function bigintspanset_lower.
func BigintspansetLower(ss *SpanSet) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspanset_lower(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// BigintspansetUpper wraps MEOS C function bigintspanset_upper.
func BigintspansetUpper(ss *SpanSet) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspanset_upper(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// BigintspansetWidth wraps MEOS C function bigintspanset_width.
func BigintspansetWidth(ss *SpanSet, boundspan bool) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspanset_width(ss._inner, C.bool(boundspan))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DatesetEndValue wraps MEOS C function dateset_end_value.
func DatesetEndValue(s *Set) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.dateset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DatesetStartValue wraps MEOS C function dateset_start_value.
func DatesetStartValue(s *Set) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.dateset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DatesetValueN wraps MEOS C function dateset_value_n.
func DatesetValueN(s *Set, n int) (_r0 bool, _r1 int32, _err error) {
	var _out_result C.DateADT
	C.meos_errno_reset()
	_cret := C.dateset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int32(_out_result), nil
}


// DatesetValues wraps MEOS C function dateset_values.
func DatesetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.dateset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// DatespanDuration wraps MEOS C function datespan_duration.
func DatespanDuration(s *Span) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.datespan_duration(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// DatespanLower wraps MEOS C function datespan_lower.
func DatespanLower(s *Span) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.datespan_lower(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DatespanUpper wraps MEOS C function datespan_upper.
func DatespanUpper(s *Span) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.datespan_upper(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DatespansetDateN wraps MEOS C function datespanset_date_n.
func DatespansetDateN(ss *SpanSet, n int) (_r0 bool, _r1 int32, _err error) {
	var _out_result C.DateADT
	C.meos_errno_reset()
	_cret := C.datespanset_date_n(ss._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int32(_out_result), nil
}


// DatespansetDates wraps MEOS C function datespanset_dates.
func DatespansetDates(ss *SpanSet) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.datespanset_dates(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// DatespansetDuration wraps MEOS C function datespanset_duration.
func DatespansetDuration(ss *SpanSet, boundspan bool) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.datespanset_duration(ss._inner, C.bool(boundspan))
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// DatespansetEndDate wraps MEOS C function datespanset_end_date.
func DatespansetEndDate(ss *SpanSet) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.datespanset_end_date(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DatespansetLower wraps MEOS C function datespanset_lower.
func DatespansetLower(ss *SpanSet) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.datespanset_lower(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DatespansetNumDates wraps MEOS C function datespanset_num_dates.
func DatespansetNumDates(ss *SpanSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.datespanset_num_dates(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DatespansetStartDate wraps MEOS C function datespanset_start_date.
func DatespansetStartDate(ss *SpanSet) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.datespanset_start_date(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DatespansetUpper wraps MEOS C function datespanset_upper.
func DatespansetUpper(ss *SpanSet) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.datespanset_upper(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// FloatsetEndValue wraps MEOS C function floatset_end_value.
func FloatsetEndValue(s *Set) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.floatset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// FloatsetStartValue wraps MEOS C function floatset_start_value.
func FloatsetStartValue(s *Set) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.floatset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// FloatsetValueN wraps MEOS C function floatset_value_n.
func FloatsetValueN(s *Set, n int) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.floatset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// FloatsetValues wraps MEOS C function floatset_values.
func FloatsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.floatset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// FloatspanLower wraps MEOS C function floatspan_lower.
func FloatspanLower(s *Span) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_lower(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// FloatspanUpper wraps MEOS C function floatspan_upper.
func FloatspanUpper(s *Span) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_upper(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// FloatspanWidth wraps MEOS C function floatspan_width.
func FloatspanWidth(s *Span) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_width(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// FloatspansetLower wraps MEOS C function floatspanset_lower.
func FloatspansetLower(ss *SpanSet) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_lower(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// FloatspansetUpper wraps MEOS C function floatspanset_upper.
func FloatspansetUpper(ss *SpanSet) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_upper(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// FloatspansetWidth wraps MEOS C function floatspanset_width.
func FloatspansetWidth(ss *SpanSet, boundspan bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_width(ss._inner, C.bool(boundspan))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// IntsetEndValue wraps MEOS C function intset_end_value.
func IntsetEndValue(s *Set) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.intset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// IntsetStartValue wraps MEOS C function intset_start_value.
func IntsetStartValue(s *Set) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.intset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// IntsetValueN wraps MEOS C function intset_value_n.
func IntsetValueN(s *Set, n int) (_r0 bool, _r1 int, _err error) {
	var _out_result C.int
	C.meos_errno_reset()
	_cret := C.intset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int(_out_result), nil
}


// IntsetValues wraps MEOS C function intset_values.
func IntsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.intset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// IntspanLower wraps MEOS C function intspan_lower.
func IntspanLower(s *Span) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.intspan_lower(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// IntspanUpper wraps MEOS C function intspan_upper.
func IntspanUpper(s *Span) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.intspan_upper(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// IntspanWidth wraps MEOS C function intspan_width.
func IntspanWidth(s *Span) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.intspan_width(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// IntspansetLower wraps MEOS C function intspanset_lower.
func IntspansetLower(ss *SpanSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.intspanset_lower(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// IntspansetUpper wraps MEOS C function intspanset_upper.
func IntspansetUpper(ss *SpanSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.intspanset_upper(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// IntspansetWidth wraps MEOS C function intspanset_width.
func IntspansetWidth(ss *SpanSet, boundspan bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.intspanset_width(ss._inner, C.bool(boundspan))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// SetHash wraps MEOS C function set_hash.
func SetHash(s *Set) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.set_hash(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// SetHashExtended wraps MEOS C function set_hash_extended.
func SetHashExtended(s *Set, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.set_hash_extended(s._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// SetNumValues wraps MEOS C function set_num_values.
func SetNumValues(s *Set) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.set_num_values(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// SpanHash wraps MEOS C function span_hash.
func SpanHash(s *Span) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.span_hash(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// SpanHashExtended wraps MEOS C function span_hash_extended.
func SpanHashExtended(s *Span, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.span_hash_extended(s._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// SpanLowerInc wraps MEOS C function span_lower_inc.
func SpanLowerInc(s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_lower_inc(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanUpperInc wraps MEOS C function span_upper_inc.
func SpanUpperInc(s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_upper_inc(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpansetEndSpan wraps MEOS C function spanset_end_span.
func SpansetEndSpan(ss *SpanSet) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_end_span(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpansetHash wraps MEOS C function spanset_hash.
func SpansetHash(ss *SpanSet) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_hash(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// SpansetHashExtended wraps MEOS C function spanset_hash_extended.
func SpansetHashExtended(ss *SpanSet, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_hash_extended(ss._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// SpansetLowerInc wraps MEOS C function spanset_lower_inc.
func SpansetLowerInc(ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_lower_inc(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpansetNumSpans wraps MEOS C function spanset_num_spans.
func SpansetNumSpans(ss *SpanSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_num_spans(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// SpansetSpan wraps MEOS C function spanset_span.
func SpansetSpan(ss *SpanSet) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_span(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpansetSpanN wraps MEOS C function spanset_span_n.
func SpansetSpanN(ss *SpanSet, i int) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_span_n(ss._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpansetSpanarr wraps MEOS C function spanset_spanarr.
func SpansetSpanarr(ss *SpanSet, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_spanarr(ss._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SpansetStartSpan wraps MEOS C function spanset_start_span.
func SpansetStartSpan(ss *SpanSet) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_start_span(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpansetUpperInc wraps MEOS C function spanset_upper_inc.
func SpansetUpperInc(ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_upper_inc(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TextsetEndValue wraps MEOS C function textset_end_value.
func TextsetEndValue(s *Set) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.textset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextsetStartValue wraps MEOS C function textset_start_value.
func TextsetStartValue(s *Set) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.textset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextsetValueN wraps MEOS C function textset_value_n.
func TextsetValueN(s *Set, n int) (_r0 bool, _r1 string, _err error) {
	var _out_result *C.text
	C.meos_errno_reset()
	_cret := C.textset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), C.GoString(C.text_to_cstring(_out_result)), nil
}


// TextsetValues wraps MEOS C function textset_values.
func TextsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.textset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TstzsetEndValue wraps MEOS C function tstzset_end_value.
func TstzsetEndValue(s *Set) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TstzsetStartValue wraps MEOS C function tstzset_start_value.
func TstzsetStartValue(s *Set) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TstzsetValueN wraps MEOS C function tstzset_value_n.
func TstzsetValueN(s *Set, n int) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.TimestampTz
	C.meos_errno_reset()
	_cret := C.tstzset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// TstzsetValues wraps MEOS C function tstzset_values.
func TstzsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TstzspanDuration wraps MEOS C function tstzspan_duration.
func TstzspanDuration(s *Span) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_duration(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// TstzspanLower wraps MEOS C function tstzspan_lower.
func TstzspanLower(s *Span) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_lower(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TstzspanUpper wraps MEOS C function tstzspan_upper.
func TstzspanUpper(s *Span) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_upper(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TstzspansetDuration wraps MEOS C function tstzspanset_duration.
func TstzspansetDuration(ss *SpanSet, boundspan bool) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_duration(ss._inner, C.bool(boundspan))
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// TstzspansetEndTimestamptz wraps MEOS C function tstzspanset_end_timestamptz.
func TstzspansetEndTimestamptz(ss *SpanSet) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_end_timestamptz(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TstzspansetLower wraps MEOS C function tstzspanset_lower.
func TstzspansetLower(ss *SpanSet) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_lower(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TstzspansetNumTimestamps wraps MEOS C function tstzspanset_num_timestamps.
func TstzspansetNumTimestamps(ss *SpanSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_num_timestamps(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TstzspansetStartTimestamptz wraps MEOS C function tstzspanset_start_timestamptz.
func TstzspansetStartTimestamptz(ss *SpanSet) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_start_timestamptz(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TstzspansetTimestamps wraps MEOS C function tstzspanset_timestamps.
func TstzspansetTimestamps(ss *SpanSet) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_timestamps(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TstzspansetTimestamptzN wraps MEOS C function tstzspanset_timestamptz_n.
func TstzspansetTimestamptzN(ss *SpanSet, n int) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.TimestampTz
	C.meos_errno_reset()
	_cret := C.tstzspanset_timestamptz_n(ss._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// TstzspansetUpper wraps MEOS C function tstzspanset_upper.
func TstzspansetUpper(ss *SpanSet) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_upper(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// BigintsetShiftScale wraps MEOS C function bigintset_shift_scale.
func BigintsetShiftScale(s *Set, shift int64, width int64, hasshift bool, haswidth bool) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintset_shift_scale(s._inner, C.int64_t(shift), C.int64_t(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// BigintspanShiftScale wraps MEOS C function bigintspan_shift_scale.
func BigintspanShiftScale(s *Span, shift int64, width int64, hasshift bool, haswidth bool) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspan_shift_scale(s._inner, C.int64_t(shift), C.int64_t(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// BigintspansetShiftScale wraps MEOS C function bigintspanset_shift_scale.
func BigintspansetShiftScale(ss *SpanSet, shift int64, width int64, hasshift bool, haswidth bool) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspanset_shift_scale(ss._inner, C.int64_t(shift), C.int64_t(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// DatesetShiftScale wraps MEOS C function dateset_shift_scale.
func DatesetShiftScale(s *Set, shift int, width int, hasshift bool, haswidth bool) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.dateset_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// DatespanShiftScale wraps MEOS C function datespan_shift_scale.
func DatespanShiftScale(s *Span, shift int, width int, hasshift bool, haswidth bool) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.datespan_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// DatespansetShiftScale wraps MEOS C function datespanset_shift_scale.
func DatespansetShiftScale(ss *SpanSet, shift int, width int, hasshift bool, haswidth bool) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.datespanset_shift_scale(ss._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// FloatsetCeil wraps MEOS C function floatset_ceil.
func FloatsetCeil(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.floatset_ceil(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// FloatsetDegrees wraps MEOS C function floatset_degrees.
func FloatsetDegrees(s *Set, normalize bool) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.floatset_degrees(s._inner, C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// FloatsetFloor wraps MEOS C function floatset_floor.
func FloatsetFloor(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.floatset_floor(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// FloatsetRadians wraps MEOS C function floatset_radians.
func FloatsetRadians(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.floatset_radians(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// FloatsetShiftScale wraps MEOS C function floatset_shift_scale.
func FloatsetShiftScale(s *Set, shift float64, width float64, hasshift bool, haswidth bool) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.floatset_shift_scale(s._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// FloatspanCeil wraps MEOS C function floatspan_ceil.
func FloatspanCeil(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_ceil(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatspanDegrees wraps MEOS C function floatspan_degrees.
func FloatspanDegrees(s *Span, normalize bool) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_degrees(s._inner, C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatspanFloor wraps MEOS C function floatspan_floor.
func FloatspanFloor(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_floor(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatspanRadians wraps MEOS C function floatspan_radians.
func FloatspanRadians(s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_radians(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatspanRound wraps MEOS C function floatspan_round.
func FloatspanRound(s *Span, maxdd int) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_round(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatspanShiftScale wraps MEOS C function floatspan_shift_scale.
func FloatspanShiftScale(s *Span, shift float64, width float64, hasshift bool, haswidth bool) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_shift_scale(s._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatspansetCeil wraps MEOS C function floatspanset_ceil.
func FloatspansetCeil(ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_ceil(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// FloatspansetFloor wraps MEOS C function floatspanset_floor.
func FloatspansetFloor(ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_floor(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// FloatspansetDegrees wraps MEOS C function floatspanset_degrees.
func FloatspansetDegrees(ss *SpanSet, normalize bool) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_degrees(ss._inner, C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// FloatspansetRadians wraps MEOS C function floatspanset_radians.
func FloatspansetRadians(ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_radians(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// FloatspansetRound wraps MEOS C function floatspanset_round.
func FloatspansetRound(ss *SpanSet, maxdd int) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_round(ss._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// FloatspansetShiftScale wraps MEOS C function floatspanset_shift_scale.
func FloatspansetShiftScale(ss *SpanSet, shift float64, width float64, hasshift bool, haswidth bool) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_shift_scale(ss._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntsetShiftScale wraps MEOS C function intset_shift_scale.
func IntsetShiftScale(s *Set, shift int, width int, hasshift bool, haswidth bool) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intset_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntspanShiftScale wraps MEOS C function intspan_shift_scale.
func IntspanShiftScale(s *Span, shift int, width int, hasshift bool, haswidth bool) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intspan_shift_scale(s._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntspansetShiftScale wraps MEOS C function intspanset_shift_scale.
func IntspansetShiftScale(ss *SpanSet, shift int, width int, hasshift bool, haswidth bool) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.intspanset_shift_scale(ss._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TstzspanExpand wraps MEOS C function tstzspan_expand.
func TstzspanExpand(s *Span, interv *Interval) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_expand(s._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SetRound wraps MEOS C function set_round.
func SetRound(s *Set, maxdd int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.set_round(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TextcatTextTextset wraps MEOS C function textcat_text_textset.
func TextcatTextTextset(txt string, s *Set) (_r0 *Set, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.textcat_text_textset(_c_txt, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TextcatTextsetText wraps MEOS C function textcat_textset_text.
func TextcatTextsetText(s *Set, txt string) (_r0 *Set, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.textcat_textset_text(s._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TextsetInitcap wraps MEOS C function textset_initcap.
func TextsetInitcap(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.textset_initcap(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TextsetLower wraps MEOS C function textset_lower.
func TextsetLower(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.textset_lower(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TextsetUpper wraps MEOS C function textset_upper.
func TextsetUpper(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.textset_upper(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TimestamptzTprecision wraps MEOS C function timestamptz_tprecision.
func TimestamptzTprecision(t int64, duration *Interval, torigin int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_tprecision(C.TimestampTz(t), duration._inner, C.TimestampTz(torigin))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TstzsetShiftScale wraps MEOS C function tstzset_shift_scale.
func TstzsetShiftScale(s *Set, shift *Interval, duration *Interval) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzset_shift_scale(s._inner, shift._inner, duration._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TstzsetTprecision wraps MEOS C function tstzset_tprecision.
func TstzsetTprecision(s *Set, duration *Interval, torigin int64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzset_tprecision(s._inner, duration._inner, C.TimestampTz(torigin))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TstzspanShiftScale wraps MEOS C function tstzspan_shift_scale.
func TstzspanShiftScale(s *Span, shift *Interval, duration *Interval) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_shift_scale(s._inner, shift._inner, duration._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TstzspanTprecision wraps MEOS C function tstzspan_tprecision.
func TstzspanTprecision(s *Span, duration *Interval, torigin int64) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_tprecision(s._inner, duration._inner, C.TimestampTz(torigin))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TstzspansetShiftScale wraps MEOS C function tstzspanset_shift_scale.
func TstzspansetShiftScale(ss *SpanSet, shift *Interval, duration *Interval) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_shift_scale(ss._inner, shift._inner, duration._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TstzspansetTprecision wraps MEOS C function tstzspanset_tprecision.
func TstzspansetTprecision(ss *SpanSet, duration *Interval, torigin int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_tprecision(ss._inner, duration._inner, C.TimestampTz(torigin))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// SetCmp wraps MEOS C function set_cmp.
func SetCmp(s1 *Set, s2 *Set) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.set_cmp(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// SetEq wraps MEOS C function set_eq.
func SetEq(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.set_eq(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SetGe wraps MEOS C function set_ge.
func SetGe(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.set_ge(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SetGt wraps MEOS C function set_gt.
func SetGt(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.set_gt(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SetLe wraps MEOS C function set_le.
func SetLe(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.set_le(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SetLt wraps MEOS C function set_lt.
func SetLt(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.set_lt(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SetNe wraps MEOS C function set_ne.
func SetNe(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.set_ne(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanCmp wraps MEOS C function span_cmp.
func SpanCmp(s1 *Span, s2 *Span) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.span_cmp(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// SpanEq wraps MEOS C function span_eq.
func SpanEq(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_eq(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanGe wraps MEOS C function span_ge.
func SpanGe(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_ge(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanGt wraps MEOS C function span_gt.
func SpanGt(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_gt(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanLe wraps MEOS C function span_le.
func SpanLe(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_le(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanLt wraps MEOS C function span_lt.
func SpanLt(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_lt(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanNe wraps MEOS C function span_ne.
func SpanNe(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_ne(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpansetCmp wraps MEOS C function spanset_cmp.
func SpansetCmp(ss1 *SpanSet, ss2 *SpanSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_cmp(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// SpansetEq wraps MEOS C function spanset_eq.
func SpansetEq(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_eq(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpansetGe wraps MEOS C function spanset_ge.
func SpansetGe(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_ge(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpansetGt wraps MEOS C function spanset_gt.
func SpansetGt(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_gt(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpansetLe wraps MEOS C function spanset_le.
func SpansetLe(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_le(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpansetLt wraps MEOS C function spanset_lt.
func SpansetLt(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_lt(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpansetNe wraps MEOS C function spanset_ne.
func SpansetNe(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_ne(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SetSpans wraps MEOS C function set_spans.
func SetSpans(s *Set, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.set_spans(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SetSplitEachNSpans wraps MEOS C function set_split_each_n_spans.
func SetSplitEachNSpans(s *Set, elems_per_span int, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.set_split_each_n_spans(s._inner, C.int(elems_per_span), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SetSplitNSpans wraps MEOS C function set_split_n_spans.
func SetSplitNSpans(s *Set, span_count int, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.set_split_n_spans(s._inner, C.int(span_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpansetSpans wraps MEOS C function spanset_spans.
func SpansetSpans(ss *SpanSet, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_spans(ss._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpansetSplitEachNSpans wraps MEOS C function spanset_split_each_n_spans.
func SpansetSplitEachNSpans(ss *SpanSet, elems_per_span int, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_split_each_n_spans(ss._inner, C.int(elems_per_span), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpansetSplitNSpans wraps MEOS C function spanset_split_n_spans.
func SpansetSplitNSpans(ss *SpanSet, span_count int, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_split_n_spans(ss._inner, C.int(span_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// AdjacentSpanBigint wraps MEOS C function adjacent_span_bigint.
func AdjacentSpanBigint(s *Span, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_span_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpanDate wraps MEOS C function adjacent_span_date.
func AdjacentSpanDate(s *Span, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_span_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpanFloat wraps MEOS C function adjacent_span_float.
func AdjacentSpanFloat(s *Span, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_span_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpanInt wraps MEOS C function adjacent_span_int.
func AdjacentSpanInt(s *Span, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_span_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpanSpan wraps MEOS C function adjacent_span_span.
func AdjacentSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpanSpanset wraps MEOS C function adjacent_span_spanset.
func AdjacentSpanSpanset(s *Span, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_span_spanset(s._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpanTimestamptz wraps MEOS C function adjacent_span_timestamptz.
func AdjacentSpanTimestamptz(s *Span, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_span_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentBigintSpan wraps MEOS C function adjacent_bigint_span.
func AdjacentBigintSpan(i int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_bigint_span(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentDateSpan wraps MEOS C function adjacent_date_span.
func AdjacentDateSpan(d int32, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_date_span(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentFloatSpan wraps MEOS C function adjacent_float_span.
func AdjacentFloatSpan(d float64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_float_span(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentIntSpan wraps MEOS C function adjacent_int_span.
func AdjacentIntSpan(i int, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_int_span(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTimestamptzSpan wraps MEOS C function adjacent_timestamptz_span.
func AdjacentTimestamptzSpan(t int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_timestamptz_span(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpansetBigint wraps MEOS C function adjacent_spanset_bigint.
func AdjacentSpansetBigint(ss *SpanSet, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_spanset_bigint(ss._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpansetDate wraps MEOS C function adjacent_spanset_date.
func AdjacentSpansetDate(ss *SpanSet, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_spanset_date(ss._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpansetFloat wraps MEOS C function adjacent_spanset_float.
func AdjacentSpansetFloat(ss *SpanSet, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_spanset_float(ss._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpansetInt wraps MEOS C function adjacent_spanset_int.
func AdjacentSpansetInt(ss *SpanSet, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_spanset_int(ss._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpansetTimestamptz wraps MEOS C function adjacent_spanset_timestamptz.
func AdjacentSpansetTimestamptz(ss *SpanSet, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpansetSpan wraps MEOS C function adjacent_spanset_span.
func AdjacentSpansetSpan(ss *SpanSet, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_spanset_span(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentSpansetSpanset wraps MEOS C function adjacent_spanset_spanset.
func AdjacentSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_spanset_spanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedBigintSet wraps MEOS C function contained_bigint_set.
func ContainedBigintSet(i int64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_bigint_set(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedBigintSpan wraps MEOS C function contained_bigint_span.
func ContainedBigintSpan(i int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_bigint_span(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedBigintSpanset wraps MEOS C function contained_bigint_spanset.
func ContainedBigintSpanset(i int64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_bigint_spanset(C.int64_t(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedDateSet wraps MEOS C function contained_date_set.
func ContainedDateSet(d int32, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_date_set(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedDateSpan wraps MEOS C function contained_date_span.
func ContainedDateSpan(d int32, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_date_span(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedDateSpanset wraps MEOS C function contained_date_spanset.
func ContainedDateSpanset(d int32, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_date_spanset(C.DateADT(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedFloatSet wraps MEOS C function contained_float_set.
func ContainedFloatSet(d float64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_float_set(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedFloatSpan wraps MEOS C function contained_float_span.
func ContainedFloatSpan(d float64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_float_span(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedFloatSpanset wraps MEOS C function contained_float_spanset.
func ContainedFloatSpanset(d float64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_float_spanset(C.double(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedIntSet wraps MEOS C function contained_int_set.
func ContainedIntSet(i int, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_int_set(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedIntSpan wraps MEOS C function contained_int_span.
func ContainedIntSpan(i int, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_int_span(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedIntSpanset wraps MEOS C function contained_int_spanset.
func ContainedIntSpanset(i int, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_int_spanset(C.int(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedSetSet wraps MEOS C function contained_set_set.
func ContainedSetSet(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_set_set(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedSpanSpan wraps MEOS C function contained_span_span.
func ContainedSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedSpanSpanset wraps MEOS C function contained_span_spanset.
func ContainedSpanSpanset(s *Span, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_span_spanset(s._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedSpansetSpan wraps MEOS C function contained_spanset_span.
func ContainedSpansetSpan(ss *SpanSet, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_spanset_span(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedSpansetSpanset wraps MEOS C function contained_spanset_spanset.
func ContainedSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_spanset_spanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTextSet wraps MEOS C function contained_text_set.
func ContainedTextSet(txt string, s *Set) (_r0 bool, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.contained_text_set(_c_txt, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTimestamptzSet wraps MEOS C function contained_timestamptz_set.
func ContainedTimestamptzSet(t int64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_timestamptz_set(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTimestamptzSpan wraps MEOS C function contained_timestamptz_span.
func ContainedTimestamptzSpan(t int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_timestamptz_span(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTimestamptzSpanset wraps MEOS C function contained_timestamptz_spanset.
func ContainedTimestamptzSpanset(t int64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetBigint wraps MEOS C function contains_set_bigint.
func ContainsSetBigint(s *Set, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetDate wraps MEOS C function contains_set_date.
func ContainsSetDate(s *Set, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetFloat wraps MEOS C function contains_set_float.
func ContainsSetFloat(s *Set, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetInt wraps MEOS C function contains_set_int.
func ContainsSetInt(s *Set, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetSet wraps MEOS C function contains_set_set.
func ContainsSetSet(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_set(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetText wraps MEOS C function contains_set_text.
func ContainsSetText(s *Set, t string) (_r0 bool, _err error) {
	_c_t := C.cstring_to_text(C.CString(t))
	defer C.free(unsafe.Pointer(_c_t))
	C.meos_errno_reset()
	_cret := C.contains_set_text(s._inner, _c_t)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetTimestamptz wraps MEOS C function contains_set_timestamptz.
func ContainsSetTimestamptz(s *Set, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpanBigint wraps MEOS C function contains_span_bigint.
func ContainsSpanBigint(s *Span, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_span_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpanDate wraps MEOS C function contains_span_date.
func ContainsSpanDate(s *Span, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_span_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpanFloat wraps MEOS C function contains_span_float.
func ContainsSpanFloat(s *Span, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_span_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpanInt wraps MEOS C function contains_span_int.
func ContainsSpanInt(s *Span, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_span_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpanSpan wraps MEOS C function contains_span_span.
func ContainsSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpanSpanset wraps MEOS C function contains_span_spanset.
func ContainsSpanSpanset(s *Span, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_span_spanset(s._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpanTimestamptz wraps MEOS C function contains_span_timestamptz.
func ContainsSpanTimestamptz(s *Span, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_span_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpansetBigint wraps MEOS C function contains_spanset_bigint.
func ContainsSpansetBigint(ss *SpanSet, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_spanset_bigint(ss._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpansetDate wraps MEOS C function contains_spanset_date.
func ContainsSpansetDate(ss *SpanSet, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_spanset_date(ss._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpansetFloat wraps MEOS C function contains_spanset_float.
func ContainsSpansetFloat(ss *SpanSet, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_spanset_float(ss._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpansetInt wraps MEOS C function contains_spanset_int.
func ContainsSpansetInt(ss *SpanSet, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_spanset_int(ss._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpansetSpan wraps MEOS C function contains_spanset_span.
func ContainsSpansetSpan(ss *SpanSet, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_spanset_span(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpansetSpanset wraps MEOS C function contains_spanset_spanset.
func ContainsSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_spanset_spanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSpansetTimestamptz wraps MEOS C function contains_spanset_timestamptz.
func ContainsSpansetTimestamptz(ss *SpanSet, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsSetSet wraps MEOS C function overlaps_set_set.
func OverlapsSetSet(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_set_set(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsSpanSpan wraps MEOS C function overlaps_span_span.
func OverlapsSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsSpanSpanset wraps MEOS C function overlaps_span_spanset.
func OverlapsSpanSpanset(s *Span, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_span_spanset(s._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsSpansetSpan wraps MEOS C function overlaps_spanset_span.
func OverlapsSpansetSpan(ss *SpanSet, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_spanset_span(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsSpansetSpanset wraps MEOS C function overlaps_spanset_spanset.
func OverlapsSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_spanset_spanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameSpanSpan wraps MEOS C function same_span_span.
func SameSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterDateSet wraps MEOS C function after_date_set.
func AfterDateSet(d int32, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_date_set(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterDateSpan wraps MEOS C function after_date_span.
func AfterDateSpan(d int32, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_date_span(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterDateSpanset wraps MEOS C function after_date_spanset.
func AfterDateSpanset(d int32, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_date_spanset(C.DateADT(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterSetDate wraps MEOS C function after_set_date.
func AfterSetDate(s *Set, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_set_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterSetTimestamptz wraps MEOS C function after_set_timestamptz.
func AfterSetTimestamptz(s *Set, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_set_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterSpanDate wraps MEOS C function after_span_date.
func AfterSpanDate(s *Span, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_span_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterSpanTimestamptz wraps MEOS C function after_span_timestamptz.
func AfterSpanTimestamptz(s *Span, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_span_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterSpansetDate wraps MEOS C function after_spanset_date.
func AfterSpansetDate(ss *SpanSet, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_spanset_date(ss._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterSpansetTimestamptz wraps MEOS C function after_spanset_timestamptz.
func AfterSpansetTimestamptz(ss *SpanSet, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTimestamptzSet wraps MEOS C function after_timestamptz_set.
func AfterTimestamptzSet(t int64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_timestamptz_set(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTimestamptzSpan wraps MEOS C function after_timestamptz_span.
func AfterTimestamptzSpan(t int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_timestamptz_span(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTimestamptzSpanset wraps MEOS C function after_timestamptz_spanset.
func AfterTimestamptzSpanset(t int64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeDateSet wraps MEOS C function before_date_set.
func BeforeDateSet(d int32, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_date_set(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeDateSpan wraps MEOS C function before_date_span.
func BeforeDateSpan(d int32, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_date_span(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeDateSpanset wraps MEOS C function before_date_spanset.
func BeforeDateSpanset(d int32, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_date_spanset(C.DateADT(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeSetDate wraps MEOS C function before_set_date.
func BeforeSetDate(s *Set, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_set_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeSetTimestamptz wraps MEOS C function before_set_timestamptz.
func BeforeSetTimestamptz(s *Set, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_set_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeSpanDate wraps MEOS C function before_span_date.
func BeforeSpanDate(s *Span, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_span_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeSpanTimestamptz wraps MEOS C function before_span_timestamptz.
func BeforeSpanTimestamptz(s *Span, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_span_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeSpansetDate wraps MEOS C function before_spanset_date.
func BeforeSpansetDate(ss *SpanSet, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_spanset_date(ss._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeSpansetTimestamptz wraps MEOS C function before_spanset_timestamptz.
func BeforeSpansetTimestamptz(ss *SpanSet, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTimestamptzSet wraps MEOS C function before_timestamptz_set.
func BeforeTimestamptzSet(t int64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_timestamptz_set(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTimestamptzSpan wraps MEOS C function before_timestamptz_span.
func BeforeTimestamptzSpan(t int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_timestamptz_span(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTimestamptzSpanset wraps MEOS C function before_timestamptz_spanset.
func BeforeTimestamptzSpanset(t int64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftBigintSet wraps MEOS C function left_bigint_set.
func LeftBigintSet(i int64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_bigint_set(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftBigintSpan wraps MEOS C function left_bigint_span.
func LeftBigintSpan(i int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_bigint_span(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftBigintSpanset wraps MEOS C function left_bigint_spanset.
func LeftBigintSpanset(i int64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_bigint_spanset(C.int64_t(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftFloatSet wraps MEOS C function left_float_set.
func LeftFloatSet(d float64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_float_set(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftFloatSpan wraps MEOS C function left_float_span.
func LeftFloatSpan(d float64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_float_span(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftFloatSpanset wraps MEOS C function left_float_spanset.
func LeftFloatSpanset(d float64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_float_spanset(C.double(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftIntSet wraps MEOS C function left_int_set.
func LeftIntSet(i int, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_int_set(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftIntSpan wraps MEOS C function left_int_span.
func LeftIntSpan(i int, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_int_span(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftIntSpanset wraps MEOS C function left_int_spanset.
func LeftIntSpanset(i int, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_int_spanset(C.int(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSetBigint wraps MEOS C function left_set_bigint.
func LeftSetBigint(s *Set, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_set_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSetFloat wraps MEOS C function left_set_float.
func LeftSetFloat(s *Set, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_set_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSetInt wraps MEOS C function left_set_int.
func LeftSetInt(s *Set, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_set_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSetSet wraps MEOS C function left_set_set.
func LeftSetSet(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_set_set(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSetText wraps MEOS C function left_set_text.
func LeftSetText(s *Set, txt string) (_r0 bool, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.left_set_text(s._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSpanBigint wraps MEOS C function left_span_bigint.
func LeftSpanBigint(s *Span, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_span_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSpanFloat wraps MEOS C function left_span_float.
func LeftSpanFloat(s *Span, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_span_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSpanInt wraps MEOS C function left_span_int.
func LeftSpanInt(s *Span, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_span_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSpanSpan wraps MEOS C function left_span_span.
func LeftSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSpanSpanset wraps MEOS C function left_span_spanset.
func LeftSpanSpanset(s *Span, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_span_spanset(s._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSpansetBigint wraps MEOS C function left_spanset_bigint.
func LeftSpansetBigint(ss *SpanSet, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_spanset_bigint(ss._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSpansetFloat wraps MEOS C function left_spanset_float.
func LeftSpansetFloat(ss *SpanSet, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_spanset_float(ss._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSpansetInt wraps MEOS C function left_spanset_int.
func LeftSpansetInt(ss *SpanSet, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_spanset_int(ss._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSpansetSpan wraps MEOS C function left_spanset_span.
func LeftSpansetSpan(ss *SpanSet, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_spanset_span(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftSpansetSpanset wraps MEOS C function left_spanset_spanset.
func LeftSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_spanset_spanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTextSet wraps MEOS C function left_text_set.
func LeftTextSet(txt string, s *Set) (_r0 bool, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.left_text_set(_c_txt, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterDateSet wraps MEOS C function overafter_date_set.
func OverafterDateSet(d int32, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_date_set(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterDateSpan wraps MEOS C function overafter_date_span.
func OverafterDateSpan(d int32, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_date_span(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterDateSpanset wraps MEOS C function overafter_date_spanset.
func OverafterDateSpanset(d int32, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_date_spanset(C.DateADT(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterSetDate wraps MEOS C function overafter_set_date.
func OverafterSetDate(s *Set, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_set_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterSetTimestamptz wraps MEOS C function overafter_set_timestamptz.
func OverafterSetTimestamptz(s *Set, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_set_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterSpanDate wraps MEOS C function overafter_span_date.
func OverafterSpanDate(s *Span, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_span_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterSpanTimestamptz wraps MEOS C function overafter_span_timestamptz.
func OverafterSpanTimestamptz(s *Span, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_span_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterSpansetDate wraps MEOS C function overafter_spanset_date.
func OverafterSpansetDate(ss *SpanSet, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_spanset_date(ss._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterSpansetTimestamptz wraps MEOS C function overafter_spanset_timestamptz.
func OverafterSpansetTimestamptz(ss *SpanSet, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTimestamptzSet wraps MEOS C function overafter_timestamptz_set.
func OverafterTimestamptzSet(t int64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_timestamptz_set(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTimestamptzSpan wraps MEOS C function overafter_timestamptz_span.
func OverafterTimestamptzSpan(t int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_timestamptz_span(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTimestamptzSpanset wraps MEOS C function overafter_timestamptz_spanset.
func OverafterTimestamptzSpanset(t int64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeDateSet wraps MEOS C function overbefore_date_set.
func OverbeforeDateSet(d int32, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_date_set(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeDateSpan wraps MEOS C function overbefore_date_span.
func OverbeforeDateSpan(d int32, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_date_span(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeDateSpanset wraps MEOS C function overbefore_date_spanset.
func OverbeforeDateSpanset(d int32, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_date_spanset(C.DateADT(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeSetDate wraps MEOS C function overbefore_set_date.
func OverbeforeSetDate(s *Set, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_set_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeSetTimestamptz wraps MEOS C function overbefore_set_timestamptz.
func OverbeforeSetTimestamptz(s *Set, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_set_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeSpanDate wraps MEOS C function overbefore_span_date.
func OverbeforeSpanDate(s *Span, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_span_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeSpanTimestamptz wraps MEOS C function overbefore_span_timestamptz.
func OverbeforeSpanTimestamptz(s *Span, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_span_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeSpansetDate wraps MEOS C function overbefore_spanset_date.
func OverbeforeSpansetDate(ss *SpanSet, d int32) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_spanset_date(ss._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeSpansetTimestamptz wraps MEOS C function overbefore_spanset_timestamptz.
func OverbeforeSpansetTimestamptz(ss *SpanSet, t int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTimestamptzSet wraps MEOS C function overbefore_timestamptz_set.
func OverbeforeTimestamptzSet(t int64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_timestamptz_set(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTimestamptzSpan wraps MEOS C function overbefore_timestamptz_span.
func OverbeforeTimestamptzSpan(t int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_timestamptz_span(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTimestamptzSpanset wraps MEOS C function overbefore_timestamptz_spanset.
func OverbeforeTimestamptzSpanset(t int64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftBigintSet wraps MEOS C function overleft_bigint_set.
func OverleftBigintSet(i int64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_bigint_set(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftBigintSpan wraps MEOS C function overleft_bigint_span.
func OverleftBigintSpan(i int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_bigint_span(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftBigintSpanset wraps MEOS C function overleft_bigint_spanset.
func OverleftBigintSpanset(i int64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_bigint_spanset(C.int64_t(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftFloatSet wraps MEOS C function overleft_float_set.
func OverleftFloatSet(d float64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_float_set(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftFloatSpan wraps MEOS C function overleft_float_span.
func OverleftFloatSpan(d float64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_float_span(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftFloatSpanset wraps MEOS C function overleft_float_spanset.
func OverleftFloatSpanset(d float64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_float_spanset(C.double(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftIntSet wraps MEOS C function overleft_int_set.
func OverleftIntSet(i int, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_int_set(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftIntSpan wraps MEOS C function overleft_int_span.
func OverleftIntSpan(i int, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_int_span(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftIntSpanset wraps MEOS C function overleft_int_spanset.
func OverleftIntSpanset(i int, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_int_spanset(C.int(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSetBigint wraps MEOS C function overleft_set_bigint.
func OverleftSetBigint(s *Set, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_set_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSetFloat wraps MEOS C function overleft_set_float.
func OverleftSetFloat(s *Set, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_set_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSetInt wraps MEOS C function overleft_set_int.
func OverleftSetInt(s *Set, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_set_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSetSet wraps MEOS C function overleft_set_set.
func OverleftSetSet(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_set_set(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSetText wraps MEOS C function overleft_set_text.
func OverleftSetText(s *Set, txt string) (_r0 bool, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.overleft_set_text(s._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSpanBigint wraps MEOS C function overleft_span_bigint.
func OverleftSpanBigint(s *Span, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_span_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSpanFloat wraps MEOS C function overleft_span_float.
func OverleftSpanFloat(s *Span, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_span_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSpanInt wraps MEOS C function overleft_span_int.
func OverleftSpanInt(s *Span, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_span_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSpanSpan wraps MEOS C function overleft_span_span.
func OverleftSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSpanSpanset wraps MEOS C function overleft_span_spanset.
func OverleftSpanSpanset(s *Span, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_span_spanset(s._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSpansetBigint wraps MEOS C function overleft_spanset_bigint.
func OverleftSpansetBigint(ss *SpanSet, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_spanset_bigint(ss._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSpansetFloat wraps MEOS C function overleft_spanset_float.
func OverleftSpansetFloat(ss *SpanSet, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_spanset_float(ss._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSpansetInt wraps MEOS C function overleft_spanset_int.
func OverleftSpansetInt(ss *SpanSet, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_spanset_int(ss._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSpansetSpan wraps MEOS C function overleft_spanset_span.
func OverleftSpansetSpan(ss *SpanSet, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_spanset_span(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftSpansetSpanset wraps MEOS C function overleft_spanset_spanset.
func OverleftSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_spanset_spanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTextSet wraps MEOS C function overleft_text_set.
func OverleftTextSet(txt string, s *Set) (_r0 bool, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.overleft_text_set(_c_txt, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightBigintSet wraps MEOS C function overright_bigint_set.
func OverrightBigintSet(i int64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_bigint_set(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightBigintSpan wraps MEOS C function overright_bigint_span.
func OverrightBigintSpan(i int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_bigint_span(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightBigintSpanset wraps MEOS C function overright_bigint_spanset.
func OverrightBigintSpanset(i int64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_bigint_spanset(C.int64_t(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightFloatSet wraps MEOS C function overright_float_set.
func OverrightFloatSet(d float64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_float_set(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightFloatSpan wraps MEOS C function overright_float_span.
func OverrightFloatSpan(d float64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_float_span(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightFloatSpanset wraps MEOS C function overright_float_spanset.
func OverrightFloatSpanset(d float64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_float_spanset(C.double(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightIntSet wraps MEOS C function overright_int_set.
func OverrightIntSet(i int, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_int_set(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightIntSpan wraps MEOS C function overright_int_span.
func OverrightIntSpan(i int, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_int_span(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightIntSpanset wraps MEOS C function overright_int_spanset.
func OverrightIntSpanset(i int, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_int_spanset(C.int(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSetBigint wraps MEOS C function overright_set_bigint.
func OverrightSetBigint(s *Set, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_set_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSetFloat wraps MEOS C function overright_set_float.
func OverrightSetFloat(s *Set, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_set_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSetInt wraps MEOS C function overright_set_int.
func OverrightSetInt(s *Set, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_set_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSetSet wraps MEOS C function overright_set_set.
func OverrightSetSet(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_set_set(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSetText wraps MEOS C function overright_set_text.
func OverrightSetText(s *Set, txt string) (_r0 bool, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.overright_set_text(s._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSpanBigint wraps MEOS C function overright_span_bigint.
func OverrightSpanBigint(s *Span, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_span_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSpanFloat wraps MEOS C function overright_span_float.
func OverrightSpanFloat(s *Span, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_span_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSpanInt wraps MEOS C function overright_span_int.
func OverrightSpanInt(s *Span, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_span_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSpanSpan wraps MEOS C function overright_span_span.
func OverrightSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSpanSpanset wraps MEOS C function overright_span_spanset.
func OverrightSpanSpanset(s *Span, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_span_spanset(s._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSpansetBigint wraps MEOS C function overright_spanset_bigint.
func OverrightSpansetBigint(ss *SpanSet, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_spanset_bigint(ss._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSpansetFloat wraps MEOS C function overright_spanset_float.
func OverrightSpansetFloat(ss *SpanSet, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_spanset_float(ss._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSpansetInt wraps MEOS C function overright_spanset_int.
func OverrightSpansetInt(ss *SpanSet, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_spanset_int(ss._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSpansetSpan wraps MEOS C function overright_spanset_span.
func OverrightSpansetSpan(ss *SpanSet, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_spanset_span(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightSpansetSpanset wraps MEOS C function overright_spanset_spanset.
func OverrightSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_spanset_spanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTextSet wraps MEOS C function overright_text_set.
func OverrightTextSet(txt string, s *Set) (_r0 bool, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.overright_text_set(_c_txt, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightBigintSet wraps MEOS C function right_bigint_set.
func RightBigintSet(i int64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_bigint_set(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightBigintSpan wraps MEOS C function right_bigint_span.
func RightBigintSpan(i int64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_bigint_span(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightBigintSpanset wraps MEOS C function right_bigint_spanset.
func RightBigintSpanset(i int64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_bigint_spanset(C.int64_t(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightFloatSet wraps MEOS C function right_float_set.
func RightFloatSet(d float64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_float_set(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightFloatSpan wraps MEOS C function right_float_span.
func RightFloatSpan(d float64, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_float_span(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightFloatSpanset wraps MEOS C function right_float_spanset.
func RightFloatSpanset(d float64, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_float_spanset(C.double(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightIntSet wraps MEOS C function right_int_set.
func RightIntSet(i int, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_int_set(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightIntSpan wraps MEOS C function right_int_span.
func RightIntSpan(i int, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_int_span(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightIntSpanset wraps MEOS C function right_int_spanset.
func RightIntSpanset(i int, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_int_spanset(C.int(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSetBigint wraps MEOS C function right_set_bigint.
func RightSetBigint(s *Set, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_set_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSetFloat wraps MEOS C function right_set_float.
func RightSetFloat(s *Set, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_set_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSetInt wraps MEOS C function right_set_int.
func RightSetInt(s *Set, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_set_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSetSet wraps MEOS C function right_set_set.
func RightSetSet(s1 *Set, s2 *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_set_set(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSetText wraps MEOS C function right_set_text.
func RightSetText(s *Set, txt string) (_r0 bool, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.right_set_text(s._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSpanBigint wraps MEOS C function right_span_bigint.
func RightSpanBigint(s *Span, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_span_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSpanFloat wraps MEOS C function right_span_float.
func RightSpanFloat(s *Span, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_span_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSpanInt wraps MEOS C function right_span_int.
func RightSpanInt(s *Span, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_span_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSpanSpan wraps MEOS C function right_span_span.
func RightSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSpanSpanset wraps MEOS C function right_span_spanset.
func RightSpanSpanset(s *Span, ss *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_span_spanset(s._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSpansetBigint wraps MEOS C function right_spanset_bigint.
func RightSpansetBigint(ss *SpanSet, i int64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_spanset_bigint(ss._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSpansetFloat wraps MEOS C function right_spanset_float.
func RightSpansetFloat(ss *SpanSet, d float64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_spanset_float(ss._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSpansetInt wraps MEOS C function right_spanset_int.
func RightSpansetInt(ss *SpanSet, i int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_spanset_int(ss._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSpansetSpan wraps MEOS C function right_spanset_span.
func RightSpansetSpan(ss *SpanSet, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_spanset_span(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightSpansetSpanset wraps MEOS C function right_spanset_spanset.
func RightSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_spanset_spanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTextSet wraps MEOS C function right_text_set.
func RightTextSet(txt string, s *Set) (_r0 bool, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.right_text_set(_c_txt, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionBigintSet wraps MEOS C function intersection_bigint_set.
func IntersectionBigintSet(i int64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_bigint_set(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionDateSet wraps MEOS C function intersection_date_set.
func IntersectionDateSet(d int32, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_date_set(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionFloatSet wraps MEOS C function intersection_float_set.
func IntersectionFloatSet(d float64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_float_set(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionIntSet wraps MEOS C function intersection_int_set.
func IntersectionIntSet(i int, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_int_set(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetBigint wraps MEOS C function intersection_set_bigint.
func IntersectionSetBigint(s *Set, i int64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetDate wraps MEOS C function intersection_set_date.
func IntersectionSetDate(s *Set, d int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetFloat wraps MEOS C function intersection_set_float.
func IntersectionSetFloat(s *Set, d float64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetInt wraps MEOS C function intersection_set_int.
func IntersectionSetInt(s *Set, i int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetSet wraps MEOS C function intersection_set_set.
func IntersectionSetSet(s1 *Set, s2 *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_set(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetText wraps MEOS C function intersection_set_text.
func IntersectionSetText(s *Set, txt string) (_r0 *Set, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.intersection_set_text(s._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetTimestamptz wraps MEOS C function intersection_set_timestamptz.
func IntersectionSetTimestamptz(s *Set, t int64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSpanBigint wraps MEOS C function intersection_span_bigint.
func IntersectionSpanBigint(s *Span, i int64) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_span_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntersectionSpanDate wraps MEOS C function intersection_span_date.
func IntersectionSpanDate(s *Span, d int32) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_span_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntersectionSpanFloat wraps MEOS C function intersection_span_float.
func IntersectionSpanFloat(s *Span, d float64) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_span_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntersectionSpanInt wraps MEOS C function intersection_span_int.
func IntersectionSpanInt(s *Span, i int) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_span_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntersectionSpanSpan wraps MEOS C function intersection_span_span.
func IntersectionSpanSpan(s1 *Span, s2 *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntersectionSpanSpanset wraps MEOS C function intersection_span_spanset.
func IntersectionSpanSpanset(s *Span, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_span_spanset(s._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntersectionSpanTimestamptz wraps MEOS C function intersection_span_timestamptz.
func IntersectionSpanTimestamptz(s *Span, t int64) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_span_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntersectionSpansetBigint wraps MEOS C function intersection_spanset_bigint.
func IntersectionSpansetBigint(ss *SpanSet, i int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_spanset_bigint(ss._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntersectionSpansetDate wraps MEOS C function intersection_spanset_date.
func IntersectionSpansetDate(ss *SpanSet, d int32) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_spanset_date(ss._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntersectionSpansetFloat wraps MEOS C function intersection_spanset_float.
func IntersectionSpansetFloat(ss *SpanSet, d float64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_spanset_float(ss._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntersectionSpansetInt wraps MEOS C function intersection_spanset_int.
func IntersectionSpansetInt(ss *SpanSet, i int) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_spanset_int(ss._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntersectionSpansetSpan wraps MEOS C function intersection_spanset_span.
func IntersectionSpansetSpan(ss *SpanSet, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_spanset_span(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntersectionSpansetSpanset wraps MEOS C function intersection_spanset_spanset.
func IntersectionSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_spanset_spanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntersectionSpansetTimestamptz wraps MEOS C function intersection_spanset_timestamptz.
func IntersectionSpansetTimestamptz(ss *SpanSet, t int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// IntersectionTextSet wraps MEOS C function intersection_text_set.
func IntersectionTextSet(txt string, s *Set) (_r0 *Set, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.intersection_text_set(_c_txt, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionTimestamptzSet wraps MEOS C function intersection_timestamptz_set.
func IntersectionTimestamptzSet(t int64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_timestamptz_set(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusBigintSet wraps MEOS C function minus_bigint_set.
func MinusBigintSet(i int64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_bigint_set(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusBigintSpan wraps MEOS C function minus_bigint_span.
func MinusBigintSpan(i int64, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_bigint_span(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusBigintSpanset wraps MEOS C function minus_bigint_spanset.
func MinusBigintSpanset(i int64, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_bigint_spanset(C.int64_t(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusDateSet wraps MEOS C function minus_date_set.
func MinusDateSet(d int32, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_date_set(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusDateSpan wraps MEOS C function minus_date_span.
func MinusDateSpan(d int32, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_date_span(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusDateSpanset wraps MEOS C function minus_date_spanset.
func MinusDateSpanset(d int32, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_date_spanset(C.DateADT(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusFloatSet wraps MEOS C function minus_float_set.
func MinusFloatSet(d float64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_float_set(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusFloatSpan wraps MEOS C function minus_float_span.
func MinusFloatSpan(d float64, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_float_span(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusFloatSpanset wraps MEOS C function minus_float_spanset.
func MinusFloatSpanset(d float64, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_float_spanset(C.double(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusIntSet wraps MEOS C function minus_int_set.
func MinusIntSet(i int, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int_set(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusIntSpan wraps MEOS C function minus_int_span.
func MinusIntSpan(i int, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int_span(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusIntSpanset wraps MEOS C function minus_int_spanset.
func MinusIntSpanset(i int, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_int_spanset(C.int(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSetBigint wraps MEOS C function minus_set_bigint.
func MinusSetBigint(s *Set, i int64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetDate wraps MEOS C function minus_set_date.
func MinusSetDate(s *Set, d int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetFloat wraps MEOS C function minus_set_float.
func MinusSetFloat(s *Set, d float64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetInt wraps MEOS C function minus_set_int.
func MinusSetInt(s *Set, i int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetSet wraps MEOS C function minus_set_set.
func MinusSetSet(s1 *Set, s2 *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_set(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetText wraps MEOS C function minus_set_text.
func MinusSetText(s *Set, txt string) (_r0 *Set, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.minus_set_text(s._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetTimestamptz wraps MEOS C function minus_set_timestamptz.
func MinusSetTimestamptz(s *Set, t int64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSpanBigint wraps MEOS C function minus_span_bigint.
func MinusSpanBigint(s *Span, i int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_span_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpanDate wraps MEOS C function minus_span_date.
func MinusSpanDate(s *Span, d int32) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_span_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpanFloat wraps MEOS C function minus_span_float.
func MinusSpanFloat(s *Span, d float64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_span_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpanInt wraps MEOS C function minus_span_int.
func MinusSpanInt(s *Span, i int) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_span_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpanSpan wraps MEOS C function minus_span_span.
func MinusSpanSpan(s1 *Span, s2 *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpanSpanset wraps MEOS C function minus_span_spanset.
func MinusSpanSpanset(s *Span, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_span_spanset(s._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpanTimestamptz wraps MEOS C function minus_span_timestamptz.
func MinusSpanTimestamptz(s *Span, t int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_span_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpansetBigint wraps MEOS C function minus_spanset_bigint.
func MinusSpansetBigint(ss *SpanSet, i int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_spanset_bigint(ss._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpansetDate wraps MEOS C function minus_spanset_date.
func MinusSpansetDate(ss *SpanSet, d int32) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_spanset_date(ss._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpansetFloat wraps MEOS C function minus_spanset_float.
func MinusSpansetFloat(ss *SpanSet, d float64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_spanset_float(ss._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpansetInt wraps MEOS C function minus_spanset_int.
func MinusSpansetInt(ss *SpanSet, i int) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_spanset_int(ss._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpansetSpan wraps MEOS C function minus_spanset_span.
func MinusSpansetSpan(ss *SpanSet, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_spanset_span(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpansetSpanset wraps MEOS C function minus_spanset_spanset.
func MinusSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_spanset_spanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusSpansetTimestamptz wraps MEOS C function minus_spanset_timestamptz.
func MinusSpansetTimestamptz(ss *SpanSet, t int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusTextSet wraps MEOS C function minus_text_set.
func MinusTextSet(txt string, s *Set) (_r0 *Set, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.minus_text_set(_c_txt, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusTimestamptzSet wraps MEOS C function minus_timestamptz_set.
func MinusTimestamptzSet(t int64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_timestamptz_set(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusTimestamptzSpan wraps MEOS C function minus_timestamptz_span.
func MinusTimestamptzSpan(t int64, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_timestamptz_span(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// MinusTimestamptzSpanset wraps MEOS C function minus_timestamptz_spanset.
func MinusTimestamptzSpanset(t int64, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionBigintSet wraps MEOS C function union_bigint_set.
func UnionBigintSet(i int64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_bigint_set(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionBigintSpan wraps MEOS C function union_bigint_span.
func UnionBigintSpan(i int64, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_bigint_span(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionBigintSpanset wraps MEOS C function union_bigint_spanset.
func UnionBigintSpanset(i int64, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_bigint_spanset(C.int64_t(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionDateSet wraps MEOS C function union_date_set.
func UnionDateSet(d int32, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_date_set(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionDateSpan wraps MEOS C function union_date_span.
func UnionDateSpan(d int32, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_date_span(C.DateADT(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionDateSpanset wraps MEOS C function union_date_spanset.
func UnionDateSpanset(d int32, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_date_spanset(C.DateADT(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionFloatSet wraps MEOS C function union_float_set.
func UnionFloatSet(d float64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_float_set(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionFloatSpan wraps MEOS C function union_float_span.
func UnionFloatSpan(d float64, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_float_span(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionFloatSpanset wraps MEOS C function union_float_spanset.
func UnionFloatSpanset(d float64, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_float_spanset(C.double(d), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionIntSet wraps MEOS C function union_int_set.
func UnionIntSet(i int, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_int_set(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionIntSpan wraps MEOS C function union_int_span.
func UnionIntSpan(i int, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_int_span(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionIntSpanset wraps MEOS C function union_int_spanset.
func UnionIntSpanset(i int, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_int_spanset(C.int(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSetBigint wraps MEOS C function union_set_bigint.
func UnionSetBigint(s *Set, i int64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetDate wraps MEOS C function union_set_date.
func UnionSetDate(s *Set, d int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetFloat wraps MEOS C function union_set_float.
func UnionSetFloat(s *Set, d float64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetInt wraps MEOS C function union_set_int.
func UnionSetInt(s *Set, i int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetSet wraps MEOS C function union_set_set.
func UnionSetSet(s1 *Set, s2 *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_set(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetText wraps MEOS C function union_set_text.
func UnionSetText(s *Set, txt string) (_r0 *Set, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.gunion_set_text(s._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetTimestamptz wraps MEOS C function union_set_timestamptz.
func UnionSetTimestamptz(s *Set, t int64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSpanBigint wraps MEOS C function union_span_bigint.
func UnionSpanBigint(s *Span, i int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_span_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpanDate wraps MEOS C function union_span_date.
func UnionSpanDate(s *Span, d int32) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_span_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpanFloat wraps MEOS C function union_span_float.
func UnionSpanFloat(s *Span, d float64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_span_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpanInt wraps MEOS C function union_span_int.
func UnionSpanInt(s *Span, i int) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_span_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpanSpan wraps MEOS C function union_span_span.
func UnionSpanSpan(s1 *Span, s2 *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpanSpanset wraps MEOS C function union_span_spanset.
func UnionSpanSpanset(s *Span, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_span_spanset(s._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpanTimestamptz wraps MEOS C function union_span_timestamptz.
func UnionSpanTimestamptz(s *Span, t int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_span_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpansetBigint wraps MEOS C function union_spanset_bigint.
func UnionSpansetBigint(ss *SpanSet, i int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_spanset_bigint(ss._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpansetDate wraps MEOS C function union_spanset_date.
func UnionSpansetDate(ss *SpanSet, d int32) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_spanset_date(ss._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpansetFloat wraps MEOS C function union_spanset_float.
func UnionSpansetFloat(ss *SpanSet, d float64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_spanset_float(ss._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpansetInt wraps MEOS C function union_spanset_int.
func UnionSpansetInt(ss *SpanSet, i int) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_spanset_int(ss._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpansetSpan wraps MEOS C function union_spanset_span.
func UnionSpansetSpan(ss *SpanSet, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_spanset_span(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpansetSpanset wraps MEOS C function union_spanset_spanset.
func UnionSpansetSpanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_spanset_spanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionSpansetTimestamptz wraps MEOS C function union_spanset_timestamptz.
func UnionSpansetTimestamptz(ss *SpanSet, t int64) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionTextSet wraps MEOS C function union_text_set.
func UnionTextSet(txt string, s *Set) (_r0 *Set, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.gunion_text_set(_c_txt, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionTimestamptzSet wraps MEOS C function union_timestamptz_set.
func UnionTimestamptzSet(t int64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_timestamptz_set(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionTimestamptzSpan wraps MEOS C function union_timestamptz_span.
func UnionTimestamptzSpan(t int64, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_timestamptz_span(C.TimestampTz(t), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// UnionTimestamptzSpanset wraps MEOS C function union_timestamptz_spanset.
func UnionTimestamptzSpanset(t int64, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_timestamptz_spanset(C.TimestampTz(t), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// DistanceBigintsetBigintset wraps MEOS C function distance_bigintset_bigintset.
func DistanceBigintsetBigintset(s1 *Set, s2 *Set) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_bigintset_bigintset(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DistanceBigintspanBigintspan wraps MEOS C function distance_bigintspan_bigintspan.
func DistanceBigintspanBigintspan(s1 *Span, s2 *Span) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_bigintspan_bigintspan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DistanceBigintspansetBigintspan wraps MEOS C function distance_bigintspanset_bigintspan.
func DistanceBigintspansetBigintspan(ss *SpanSet, s *Span) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_bigintspanset_bigintspan(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DistanceBigintspansetBigintspanset wraps MEOS C function distance_bigintspanset_bigintspanset.
func DistanceBigintspansetBigintspanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_bigintspanset_bigintspanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DistanceDatesetDateset wraps MEOS C function distance_dateset_dateset.
func DistanceDatesetDateset(s1 *Set, s2 *Set) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_dateset_dateset(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceDatespanDatespan wraps MEOS C function distance_datespan_datespan.
func DistanceDatespanDatespan(s1 *Span, s2 *Span) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_datespan_datespan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceDatespansetDatespan wraps MEOS C function distance_datespanset_datespan.
func DistanceDatespansetDatespan(ss *SpanSet, s *Span) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_datespanset_datespan(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceDatespansetDatespanset wraps MEOS C function distance_datespanset_datespanset.
func DistanceDatespansetDatespanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_datespanset_datespanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceFloatsetFloatset wraps MEOS C function distance_floatset_floatset.
func DistanceFloatsetFloatset(s1 *Set, s2 *Set) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_floatset_floatset(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceFloatspanFloatspan wraps MEOS C function distance_floatspan_floatspan.
func DistanceFloatspanFloatspan(s1 *Span, s2 *Span) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_floatspan_floatspan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceFloatspansetFloatspan wraps MEOS C function distance_floatspanset_floatspan.
func DistanceFloatspansetFloatspan(ss *SpanSet, s *Span) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_floatspanset_floatspan(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceFloatspansetFloatspanset wraps MEOS C function distance_floatspanset_floatspanset.
func DistanceFloatspansetFloatspanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_floatspanset_floatspanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceIntsetIntset wraps MEOS C function distance_intset_intset.
func DistanceIntsetIntset(s1 *Set, s2 *Set) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_intset_intset(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceIntspanIntspan wraps MEOS C function distance_intspan_intspan.
func DistanceIntspanIntspan(s1 *Span, s2 *Span) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_intspan_intspan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceIntspansetIntspan wraps MEOS C function distance_intspanset_intspan.
func DistanceIntspansetIntspan(ss *SpanSet, s *Span) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_intspanset_intspan(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceIntspansetIntspanset wraps MEOS C function distance_intspanset_intspanset.
func DistanceIntspansetIntspanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_intspanset_intspanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceSetBigint wraps MEOS C function distance_set_bigint.
func DistanceSetBigint(s *Set, i int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_set_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DistanceSetDate wraps MEOS C function distance_set_date.
func DistanceSetDate(s *Set, d int32) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_set_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceSetFloat wraps MEOS C function distance_set_float.
func DistanceSetFloat(s *Set, d float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_set_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceSetInt wraps MEOS C function distance_set_int.
func DistanceSetInt(s *Set, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_set_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceSetTimestamptz wraps MEOS C function distance_set_timestamptz.
func DistanceSetTimestamptz(s *Set, t int64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_set_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceSpanBigint wraps MEOS C function distance_span_bigint.
func DistanceSpanBigint(s *Span, i int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_span_bigint(s._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DistanceSpanDate wraps MEOS C function distance_span_date.
func DistanceSpanDate(s *Span, d int32) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_span_date(s._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceSpanFloat wraps MEOS C function distance_span_float.
func DistanceSpanFloat(s *Span, d float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_span_float(s._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceSpanInt wraps MEOS C function distance_span_int.
func DistanceSpanInt(s *Span, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_span_int(s._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceSpanTimestamptz wraps MEOS C function distance_span_timestamptz.
func DistanceSpanTimestamptz(s *Span, t int64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_span_timestamptz(s._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceSpansetBigint wraps MEOS C function distance_spanset_bigint.
func DistanceSpansetBigint(ss *SpanSet, i int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_spanset_bigint(ss._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// DistanceSpansetDate wraps MEOS C function distance_spanset_date.
func DistanceSpansetDate(ss *SpanSet, d int32) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_spanset_date(ss._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceSpansetFloat wraps MEOS C function distance_spanset_float.
func DistanceSpansetFloat(ss *SpanSet, d float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_spanset_float(ss._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceSpansetInt wraps MEOS C function distance_spanset_int.
func DistanceSpansetInt(ss *SpanSet, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_spanset_int(ss._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// DistanceSpansetTimestamptz wraps MEOS C function distance_spanset_timestamptz.
func DistanceSpansetTimestamptz(ss *SpanSet, t int64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_spanset_timestamptz(ss._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceTstzsetTstzset wraps MEOS C function distance_tstzset_tstzset.
func DistanceTstzsetTstzset(s1 *Set, s2 *Set) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_tstzset_tstzset(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceTstzspanTstzspan wraps MEOS C function distance_tstzspan_tstzspan.
func DistanceTstzspanTstzspan(s1 *Span, s2 *Span) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_tstzspan_tstzspan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceTstzspansetTstzspan wraps MEOS C function distance_tstzspanset_tstzspan.
func DistanceTstzspansetTstzspan(ss *SpanSet, s *Span) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_tstzspanset_tstzspan(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// DistanceTstzspansetTstzspanset wraps MEOS C function distance_tstzspanset_tstzspanset.
func DistanceTstzspansetTstzspanset(ss1 *SpanSet, ss2 *SpanSet) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.distance_tstzspanset_tstzspanset(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// BigintExtentTransfn wraps MEOS C function bigint_extent_transfn.
func BigintExtentTransfn(state *Span, i int64) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_extent_transfn(state._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// BigintUnionTransfn wraps MEOS C function bigint_union_transfn.
func BigintUnionTransfn(state *Set, i int64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_union_transfn(state._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// DateExtentTransfn wraps MEOS C function date_extent_transfn.
func DateExtentTransfn(state *Span, d int32) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.date_extent_transfn(state._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// DateUnionTransfn wraps MEOS C function date_union_transfn.
func DateUnionTransfn(state *Set, d int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.date_union_transfn(state._inner, C.DateADT(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// FloatExtentTransfn wraps MEOS C function float_extent_transfn.
func FloatExtentTransfn(state *Span, d float64) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.float_extent_transfn(state._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatUnionTransfn wraps MEOS C function float_union_transfn.
func FloatUnionTransfn(state *Set, d float64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.float_union_transfn(state._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntExtentTransfn wraps MEOS C function int_extent_transfn.
func IntExtentTransfn(state *Span, i int) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.int_extent_transfn(state._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntUnionTransfn wraps MEOS C function int_union_transfn.
func IntUnionTransfn(state *Set, i int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.int_union_transfn(state._inner, C.int32(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// SetExtentTransfn wraps MEOS C function set_extent_transfn.
func SetExtentTransfn(state *Span, s *Set) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.set_extent_transfn(state._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SetUnionFinalfn wraps MEOS C function set_union_finalfn.
func SetUnionFinalfn(state *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.set_union_finalfn(state._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// SetUnionTransfn wraps MEOS C function set_union_transfn.
func SetUnionTransfn(state *Set, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.set_union_transfn(state._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// SpanExtentTransfn wraps MEOS C function span_extent_transfn.
func SpanExtentTransfn(state *Span, s *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.span_extent_transfn(state._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpanUnionTransfn wraps MEOS C function span_union_transfn.
func SpanUnionTransfn(state *SpanSet, s *Span) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.span_union_transfn(state._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// SpansetExtentTransfn wraps MEOS C function spanset_extent_transfn.
func SpansetExtentTransfn(state *Span, ss *SpanSet) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_extent_transfn(state._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpansetUnionFinalfn wraps MEOS C function spanset_union_finalfn.
func SpansetUnionFinalfn(state *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_union_finalfn(state._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// SpansetUnionTransfn wraps MEOS C function spanset_union_transfn.
func SpansetUnionTransfn(state *SpanSet, ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_union_transfn(state._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TextUnionTransfn wraps MEOS C function text_union_transfn.
func TextUnionTransfn(state *Set, txt string) (_r0 *Set, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_union_transfn(state._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TimestamptzExtentTransfn wraps MEOS C function timestamptz_extent_transfn.
func TimestamptzExtentTransfn(state *Span, t int64) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_extent_transfn(state._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TimestamptzUnionTransfn wraps MEOS C function timestamptz_union_transfn.
func TimestamptzUnionTransfn(state *Set, t int64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_union_transfn(state._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// BigintGetBin wraps MEOS C function bigint_get_bin.
func BigintGetBin(value int64, vsize int64, vorigin int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_get_bin(C.int64_t(value), C.int64_t(vsize), C.int64_t(vorigin))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// BigintspanBins wraps MEOS C function bigintspan_bins.
func BigintspanBins(s *Span, vsize int64, vorigin int64, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspan_bins(s._inner, C.int64_t(vsize), C.int64_t(vorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// BigintspansetBins wraps MEOS C function bigintspanset_bins.
func BigintspansetBins(ss *SpanSet, vsize int64, vorigin int64, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.bigintspanset_bins(ss._inner, C.int64_t(vsize), C.int64_t(vorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// DateGetBin wraps MEOS C function date_get_bin.
func DateGetBin(d int32, duration *Interval, torigin int32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.date_get_bin(C.DateADT(d), duration._inner, C.DateADT(torigin))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// DatespanBins wraps MEOS C function datespan_bins.
func DatespanBins(s *Span, duration *Interval, torigin int32, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.datespan_bins(s._inner, duration._inner, C.DateADT(torigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// DatespansetBins wraps MEOS C function datespanset_bins.
func DatespansetBins(ss *SpanSet, duration *Interval, torigin int32, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.datespanset_bins(ss._inner, duration._inner, C.DateADT(torigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatGetBin wraps MEOS C function float_get_bin.
func FloatGetBin(value float64, vsize float64, vorigin float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float_get_bin(C.double(value), C.double(vsize), C.double(vorigin))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// FloatspanBins wraps MEOS C function floatspan_bins.
func FloatspanBins(s *Span, vsize float64, vorigin float64, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspan_bins(s._inner, C.double(vsize), C.double(vorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// FloatspansetBins wraps MEOS C function floatspanset_bins.
func FloatspansetBins(ss *SpanSet, vsize float64, vorigin float64, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.floatspanset_bins(ss._inner, C.double(vsize), C.double(vorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntGetBin wraps MEOS C function int_get_bin.
func IntGetBin(value int, vsize int, vorigin int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.int_get_bin(C.int(value), C.int(vsize), C.int(vorigin))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// IntspanBins wraps MEOS C function intspan_bins.
func IntspanBins(s *Span, vsize int, vorigin int, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intspan_bins(s._inner, C.int(vsize), C.int(vorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// IntspansetBins wraps MEOS C function intspanset_bins.
func IntspansetBins(ss *SpanSet, vsize int, vorigin int, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.intspanset_bins(ss._inner, C.int(vsize), C.int(vorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TimestamptzGetBin wraps MEOS C function timestamptz_get_bin.
func TimestamptzGetBin(t int64, duration *Interval, torigin int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_get_bin(C.TimestampTz(t), duration._inner, C.TimestampTz(torigin))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TstzspanBins wraps MEOS C function tstzspan_bins.
func TstzspanBins(s *Span, duration *Interval, origin int64, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_bins(s._inner, duration._inner, C.TimestampTz(origin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TstzspansetBins wraps MEOS C function tstzspanset_bins.
func TstzspansetBins(ss *SpanSet, duration *Interval, torigin int64, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_bins(ss._inner, duration._inner, C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TBOXAsHexwkb wraps MEOS C function tbox_as_hexwkb.
func TBOXAsHexwkb(box *TBox, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_as_hexwkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TBOXAsWKB wraps MEOS C function tbox_as_wkb.
func TBOXAsWKB(box *TBox, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_as_wkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TBOXFromHexwkb wraps MEOS C function tbox_from_hexwkb.
func TBOXFromHexwkb(hexwkb string) (_r0 *TBox, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.tbox_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TBOXFromWKB wraps MEOS C function tbox_from_wkb.
func TBOXFromWKB(wkb unsafe.Pointer, size uint) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TBOXIn wraps MEOS C function tbox_in.
func TBOXIn(str string) (_r0 *TBox, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tbox_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TBOXOut wraps MEOS C function tbox_out.
func TBOXOut(box *TBox, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_out(box._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// FloatTimestamptzToTBOX wraps MEOS C function float_timestamptz_to_tbox.
func FloatTimestamptzToTBOX(d float64, t int64) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.float_timestamptz_to_tbox(C.double(d), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// FloatTstzspanToTBOX wraps MEOS C function float_tstzspan_to_tbox.
func FloatTstzspanToTBOX(d float64, s *Span) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.float_tstzspan_to_tbox(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// IntTimestamptzToTBOX wraps MEOS C function int_timestamptz_to_tbox.
func IntTimestamptzToTBOX(i int, t int64) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.int_timestamptz_to_tbox(C.int(i), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// BigintTimestamptzToTBOX wraps MEOS C function bigint_timestamptz_to_tbox.
func BigintTimestamptzToTBOX(i int64, t int64) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_timestamptz_to_tbox(C.int64_t(i), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// IntTstzspanToTBOX wraps MEOS C function int_tstzspan_to_tbox.
func IntTstzspanToTBOX(i int, s *Span) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.int_tstzspan_to_tbox(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// BigintTstzspanToTBOX wraps MEOS C function bigint_tstzspan_to_tbox.
func BigintTstzspanToTBOX(i int64, s *Span) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_tstzspan_to_tbox(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// NumspanTstzspanToTBOX wraps MEOS C function numspan_tstzspan_to_tbox.
func NumspanTstzspanToTBOX(span *Span, s *Span) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.numspan_tstzspan_to_tbox(span._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// NumspanTimestamptzToTBOX wraps MEOS C function numspan_timestamptz_to_tbox.
func NumspanTimestamptzToTBOX(span *Span, t int64) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.numspan_timestamptz_to_tbox(span._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TBOXCopy wraps MEOS C function tbox_copy.
func TBOXCopy(box *TBox) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_copy(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TBOXMake wraps MEOS C function tbox_make.
func TBOXMake(s *Span, p *Span) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_make(s._inner, p._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// FloatToTBOX wraps MEOS C function float_to_tbox.
func FloatToTBOX(d float64) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.float_to_tbox(C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// IntToTBOX wraps MEOS C function int_to_tbox.
func IntToTBOX(i int) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.int_to_tbox(C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// BigintToTBOX wraps MEOS C function bigint_to_tbox.
func BigintToTBOX(i int64) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_to_tbox(C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// SetToTBOX wraps MEOS C function set_to_tbox.
func SetToTBOX(s *Set) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.set_to_tbox(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// SpanToTBOX wraps MEOS C function span_to_tbox.
func SpanToTBOX(s *Span) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.span_to_tbox(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// SpansetToTBOX wraps MEOS C function spanset_to_tbox.
func SpansetToTBOX(ss *SpanSet) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_to_tbox(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TBOXToIntspan wraps MEOS C function tbox_to_intspan.
func TBOXToIntspan(box *TBox) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_to_intspan(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TBOXToBigintspan wraps MEOS C function tbox_to_bigintspan.
func TBOXToBigintspan(box *TBox) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_to_bigintspan(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TBOXToFloatspan wraps MEOS C function tbox_to_floatspan.
func TBOXToFloatspan(box *TBox) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_to_floatspan(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TBOXToTstzspan wraps MEOS C function tbox_to_tstzspan.
func TBOXToTstzspan(box *TBox) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_to_tstzspan(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TimestamptzToTBOX wraps MEOS C function timestamptz_to_tbox.
func TimestamptzToTBOX(t int64) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_to_tbox(C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TBOXHash wraps MEOS C function tbox_hash.
func TBOXHash(box *TBox) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_hash(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TBOXHashExtended wraps MEOS C function tbox_hash_extended.
func TBOXHashExtended(box *TBox, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_hash_extended(box._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TBOXHast wraps MEOS C function tbox_hast.
func TBOXHast(box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_hast(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXHasx wraps MEOS C function tbox_hasx.
func TBOXHasx(box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_hasx(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXTmax wraps MEOS C function tbox_tmax.
func TBOXTmax(box *TBox) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.TimestampTz
	C.meos_errno_reset()
	_cret := C.tbox_tmax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// TBOXTmaxInc wraps MEOS C function tbox_tmax_inc.
func TBOXTmaxInc(box *TBox) (_r0 bool, _r1 bool, _err error) {
	var _out_result C.bool
	C.meos_errno_reset()
	_cret := C.tbox_tmax_inc(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), bool(_out_result), nil
}


// TBOXTmin wraps MEOS C function tbox_tmin.
func TBOXTmin(box *TBox) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.TimestampTz
	C.meos_errno_reset()
	_cret := C.tbox_tmin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// TBOXTminInc wraps MEOS C function tbox_tmin_inc.
func TBOXTminInc(box *TBox) (_r0 bool, _r1 bool, _err error) {
	var _out_result C.bool
	C.meos_errno_reset()
	_cret := C.tbox_tmin_inc(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), bool(_out_result), nil
}


// TBOXXmax wraps MEOS C function tbox_xmax.
func TBOXXmax(box *TBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tbox_xmax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TBOXXmaxInc wraps MEOS C function tbox_xmax_inc.
func TBOXXmaxInc(box *TBox) (_r0 bool, _r1 bool, _err error) {
	var _out_result C.bool
	C.meos_errno_reset()
	_cret := C.tbox_xmax_inc(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), bool(_out_result), nil
}


// TBOXXmin wraps MEOS C function tbox_xmin.
func TBOXXmin(box *TBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tbox_xmin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TBOXXminInc wraps MEOS C function tbox_xmin_inc.
func TBOXXminInc(box *TBox) (_r0 bool, _r1 bool, _err error) {
	var _out_result C.bool
	C.meos_errno_reset()
	_cret := C.tbox_xmin_inc(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), bool(_out_result), nil
}


// TboxfloatXmax wraps MEOS C function tboxfloat_xmax.
func TboxfloatXmax(box *TBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tboxfloat_xmax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TboxfloatXmin wraps MEOS C function tboxfloat_xmin.
func TboxfloatXmin(box *TBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tboxfloat_xmin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TboxintXmax wraps MEOS C function tboxint_xmax.
func TboxintXmax(box *TBox) (_r0 bool, _r1 int, _err error) {
	var _out_result C.int
	C.meos_errno_reset()
	_cret := C.tboxint_xmax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int(_out_result), nil
}


// TboxbigintXmax wraps MEOS C function tboxbigint_xmax.
func TboxbigintXmax(box *TBox) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.int64_t
	C.meos_errno_reset()
	_cret := C.tboxbigint_xmax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// TboxintXmin wraps MEOS C function tboxint_xmin.
func TboxintXmin(box *TBox) (_r0 bool, _r1 int, _err error) {
	var _out_result C.int
	C.meos_errno_reset()
	_cret := C.tboxint_xmin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int(_out_result), nil
}


// TboxbigintXmin wraps MEOS C function tboxbigint_xmin.
func TboxbigintXmin(box *TBox) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.int64_t
	C.meos_errno_reset()
	_cret := C.tboxbigint_xmin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// TfloatboxExpand wraps MEOS C function tfloatbox_expand.
func TfloatboxExpand(box *TBox, d float64) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloatbox_expand(box._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TintboxExpand wraps MEOS C function tintbox_expand.
func TintboxExpand(box *TBox, i int) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tintbox_expand(box._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TBOXExpandTime wraps MEOS C function tbox_expand_time.
func TBOXExpandTime(box *TBox, interv *Interval) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_expand_time(box._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TBOXRound wraps MEOS C function tbox_round.
func TBOXRound(box *TBox, maxdd int) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_round(box._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TfloatboxShiftScale wraps MEOS C function tfloatbox_shift_scale.
func TfloatboxShiftScale(box *TBox, shift float64, width float64, hasshift bool, haswidth bool) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloatbox_shift_scale(box._inner, C.double(shift), C.double(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TintboxShiftScale wraps MEOS C function tintbox_shift_scale.
func TintboxShiftScale(box *TBox, shift int, width int, hasshift bool, haswidth bool) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tintbox_shift_scale(box._inner, C.int(shift), C.int(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TBOXShiftScaleTime wraps MEOS C function tbox_shift_scale_time.
func TBOXShiftScaleTime(box *TBox, shift *Interval, duration *Interval) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_shift_scale_time(box._inner, shift._inner, duration._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TbigintboxExpand wraps MEOS C function tbigintbox_expand.
func TbigintboxExpand(box *TBox, i int64) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigintbox_expand(box._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TbigintboxShiftScale wraps MEOS C function tbigintbox_shift_scale.
func TbigintboxShiftScale(box *TBox, shift int64, width int64, hasshift bool, haswidth bool) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigintbox_shift_scale(box._inner, C.int64_t(shift), C.int64_t(width), C.bool(hasshift), C.bool(haswidth))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// UnionTBOXTBOX wraps MEOS C function union_tbox_tbox.
func UnionTBOXTBOX(box1 *TBox, box2 *TBox, strict bool) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_tbox_tbox(box1._inner, box2._inner, C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// IntersectionTBOXTBOX wraps MEOS C function intersection_tbox_tbox.
func IntersectionTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// AdjacentTBOXTBOX wraps MEOS C function adjacent_tbox_tbox.
func AdjacentTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTBOXTBOX wraps MEOS C function contained_tbox_tbox.
func ContainedTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTBOXTBOX wraps MEOS C function contains_tbox_tbox.
func ContainsTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTBOXTBOX wraps MEOS C function overlaps_tbox_tbox.
func OverlapsTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTBOXTBOX wraps MEOS C function same_tbox_tbox.
func SameTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTBOXTBOX wraps MEOS C function after_tbox_tbox.
func AfterTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTBOXTBOX wraps MEOS C function before_tbox_tbox.
func BeforeTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTBOXTBOX wraps MEOS C function left_tbox_tbox.
func LeftTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTBOXTBOX wraps MEOS C function overafter_tbox_tbox.
func OverafterTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTBOXTBOX wraps MEOS C function overbefore_tbox_tbox.
func OverbeforeTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTBOXTBOX wraps MEOS C function overleft_tbox_tbox.
func OverleftTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTBOXTBOX wraps MEOS C function overright_tbox_tbox.
func OverrightTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTBOXTBOX wraps MEOS C function right_tbox_tbox.
func RightTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_tbox_tbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXCmp wraps MEOS C function tbox_cmp.
func TBOXCmp(box1 *TBox, box2 *TBox) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_cmp(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TBOXEq wraps MEOS C function tbox_eq.
func TBOXEq(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_eq(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXGe wraps MEOS C function tbox_ge.
func TBOXGe(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_ge(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXGt wraps MEOS C function tbox_gt.
func TBOXGt(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_gt(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXLe wraps MEOS C function tbox_le.
func TBOXLe(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_le(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXLt wraps MEOS C function tbox_lt.
func TBOXLt(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_lt(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXNe wraps MEOS C function tbox_ne.
func TBOXNe(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_ne(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TboolFromMFJSON wraps MEOS C function tbool_from_mfjson.
func TboolFromMFJSON(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tbool_from_mfjson(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TboolIn wraps MEOS C function tbool_in.
func TboolIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tbool_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TboolOut wraps MEOS C function tbool_out.
func TboolOut(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_out(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TemporalAsHexwkb wraps MEOS C function temporal_as_hexwkb.
func TemporalAsHexwkb(temp *Temporal, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_as_hexwkb(temp._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TemporalAsMFJSON wraps MEOS C function temporal_as_mfjson.
func TemporalAsMFJSON(temp *Temporal, with_bbox bool, flags int, precision int, srs string) (_r0 string, _err error) {
	_c_srs := C.CString(srs)
	defer C.free(unsafe.Pointer(_c_srs))
	C.meos_errno_reset()
	_cret := C.temporal_as_mfjson(temp._inner, C.bool(with_bbox), C.int(flags), C.int(precision), _c_srs)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TemporalAsWKB wraps MEOS C function temporal_as_wkb.
func TemporalAsWKB(temp *Temporal, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_as_wkb(temp._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// WKBVariantFromEndian wraps MEOS C function wkb_variant_from_endian.
func WKBVariantFromEndian(endian string) (_r0 uint8, _err error) {
	_c_endian := C.CString(endian)
	defer C.free(unsafe.Pointer(_c_endian))
	C.meos_errno_reset()
	_cret := C.wkb_variant_from_endian(_c_endian)
	if _err = meosError(); _err != nil {
		return
	}
	return uint8(_cret), nil
}


// TemporalFromHexwkb wraps MEOS C function temporal_from_hexwkb.
func TemporalFromHexwkb(hexwkb string) (_r0 *Temporal, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.temporal_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalFromWKB wraps MEOS C function temporal_from_wkb.
func TemporalFromWKB(wkb unsafe.Pointer, size uint) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatFromMFJSON wraps MEOS C function tfloat_from_mfjson.
func TfloatFromMFJSON(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tfloat_from_mfjson(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatIn wraps MEOS C function tfloat_in.
func TfloatIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tfloat_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatOut wraps MEOS C function tfloat_out.
func TfloatOut(temp *Temporal, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_out(temp._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TintFromMFJSON wraps MEOS C function tint_from_mfjson.
func TintFromMFJSON(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tint_from_mfjson(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TbigintFromMFJSON wraps MEOS C function tbigint_from_mfjson.
func TbigintFromMFJSON(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tbigint_from_mfjson(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintIn wraps MEOS C function tint_in.
func TintIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tint_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TbigintIn wraps MEOS C function tbigint_in.
func TbigintIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tbigint_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintOut wraps MEOS C function tint_out.
func TintOut(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_out(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TbigintOut wraps MEOS C function tbigint_out.
func TbigintOut(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_out(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TtextFromMFJSON wraps MEOS C function ttext_from_mfjson.
func TtextFromMFJSON(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.ttext_from_mfjson(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtextIn wraps MEOS C function ttext_in.
func TtextIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.ttext_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtextOut wraps MEOS C function ttext_out.
func TtextOut(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_out(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TboolFromBaseTemp wraps MEOS C function tbool_from_base_temp.
func TboolFromBaseTemp(b bool, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_from_base_temp(C.bool(b), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TboolinstMake wraps MEOS C function tboolinst_make.
func TboolinstMake(b bool, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tboolinst_make(C.bool(b), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TboolseqFromBaseTstzset wraps MEOS C function tboolseq_from_base_tstzset.
func TboolseqFromBaseTstzset(b bool, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tboolseq_from_base_tstzset(C.bool(b), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TboolseqFromBaseTstzspan wraps MEOS C function tboolseq_from_base_tstzspan.
func TboolseqFromBaseTstzspan(b bool, s *Span) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tboolseq_from_base_tstzspan(C.bool(b), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TboolseqsetFromBaseTstzspanset wraps MEOS C function tboolseqset_from_base_tstzspanset.
func TboolseqsetFromBaseTstzspanset(b bool, ss *SpanSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tboolseqset_from_base_tstzspanset(C.bool(b), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TemporalCopy wraps MEOS C function temporal_copy.
func TemporalCopy(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_copy(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatFromBaseTemp wraps MEOS C function tfloat_from_base_temp.
func TfloatFromBaseTemp(d float64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_from_base_temp(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatinstMake wraps MEOS C function tfloatinst_make.
func TfloatinstMake(d float64, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloatinst_make(C.double(d), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TfloatseqFromBaseTstzset wraps MEOS C function tfloatseq_from_base_tstzset.
func TfloatseqFromBaseTstzset(d float64, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloatseq_from_base_tstzset(C.double(d), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TfloatseqFromBaseTstzspan wraps MEOS C function tfloatseq_from_base_tstzspan.
func TfloatseqFromBaseTstzspan(d float64, s *Span, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloatseq_from_base_tstzspan(C.double(d), s._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TfloatseqsetFromBaseTstzspanset wraps MEOS C function tfloatseqset_from_base_tstzspanset.
func TfloatseqsetFromBaseTstzspanset(d float64, ss *SpanSet, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloatseqset_from_base_tstzspanset(C.double(d), ss._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TintFromBaseTemp wraps MEOS C function tint_from_base_temp.
func TintFromBaseTemp(i int, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_from_base_temp(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TbigintFromBaseTemp wraps MEOS C function tbigint_from_base_temp.
func TbigintFromBaseTemp(i int64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_from_base_temp(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintinstMake wraps MEOS C function tintinst_make.
func TintinstMake(i int, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tintinst_make(C.int(i), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TbigintinstMake wraps MEOS C function tbigintinst_make.
func TbigintinstMake(i int64, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigintinst_make(C.int64_t(i), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TintseqFromBaseTstzset wraps MEOS C function tintseq_from_base_tstzset.
func TintseqFromBaseTstzset(i int, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tintseq_from_base_tstzset(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TbigintseqFromBaseTstzset wraps MEOS C function tbigintseq_from_base_tstzset.
func TbigintseqFromBaseTstzset(i int64, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigintseq_from_base_tstzset(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TintseqFromBaseTstzspan wraps MEOS C function tintseq_from_base_tstzspan.
func TintseqFromBaseTstzspan(i int, s *Span) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tintseq_from_base_tstzspan(C.int(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TbigintseqFromBaseTstzspan wraps MEOS C function tbigintseq_from_base_tstzspan.
func TbigintseqFromBaseTstzspan(i int64, s *Span) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigintseq_from_base_tstzspan(C.int64_t(i), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TintseqsetFromBaseTstzspanset wraps MEOS C function tintseqset_from_base_tstzspanset.
func TintseqsetFromBaseTstzspanset(i int, ss *SpanSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tintseqset_from_base_tstzspanset(C.int(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TbigintseqsetFromBaseTstzspanset wraps MEOS C function tbigintseqset_from_base_tstzspanset.
func TbigintseqsetFromBaseTstzspanset(i int64, ss *SpanSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigintseqset_from_base_tstzspanset(C.int64_t(i), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequenceMake wraps MEOS C function tsequence_make.
func TsequenceMake(instants unsafe.Pointer, count int, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_make((**C.TInstant)(unsafe.Pointer(instants)), C.int(count), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TsequencesetMake wraps MEOS C function tsequenceset_make.
func TsequencesetMake(sequences unsafe.Pointer, count int, normalize bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_make((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count), C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetMakeGaps wraps MEOS C function tsequenceset_make_gaps.
func TsequencesetMakeGaps(instants unsafe.Pointer, count int, interp Interpolation, maxt *Interval, maxdist float64) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_make_gaps((**C.TInstant)(unsafe.Pointer(instants)), C.int(count), C.interpType(interp), maxt._inner, C.double(maxdist))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TtextFromBaseTemp wraps MEOS C function ttext_from_base_temp.
func TtextFromBaseTemp(txt string, temp *Temporal) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ttext_from_base_temp(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtextinstMake wraps MEOS C function ttextinst_make.
func TtextinstMake(txt string, t int64) (_r0 *TInstant, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ttextinst_make(_c_txt, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TtextseqFromBaseTstzset wraps MEOS C function ttextseq_from_base_tstzset.
func TtextseqFromBaseTstzset(txt string, s *Set) (_r0 *TSequence, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ttextseq_from_base_tstzset(_c_txt, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TtextseqFromBaseTstzspan wraps MEOS C function ttextseq_from_base_tstzspan.
func TtextseqFromBaseTstzspan(txt string, s *Span) (_r0 *TSequence, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ttextseq_from_base_tstzspan(_c_txt, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TtextseqsetFromBaseTstzspanset wraps MEOS C function ttextseqset_from_base_tstzspanset.
func TtextseqsetFromBaseTstzspanset(txt string, ss *SpanSet) (_r0 *TSequenceSet, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ttextseqset_from_base_tstzspanset(_c_txt, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TboolToTint wraps MEOS C function tbool_to_tint.
func TboolToTint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_to_tint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalToTstzspan wraps MEOS C function temporal_to_tstzspan.
func TemporalToTstzspan(temp *Temporal) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_to_tstzspan(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TfloatToTint wraps MEOS C function tfloat_to_tint.
func TfloatToTint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_to_tint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatToTbigint wraps MEOS C function tfloat_to_tbigint.
func TfloatToTbigint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_to_tbigint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintToTfloat wraps MEOS C function tint_to_tfloat.
func TintToTfloat(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_to_tfloat(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintToTbigint wraps MEOS C function tint_to_tbigint.
func TintToTbigint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_to_tbigint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TbigintToTint wraps MEOS C function tbigint_to_tint.
func TbigintToTint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_to_tint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TbigintToTfloat wraps MEOS C function tbigint_to_tfloat.
func TbigintToTfloat(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_to_tfloat(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberToSpan wraps MEOS C function tnumber_to_span.
func TnumberToSpan(temp *Temporal) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_to_span(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TnumberToTBOX wraps MEOS C function tnumber_to_tbox.
func TnumberToTBOX(temp *Temporal) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_to_tbox(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TboolEndValue wraps MEOS C function tbool_end_value.
func TboolEndValue(temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TboolStartValue wraps MEOS C function tbool_start_value.
func TboolStartValue(temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TboolValueAtTimestamptz wraps MEOS C function tbool_value_at_timestamptz.
func TboolValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 bool, _err error) {
	var _out_value C.bool
	C.meos_errno_reset()
	_cret := C.tbool_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), bool(_out_value), nil
}


// TboolValueN wraps MEOS C function tbool_value_n.
func TboolValueN(temp *Temporal, n int) (_r0 bool, _r1 bool, _err error) {
	var _out_result C.bool
	C.meos_errno_reset()
	_cret := C.tbool_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), bool(_out_result), nil
}


// TboolValues wraps MEOS C function tbool_values.
func TboolValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TemporalDuration wraps MEOS C function temporal_duration.
func TemporalDuration(temp *Temporal, boundspan bool) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_duration(temp._inner, C.bool(boundspan))
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// TemporalEndInstant wraps MEOS C function temporal_end_instant.
func TemporalEndInstant(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_end_instant(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TemporalEndSequence wraps MEOS C function temporal_end_sequence.
func TemporalEndSequence(temp *Temporal) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_end_sequence(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TemporalEndTimestamptz wraps MEOS C function temporal_end_timestamptz.
func TemporalEndTimestamptz(temp *Temporal) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_end_timestamptz(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TemporalHash wraps MEOS C function temporal_hash.
func TemporalHash(temp *Temporal) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_hash(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TemporalHashExtended wraps MEOS C function temporal_hash_extended.
func TemporalHashExtended(temp *Temporal, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_hash_extended(temp._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TemporalInstantN wraps MEOS C function temporal_instant_n.
func TemporalInstantN(temp *Temporal, n int) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_instant_n(temp._inner, C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TemporalInstants wraps MEOS C function temporal_instants.
func TemporalInstants(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_instants(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TemporalInterp wraps MEOS C function temporal_interp.
func TemporalInterp(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_interp(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// TemporalLowerInc wraps MEOS C function temporal_lower_inc.
func TemporalLowerInc(temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_lower_inc(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemporalMaxInstant wraps MEOS C function temporal_max_instant.
func TemporalMaxInstant(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_max_instant(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TemporalMinInstant wraps MEOS C function temporal_min_instant.
func TemporalMinInstant(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_min_instant(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TemporalNumInstants wraps MEOS C function temporal_num_instants.
func TemporalNumInstants(temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_num_instants(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TemporalNumSequences wraps MEOS C function temporal_num_sequences.
func TemporalNumSequences(temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_num_sequences(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TemporalNumTimestamps wraps MEOS C function temporal_num_timestamps.
func TemporalNumTimestamps(temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_num_timestamps(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TemporalSegmDuration wraps MEOS C function temporal_segm_duration.
func TemporalSegmDuration(temp *Temporal, duration *Interval, atleast bool, strict bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_segm_duration(temp._inner, duration._inner, C.bool(atleast), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TemporalSegments wraps MEOS C function temporal_segments.
func TemporalSegments(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_segments(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TemporalSequenceN wraps MEOS C function temporal_sequence_n.
func TemporalSequenceN(temp *Temporal, i int) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_sequence_n(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TemporalSequences wraps MEOS C function temporal_sequences.
func TemporalSequences(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_sequences(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TemporalStartInstant wraps MEOS C function temporal_start_instant.
func TemporalStartInstant(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_start_instant(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TemporalStartSequence wraps MEOS C function temporal_start_sequence.
func TemporalStartSequence(temp *Temporal) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_start_sequence(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TemporalStartTimestamptz wraps MEOS C function temporal_start_timestamptz.
func TemporalStartTimestamptz(temp *Temporal) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_start_timestamptz(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TemporalStops wraps MEOS C function temporal_stops.
func TemporalStops(temp *Temporal, maxdist float64, minduration *Interval) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_stops(temp._inner, C.double(maxdist), minduration._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TemporalSubtype wraps MEOS C function temporal_subtype.
func TemporalSubtype(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_subtype(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// TemporalBasetypeName wraps MEOS C function temporal_basetype_name.
func TemporalBasetypeName(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_basetype_name(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// TemporalTime wraps MEOS C function temporal_time.
func TemporalTime(temp *Temporal) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_time(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TemporalTimestamps wraps MEOS C function temporal_timestamps.
func TemporalTimestamps(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_timestamps(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TemporalTimestamptzN wraps MEOS C function temporal_timestamptz_n.
func TemporalTimestamptzN(temp *Temporal, n int) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.TimestampTz
	C.meos_errno_reset()
	_cret := C.temporal_timestamptz_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// TemporalUpperInc wraps MEOS C function temporal_upper_inc.
func TemporalUpperInc(temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_upper_inc(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TfloatEndValue wraps MEOS C function tfloat_end_value.
func TfloatEndValue(temp *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TfloatMinValue wraps MEOS C function tfloat_min_value.
func TfloatMinValue(temp *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_min_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TfloatMaxValue wraps MEOS C function tfloat_max_value.
func TfloatMaxValue(temp *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_max_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TfloatStartValue wraps MEOS C function tfloat_start_value.
func TfloatStartValue(temp *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TfloatValueAtTimestamptz wraps MEOS C function tfloat_value_at_timestamptz.
func TfloatValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 float64, _err error) {
	var _out_value C.double
	C.meos_errno_reset()
	_cret := C.tfloat_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_value), nil
}


// TfloatValueN wraps MEOS C function tfloat_value_n.
func TfloatValueN(temp *Temporal, n int) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tfloat_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TfloatValues wraps MEOS C function tfloat_values.
func TfloatValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TintEndValue wraps MEOS C function tint_end_value.
func TintEndValue(temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TbigintEndValue wraps MEOS C function tbigint_end_value.
func TbigintEndValue(temp *Temporal) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TintMaxValue wraps MEOS C function tint_max_value.
func TintMaxValue(temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_max_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TbigintMaxValue wraps MEOS C function tbigint_max_value.
func TbigintMaxValue(temp *Temporal) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_max_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TintMinValue wraps MEOS C function tint_min_value.
func TintMinValue(temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_min_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TbigintMinValue wraps MEOS C function tbigint_min_value.
func TbigintMinValue(temp *Temporal) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_min_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TintStartValue wraps MEOS C function tint_start_value.
func TintStartValue(temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TbigintStartValue wraps MEOS C function tbigint_start_value.
func TbigintStartValue(temp *Temporal) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TbigintValueAtTimestamptz wraps MEOS C function tbigint_value_at_timestamptz.
func TbigintValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 int64, _err error) {
	var _out_value C.int64_t
	C.meos_errno_reset()
	_cret := C.tbigint_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_value), nil
}


// TintValueAtTimestamptz wraps MEOS C function tint_value_at_timestamptz.
func TintValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 int, _err error) {
	var _out_value C.int
	C.meos_errno_reset()
	_cret := C.tint_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int(_out_value), nil
}


// TintValueN wraps MEOS C function tint_value_n.
func TintValueN(temp *Temporal, n int) (_r0 bool, _r1 int, _err error) {
	var _out_result C.int
	C.meos_errno_reset()
	_cret := C.tint_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int(_out_result), nil
}


// TbigintValueN wraps MEOS C function tbigint_value_n.
func TbigintValueN(temp *Temporal, n int64) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.int64_t
	C.meos_errno_reset()
	_cret := C.tbigint_value_n(temp._inner, C.int64_t(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// TintValues wraps MEOS C function tint_values.
func TintValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TbigintValues wraps MEOS C function tbigint_values.
func TbigintValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TnumberAvgValue wraps MEOS C function tnumber_avg_value.
func TnumberAvgValue(temp *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_avg_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TnumberIntegral wraps MEOS C function tnumber_integral.
func TnumberIntegral(temp *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_integral(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TnumberTwavg wraps MEOS C function tnumber_twavg.
func TnumberTwavg(temp *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_twavg(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TnumberValuespans wraps MEOS C function tnumber_valuespans.
func TnumberValuespans(temp *Temporal) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_valuespans(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TtextEndValue wraps MEOS C function ttext_end_value.
func TtextEndValue(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TtextMaxValue wraps MEOS C function ttext_max_value.
func TtextMaxValue(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_max_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TtextMinValue wraps MEOS C function ttext_min_value.
func TtextMinValue(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_min_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TtextStartValue wraps MEOS C function ttext_start_value.
func TtextStartValue(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TtextValueAtTimestamptz wraps MEOS C function ttext_value_at_timestamptz.
func TtextValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 string, _err error) {
	var _out_value *C.text
	C.meos_errno_reset()
	_cret := C.ttext_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), C.GoString(C.text_to_cstring(_out_value)), nil
}


// TtextValueN wraps MEOS C function ttext_value_n.
func TtextValueN(temp *Temporal, n int) (_r0 bool, _r1 string, _err error) {
	var _out_result *C.text
	C.meos_errno_reset()
	_cret := C.ttext_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), C.GoString(C.text_to_cstring(_out_result)), nil
}


// TtextValues wraps MEOS C function ttext_values.
func TtextValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// FloatDegrees wraps MEOS C function float_degrees.
func FloatDegrees(value float64, normalize bool) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float_degrees(C.double(value), C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TemparrRound wraps MEOS C function temparr_round.
func TemparrRound(temp unsafe.Pointer, count int, maxdd int) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.temparr_round((**C.Temporal)(unsafe.Pointer(temp)), C.int(count), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TemporalRound wraps MEOS C function temporal_round.
func TemporalRound(temp *Temporal, maxdd int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_round(temp._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalScaleTime wraps MEOS C function temporal_scale_time.
func TemporalScaleTime(temp *Temporal, duration *Interval) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_scale_time(temp._inner, duration._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalSetInterp wraps MEOS C function temporal_set_interp.
func TemporalSetInterp(temp *Temporal, interp Interpolation) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_set_interp(temp._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalShiftScaleTime wraps MEOS C function temporal_shift_scale_time.
func TemporalShiftScaleTime(temp *Temporal, shift *Interval, duration *Interval) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_shift_scale_time(temp._inner, shift._inner, duration._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalShiftTime wraps MEOS C function temporal_shift_time.
func TemporalShiftTime(temp *Temporal, shift *Interval) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_shift_time(temp._inner, shift._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAsTinstant wraps MEOS C function temporal_as_tinstant.
func TemporalAsTinstant(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_as_tinstant(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TemporalAsTsequence wraps MEOS C function temporal_as_tsequence.
func TemporalAsTsequence(temp *Temporal, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_as_tsequence(temp._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TemporalAsTsequenceset wraps MEOS C function temporal_as_tsequenceset.
func TemporalAsTsequenceset(temp *Temporal, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_as_tsequenceset(temp._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TfloatCeil wraps MEOS C function tfloat_ceil.
func TfloatCeil(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_ceil(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatDegrees wraps MEOS C function tfloat_degrees.
func TfloatDegrees(temp *Temporal, normalize bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_degrees(temp._inner, C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatFloor wraps MEOS C function tfloat_floor.
func TfloatFloor(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_floor(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatRadians wraps MEOS C function tfloat_radians.
func TfloatRadians(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_radians(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatScaleValue wraps MEOS C function tfloat_scale_value.
func TfloatScaleValue(temp *Temporal, width float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_scale_value(temp._inner, C.double(width))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatShiftScaleValue wraps MEOS C function tfloat_shift_scale_value.
func TfloatShiftScaleValue(temp *Temporal, shift float64, width float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_shift_scale_value(temp._inner, C.double(shift), C.double(width))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatShiftValue wraps MEOS C function tfloat_shift_value.
func TfloatShiftValue(temp *Temporal, shift float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_shift_value(temp._inner, C.double(shift))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintScaleValue wraps MEOS C function tint_scale_value.
func TintScaleValue(temp *Temporal, width int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_scale_value(temp._inner, C.int(width))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TbigintScaleValue wraps MEOS C function tbigint_scale_value.
func TbigintScaleValue(temp *Temporal, width int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_scale_value(temp._inner, C.int64_t(width))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintShiftScaleValue wraps MEOS C function tint_shift_scale_value.
func TintShiftScaleValue(temp *Temporal, shift int, width int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_shift_scale_value(temp._inner, C.int(shift), C.int(width))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TbigintShiftScaleValue wraps MEOS C function tbigint_shift_scale_value.
func TbigintShiftScaleValue(temp *Temporal, shift int64, width int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_shift_scale_value(temp._inner, C.int64_t(shift), C.int64_t(width))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintShiftValue wraps MEOS C function tint_shift_value.
func TintShiftValue(temp *Temporal, shift int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_shift_value(temp._inner, C.int(shift))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TbigintShiftValue wraps MEOS C function tbigint_shift_value.
func TbigintShiftValue(temp *Temporal, shift int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_shift_value(temp._inner, C.int64_t(shift))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAppendTinstant wraps MEOS C function temporal_append_tinstant.
func TemporalAppendTinstant(temp *Temporal, inst *TInstant, interp Interpolation, maxdist float64, maxt *Interval, expand bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_append_tinstant(temp._inner, inst._inner, C.interpType(interp), C.double(maxdist), maxt._inner, C.bool(expand))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAppendTsequence wraps MEOS C function temporal_append_tsequence.
func TemporalAppendTsequence(temp *Temporal, seq *TSequence, expand bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_append_tsequence(temp._inner, seq._inner, C.bool(expand))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalDeleteTimestamptz wraps MEOS C function temporal_delete_timestamptz.
func TemporalDeleteTimestamptz(temp *Temporal, t int64, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_delete_timestamptz(temp._inner, C.TimestampTz(t), C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalDeleteTstzset wraps MEOS C function temporal_delete_tstzset.
func TemporalDeleteTstzset(temp *Temporal, s *Set, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_delete_tstzset(temp._inner, s._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalDeleteTstzspan wraps MEOS C function temporal_delete_tstzspan.
func TemporalDeleteTstzspan(temp *Temporal, s *Span, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_delete_tstzspan(temp._inner, s._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalDeleteTstzspanset wraps MEOS C function temporal_delete_tstzspanset.
func TemporalDeleteTstzspanset(temp *Temporal, ss *SpanSet, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_delete_tstzspanset(temp._inner, ss._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalInsert wraps MEOS C function temporal_insert.
func TemporalInsert(temp1 *Temporal, temp2 *Temporal, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_insert(temp1._inner, temp2._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalMerge wraps MEOS C function temporal_merge.
func TemporalMerge(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_merge(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalMergeArray wraps MEOS C function temporal_merge_array.
func TemporalMergeArray(temparr unsafe.Pointer, count int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_merge_array((**C.Temporal)(unsafe.Pointer(temparr)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalUpdate wraps MEOS C function temporal_update.
func TemporalUpdate(temp1 *Temporal, temp2 *Temporal, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_update(temp1._inner, temp2._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TboolAtValue wraps MEOS C function tbool_at_value.
func TboolAtValue(temp *Temporal, b bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_at_value(temp._inner, C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TboolMinusValue wraps MEOS C function tbool_minus_value.
func TboolMinusValue(temp *Temporal, b bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_minus_value(temp._inner, C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAfterTimestamptz wraps MEOS C function temporal_after_timestamptz.
func TemporalAfterTimestamptz(temp *Temporal, t int64, strict bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_after_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAtMax wraps MEOS C function temporal_at_max.
func TemporalAtMax(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_at_max(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAtMin wraps MEOS C function temporal_at_min.
func TemporalAtMin(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_at_min(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAtTimestamptz wraps MEOS C function temporal_at_timestamptz.
func TemporalAtTimestamptz(temp *Temporal, t int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_at_timestamptz(temp._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAtTstzset wraps MEOS C function temporal_at_tstzset.
func TemporalAtTstzset(temp *Temporal, s *Set) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_at_tstzset(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAtTstzspan wraps MEOS C function temporal_at_tstzspan.
func TemporalAtTstzspan(temp *Temporal, s *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_at_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAtTstzspanset wraps MEOS C function temporal_at_tstzspanset.
func TemporalAtTstzspanset(temp *Temporal, ss *SpanSet) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_at_tstzspanset(temp._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAtValues wraps MEOS C function temporal_at_values.
func TemporalAtValues(temp *Temporal, set *Set) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_at_values(temp._inner, set._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalBeforeTimestamptz wraps MEOS C function temporal_before_timestamptz.
func TemporalBeforeTimestamptz(temp *Temporal, t int64, strict bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_before_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalMinusMax wraps MEOS C function temporal_minus_max.
func TemporalMinusMax(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_minus_max(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalMinusMin wraps MEOS C function temporal_minus_min.
func TemporalMinusMin(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_minus_min(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalMinusTimestamptz wraps MEOS C function temporal_minus_timestamptz.
func TemporalMinusTimestamptz(temp *Temporal, t int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_minus_timestamptz(temp._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalMinusTstzset wraps MEOS C function temporal_minus_tstzset.
func TemporalMinusTstzset(temp *Temporal, s *Set) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_minus_tstzset(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalMinusTstzspan wraps MEOS C function temporal_minus_tstzspan.
func TemporalMinusTstzspan(temp *Temporal, s *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_minus_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalMinusTstzspanset wraps MEOS C function temporal_minus_tstzspanset.
func TemporalMinusTstzspanset(temp *Temporal, ss *SpanSet) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_minus_tstzspanset(temp._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalMinusValues wraps MEOS C function temporal_minus_values.
func TemporalMinusValues(temp *Temporal, set *Set) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_minus_values(temp._inner, set._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatAtValue wraps MEOS C function tfloat_at_value.
func TfloatAtValue(temp *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_at_value(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatMinusValue wraps MEOS C function tfloat_minus_value.
func TfloatMinusValue(temp *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_minus_value(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintAtValue wraps MEOS C function tint_at_value.
func TintAtValue(temp *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_at_value(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TintMinusValue wraps MEOS C function tint_minus_value.
func TintMinusValue(temp *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_minus_value(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberAtSpan wraps MEOS C function tnumber_at_span.
func TnumberAtSpan(temp *Temporal, span *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_at_span(temp._inner, span._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberAtSpanset wraps MEOS C function tnumber_at_spanset.
func TnumberAtSpanset(temp *Temporal, ss *SpanSet) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_at_spanset(temp._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberAtTBOX wraps MEOS C function tnumber_at_tbox.
func TnumberAtTBOX(temp *Temporal, box *TBox) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_at_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberMinusSpan wraps MEOS C function tnumber_minus_span.
func TnumberMinusSpan(temp *Temporal, span *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_minus_span(temp._inner, span._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberMinusSpanset wraps MEOS C function tnumber_minus_spanset.
func TnumberMinusSpanset(temp *Temporal, ss *SpanSet) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_minus_spanset(temp._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberMinusTBOX wraps MEOS C function tnumber_minus_tbox.
func TnumberMinusTBOX(temp *Temporal, box *TBox) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_minus_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtextAtValue wraps MEOS C function ttext_at_value.
func TtextAtValue(temp *Temporal, txt string) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ttext_at_value(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtextMinusValue wraps MEOS C function ttext_minus_value.
func TtextMinusValue(temp *Temporal, txt string) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ttext_minus_value(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalCmp wraps MEOS C function temporal_cmp.
func TemporalCmp(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_cmp(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TemporalEq wraps MEOS C function temporal_eq.
func TemporalEq(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_eq(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemporalGe wraps MEOS C function temporal_ge.
func TemporalGe(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_ge(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemporalGt wraps MEOS C function temporal_gt.
func TemporalGt(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_gt(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemporalLe wraps MEOS C function temporal_le.
func TemporalLe(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_le(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemporalLt wraps MEOS C function temporal_lt.
func TemporalLt(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_lt(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemporalNe wraps MEOS C function temporal_ne.
func TemporalNe(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_ne(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AlwaysEqBoolTbool wraps MEOS C function always_eq_bool_tbool.
func AlwaysEqBoolTbool(b bool, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_bool_tbool(C.bool(b), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqFloatTfloat wraps MEOS C function always_eq_float_tfloat.
func AlwaysEqFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqIntTint wraps MEOS C function always_eq_int_tint.
func AlwaysEqIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTboolBool wraps MEOS C function always_eq_tbool_bool.
func AlwaysEqTboolBool(temp *Temporal, b bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tbool_bool(temp._inner, C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTemporalTemporal wraps MEOS C function always_eq_temporal_temporal.
func AlwaysEqTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTextTtext wraps MEOS C function always_eq_text_ttext.
func AlwaysEqTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_eq_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTfloatFloat wraps MEOS C function always_eq_tfloat_float.
func AlwaysEqTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTintInt wraps MEOS C function always_eq_tint_int.
func AlwaysEqTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqBigintTbigint wraps MEOS C function always_eq_bigint_tbigint.
func AlwaysEqBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTbigintBigint wraps MEOS C function always_eq_tbigint_bigint.
func AlwaysEqTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTtextText wraps MEOS C function always_eq_ttext_text.
func AlwaysEqTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_eq_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGeFloatTfloat wraps MEOS C function always_ge_float_tfloat.
func AlwaysGeFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ge_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGeIntTint wraps MEOS C function always_ge_int_tint.
func AlwaysGeIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ge_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGeTemporalTemporal wraps MEOS C function always_ge_temporal_temporal.
func AlwaysGeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ge_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGeTextTtext wraps MEOS C function always_ge_text_ttext.
func AlwaysGeTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_ge_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGeTfloatFloat wraps MEOS C function always_ge_tfloat_float.
func AlwaysGeTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ge_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGeTintInt wraps MEOS C function always_ge_tint_int.
func AlwaysGeTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ge_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGeBigintTbigint wraps MEOS C function always_ge_bigint_tbigint.
func AlwaysGeBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ge_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGeTbigintBigint wraps MEOS C function always_ge_tbigint_bigint.
func AlwaysGeTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ge_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGeTtextText wraps MEOS C function always_ge_ttext_text.
func AlwaysGeTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_ge_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGtFloatTfloat wraps MEOS C function always_gt_float_tfloat.
func AlwaysGtFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_gt_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGtIntTint wraps MEOS C function always_gt_int_tint.
func AlwaysGtIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_gt_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGtTemporalTemporal wraps MEOS C function always_gt_temporal_temporal.
func AlwaysGtTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_gt_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGtTextTtext wraps MEOS C function always_gt_text_ttext.
func AlwaysGtTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_gt_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGtTfloatFloat wraps MEOS C function always_gt_tfloat_float.
func AlwaysGtTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_gt_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGtTintInt wraps MEOS C function always_gt_tint_int.
func AlwaysGtTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_gt_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGtBigintTbigint wraps MEOS C function always_gt_bigint_tbigint.
func AlwaysGtBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_gt_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGtTbigintBigint wraps MEOS C function always_gt_tbigint_bigint.
func AlwaysGtTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_gt_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysGtTtextText wraps MEOS C function always_gt_ttext_text.
func AlwaysGtTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_gt_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLeFloatTfloat wraps MEOS C function always_le_float_tfloat.
func AlwaysLeFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_le_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLeIntTint wraps MEOS C function always_le_int_tint.
func AlwaysLeIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_le_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLeTemporalTemporal wraps MEOS C function always_le_temporal_temporal.
func AlwaysLeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_le_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLeTextTtext wraps MEOS C function always_le_text_ttext.
func AlwaysLeTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_le_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLeTfloatFloat wraps MEOS C function always_le_tfloat_float.
func AlwaysLeTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_le_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLeTintInt wraps MEOS C function always_le_tint_int.
func AlwaysLeTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_le_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLeBigintTbigint wraps MEOS C function always_le_bigint_tbigint.
func AlwaysLeBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_le_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLeTbigintBigint wraps MEOS C function always_le_tbigint_bigint.
func AlwaysLeTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_le_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLeTtextText wraps MEOS C function always_le_ttext_text.
func AlwaysLeTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_le_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLtFloatTfloat wraps MEOS C function always_lt_float_tfloat.
func AlwaysLtFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_lt_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLtIntTint wraps MEOS C function always_lt_int_tint.
func AlwaysLtIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_lt_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLtTemporalTemporal wraps MEOS C function always_lt_temporal_temporal.
func AlwaysLtTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_lt_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLtTextTtext wraps MEOS C function always_lt_text_ttext.
func AlwaysLtTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_lt_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLtTfloatFloat wraps MEOS C function always_lt_tfloat_float.
func AlwaysLtTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_lt_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLtTintInt wraps MEOS C function always_lt_tint_int.
func AlwaysLtTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_lt_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLtBigintTbigint wraps MEOS C function always_lt_bigint_tbigint.
func AlwaysLtBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_lt_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLtTbigintBigint wraps MEOS C function always_lt_tbigint_bigint.
func AlwaysLtTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_lt_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysLtTtextText wraps MEOS C function always_lt_ttext_text.
func AlwaysLtTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_lt_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeBoolTbool wraps MEOS C function always_ne_bool_tbool.
func AlwaysNeBoolTbool(b bool, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_bool_tbool(C.bool(b), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeFloatTfloat wraps MEOS C function always_ne_float_tfloat.
func AlwaysNeFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeIntTint wraps MEOS C function always_ne_int_tint.
func AlwaysNeIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTboolBool wraps MEOS C function always_ne_tbool_bool.
func AlwaysNeTboolBool(temp *Temporal, b bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tbool_bool(temp._inner, C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTemporalTemporal wraps MEOS C function always_ne_temporal_temporal.
func AlwaysNeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTextTtext wraps MEOS C function always_ne_text_ttext.
func AlwaysNeTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_ne_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTfloatFloat wraps MEOS C function always_ne_tfloat_float.
func AlwaysNeTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTintInt wraps MEOS C function always_ne_tint_int.
func AlwaysNeTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeBigintTbigint wraps MEOS C function always_ne_bigint_tbigint.
func AlwaysNeBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTbigintBigint wraps MEOS C function always_ne_tbigint_bigint.
func AlwaysNeTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTtextText wraps MEOS C function always_ne_ttext_text.
func AlwaysNeTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.always_ne_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqBoolTbool wraps MEOS C function ever_eq_bool_tbool.
func EverEqBoolTbool(b bool, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_bool_tbool(C.bool(b), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqFloatTfloat wraps MEOS C function ever_eq_float_tfloat.
func EverEqFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqIntTint wraps MEOS C function ever_eq_int_tint.
func EverEqIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTboolBool wraps MEOS C function ever_eq_tbool_bool.
func EverEqTboolBool(temp *Temporal, b bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tbool_bool(temp._inner, C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTemporalTemporal wraps MEOS C function ever_eq_temporal_temporal.
func EverEqTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTextTtext wraps MEOS C function ever_eq_text_ttext.
func EverEqTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_eq_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTfloatFloat wraps MEOS C function ever_eq_tfloat_float.
func EverEqTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTintInt wraps MEOS C function ever_eq_tint_int.
func EverEqTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqBigintTbigint wraps MEOS C function ever_eq_bigint_tbigint.
func EverEqBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTbigintBigint wraps MEOS C function ever_eq_tbigint_bigint.
func EverEqTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTtextText wraps MEOS C function ever_eq_ttext_text.
func EverEqTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_eq_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGeFloatTfloat wraps MEOS C function ever_ge_float_tfloat.
func EverGeFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ge_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGeIntTint wraps MEOS C function ever_ge_int_tint.
func EverGeIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ge_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGeTemporalTemporal wraps MEOS C function ever_ge_temporal_temporal.
func EverGeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ge_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGeTextTtext wraps MEOS C function ever_ge_text_ttext.
func EverGeTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_ge_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGeTfloatFloat wraps MEOS C function ever_ge_tfloat_float.
func EverGeTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ge_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGeTintInt wraps MEOS C function ever_ge_tint_int.
func EverGeTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ge_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGeBigintTbigint wraps MEOS C function ever_ge_bigint_tbigint.
func EverGeBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ge_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGeTbigintBigint wraps MEOS C function ever_ge_tbigint_bigint.
func EverGeTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ge_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGeTtextText wraps MEOS C function ever_ge_ttext_text.
func EverGeTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_ge_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGtFloatTfloat wraps MEOS C function ever_gt_float_tfloat.
func EverGtFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_gt_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGtIntTint wraps MEOS C function ever_gt_int_tint.
func EverGtIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_gt_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGtTemporalTemporal wraps MEOS C function ever_gt_temporal_temporal.
func EverGtTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_gt_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGtTextTtext wraps MEOS C function ever_gt_text_ttext.
func EverGtTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_gt_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGtTfloatFloat wraps MEOS C function ever_gt_tfloat_float.
func EverGtTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_gt_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGtTintInt wraps MEOS C function ever_gt_tint_int.
func EverGtTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_gt_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGtBigintTbigint wraps MEOS C function ever_gt_bigint_tbigint.
func EverGtBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_gt_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGtTbigintBigint wraps MEOS C function ever_gt_tbigint_bigint.
func EverGtTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_gt_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverGtTtextText wraps MEOS C function ever_gt_ttext_text.
func EverGtTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_gt_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLeFloatTfloat wraps MEOS C function ever_le_float_tfloat.
func EverLeFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_le_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLeIntTint wraps MEOS C function ever_le_int_tint.
func EverLeIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_le_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLeTemporalTemporal wraps MEOS C function ever_le_temporal_temporal.
func EverLeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_le_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLeTextTtext wraps MEOS C function ever_le_text_ttext.
func EverLeTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_le_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLeTfloatFloat wraps MEOS C function ever_le_tfloat_float.
func EverLeTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_le_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLeTintInt wraps MEOS C function ever_le_tint_int.
func EverLeTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_le_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLeBigintTbigint wraps MEOS C function ever_le_bigint_tbigint.
func EverLeBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_le_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLeTbigintBigint wraps MEOS C function ever_le_tbigint_bigint.
func EverLeTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_le_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLeTtextText wraps MEOS C function ever_le_ttext_text.
func EverLeTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_le_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLtFloatTfloat wraps MEOS C function ever_lt_float_tfloat.
func EverLtFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_lt_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLtIntTint wraps MEOS C function ever_lt_int_tint.
func EverLtIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_lt_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLtTemporalTemporal wraps MEOS C function ever_lt_temporal_temporal.
func EverLtTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_lt_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLtTextTtext wraps MEOS C function ever_lt_text_ttext.
func EverLtTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_lt_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLtTfloatFloat wraps MEOS C function ever_lt_tfloat_float.
func EverLtTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_lt_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLtTintInt wraps MEOS C function ever_lt_tint_int.
func EverLtTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_lt_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLtBigintTbigint wraps MEOS C function ever_lt_bigint_tbigint.
func EverLtBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_lt_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLtTbigintBigint wraps MEOS C function ever_lt_tbigint_bigint.
func EverLtTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_lt_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverLtTtextText wraps MEOS C function ever_lt_ttext_text.
func EverLtTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_lt_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeBoolTbool wraps MEOS C function ever_ne_bool_tbool.
func EverNeBoolTbool(b bool, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_bool_tbool(C.bool(b), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeFloatTfloat wraps MEOS C function ever_ne_float_tfloat.
func EverNeFloatTfloat(d float64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeIntTint wraps MEOS C function ever_ne_int_tint.
func EverNeIntTint(i int, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTboolBool wraps MEOS C function ever_ne_tbool_bool.
func EverNeTboolBool(temp *Temporal, b bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tbool_bool(temp._inner, C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTemporalTemporal wraps MEOS C function ever_ne_temporal_temporal.
func EverNeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTextTtext wraps MEOS C function ever_ne_text_ttext.
func EverNeTextTtext(txt string, temp *Temporal) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_ne_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTfloatFloat wraps MEOS C function ever_ne_tfloat_float.
func EverNeTfloatFloat(temp *Temporal, d float64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTintInt wraps MEOS C function ever_ne_tint_int.
func EverNeTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeBigintTbigint wraps MEOS C function ever_ne_bigint_tbigint.
func EverNeBigintTbigint(i int64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTbigintBigint wraps MEOS C function ever_ne_tbigint_bigint.
func EverNeTbigintBigint(temp *Temporal, i int64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTtextText wraps MEOS C function ever_ne_ttext_text.
func EverNeTtextText(temp *Temporal, txt string) (_r0 int, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.ever_ne_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqBigintTbigint wraps MEOS C function teq_bigint_tbigint.
func TeqBigintTbigint(i int64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqBoolTbool wraps MEOS C function teq_bool_tbool.
func TeqBoolTbool(b bool, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_bool_tbool(C.bool(b), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqFloatTfloat wraps MEOS C function teq_float_tfloat.
func TeqFloatTfloat(d float64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqIntTint wraps MEOS C function teq_int_tint.
func TeqIntTint(i int, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTbigintBigint wraps MEOS C function teq_tbigint_bigint.
func TeqTbigintBigint(temp *Temporal, i int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTboolBool wraps MEOS C function teq_tbool_bool.
func TeqTboolBool(temp *Temporal, b bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tbool_bool(temp._inner, C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTemporalTemporal wraps MEOS C function teq_temporal_temporal.
func TeqTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTextTtext wraps MEOS C function teq_text_ttext.
func TeqTextTtext(txt string, temp *Temporal) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.teq_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTfloatFloat wraps MEOS C function teq_tfloat_float.
func TeqTfloatFloat(temp *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTintInt wraps MEOS C function teq_tint_int.
func TeqTintInt(temp *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTtextText wraps MEOS C function teq_ttext_text.
func TeqTtextText(temp *Temporal, txt string) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.teq_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeBigintTbigint wraps MEOS C function tge_bigint_tbigint.
func TgeBigintTbigint(i int64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tge_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeFloatTfloat wraps MEOS C function tge_float_tfloat.
func TgeFloatTfloat(d float64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tge_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeIntTint wraps MEOS C function tge_int_tint.
func TgeIntTint(i int, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tge_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeTemporalTemporal wraps MEOS C function tge_temporal_temporal.
func TgeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tge_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeTextTtext wraps MEOS C function tge_text_ttext.
func TgeTextTtext(txt string, temp *Temporal) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.tge_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeTbigintBigint wraps MEOS C function tge_tbigint_bigint.
func TgeTbigintBigint(temp *Temporal, i int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tge_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeTfloatFloat wraps MEOS C function tge_tfloat_float.
func TgeTfloatFloat(temp *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tge_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeTintInt wraps MEOS C function tge_tint_int.
func TgeTintInt(temp *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tge_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeTtextText wraps MEOS C function tge_ttext_text.
func TgeTtextText(temp *Temporal, txt string) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.tge_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgtBigintTbigint wraps MEOS C function tgt_bigint_tbigint.
func TgtBigintTbigint(i int64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgt_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgtFloatTfloat wraps MEOS C function tgt_float_tfloat.
func TgtFloatTfloat(d float64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgt_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgtIntTint wraps MEOS C function tgt_int_tint.
func TgtIntTint(i int, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgt_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgtTemporalTemporal wraps MEOS C function tgt_temporal_temporal.
func TgtTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgt_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgtTextTtext wraps MEOS C function tgt_text_ttext.
func TgtTextTtext(txt string, temp *Temporal) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.tgt_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgtTbigintBigint wraps MEOS C function tgt_tbigint_bigint.
func TgtTbigintBigint(temp *Temporal, i int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgt_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgtTfloatFloat wraps MEOS C function tgt_tfloat_float.
func TgtTfloatFloat(temp *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgt_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgtTintInt wraps MEOS C function tgt_tint_int.
func TgtTintInt(temp *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgt_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgtTtextText wraps MEOS C function tgt_ttext_text.
func TgtTtextText(temp *Temporal, txt string) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.tgt_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TleBigintTbigint wraps MEOS C function tle_bigint_tbigint.
func TleBigintTbigint(i int64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tle_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TleFloatTfloat wraps MEOS C function tle_float_tfloat.
func TleFloatTfloat(d float64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tle_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TleIntTint wraps MEOS C function tle_int_tint.
func TleIntTint(i int, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tle_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TleTemporalTemporal wraps MEOS C function tle_temporal_temporal.
func TleTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tle_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TleTextTtext wraps MEOS C function tle_text_ttext.
func TleTextTtext(txt string, temp *Temporal) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.tle_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TleTbigintBigint wraps MEOS C function tle_tbigint_bigint.
func TleTbigintBigint(temp *Temporal, i int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tle_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TleTfloatFloat wraps MEOS C function tle_tfloat_float.
func TleTfloatFloat(temp *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tle_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TleTintInt wraps MEOS C function tle_tint_int.
func TleTintInt(temp *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tle_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TleTtextText wraps MEOS C function tle_ttext_text.
func TleTtextText(temp *Temporal, txt string) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.tle_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TltBigintTbigint wraps MEOS C function tlt_bigint_tbigint.
func TltBigintTbigint(i int64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tlt_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TltFloatTfloat wraps MEOS C function tlt_float_tfloat.
func TltFloatTfloat(d float64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tlt_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TltIntTint wraps MEOS C function tlt_int_tint.
func TltIntTint(i int, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tlt_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TltTemporalTemporal wraps MEOS C function tlt_temporal_temporal.
func TltTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tlt_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TltTextTtext wraps MEOS C function tlt_text_ttext.
func TltTextTtext(txt string, temp *Temporal) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.tlt_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TltTbigintBigint wraps MEOS C function tlt_tbigint_bigint.
func TltTbigintBigint(temp *Temporal, i int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tlt_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TltTfloatFloat wraps MEOS C function tlt_tfloat_float.
func TltTfloatFloat(temp *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tlt_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TltTintInt wraps MEOS C function tlt_tint_int.
func TltTintInt(temp *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tlt_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TltTtextText wraps MEOS C function tlt_ttext_text.
func TltTtextText(temp *Temporal, txt string) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.tlt_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneBigintTbigint wraps MEOS C function tne_bigint_tbigint.
func TneBigintTbigint(i int64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_bigint_tbigint(C.int64_t(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneBoolTbool wraps MEOS C function tne_bool_tbool.
func TneBoolTbool(b bool, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_bool_tbool(C.bool(b), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneFloatTfloat wraps MEOS C function tne_float_tfloat.
func TneFloatTfloat(d float64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_float_tfloat(C.double(d), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneIntTint wraps MEOS C function tne_int_tint.
func TneIntTint(i int, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_int_tint(C.int(i), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTbigintBigint wraps MEOS C function tne_tbigint_bigint.
func TneTbigintBigint(temp *Temporal, i int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTboolBool wraps MEOS C function tne_tbool_bool.
func TneTboolBool(temp *Temporal, b bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tbool_bool(temp._inner, C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTemporalTemporal wraps MEOS C function tne_temporal_temporal.
func TneTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTextTtext wraps MEOS C function tne_text_ttext.
func TneTextTtext(txt string, temp *Temporal) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.tne_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTfloatFloat wraps MEOS C function tne_tfloat_float.
func TneTfloatFloat(temp *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTintInt wraps MEOS C function tne_tint_int.
func TneTintInt(temp *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTtextText wraps MEOS C function tne_ttext_text.
func TneTtextText(temp *Temporal, txt string) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.tne_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalSpans wraps MEOS C function temporal_spans.
func TemporalSpans(temp *Temporal, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_spans(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TemporalSplitEachNSpans wraps MEOS C function temporal_split_each_n_spans.
func TemporalSplitEachNSpans(temp *Temporal, elem_count int, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_split_each_n_spans(temp._inner, C.int(elem_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TemporalSplitNSpans wraps MEOS C function temporal_split_n_spans.
func TemporalSplitNSpans(temp *Temporal, span_count int, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_split_n_spans(temp._inner, C.int(span_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TnumberSplitEachNTboxes wraps MEOS C function tnumber_split_each_n_tboxes.
func TnumberSplitEachNTboxes(temp *Temporal, elem_count int, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_split_each_n_tboxes(temp._inner, C.int(elem_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TnumberSplitNTboxes wraps MEOS C function tnumber_split_n_tboxes.
func TnumberSplitNTboxes(temp *Temporal, box_count int, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_split_n_tboxes(temp._inner, C.int(box_count), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TnumberTboxes wraps MEOS C function tnumber_tboxes.
func TnumberTboxes(temp *Temporal, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_tboxes(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// AdjacentNumspanTnumber wraps MEOS C function adjacent_numspan_tnumber.
func AdjacentNumspanTnumber(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_numspan_tnumber(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTBOXTnumber wraps MEOS C function adjacent_tbox_tnumber.
func AdjacentTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTemporalTemporal wraps MEOS C function adjacent_temporal_temporal.
func AdjacentTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTemporalTstzspan wraps MEOS C function adjacent_temporal_tstzspan.
func AdjacentTemporalTstzspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_temporal_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTnumberNumspan wraps MEOS C function adjacent_tnumber_numspan.
func AdjacentTnumberNumspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tnumber_numspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTnumberTBOX wraps MEOS C function adjacent_tnumber_tbox.
func AdjacentTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTnumberTnumber wraps MEOS C function adjacent_tnumber_tnumber.
func AdjacentTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTstzspanTemporal wraps MEOS C function adjacent_tstzspan_temporal.
func AdjacentTstzspanTemporal(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tstzspan_temporal(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedNumspanTnumber wraps MEOS C function contained_numspan_tnumber.
func ContainedNumspanTnumber(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_numspan_tnumber(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTBOXTnumber wraps MEOS C function contained_tbox_tnumber.
func ContainedTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTemporalTemporal wraps MEOS C function contained_temporal_temporal.
func ContainedTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTemporalTstzspan wraps MEOS C function contained_temporal_tstzspan.
func ContainedTemporalTstzspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_temporal_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTnumberNumspan wraps MEOS C function contained_tnumber_numspan.
func ContainedTnumberNumspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tnumber_numspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTnumberTBOX wraps MEOS C function contained_tnumber_tbox.
func ContainedTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTnumberTnumber wraps MEOS C function contained_tnumber_tnumber.
func ContainedTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTstzspanTemporal wraps MEOS C function contained_tstzspan_temporal.
func ContainedTstzspanTemporal(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tstzspan_temporal(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsNumspanTnumber wraps MEOS C function contains_numspan_tnumber.
func ContainsNumspanTnumber(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_numspan_tnumber(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTBOXTnumber wraps MEOS C function contains_tbox_tnumber.
func ContainsTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTemporalTstzspan wraps MEOS C function contains_temporal_tstzspan.
func ContainsTemporalTstzspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_temporal_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTemporalTemporal wraps MEOS C function contains_temporal_temporal.
func ContainsTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTnumberNumspan wraps MEOS C function contains_tnumber_numspan.
func ContainsTnumberNumspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tnumber_numspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTnumberTBOX wraps MEOS C function contains_tnumber_tbox.
func ContainsTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTnumberTnumber wraps MEOS C function contains_tnumber_tnumber.
func ContainsTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTstzspanTemporal wraps MEOS C function contains_tstzspan_temporal.
func ContainsTstzspanTemporal(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tstzspan_temporal(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsNumspanTnumber wraps MEOS C function overlaps_numspan_tnumber.
func OverlapsNumspanTnumber(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_numspan_tnumber(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTBOXTnumber wraps MEOS C function overlaps_tbox_tnumber.
func OverlapsTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTemporalTemporal wraps MEOS C function overlaps_temporal_temporal.
func OverlapsTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemporalTimeOverlaps wraps MEOS C function temporal_time_overlaps.
func TemporalTimeOverlaps(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_time_overlaps(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTemporalTstzspan wraps MEOS C function overlaps_temporal_tstzspan.
func OverlapsTemporalTstzspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_temporal_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTnumberNumspan wraps MEOS C function overlaps_tnumber_numspan.
func OverlapsTnumberNumspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tnumber_numspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTnumberTBOX wraps MEOS C function overlaps_tnumber_tbox.
func OverlapsTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTnumberTnumber wraps MEOS C function overlaps_tnumber_tnumber.
func OverlapsTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTstzspanTemporal wraps MEOS C function overlaps_tstzspan_temporal.
func OverlapsTstzspanTemporal(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tstzspan_temporal(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameNumspanTnumber wraps MEOS C function same_numspan_tnumber.
func SameNumspanTnumber(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_numspan_tnumber(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTBOXTnumber wraps MEOS C function same_tbox_tnumber.
func SameTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTemporalTemporal wraps MEOS C function same_temporal_temporal.
func SameTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTemporalTstzspan wraps MEOS C function same_temporal_tstzspan.
func SameTemporalTstzspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_temporal_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTnumberNumspan wraps MEOS C function same_tnumber_numspan.
func SameTnumberNumspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tnumber_numspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTnumberTBOX wraps MEOS C function same_tnumber_tbox.
func SameTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTnumberTnumber wraps MEOS C function same_tnumber_tnumber.
func SameTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTstzspanTemporal wraps MEOS C function same_tstzspan_temporal.
func SameTstzspanTemporal(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tstzspan_temporal(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTBOXTnumber wraps MEOS C function after_tbox_tnumber.
func AfterTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTemporalTstzspan wraps MEOS C function after_temporal_tstzspan.
func AfterTemporalTstzspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_temporal_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTemporalTemporal wraps MEOS C function after_temporal_temporal.
func AfterTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTnumberTBOX wraps MEOS C function after_tnumber_tbox.
func AfterTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTnumberTnumber wraps MEOS C function after_tnumber_tnumber.
func AfterTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTstzspanTemporal wraps MEOS C function after_tstzspan_temporal.
func AfterTstzspanTemporal(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_tstzspan_temporal(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTBOXTnumber wraps MEOS C function before_tbox_tnumber.
func BeforeTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTemporalTstzspan wraps MEOS C function before_temporal_tstzspan.
func BeforeTemporalTstzspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_temporal_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTemporalTemporal wraps MEOS C function before_temporal_temporal.
func BeforeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTnumberTBOX wraps MEOS C function before_tnumber_tbox.
func BeforeTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTnumberTnumber wraps MEOS C function before_tnumber_tnumber.
func BeforeTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTstzspanTemporal wraps MEOS C function before_tstzspan_temporal.
func BeforeTstzspanTemporal(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_tstzspan_temporal(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTBOXTnumber wraps MEOS C function left_tbox_tnumber.
func LeftTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftNumspanTnumber wraps MEOS C function left_numspan_tnumber.
func LeftNumspanTnumber(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_numspan_tnumber(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTnumberNumspan wraps MEOS C function left_tnumber_numspan.
func LeftTnumberNumspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_tnumber_numspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTnumberTBOX wraps MEOS C function left_tnumber_tbox.
func LeftTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTnumberTnumber wraps MEOS C function left_tnumber_tnumber.
func LeftTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTBOXTnumber wraps MEOS C function overafter_tbox_tnumber.
func OverafterTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTemporalTstzspan wraps MEOS C function overafter_temporal_tstzspan.
func OverafterTemporalTstzspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_temporal_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTemporalTemporal wraps MEOS C function overafter_temporal_temporal.
func OverafterTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTnumberTBOX wraps MEOS C function overafter_tnumber_tbox.
func OverafterTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTnumberTnumber wraps MEOS C function overafter_tnumber_tnumber.
func OverafterTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTstzspanTemporal wraps MEOS C function overafter_tstzspan_temporal.
func OverafterTstzspanTemporal(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_tstzspan_temporal(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTBOXTnumber wraps MEOS C function overbefore_tbox_tnumber.
func OverbeforeTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTemporalTstzspan wraps MEOS C function overbefore_temporal_tstzspan.
func OverbeforeTemporalTstzspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_temporal_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTemporalTemporal wraps MEOS C function overbefore_temporal_temporal.
func OverbeforeTemporalTemporal(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_temporal_temporal(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTnumberTBOX wraps MEOS C function overbefore_tnumber_tbox.
func OverbeforeTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTnumberTnumber wraps MEOS C function overbefore_tnumber_tnumber.
func OverbeforeTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTstzspanTemporal wraps MEOS C function overbefore_tstzspan_temporal.
func OverbeforeTstzspanTemporal(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_tstzspan_temporal(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftNumspanTnumber wraps MEOS C function overleft_numspan_tnumber.
func OverleftNumspanTnumber(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_numspan_tnumber(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTBOXTnumber wraps MEOS C function overleft_tbox_tnumber.
func OverleftTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTnumberNumspan wraps MEOS C function overleft_tnumber_numspan.
func OverleftTnumberNumspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_tnumber_numspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTnumberTBOX wraps MEOS C function overleft_tnumber_tbox.
func OverleftTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTnumberTnumber wraps MEOS C function overleft_tnumber_tnumber.
func OverleftTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightNumspanTnumber wraps MEOS C function overright_numspan_tnumber.
func OverrightNumspanTnumber(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_numspan_tnumber(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTBOXTnumber wraps MEOS C function overright_tbox_tnumber.
func OverrightTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTnumberNumspan wraps MEOS C function overright_tnumber_numspan.
func OverrightTnumberNumspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_tnumber_numspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTnumberTBOX wraps MEOS C function overright_tnumber_tbox.
func OverrightTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTnumberTnumber wraps MEOS C function overright_tnumber_tnumber.
func OverrightTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightNumspanTnumber wraps MEOS C function right_numspan_tnumber.
func RightNumspanTnumber(s *Span, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_numspan_tnumber(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTBOXTnumber wraps MEOS C function right_tbox_tnumber.
func RightTBOXTnumber(box *TBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_tbox_tnumber(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTnumberNumspan wraps MEOS C function right_tnumber_numspan.
func RightTnumberNumspan(temp *Temporal, s *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_tnumber_numspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTnumberTBOX wraps MEOS C function right_tnumber_tbox.
func RightTnumberTBOX(temp *Temporal, box *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_tnumber_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTnumberTnumber wraps MEOS C function right_tnumber_tnumber.
func RightTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TandBoolTbool wraps MEOS C function tand_bool_tbool.
func TandBoolTbool(b bool, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tand_bool_tbool(C.bool(b), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TandTboolBool wraps MEOS C function tand_tbool_bool.
func TandTboolBool(temp *Temporal, b bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tand_tbool_bool(temp._inner, C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TandTboolTbool wraps MEOS C function tand_tbool_tbool.
func TandTboolTbool(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tand_tbool_tbool(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TboolWhenTrue wraps MEOS C function tbool_when_true.
func TboolWhenTrue(temp *Temporal) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_when_true(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TnotTbool wraps MEOS C function tnot_tbool.
func TnotTbool(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnot_tbool(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TorBoolTbool wraps MEOS C function tor_bool_tbool.
func TorBoolTbool(b bool, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tor_bool_tbool(C.bool(b), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TorTboolBool wraps MEOS C function tor_tbool_bool.
func TorTboolBool(temp *Temporal, b bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tor_tbool_bool(temp._inner, C.bool(b))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TorTboolTbool wraps MEOS C function tor_tbool_tbool.
func TorTboolTbool(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tor_tbool_tbool(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AddFloatTfloat wraps MEOS C function add_float_tfloat.
func AddFloatTfloat(d float64, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.add_float_tfloat(C.double(d), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AddIntTint wraps MEOS C function add_int_tint.
func AddIntTint(i int, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.add_int_tint(C.int(i), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AddTfloatFloat wraps MEOS C function add_tfloat_float.
func AddTfloatFloat(tnumber *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.add_tfloat_float(tnumber._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AddTintInt wraps MEOS C function add_tint_int.
func AddTintInt(tnumber *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.add_tint_int(tnumber._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AddBigintTbigint wraps MEOS C function add_bigint_tbigint.
func AddBigintTbigint(i int64, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.add_bigint_tbigint(C.int64_t(i), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AddTbigintBigint wraps MEOS C function add_tbigint_bigint.
func AddTbigintBigint(tnumber *Temporal, i int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.add_tbigint_bigint(tnumber._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AddTnumberTnumber wraps MEOS C function add_tnumber_tnumber.
func AddTnumberTnumber(tnumber1 *Temporal, tnumber2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.add_tnumber_tnumber(tnumber1._inner, tnumber2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// DivFloatTfloat wraps MEOS C function div_float_tfloat.
func DivFloatTfloat(d float64, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.div_float_tfloat(C.double(d), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// DivIntTint wraps MEOS C function div_int_tint.
func DivIntTint(i int, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.div_int_tint(C.int(i), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// DivTfloatFloat wraps MEOS C function div_tfloat_float.
func DivTfloatFloat(tnumber *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.div_tfloat_float(tnumber._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// DivTintInt wraps MEOS C function div_tint_int.
func DivTintInt(tnumber *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.div_tint_int(tnumber._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// DivBigintTbigint wraps MEOS C function div_bigint_tbigint.
func DivBigintTbigint(i int64, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.div_bigint_tbigint(C.int64_t(i), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// DivTbigintBigint wraps MEOS C function div_tbigint_bigint.
func DivTbigintBigint(tnumber *Temporal, i int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.div_tbigint_bigint(tnumber._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// DivTnumberTnumber wraps MEOS C function div_tnumber_tnumber.
func DivTnumberTnumber(tnumber1 *Temporal, tnumber2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.div_tnumber_tnumber(tnumber1._inner, tnumber2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// MulFloatTfloat wraps MEOS C function mul_float_tfloat.
func MulFloatTfloat(d float64, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_float_tfloat(C.double(d), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// MulIntTint wraps MEOS C function mul_int_tint.
func MulIntTint(i int, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_int_tint(C.int(i), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// MulTfloatFloat wraps MEOS C function mul_tfloat_float.
func MulTfloatFloat(tnumber *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_tfloat_float(tnumber._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// MulTintInt wraps MEOS C function mul_tint_int.
func MulTintInt(tnumber *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_tint_int(tnumber._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// MulBigintTbigint wraps MEOS C function mul_bigint_tbigint.
func MulBigintTbigint(i int64, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_bigint_tbigint(C.int64_t(i), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// MulTbigintBigint wraps MEOS C function mul_tbigint_bigint.
func MulTbigintBigint(tnumber *Temporal, i int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_tbigint_bigint(tnumber._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// MulTnumberTnumber wraps MEOS C function mul_tnumber_tnumber.
func MulTnumberTnumber(tnumber1 *Temporal, tnumber2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_tnumber_tnumber(tnumber1._inner, tnumber2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// SubFloatTfloat wraps MEOS C function sub_float_tfloat.
func SubFloatTfloat(d float64, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.sub_float_tfloat(C.double(d), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// SubIntTint wraps MEOS C function sub_int_tint.
func SubIntTint(i int, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.sub_int_tint(C.int(i), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// SubTfloatFloat wraps MEOS C function sub_tfloat_float.
func SubTfloatFloat(tnumber *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.sub_tfloat_float(tnumber._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// SubTintInt wraps MEOS C function sub_tint_int.
func SubTintInt(tnumber *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.sub_tint_int(tnumber._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// SubBigintTbigint wraps MEOS C function sub_bigint_tbigint.
func SubBigintTbigint(i int64, tnumber *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.sub_bigint_tbigint(C.int64_t(i), tnumber._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// SubTbigintBigint wraps MEOS C function sub_tbigint_bigint.
func SubTbigintBigint(tnumber *Temporal, i int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.sub_tbigint_bigint(tnumber._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// SubTnumberTnumber wraps MEOS C function sub_tnumber_tnumber.
func SubTnumberTnumber(tnumber1 *Temporal, tnumber2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.sub_tnumber_tnumber(tnumber1._inner, tnumber2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalDerivative wraps MEOS C function temporal_derivative.
func TemporalDerivative(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_derivative(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatExp wraps MEOS C function tfloat_exp.
func TfloatExp(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_exp(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatLn wraps MEOS C function tfloat_ln.
func TfloatLn(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_ln(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatLog10 wraps MEOS C function tfloat_log10.
func TfloatLog10(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_log10(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatSin wraps MEOS C function tfloat_sin.
func TfloatSin(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_sin(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatCos wraps MEOS C function tfloat_cos.
func TfloatCos(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_cos(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TfloatTan wraps MEOS C function tfloat_tan.
func TfloatTan(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_tan(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberAbs wraps MEOS C function tnumber_abs.
func TnumberAbs(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_abs(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberTrend wraps MEOS C function tnumber_trend.
func TnumberTrend(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_trend(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// FloatAngularDifference wraps MEOS C function float_angular_difference.
func FloatAngularDifference(degrees1 float64, degrees2 float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.float_angular_difference(C.double(degrees1), C.double(degrees2))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TnumberAngularDifference wraps MEOS C function tnumber_angular_difference.
func TnumberAngularDifference(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_angular_difference(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberDeltaValue wraps MEOS C function tnumber_delta_value.
func TnumberDeltaValue(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_delta_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TextcatTextTtext wraps MEOS C function textcat_text_ttext.
func TextcatTextTtext(txt string, temp *Temporal) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.textcat_text_ttext(_c_txt, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TextcatTtextText wraps MEOS C function textcat_ttext_text.
func TextcatTtextText(temp *Temporal, txt string) (_r0 *Temporal, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.textcat_ttext_text(temp._inner, _c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TextcatTtextTtext wraps MEOS C function textcat_ttext_ttext.
func TextcatTtextTtext(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.textcat_ttext_ttext(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtextInitcap wraps MEOS C function ttext_initcap.
func TtextInitcap(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_initcap(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtextUpper wraps MEOS C function ttext_upper.
func TtextUpper(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_upper(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtextLower wraps MEOS C function ttext_lower.
func TtextLower(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_lower(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTfloatFloat wraps MEOS C function tdistance_tfloat_float.
func TdistanceTfloatFloat(temp *Temporal, d float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTintInt wraps MEOS C function tdistance_tint_int.
func TdistanceTintInt(temp *Temporal, i int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TdistanceTnumberTnumber wraps MEOS C function tdistance_tnumber_tnumber.
func TdistanceTnumberTnumber(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tdistance_tnumber_tnumber(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// NadTboxfloatTboxfloat wraps MEOS C function nad_tboxfloat_tboxfloat.
func NadTboxfloatTboxfloat(box1 *TBox, box2 *TBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tboxfloat_tboxfloat(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTboxintTboxint wraps MEOS C function nad_tboxint_tboxint.
func NadTboxintTboxint(box1 *TBox, box2 *TBox) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tboxint_tboxint(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// NadTfloatFloat wraps MEOS C function nad_tfloat_float.
func NadTfloatFloat(temp *Temporal, d float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tfloat_float(temp._inner, C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTfloatTfloat wraps MEOS C function nad_tfloat_tfloat.
func NadTfloatTfloat(temp1 *Temporal, temp2 *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tfloat_tfloat(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTfloatTBOX wraps MEOS C function nad_tfloat_tbox.
func NadTfloatTBOX(temp *Temporal, box *TBox) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tfloat_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// NadTbigintBigint wraps MEOS C function nad_tbigint_bigint.
func NadTbigintBigint(temp *Temporal, i int64) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tbigint_bigint(temp._inner, C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// NadTbigintTBOX wraps MEOS C function nad_tbigint_tbox.
func NadTbigintTBOX(temp *Temporal, box *TBox) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tbigint_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// NadTbigintTbigint wraps MEOS C function nad_tbigint_tbigint.
func NadTbigintTbigint(temp1 *Temporal, temp2 *Temporal) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tbigint_tbigint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// NadTboxbigintTboxbigint wraps MEOS C function nad_tboxbigint_tboxbigint.
func NadTboxbigintTboxbigint(box1 *TBox, box2 *TBox) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tboxbigint_tboxbigint(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// NadTintInt wraps MEOS C function nad_tint_int.
func NadTintInt(temp *Temporal, i int) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tint_int(temp._inner, C.int(i))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// NadTintTBOX wraps MEOS C function nad_tint_tbox.
func NadTintTBOX(temp *Temporal, box *TBox) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tint_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// NadTintTint wraps MEOS C function nad_tint_tint.
func NadTintTint(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tint_tint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TbigintTmaxTransfn wraps MEOS C function tbigint_tmax_transfn.
func TbigintTmaxTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_tmax_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TbigintTmaxCombinefn wraps MEOS C function tbigint_tmax_combinefn.
func TbigintTmaxCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_tmax_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TbigintTminTransfn wraps MEOS C function tbigint_tmin_transfn.
func TbigintTminTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_tmin_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TbigintTminCombinefn wraps MEOS C function tbigint_tmin_combinefn.
func TbigintTminCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_tmin_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TbigintTsumTransfn wraps MEOS C function tbigint_tsum_transfn.
func TbigintTsumTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_tsum_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TbigintTsumCombinefn wraps MEOS C function tbigint_tsum_combinefn.
func TbigintTsumCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_tsum_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TbigintWmaxTransfn wraps MEOS C function tbigint_wmax_transfn.
func TbigintWmaxTransfn(state *SkipList, temp *Temporal, interv *Interval) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_wmax_transfn(state._inner, temp._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TbigintWminTransfn wraps MEOS C function tbigint_wmin_transfn.
func TbigintWminTransfn(state *SkipList, temp *Temporal, interv *Interval) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_wmin_transfn(state._inner, temp._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TbigintWsumTransfn wraps MEOS C function tbigint_wsum_transfn.
func TbigintWsumTransfn(state *SkipList, temp *Temporal, interv *Interval) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_wsum_transfn(state._inner, temp._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TboolTandTransfn wraps MEOS C function tbool_tand_transfn.
func TboolTandTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_tand_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TboolTandCombinefn wraps MEOS C function tbool_tand_combinefn.
func TboolTandCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_tand_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TboolTorTransfn wraps MEOS C function tbool_tor_transfn.
func TboolTorTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_tor_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TboolTorCombinefn wraps MEOS C function tbool_tor_combinefn.
func TboolTorCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tbool_tor_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TemporalExtentTransfn wraps MEOS C function temporal_extent_transfn.
func TemporalExtentTransfn(s *Span, temp *Temporal) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_extent_transfn(s._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TemporalTaggFinalfn wraps MEOS C function temporal_tagg_finalfn.
func TemporalTaggFinalfn(state *SkipList) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_tagg_finalfn(state._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalTcountTransfn wraps MEOS C function temporal_tcount_transfn.
func TemporalTcountTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_tcount_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TemporalTcountCombinefn wraps MEOS C function temporal_tcount_combinefn.
func TemporalTcountCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_tcount_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TfloatTmaxTransfn wraps MEOS C function tfloat_tmax_transfn.
func TfloatTmaxTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_tmax_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TfloatTmaxCombinefn wraps MEOS C function tfloat_tmax_combinefn.
func TfloatTmaxCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_tmax_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TfloatTminTransfn wraps MEOS C function tfloat_tmin_transfn.
func TfloatTminTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_tmin_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TfloatTminCombinefn wraps MEOS C function tfloat_tmin_combinefn.
func TfloatTminCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_tmin_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TfloatTsumTransfn wraps MEOS C function tfloat_tsum_transfn.
func TfloatTsumTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_tsum_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TfloatTsumCombinefn wraps MEOS C function tfloat_tsum_combinefn.
func TfloatTsumCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_tsum_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TfloatWmaxTransfn wraps MEOS C function tfloat_wmax_transfn.
func TfloatWmaxTransfn(state *SkipList, temp *Temporal, interv *Interval) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_wmax_transfn(state._inner, temp._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TfloatWminTransfn wraps MEOS C function tfloat_wmin_transfn.
func TfloatWminTransfn(state *SkipList, temp *Temporal, interv *Interval) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_wmin_transfn(state._inner, temp._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TfloatWsumTransfn wraps MEOS C function tfloat_wsum_transfn.
func TfloatWsumTransfn(state *SkipList, temp *Temporal, interv *Interval) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_wsum_transfn(state._inner, temp._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TimestamptzTcountTransfn wraps MEOS C function timestamptz_tcount_transfn.
func TimestamptzTcountTransfn(state *SkipList, t int64) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.timestamptz_tcount_transfn(state._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TintTmaxTransfn wraps MEOS C function tint_tmax_transfn.
func TintTmaxTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_tmax_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TintTmaxCombinefn wraps MEOS C function tint_tmax_combinefn.
func TintTmaxCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_tmax_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TintTminTransfn wraps MEOS C function tint_tmin_transfn.
func TintTminTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_tmin_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TintTminCombinefn wraps MEOS C function tint_tmin_combinefn.
func TintTminCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_tmin_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TintTsumTransfn wraps MEOS C function tint_tsum_transfn.
func TintTsumTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_tsum_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TintTsumCombinefn wraps MEOS C function tint_tsum_combinefn.
func TintTsumCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_tsum_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TintWmaxTransfn wraps MEOS C function tint_wmax_transfn.
func TintWmaxTransfn(state *SkipList, temp *Temporal, interv *Interval) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_wmax_transfn(state._inner, temp._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TintWminTransfn wraps MEOS C function tint_wmin_transfn.
func TintWminTransfn(state *SkipList, temp *Temporal, interv *Interval) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_wmin_transfn(state._inner, temp._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TintWsumTransfn wraps MEOS C function tint_wsum_transfn.
func TintWsumTransfn(state *SkipList, temp *Temporal, interv *Interval) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_wsum_transfn(state._inner, temp._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TnumberExtentTransfn wraps MEOS C function tnumber_extent_transfn.
func TnumberExtentTransfn(box *TBox, temp *Temporal) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_extent_transfn(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TnumberTavgFinalfn wraps MEOS C function tnumber_tavg_finalfn.
func TnumberTavgFinalfn(state *SkipList) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_tavg_finalfn(state._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberTavgTransfn wraps MEOS C function tnumber_tavg_transfn.
func TnumberTavgTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_tavg_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TnumberTavgCombinefn wraps MEOS C function tnumber_tavg_combinefn.
func TnumberTavgCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_tavg_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TnumberWavgTransfn wraps MEOS C function tnumber_wavg_transfn.
func TnumberWavgTransfn(state *SkipList, temp *Temporal, interv *Interval) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_wavg_transfn(state._inner, temp._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TstzsetTcountTransfn wraps MEOS C function tstzset_tcount_transfn.
func TstzsetTcountTransfn(state *SkipList, s *Set) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzset_tcount_transfn(state._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TstzspanTcountTransfn wraps MEOS C function tstzspan_tcount_transfn.
func TstzspanTcountTransfn(state *SkipList, s *Span) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspan_tcount_transfn(state._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TstzspansetTcountTransfn wraps MEOS C function tstzspanset_tcount_transfn.
func TstzspansetTcountTransfn(state *SkipList, ss *SpanSet) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tstzspanset_tcount_transfn(state._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TemporalMergeTransfn wraps MEOS C function temporal_merge_transfn.
func TemporalMergeTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_merge_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TemporalMergeCombinefn wraps MEOS C function temporal_merge_combinefn.
func TemporalMergeCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_merge_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TtextTmaxTransfn wraps MEOS C function ttext_tmax_transfn.
func TtextTmaxTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_tmax_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TtextTmaxCombinefn wraps MEOS C function ttext_tmax_combinefn.
func TtextTmaxCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_tmax_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TtextTminTransfn wraps MEOS C function ttext_tmin_transfn.
func TtextTminTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_tmin_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TtextTminCombinefn wraps MEOS C function ttext_tmin_combinefn.
func TtextTminCombinefn(state1 *SkipList, state2 *SkipList) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_tmin_combinefn(state1._inner, state2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TemporalSimplifyDp wraps MEOS C function temporal_simplify_dp.
func TemporalSimplifyDp(temp *Temporal, dist float64, synchronized bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_simplify_dp(temp._inner, C.double(dist), C.bool(synchronized))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalSimplifyMaxDist wraps MEOS C function temporal_simplify_max_dist.
func TemporalSimplifyMaxDist(temp *Temporal, dist float64, synchronized bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_simplify_max_dist(temp._inner, C.double(dist), C.bool(synchronized))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalSimplifyMinDist wraps MEOS C function temporal_simplify_min_dist.
func TemporalSimplifyMinDist(temp *Temporal, dist float64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_simplify_min_dist(temp._inner, C.double(dist))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalSimplifyMinTdelta wraps MEOS C function temporal_simplify_min_tdelta.
func TemporalSimplifyMinTdelta(temp *Temporal, mint *Interval) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_simplify_min_tdelta(temp._inner, mint._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalTprecision wraps MEOS C function temporal_tprecision.
func TemporalTprecision(temp *Temporal, duration *Interval, origin int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_tprecision(temp._inner, duration._inner, C.TimestampTz(origin))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalTsample wraps MEOS C function temporal_tsample.
func TemporalTsample(temp *Temporal, duration *Interval, origin int64, interp Interpolation) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_tsample(temp._inner, duration._inner, C.TimestampTz(origin), C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalDyntimewarpDistance wraps MEOS C function temporal_dyntimewarp_distance.
func TemporalDyntimewarpDistance(temp1 *Temporal, temp2 *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_dyntimewarp_distance(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TemporalDyntimewarpPath wraps MEOS C function temporal_dyntimewarp_path.
func TemporalDyntimewarpPath(temp1 *Temporal, temp2 *Temporal, count unsafe.Pointer) (_r0 *Match, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_dyntimewarp_path(temp1._inner, temp2._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Match{_inner: _cret}, nil
}


// TemporalFrechetDistance wraps MEOS C function temporal_frechet_distance.
func TemporalFrechetDistance(temp1 *Temporal, temp2 *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_frechet_distance(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TemporalFrechetPath wraps MEOS C function temporal_frechet_path.
func TemporalFrechetPath(temp1 *Temporal, temp2 *Temporal, count unsafe.Pointer) (_r0 *Match, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_frechet_path(temp1._inner, temp2._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Match{_inner: _cret}, nil
}


// TemporalHausdorffDistance wraps MEOS C function temporal_hausdorff_distance.
func TemporalHausdorffDistance(temp1 *Temporal, temp2 *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_hausdorff_distance(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TemporalAverageHausdorffDistance wraps MEOS C function temporal_average_hausdorff_distance.
func TemporalAverageHausdorffDistance(temp1 *Temporal, temp2 *Temporal) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_average_hausdorff_distance(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TemporalLcssDistance wraps MEOS C function temporal_lcss_distance.
func TemporalLcssDistance(temp1 *Temporal, temp2 *Temporal, epsilon float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_lcss_distance(temp1._inner, temp2._inner, C.double(epsilon))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TemporalExtKalmanFilter wraps MEOS C function temporal_ext_kalman_filter.
func TemporalExtKalmanFilter(temp *Temporal, gate float64, q float64, variance float64, to_drop bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_ext_kalman_filter(temp._inner, C.double(gate), C.double(q), C.double(variance), C.bool(to_drop))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalTimeBins wraps MEOS C function temporal_time_bins.
func TemporalTimeBins(temp *Temporal, duration *Interval, origin int64, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_time_bins(temp._inner, duration._inner, C.TimestampTz(origin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TemporalTimeSplit wraps MEOS C function temporal_time_split.
func TemporalTimeSplit(temp *Temporal, duration *Interval, torigin int64, bins unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_time_split(temp._inner, duration._inner, C.TimestampTz(torigin), (**C.TimestampTz)(unsafe.Pointer(bins)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TfloatTimeBoxes wraps MEOS C function tfloat_time_boxes.
func TfloatTimeBoxes(temp *Temporal, duration *Interval, torigin int64, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_time_boxes(temp._inner, duration._inner, C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TfloatValueBins wraps MEOS C function tfloat_value_bins.
func TfloatValueBins(temp *Temporal, vsize float64, vorigin float64, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_value_bins(temp._inner, C.double(vsize), C.double(vorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TfloatValueBoxes wraps MEOS C function tfloat_value_boxes.
func TfloatValueBoxes(temp *Temporal, vsize float64, vorigin float64, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_value_boxes(temp._inner, C.double(vsize), C.double(vorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TfloatValueSplit wraps MEOS C function tfloat_value_split.
func TfloatValueSplit(temp *Temporal, size float64, origin float64, bins unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_value_split(temp._inner, C.double(size), C.double(origin), (**C.double)(unsafe.Pointer(bins)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TfloatValueTimeBoxes wraps MEOS C function tfloat_value_time_boxes.
func TfloatValueTimeBoxes(temp *Temporal, vsize float64, duration *Interval, vorigin float64, torigin int64, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_value_time_boxes(temp._inner, C.double(vsize), duration._inner, C.double(vorigin), C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TfloatValueTimeSplit wraps MEOS C function tfloat_value_time_split.
func TfloatValueTimeSplit(temp *Temporal, vsize float64, duration *Interval, vorigin float64, torigin int64, value_bins unsafe.Pointer, time_bins unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloat_value_time_split(temp._inner, C.double(vsize), duration._inner, C.double(vorigin), C.TimestampTz(torigin), (**C.double)(unsafe.Pointer(value_bins)), (**C.TimestampTz)(unsafe.Pointer(time_bins)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TfloatboxTimeTiles wraps MEOS C function tfloatbox_time_tiles.
func TfloatboxTimeTiles(box *TBox, duration *Interval, torigin int64, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloatbox_time_tiles(box._inner, duration._inner, C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TfloatboxValueTiles wraps MEOS C function tfloatbox_value_tiles.
func TfloatboxValueTiles(box *TBox, vsize float64, vorigin float64, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloatbox_value_tiles(box._inner, C.double(vsize), C.double(vorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TfloatboxValueTimeTiles wraps MEOS C function tfloatbox_value_time_tiles.
func TfloatboxValueTimeTiles(box *TBox, vsize float64, duration *Interval, vorigin float64, torigin int64, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tfloatbox_value_time_tiles(box._inner, C.double(vsize), duration._inner, C.double(vorigin), C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TintTimeBoxes wraps MEOS C function tint_time_boxes.
func TintTimeBoxes(temp *Temporal, duration *Interval, torigin int64, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_time_boxes(temp._inner, duration._inner, C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TintValueBins wraps MEOS C function tint_value_bins.
func TintValueBins(temp *Temporal, vsize int, vorigin int, count unsafe.Pointer) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_value_bins(temp._inner, C.int(vsize), C.int(vorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// TintValueBoxes wraps MEOS C function tint_value_boxes.
func TintValueBoxes(temp *Temporal, vsize int, vorigin int, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_value_boxes(temp._inner, C.int(vsize), C.int(vorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TintValueSplit wraps MEOS C function tint_value_split.
func TintValueSplit(temp *Temporal, vsize int, vorigin int, bins unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_value_split(temp._inner, C.int(vsize), C.int(vorigin), (**C.int)(unsafe.Pointer(bins)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TintValueTimeBoxes wraps MEOS C function tint_value_time_boxes.
func TintValueTimeBoxes(temp *Temporal, vsize int, duration *Interval, vorigin int, torigin int64, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_value_time_boxes(temp._inner, C.int(vsize), duration._inner, C.int(vorigin), C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TintValueTimeSplit wraps MEOS C function tint_value_time_split.
func TintValueTimeSplit(temp *Temporal, size int, duration *Interval, vorigin int, torigin int64, value_bins unsafe.Pointer, time_bins unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tint_value_time_split(temp._inner, C.int(size), duration._inner, C.int(vorigin), C.TimestampTz(torigin), (**C.int)(unsafe.Pointer(value_bins)), (**C.TimestampTz)(unsafe.Pointer(time_bins)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TintboxTimeTiles wraps MEOS C function tintbox_time_tiles.
func TintboxTimeTiles(box *TBox, duration *Interval, torigin int64, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tintbox_time_tiles(box._inner, duration._inner, C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TintboxValueTiles wraps MEOS C function tintbox_value_tiles.
func TintboxValueTiles(box *TBox, xsize int, xorigin int, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tintbox_value_tiles(box._inner, C.int(xsize), C.int(xorigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}


// TintboxValueTimeTiles wraps MEOS C function tintbox_value_time_tiles.
func TintboxValueTimeTiles(box *TBox, xsize int, duration *Interval, xorigin int, torigin int64, count unsafe.Pointer) (_r0 *TBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tintbox_value_time_tiles(box._inner, C.int(xsize), duration._inner, C.int(xorigin), C.TimestampTz(torigin), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return &TBox{_inner: _cret}, nil
}

