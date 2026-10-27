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

// JSONIn wraps MEOS C function json_in.
func JSONIn(str string) (_r0 string, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.json_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JSONOut wraps MEOS C function json_out.
func JSONOut(js string) (_r0 string, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_out(_c_js)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// JsonbFromText wraps MEOS C function jsonb_from_text.
func JsonbFromText(txt string, unique_keys bool) (_r0 *Jsonb, _err error) {
	_c_txt := C.cstring_to_text(C.CString(txt))
	defer C.free(unsafe.Pointer(_c_txt))
	C.meos_errno_reset()
	_cret := C.jsonb_from_text(_c_txt, C.bool(unique_keys))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbIn wraps MEOS C function jsonb_in.
func JsonbIn(str string) (_r0 *Jsonb, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.jsonb_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbOut wraps MEOS C function jsonb_out.
func JsonbOut(jb *Jsonb) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_out(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// JSONMake wraps MEOS C function json_make.
func JSONMake(keys_vals unsafe.Pointer, count int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.json_make((**C.text)(unsafe.Pointer(keys_vals)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JSONMakeTwoArg wraps MEOS C function json_make_two_arg.
func JSONMakeTwoArg(keys unsafe.Pointer, values unsafe.Pointer, count int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.json_make_two_arg((**C.text)(unsafe.Pointer(keys)), (**C.text)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JsonbCopy wraps MEOS C function jsonb_copy.
func JsonbCopy(jb *Jsonb) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_copy(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbMake wraps MEOS C function jsonb_make.
func JsonbMake(keys_vals unsafe.Pointer, count int) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_make((**C.text)(unsafe.Pointer(keys_vals)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbMakeTwoArg wraps MEOS C function jsonb_make_two_arg.
func JsonbMakeTwoArg(keys unsafe.Pointer, values unsafe.Pointer, count int) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_make_two_arg((**C.text)(unsafe.Pointer(keys)), (**C.text)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbToBool wraps MEOS C function jsonb_to_bool.
func JsonbToBool(jb *Jsonb) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_to_bool(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbToCstring wraps MEOS C function jsonb_to_cstring.
func JsonbToCstring(jb *Jsonb) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_to_cstring(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TODO jsonb_to_float4: unsupported return type float4
// func JsonbToFloat4(...) { /* not yet handled by codegen */ }


// TODO jsonb_to_float8: unsupported return type float8
// func JsonbToFloat8(...) { /* not yet handled by codegen */ }


// JsonbToInt16 wraps MEOS C function jsonb_to_int16.
func JsonbToInt16(jb *Jsonb) (_r0 int16, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_to_int16(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int16(_cret), nil
}


// JsonbToInt32 wraps MEOS C function jsonb_to_int32.
func JsonbToInt32(jb *Jsonb) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_to_int32(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// JsonbToInt64 wraps MEOS C function jsonb_to_int64.
func JsonbToInt64(jb *Jsonb) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_to_int64(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// TODO jsonb_to_numeric: unsupported return type Numeric
// func JsonbToNumeric(...) { /* not yet handled by codegen */ }


// JsonbToText wraps MEOS C function jsonb_to_text.
func JsonbToText(jb *Jsonb) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_to_text(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JSONArrayElement wraps MEOS C function json_array_element.
func JSONArrayElement(js string, element int) (_r0 string, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_array_element(_c_js, C.int(element))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JSONArrayElementText wraps MEOS C function json_array_element_text.
func JSONArrayElementText(js string, element int) (_r0 string, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_array_element_text(_c_js, C.int(element))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JSONArrayElements wraps MEOS C function json_array_elements.
func JSONArrayElements(js string, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_array_elements(_c_js, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JSONArrayElementsText wraps MEOS C function json_array_elements_text.
func JSONArrayElementsText(js string, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_array_elements_text(_c_js, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JSONArrayLength wraps MEOS C function json_array_length.
func JSONArrayLength(js string) (_r0 int, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_array_length(_c_js)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// JSONEach wraps MEOS C function json_each.
func JSONEach(js string, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_each(_c_js, (***C.text)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JSONEachText wraps MEOS C function json_each_text.
func JSONEachText(js string, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_each_text(_c_js, (***C.text)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JSONExtractPath wraps MEOS C function json_extract_path.
func JSONExtractPath(js string, path_elems unsafe.Pointer, path_len int) (_r0 string, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_extract_path(_c_js, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JSONExtractPathText wraps MEOS C function json_extract_path_text.
func JSONExtractPathText(js string, path_elems unsafe.Pointer, path_len int) (_r0 string, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_extract_path_text(_c_js, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JSONObjectField wraps MEOS C function json_object_field.
func JSONObjectField(js string, key string) (_r0 string, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.json_object_field(_c_js, _c_key)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JSONObjectFieldText wraps MEOS C function json_object_field_text.
func JSONObjectFieldText(js string, key string) (_r0 string, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.json_object_field_text(_c_js, _c_key)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JSONObjectKeys wraps MEOS C function json_object_keys.
func JSONObjectKeys(js string, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_object_keys(_c_js, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JSONTypeof wraps MEOS C function json_typeof.
func JSONTypeof(js string) (_r0 string, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_typeof(_c_js)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JsonbArrayElement wraps MEOS C function jsonb_array_element.
func JsonbArrayElement(jb *Jsonb, element int) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_array_element(jb._inner, C.int(element))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbArrayElementText wraps MEOS C function jsonb_array_element_text.
func JsonbArrayElementText(jb *Jsonb, element int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_array_element_text(jb._inner, C.int(element))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JsonbArrayElements wraps MEOS C function jsonb_array_elements.
func JsonbArrayElements(jb *Jsonb, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_array_elements(jb._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JsonbArrayElementsText wraps MEOS C function jsonb_array_elements_text.
func JsonbArrayElementsText(jb *Jsonb, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_array_elements_text(jb._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JsonbArrayLength wraps MEOS C function jsonb_array_length.
func JsonbArrayLength(jb *Jsonb) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_array_length(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// JsonbContained wraps MEOS C function jsonb_contained.
func JsonbContained(jb1 *Jsonb, jb2 *Jsonb) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_contained(jb1._inner, jb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbContains wraps MEOS C function jsonb_contains.
func JsonbContains(jb1 *Jsonb, jb2 *Jsonb) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_contains(jb1._inner, jb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbEach wraps MEOS C function jsonb_each.
func JsonbEach(jb *Jsonb, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_each(jb._inner, (***C.Jsonb)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JsonbEachText wraps MEOS C function jsonb_each_text.
func JsonbEachText(jb *Jsonb, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_each_text(jb._inner, (***C.text)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JsonbExists wraps MEOS C function jsonb_exists.
func JsonbExists(jb *Jsonb, key string) (_r0 bool, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonb_exists(jb._inner, _c_key)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbExistsArray wraps MEOS C function jsonb_exists_array.
func JsonbExistsArray(jb *Jsonb, keys_elems unsafe.Pointer, keys_len int, any bool) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_exists_array(jb._inner, (**C.text)(unsafe.Pointer(keys_elems)), C.int(keys_len), C.bool(any))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbExtractPath wraps MEOS C function jsonb_extract_path.
func JsonbExtractPath(jb *Jsonb, path_elems unsafe.Pointer, path_len int) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_extract_path(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbExtractPathText wraps MEOS C function jsonb_extract_path_text.
func JsonbExtractPathText(jb *Jsonb, path_elems unsafe.Pointer, path_len int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_extract_path_text(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JsonbHash wraps MEOS C function jsonb_hash.
func JsonbHash(jb *Jsonb) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_hash(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// JsonbHashExtended wraps MEOS C function jsonb_hash_extended.
func JsonbHashExtended(jb *Jsonb, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_hash_extended(jb._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// JsonbMemSize wraps MEOS C function jsonb_mem_size.
func JsonbMemSize(jb *Jsonb) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_mem_size(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// JsonbObjectField wraps MEOS C function jsonb_object_field.
func JsonbObjectField(jb *Jsonb, key string) (_r0 *Jsonb, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonb_object_field(jb._inner, _c_key)
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbObjectFieldText wraps MEOS C function jsonb_object_field_text.
func JsonbObjectFieldText(jb *Jsonb, key string) (_r0 string, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonb_object_field_text(jb._inner, _c_key)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JsonbObjectKeys wraps MEOS C function jsonb_object_keys.
func JsonbObjectKeys(jb *Jsonb, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_object_keys(jb._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JSONStripNulls wraps MEOS C function json_strip_nulls.
func JSONStripNulls(js string, strip_in_arrays bool) (_r0 string, _err error) {
	_c_js := C.cstring_to_text(C.CString(js))
	defer C.free(unsafe.Pointer(_c_js))
	C.meos_errno_reset()
	_cret := C.json_strip_nulls(_c_js, C.bool(strip_in_arrays))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JsonbConcat wraps MEOS C function jsonb_concat.
func JsonbConcat(jb1 *Jsonb, jb2 *Jsonb) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_concat(jb1._inner, jb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbDelete wraps MEOS C function jsonb_delete.
func JsonbDelete(jb *Jsonb, key string) (_r0 *Jsonb, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonb_delete(jb._inner, _c_key)
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbDeleteArray wraps MEOS C function jsonb_delete_array.
func JsonbDeleteArray(jb *Jsonb, keys_elems unsafe.Pointer, keys_len int) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_delete_array(jb._inner, (**C.text)(unsafe.Pointer(keys_elems)), C.int(keys_len))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbDeleteIndex wraps MEOS C function jsonb_delete_index.
func JsonbDeleteIndex(jb *Jsonb, idx int) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_delete_index(jb._inner, C.int(idx))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbDeletePath wraps MEOS C function jsonb_delete_path.
func JsonbDeletePath(jb *Jsonb, path_elems unsafe.Pointer, path_len int) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_delete_path(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbInsert wraps MEOS C function jsonb_insert.
func JsonbInsert(jb *Jsonb, path_elems unsafe.Pointer, path_len int, newjb *Jsonb, after bool) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_insert(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), newjb._inner, C.bool(after))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbPretty wraps MEOS C function jsonb_pretty.
func JsonbPretty(jb *Jsonb) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_pretty(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(C.text_to_cstring(_cret)), nil
}


// JsonbSet wraps MEOS C function jsonb_set.
func JsonbSet(jb *Jsonb, path_elems unsafe.Pointer, path_len int, newjb *Jsonb, create bool) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_set(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), newjb._inner, C.bool(create))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbSetLax wraps MEOS C function jsonb_set_lax.
func JsonbSetLax(jb *Jsonb, path_elems unsafe.Pointer, path_len int, newjb *Jsonb, create bool, handle_null string) (_r0 *Jsonb, _err error) {
	_c_handle_null := C.cstring_to_text(C.CString(handle_null))
	defer C.free(unsafe.Pointer(_c_handle_null))
	C.meos_errno_reset()
	_cret := C.jsonb_set_lax(jb._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), newjb._inner, C.bool(create), _c_handle_null)
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbStripNulls wraps MEOS C function jsonb_strip_nulls.
func JsonbStripNulls(jb *Jsonb, strip_in_arrays bool) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_strip_nulls(jb._inner, C.bool(strip_in_arrays))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbCmp wraps MEOS C function jsonb_cmp.
func JsonbCmp(jb1 *Jsonb, jb2 *Jsonb) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_cmp(jb1._inner, jb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// JsonbEq wraps MEOS C function jsonb_eq.
func JsonbEq(jb1 *Jsonb, jb2 *Jsonb) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_eq(jb1._inner, jb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbGe wraps MEOS C function jsonb_ge.
func JsonbGe(jb1 *Jsonb, jb2 *Jsonb) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_ge(jb1._inner, jb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbGt wraps MEOS C function jsonb_gt.
func JsonbGt(jb1 *Jsonb, jb2 *Jsonb) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_gt(jb1._inner, jb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbLe wraps MEOS C function jsonb_le.
func JsonbLe(jb1 *Jsonb, jb2 *Jsonb) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_le(jb1._inner, jb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbLt wraps MEOS C function jsonb_lt.
func JsonbLt(jb1 *Jsonb, jb2 *Jsonb) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_lt(jb1._inner, jb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbNe wraps MEOS C function jsonb_ne.
func JsonbNe(jb1 *Jsonb, jb2 *Jsonb) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_ne(jb1._inner, jb2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbPathExists wraps MEOS C function jsonb_path_exists.
func JsonbPathExists(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_path_exists(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// JsonbPathMatch wraps MEOS C function jsonb_path_match.
func JsonbPathMatch(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_path_match(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// JsonbPathQueryAll wraps MEOS C function jsonb_path_query_all.
func JsonbPathQueryAll(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_path_query_all(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JsonbPathQueryArray wraps MEOS C function jsonb_path_query_array.
func JsonbPathQueryArray(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_path_query_array(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbPathQueryFirst wraps MEOS C function jsonb_path_query_first.
func JsonbPathQueryFirst(jb *Jsonb, jp *JsonPath, vars *Jsonb, silent bool, tz bool) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_path_query_first(jb._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonpathIn wraps MEOS C function jsonpath_in.
func JsonpathIn(str string) (_r0 *JsonPath, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.jsonpath_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &JsonPath{_inner: _cret}, nil
}


// JsonpathCopy wraps MEOS C function jsonpath_copy.
func JsonpathCopy(jp *JsonPath) (_r0 *JsonPath, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonpath_copy(jp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &JsonPath{_inner: _cret}, nil
}


// JsonpathMemSize wraps MEOS C function jsonpath_mem_size.
func JsonpathMemSize(jp *JsonPath) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonpath_mem_size(jp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// JsonpathOut wraps MEOS C function jsonpath_out.
func JsonpathOut(jp *JsonPath) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonpath_out(jp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// JsonbsetIn wraps MEOS C function jsonbset_in.
func JsonbsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.jsonbset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetOut wraps MEOS C function jsonbset_out.
func JsonbsetOut(s *Set) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_out(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// JsonbsetMake wraps MEOS C function jsonbset_make.
func JsonbsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_make((**C.Jsonb)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbToSet wraps MEOS C function jsonb_to_set.
func JsonbToSet(jb *Jsonb) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_to_set(jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetToTextset wraps MEOS C function jsonbset_to_textset.
func JsonbsetToTextset(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_to_textset(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TextsetToJsonbset wraps MEOS C function textset_to_jsonbset.
func TextsetToJsonbset(s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.textset_to_jsonbset(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetEndValue wraps MEOS C function jsonbset_end_value.
func JsonbsetEndValue(s *Set) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbsetStartValue wraps MEOS C function jsonbset_start_value.
func JsonbsetStartValue(s *Set) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// JsonbsetValueN wraps MEOS C function jsonbset_value_n.
func JsonbsetValueN(s *Set, n int) (_r0 bool, _r1 *Jsonb, _err error) {
	var _out_result *C.Jsonb
	C.meos_errno_reset()
	_cret := C.jsonbset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Jsonb{_inner: _out_result}, nil
}


// JsonbsetValues wraps MEOS C function jsonbset_values.
func JsonbsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ConcatJsonbsetJsonb wraps MEOS C function concat_jsonbset_jsonb.
func ConcatJsonbsetJsonb(s *Set, jb *Jsonb, invert bool) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.concat_jsonbset_jsonb(s._inner, jb._inner, C.bool(invert))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetArrayLength wraps MEOS C function jsonbset_array_length.
func JsonbsetArrayLength(set *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_array_length(set._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetObjectField wraps MEOS C function jsonbset_object_field.
func JsonbsetObjectField(set *Set, key string, null_handle NullHandleType) (_r0 *Set, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonbset_object_field(set._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetObjectFieldText wraps MEOS C function jsonbset_object_field_text.
func JsonbsetObjectFieldText(set *Set, key string, null_handle NullHandleType) (_r0 *Set, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonbset_object_field_text(set._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetArrayElement wraps MEOS C function jsonbset_array_element.
func JsonbsetArrayElement(set *Set, idx int, null_handle NullHandleType) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_array_element(set._inner, C.int(idx), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetArrayElementText wraps MEOS C function jsonbset_array_element_text.
func JsonbsetArrayElementText(set *Set, idx int, null_handle NullHandleType) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_array_element_text(set._inner, C.int(idx), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetDeleteIndex wraps MEOS C function jsonbset_delete_index.
func JsonbsetDeleteIndex(set *Set, idx int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_delete_index(set._inner, C.int(idx))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetDelete wraps MEOS C function jsonbset_delete.
func JsonbsetDelete(set *Set, key string) (_r0 *Set, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonbset_delete(set._inner, _c_key)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetDeleteArray wraps MEOS C function jsonbset_delete_array.
func JsonbsetDeleteArray(set *Set, keys unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_delete_array(set._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetExists wraps MEOS C function jsonbset_exists.
func JsonbsetExists(set *Set, key string, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonbset_exists(set._inner, _c_key, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JsonbsetExistsArray wraps MEOS C function jsonbset_exists_array.
func JsonbsetExistsArray(set *Set, keys unsafe.Pointer, count int, any bool, rescount unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_exists_array(set._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count), C.bool(any), (*C.int)(unsafe.Pointer(rescount)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JsonbsetSet wraps MEOS C function jsonbset_set.
func JsonbsetSet(set *Set, keys unsafe.Pointer, count int, newjb *Jsonb, create bool, null_handle string, lax bool) (_r0 *Set, _err error) {
	_c_null_handle := C.cstring_to_text(C.CString(null_handle))
	defer C.free(unsafe.Pointer(_c_null_handle))
	C.meos_errno_reset()
	_cret := C.jsonbset_set(set._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count), newjb._inner, C.bool(create), _c_null_handle, C.bool(lax))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetToAlphanumset wraps MEOS C function jsonbset_to_alphanumset.
func JsonbsetToAlphanumset(set *Set, key string, settype MeosType, null_handle NullHandleType) (_r0 *Set, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonbset_to_alphanumset(set._inner, _c_key, C.MeosType(settype), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetToIntset wraps MEOS C function jsonbset_to_intset.
func JsonbsetToIntset(set *Set, key string, null_handle NullHandleType) (_r0 *Set, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonbset_to_intset(set._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetToBigintset wraps MEOS C function jsonbset_to_bigintset.
func JsonbsetToBigintset(set *Set, key string, null_handle NullHandleType) (_r0 *Set, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonbset_to_bigintset(set._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetToFloatset wraps MEOS C function jsonbset_to_floatset.
func JsonbsetToFloatset(set *Set, key string, null_handle NullHandleType) (_r0 *Set, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonbset_to_floatset(set._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetToTextsetKey wraps MEOS C function jsonbset_to_textset_key.
func JsonbsetToTextsetKey(set *Set, key string, null_handle NullHandleType) (_r0 *Set, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonbset_to_textset_key(set._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetStripNulls wraps MEOS C function jsonbset_strip_nulls.
func JsonbsetStripNulls(set *Set, strip_in_arrays bool) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_strip_nulls(set._inner, C.bool(strip_in_arrays))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetPretty wraps MEOS C function jsonbset_pretty.
func JsonbsetPretty(set *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_pretty(set._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetDeletePath wraps MEOS C function jsonbset_delete_path.
func JsonbsetDeletePath(set *Set, path_elems unsafe.Pointer, path_len int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_delete_path(set._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetExtractPath wraps MEOS C function jsonbset_extract_path.
func JsonbsetExtractPath(set *Set, path_elems unsafe.Pointer, path_len int, null_handle NullHandleType) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_extract_path(set._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetExtractPathText wraps MEOS C function jsonbset_extract_path_text.
func JsonbsetExtractPathText(set *Set, path_elems unsafe.Pointer, path_len int, null_handle NullHandleType) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_extract_path_text(set._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetInsert wraps MEOS C function jsonbset_insert.
func JsonbsetInsert(set *Set, path_elems unsafe.Pointer, path_len int, newjb *Jsonb, after bool) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_insert(set._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), newjb._inner, C.bool(after))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetPathExists wraps MEOS C function jsonbset_path_exists.
func JsonbsetPathExists(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_path_exists(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JsonbsetPathMatch wraps MEOS C function jsonbset_path_match.
func JsonbsetPathMatch(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_path_match(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// JsonbsetPathQueryArray wraps MEOS C function jsonbset_path_query_array.
func JsonbsetPathQueryArray(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_path_query_array(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetPathQueryFirst wraps MEOS C function jsonbset_path_query_first.
func JsonbsetPathQueryFirst(set *Set, jp *JsonPath, vars *Jsonb, silent bool, tz bool) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_path_query_first(set._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetArrayElementCommon wraps MEOS C function jsonbset_array_element_common.
func JsonbsetArrayElementCommon(set *Set, idx int, astext bool, null_handle NullHandleType) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_array_element_common(set._inner, C.int(idx), C.bool(astext), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetExtractPathCommon wraps MEOS C function jsonbset_extract_path_common.
func JsonbsetExtractPathCommon(set *Set, path_elems unsafe.Pointer, path_len int, astext bool, null_handle NullHandleType) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonbset_extract_path_common(set._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), C.bool(astext), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbsetObjectFieldCommon wraps MEOS C function jsonbset_object_field_common.
func JsonbsetObjectFieldCommon(set *Set, key string, astext bool, null_handle NullHandleType) (_r0 *Set, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.jsonbset_object_field_common(set._inner, _c_key, C.bool(astext), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// ContainedJsonbSet wraps MEOS C function contained_jsonb_set.
func ContainedJsonbSet(jb *Jsonb, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_jsonb_set(jb._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetJsonb wraps MEOS C function contains_set_jsonb.
func ContainsSetJsonb(s *Set, jb *Jsonb) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_jsonb(s._inner, jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionJsonbSet wraps MEOS C function intersection_jsonb_set.
func IntersectionJsonbSet(jb *Jsonb, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_jsonb_set(jb._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetJsonb wraps MEOS C function intersection_set_jsonb.
func IntersectionSetJsonb(s *Set, jb *Jsonb) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_jsonb(s._inner, jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// JsonbUnionTransfn wraps MEOS C function jsonb_union_transfn.
func JsonbUnionTransfn(state *Set, jb *Jsonb) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.jsonb_union_transfn(state._inner, jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusJsonbSet wraps MEOS C function minus_jsonb_set.
func MinusJsonbSet(jb *Jsonb, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_jsonb_set(jb._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetJsonb wraps MEOS C function minus_set_jsonb.
func MinusSetJsonb(s *Set, jb *Jsonb) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_jsonb(s._inner, jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionJsonbSet wraps MEOS C function union_jsonb_set.
func UnionJsonbSet(jb *Jsonb, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_jsonb_set(jb._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetJsonb wraps MEOS C function union_set_jsonb.
func UnionSetJsonb(s *Set, jb *Jsonb) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_jsonb(s._inner, jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TjsonbFromMFJSON wraps MEOS C function tjsonb_from_mfjson.
func TjsonbFromMFJSON(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tjsonb_from_mfjson(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbIn wraps MEOS C function tjsonb_in.
func TjsonbIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tjsonb_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbOut wraps MEOS C function tjsonb_out.
func TjsonbOut(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_out(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TjsonbinstIn wraps MEOS C function tjsonbinst_in.
func TjsonbinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tjsonbinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TjsonbseqIn wraps MEOS C function tjsonbseq_in.
func TjsonbseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tjsonbseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TjsonbseqsetIn wraps MEOS C function tjsonbseqset_in.
func TjsonbseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tjsonbseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TjsonbFromBaseTemp wraps MEOS C function tjsonb_from_base_temp.
func TjsonbFromBaseTemp(jsonb *Jsonb, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_from_base_temp(jsonb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbinstMake wraps MEOS C function tjsonbinst_make.
func TjsonbinstMake(jsonb *Jsonb, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonbinst_make(jsonb._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TjsonbseqFromBaseTstzset wraps MEOS C function tjsonbseq_from_base_tstzset.
func TjsonbseqFromBaseTstzset(jsonb *Jsonb, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonbseq_from_base_tstzset(jsonb._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TjsonbseqFromBaseTstzspan wraps MEOS C function tjsonbseq_from_base_tstzspan.
func TjsonbseqFromBaseTstzspan(jsonb *Jsonb, sp *Span) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonbseq_from_base_tstzspan(jsonb._inner, sp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TjsonbseqsetFromBaseTstzspanset wraps MEOS C function tjsonbseqset_from_base_tstzspanset.
func TjsonbseqsetFromBaseTstzspanset(jsonb *Jsonb, ss *SpanSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonbseqset_from_base_tstzspanset(jsonb._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TjsonbToTtext wraps MEOS C function tjsonb_to_ttext.
func TjsonbToTtext(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_to_ttext(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TtextToTjsonb wraps MEOS C function ttext_to_tjsonb.
func TtextToTjsonb(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ttext_to_tjsonb(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbEndValue wraps MEOS C function tjsonb_end_value.
func TjsonbEndValue(temp *Temporal) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// TjsonbStartValue wraps MEOS C function tjsonb_start_value.
func TjsonbStartValue(temp *Temporal) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// TjsonbValue wraps MEOS C function tjsonb_value.
func TjsonbValue(temp *Temporal) (_r0 *Jsonb, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Jsonb{_inner: _cret}, nil
}


// TjsonbValueAtTimestamptz wraps MEOS C function tjsonb_value_at_timestamptz.
func TjsonbValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 *Jsonb, _err error) {
	var _out_value *C.Jsonb
	C.meos_errno_reset()
	_cret := C.tjsonb_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Jsonb{_inner: _out_value}, nil
}


// TjsonbValueN wraps MEOS C function tjsonb_value_n.
func TjsonbValueN(temp *Temporal, n int) (_r0 bool, _r1 *Jsonb, _err error) {
	var _out_result *C.Jsonb
	C.meos_errno_reset()
	_cret := C.tjsonb_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Jsonb{_inner: _out_result}, nil
}


// TjsonbValues wraps MEOS C function tjsonb_values.
func TjsonbValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TjsonbUnnest wraps MEOS C function tjsonb_unnest.
func TjsonbUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_unnest(temp._inner, (***C.Jsonb)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ConcatTjsonbJsonb wraps MEOS C function concat_tjsonb_jsonb.
func ConcatTjsonbJsonb(temp *Temporal, jb *Jsonb, invert bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.concat_tjsonb_jsonb(temp._inner, jb._inner, C.bool(invert))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// ConcatTjsonbTjsonb wraps MEOS C function concat_tjsonb_tjsonb.
func ConcatTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.concat_tjsonb_tjsonb(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// ContainsTjsonbJsonb wraps MEOS C function contains_tjsonb_jsonb.
func ContainsTjsonbJsonb(temp *Temporal, jb *Jsonb, invert bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tjsonb_jsonb(temp._inner, jb._inner, C.bool(invert))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// ContainsTjsonbTjsonb wraps MEOS C function contains_tjsonb_tjsonb.
func ContainsTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tjsonb_tjsonb(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// NullHandleTypeFromString wraps MEOS C function null_handle_type_from_string.
func NullHandleTypeFromString(str string) (_r0 NullHandleType, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.null_handle_type_from_string(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return NullHandleType(_cret), nil
}


// TjsonArrayElement wraps MEOS C function tjson_array_element.
func TjsonArrayElement(temp *Temporal, idx int, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjson_array_element(temp._inner, C.int(idx), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonArrayLength wraps MEOS C function tjson_array_length.
func TjsonArrayLength(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjson_array_length(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonExtractPath wraps MEOS C function tjson_extract_path.
func TjsonExtractPath(temp *Temporal, path_elems unsafe.Pointer, path_len int, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjson_extract_path(temp._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonObjectField wraps MEOS C function tjson_object_field.
func TjsonObjectField(temp *Temporal, key string, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjson_object_field(temp._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonStripNulls wraps MEOS C function tjson_strip_nulls.
func TjsonStripNulls(temp *Temporal, strip_in_arrays bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjson_strip_nulls(temp._inner, C.bool(strip_in_arrays))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbArrayElement wraps MEOS C function tjsonb_array_element.
func TjsonbArrayElement(temp *Temporal, idx int, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_array_element(temp._inner, C.int(idx), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbArrayElementText wraps MEOS C function tjsonb_array_element_text.
func TjsonbArrayElementText(temp *Temporal, idx int, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_array_element_text(temp._inner, C.int(idx), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbArrayLength wraps MEOS C function tjsonb_array_length.
func TjsonbArrayLength(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_array_length(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbDelete wraps MEOS C function tjsonb_delete.
func TjsonbDelete(temp *Temporal, key string) (_r0 *Temporal, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_delete(temp._inner, _c_key)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbDeleteArray wraps MEOS C function tjsonb_delete_array.
func TjsonbDeleteArray(temp *Temporal, keys unsafe.Pointer, count int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_delete_array(temp._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbDeleteIndex wraps MEOS C function tjsonb_delete_index.
func TjsonbDeleteIndex(temp *Temporal, idx int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_delete_index(temp._inner, C.int(idx))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbDeletePath wraps MEOS C function tjsonb_delete_path.
func TjsonbDeletePath(temp *Temporal, path_elems unsafe.Pointer, path_len int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_delete_path(temp._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbExists wraps MEOS C function tjsonb_exists.
func TjsonbExists(temp *Temporal, key string) (_r0 *Temporal, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_exists(temp._inner, _c_key)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbExistsAll wraps MEOS C function tjsonb_exists_all.
func TjsonbExistsAll(temp *Temporal, keys unsafe.Pointer, count int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_exists_all(temp._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbExistsAny wraps MEOS C function tjsonb_exists_any.
func TjsonbExistsAny(temp *Temporal, keys unsafe.Pointer, count int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_exists_any(temp._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbExistsArray wraps MEOS C function tjsonb_exists_array.
func TjsonbExistsArray(temp *Temporal, keys unsafe.Pointer, count int, any bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_exists_array(temp._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count), C.bool(any))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbExtractPath wraps MEOS C function tjsonb_extract_path.
func TjsonbExtractPath(temp *Temporal, path_elems unsafe.Pointer, path_len int, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_extract_path(temp._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbExtractPathText wraps MEOS C function tjsonb_extract_path_text.
func TjsonbExtractPathText(temp *Temporal, path_elems unsafe.Pointer, path_len int, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_extract_path_text(temp._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbInsert wraps MEOS C function tjsonb_insert.
func TjsonbInsert(temp *Temporal, keys unsafe.Pointer, count int, newjb *Jsonb, after bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_insert(temp._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count), newjb._inner, C.bool(after))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbObjectField wraps MEOS C function tjsonb_object_field.
func TjsonbObjectField(temp *Temporal, key string, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_object_field(temp._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbObjectFieldText wraps MEOS C function tjsonb_object_field_text.
func TjsonbObjectFieldText(temp *Temporal, key string, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_object_field_text(temp._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbPathExists wraps MEOS C function tjsonb_path_exists.
func TjsonbPathExists(temp *Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_path_exists(temp._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbPathMatch wraps MEOS C function tjsonb_path_match.
func TjsonbPathMatch(temp *Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_path_match(temp._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbPathQueryArray wraps MEOS C function tjsonb_path_query_array.
func TjsonbPathQueryArray(temp *Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_path_query_array(temp._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbPathQueryFirst wraps MEOS C function tjsonb_path_query_first.
func TjsonbPathQueryFirst(temp *Temporal, jp *JsonPath, vars *Jsonb, silent bool, tz bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_path_query_first(temp._inner, jp._inner, vars._inner, C.bool(silent), C.bool(tz))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbPretty wraps MEOS C function tjsonb_pretty.
func TjsonbPretty(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_pretty(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbSet wraps MEOS C function tjsonb_set.
func TjsonbSet(temp *Temporal, keys unsafe.Pointer, count int, newjb *Jsonb, create bool, handle_null string, lax bool) (_r0 *Temporal, _err error) {
	_c_handle_null := C.cstring_to_text(C.CString(handle_null))
	defer C.free(unsafe.Pointer(_c_handle_null))
	C.meos_errno_reset()
	_cret := C.tjsonb_set(temp._inner, (**C.text)(unsafe.Pointer(keys)), C.int(count), newjb._inner, C.bool(create), _c_handle_null, C.bool(lax))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbStripNulls wraps MEOS C function tjsonb_strip_nulls.
func TjsonbStripNulls(temp *Temporal, strip_in_arrays bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_strip_nulls(temp._inner, C.bool(strip_in_arrays))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbToTbigint wraps MEOS C function tjsonb_to_tbigint.
func TjsonbToTbigint(temp *Temporal, key string, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_to_tbigint(temp._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbToTbool wraps MEOS C function tjsonb_to_tbool.
func TjsonbToTbool(temp *Temporal, key string, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_to_tbool(temp._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbToTfloat wraps MEOS C function tjsonb_to_tfloat.
func TjsonbToTfloat(temp *Temporal, key string, interp Interpolation, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_to_tfloat(temp._inner, _c_key, C.interpType(interp), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbToTgeogpoint wraps MEOS C function tjsonb_to_tgeogpoint.
func TjsonbToTgeogpoint(temp *Temporal, key string, interp Interpolation, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_to_tgeogpoint(temp._inner, _c_key, C.interpType(interp), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbToTgeography wraps MEOS C function tjsonb_to_tgeography.
func TjsonbToTgeography(temp *Temporal, key string, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_to_tgeography(temp._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbToTgeometry wraps MEOS C function tjsonb_to_tgeometry.
func TjsonbToTgeometry(temp *Temporal, key string, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_to_tgeometry(temp._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbToTgeompoint wraps MEOS C function tjsonb_to_tgeompoint.
func TjsonbToTgeompoint(temp *Temporal, key string, interp Interpolation, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_to_tgeompoint(temp._inner, _c_key, C.interpType(interp), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbToTint wraps MEOS C function tjsonb_to_tint.
func TjsonbToTint(temp *Temporal, key string, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_to_tint(temp._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbToTtextKey wraps MEOS C function tjsonb_to_ttext_key.
func TjsonbToTtextKey(temp *Temporal, key string, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.CString(key)
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_to_ttext_key(temp._inner, _c_key, C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonObjectFieldCommon wraps MEOS C function tjson_object_field_common.
func TjsonObjectFieldCommon(temp *Temporal, key string, astext bool, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjson_object_field_common(temp._inner, _c_key, C.bool(astext), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbArrayElementCommon wraps MEOS C function tjsonb_array_element_common.
func TjsonbArrayElementCommon(temp *Temporal, idx int, astext bool, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_array_element_common(temp._inner, C.int(idx), C.bool(astext), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbExtractPathCommon wraps MEOS C function tjsonb_extract_path_common.
func TjsonbExtractPathCommon(temp *Temporal, path_elems unsafe.Pointer, path_len int, astext bool, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_extract_path_common(temp._inner, (**C.text)(unsafe.Pointer(path_elems)), C.int(path_len), C.bool(astext), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbObjectFieldCommon wraps MEOS C function tjsonb_object_field_common.
func TjsonbObjectFieldCommon(temp *Temporal, key string, astext bool, null_handle NullHandleType) (_r0 *Temporal, _err error) {
	_c_key := C.cstring_to_text(C.CString(key))
	defer C.free(unsafe.Pointer(_c_key))
	C.meos_errno_reset()
	_cret := C.tjsonb_object_field_common(temp._inner, _c_key, C.bool(astext), C.nullHandleType(null_handle))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbAtValue wraps MEOS C function tjsonb_at_value.
func TjsonbAtValue(temp *Temporal, jsb *Jsonb) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_at_value(temp._inner, jsb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TjsonbMinusValue wraps MEOS C function tjsonb_minus_value.
func TjsonbMinusValue(temp *Temporal, jsb *Jsonb) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tjsonb_minus_value(temp._inner, jsb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AlwaysEqJsonbTjsonb wraps MEOS C function always_eq_jsonb_tjsonb.
func AlwaysEqJsonbTjsonb(jb *Jsonb, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_jsonb_tjsonb(jb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTjsonbJsonb wraps MEOS C function always_eq_tjsonb_jsonb.
func AlwaysEqTjsonbJsonb(temp *Temporal, jb *Jsonb) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tjsonb_jsonb(temp._inner, jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTjsonbTjsonb wraps MEOS C function always_eq_tjsonb_tjsonb.
func AlwaysEqTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tjsonb_tjsonb(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeJsonbTjsonb wraps MEOS C function always_ne_jsonb_tjsonb.
func AlwaysNeJsonbTjsonb(jb *Jsonb, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_jsonb_tjsonb(jb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTjsonbJsonb wraps MEOS C function always_ne_tjsonb_jsonb.
func AlwaysNeTjsonbJsonb(temp *Temporal, jb *Jsonb) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tjsonb_jsonb(temp._inner, jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTjsonbTjsonb wraps MEOS C function always_ne_tjsonb_tjsonb.
func AlwaysNeTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tjsonb_tjsonb(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqJsonbTjsonb wraps MEOS C function ever_eq_jsonb_tjsonb.
func EverEqJsonbTjsonb(jb *Jsonb, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_jsonb_tjsonb(jb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTjsonbJsonb wraps MEOS C function ever_eq_tjsonb_jsonb.
func EverEqTjsonbJsonb(temp *Temporal, jb *Jsonb) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tjsonb_jsonb(temp._inner, jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTjsonbTjsonb wraps MEOS C function ever_eq_tjsonb_tjsonb.
func EverEqTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tjsonb_tjsonb(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeJsonbTjsonb wraps MEOS C function ever_ne_jsonb_tjsonb.
func EverNeJsonbTjsonb(jb *Jsonb, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_jsonb_tjsonb(jb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTjsonbJsonb wraps MEOS C function ever_ne_tjsonb_jsonb.
func EverNeTjsonbJsonb(temp *Temporal, jb *Jsonb) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tjsonb_jsonb(temp._inner, jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTjsonbTjsonb wraps MEOS C function ever_ne_tjsonb_tjsonb.
func EverNeTjsonbTjsonb(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tjsonb_tjsonb(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqJsonbTjsonb wraps MEOS C function teq_jsonb_tjsonb.
func TeqJsonbTjsonb(jb *Jsonb, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_jsonb_tjsonb(jb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTjsonbJsonb wraps MEOS C function teq_tjsonb_jsonb.
func TeqTjsonbJsonb(temp *Temporal, jb *Jsonb) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tjsonb_jsonb(temp._inner, jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneJsonbTjsonb wraps MEOS C function tne_jsonb_tjsonb.
func TneJsonbTjsonb(jb *Jsonb, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_jsonb_tjsonb(jb._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTjsonbJsonb wraps MEOS C function tne_tjsonb_jsonb.
func TneTjsonbJsonb(temp *Temporal, jb *Jsonb) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tjsonb_jsonb(temp._inner, jb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}

