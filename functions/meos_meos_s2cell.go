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

// S2cellAsHexwkb wraps MEOS C function s2cell_as_hexwkb.
func S2cellAsHexwkb(cell uint64, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_as_hexwkb(C.uint64_t(cell), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// S2cellAsWKB wraps MEOS C function s2cell_as_wkb.
func S2cellAsWKB(cell uint64, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_as_wkb(C.uint64_t(cell), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// S2cellFromHexwkb wraps MEOS C function s2cell_from_hexwkb.
func S2cellFromHexwkb(hexwkb string) (_r0 uint64, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.s2cell_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellFromWKB wraps MEOS C function s2cell_from_wkb.
func S2cellFromWKB(wkb unsafe.Pointer, size uint) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellIn wraps MEOS C function s2cell_in.
func S2cellIn(str string) (_r0 uint64, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.s2cell_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// BigintToS2cell wraps MEOS C function bigint_to_s2cell.
func BigintToS2cell(i int64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_to_s2cell(C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellOut wraps MEOS C function s2cell_out.
func S2cellOut(cell uint64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_out(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// S2cellEq wraps MEOS C function s2cell_eq.
func S2cellEq(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_eq(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// S2cellNe wraps MEOS C function s2cell_ne.
func S2cellNe(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_ne(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// S2cellLt wraps MEOS C function s2cell_lt.
func S2cellLt(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_lt(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// S2cellLe wraps MEOS C function s2cell_le.
func S2cellLe(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_le(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// S2cellGt wraps MEOS C function s2cell_gt.
func S2cellGt(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_gt(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// S2cellGe wraps MEOS C function s2cell_ge.
func S2cellGe(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_ge(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// S2cellCmp wraps MEOS C function s2cell_cmp.
func S2cellCmp(a uint64, b uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_cmp(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// S2cellHash wraps MEOS C function s2cell_hash.
func S2cellHash(cell uint64) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_hash(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// S2cellHashExtended wraps MEOS C function s2cell_hash_extended.
func S2cellHashExtended(cell uint64, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_hash_extended(C.uint64_t(cell), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellIsValidCell wraps MEOS C function s2cell_is_valid_cell.
func S2cellIsValidCell(cell uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_is_valid_cell(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// S2cellCellToToken wraps MEOS C function s2cell_cell_to_token.
func S2cellCellToToken(cell uint64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_cell_to_token(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// S2cellTokenToCell wraps MEOS C function s2cell_token_to_cell.
func S2cellTokenToCell(token string) (_r0 uint64, _err error) {
	_c_token := C.CString(token)
	defer C.free(unsafe.Pointer(_c_token))
	C.meos_errno_reset()
	_cret := C.s2cell_token_to_cell(_c_token)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellGetResolution wraps MEOS C function s2cell_get_resolution.
func S2cellGetResolution(cell uint64) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_get_resolution(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// S2cellGetFace wraps MEOS C function s2cell_get_face.
func S2cellGetFace(cell uint64) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_get_face(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// S2cellCellToParent wraps MEOS C function s2cell_cell_to_parent.
func S2cellCellToParent(cell uint64, level uint32) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_cell_to_parent(C.uint64_t(cell), C.uint32_t(level))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellCellToChild wraps MEOS C function s2cell_cell_to_child.
func S2cellCellToChild(cell uint64, level uint32, position uint32) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_cell_to_child(C.uint64_t(cell), C.uint32_t(level), C.uint32_t(position))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellCellToChildren wraps MEOS C function s2cell_cell_to_children.
func S2cellCellToChildren(cell uint64, level uint32, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_cell_to_children(C.uint64_t(cell), C.uint32_t(level), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// S2cellCellContains wraps MEOS C function s2cell_cell_contains.
func S2cellCellContains(cell uint64, other uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_cell_contains(C.uint64_t(cell), C.uint64_t(other))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// S2cellCommonAncestorLevel wraps MEOS C function s2cell_common_ancestor_level.
func S2cellCommonAncestorLevel(a uint64, b uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_common_ancestor_level(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// S2cellRangeMin wraps MEOS C function s2cell_range_min.
func S2cellRangeMin(cell uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_range_min(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellRangeMax wraps MEOS C function s2cell_range_max.
func S2cellRangeMax(cell uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_range_max(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellEdgeNeighbors wraps MEOS C function s2cell_edge_neighbors.
func S2cellEdgeNeighbors(cell uint64, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_edge_neighbors(C.uint64_t(cell), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// S2cellPointToCell wraps MEOS C function s2cell_point_to_cell.
func S2cellPointToCell(longitude float64, latitude float64, level uint32) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_point_to_cell(C.double(longitude), C.double(latitude), C.uint32_t(level))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellCellArea wraps MEOS C function s2cell_cell_area.
func S2cellCellArea(cell uint64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_cell_area(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// S2cellEdgeLength wraps MEOS C function s2cell_edge_length.
func S2cellEdgeLength(cell uint64, edge uint32) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_edge_length(C.uint64_t(cell), C.uint32_t(edge))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// GeoToS2cellCell wraps MEOS C function geo_to_s2cell_cell.
func GeoToS2cellCell(point *Geom, level int32) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_to_s2cell_cell(point._inner, C.int32(level))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// GeoToS2cellSet wraps MEOS C function geo_to_s2cell_set.
func GeoToS2cellSet(gs *Geom, level int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_to_s2cell_set(gs._inner, C.int32(level))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// S2cellCellToGeogpoint wraps MEOS C function s2cell_cell_to_geogpoint.
func S2cellCellToGeogpoint(cell uint64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_cell_to_geogpoint(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// S2cellCellToGeog wraps MEOS C function s2cell_cell_to_geog.
func S2cellCellToGeog(cell uint64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_cell_to_geog(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// EverEqS2cellsetTs2cell wraps MEOS C function ever_eq_s2cellset_ts2cell.
func EverEqS2cellsetTs2cell(cells *Set, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_s2cellset_ts2cell(cells._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTs2cellS2cellset wraps MEOS C function ever_eq_ts2cell_s2cellset.
func EverEqTs2cellS2cellset(temp *Temporal, cells *Set) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_ts2cell_s2cellset(temp._inner, cells._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// S2cellToSTBOX wraps MEOS C function s2cell_to_stbox.
func S2cellToSTBOX(cell uint64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_to_stbox(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// S2cellTimestamptzToSTBOX wraps MEOS C function s2cell_timestamptz_to_stbox.
func S2cellTimestamptzToSTBOX(cell uint64, t int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_timestamptz_to_stbox(C.uint64_t(cell), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// S2cellTstzspanToSTBOX wraps MEOS C function s2cell_tstzspan_to_stbox.
func S2cellTstzspanToSTBOX(cell uint64, s *Span) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_tstzspan_to_stbox(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// S2cellEdgeNeighborsSet wraps MEOS C function s2cell_edge_neighbors_set.
func S2cellEdgeNeighborsSet(cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_edge_neighbors_set(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// S2cellCellToChildrenSet wraps MEOS C function s2cell_cell_to_children_set.
func S2cellCellToChildrenSet(cell uint64, children_level int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_cell_to_children_set(C.uint64_t(cell), C.int(children_level))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// S2cellsetCompactCells wraps MEOS C function s2cellset_compact_cells.
func S2cellsetCompactCells(cells *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cellset_compact_cells(cells._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// S2cellsetUncompactCells wraps MEOS C function s2cellset_uncompact_cells.
func S2cellsetUncompactCells(cells *Set, resolution int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cellset_uncompact_cells(cells._inner, C.int(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// S2cellCellToCover wraps MEOS C function s2cell_cell_to_cover.
func S2cellCellToCover(cell uint64, resolution int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_cell_to_cover(C.uint64_t(cell), C.int(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// S2cellsetCoverCells wraps MEOS C function s2cellset_cover_cells.
func S2cellsetCoverCells(cells *Set, resolution int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cellset_cover_cells(cells._inner, C.int(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// S2cellsetIn wraps MEOS C function s2cellset_in.
func S2cellsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.s2cellset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// S2cellsetOut wraps MEOS C function s2cellset_out.
func S2cellsetOut(s *Set) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cellset_out(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// S2cellsetMake wraps MEOS C function s2cellset_make.
func S2cellsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cellset_make((*C.uint64_t)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// S2cellsetEndValue wraps MEOS C function s2cellset_end_value.
func S2cellsetEndValue(s *Set) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cellset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellsetStartValue wraps MEOS C function s2cellset_start_value.
func S2cellsetStartValue(s *Set) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cellset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// S2cellsetValueN wraps MEOS C function s2cellset_value_n.
func S2cellsetValueN(s *Set, n int) (_r0 bool, _r1 uint64, _err error) {
	var _out_result C.uint64_t
	C.meos_errno_reset()
	_cret := C.s2cellset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), uint64(_out_result), nil
}


// S2cellsetValues wraps MEOS C function s2cellset_values.
func S2cellsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cellset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ContainedS2cellSet wraps MEOS C function contained_s2cell_set.
func ContainedS2cellSet(cell uint64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_s2cell_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetS2cell wraps MEOS C function contains_set_s2cell.
func ContainsSetS2cell(s *Set, cell uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_s2cell(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionS2cellSet wraps MEOS C function intersection_s2cell_set.
func IntersectionS2cellSet(cell uint64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_s2cell_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetS2cell wraps MEOS C function intersection_set_s2cell.
func IntersectionSetS2cell(s *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_s2cell(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusS2cellSet wraps MEOS C function minus_s2cell_set.
func MinusS2cellSet(cell uint64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_s2cell_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetS2cell wraps MEOS C function minus_set_s2cell.
func MinusSetS2cell(s *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_s2cell(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionS2cellSet wraps MEOS C function union_s2cell_set.
func UnionS2cellSet(cell uint64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_s2cell_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetS2cell wraps MEOS C function union_set_s2cell.
func UnionSetS2cell(s *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_s2cell(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// S2cellUnionTransfn wraps MEOS C function s2cell_union_transfn.
func S2cellUnionTransfn(state *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_union_transfn(state._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// Ts2cellIn wraps MEOS C function ts2cell_in.
func Ts2cellIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.ts2cell_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Ts2cellinstIn wraps MEOS C function ts2cellinst_in.
func Ts2cellinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.ts2cellinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// Ts2cellseqIn wraps MEOS C function ts2cellseq_in.
func Ts2cellseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.ts2cellseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// Ts2cellseqsetIn wraps MEOS C function ts2cellseqset_in.
func Ts2cellseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.ts2cellseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// Ts2cellOut wraps MEOS C function ts2cell_out.
func Ts2cellOut(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_out(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// Ts2cellFromMFJSON wraps MEOS C function ts2cell_from_mfjson.
func Ts2cellFromMFJSON(mfjson string) (_r0 *Temporal, _err error) {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	C.meos_errno_reset()
	_cret := C.ts2cell_from_mfjson(_c_mfjson)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Ts2cellMake wraps MEOS C function ts2cell_make.
func Ts2cellMake(value uint64, t int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_make(C.uint64_t(value), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Ts2cellinstMake wraps MEOS C function ts2cellinst_make.
func Ts2cellinstMake(value uint64, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cellinst_make(C.uint64_t(value), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// Ts2cellFromBaseTemp wraps MEOS C function ts2cell_from_base_temp.
func Ts2cellFromBaseTemp(value uint64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_from_base_temp(C.uint64_t(value), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Ts2cellseqFromBaseTstzset wraps MEOS C function ts2cellseq_from_base_tstzset.
func Ts2cellseqFromBaseTstzset(value uint64, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cellseq_from_base_tstzset(C.uint64_t(value), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// Ts2cellseqFromBaseTstzspan wraps MEOS C function ts2cellseq_from_base_tstzspan.
func Ts2cellseqFromBaseTstzspan(value uint64, s *Span) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cellseq_from_base_tstzspan(C.uint64_t(value), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// Ts2cellseqsetFromBaseTstzspanset wraps MEOS C function ts2cellseqset_from_base_tstzspanset.
func Ts2cellseqsetFromBaseTstzspanset(value uint64, ss *SpanSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cellseqset_from_base_tstzspanset(C.uint64_t(value), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// Ts2cellseqMake wraps MEOS C function ts2cellseq_make.
func Ts2cellseqMake(values unsafe.Pointer, times unsafe.Pointer, count int, lower_inc bool, upper_inc bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cellseq_make((*C.uint64_t)(unsafe.Pointer(values)), (*C.TimestampTz)(unsafe.Pointer(times)), C.int(count), C.bool(lower_inc), C.bool(upper_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// Ts2cellseqsetMake wraps MEOS C function ts2cellseqset_make.
func Ts2cellseqsetMake(sequences unsafe.Pointer, count int) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cellseqset_make((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// Ts2cellStartValue wraps MEOS C function ts2cell_start_value.
func Ts2cellStartValue(temp *Temporal) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// Ts2cellEndValue wraps MEOS C function ts2cell_end_value.
func Ts2cellEndValue(temp *Temporal) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// Ts2cellValue wraps MEOS C function ts2cell_value.
func Ts2cellValue(temp *Temporal) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// Ts2cellValueN wraps MEOS C function ts2cell_value_n.
func Ts2cellValueN(temp *Temporal, n int) (_r0 bool, _r1 uint64, _err error) {
	var _out_result C.uint64_t
	C.meos_errno_reset()
	_cret := C.ts2cell_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), uint64(_out_result), nil
}


// Ts2cellValues wraps MEOS C function ts2cell_values.
func Ts2cellValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// Ts2cellUnnest wraps MEOS C function ts2cell_unnest.
func Ts2cellUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_unnest(temp._inner, (**C.uint64_t)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// Ts2cellValueAtTimestamptz wraps MEOS C function ts2cell_value_at_timestamptz.
func Ts2cellValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 uint64, _err error) {
	var _out_result C.uint64_t
	C.meos_errno_reset()
	_cret := C.ts2cell_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), uint64(_out_result), nil
}


// S2cellToSet wraps MEOS C function s2cell_to_set.
func S2cellToSet(cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.s2cell_to_set(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TbigintToTs2cell wraps MEOS C function tbigint_to_ts2cell.
func TbigintToTs2cell(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_to_ts2cell(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Ts2cellToTbigint wraps MEOS C function ts2cell_to_tbigint.
func Ts2cellToTbigint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_to_tbigint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeogpointToTs2cell wraps MEOS C function tgeogpoint_to_ts2cell.
func TgeogpointToTs2cell(temp *Temporal, level int32, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeogpoint_to_ts2cell(temp._inner, C.int32(level), C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeompointToTs2cell wraps MEOS C function tgeompoint_to_ts2cell.
func TgeompointToTs2cell(temp *Temporal, level int32, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeompoint_to_ts2cell(temp._inner, C.int32(level), C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeogpointS2cellSplit wraps MEOS C function tgeogpoint_s2cell_split.
func TgeogpointS2cellSplit(temp *Temporal, level int32, border_inc bool, cells unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeogpoint_s2cell_split(temp._inner, C.int32(level), C.bool(border_inc), (**C.uint64_t)(unsafe.Pointer(cells)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TgeompointS2cellSplit wraps MEOS C function tgeompoint_s2cell_split.
func TgeompointS2cellSplit(temp *Temporal, level int32, border_inc bool, cells unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeompoint_s2cell_split(temp._inner, C.int32(level), C.bool(border_inc), (**C.uint64_t)(unsafe.Pointer(cells)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// Ts2cellAtValue wraps MEOS C function ts2cell_at_value.
func Ts2cellAtValue(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_at_value(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Ts2cellMinusValue wraps MEOS C function ts2cell_minus_value.
func Ts2cellMinusValue(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_minus_value(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// EverEqS2cellTs2cell wraps MEOS C function ever_eq_s2cell_ts2cell.
func EverEqS2cellTs2cell(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_s2cell_ts2cell(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTs2cellS2cell wraps MEOS C function ever_eq_ts2cell_s2cell.
func EverEqTs2cellS2cell(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_ts2cell_s2cell(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeS2cellTs2cell wraps MEOS C function ever_ne_s2cell_ts2cell.
func EverNeS2cellTs2cell(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_s2cell_ts2cell(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTs2cellS2cell wraps MEOS C function ever_ne_ts2cell_s2cell.
func EverNeTs2cellS2cell(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_ts2cell_s2cell(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqS2cellTs2cell wraps MEOS C function always_eq_s2cell_ts2cell.
func AlwaysEqS2cellTs2cell(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_s2cell_ts2cell(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTs2cellS2cell wraps MEOS C function always_eq_ts2cell_s2cell.
func AlwaysEqTs2cellS2cell(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_ts2cell_s2cell(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeS2cellTs2cell wraps MEOS C function always_ne_s2cell_ts2cell.
func AlwaysNeS2cellTs2cell(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_s2cell_ts2cell(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTs2cellS2cell wraps MEOS C function always_ne_ts2cell_s2cell.
func AlwaysNeTs2cellS2cell(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_ts2cell_s2cell(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTs2cellTs2cell wraps MEOS C function ever_eq_ts2cell_ts2cell.
func EverEqTs2cellTs2cell(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_ts2cell_ts2cell(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTs2cellTs2cell wraps MEOS C function ever_ne_ts2cell_ts2cell.
func EverNeTs2cellTs2cell(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_ts2cell_ts2cell(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTs2cellTs2cell wraps MEOS C function always_eq_ts2cell_ts2cell.
func AlwaysEqTs2cellTs2cell(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_ts2cell_ts2cell(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTs2cellTs2cell wraps MEOS C function always_ne_ts2cell_ts2cell.
func AlwaysNeTs2cellTs2cell(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_ts2cell_ts2cell(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqS2cellTs2cell wraps MEOS C function teq_s2cell_ts2cell.
func TeqS2cellTs2cell(cell uint64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_s2cell_ts2cell(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTs2cellS2cell wraps MEOS C function teq_ts2cell_s2cell.
func TeqTs2cellS2cell(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_ts2cell_s2cell(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTs2cellTs2cell wraps MEOS C function teq_ts2cell_ts2cell.
func TeqTs2cellTs2cell(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_ts2cell_ts2cell(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneS2cellTs2cell wraps MEOS C function tne_s2cell_ts2cell.
func TneS2cellTs2cell(cell uint64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_s2cell_ts2cell(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTs2cellS2cell wraps MEOS C function tne_ts2cell_s2cell.
func TneTs2cellS2cell(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_ts2cell_s2cell(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTs2cellTs2cell wraps MEOS C function tne_ts2cell_ts2cell.
func TneTs2cellTs2cell(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_ts2cell_ts2cell(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Ts2cellCellToToken wraps MEOS C function ts2cell_cell_to_token.
func Ts2cellCellToToken(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.ts2cell_cell_to_token(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}

