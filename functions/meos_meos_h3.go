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

// H3indexIn wraps MEOS C function h3index_in.
func H3indexIn(str string) (_r0 uint64, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.h3index_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// BigintToH3index wraps MEOS C function bigint_to_h3index.
func BigintToH3index(i int64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.bigint_to_h3index(C.int64_t(i))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// H3indexOut wraps MEOS C function h3index_out.
func H3indexOut(cell uint64) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_out(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// H3IsValidCell wraps MEOS C function h3_is_valid_cell.
func H3IsValidCell(cell uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_is_valid_cell(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// H3IsValidDirectedEdge wraps MEOS C function h3_is_valid_directed_edge.
func H3IsValidDirectedEdge(edge uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_is_valid_directed_edge(C.uint64_t(edge))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// H3IsValidVertex wraps MEOS C function h3_is_valid_vertex.
func H3IsValidVertex(vertex uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_is_valid_vertex(C.uint64_t(vertex))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// H3indexEq wraps MEOS C function h3index_eq.
func H3indexEq(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_eq(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// H3indexNe wraps MEOS C function h3index_ne.
func H3indexNe(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_ne(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// H3indexLt wraps MEOS C function h3index_lt.
func H3indexLt(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_lt(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// H3indexLe wraps MEOS C function h3index_le.
func H3indexLe(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_le(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// H3indexGt wraps MEOS C function h3index_gt.
func H3indexGt(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_gt(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// H3indexGe wraps MEOS C function h3index_ge.
func H3indexGe(a uint64, b uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_ge(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// H3indexCmp wraps MEOS C function h3index_cmp.
func H3indexCmp(a uint64, b uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_cmp(C.uint64_t(a), C.uint64_t(b))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// H3indexHash wraps MEOS C function h3index_hash.
func H3indexHash(cell uint64) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_hash(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// H3indexHashExtended wraps MEOS C function h3index_hash_extended.
func H3indexHashExtended(cell uint64, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_hash_extended(C.uint64_t(cell), C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// H3GridDisk wraps MEOS C function h3_grid_disk.
func H3GridDisk(origin uint64, k int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_grid_disk(C.uint64_t(origin), C.int(k))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3CellToChildren wraps MEOS C function h3_cell_to_children.
func H3CellToChildren(origin uint64, childRes int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_cell_to_children(C.uint64_t(origin), C.int(childRes))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3CompactCells wraps MEOS C function h3_compact_cells.
func H3CompactCells(cells *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_compact_cells(cells._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3UncompactCells wraps MEOS C function h3_uncompact_cells.
func H3UncompactCells(cells *Set, res int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_uncompact_cells(cells._inner, C.int(res))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3GridRing wraps MEOS C function h3_grid_ring.
func H3GridRing(origin uint64, k int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_grid_ring(C.uint64_t(origin), C.int(k))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3GridPathCells wraps MEOS C function h3_grid_path_cells.
func H3GridPathCells(start uint64, end uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_grid_path_cells(C.uint64_t(start), C.uint64_t(end))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3OriginToDirectedEdges wraps MEOS C function h3_origin_to_directed_edges.
func H3OriginToDirectedEdges(origin uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_origin_to_directed_edges(C.uint64_t(origin))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3CellToVertexes wraps MEOS C function h3_cell_to_vertexes.
func H3CellToVertexes(cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_cell_to_vertexes(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3GetIcosahedronFaces wraps MEOS C function h3_get_icosahedron_faces.
func H3GetIcosahedronFaces(cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3_get_icosahedron_faces(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3indexFromWKB wraps MEOS C function h3index_from_wkb.
func H3indexFromWKB(wkb unsafe.Pointer, size uint) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// H3indexFromHexwkb wraps MEOS C function h3index_from_hexwkb.
func H3indexFromHexwkb(hexwkb string) (_r0 uint64, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.h3index_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// H3indexAsWKB wraps MEOS C function h3index_as_wkb.
func H3indexAsWKB(cell uint64, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_as_wkb(C.uint64_t(cell), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// H3indexAsHexwkb wraps MEOS C function h3index_as_hexwkb.
func H3indexAsHexwkb(cell uint64, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_as_hexwkb(C.uint64_t(cell), C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// H3indexGetResolution wraps MEOS C function h3index_get_resolution.
func H3indexGetResolution(cell uint64) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_get_resolution(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// H3indexCellToParent wraps MEOS C function h3index_cell_to_parent.
func H3indexCellToParent(cell uint64, parent_resolution uint32) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_cell_to_parent(C.uint64_t(cell), C.uint32_t(parent_resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// H3indexCellToPoint wraps MEOS C function h3index_cell_to_point.
func H3indexCellToPoint(cell uint64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_cell_to_point(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// H3indexCellToBoundary wraps MEOS C function h3index_cell_to_boundary.
func H3indexCellToBoundary(cell uint64) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_cell_to_boundary(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// H3indexCellArea wraps MEOS C function h3index_cell_area.
func H3indexCellArea(cell uint64) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_cell_area(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}


// Th3indexIn wraps MEOS C function th3index_in.
func Th3indexIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.th3index_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexinstIn wraps MEOS C function th3indexinst_in.
func Th3indexinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.th3indexinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// Th3indexseqIn wraps MEOS C function th3indexseq_in.
func Th3indexseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.th3indexseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// Th3indexseqsetIn wraps MEOS C function th3indexseqset_in.
func Th3indexseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.th3indexseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// Th3indexOut wraps MEOS C function th3index_out.
func Th3indexOut(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_out(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// Th3indexFromMFJSON wraps MEOS C function th3index_from_mfjson.
func Th3indexFromMFJSON(mfjson string) (_r0 *Temporal, _err error) {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	C.meos_errno_reset()
	_cret := C.th3index_from_mfjson(_c_mfjson)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexMake wraps MEOS C function th3index_make.
func Th3indexMake(value uint64, t int64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_make(C.uint64_t(value), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexinstMake wraps MEOS C function th3indexinst_make.
func Th3indexinstMake(value uint64, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.th3indexinst_make(C.uint64_t(value), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// Th3indexFromBaseTemp wraps MEOS C function th3index_from_base_temp.
func Th3indexFromBaseTemp(value uint64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_from_base_temp(C.uint64_t(value), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexseqFromBaseTstzset wraps MEOS C function th3indexseq_from_base_tstzset.
func Th3indexseqFromBaseTstzset(value uint64, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.th3indexseq_from_base_tstzset(C.uint64_t(value), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// Th3indexseqFromBaseTstzspan wraps MEOS C function th3indexseq_from_base_tstzspan.
func Th3indexseqFromBaseTstzspan(value uint64, s *Span) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.th3indexseq_from_base_tstzspan(C.uint64_t(value), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// Th3indexseqsetFromBaseTstzspanset wraps MEOS C function th3indexseqset_from_base_tstzspanset.
func Th3indexseqsetFromBaseTstzspanset(value uint64, ss *SpanSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.th3indexseqset_from_base_tstzspanset(C.uint64_t(value), ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// Th3indexseqMake wraps MEOS C function th3indexseq_make.
func Th3indexseqMake(values unsafe.Pointer, times unsafe.Pointer, count int, lower_inc bool, upper_inc bool) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.th3indexseq_make((*C.uint64_t)(unsafe.Pointer(values)), (*C.TimestampTz)(unsafe.Pointer(times)), C.int(count), C.bool(lower_inc), C.bool(upper_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// Th3indexseqsetMake wraps MEOS C function th3indexseqset_make.
func Th3indexseqsetMake(sequences unsafe.Pointer, count int) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.th3indexseqset_make((**C.TSequence)(unsafe.Pointer(sequences)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// Th3indexStartValue wraps MEOS C function th3index_start_value.
func Th3indexStartValue(temp *Temporal) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// Th3indexEndValue wraps MEOS C function th3index_end_value.
func Th3indexEndValue(temp *Temporal) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// Th3indexValue wraps MEOS C function th3index_value.
func Th3indexValue(temp *Temporal) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// Th3indexValueN wraps MEOS C function th3index_value_n.
func Th3indexValueN(temp *Temporal, n int) (_r0 bool, _r1 uint64, _err error) {
	var _out_result C.uint64_t
	C.meos_errno_reset()
	_cret := C.th3index_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), uint64(_out_result), nil
}


// Th3indexValues wraps MEOS C function th3index_values.
func Th3indexValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// Th3indexUnnest wraps MEOS C function th3index_unnest.
func Th3indexUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_unnest(temp._inner, (**C.uint64_t)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// Th3indexValueAtTimestamptz wraps MEOS C function th3index_value_at_timestamptz.
func Th3indexValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 uint64, _err error) {
	var _out_result C.uint64_t
	C.meos_errno_reset()
	_cret := C.th3index_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), uint64(_out_result), nil
}


// TbigintToTh3index wraps MEOS C function tbigint_to_th3index.
func TbigintToTh3index(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tbigint_to_th3index(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexToTbigint wraps MEOS C function th3index_to_tbigint.
func Th3indexToTbigint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_to_tbigint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexAtValue wraps MEOS C function th3index_at_value.
func Th3indexAtValue(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_at_value(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexMinusValue wraps MEOS C function th3index_minus_value.
func Th3indexMinusValue(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_minus_value(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// EverEqH3indexTh3index wraps MEOS C function ever_eq_h3index_th3index.
func EverEqH3indexTh3index(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_h3index_th3index(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTh3indexH3index wraps MEOS C function ever_eq_th3index_h3index.
func EverEqTh3indexH3index(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_th3index_h3index(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeH3indexTh3index wraps MEOS C function ever_ne_h3index_th3index.
func EverNeH3indexTh3index(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_h3index_th3index(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTh3indexH3index wraps MEOS C function ever_ne_th3index_h3index.
func EverNeTh3indexH3index(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_th3index_h3index(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqH3indexTh3index wraps MEOS C function always_eq_h3index_th3index.
func AlwaysEqH3indexTh3index(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_h3index_th3index(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTh3indexH3index wraps MEOS C function always_eq_th3index_h3index.
func AlwaysEqTh3indexH3index(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_th3index_h3index(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeH3indexTh3index wraps MEOS C function always_ne_h3index_th3index.
func AlwaysNeH3indexTh3index(cell uint64, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_h3index_th3index(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTh3indexH3index wraps MEOS C function always_ne_th3index_h3index.
func AlwaysNeTh3indexH3index(temp *Temporal, cell uint64) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_th3index_h3index(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTh3indexTh3index wraps MEOS C function ever_eq_th3index_th3index.
func EverEqTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_th3index_th3index(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTh3indexTh3index wraps MEOS C function ever_ne_th3index_th3index.
func EverNeTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_th3index_th3index(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTh3indexTh3index wraps MEOS C function always_eq_th3index_th3index.
func AlwaysEqTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_th3index_th3index(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTh3indexTh3index wraps MEOS C function always_ne_th3index_th3index.
func AlwaysNeTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_th3index_th3index(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqH3indexTh3index wraps MEOS C function teq_h3index_th3index.
func TeqH3indexTh3index(cell uint64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_h3index_th3index(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTh3indexH3index wraps MEOS C function teq_th3index_h3index.
func TeqTh3indexH3index(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_th3index_h3index(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTh3indexTh3index wraps MEOS C function teq_th3index_th3index.
func TeqTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_th3index_th3index(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneH3indexTh3index wraps MEOS C function tne_h3index_th3index.
func TneH3indexTh3index(cell uint64, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_h3index_th3index(C.uint64_t(cell), temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTh3indexH3index wraps MEOS C function tne_th3index_h3index.
func TneTh3indexH3index(temp *Temporal, cell uint64) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_th3index_h3index(temp._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTh3indexTh3index wraps MEOS C function tne_th3index_th3index.
func TneTh3indexTh3index(temp1 *Temporal, temp2 *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_th3index_th3index(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexGetBaseCellNumber wraps MEOS C function th3index_get_base_cell_number.
func Th3indexGetBaseCellNumber(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_get_base_cell_number(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexIsResClassIii wraps MEOS C function th3index_is_res_class_iii.
func Th3indexIsResClassIii(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_is_res_class_iii(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexIsPentagon wraps MEOS C function th3index_is_pentagon.
func Th3indexIsPentagon(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_is_pentagon(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexCellToParentNext wraps MEOS C function th3index_cell_to_parent_next.
func Th3indexCellToParentNext(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_cell_to_parent_next(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexCellToCenterChild wraps MEOS C function th3index_cell_to_center_child.
func Th3indexCellToCenterChild(temp *Temporal, resolution int32) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_cell_to_center_child(temp._inner, C.int32(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexCellToCenterChildNext wraps MEOS C function th3index_cell_to_center_child_next.
func Th3indexCellToCenterChildNext(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_cell_to_center_child_next(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexCellToChildPos wraps MEOS C function th3index_cell_to_child_pos.
func Th3indexCellToChildPos(temp *Temporal, parent_res int32) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_cell_to_child_pos(temp._inner, C.int32(parent_res))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexChildPosToCell wraps MEOS C function th3index_child_pos_to_cell.
func Th3indexChildPosToCell(child_pos *Temporal, parent *Temporal, child_res int32) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_child_pos_to_cell(child_pos._inner, parent._inner, C.int32(child_res))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeogpointToTh3index wraps MEOS C function tgeogpoint_to_th3index.
func TgeogpointToTh3index(temp *Temporal, resolution int32, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeogpoint_to_th3index(temp._inner, C.int32(resolution), C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeompointToTh3index wraps MEOS C function tgeompoint_to_th3index.
func TgeompointToTh3index(temp *Temporal, resolution int32, border_inc bool) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeompoint_to_th3index(temp._inner, C.int32(resolution), C.bool(border_inc))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeogpointH3indexSplit wraps MEOS C function tgeogpoint_h3index_split.
func TgeogpointH3indexSplit(temp *Temporal, resolution int32, border_inc bool, cells unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeogpoint_h3index_split(temp._inner, C.int32(resolution), C.bool(border_inc), (**C.uint64_t)(unsafe.Pointer(cells)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TgeompointH3indexSplit wraps MEOS C function tgeompoint_h3index_split.
func TgeompointH3indexSplit(temp *Temporal, resolution int32, border_inc bool, cells unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeompoint_h3index_split(temp._inner, C.int32(resolution), C.bool(border_inc), (**C.uint64_t)(unsafe.Pointer(cells)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// Th3indexToTgeogpoint wraps MEOS C function th3index_to_tgeogpoint.
func Th3indexToTgeogpoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_to_tgeogpoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexToTgeompoint wraps MEOS C function th3index_to_tgeompoint.
func Th3indexToTgeompoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_to_tgeompoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// H3indexToSet wraps MEOS C function h3index_to_set.
func H3indexToSet(cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_to_set(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3indexsetIn wraps MEOS C function h3indexset_in.
func H3indexsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.h3indexset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3indexsetOut wraps MEOS C function h3indexset_out.
func H3indexsetOut(s *Set) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.h3indexset_out(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// H3indexsetMake wraps MEOS C function h3indexset_make.
func H3indexsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3indexset_make((*C.uint64_t)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3indexsetEndValue wraps MEOS C function h3indexset_end_value.
func H3indexsetEndValue(s *Set) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.h3indexset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// H3indexsetStartValue wraps MEOS C function h3indexset_start_value.
func H3indexsetStartValue(s *Set) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.h3indexset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// H3indexsetValueN wraps MEOS C function h3indexset_value_n.
func H3indexsetValueN(s *Set, n int) (_r0 bool, _r1 uint64, _err error) {
	var _out_result C.uint64_t
	C.meos_errno_reset()
	_cret := C.h3indexset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), uint64(_out_result), nil
}


// H3indexsetValues wraps MEOS C function h3indexset_values.
func H3indexsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.h3indexset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ContainedH3indexSet wraps MEOS C function contained_h3index_set.
func ContainedH3indexSet(cell uint64, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_h3index_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsSetH3index wraps MEOS C function contains_set_h3index.
func ContainsSetH3index(s *Set, cell uint64) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_h3index(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionH3indexSet wraps MEOS C function intersection_h3index_set.
func IntersectionH3indexSet(cell uint64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_h3index_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetH3index wraps MEOS C function intersection_set_h3index.
func IntersectionSetH3index(s *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_h3index(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusH3indexSet wraps MEOS C function minus_h3index_set.
func MinusH3indexSet(cell uint64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_h3index_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetH3index wraps MEOS C function minus_set_h3index.
func MinusSetH3index(s *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_h3index(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionH3indexSet wraps MEOS C function union_h3index_set.
func UnionH3indexSet(cell uint64, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_h3index_set(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetH3index wraps MEOS C function union_set_h3index.
func UnionSetH3index(s *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_h3index(s._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3indexUnionTransfn wraps MEOS C function h3index_union_transfn.
func H3indexUnionTransfn(state *Set, cell uint64) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_union_transfn(state._inner, C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3indexsetGridDisk wraps MEOS C function h3indexset_grid_disk.
func H3indexsetGridDisk(cells *Set, k int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3indexset_grid_disk(cells._inner, C.int(k))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3indexsetCoverCells wraps MEOS C function h3indexset_cover_cells.
func H3indexsetCoverCells(cells *Set, resolution int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3indexset_cover_cells(cells._inner, C.int32(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// GeoToH3indexCell wraps MEOS C function geo_to_h3index_cell.
func GeoToH3indexCell(point *Geom, resolution int32) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_to_h3index_cell(point._inner, C.int32(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// GeoToH3indexSet wraps MEOS C function geo_to_h3index_set.
func GeoToH3indexSet(gs *Geom, resolution int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.geo_to_h3index_set(gs._inner, C.int32(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3indexCellToCover wraps MEOS C function h3index_cell_to_cover.
func H3indexCellToCover(cell uint64, resolution int32) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_cell_to_cover(C.uint64_t(cell), C.int32(resolution))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// H3indexToSTBOX wraps MEOS C function h3index_to_stbox.
func H3indexToSTBOX(cell uint64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_to_stbox(C.uint64_t(cell))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// H3indexTimestamptzToSTBOX wraps MEOS C function h3index_timestamptz_to_stbox.
func H3indexTimestamptzToSTBOX(cell uint64, t int64) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_timestamptz_to_stbox(C.uint64_t(cell), C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// H3indexTstzspanToSTBOX wraps MEOS C function h3index_tstzspan_to_stbox.
func H3indexTstzspanToSTBOX(cell uint64, s *Span) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.h3index_tstzspan_to_stbox(C.uint64_t(cell), s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// EverEqH3indexsetTh3index wraps MEOS C function ever_eq_h3indexset_th3index.
func EverEqH3indexsetTh3index(cells *Set, th3idx *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_h3indexset_th3index(cells._inner, th3idx._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTh3indexH3indexset wraps MEOS C function ever_eq_th3index_h3indexset.
func EverEqTh3indexH3indexset(th3idx *Temporal, cells *Set) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_th3index_h3indexset(th3idx._inner, cells._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// Th3indexAreNeighborCells wraps MEOS C function th3index_are_neighbor_cells.
func Th3indexAreNeighborCells(origin *Temporal, dest *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_are_neighbor_cells(origin._inner, dest._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexCellsToDirectedEdge wraps MEOS C function th3index_cells_to_directed_edge.
func Th3indexCellsToDirectedEdge(origin *Temporal, dest *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_cells_to_directed_edge(origin._inner, dest._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexIsValidDirectedEdge wraps MEOS C function th3index_is_valid_directed_edge.
func Th3indexIsValidDirectedEdge(edge *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_is_valid_directed_edge(edge._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexGetDirectedEdgeOrigin wraps MEOS C function th3index_get_directed_edge_origin.
func Th3indexGetDirectedEdgeOrigin(edge *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_get_directed_edge_origin(edge._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexGetDirectedEdgeDestination wraps MEOS C function th3index_get_directed_edge_destination.
func Th3indexGetDirectedEdgeDestination(edge *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_get_directed_edge_destination(edge._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexDirectedEdgeToBoundary wraps MEOS C function th3index_directed_edge_to_boundary.
func Th3indexDirectedEdgeToBoundary(edge *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_directed_edge_to_boundary(edge._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexCellToVertex wraps MEOS C function th3index_cell_to_vertex.
func Th3indexCellToVertex(temp *Temporal, vertex_num int32) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_cell_to_vertex(temp._inner, C.int32(vertex_num))
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexVertexToLatlng wraps MEOS C function th3index_vertex_to_latlng.
func Th3indexVertexToLatlng(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_vertex_to_latlng(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexIsValidVertex wraps MEOS C function th3index_is_valid_vertex.
func Th3indexIsValidVertex(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_is_valid_vertex(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexGridDistance wraps MEOS C function th3index_grid_distance.
func Th3indexGridDistance(origin *Temporal, dest *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_grid_distance(origin._inner, dest._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexCellToLocalIj wraps MEOS C function th3index_cell_to_local_ij.
func Th3indexCellToLocalIj(origin *Temporal, cell *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_cell_to_local_ij(origin._inner, cell._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexLocalIjToCell wraps MEOS C function th3index_local_ij_to_cell.
func Th3indexLocalIjToCell(origin *Temporal, coord *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_local_ij_to_cell(origin._inner, coord._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// Th3indexEdgeLength wraps MEOS C function th3index_edge_length.
func Th3indexEdgeLength(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.th3index_edge_length(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TgeogpointGreatCircleDistance wraps MEOS C function tgeogpoint_great_circle_distance.
func TgeogpointGreatCircleDistance(a *Temporal, b *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tgeogpoint_great_circle_distance(a._inner, b._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}

