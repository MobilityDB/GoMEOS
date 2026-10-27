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

// TODO char_hash: unsupported param char
// func CharHash(...) { /* not yet handled by codegen */ }


// TODO char_hash_extended: unsupported param char
// func CharHashExtended(...) { /* not yet handled by codegen */ }


// IcuUnicodeVersion wraps MEOS C function icu_unicode_version.
func IcuUnicodeVersion() (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.icu_unicode_version()
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// Int32ToBin wraps MEOS C function int32_to_bin.
func Int32ToBin(num int32) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_to_bin(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// Int32ToHex wraps MEOS C function int32_to_hex.
func Int32ToHex(num int32) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_to_hex(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// Int32ToOct wraps MEOS C function int32_to_oct.
func Int32ToOct(num int32) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.int32_to_oct(C.int32(num))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// Int64ToBin wraps MEOS C function int64_to_bin.
func Int64ToBin(num int64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_to_bin(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// Int64ToHex wraps MEOS C function int64_to_hex.
func Int64ToHex(num int64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_to_hex(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// Int64ToOct wraps MEOS C function int64_to_oct.
func Int64ToOct(num int64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.int64_to_oct(C.int64_t(num))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextCat wraps MEOS C function text_cat.
func TextCat(txt1 string, txt2 string) (_r0 string, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_cat(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextConcat wraps MEOS C function text_concat.
func TextConcat(textarr unsafe.Pointer, count int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.text_concat((**C.text)(unsafe.Pointer(textarr)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextConcatWs wraps MEOS C function text_concat_ws.
func TextConcatWs(textarr unsafe.Pointer, count int, sep string) (_r0 string, _err error) {
	_c_sep := C.cstring_to_text(C.CString(sep))
	defer C.free(unsafe.Pointer(_c_sep))
	C.meos_errno_reset()
	_cret := C.text_concat_ws((**C.text)(unsafe.Pointer(textarr)), C.int(count), _c_sep)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextEq wraps MEOS C function text_eq.
func TextEq(txt1 string, txt2 string) (_r0 bool, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_eq(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TextGe wraps MEOS C function text_ge.
func TextGe(txt1 string, txt2 string) (_r0 bool, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_ge(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TextGt wraps MEOS C function text_gt.
func TextGt(txt1 string, txt2 string) (_r0 bool, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_gt(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TODO text_hash: unsupported param Oid
// func TextHash(...) { /* not yet handled by codegen */ }


// TODO text_hash_extended: unsupported param Oid
// func TextHashExtended(...) { /* not yet handled by codegen */ }


// TextLarger wraps MEOS C function text_larger.
func TextLarger(txt1 string, txt2 string) (_r0 string, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_larger(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextLe wraps MEOS C function text_le.
func TextLe(txt1 string, txt2 string) (_r0 bool, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_le(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TextLeft wraps MEOS C function text_left.
func TextLeft(txt string, n int) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_left(_c_txt, C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextLen wraps MEOS C function text_len.
func TextLen(txt string) (_r0 int32, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_len(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// TextLt wraps MEOS C function text_lt.
func TextLt(txt1 string, txt2 string) (_r0 bool, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_lt(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TextNe wraps MEOS C function text_ne.
func TextNe(txt1 string, txt2 string) (_r0 bool, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_ne(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TextOctetlen wraps MEOS C function text_octetlen.
func TextOctetlen(txt string) (_r0 int32, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_octetlen(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// TextOverlay wraps MEOS C function text_overlay.
func TextOverlay(txt1 string, txt2 string, from int, count int) (_r0 string, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_overlay(_c_txt1, _c_txt2, C.int(from), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextOverlayNoLen wraps MEOS C function text_overlay_no_len.
func TextOverlayNoLen(txt1 string, txt2 string, from int) (_r0 string, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_overlay_no_len(_c_txt1, _c_txt2, C.int(from))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextPatternGe wraps MEOS C function text_pattern_ge.
func TextPatternGe(txt1 string, txt2 string) (_r0 bool, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_pattern_ge(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TextPatternGt wraps MEOS C function text_pattern_gt.
func TextPatternGt(txt1 string, txt2 string) (_r0 bool, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_pattern_gt(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TextPatternLe wraps MEOS C function text_pattern_le.
func TextPatternLe(txt1 string, txt2 string) (_r0 bool, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_pattern_le(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TextPatternLt wraps MEOS C function text_pattern_lt.
func TextPatternLt(txt1 string, txt2 string) (_r0 bool, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_pattern_lt(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TextPos wraps MEOS C function text_pos.
func TextPos(txt string, search string) (_r0 int32, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_c_search := C.cstring_to_text(C.CString(search))
	defer C.free(unsafe.Pointer(_c_search))
	C.meos_errno_reset()
	_cret := C.text_pos(_c_txt, _c_search)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// TextReplace wraps MEOS C function text_replace.
func TextReplace(txt string, from string, to string) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_c_from := C.cstring_to_text(C.CString(from))
	defer C.free(unsafe.Pointer(_c_from))
	_c_to := C.cstring_to_text(C.CString(to))
	defer C.free(unsafe.Pointer(_c_to))
	C.meos_errno_reset()
	_cret := C.text_replace(_c_txt, _c_from, _c_to)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextReverse wraps MEOS C function text_reverse.
func TextReverse(txt string) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_reverse(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextRight wraps MEOS C function text_right.
func TextRight(txt string, n int) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_right(_c_txt, C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextSmaller wraps MEOS C function text_smaller.
func TextSmaller(txt1 string, txt2 string) (_r0 string, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_smaller(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextSplitPart wraps MEOS C function text_split_part.
func TextSplitPart(txt string, sep string, fldnum int) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_c_sep := C.cstring_to_text(C.CString(sep))
	defer C.free(unsafe.Pointer(_c_sep))
	C.meos_errno_reset()
	_cret := C.text_split_part(_c_txt, _c_sep, C.int(fldnum))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextStartsWith wraps MEOS C function text_starts_with.
func TextStartsWith(txt1 string, txt2 string) (_r0 bool, _err error) {
	_c_txt1 := C.cstring_to_text(C.CString(txt1))
	defer C.free(unsafe.Pointer(_c_txt1))
	_c_txt2 := C.cstring_to_text(C.CString(txt2))
	defer C.free(unsafe.Pointer(_c_txt2))
	C.meos_errno_reset()
	_cret := C.text_starts_with(_c_txt1, _c_txt2)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TextSubstr wraps MEOS C function text_substr.
func TextSubstr(txt string, start int32, length int32) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_substr(_c_txt, C.int32(start), C.int32(length))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// TextSubstrNoLen wraps MEOS C function text_substr_no_len.
func TextSubstrNoLen(txt string, start int32) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.text_substr_no_len(_c_txt, C.int32(start))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// UnicodeAssigned wraps MEOS C function unicode_assigned.
func UnicodeAssigned(txt string) (_r0 bool, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.unicode_assigned(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// UnicodeIsNormalized wraps MEOS C function unicode_is_normalized.
func UnicodeIsNormalized(txt string, fmt string) (_r0 bool, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_c_fmt := C.cstring_to_text(C.CString(fmt))
	defer C.free(unsafe.Pointer(_c_fmt))
	C.meos_errno_reset()
	_cret := C.unicode_is_normalized(_c_txt, _c_fmt)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// UnicodeNormalizeFunc wraps MEOS C function unicode_normalize_func.
func UnicodeNormalizeFunc(txt string, fmt string) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	_c_fmt := C.cstring_to_text(C.CString(fmt))
	defer C.free(unsafe.Pointer(_c_fmt))
	C.meos_errno_reset()
	_cret := C.unicode_normalize_func(_c_txt, _c_fmt)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// UnicodeVersion wraps MEOS C function unicode_version.
func UnicodeVersion() (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.unicode_version()
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// Unistr wraps MEOS C function unistr.
func Unistr(txt string) (_r0 string, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.unistr(_c_txt)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}

