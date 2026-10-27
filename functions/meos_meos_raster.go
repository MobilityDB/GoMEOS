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

// RaquetPixtypeSize wraps MEOS C function raquet_pixtype_size.
func RaquetPixtypeSize(pixtype MeosPixType) (_r0 uint, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_pixtype_size(C.MeosPixType(pixtype))
	if _err = meosError(); _err != nil {
		return
	}
	return uint(_cret), nil
}


// RaquetPixtypeFromString wraps MEOS C function raquet_pixtype_from_string.
func RaquetPixtypeFromString(str string) (_r0 MeosPixType, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.raquet_pixtype_from_string(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return MeosPixType(_cret), nil
}


// RaquetIn wraps MEOS C function raquet_in.
func RaquetIn(str string) (_r0 unsafe.Pointer, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.raquet_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RaquetOut wraps MEOS C function raquet_out.
func RaquetOut(rq unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_out((*C.Raquet)(unsafe.Pointer(rq)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// RaquetFromWKB wraps MEOS C function raquet_from_wkb.
func RaquetFromWKB(wkb unsafe.Pointer, size uint) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RaquetFromHexwkb wraps MEOS C function raquet_from_hexwkb.
func RaquetFromHexwkb(hexwkb string) (_r0 unsafe.Pointer, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.raquet_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RaquetAsWKB wraps MEOS C function raquet_as_wkb.
func RaquetAsWKB(rq unsafe.Pointer, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_as_wkb((*C.Raquet)(unsafe.Pointer(rq)), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RaquetAsHexwkb wraps MEOS C function raquet_as_hexwkb.
func RaquetAsHexwkb(rq unsafe.Pointer, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_as_hexwkb((*C.Raquet)(unsafe.Pointer(rq)), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// RaquetMake wraps MEOS C function raquet_make.
func RaquetMake(quadbin uint64, width int32, height int32, pixtype MeosPixType, nodata float64, has_nodata bool, pixels unsafe.Pointer, pixels_size uint) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_make(C.uint64_t(quadbin), C.int32(width), C.int32(height), C.MeosPixType(pixtype), C.double(nodata), C.bool(has_nodata), (*C.uint8_t)(unsafe.Pointer(pixels)), C.size_t(pixels_size))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RaquetCopy wraps MEOS C function raquet_copy.
func RaquetCopy(rq unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_copy((*C.Raquet)(unsafe.Pointer(rq)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RaquetRead wraps MEOS C function raquet_read.
func RaquetRead(path string, quadbin uint64) (_r0 unsafe.Pointer, _err error) {
	_c_path := C.CString(path)
	defer C.free(unsafe.Pointer(_c_path))
	C.meos_errno_reset()
	_cret := C.raquet_read(_c_path, C.uint64_t(quadbin))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RaquetReadBytes wraps MEOS C function raquet_read_bytes.
func RaquetReadBytes(data unsafe.Pointer, size uint, quadbin uint64) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_read_bytes((*C.uint8_t)(unsafe.Pointer(data)), C.size_t(size), C.uint64_t(quadbin))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RaquetQuadbin wraps MEOS C function raquet_quadbin.
func RaquetQuadbin(rq unsafe.Pointer) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_quadbin((*C.Raquet)(unsafe.Pointer(rq)))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// RaquetWidth wraps MEOS C function raquet_width.
func RaquetWidth(rq unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_width((*C.Raquet)(unsafe.Pointer(rq)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RaquetHeight wraps MEOS C function raquet_height.
func RaquetHeight(rq unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_height((*C.Raquet)(unsafe.Pointer(rq)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RaquetBandPixelType wraps MEOS C function raquet_band_pixel_type.
func RaquetBandPixelType(rq unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_band_pixel_type((*C.Raquet)(unsafe.Pointer(rq)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// RaquetBandHasNodataValue wraps MEOS C function raquet_band_has_nodata_value.
func RaquetBandHasNodataValue(rq unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_band_has_nodata_value((*C.Raquet)(unsafe.Pointer(rq)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RaquetBandNodataValue wraps MEOS C function raquet_band_nodata_value.
func RaquetBandNodataValue(rq unsafe.Pointer) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.raquet_band_nodata_value((*C.Raquet)(unsafe.Pointer(rq)), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// RaquetPixels wraps MEOS C function raquet_pixels.
func RaquetPixels(rq unsafe.Pointer, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_pixels((*C.Raquet)(unsafe.Pointer(rq)), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RaquetHash wraps MEOS C function raquet_hash.
func RaquetHash(rq unsafe.Pointer) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_hash((*C.Raquet)(unsafe.Pointer(rq)))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// RaquetHashExtended wraps MEOS C function raquet_hash_extended.
func RaquetHashExtended(rq unsafe.Pointer, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_hash_extended((*C.Raquet)(unsafe.Pointer(rq)), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// RaquetMemSize wraps MEOS C function raquet_mem_size.
func RaquetMemSize(rq unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_mem_size((*C.Raquet)(unsafe.Pointer(rq)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RasterFromWKB wraps MEOS C function raster_from_wkb.
func RasterFromWKB(wkb unsafe.Pointer, size uint) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterFromHexwkb wraps MEOS C function raster_from_hexwkb.
func RasterFromHexwkb(hexwkb string) (_r0 unsafe.Pointer, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.raster_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterAsWKB wraps MEOS C function raster_as_wkb.
func RasterAsWKB(rast unsafe.Pointer, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_as_wkb((*C.Raster)(unsafe.Pointer(rast)), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterAsHexwkb wraps MEOS C function raster_as_hexwkb.
func RasterAsHexwkb(rast unsafe.Pointer, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_as_hexwkb((*C.Raster)(unsafe.Pointer(rast)), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// RasterCopy wraps MEOS C function raster_copy.
func RasterCopy(rast unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_copy((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterMake wraps MEOS C function raster_make.
func RasterMake(width int, height int, upperleftx float64, upperlefty float64, scalex float64, scaley float64, skewx float64, skewy float64, srid int32, pixeltype string, nodataval float64, has_nodata bool, pixels unsafe.Pointer, size uint) (_r0 unsafe.Pointer, _err error) {
	_c_pixeltype := C.CString(pixeltype)
	defer C.free(unsafe.Pointer(_c_pixeltype))
	C.meos_errno_reset()
	_cret := C.raster_make(C.int(width), C.int(height), C.double(upperleftx), C.double(upperlefty), C.double(scalex), C.double(scaley), C.double(skewx), C.double(skewy), C.int32_t(srid), _c_pixeltype, C.double(nodataval), C.bool(has_nodata), (*C.uint8_t)(unsafe.Pointer(pixels)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterMakeEmpty wraps MEOS C function raster_make_empty.
func RasterMakeEmpty(width int, height int, upperleftx float64, upperlefty float64, scalex float64, scaley float64, skewx float64, skewy float64, srid int32) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_make_empty(C.int(width), C.int(height), C.double(upperleftx), C.double(upperlefty), C.double(scalex), C.double(scaley), C.double(skewx), C.double(skewy), C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterAddBand wraps MEOS C function raster_add_band.
func RasterAddBand(rast unsafe.Pointer, index int, pixeltype string, initialvalue float64, nodataval float64, has_nodata bool) (_r0 unsafe.Pointer, _err error) {
	_c_pixeltype := C.CString(pixeltype)
	defer C.free(unsafe.Pointer(_c_pixeltype))
	C.meos_errno_reset()
	_cret := C.raster_add_band((*C.Raster)(unsafe.Pointer(rast)), C.int(index), _c_pixeltype, C.double(initialvalue), C.double(nodataval), C.bool(has_nodata))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterAddBandRaster wraps MEOS C function raster_add_band_raster.
func RasterAddBandRaster(torast unsafe.Pointer, fromrast unsafe.Pointer, fromband int, toindex int) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_add_band_raster((*C.Raster)(unsafe.Pointer(torast)), (*C.Raster)(unsafe.Pointer(fromrast)), C.int(fromband), C.int(toindex))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterMemSize wraps MEOS C function raster_mem_size.
func RasterMemSize(rast unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_mem_size((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RasterNumBands wraps MEOS C function raster_num_bands.
func RasterNumBands(rast unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_num_bands((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RasterWidth wraps MEOS C function raster_width.
func RasterWidth(rast unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_width((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RasterHeight wraps MEOS C function raster_height.
func RasterHeight(rast unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_height((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RasterSRID wraps MEOS C function raster_srid.
func RasterSRID(rast unsafe.Pointer) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_srid((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// RasterUpperLeftX wraps MEOS C function raster_upper_left_x.
func RasterUpperLeftX(rast unsafe.Pointer) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_upper_left_x((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// RasterUpperLeftY wraps MEOS C function raster_upper_left_y.
func RasterUpperLeftY(rast unsafe.Pointer) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_upper_left_y((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// RasterScaleX wraps MEOS C function raster_scale_x.
func RasterScaleX(rast unsafe.Pointer) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_scale_x((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// RasterScaleY wraps MEOS C function raster_scale_y.
func RasterScaleY(rast unsafe.Pointer) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_scale_y((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// RasterSkewX wraps MEOS C function raster_skew_x.
func RasterSkewX(rast unsafe.Pointer) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_skew_x((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// RasterSkewY wraps MEOS C function raster_skew_y.
func RasterSkewY(rast unsafe.Pointer) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_skew_y((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// RasterBandPixelType wraps MEOS C function raster_band_pixel_type.
func RasterBandPixelType(rast unsafe.Pointer, band int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_band_pixel_type((*C.Raster)(unsafe.Pointer(rast)), C.int(band))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// RasterBandHasNodataValue wraps MEOS C function raster_band_has_nodata_value.
func RasterBandHasNodataValue(rast unsafe.Pointer, band int) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_band_has_nodata_value((*C.Raster)(unsafe.Pointer(rast)), C.int(band))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RasterBandNodataValue wraps MEOS C function raster_band_nodata_value.
func RasterBandNodataValue(rast unsafe.Pointer, band int) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.raster_band_nodata_value((*C.Raster)(unsafe.Pointer(rast)), C.int(band), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// RasterBandPixels wraps MEOS C function raster_band_pixels.
func RasterBandPixels(rast unsafe.Pointer, band int, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_band_pixels((*C.Raster)(unsafe.Pointer(rast)), C.int(band), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// BandstatsCount wraps MEOS C function bandstats_count.
func BandstatsCount(stats unsafe.Pointer) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.bandstats_count((*C.BandStats)(unsafe.Pointer(stats)))
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// BandstatsSum wraps MEOS C function bandstats_sum.
func BandstatsSum(stats unsafe.Pointer) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.bandstats_sum((*C.BandStats)(unsafe.Pointer(stats)), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// BandstatsMean wraps MEOS C function bandstats_mean.
func BandstatsMean(stats unsafe.Pointer) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.bandstats_mean((*C.BandStats)(unsafe.Pointer(stats)), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// BandstatsStddev wraps MEOS C function bandstats_stddev.
func BandstatsStddev(stats unsafe.Pointer) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.bandstats_stddev((*C.BandStats)(unsafe.Pointer(stats)), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// BandstatsMin wraps MEOS C function bandstats_min.
func BandstatsMin(stats unsafe.Pointer) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.bandstats_min((*C.BandStats)(unsafe.Pointer(stats)), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// BandstatsMax wraps MEOS C function bandstats_max.
func BandstatsMax(stats unsafe.Pointer) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.bandstats_max((*C.BandStats)(unsafe.Pointer(stats)), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// RasterClip wraps MEOS C function raster_clip.
func RasterClip(rast unsafe.Pointer, bands unsafe.Pointer, nbands int, gs *Geom, nodata unsafe.Pointer, nnodata int, crop bool, touched bool) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_clip((*C.Raster)(unsafe.Pointer(rast)), (*C.int)(unsafe.Pointer(bands)), C.int(nbands), gs._inner, (*C.double)(unsafe.Pointer(nodata)), C.int(nnodata), C.bool(crop), C.bool(touched))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterTransform wraps MEOS C function raster_transform.
func RasterTransform(rast unsafe.Pointer, srid int32, algorithm *string, max_err float64, scale_x float64, scale_y float64) (_r0 unsafe.Pointer, _err error) {
	var _c_algorithm *C.char
	if algorithm != nil {
		_c_algorithm = C.CString(*algorithm)
		defer C.free(unsafe.Pointer(_c_algorithm))
	}
	C.meos_errno_reset()
	_cret := C.raster_transform((*C.Raster)(unsafe.Pointer(rast)), C.int32_t(srid), _c_algorithm, C.double(max_err), C.double(scale_x), C.double(scale_y))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterTransformPipeline wraps MEOS C function raster_transform_pipeline.
func RasterTransformPipeline(rast unsafe.Pointer, pipelinestr string, srid int32, is_forward bool, algorithm *string, max_err float64, scale_x float64, scale_y float64) (_r0 unsafe.Pointer, _err error) {
	_c_pipelinestr := C.CString(pipelinestr)
	defer C.free(unsafe.Pointer(_c_pipelinestr))
	var _c_algorithm *C.char
	if algorithm != nil {
		_c_algorithm = C.CString(*algorithm)
		defer C.free(unsafe.Pointer(_c_algorithm))
	}
	C.meos_errno_reset()
	_cret := C.raster_transform_pipeline((*C.Raster)(unsafe.Pointer(rast)), _c_pipelinestr, C.int32_t(srid), C.bool(is_forward), _c_algorithm, C.double(max_err), C.double(scale_x), C.double(scale_y))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterTransformRaster wraps MEOS C function raster_transform_raster.
func RasterTransformRaster(rast unsafe.Pointer, alignto unsafe.Pointer, algorithm *string, max_err float64) (_r0 unsafe.Pointer, _err error) {
	var _c_algorithm *C.char
	if algorithm != nil {
		_c_algorithm = C.CString(*algorithm)
		defer C.free(unsafe.Pointer(_c_algorithm))
	}
	C.meos_errno_reset()
	_cret := C.raster_transform_raster((*C.Raster)(unsafe.Pointer(rast)), (*C.Raster)(unsafe.Pointer(alignto)), _c_algorithm, C.double(max_err))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterRescale wraps MEOS C function raster_rescale.
func RasterRescale(rast unsafe.Pointer, scale_x float64, scale_y float64, algorithm *string, max_err float64) (_r0 unsafe.Pointer, _err error) {
	var _c_algorithm *C.char
	if algorithm != nil {
		_c_algorithm = C.CString(*algorithm)
		defer C.free(unsafe.Pointer(_c_algorithm))
	}
	C.meos_errno_reset()
	_cret := C.raster_rescale((*C.Raster)(unsafe.Pointer(rast)), C.double(scale_x), C.double(scale_y), _c_algorithm, C.double(max_err))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterDumpAsPolygons wraps MEOS C function raster_dump_as_polygons.
func RasterDumpAsPolygons(rast unsafe.Pointer, band int, exclude_nodata bool, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_dump_as_polygons((*C.Raster)(unsafe.Pointer(rast)), C.int(band), C.bool(exclude_nodata), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterDumpValues wraps MEOS C function raster_dump_values.
func RasterDumpValues(rast unsafe.Pointer, band int, exclude_nodata bool, nodata unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_dump_values((*C.Raster)(unsafe.Pointer(rast)), C.int(band), C.bool(exclude_nodata), (**C.bool)(unsafe.Pointer(nodata)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterPixelAsPoints wraps MEOS C function raster_pixel_as_points.
func RasterPixelAsPoints(rast unsafe.Pointer, band int, exclude_nodata bool, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_pixel_as_points((*C.Raster)(unsafe.Pointer(rast)), C.int(band), C.bool(exclude_nodata), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterSummaryStats wraps MEOS C function raster_summary_stats.
func RasterSummaryStats(rast unsafe.Pointer, band int, exclude_nodata bool) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_summary_stats((*C.Raster)(unsafe.Pointer(rast)), C.int(band), C.bool(exclude_nodata))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterValueCount wraps MEOS C function raster_value_count.
func RasterValueCount(rast unsafe.Pointer, band int, exclude_nodata bool, searchvalues unsafe.Pointer, nsearch int, roundto float64, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_value_count((*C.Raster)(unsafe.Pointer(rast)), C.int(band), C.bool(exclude_nodata), (*C.double)(unsafe.Pointer(searchvalues)), C.int(nsearch), C.double(roundto), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterSetValues wraps MEOS C function raster_set_values.
func RasterSetValues(rast unsafe.Pointer, band int, x int, y int, width int, height int, values unsafe.Pointer, nodata unsafe.Pointer, noset unsafe.Pointer, count int, keepnodata bool) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_set_values((*C.Raster)(unsafe.Pointer(rast)), C.int(band), C.int(x), C.int(y), C.int(width), C.int(height), (*C.double)(unsafe.Pointer(values)), (*C.bool)(unsafe.Pointer(nodata)), (*C.bool)(unsafe.Pointer(noset)), C.int(count), C.bool(keepnodata))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterReclass wraps MEOS C function raster_reclass.
func RasterReclass(rast unsafe.Pointer, band int, expr string, pixeltype string, has_nodata bool, nodataval float64) (_r0 unsafe.Pointer, _err error) {
	_c_expr := C.CString(expr)
	defer C.free(unsafe.Pointer(_c_expr))
	_c_pixeltype := C.CString(pixeltype)
	defer C.free(unsafe.Pointer(_c_pixeltype))
	C.meos_errno_reset()
	_cret := C.raster_reclass((*C.Raster)(unsafe.Pointer(rast)), C.int(band), _c_expr, _c_pixeltype, C.bool(has_nodata), C.double(nodataval))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterSlope wraps MEOS C function raster_slope.
func RasterSlope(rast unsafe.Pointer, band int, customextent unsafe.Pointer, pixeltype string, units string, scale float64, interpolate_nodata bool) (_r0 unsafe.Pointer, _err error) {
	_c_pixeltype := C.CString(pixeltype)
	defer C.free(unsafe.Pointer(_c_pixeltype))
	_c_units := C.CString(units)
	defer C.free(unsafe.Pointer(_c_units))
	C.meos_errno_reset()
	_cret := C.raster_slope((*C.Raster)(unsafe.Pointer(rast)), C.int(band), (*C.Raster)(unsafe.Pointer(customextent)), _c_pixeltype, _c_units, C.double(scale), C.bool(interpolate_nodata))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterTpi wraps MEOS C function raster_tpi.
func RasterTpi(rast unsafe.Pointer, band int, customextent unsafe.Pointer, pixeltype string, interpolate_nodata bool) (_r0 unsafe.Pointer, _err error) {
	_c_pixeltype := C.CString(pixeltype)
	defer C.free(unsafe.Pointer(_c_pixeltype))
	C.meos_errno_reset()
	_cret := C.raster_tpi((*C.Raster)(unsafe.Pointer(rast)), C.int(band), (*C.Raster)(unsafe.Pointer(customextent)), _c_pixeltype, C.bool(interpolate_nodata))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterTri wraps MEOS C function raster_tri.
func RasterTri(rast unsafe.Pointer, band int, customextent unsafe.Pointer, pixeltype string, interpolate_nodata bool) (_r0 unsafe.Pointer, _err error) {
	_c_pixeltype := C.CString(pixeltype)
	defer C.free(unsafe.Pointer(_c_pixeltype))
	C.meos_errno_reset()
	_cret := C.raster_tri((*C.Raster)(unsafe.Pointer(rast)), C.int(band), (*C.Raster)(unsafe.Pointer(customextent)), _c_pixeltype, C.bool(interpolate_nodata))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterMapAlgebra wraps MEOS C function raster_map_algebra.
func RasterMapAlgebra(rast unsafe.Pointer, band int, stat string, pixeltype *string, extenttype *string, customextent unsafe.Pointer, distancex int32, distancey int32) (_r0 unsafe.Pointer, _err error) {
	_c_stat := C.CString(stat)
	defer C.free(unsafe.Pointer(_c_stat))
	var _c_pixeltype *C.char
	if pixeltype != nil {
		_c_pixeltype = C.CString(*pixeltype)
		defer C.free(unsafe.Pointer(_c_pixeltype))
	}
	var _c_extenttype *C.char
	if extenttype != nil {
		_c_extenttype = C.CString(*extenttype)
		defer C.free(unsafe.Pointer(_c_extenttype))
	}
	C.meos_errno_reset()
	_cret := C.raster_map_algebra((*C.Raster)(unsafe.Pointer(rast)), C.int(band), _c_stat, _c_pixeltype, _c_extenttype, (*C.Raster)(unsafe.Pointer(customextent)), C.int32(distancex), C.int32(distancey))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// GeomvalArrFree wraps MEOS C function geomval_arr_free.
func GeomvalArrFree(gvarr unsafe.Pointer, count int) (_err error) {
	C.meos_errno_reset()
	C.geomval_arr_free((*C.GeomVal)(unsafe.Pointer(gvarr)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// PixelArrFree wraps MEOS C function pixel_arr_free.
func PixelArrFree(pixarr unsafe.Pointer, count int) (_err error) {
	C.meos_errno_reset()
	C.pixel_arr_free((*C.Pixel)(unsafe.Pointer(pixarr)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// AddFloatRaster wraps MEOS C function add_float_raster.
func AddFloatRaster(d float64, rast unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.add_float_raster(C.double(d), (*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// AddRasterFloat wraps MEOS C function add_raster_float.
func AddRasterFloat(rast unsafe.Pointer, d float64) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.add_raster_float((*C.Raster)(unsafe.Pointer(rast)), C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// AddRasterRaster wraps MEOS C function add_raster_raster.
func AddRasterRaster(rast1 unsafe.Pointer, rast2 unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.add_raster_raster((*C.Raster)(unsafe.Pointer(rast1)), (*C.Raster)(unsafe.Pointer(rast2)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// DivFloatRaster wraps MEOS C function div_float_raster.
func DivFloatRaster(d float64, rast unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.div_float_raster(C.double(d), (*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// DivRasterFloat wraps MEOS C function div_raster_float.
func DivRasterFloat(rast unsafe.Pointer, d float64) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.div_raster_float((*C.Raster)(unsafe.Pointer(rast)), C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// DivRasterRaster wraps MEOS C function div_raster_raster.
func DivRasterRaster(rast1 unsafe.Pointer, rast2 unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.div_raster_raster((*C.Raster)(unsafe.Pointer(rast1)), (*C.Raster)(unsafe.Pointer(rast2)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// MulFloatRaster wraps MEOS C function mul_float_raster.
func MulFloatRaster(d float64, rast unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_float_raster(C.double(d), (*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// MulRasterFloat wraps MEOS C function mul_raster_float.
func MulRasterFloat(rast unsafe.Pointer, d float64) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_raster_float((*C.Raster)(unsafe.Pointer(rast)), C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// MulRasterRaster wraps MEOS C function mul_raster_raster.
func MulRasterRaster(rast1 unsafe.Pointer, rast2 unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.mul_raster_raster((*C.Raster)(unsafe.Pointer(rast1)), (*C.Raster)(unsafe.Pointer(rast2)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SubFloatRaster wraps MEOS C function sub_float_raster.
func SubFloatRaster(d float64, rast unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sub_float_raster(C.double(d), (*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SubRasterFloat wraps MEOS C function sub_raster_float.
func SubRasterFloat(rast unsafe.Pointer, d float64) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sub_raster_float((*C.Raster)(unsafe.Pointer(rast)), C.double(d))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// SubRasterRaster wraps MEOS C function sub_raster_raster.
func SubRasterRaster(rast1 unsafe.Pointer, rast2 unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.sub_raster_raster((*C.Raster)(unsafe.Pointer(rast1)), (*C.Raster)(unsafe.Pointer(rast2)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterSameAlignment wraps MEOS C function raster_same_alignment.
func RasterSameAlignment(rast1 unsafe.Pointer, rast2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_same_alignment((*C.Raster)(unsafe.Pointer(rast1)), (*C.Raster)(unsafe.Pointer(rast2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RasterNotSameAlignmentReason wraps MEOS C function raster_not_same_alignment_reason.
func RasterNotSameAlignmentReason(rast1 unsafe.Pointer, rast2 unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_not_same_alignment_reason((*C.Raster)(unsafe.Pointer(rast1)), (*C.Raster)(unsafe.Pointer(rast2)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// RasterSameAlignmentCombinefn wraps MEOS C function raster_same_alignment_combinefn.
func RasterSameAlignmentCombinefn(state1 unsafe.Pointer, state2 unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_same_alignment_combinefn((*C.RasterAlignState)(unsafe.Pointer(state1)), (*C.RasterAlignState)(unsafe.Pointer(state2)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterSameAlignmentFinalfn wraps MEOS C function raster_same_alignment_finalfn.
func RasterSameAlignmentFinalfn(state unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_same_alignment_finalfn((*C.RasterAlignState)(unsafe.Pointer(state)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RasterSameAlignmentTransfn wraps MEOS C function raster_same_alignment_transfn.
func RasterSameAlignmentTransfn(state unsafe.Pointer, rast unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_same_alignment_transfn((*C.RasterAlignState)(unsafe.Pointer(state)), (*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterSummaryStatsCombinefn wraps MEOS C function raster_summary_stats_combinefn.
func RasterSummaryStatsCombinefn(state1 unsafe.Pointer, state2 unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_summary_stats_combinefn((*C.RasterStatsState)(unsafe.Pointer(state1)), (*C.RasterStatsState)(unsafe.Pointer(state2)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterSummaryStatsFinalfn wraps MEOS C function raster_summary_stats_finalfn.
func RasterSummaryStatsFinalfn(state unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_summary_stats_finalfn((*C.RasterStatsState)(unsafe.Pointer(state)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterSummaryStatsTransfn wraps MEOS C function raster_summary_stats_transfn.
func RasterSummaryStatsTransfn(state unsafe.Pointer, rast unsafe.Pointer, band int, exclude_nodata bool) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_summary_stats_transfn((*C.RasterStatsState)(unsafe.Pointer(state)), (*C.Raster)(unsafe.Pointer(rast)), C.int(band), C.bool(exclude_nodata))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterUnionCombinefn wraps MEOS C function raster_union_combinefn.
func RasterUnionCombinefn(state1 unsafe.Pointer, state2 unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_union_combinefn((*C.RasterUnionState)(unsafe.Pointer(state1)), (*C.RasterUnionState)(unsafe.Pointer(state2)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterUnionFinalfn wraps MEOS C function raster_union_finalfn.
func RasterUnionFinalfn(state unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_union_finalfn((*C.RasterUnionState)(unsafe.Pointer(state)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterUnionTransfn wraps MEOS C function raster_union_transfn.
func RasterUnionTransfn(state unsafe.Pointer, rast unsafe.Pointer, band int, uniontype *string) (_r0 unsafe.Pointer, _err error) {
	var _c_uniontype *C.char
	if uniontype != nil {
		_c_uniontype = C.CString(*uniontype)
		defer C.free(unsafe.Pointer(_c_uniontype))
	}
	C.meos_errno_reset()
	_cret := C.raster_union_transfn((*C.RasterUnionState)(unsafe.Pointer(state)), (*C.Raster)(unsafe.Pointer(rast)), C.int(band), _c_uniontype)
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasteralignstateDeserialize wraps MEOS C function rasteralignstate_deserialize.
func RasteralignstateDeserialize(bytes unsafe.Pointer, size uint) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.rasteralignstate_deserialize((*C.uint8_t)(unsafe.Pointer(bytes)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasteralignstateSerialize wraps MEOS C function rasteralignstate_serialize.
func RasteralignstateSerialize(state unsafe.Pointer, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.rasteralignstate_serialize((*C.RasterAlignState)(unsafe.Pointer(state)), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterstatsstateDeserialize wraps MEOS C function rasterstatsstate_deserialize.
func RasterstatsstateDeserialize(bytes unsafe.Pointer, size uint) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.rasterstatsstate_deserialize((*C.uint8_t)(unsafe.Pointer(bytes)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterstatsstateSerialize wraps MEOS C function rasterstatsstate_serialize.
func RasterstatsstateSerialize(state unsafe.Pointer, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.rasterstatsstate_serialize((*C.RasterStatsState)(unsafe.Pointer(state)), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterunionstateDeserialize wraps MEOS C function rasterunionstate_deserialize.
func RasterunionstateDeserialize(bytes unsafe.Pointer, size uint) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.rasterunionstate_deserialize((*C.uint8_t)(unsafe.Pointer(bytes)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RasterunionstateSerialize wraps MEOS C function rasterunionstate_serialize.
func RasterunionstateSerialize(state unsafe.Pointer, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.rasterunionstate_serialize((*C.RasterUnionState)(unsafe.Pointer(state)), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// RaquetToSTBOX wraps MEOS C function raquet_to_stbox.
func RaquetToSTBOX(rq unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_to_stbox((*C.Raquet)(unsafe.Pointer(rq)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// RasterToSTBOX wraps MEOS C function raster_to_stbox.
func RasterToSTBOX(rast unsafe.Pointer) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_to_stbox((*C.Raster)(unsafe.Pointer(rast)))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// RaquetCmp wraps MEOS C function raquet_cmp.
func RaquetCmp(rq1 unsafe.Pointer, rq2 unsafe.Pointer) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_cmp((*C.Raquet)(unsafe.Pointer(rq1)), (*C.Raquet)(unsafe.Pointer(rq2)))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RaquetEq wraps MEOS C function raquet_eq.
func RaquetEq(rq1 unsafe.Pointer, rq2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_eq((*C.Raquet)(unsafe.Pointer(rq1)), (*C.Raquet)(unsafe.Pointer(rq2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RaquetNe wraps MEOS C function raquet_ne.
func RaquetNe(rq1 unsafe.Pointer, rq2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_ne((*C.Raquet)(unsafe.Pointer(rq1)), (*C.Raquet)(unsafe.Pointer(rq2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RaquetLt wraps MEOS C function raquet_lt.
func RaquetLt(rq1 unsafe.Pointer, rq2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_lt((*C.Raquet)(unsafe.Pointer(rq1)), (*C.Raquet)(unsafe.Pointer(rq2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RaquetLe wraps MEOS C function raquet_le.
func RaquetLe(rq1 unsafe.Pointer, rq2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_le((*C.Raquet)(unsafe.Pointer(rq1)), (*C.Raquet)(unsafe.Pointer(rq2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RaquetGe wraps MEOS C function raquet_ge.
func RaquetGe(rq1 unsafe.Pointer, rq2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_ge((*C.Raquet)(unsafe.Pointer(rq1)), (*C.Raquet)(unsafe.Pointer(rq2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RaquetGt wraps MEOS C function raquet_gt.
func RaquetGt(rq1 unsafe.Pointer, rq2 unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.raquet_gt((*C.Raquet)(unsafe.Pointer(rq1)), (*C.Raquet)(unsafe.Pointer(rq2)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RasterValue wraps MEOS C function raster_value.
func RasterValue(traj *Temporal, rast unsafe.Pointer, band int, exclude_nodata bool, resample *string) (_r0 *Temporal, _err error) {
	var _c_resample *C.char
	if resample != nil {
		_c_resample = C.CString(*resample)
		defer C.free(unsafe.Pointer(_c_resample))
	}
	C.meos_errno_reset()
	_cret := C.raster_value(traj._inner, (*C.Raster)(unsafe.Pointer(rast)), C.int(band), C.bool(exclude_nodata), _c_resample)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// RasterAtValue wraps MEOS C function raster_at_value.
func RasterAtValue(traj *Temporal, rast unsafe.Pointer, band int, vspan *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_at_value(traj._inner, (*C.Raster)(unsafe.Pointer(rast)), C.int(band), vspan._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// RasterMinusValue wraps MEOS C function raster_minus_value.
func RasterMinusValue(traj *Temporal, rast unsafe.Pointer, band int, vspan *Span) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_minus_value(traj._inner, (*C.Raster)(unsafe.Pointer(rast)), C.int(band), vspan._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// ErasterValue wraps MEOS C function eraster_value.
func ErasterValue(traj *Temporal, rast unsafe.Pointer, band int, vspan *Span) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.eraster_value(traj._inner, (*C.Raster)(unsafe.Pointer(rast)), C.int(band), vspan._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// ArasterValue wraps MEOS C function araster_value.
func ArasterValue(traj *Temporal, rast unsafe.Pointer, band int, vspan *Span) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.araster_value(traj._inner, (*C.Raster)(unsafe.Pointer(rast)), C.int(band), vspan._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RasterValueGdal wraps MEOS C function raster_value_gdal.
func RasterValueGdal(traj *Temporal, path string, band int) (_r0 *Temporal, _err error) {
	_c_path := C.CString(path)
	defer C.free(unsafe.Pointer(_c_path))
	C.meos_errno_reset()
	_cret := C.raster_value_gdal(traj._inner, _c_path, C.int(band))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// RasterAtValueGdal wraps MEOS C function raster_at_value_gdal.
func RasterAtValueGdal(traj *Temporal, path string, band int, vspan *Span) (_r0 *Temporal, _err error) {
	_c_path := C.CString(path)
	defer C.free(unsafe.Pointer(_c_path))
	C.meos_errno_reset()
	_cret := C.raster_at_value_gdal(traj._inner, _c_path, C.int(band), vspan._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// RasterMinusValueGdal wraps MEOS C function raster_minus_value_gdal.
func RasterMinusValueGdal(traj *Temporal, path string, band int, vspan *Span) (_r0 *Temporal, _err error) {
	_c_path := C.CString(path)
	defer C.free(unsafe.Pointer(_c_path))
	C.meos_errno_reset()
	_cret := C.raster_minus_value_gdal(traj._inner, _c_path, C.int(band), vspan._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// ErasterValueGdal wraps MEOS C function eraster_value_gdal.
func ErasterValueGdal(traj *Temporal, path string, band int, vspan *Span) (_r0 int, _err error) {
	_c_path := C.CString(path)
	defer C.free(unsafe.Pointer(_c_path))
	C.meos_errno_reset()
	_cret := C.eraster_value_gdal(traj._inner, _c_path, C.int(band), vspan._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// ArasterValueGdal wraps MEOS C function araster_value_gdal.
func ArasterValueGdal(traj *Temporal, path string, band int, vspan *Span) (_r0 int, _err error) {
	_c_path := C.CString(path)
	defer C.free(unsafe.Pointer(_c_path))
	C.meos_errno_reset()
	_cret := C.araster_value_gdal(traj._inner, _c_path, C.int(band), vspan._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// RasterTileValueQuadbin wraps MEOS C function raster_tile_value_quadbin.
func RasterTileValueQuadbin(traj *Temporal, pixels unsafe.Pointer, pixels_size uint, width int32, height int32, quadbin uint64, pixtype MeosPixType, nodata float64, has_nodata bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_tile_value_quadbin(traj._inner, (*C.uint8_t)(unsafe.Pointer(pixels)), C.size_t(pixels_size), C.int32(width), C.int32(height), C.uint64_t(quadbin), C.MeosPixType(pixtype), C.double(nodata), C.bool(has_nodata))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// RasterTileValue wraps MEOS C function raster_tile_value.
func RasterTileValue(traj *Temporal, rq unsafe.Pointer) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_tile_value(traj._inner, (*C.Raquet)(unsafe.Pointer(rq)))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// RasterTileValueArray wraps MEOS C function raster_tile_value_array.
func RasterTileValueArray(traj *Temporal, rqarr unsafe.Pointer, count int) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.raster_tile_value_array(traj._inner, (**C.Raquet)(unsafe.Pointer(rqarr)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}

