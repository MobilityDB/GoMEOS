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

// QuadbinAsHexwkb wraps MEOS C function quadbin_as_hexwkb.
func QuadbinAsHexwkb(cell uint64, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_as_hexwkb(C.uint64_t(cell), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// QuadbinAsWKB wraps MEOS C function quadbin_as_wkb.
func QuadbinAsWKB(cell uint64, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_as_wkb(C.uint64_t(cell), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// QuadbinFromHexwkb wraps MEOS C function quadbin_from_hexwkb.
func QuadbinFromHexwkb(hexwkb string) (_r0 uint64, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.quadbin_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// QuadbinFromWKB wraps MEOS C function quadbin_from_wkb.
func QuadbinFromWKB(wkb unsafe.Pointer, size uint) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// QuadbinIn wraps MEOS C function quadbin_in.
func QuadbinIn(str string) (_r0 uint64, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.quadbin_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// BigintToQuadbin wraps MEOS C function bigint_to_quadbin.
func BigintToQuadbin(i int64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_to_quadbin(C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// QuadbinOut wraps MEOS C function quadbin_out.
func QuadbinOut(cell uint64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_out(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// QuadbinEq wraps MEOS C function quadbin_eq.
func QuadbinEq(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_eq(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// QuadbinNe wraps MEOS C function quadbin_ne.
func QuadbinNe(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_ne(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// QuadbinLt wraps MEOS C function quadbin_lt.
func QuadbinLt(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_lt(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// QuadbinLe wraps MEOS C function quadbin_le.
func QuadbinLe(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_le(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// QuadbinGt wraps MEOS C function quadbin_gt.
func QuadbinGt(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_gt(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// QuadbinGe wraps MEOS C function quadbin_ge.
func QuadbinGe(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_ge(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// QuadbinCmp wraps MEOS C function quadbin_cmp.
func QuadbinCmp(a uint64, b uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_cmp(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// QuadbinHash wraps MEOS C function quadbin_hash.
func QuadbinHash(cell uint64) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_hash(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// QuadbinHashExtended wraps MEOS C function quadbin_hash_extended.
func QuadbinHashExtended(cell uint64, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_hash_extended(C.uint64_t(cell), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// QuadbinIsValidIndex wraps MEOS C function quadbin_is_valid_index.
func QuadbinIsValidIndex(index uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_is_valid_index(C.uint64_t(index))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// QuadbinIsValidCell wraps MEOS C function quadbin_is_valid_cell.
func QuadbinIsValidCell(cell uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_is_valid_cell(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// QuadbinTileToCell wraps MEOS C function quadbin_tile_to_cell.
func QuadbinTileToCell(x uint32, y uint32, z uint32) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_tile_to_cell(C.uint32_t(x), C.uint32_t(y), C.uint32_t(z))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// QuadbinCellToTile wraps MEOS C function quadbin_cell_to_tile.
func QuadbinCellToTile(cell uint64, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_cell_to_tile(C.uint64_t(cell), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// QuadbinGetResolution wraps MEOS C function quadbin_get_resolution.
func QuadbinGetResolution(cell uint64) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_get_resolution(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// QuadbinCellToParent wraps MEOS C function quadbin_cell_to_parent.
func QuadbinCellToParent(cell uint64, parent_resolution uint32) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_cell_to_parent(C.uint64_t(cell), C.uint32_t(parent_resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// QuadbinCellToChildren wraps MEOS C function quadbin_cell_to_children.
func QuadbinCellToChildren(cell uint64, children_resolution uint32, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_cell_to_children(C.uint64_t(cell), C.uint32_t(children_resolution), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// QuadbinCellSibling wraps MEOS C function quadbin_cell_sibling.
func QuadbinCellSibling(cell uint64, direction string) (_r0 uint64, _err error) {
	_c_direction := C.CString(direction)
	defer C.free(unsafe.Pointer(_c_direction))
	C.meos_errno_reset()
	_cret := C.quadbin_cell_sibling(C.uint64_t(cell), _c_direction)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// QuadbinKRing wraps MEOS C function quadbin_k_ring.
func QuadbinKRing(cell uint64, k int, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_k_ring(C.uint64_t(cell), C.int(k), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// QuadbinPointToCell wraps MEOS C function quadbin_point_to_cell.
func QuadbinPointToCell(longitude float64, latitude float64, resolution uint32) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_point_to_cell(C.double(longitude), C.double(latitude), C.uint32_t(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// GeoToQuadbinCell wraps MEOS C function geo_to_quadbin_cell.
func GeoToQuadbinCell(point *Geom, resolution int32) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_to_quadbin_cell(point._inner, C.int32(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// GeoToQuadbinSet wraps MEOS C function geo_to_quadbin_set.
func GeoToQuadbinSet(gs *Geom, resolution int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_to_quadbin_set(gs._inner, C.int32(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// QuadbinCellToGeompoint wraps MEOS C function quadbin_cell_to_geompoint.
func QuadbinCellToGeompoint(cell uint64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_cell_to_geompoint(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// QuadbinCellToGeom wraps MEOS C function quadbin_cell_to_geom.
func QuadbinCellToGeom(cell uint64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_cell_to_geom(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// EverEqQuadbinsetTquadbin wraps MEOS C function ever_eq_quadbinset_tquadbin.
func EverEqQuadbinsetTquadbin(cells *Set, tqb *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_quadbinset_tquadbin(cells._inner, tqb._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTquadbinQuadbinset wraps MEOS C function ever_eq_tquadbin_quadbinset.
func EverEqTquadbinQuadbinset(tqb *Temporal, cells *Set) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tquadbin_quadbinset(tqb._inner, cells._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// QuadbinToSTBOX wraps MEOS C function quadbin_to_stbox.
func QuadbinToSTBOX(cell uint64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_to_stbox(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// QuadbinTimestamptzToSTBOX wraps MEOS C function quadbin_timestamptz_to_stbox.
func QuadbinTimestamptzToSTBOX(cell uint64, t int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_timestamptz_to_stbox(C.uint64_t(cell), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// QuadbinTstzspanToSTBOX wraps MEOS C function quadbin_tstzspan_to_stbox.
func QuadbinTstzspanToSTBOX(cell uint64, s *Span) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_tstzspan_to_stbox(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// QuadbinCellArea wraps MEOS C function quadbin_cell_area.
func QuadbinCellArea(cell uint64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_cell_area(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// QuadbinIndexToString wraps MEOS C function quadbin_index_to_string.
func QuadbinIndexToString(index uint64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_index_to_string(C.uint64_t(index))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// QuadbinStringToIndex wraps MEOS C function quadbin_string_to_index.
func QuadbinStringToIndex(str string) (_r0 uint64, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.quadbin_string_to_index(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// QuadbinCellToQuadkey wraps MEOS C function quadbin_cell_to_quadkey.
func QuadbinCellToQuadkey(cell uint64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_cell_to_quadkey(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// QuadbinGridDisk wraps MEOS C function quadbin_grid_disk.
func QuadbinGridDisk(origin uint64, k int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_grid_disk(C.uint64_t(origin), C.int(k))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// QuadbinsetGridDisk wraps MEOS C function quadbinset_grid_disk.
func QuadbinsetGridDisk(cells *Set, k int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbinset_grid_disk(cells._inner, C.int(k))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// QuadbinCellToChildrenSet wraps MEOS C function quadbin_cell_to_children_set.
func QuadbinCellToChildrenSet(origin uint64, children_resolution int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_cell_to_children_set(C.uint64_t(origin), C.int(children_resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// QuadbinsetCompactCells wraps MEOS C function quadbinset_compact_cells.
func QuadbinsetCompactCells(cells *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbinset_compact_cells(cells._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// QuadbinsetUncompactCells wraps MEOS C function quadbinset_uncompact_cells.
func QuadbinsetUncompactCells(cells *Set, resolution int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbinset_uncompact_cells(cells._inner, C.int(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// QuadbinCellToCover wraps MEOS C function quadbin_cell_to_cover.
func QuadbinCellToCover(cell uint64, resolution int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_cell_to_cover(C.uint64_t(cell), C.int(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// QuadbinsetCoverCells wraps MEOS C function quadbinset_cover_cells.
func QuadbinsetCoverCells(cells *Set, resolution int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbinset_cover_cells(cells._inner, C.int(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// QuadbinsetIn wraps MEOS C function quadbinset_in.
func QuadbinsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.quadbinset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// QuadbinsetOut wraps MEOS C function quadbinset_out.
func QuadbinsetOut(s *Set) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbinset_out(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// QuadbinsetMake wraps MEOS C function quadbinset_make.
func QuadbinsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbinset_make((*C.uint64_t)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// QuadbinsetEndValue wraps MEOS C function quadbinset_end_value.
func QuadbinsetEndValue(s *Set) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbinset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// QuadbinsetStartValue wraps MEOS C function quadbinset_start_value.
func QuadbinsetStartValue(s *Set) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbinset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// QuadbinsetValueN wraps MEOS C function quadbinset_value_n.
func QuadbinsetValueN(s *Set, n int) (_r0 bool, _r1 uint64, _err error) {
	var _out_result C.uint64_t
	C.meos_errno_reset()
	_cret := C.quadbinset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), uint64(_out_result), nil
}


// QuadbinsetValues wraps MEOS C function quadbinset_values.
func QuadbinsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbinset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ContainedQuadbinSet wraps MEOS C function contained_quadbin_set.
func ContainedQuadbinSet(cell uint64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_quadbin_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetQuadbin wraps MEOS C function contains_set_quadbin.
func ContainsSetQuadbin(s *Set, cell uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_quadbin(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionQuadbinSet wraps MEOS C function intersection_quadbin_set.
func IntersectionQuadbinSet(cell uint64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_quadbin_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetQuadbin wraps MEOS C function intersection_set_quadbin.
func IntersectionSetQuadbin(s *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_quadbin(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusQuadbinSet wraps MEOS C function minus_quadbin_set.
func MinusQuadbinSet(cell uint64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_quadbin_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetQuadbin wraps MEOS C function minus_set_quadbin.
func MinusSetQuadbin(s *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_quadbin(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionQuadbinSet wraps MEOS C function union_quadbin_set.
func UnionQuadbinSet(cell uint64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_quadbin_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetQuadbin wraps MEOS C function union_set_quadbin.
func UnionSetQuadbin(s *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_quadbin(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// QuadbinUnionTransfn wraps MEOS C function quadbin_union_transfn.
func QuadbinUnionTransfn(state *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_union_transfn(state._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TquadbinIn wraps MEOS C function tquadbin_in.
func TquadbinIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tquadbin_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TquadbininstIn wraps MEOS C function tquadbininst_in.
func TquadbininstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tquadbininst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TquadbinseqIn wraps MEOS C function tquadbinseq_in.
func TquadbinseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tquadbinseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TquadbinseqsetIn wraps MEOS C function tquadbinseqset_in.
func TquadbinseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tquadbinseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TquadbinOut wraps MEOS C function tquadbin_out.
func TquadbinOut(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_out(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TquadbinFromMFJSON wraps MEOS C function tquadbin_from_mfjson.
func TquadbinFromMFJSON(mfjson string) (_r0 *Temporal, _err error) {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	C.meos_errno_reset()
	_cret := C.tquadbin_from_mfjson(_c_mfjson)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TquadbinMake wraps MEOS C function tquadbin_make.
func TquadbinMake(value uint64, t int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_make(C.uint64_t(value), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TquadbininstMake wraps MEOS C function tquadbininst_make.
func TquadbininstMake(value uint64, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbininst_make(C.uint64_t(value), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TquadbinFromBaseTemp wraps MEOS C function tquadbin_from_base_temp.
func TquadbinFromBaseTemp(value uint64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_from_base_temp(C.uint64_t(value), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TquadbinseqFromBaseTstzset wraps MEOS C function tquadbinseq_from_base_tstzset.
func TquadbinseqFromBaseTstzset(value uint64, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbinseq_from_base_tstzset(C.uint64_t(value), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TquadbinseqFromBaseTstzspan wraps MEOS C function tquadbinseq_from_base_tstzspan.
func TquadbinseqFromBaseTstzspan(value uint64, s *Span) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbinseq_from_base_tstzspan(C.uint64_t(value), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TquadbinseqsetFromBaseTstzspanset wraps MEOS C function tquadbinseqset_from_base_tstzspanset.
func TquadbinseqsetFromBaseTstzspanset(value uint64, ss *SpanSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbinseqset_from_base_tstzspanset(C.uint64_t(value), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TquadbinseqMake wraps MEOS C function tquadbinseq_make.
func TquadbinseqMake(values unsafe.Pointer, times unsafe.Pointer, count int, lower_inc bool, upper_inc bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbinseq_make((*C.uint64_t)(unsafe.Pointer(values)), (*C.TimestampTz)(unsafe.Pointer(times)), C.int(count), C.bool(lower_inc), C.bool(upper_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TquadbinseqsetMake wraps MEOS C function tquadbinseqset_make.
func TquadbinseqsetMake(sequences unsafe.Pointer, count int) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbinseqset_make((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TquadbinStartValue wraps MEOS C function tquadbin_start_value.
func TquadbinStartValue(temp *Temporal) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TquadbinEndValue wraps MEOS C function tquadbin_end_value.
func TquadbinEndValue(temp *Temporal) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TquadbinValue wraps MEOS C function tquadbin_value.
func TquadbinValue(temp *Temporal) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// TquadbinValueN wraps MEOS C function tquadbin_value_n.
func TquadbinValueN(temp *Temporal, n int) (_r0 bool, _r1 uint64, _err error) {
	var _out_result C.uint64_t
	C.meos_errno_reset()
	_cret := C.tquadbin_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), uint64(_out_result), nil
}


// TquadbinValues wraps MEOS C function tquadbin_values.
func TquadbinValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TquadbinUnnest wraps MEOS C function tquadbin_unnest.
func TquadbinUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_unnest(temp._inner, (**C.uint64_t)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TquadbinValueAtTimestamptz wraps MEOS C function tquadbin_value_at_timestamptz.
func TquadbinValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 uint64, _err error) {
	var _out_result C.uint64_t
	C.meos_errno_reset()
	_cret := C.tquadbin_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), uint64(_out_result), nil
}


// QuadbinToSet wraps MEOS C function quadbin_to_set.
func QuadbinToSet(cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.quadbin_to_set(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TbigintToTquadbin wraps MEOS C function tbigint_to_tquadbin.
func TbigintToTquadbin(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_to_tquadbin(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TquadbinToTbigint wraps MEOS C function tquadbin_to_tbigint.
func TquadbinToTbigint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_to_tbigint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TquadbinAtValue wraps MEOS C function tquadbin_at_value.
func TquadbinAtValue(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_at_value(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TquadbinMinusValue wraps MEOS C function tquadbin_minus_value.
func TquadbinMinusValue(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_minus_value(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// EverEqQuadbinTquadbin wraps MEOS C function ever_eq_quadbin_tquadbin.
func EverEqQuadbinTquadbin(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTquadbinQuadbin wraps MEOS C function ever_eq_tquadbin_quadbin.
func EverEqTquadbinQuadbin(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeQuadbinTquadbin wraps MEOS C function ever_ne_quadbin_tquadbin.
func EverNeQuadbinTquadbin(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTquadbinQuadbin wraps MEOS C function ever_ne_tquadbin_quadbin.
func EverNeTquadbinQuadbin(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqQuadbinTquadbin wraps MEOS C function always_eq_quadbin_tquadbin.
func AlwaysEqQuadbinTquadbin(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTquadbinQuadbin wraps MEOS C function always_eq_tquadbin_quadbin.
func AlwaysEqTquadbinQuadbin(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeQuadbinTquadbin wraps MEOS C function always_ne_quadbin_tquadbin.
func AlwaysNeQuadbinTquadbin(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTquadbinQuadbin wraps MEOS C function always_ne_tquadbin_quadbin.
func AlwaysNeTquadbinQuadbin(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTquadbinTquadbin wraps MEOS C function ever_eq_tquadbin_tquadbin.
func EverEqTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tquadbin_tquadbin(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTquadbinTquadbin wraps MEOS C function ever_ne_tquadbin_tquadbin.
func EverNeTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tquadbin_tquadbin(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTquadbinTquadbin wraps MEOS C function always_eq_tquadbin_tquadbin.
func AlwaysEqTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tquadbin_tquadbin(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTquadbinTquadbin wraps MEOS C function always_ne_tquadbin_tquadbin.
func AlwaysNeTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tquadbin_tquadbin(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqQuadbinTquadbin wraps MEOS C function teq_quadbin_tquadbin.
func TeqQuadbinTquadbin(cell uint64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTquadbinQuadbin wraps MEOS C function teq_tquadbin_quadbin.
func TeqTquadbinQuadbin(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTquadbinTquadbin wraps MEOS C function teq_tquadbin_tquadbin.
func TeqTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tquadbin_tquadbin(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneQuadbinTquadbin wraps MEOS C function tne_quadbin_tquadbin.
func TneQuadbinTquadbin(cell uint64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_quadbin_tquadbin(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTquadbinQuadbin wraps MEOS C function tne_tquadbin_quadbin.
func TneTquadbinQuadbin(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tquadbin_quadbin(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTquadbinTquadbin wraps MEOS C function tne_tquadbin_tquadbin.
func TneTquadbinTquadbin(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tquadbin_tquadbin(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeogpointToTquadbin wraps MEOS C function tgeogpoint_to_tquadbin.
func TgeogpointToTquadbin(temp *Temporal, resolution int32, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeogpoint_to_tquadbin(temp._inner, C.int32(resolution), C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeompointToTquadbin wraps MEOS C function tgeompoint_to_tquadbin.
func TgeompointToTquadbin(temp *Temporal, resolution int32, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeompoint_to_tquadbin(temp._inner, C.int32(resolution), C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeogpointQuadbinSplit wraps MEOS C function tgeogpoint_quadbin_split.
func TgeogpointQuadbinSplit(temp *Temporal, resolution int32, border_inc bool, cells unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeogpoint_quadbin_split(temp._inner, C.int32(resolution), C.bool(border_inc), (**C.uint64_t)(unsafe.Pointer(cells)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TgeompointQuadbinSplit wraps MEOS C function tgeompoint_quadbin_split.
func TgeompointQuadbinSplit(temp *Temporal, resolution int32, border_inc bool, cells unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeompoint_quadbin_split(temp._inner, C.int32(resolution), C.bool(border_inc), (**C.uint64_t)(unsafe.Pointer(cells)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TquadbinCellToQuadkey wraps MEOS C function tquadbin_cell_to_quadkey.
func TquadbinCellToQuadkey(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tquadbin_cell_to_quadkey(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}

