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

// PrngGetGenerationRng wraps MEOS C function prng_get_generation_rng.
func PrngGetGenerationRng() (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.prng_get_generation_rng()
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PrngGetAggregationRng wraps MEOS C function prng_get_aggregation_rng.
func PrngGetAggregationRng() (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.prng_get_aggregation_rng()
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// MeosRandomDouble wraps MEOS C function meos_random_double.
func MeosRandomDouble(rng unsafe.Pointer) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_random_double((*C.pg_prng_state)(unsafe.Pointer(rng)))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// MeosRandomExponential wraps MEOS C function meos_random_exponential.
func MeosRandomExponential(rng unsafe.Pointer, mean float64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_random_exponential((*C.pg_prng_state)(unsafe.Pointer(rng)), C.double(mean))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// MeosRandomBinomial20Half wraps MEOS C function meos_random_binomial20_half.
func MeosRandomBinomial20Half(rng unsafe.Pointer) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_random_binomial20_half((*C.pg_prng_state)(unsafe.Pointer(rng)))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// FloatspanRoundSet wraps MEOS C function floatspan_round_set.
func FloatspanRoundSet(s *Span, maxdd int) (_r0 *Span, _err error) {
	var _out_result C.Span
	C.meos_errno_reset()
	C.floatspan_round_set(s._inner, C.int(maxdd), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: &_out_result}, nil
}


// SetIn wraps MEOS C function set_in.
func SetIn(str string, basetype MeosType) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.set_in(_c_str, C.MeosType(basetype))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// SetOut wraps MEOS C function set_out.
func SetOut(s *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.set_out(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// SpanIn wraps MEOS C function span_in.
func SpanIn(str string, spantype MeosType) (_r0 *Span, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.span_in(_c_str, C.MeosType(spantype))
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpanOut wraps MEOS C function span_out.
func SpanOut(s *Span, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.span_out(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// SpansetIn wraps MEOS C function spanset_in.
func SpansetIn(str string, spantype MeosType) (_r0 *SpanSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.spanset_in(_c_str, C.MeosType(spantype))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// SpansetOut wraps MEOS C function spanset_out.
func SpansetOut(ss *SpanSet, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_out(ss._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// SpansetMakeExp wraps MEOS C function spanset_make_exp.
func SpansetMakeExp(spans *Span, count int, maxcount int, normalize bool, order bool) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_make_exp(spans._inner, C.int(count), C.int(maxcount), C.bool(normalize), C.bool(order))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// SpansetMakeFree wraps MEOS C function spanset_make_free.
func SpansetMakeFree(spans *Span, count int, normalize bool, order bool) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_make_free(spans._inner, C.int(count), C.bool(normalize), C.bool(order))
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// SetSpan wraps MEOS C function set_span.
func SetSpan(s *Set) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.set_span(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SetSpanset wraps MEOS C function set_spanset.
func SetSpanset(s *Set) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.set_spanset(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// SetMemSize wraps MEOS C function set_mem_size.
func SetMemSize(s *Set) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.set_mem_size(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// SetSetSubspan wraps MEOS C function set_set_subspan.
func SetSetSubspan(s *Set, minidx int, maxidx int) (_r0 *Span, _err error) {
	var _out_result C.Span
	C.meos_errno_reset()
	C.set_set_subspan(s._inner, C.int(minidx), C.int(maxidx), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: &_out_result}, nil
}


// SetSetSpan wraps MEOS C function set_set_span.
func SetSetSpan(s *Set) (_r0 *Span, _err error) {
	var _out_result C.Span
	C.meos_errno_reset()
	C.set_set_span(s._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: &_out_result}, nil
}


// SpansetMemSize wraps MEOS C function spanset_mem_size.
func SpansetMemSize(ss *SpanSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_mem_size(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// SpansetSps wraps MEOS C function spanset_sps.
func SpansetSps(ss *SpanSet, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_sps(ss._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// DatespanSetTstzspan wraps MEOS C function datespan_set_tstzspan.
func DatespanSetTstzspan(s1 *Span, s2 *Span) (_err error) {
	C.meos_errno_reset()
	C.datespan_set_tstzspan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// BigintspanSetFloatspan wraps MEOS C function bigintspan_set_floatspan.
func BigintspanSetFloatspan(s1 *Span, s2 *Span) (_err error) {
	C.meos_errno_reset()
	C.bigintspan_set_floatspan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// BigintspanSetIntspan wraps MEOS C function bigintspan_set_intspan.
func BigintspanSetIntspan(s1 *Span, s2 *Span) (_err error) {
	C.meos_errno_reset()
	C.bigintspan_set_intspan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// FloatspanSetBigintspan wraps MEOS C function floatspan_set_bigintspan.
func FloatspanSetBigintspan(s1 *Span, s2 *Span) (_err error) {
	C.meos_errno_reset()
	C.floatspan_set_bigintspan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// FloatspanSetIntspan wraps MEOS C function floatspan_set_intspan.
func FloatspanSetIntspan(s1 *Span, s2 *Span) (_err error) {
	C.meos_errno_reset()
	C.floatspan_set_intspan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// IntspanSetBigintspan wraps MEOS C function intspan_set_bigintspan.
func IntspanSetBigintspan(s1 *Span, s2 *Span) (_err error) {
	C.meos_errno_reset()
	C.intspan_set_bigintspan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// IntspanSetFloatspan wraps MEOS C function intspan_set_floatspan.
func IntspanSetFloatspan(s1 *Span, s2 *Span) (_err error) {
	C.meos_errno_reset()
	C.intspan_set_floatspan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// SetCompact wraps MEOS C function set_compact.
func SetCompact(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.set_compact(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// SpanExpand wraps MEOS C function span_expand.
func SpanExpand(s1 *Span, s2 *Span) (_err error) {
	C.meos_errno_reset()
	C.span_expand(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// SuperUnionSpanSpan wraps MEOS C function super_union_span_span.
func SuperUnionSpanSpan(s1 *Span, s2 *Span) (_r0 *Span, _err error) {
	C.meos_errno_reset()
	_cret := C.super_union_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: _cret}, nil
}


// SpansetCompact wraps MEOS C function spanset_compact.
func SpansetCompact(ss *SpanSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.spanset_compact(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TextcatTextsetTextCommon wraps MEOS C function textcat_textset_text_common.
func TextcatTextsetTextCommon(s *Set, txt string, invert bool) (_r0 *Set, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.textcat_textset_text_common(s._inner, _c_txt, C.bool(invert))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TstzspanSetDatespan wraps MEOS C function tstzspan_set_datespan.
func TstzspanSetDatespan(s1 *Span, s2 *Span) (_err error) {
	C.meos_errno_reset()
	C.tstzspan_set_datespan(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// OvadjSpanSpan wraps MEOS C function ovadj_span_span.
func OvadjSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ovadj_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LfnadjSpanSpan wraps MEOS C function lfnadj_span_span.
func LfnadjSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.lfnadj_span_span(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BboxType wraps MEOS C function bbox_type.
func BboxType(bboxtype MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.bbox_type(C.MeosType(bboxtype))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BboxGetSize wraps MEOS C function bbox_get_size.
func BboxGetSize(bboxtype MeosType) (_r0 uint, _err error) {
	C.meos_errno_reset()
	_cret := C.bbox_get_size(C.MeosType(bboxtype))
	if _err = meosError(); _err != nil {
		return
	}
	return uint(_cret), nil
}


// BboxMaxDims wraps MEOS C function bbox_max_dims.
func BboxMaxDims(bboxtype MeosType) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.bbox_max_dims(C.MeosType(bboxtype))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TemporalBboxEq wraps MEOS C function temporal_bbox_eq.
func TemporalBboxEq(box1 unsafe.Pointer, box2 unsafe.Pointer, temptype MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_bbox_eq(unsafe.Pointer(box1), unsafe.Pointer(box2), C.MeosType(temptype))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemporalBboxCmp wraps MEOS C function temporal_bbox_cmp.
func TemporalBboxCmp(box1 unsafe.Pointer, box2 unsafe.Pointer, temptype MeosType) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_bbox_cmp(unsafe.Pointer(box1), unsafe.Pointer(box2), C.MeosType(temptype))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EnsureBboxTemporalCompatible wraps MEOS C function ensure_bbox_temporal_compatible.
func EnsureBboxTemporalCompatible(bboxtype MeosType, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_bbox_temporal_compatible(C.MeosType(bboxtype), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureSameIndexBboxtype wraps MEOS C function ensure_same_index_bboxtype.
func EnsureSameIndexBboxtype(bboxtype1 MeosType, bboxtype2 MeosType) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_same_index_bboxtype(C.MeosType(bboxtype1), C.MeosType(bboxtype2))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// EnsureIndexJoinOp wraps MEOS C function ensure_index_join_op.
func EnsureIndexJoinOp(op IndexSearchOp) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.ensure_index_join_op(C.IndexSearchOp(op))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BboxTemporalSplitBoxes wraps MEOS C function bbox_temporal_split_boxes.
func BboxTemporalSplitBoxes(bboxtype MeosType, boxsize uint, temp *Temporal, maxboxes int, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.bbox_temporal_split_boxes(C.MeosType(bboxtype), C.size_t(boxsize), temp._inner, C.int(maxboxes), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SpanContains wraps MEOS C function span_contains.
func SpanContains(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_contains(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanContained wraps MEOS C function span_contained.
func SpanContained(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_contained(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanOverlaps wraps MEOS C function span_overlaps.
func SpanOverlaps(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_overlaps(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanSame wraps MEOS C function span_same.
func SpanSame(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_same(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanAdjacent wraps MEOS C function span_adjacent.
func SpanAdjacent(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_adjacent(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanLeft wraps MEOS C function span_left.
func SpanLeft(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_left(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanRight wraps MEOS C function span_right.
func SpanRight(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_right(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanOverleft wraps MEOS C function span_overleft.
func SpanOverleft(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_overleft(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SpanOverright wraps MEOS C function span_overright.
func SpanOverright(s1 *Span, s2 *Span) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.span_overright(s1._inner, s2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BboxUnionSpanSpan wraps MEOS C function bbox_union_span_span.
func BboxUnionSpanSpan(s1 *Span, s2 *Span) (_r0 *Span, _err error) {
	var _out_result C.Span
	C.meos_errno_reset()
	C.bbox_union_span_span(s1._inner, s2._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return &Span{_inner: &_out_result}, nil
}


// InterSpanSpan wraps MEOS C function inter_span_span.
func InterSpanSpan(s1 *Span, s2 *Span) (_r0 bool, _r1 *Span, _err error) {
	var _out_result C.Span
	C.meos_errno_reset()
	_cret := C.inter_span_span(s1._inner, s2._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Span{_inner: &_out_result}, nil
}


// MiSpanSpan wraps MEOS C function mi_span_span.
func MiSpanSpan(s1 *Span, s2 *Span) (_r0 int, _r1 *Span, _err error) {
	var _out_result C.Span
	C.meos_errno_reset()
	_cret := C.mi_span_span(s1._inner, s2._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), &Span{_inner: &_out_result}, nil
}


// TBOXSet wraps MEOS C function tbox_set.
func TBOXSet(s *Span, p *Span, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.tbox_set(s._inner, p._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// FloatSetTBOX wraps MEOS C function float_set_tbox.
func FloatSetTBOX(d float64, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.float_set_tbox(C.double(d), box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// IntSetTBOX wraps MEOS C function int_set_tbox.
func IntSetTBOX(i int, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.int_set_tbox(C.int(i), box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// NumsetSetTBOX wraps MEOS C function numset_set_tbox.
func NumsetSetTBOX(s *Set, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.numset_set_tbox(s._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// NumspanSetTBOX wraps MEOS C function numspan_set_tbox.
func NumspanSetTBOX(span *Span, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.numspan_set_tbox(span._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TimestamptzSetTBOX wraps MEOS C function timestamptz_set_tbox.
func TimestamptzSetTBOX(t int64, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.timestamptz_set_tbox(C.TimestampTz(t), box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TstzsetSetTBOX wraps MEOS C function tstzset_set_tbox.
func TstzsetSetTBOX(s *Set, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.tstzset_set_tbox(s._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TstzspanSetTBOX wraps MEOS C function tstzspan_set_tbox.
func TstzspanSetTBOX(s *Span, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.tstzspan_set_tbox(s._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TBOXExpand wraps MEOS C function tbox_expand.
func TBOXExpand(box1 *TBox, box2 *TBox) (_err error) {
	C.meos_errno_reset()
	C.tbox_expand(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TBOXContains wraps MEOS C function tbox_contains.
func TBOXContains(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_contains(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXContained wraps MEOS C function tbox_contained.
func TBOXContained(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_contained(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXOverlaps wraps MEOS C function tbox_overlaps.
func TBOXOverlaps(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_overlaps(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXSame wraps MEOS C function tbox_same.
func TBOXSame(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_same(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXAdjacent wraps MEOS C function tbox_adjacent.
func TBOXAdjacent(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_adjacent(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXLeft wraps MEOS C function tbox_left.
func TBOXLeft(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_left(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXRight wraps MEOS C function tbox_right.
func TBOXRight(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_right(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXOverleft wraps MEOS C function tbox_overleft.
func TBOXOverleft(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_overleft(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXOverright wraps MEOS C function tbox_overright.
func TBOXOverright(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_overright(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXBefore wraps MEOS C function tbox_before.
func TBOXBefore(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_before(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXAfter wraps MEOS C function tbox_after.
func TBOXAfter(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_after(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXOverbefore wraps MEOS C function tbox_overbefore.
func TBOXOverbefore(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_overbefore(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TBOXOverafter wraps MEOS C function tbox_overafter.
func TBOXOverafter(box1 *TBox, box2 *TBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tbox_overafter(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// InterTBOXTBOX wraps MEOS C function inter_tbox_tbox.
func InterTBOXTBOX(box1 *TBox, box2 *TBox) (_r0 bool, _r1 *TBox, _err error) {
	var _out_result C.TBox
	C.meos_errno_reset()
	_cret := C.inter_tbox_tbox(box1._inner, box2._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &TBox{_inner: &_out_result}, nil
}


// TboolinstIn wraps MEOS C function tboolinst_in.
func TboolinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tboolinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TboolseqIn wraps MEOS C function tboolseq_in.
func TboolseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tboolseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TboolseqsetIn wraps MEOS C function tboolseqset_in.
func TboolseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tboolseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TemporalIn wraps MEOS C function temporal_in.
func TemporalIn(str string, temptype MeosType) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.temporal_in(_c_str, C.MeosType(temptype))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalOut wraps MEOS C function temporal_out.
func TemporalOut(temp *Temporal, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_out(temp._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TemparrOut wraps MEOS C function temparr_out.
func TemparrOut(temparr unsafe.Pointer, count int, maxdd int) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.temparr_out((**C.Temporal)(unsafe.Pointer(temparr)), C.int(count), C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TfloatinstIn wraps MEOS C function tfloatinst_in.
func TfloatinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tfloatinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TfloatseqIn wraps MEOS C function tfloatseq_in.
func TfloatseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tfloatseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TfloatseqsetIn wraps MEOS C function tfloatseqset_in.
func TfloatseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tfloatseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TinstantIn wraps MEOS C function tinstant_in.
func TinstantIn(str string, temptype MeosType) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tinstant_in(_c_str, C.MeosType(temptype))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TinstantOut wraps MEOS C function tinstant_out.
func TinstantOut(inst *TInstant, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_out(inst._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TbigintinstIn wraps MEOS C function tbigintinst_in.
func TbigintinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tbigintinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TbigintseqsetIn wraps MEOS C function tbigintseqset_in.
func TbigintseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tbigintseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TintinstIn wraps MEOS C function tintinst_in.
func TintinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tintinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TintseqIn wraps MEOS C function tintseq_in.
func TintseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tintseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TintseqsetIn wraps MEOS C function tintseqset_in.
func TintseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tintseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequenceIn wraps MEOS C function tsequence_in.
func TsequenceIn(str string, temptype MeosType, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tsequence_in(_c_str, C.MeosType(temptype), C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TsequenceOut wraps MEOS C function tsequence_out.
func TsequenceOut(seq *TSequence, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_out(seq._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TsequencesetIn wraps MEOS C function tsequenceset_in.
func TsequencesetIn(str string, temptype MeosType, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tsequenceset_in(_c_str, C.MeosType(temptype), C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetOut wraps MEOS C function tsequenceset_out.
func TsequencesetOut(ss *TSequenceSet, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_out(ss._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TtextinstIn wraps MEOS C function ttextinst_in.
func TtextinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.ttextinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TtextseqIn wraps MEOS C function ttextseq_in.
func TtextseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.ttextseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TtextseqsetIn wraps MEOS C function ttextseqset_in.
func TtextseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.ttextseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TemporalFromMFJSON wraps MEOS C function temporal_from_mfjson.
func TemporalFromMFJSON(mfjson string, temptype MeosType) (_r0 *Temporal, _err error) {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	C.meos_errno_reset()
	_cret := C.temporal_from_mfjson(_c_mfjson, C.MeosType(temptype))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TinstantCopy wraps MEOS C function tinstant_copy.
func TinstantCopy(inst *TInstant) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_copy(inst._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TsequenceCopy wraps MEOS C function tsequence_copy.
func TsequenceCopy(seq *TSequence) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_copy(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TsequenceMakeExp wraps MEOS C function tsequence_make_exp.
func TsequenceMakeExp(instants unsafe.Pointer, count int, maxcount int, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_make_exp((**C.TInstant)(unsafe.Pointer(instants)), C.int(count), C.int(maxcount), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TsequenceMakeFree wraps MEOS C function tsequence_make_free.
func TsequenceMakeFree(instants unsafe.Pointer, count int, lower_inc bool, upper_inc bool, interp Interpolation, normalize bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_make_free((**C.TInstant)(unsafe.Pointer(instants)), C.int(count), C.bool(lower_inc), C.bool(upper_inc), C.interpType(interp), C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TsequencesetCopy wraps MEOS C function tsequenceset_copy.
func TsequencesetCopy(ss *TSequenceSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_copy(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TseqsetarrToTseqset wraps MEOS C function tseqsetarr_to_tseqset.
func TseqsetarrToTseqset(seqsets unsafe.Pointer, count int, totalseqs int) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tseqsetarr_to_tseqset((**C.TSequenceSet)(unsafe.Pointer(seqsets)), C.int(count), C.int(totalseqs))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetMakeExp wraps MEOS C function tsequenceset_make_exp.
func TsequencesetMakeExp(sequences unsafe.Pointer, count int, maxcount int, normalize bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_make_exp((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count), C.int(maxcount), C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetMakeFree wraps MEOS C function tsequenceset_make_free.
func TsequencesetMakeFree(sequences unsafe.Pointer, count int, normalize bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_make_free((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count), C.bool(normalize))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TemporalSetTstzspan wraps MEOS C function temporal_set_tstzspan.
func TemporalSetTstzspan(temp *Temporal, s *Span) (_err error) {
	C.meos_errno_reset()
	C.temporal_set_tstzspan(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TinstantSetTstzspan wraps MEOS C function tinstant_set_tstzspan.
func TinstantSetTstzspan(inst *TInstant, s *Span) (_err error) {
	C.meos_errno_reset()
	C.tinstant_set_tstzspan(inst._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TnumberSetTBOX wraps MEOS C function tnumber_set_tbox.
func TnumberSetTBOX(temp *Temporal, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.tnumber_set_tbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TnumberinstSetTBOX wraps MEOS C function tnumberinst_set_tbox.
func TnumberinstSetTBOX(inst *TInstant, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.tnumberinst_set_tbox(inst._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TnumberseqSetTBOX wraps MEOS C function tnumberseq_set_tbox.
func TnumberseqSetTBOX(seq *TSequence, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.tnumberseq_set_tbox(seq._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TnumberseqsetSetTBOX wraps MEOS C function tnumberseqset_set_tbox.
func TnumberseqsetSetTBOX(ss *TSequenceSet, box *TBox) (_err error) {
	C.meos_errno_reset()
	C.tnumberseqset_set_tbox(ss._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TsequenceSetTstzspan wraps MEOS C function tsequence_set_tstzspan.
func TsequenceSetTstzspan(seq *TSequence, s *Span) (_err error) {
	C.meos_errno_reset()
	C.tsequence_set_tstzspan(seq._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TsequencesetSetTstzspan wraps MEOS C function tsequenceset_set_tstzspan.
func TsequencesetSetTstzspan(ss *TSequenceSet, s *Span) (_err error) {
	C.meos_errno_reset()
	C.tsequenceset_set_tstzspan(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TemporalEndInst wraps MEOS C function temporal_end_inst.
func TemporalEndInst(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_end_inst(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TemporalInstN wraps MEOS C function temporal_inst_n.
func TemporalInstN(temp *Temporal, n int) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_inst_n(temp._inner, C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TemporalInstsP wraps MEOS C function temporal_insts_p.
func TemporalInstsP(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_insts_p(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TemporalMaxInstP wraps MEOS C function temporal_max_inst_p.
func TemporalMaxInstP(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_max_inst_p(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TemporalMemSize wraps MEOS C function temporal_mem_size.
func TemporalMemSize(temp *Temporal) (_r0 uint, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_mem_size(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint(_cret), nil
}


// TemporalMinInstP wraps MEOS C function temporal_min_inst_p.
func TemporalMinInstP(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_min_inst_p(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TemporalSequencesP wraps MEOS C function temporal_sequences_p.
func TemporalSequencesP(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_sequences_p(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TemporalSetBbox wraps MEOS C function temporal_set_bbox.
func TemporalSetBbox(temp *Temporal, box unsafe.Pointer) (_err error) {
	C.meos_errno_reset()
	C.temporal_set_bbox(temp._inner, unsafe.Pointer(box))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TemporalStartInst wraps MEOS C function temporal_start_inst.
func TemporalStartInst(temp *Temporal) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_start_inst(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TinstantHash wraps MEOS C function tinstant_hash.
func TinstantHash(inst *TInstant) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_hash(inst._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TinstantHashExtended wraps MEOS C function tinstant_hash_extended.
func TinstantHashExtended(inst *TInstant, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_hash_extended(inst._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TinstantInsts wraps MEOS C function tinstant_insts.
func TinstantInsts(inst *TInstant, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_insts(inst._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TinstantSetBbox wraps MEOS C function tinstant_set_bbox.
func TinstantSetBbox(inst *TInstant, box unsafe.Pointer) (_err error) {
	C.meos_errno_reset()
	C.tinstant_set_bbox(inst._inner, unsafe.Pointer(box))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TinstantTime wraps MEOS C function tinstant_time.
func TinstantTime(inst *TInstant) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_time(inst._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TinstantTimestamps wraps MEOS C function tinstant_timestamps.
func TinstantTimestamps(inst *TInstant, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_timestamps(inst._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TnumberSetSpan wraps MEOS C function tnumber_set_span.
func TnumberSetSpan(temp *Temporal, s *Span) (_err error) {
	C.meos_errno_reset()
	C.tnumber_set_span(temp._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TnumberinstValuespans wraps MEOS C function tnumberinst_valuespans.
func TnumberinstValuespans(inst *TInstant) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberinst_valuespans(inst._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TnumberseqAvgVal wraps MEOS C function tnumberseq_avg_val.
func TnumberseqAvgVal(seq *TSequence) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseq_avg_val(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TnumberseqValuespans wraps MEOS C function tnumberseq_valuespans.
func TnumberseqValuespans(seq *TSequence) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseq_valuespans(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TnumberseqsetAvgVal wraps MEOS C function tnumberseqset_avg_val.
func TnumberseqsetAvgVal(ss *TSequenceSet) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseqset_avg_val(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TnumberseqsetValuespans wraps MEOS C function tnumberseqset_valuespans.
func TnumberseqsetValuespans(ss *TSequenceSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseqset_valuespans(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TsequenceDuration wraps MEOS C function tsequence_duration.
func TsequenceDuration(seq *TSequence) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_duration(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// TsequenceEndTimestamptz wraps MEOS C function tsequence_end_timestamptz.
func TsequenceEndTimestamptz(seq *TSequence) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_end_timestamptz(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TsequenceHash wraps MEOS C function tsequence_hash.
func TsequenceHash(seq *TSequence) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_hash(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TsequenceHashExtended wraps MEOS C function tsequence_hash_extended.
func TsequenceHashExtended(seq *TSequence, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_hash_extended(seq._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TsequenceInstsP wraps MEOS C function tsequence_insts_p.
func TsequenceInstsP(seq *TSequence, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_insts_p(seq._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TsequenceMaxInstP wraps MEOS C function tsequence_max_inst_p.
func TsequenceMaxInstP(seq *TSequence) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_max_inst_p(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TsequenceMinInstP wraps MEOS C function tsequence_min_inst_p.
func TsequenceMinInstP(seq *TSequence) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_min_inst_p(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TsequenceSegments wraps MEOS C function tsequence_segments.
func TsequenceSegments(seq *TSequence, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_segments(seq._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TsequenceSeqs wraps MEOS C function tsequence_seqs.
func TsequenceSeqs(seq *TSequence, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_seqs(seq._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TsequenceStartTimestamptz wraps MEOS C function tsequence_start_timestamptz.
func TsequenceStartTimestamptz(seq *TSequence) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_start_timestamptz(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TsequenceTime wraps MEOS C function tsequence_time.
func TsequenceTime(seq *TSequence) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_time(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TsequenceTimestamps wraps MEOS C function tsequence_timestamps.
func TsequenceTimestamps(seq *TSequence, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_timestamps(seq._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TsequencesetDuration wraps MEOS C function tsequenceset_duration.
func TsequencesetDuration(ss *TSequenceSet, boundspan bool) (_r0 *Interval, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_duration(ss._inner, C.bool(boundspan))
	if _err = meosError(); _err != nil {
		return
	}
	return &Interval{_inner: _cret}, nil
}


// TsequencesetEndTimestamptz wraps MEOS C function tsequenceset_end_timestamptz.
func TsequencesetEndTimestamptz(ss *TSequenceSet) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_end_timestamptz(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TsequencesetHash wraps MEOS C function tsequenceset_hash.
func TsequencesetHash(ss *TSequenceSet) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_hash(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TsequencesetHashExtended wraps MEOS C function tsequenceset_hash_extended.
func TsequencesetHashExtended(ss *TSequenceSet, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_hash_extended(ss._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TsequencesetInstN wraps MEOS C function tsequenceset_inst_n.
func TsequencesetInstN(ss *TSequenceSet, n int) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_inst_n(ss._inner, C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TsequencesetInstsP wraps MEOS C function tsequenceset_insts_p.
func TsequencesetInstsP(ss *TSequenceSet, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_insts_p(ss._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TsequencesetMaxInstP wraps MEOS C function tsequenceset_max_inst_p.
func TsequencesetMaxInstP(ss *TSequenceSet) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_max_inst_p(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TsequencesetMinInstP wraps MEOS C function tsequenceset_min_inst_p.
func TsequencesetMinInstP(ss *TSequenceSet) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_min_inst_p(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TsequencesetNumInstants wraps MEOS C function tsequenceset_num_instants.
func TsequencesetNumInstants(ss *TSequenceSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_num_instants(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TsequencesetNumTimestamps wraps MEOS C function tsequenceset_num_timestamps.
func TsequencesetNumTimestamps(ss *TSequenceSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_num_timestamps(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TsequencesetSegments wraps MEOS C function tsequenceset_segments.
func TsequencesetSegments(ss *TSequenceSet, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_segments(ss._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TsequencesetSequencesP wraps MEOS C function tsequenceset_sequences_p.
func TsequencesetSequencesP(ss *TSequenceSet) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_sequences_p(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TsequencesetStartTimestamptz wraps MEOS C function tsequenceset_start_timestamptz.
func TsequencesetStartTimestamptz(ss *TSequenceSet) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_start_timestamptz(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TsequencesetTime wraps MEOS C function tsequenceset_time.
func TsequencesetTime(ss *TSequenceSet) (_r0 *SpanSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_time(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SpanSet{_inner: _cret}, nil
}


// TsequencesetTimestamptzN wraps MEOS C function tsequenceset_timestamptz_n.
func TsequencesetTimestamptzN(ss *TSequenceSet, n int) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.TimestampTz
	C.meos_errno_reset()
	_cret := C.tsequenceset_timestamptz_n(ss._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// TsequencesetTimestamps wraps MEOS C function tsequenceset_timestamps.
func TsequencesetTimestamps(ss *TSequenceSet, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_timestamps(ss._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TemporalRestart wraps MEOS C function temporal_restart.
func TemporalRestart(temp *Temporal, count int) (_err error) {
	C.meos_errno_reset()
	C.temporal_restart(temp._inner, C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TemporalTsequence wraps MEOS C function temporal_tsequence.
func TemporalTsequence(temp *Temporal, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_tsequence(temp._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TemporalTsequenceset wraps MEOS C function temporal_tsequenceset.
func TemporalTsequenceset(temp *Temporal, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_tsequenceset(temp._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TinstantShiftTime wraps MEOS C function tinstant_shift_time.
func TinstantShiftTime(inst *TInstant, interv *Interval) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_shift_time(inst._inner, interv._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TinstantAsTsequence wraps MEOS C function tinstant_as_tsequence.
func TinstantAsTsequence(inst *TInstant, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_as_tsequence(inst._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TinstantToTsequenceFree wraps MEOS C function tinstant_to_tsequence_free.
func TinstantToTsequenceFree(inst *TInstant, interp Interpolation) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_to_tsequence_free(inst._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TinstantAsTsequenceset wraps MEOS C function tinstant_as_tsequenceset.
func TinstantAsTsequenceset(inst *TInstant, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_as_tsequenceset(inst._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequenceRestart wraps MEOS C function tsequence_restart.
func TsequenceRestart(seq *TSequence, count int) (_err error) {
	C.meos_errno_reset()
	C.tsequence_restart(seq._inner, C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TsequenceSetInterp wraps MEOS C function tsequence_set_interp.
func TsequenceSetInterp(seq *TSequence, interp Interpolation) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_set_interp(seq._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceShiftScaleTime wraps MEOS C function tsequence_shift_scale_time.
func TsequenceShiftScaleTime(seq *TSequence, shift *Interval, duration *Interval) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_shift_scale_time(seq._inner, shift._inner, duration._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TsequenceSubseq wraps MEOS C function tsequence_subseq.
func TsequenceSubseq(seq *TSequence, from int, to int, lower_inc bool, upper_inc bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_subseq(seq._inner, C.int(from), C.int(to), C.bool(lower_inc), C.bool(upper_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TsequenceAsTinstant wraps MEOS C function tsequence_as_tinstant.
func TsequenceAsTinstant(seq *TSequence) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_as_tinstant(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TsequenceAsTsequenceset wraps MEOS C function tsequence_as_tsequenceset.
func TsequenceAsTsequenceset(seq *TSequence) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_as_tsequenceset(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequenceToTsequencesetFree wraps MEOS C function tsequence_to_tsequenceset_free.
func TsequenceToTsequencesetFree(seq *TSequence) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_to_tsequenceset_free(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequenceToTsequencesetInterp wraps MEOS C function tsequence_to_tsequenceset_interp.
func TsequenceToTsequencesetInterp(seq *TSequence, interp Interpolation) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_to_tsequenceset_interp(seq._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetRestart wraps MEOS C function tsequenceset_restart.
func TsequencesetRestart(ss *TSequenceSet, count int) (_err error) {
	C.meos_errno_reset()
	C.tsequenceset_restart(ss._inner, C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TsequencesetSetInterp wraps MEOS C function tsequenceset_set_interp.
func TsequencesetSetInterp(ss *TSequenceSet, interp Interpolation) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_set_interp(ss._inner, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequencesetShiftScaleTime wraps MEOS C function tsequenceset_shift_scale_time.
func TsequencesetShiftScaleTime(ss *TSequenceSet, start *Interval, duration *Interval) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_shift_scale_time(ss._inner, start._inner, duration._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetToDiscrete wraps MEOS C function tsequenceset_to_discrete.
func TsequencesetToDiscrete(ss *TSequenceSet) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_to_discrete(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TsequencesetToLinear wraps MEOS C function tsequenceset_to_linear.
func TsequencesetToLinear(ss *TSequenceSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_to_linear(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetToStep wraps MEOS C function tsequenceset_to_step.
func TsequencesetToStep(ss *TSequenceSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_to_step(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetAsTinstant wraps MEOS C function tsequenceset_as_tinstant.
func TsequencesetAsTinstant(ss *TSequenceSet) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_as_tinstant(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TsequencesetAsTsequence wraps MEOS C function tsequenceset_as_tsequence.
func TsequencesetAsTsequence(ss *TSequenceSet) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_as_tsequence(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TinstantMerge wraps MEOS C function tinstant_merge.
func TinstantMerge(inst1 *TInstant, inst2 *TInstant) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_merge(inst1._inner, inst2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TinstantMergeArray wraps MEOS C function tinstant_merge_array.
func TinstantMergeArray(instants unsafe.Pointer, count int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_merge_array((**C.TInstant)(unsafe.Pointer(instants)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceAppendTinstant wraps MEOS C function tsequence_append_tinstant.
func TsequenceAppendTinstant(seq *TSequence, inst *TInstant, maxdist float64, maxt *Interval, expand bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_append_tinstant(seq._inner, inst._inner, C.double(maxdist), maxt._inner, C.bool(expand))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceAppendTsequence wraps MEOS C function tsequence_append_tsequence.
func TsequenceAppendTsequence(seq1 *TSequence, seq2 *TSequence, expand bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_append_tsequence(seq1._inner, seq2._inner, C.bool(expand))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceDeleteTimestamptz wraps MEOS C function tsequence_delete_timestamptz.
func TsequenceDeleteTimestamptz(seq *TSequence, t int64, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_delete_timestamptz(seq._inner, C.TimestampTz(t), C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceDeleteTstzset wraps MEOS C function tsequence_delete_tstzset.
func TsequenceDeleteTstzset(seq *TSequence, s *Set, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_delete_tstzset(seq._inner, s._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceDeleteTstzspan wraps MEOS C function tsequence_delete_tstzspan.
func TsequenceDeleteTstzspan(seq *TSequence, s *Span, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_delete_tstzspan(seq._inner, s._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceDeleteTstzspanset wraps MEOS C function tsequence_delete_tstzspanset.
func TsequenceDeleteTstzspanset(seq *TSequence, ss *SpanSet, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_delete_tstzspanset(seq._inner, ss._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceInsert wraps MEOS C function tsequence_insert.
func TsequenceInsert(seq1 *TSequence, seq2 *TSequence, connect bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_insert(seq1._inner, seq2._inner, C.bool(connect))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceMerge wraps MEOS C function tsequence_merge.
func TsequenceMerge(seq1 *TSequence, seq2 *TSequence) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_merge(seq1._inner, seq2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceMergeArray wraps MEOS C function tsequence_merge_array.
func TsequenceMergeArray(sequences unsafe.Pointer, count int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_merge_array((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequencesetAppendTinstant wraps MEOS C function tsequenceset_append_tinstant.
func TsequencesetAppendTinstant(ss *TSequenceSet, inst *TInstant, maxdist float64, maxt *Interval, expand bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_append_tinstant(ss._inner, inst._inner, C.double(maxdist), maxt._inner, C.bool(expand))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetAppendTsequence wraps MEOS C function tsequenceset_append_tsequence.
func TsequencesetAppendTsequence(ss *TSequenceSet, seq *TSequence, expand bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_append_tsequence(ss._inner, seq._inner, C.bool(expand))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetDeleteTimestamptz wraps MEOS C function tsequenceset_delete_timestamptz.
func TsequencesetDeleteTimestamptz(ss *TSequenceSet, t int64) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_delete_timestamptz(ss._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetDeleteTstzset wraps MEOS C function tsequenceset_delete_tstzset.
func TsequencesetDeleteTstzset(ss *TSequenceSet, s *Set) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_delete_tstzset(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetDeleteTstzspan wraps MEOS C function tsequenceset_delete_tstzspan.
func TsequencesetDeleteTstzspan(ss *TSequenceSet, s *Span) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_delete_tstzspan(ss._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetDeleteTstzspanset wraps MEOS C function tsequenceset_delete_tstzspanset.
func TsequencesetDeleteTstzspanset(ss *TSequenceSet, ps *SpanSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_delete_tstzspanset(ss._inner, ps._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetInsert wraps MEOS C function tsequenceset_insert.
func TsequencesetInsert(ss1 *TSequenceSet, ss2 *TSequenceSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_insert(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetMerge wraps MEOS C function tsequenceset_merge.
func TsequencesetMerge(ss1 *TSequenceSet, ss2 *TSequenceSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_merge(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetMergeArray wraps MEOS C function tsequenceset_merge_array.
func TsequencesetMergeArray(seqsets unsafe.Pointer, count int) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_merge_array((**C.TSequenceSet)(unsafe.Pointer(seqsets)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequenceExpandBbox wraps MEOS C function tsequence_expand_bbox.
func TsequenceExpandBbox(seq *TSequence, inst *TInstant) (_err error) {
	C.meos_errno_reset()
	C.tsequence_expand_bbox(seq._inner, inst._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TsequenceSetBbox wraps MEOS C function tsequence_set_bbox.
func TsequenceSetBbox(seq *TSequence, box unsafe.Pointer) (_err error) {
	C.meos_errno_reset()
	C.tsequence_set_bbox(seq._inner, unsafe.Pointer(box))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TsequencesetExpandBbox wraps MEOS C function tsequenceset_expand_bbox.
func TsequencesetExpandBbox(ss *TSequenceSet, seq *TSequence) (_err error) {
	C.meos_errno_reset()
	C.tsequenceset_expand_bbox(ss._inner, seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TsequencesetSetBbox wraps MEOS C function tsequenceset_set_bbox.
func TsequencesetSetBbox(ss *TSequenceSet, box unsafe.Pointer) (_err error) {
	C.meos_errno_reset()
	C.tsequenceset_set_bbox(ss._inner, unsafe.Pointer(box))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TcontseqAfterTimestamptz wraps MEOS C function tcontseq_after_timestamptz.
func TcontseqAfterTimestamptz(seq *TSequence, t int64, strict bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tcontseq_after_timestamptz(seq._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TcontseqBeforeTimestamptz wraps MEOS C function tcontseq_before_timestamptz.
func TcontseqBeforeTimestamptz(seq *TSequence, t int64, strict bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tcontseq_before_timestamptz(seq._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TcontseqRestrictMinmax wraps MEOS C function tcontseq_restrict_minmax.
func TcontseqRestrictMinmax(seq *TSequence, min bool, atfunc bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tcontseq_restrict_minmax(seq._inner, C.bool(min), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TdiscseqAfterTimestamptz wraps MEOS C function tdiscseq_after_timestamptz.
func TdiscseqAfterTimestamptz(seq *TSequence, t int64, strict bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tdiscseq_after_timestamptz(seq._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TdiscseqBeforeTimestamptz wraps MEOS C function tdiscseq_before_timestamptz.
func TdiscseqBeforeTimestamptz(seq *TSequence, t int64, strict bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tdiscseq_before_timestamptz(seq._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TdiscseqRestrictMinmax wraps MEOS C function tdiscseq_restrict_minmax.
func TdiscseqRestrictMinmax(seq *TSequence, min bool, atfunc bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tdiscseq_restrict_minmax(seq._inner, C.bool(min), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TemporalBboxRestrictSet wraps MEOS C function temporal_bbox_restrict_set.
func TemporalBboxRestrictSet(temp *Temporal, set *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_bbox_restrict_set(temp._inner, set._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TemporalRestrictMinmax wraps MEOS C function temporal_restrict_minmax.
func TemporalRestrictMinmax(temp *Temporal, min bool, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_restrict_minmax(temp._inner, C.bool(min), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalRestrictTimestamptz wraps MEOS C function temporal_restrict_timestamptz.
func TemporalRestrictTimestamptz(temp *Temporal, t int64, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_restrict_timestamptz(temp._inner, C.TimestampTz(t), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalRestrictTstzset wraps MEOS C function temporal_restrict_tstzset.
func TemporalRestrictTstzset(temp *Temporal, s *Set, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_restrict_tstzset(temp._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalRestrictTstzspan wraps MEOS C function temporal_restrict_tstzspan.
func TemporalRestrictTstzspan(temp *Temporal, s *Span, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_restrict_tstzspan(temp._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalRestrictTstzspanset wraps MEOS C function temporal_restrict_tstzspanset.
func TemporalRestrictTstzspanset(temp *Temporal, ss *SpanSet, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_restrict_tstzspanset(temp._inner, ss._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalRestrictValues wraps MEOS C function temporal_restrict_values.
func TemporalRestrictValues(temp *Temporal, set *Set, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_restrict_values(temp._inner, set._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TinstantAfterTimestamptz wraps MEOS C function tinstant_after_timestamptz.
func TinstantAfterTimestamptz(inst *TInstant, t int64, strict bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_after_timestamptz(inst._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TinstantBeforeTimestamptz wraps MEOS C function tinstant_before_timestamptz.
func TinstantBeforeTimestamptz(inst *TInstant, t int64, strict bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_before_timestamptz(inst._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TinstantRestrictTstzspan wraps MEOS C function tinstant_restrict_tstzspan.
func TinstantRestrictTstzspan(inst *TInstant, period *Span, atfunc bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_restrict_tstzspan(inst._inner, period._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TinstantRestrictTstzspanset wraps MEOS C function tinstant_restrict_tstzspanset.
func TinstantRestrictTstzspanset(inst *TInstant, ss *SpanSet, atfunc bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_restrict_tstzspanset(inst._inner, ss._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TinstantRestrictTimestamptz wraps MEOS C function tinstant_restrict_timestamptz.
func TinstantRestrictTimestamptz(inst *TInstant, t int64, atfunc bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_restrict_timestamptz(inst._inner, C.TimestampTz(t), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TinstantRestrictTstzset wraps MEOS C function tinstant_restrict_tstzset.
func TinstantRestrictTstzset(inst *TInstant, s *Set, atfunc bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_restrict_tstzset(inst._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TinstantRestrictValues wraps MEOS C function tinstant_restrict_values.
func TinstantRestrictValues(inst *TInstant, set *Set, atfunc bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_restrict_values(inst._inner, set._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TnumberRestrictSpan wraps MEOS C function tnumber_restrict_span.
func TnumberRestrictSpan(temp *Temporal, span *Span, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_restrict_span(temp._inner, span._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberRestrictSpanset wraps MEOS C function tnumber_restrict_spanset.
func TnumberRestrictSpanset(temp *Temporal, ss *SpanSet, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumber_restrict_spanset(temp._inner, ss._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnumberinstRestrictSpan wraps MEOS C function tnumberinst_restrict_span.
func TnumberinstRestrictSpan(inst *TInstant, span *Span, atfunc bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberinst_restrict_span(inst._inner, span._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TnumberinstRestrictSpanset wraps MEOS C function tnumberinst_restrict_spanset.
func TnumberinstRestrictSpanset(inst *TInstant, ss *SpanSet, atfunc bool) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberinst_restrict_spanset(inst._inner, ss._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TnumberseqsetRestrictSpan wraps MEOS C function tnumberseqset_restrict_span.
func TnumberseqsetRestrictSpan(ss *TSequenceSet, span *Span, atfunc bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseqset_restrict_span(ss._inner, span._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TnumberseqsetRestrictSpanset wraps MEOS C function tnumberseqset_restrict_spanset.
func TnumberseqsetRestrictSpanset(ss *TSequenceSet, spanset *SpanSet, atfunc bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseqset_restrict_spanset(ss._inner, spanset._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequenceAtTimestamptz wraps MEOS C function tsequence_at_timestamptz.
func TsequenceAtTimestamptz(seq *TSequence, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_at_timestamptz(seq._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TsequenceRestrictTstzspan wraps MEOS C function tsequence_restrict_tstzspan.
func TsequenceRestrictTstzspan(seq *TSequence, s *Span, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_restrict_tstzspan(seq._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceRestrictTstzspanset wraps MEOS C function tsequence_restrict_tstzspanset.
func TsequenceRestrictTstzspanset(seq *TSequence, ss *SpanSet, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_restrict_tstzspanset(seq._inner, ss._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequencesetAfterTimestamptz wraps MEOS C function tsequenceset_after_timestamptz.
func TsequencesetAfterTimestamptz(ss *TSequenceSet, t int64, strict bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_after_timestamptz(ss._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetBeforeTimestamptz wraps MEOS C function tsequenceset_before_timestamptz.
func TsequencesetBeforeTimestamptz(ss *TSequenceSet, t int64, strict bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_before_timestamptz(ss._inner, C.TimestampTz(t), C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetRestrictMinmax wraps MEOS C function tsequenceset_restrict_minmax.
func TsequencesetRestrictMinmax(ss *TSequenceSet, min bool, atfunc bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_restrict_minmax(ss._inner, C.bool(min), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetRestrictTstzspan wraps MEOS C function tsequenceset_restrict_tstzspan.
func TsequencesetRestrictTstzspan(ss *TSequenceSet, s *Span, atfunc bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_restrict_tstzspan(ss._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetRestrictTstzspanset wraps MEOS C function tsequenceset_restrict_tstzspanset.
func TsequencesetRestrictTstzspanset(ss *TSequenceSet, ps *SpanSet, atfunc bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_restrict_tstzspanset(ss._inner, ps._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TsequencesetRestrictTimestamptz wraps MEOS C function tsequenceset_restrict_timestamptz.
func TsequencesetRestrictTimestamptz(ss *TSequenceSet, t int64, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_restrict_timestamptz(ss._inner, C.TimestampTz(t), C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequencesetRestrictTstzset wraps MEOS C function tsequenceset_restrict_tstzset.
func TsequencesetRestrictTstzset(ss *TSequenceSet, s *Set, atfunc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_restrict_tstzset(ss._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequencesetRestrictValues wraps MEOS C function tsequenceset_restrict_values.
func TsequencesetRestrictValues(ss *TSequenceSet, s *Set, atfunc bool) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_restrict_values(ss._inner, s._inner, C.bool(atfunc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TinstantCmp wraps MEOS C function tinstant_cmp.
func TinstantCmp(inst1 *TInstant, inst2 *TInstant) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_cmp(inst1._inner, inst2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TinstantEq wraps MEOS C function tinstant_eq.
func TinstantEq(inst1 *TInstant, inst2 *TInstant) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tinstant_eq(inst1._inner, inst2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TsequenceCmp wraps MEOS C function tsequence_cmp.
func TsequenceCmp(seq1 *TSequence, seq2 *TSequence) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_cmp(seq1._inner, seq2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TsequenceEq wraps MEOS C function tsequence_eq.
func TsequenceEq(seq1 *TSequence, seq2 *TSequence) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_eq(seq1._inner, seq2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TsequencesetCmp wraps MEOS C function tsequenceset_cmp.
func TsequencesetCmp(ss1 *TSequenceSet, ss2 *TSequenceSet) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_cmp(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TsequencesetEq wraps MEOS C function tsequenceset_eq.
func TsequencesetEq(ss1 *TSequenceSet, ss2 *TSequenceSet) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_eq(ss1._inner, ss2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TnumberinstAbs wraps MEOS C function tnumberinst_abs.
func TnumberinstAbs(inst *TInstant) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberinst_abs(inst._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TnumberinstDistance wraps MEOS C function tnumberinst_distance.
func TnumberinstDistance(inst1 *TInstant, inst2 *TInstant) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberinst_distance(inst1._inner, inst2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TnumberseqAbs wraps MEOS C function tnumberseq_abs.
func TnumberseqAbs(seq *TSequence) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseq_abs(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TnumberseqAngularDifference wraps MEOS C function tnumberseq_angular_difference.
func TnumberseqAngularDifference(seq *TSequence) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseq_angular_difference(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TnumberseqDeltaValue wraps MEOS C function tnumberseq_delta_value.
func TnumberseqDeltaValue(seq *TSequence) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseq_delta_value(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TnumberseqsetAbs wraps MEOS C function tnumberseqset_abs.
func TnumberseqsetAbs(ss *TSequenceSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseqset_abs(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TnumberseqsetAngularDifference wraps MEOS C function tnumberseqset_angular_difference.
func TnumberseqsetAngularDifference(ss *TSequenceSet) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseqset_angular_difference(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TnumberseqsetDeltaValue wraps MEOS C function tnumberseqset_delta_value.
func TnumberseqsetDeltaValue(ss *TSequenceSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseqset_delta_value(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TnumberseqIntegral wraps MEOS C function tnumberseq_integral.
func TnumberseqIntegral(seq *TSequence) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseq_integral(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TnumberseqTwavg wraps MEOS C function tnumberseq_twavg.
func TnumberseqTwavg(seq *TSequence) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseq_twavg(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TnumberseqsetIntegral wraps MEOS C function tnumberseqset_integral.
func TnumberseqsetIntegral(ss *TSequenceSet) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseqset_integral(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TnumberseqsetTwavg wraps MEOS C function tnumberseqset_twavg.
func TnumberseqsetTwavg(ss *TSequenceSet) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.tnumberseqset_twavg(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// TemporalCompact wraps MEOS C function temporal_compact.
func TemporalCompact(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_compact(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TsequenceCompact wraps MEOS C function tsequence_compact.
func TsequenceCompact(seq *TSequence) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequence_compact(seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TsequencesetCompact wraps MEOS C function tsequenceset_compact.
func TsequencesetCompact(ss *TSequenceSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tsequenceset_compact(ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TemporalSkiplistMake wraps MEOS C function temporal_skiplist_make.
func TemporalSkiplistMake() (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_skiplist_make()
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// SkiplistMake wraps MEOS C function skiplist_make.
func SkiplistMake(key_size uint, value_size uint, comp_fn unsafe.Pointer, merge_fn unsafe.Pointer) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.skiplist_make(C.size_t(key_size), C.size_t(value_size), (*[0]byte)(comp_fn), (*[0]byte)(merge_fn))
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// SkiplistSearch wraps MEOS C function skiplist_search.
func SkiplistSearch(list *SkipList, key unsafe.Pointer, value unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.skiplist_search(list._inner, unsafe.Pointer(key), unsafe.Pointer(value))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// SkiplistFree wraps MEOS C function skiplist_free.
func SkiplistFree(list *SkipList) (_err error) {
	C.meos_errno_reset()
	C.skiplist_free(list._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// SkiplistSplice wraps MEOS C function skiplist_splice.
func SkiplistSplice(list *SkipList, keys unsafe.Pointer, values unsafe.Pointer, count int, func_ unsafe.Pointer, crossings bool, sktype SkipListType) (_err error) {
	C.meos_errno_reset()
	C.skiplist_splice(list._inner, (*unsafe.Pointer)(unsafe.Pointer(keys)), (*unsafe.Pointer)(unsafe.Pointer(values)), C.int(count), (*[0]byte)(func_), C.bool(crossings), C.SkipListType(sktype))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// TemporalSkiplistSplice wraps MEOS C function temporal_skiplist_splice.
func TemporalSkiplistSplice(list *SkipList, values unsafe.Pointer, count int, func_ unsafe.Pointer, crossings bool) (_err error) {
	C.meos_errno_reset()
	C.temporal_skiplist_splice(list._inner, (*unsafe.Pointer)(unsafe.Pointer(values)), C.int(count), (*[0]byte)(func_), C.bool(crossings))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// SkiplistValues wraps MEOS C function skiplist_values.
func SkiplistValues(list *SkipList) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.skiplist_values(list._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SkiplistKeysValues wraps MEOS C function skiplist_keys_values.
func SkiplistKeysValues(list *SkipList, values unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.skiplist_keys_values(list._inner, (*unsafe.Pointer)(unsafe.Pointer(values)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TemporalAppTinstTransfn wraps MEOS C function temporal_app_tinst_transfn.
func TemporalAppTinstTransfn(state *Temporal, inst *TInstant, interp Interpolation, maxdist float64, maxt *Interval) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_app_tinst_transfn(state._inner, inst._inner, C.interpType(interp), C.double(maxdist), maxt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TemporalAppTseqTransfn wraps MEOS C function temporal_app_tseq_transfn.
func TemporalAppTseqTransfn(state *Temporal, seq *TSequence) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.temporal_app_tseq_transfn(state._inner, seq._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}

