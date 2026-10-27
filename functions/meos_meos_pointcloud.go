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

// PcpointHexIn wraps MEOS C function pcpoint_hex_in.
func PcpointHexIn(str string) (_r0 *Pcpoint, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.pcpoint_hex_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpoint{_inner: _cret}, nil
}


// PcpointHexOut wraps MEOS C function pcpoint_hex_out.
func PcpointHexOut(pt *Pcpoint, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_hex_out(pt._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PcpointFromHexwkb wraps MEOS C function pcpoint_from_hexwkb.
func PcpointFromHexwkb(hexwkb string) (_r0 *Pcpoint, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.pcpoint_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpoint{_inner: _cret}, nil
}


// PcpointAsHexwkb wraps MEOS C function pcpoint_as_hexwkb.
func PcpointAsHexwkb(pt *Pcpoint) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_as_hexwkb(pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PcpointMake wraps MEOS C function pcpoint_make.
func PcpointMake(pcid uint32, values unsafe.Pointer, count int) (_r0 *Pcpoint, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_make(C.uint32_t(pcid), (*C.double)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpoint{_inner: _cret}, nil
}


// PcpointCopy wraps MEOS C function pcpoint_copy.
func PcpointCopy(pt *Pcpoint) (_r0 *Pcpoint, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_copy(pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpoint{_inner: _cret}, nil
}


// PcpointGetPcid wraps MEOS C function pcpoint_get_pcid.
func PcpointGetPcid(pt *Pcpoint) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_get_pcid(pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// PcpointHash wraps MEOS C function pcpoint_hash.
func PcpointHash(pt *Pcpoint) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_hash(pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// PcpointHashExtended wraps MEOS C function pcpoint_hash_extended.
func PcpointHashExtended(pt *Pcpoint, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_hash_extended(pt._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// PcpointMemSize wraps MEOS C function pcpoint_mem_size.
func PcpointMemSize(pt *Pcpoint) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_mem_size(pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// PcpointGetX wraps MEOS C function pcpoint_get_x.
func PcpointGetX(pt *Pcpoint, out unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_get_x(pt._inner, (*C.double)(unsafe.Pointer(out)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpointGetY wraps MEOS C function pcpoint_get_y.
func PcpointGetY(pt *Pcpoint, out unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_get_y(pt._inner, (*C.double)(unsafe.Pointer(out)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpointGetZ wraps MEOS C function pcpoint_get_z.
func PcpointGetZ(pt *Pcpoint, out unsafe.Pointer) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_get_z(pt._inner, (*C.double)(unsafe.Pointer(out)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpointGetDim wraps MEOS C function pcpoint_get_dim.
func PcpointGetDim(pt *Pcpoint, name string, out unsafe.Pointer) (_r0 bool, _err error) {
	_c_name := C.CString(name)
	defer C.free(unsafe.Pointer(_c_name))
	C.meos_errno_reset()
	_cret := C.pcpoint_get_dim(pt._inner, _c_name, (*C.double)(unsafe.Pointer(out)))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpointToTpcbox wraps MEOS C function pcpoint_to_tpcbox.
func PcpointToTpcbox(pt *Pcpoint) (_r0 *TPCBox, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_to_tpcbox(pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// MeosPcSchema wraps MEOS C function meos_pc_schema.
func MeosPcSchema(pcid uint32) (_r0 *PCSchema, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_pc_schema(C.uint32_t(pcid))
	if _err = meosError(); _err != nil {
		return
	}
	return &PCSchema{_inner: _cret}, nil
}


// MeosPcSchemaRegister wraps MEOS C function meos_pc_schema_register.
func MeosPcSchemaRegister(pcid uint32, schema *PCSchema) (_err error) {
	C.meos_errno_reset()
	C.meos_pc_schema_register(C.uint32_t(pcid), schema._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosPcSchemaFromDims wraps MEOS C function meos_pc_schema_from_dims.
func MeosPcSchemaFromDims(pcid uint32, srid int32, compression string, dims unsafe.Pointer, ndims int) (_r0 *PCSchema, _err error) {
	_c_compression := C.CString(compression)
	defer C.free(unsafe.Pointer(_c_compression))
	C.meos_errno_reset()
	_cret := C.meos_pc_schema_from_dims(C.uint32_t(pcid), C.int32_t(srid), _c_compression, (*C.PCDimensionSpec)(unsafe.Pointer(dims)), C.int(ndims))
	if _err = meosError(); _err != nil {
		return
	}
	return &PCSchema{_inner: _cret}, nil
}


// MeosPcSchemaRegisterDims wraps MEOS C function meos_pc_schema_register_dims.
func MeosPcSchemaRegisterDims(pcid uint32, srid int32, compression string, dims unsafe.Pointer, ndims int) (_r0 bool, _err error) {
	_c_compression := C.CString(compression)
	defer C.free(unsafe.Pointer(_c_compression))
	C.meos_errno_reset()
	_cret := C.meos_pc_schema_register_dims(C.uint32_t(pcid), C.int32_t(srid), _c_compression, (*C.PCDimensionSpec)(unsafe.Pointer(dims)), C.int(ndims))
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// MeosPcSchemaRegisterXml wraps MEOS C function meos_pc_schema_register_xml.
func MeosPcSchemaRegisterXml(pcid uint32, schema *PCSchema, xml_text string) (_err error) {
	_c_xml_text := C.CString(xml_text)
	defer C.free(unsafe.Pointer(_c_xml_text))
	C.meos_errno_reset()
	C.meos_pc_schema_register_xml(C.uint32_t(pcid), schema._inner, _c_xml_text)
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosPcSchemaXml wraps MEOS C function meos_pc_schema_xml.
func MeosPcSchemaXml(pcid uint32) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_pc_schema_xml(C.uint32_t(pcid))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// MeosPcSchemaClear wraps MEOS C function meos_pc_schema_clear.
func MeosPcSchemaClear() (_err error) {
	C.meos_errno_reset()
	C.meos_pc_schema_clear()
	if _err = meosError(); _err != nil {
		return
	}
	return nil
}


// MeosPcSchemaSRID wraps MEOS C function meos_pc_schema_srid.
func MeosPcSchemaSRID(pcid uint32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_pc_schema_srid(C.uint32_t(pcid))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// MeosPcSchemaCompression wraps MEOS C function meos_pc_schema_compression.
func MeosPcSchemaCompression(pcid uint32) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_pc_schema_compression(C.uint32_t(pcid))
	if _err = meosError(); _err != nil {
		return
	}
	return C.GoString(_cret), nil
}


// MeosPcSchemaNdims wraps MEOS C function meos_pc_schema_ndims.
func MeosPcSchemaNdims(pcid uint32) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.meos_pc_schema_ndims(C.uint32_t(pcid))
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// MeosSetPointcloudSchemasXml wraps MEOS C function meos_set_pointcloud_schemas_xml.
func MeosSetPointcloudSchemasXml(path *string) (_err error) {
	var _c_path *C.char
	if path != nil {
		_c_path = C.CString(*path)
		defer C.free(unsafe.Pointer(_c_path))
	}
	C.meos_errno_reset()
	C.meos_set_pointcloud_schemas_xml(_c_path)
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


// RtreeCreateTpcbox wraps MEOS C function rtree_create_tpcbox.
func RtreeCreateTpcbox() (_r0 *RTree, _err error) {
	C.meos_errno_reset()
	_cret := C.rtree_create_tpcbox()
	if _err = meosError(); _err != nil {
		return
	}
	return &RTree{_inner: _cret}, nil
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


// PcpointCmp wraps MEOS C function pcpoint_cmp.
func PcpointCmp(pt1 *Pcpoint, pt2 *Pcpoint) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_cmp(pt1._inner, pt2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// PcpointEq wraps MEOS C function pcpoint_eq.
func PcpointEq(pt1 *Pcpoint, pt2 *Pcpoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_eq(pt1._inner, pt2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpointNe wraps MEOS C function pcpoint_ne.
func PcpointNe(pt1 *Pcpoint, pt2 *Pcpoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_ne(pt1._inner, pt2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpointLt wraps MEOS C function pcpoint_lt.
func PcpointLt(pt1 *Pcpoint, pt2 *Pcpoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_lt(pt1._inner, pt2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpointLe wraps MEOS C function pcpoint_le.
func PcpointLe(pt1 *Pcpoint, pt2 *Pcpoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_le(pt1._inner, pt2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpointGt wraps MEOS C function pcpoint_gt.
func PcpointGt(pt1 *Pcpoint, pt2 *Pcpoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_gt(pt1._inner, pt2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpointGe wraps MEOS C function pcpoint_ge.
func PcpointGe(pt1 *Pcpoint, pt2 *Pcpoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_ge(pt1._inner, pt2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpatchHexIn wraps MEOS C function pcpatch_hex_in.
func PcpatchHexIn(str string) (_r0 *Pcpatch, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.pcpatch_hex_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpatch{_inner: _cret}, nil
}


// PcpatchHexOut wraps MEOS C function pcpatch_hex_out.
func PcpatchHexOut(pa *Pcpatch, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_hex_out(pa._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PcpatchFromHexwkb wraps MEOS C function pcpatch_from_hexwkb.
func PcpatchFromHexwkb(hexwkb string) (_r0 *Pcpatch, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.pcpatch_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpatch{_inner: _cret}, nil
}


// PcpatchAsHexwkb wraps MEOS C function pcpatch_as_hexwkb.
func PcpatchAsHexwkb(pa *Pcpatch) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_as_hexwkb(pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PcpatchMake wraps MEOS C function pcpatch_make.
func PcpatchMake(points unsafe.Pointer, count int) (_r0 *Pcpatch, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_make((**C.Pcpoint)(unsafe.Pointer(points)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpatch{_inner: _cret}, nil
}


// PcpatchMakeCoords wraps MEOS C function pcpatch_make_coords.
func PcpatchMakeCoords(pcid uint32, values unsafe.Pointer, count int) (_r0 *Pcpatch, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_make_coords(C.uint32_t(pcid), (*C.double)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpatch{_inner: _cret}, nil
}


// PcpatchCopy wraps MEOS C function pcpatch_copy.
func PcpatchCopy(pa *Pcpatch) (_r0 *Pcpatch, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_copy(pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpatch{_inner: _cret}, nil
}


// PcpatchGetPcid wraps MEOS C function pcpatch_get_pcid.
func PcpatchGetPcid(pa *Pcpatch) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_get_pcid(pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// PcpatchNpoints wraps MEOS C function pcpatch_npoints.
func PcpatchNpoints(pa *Pcpatch) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_npoints(pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// PcpatchPointN wraps MEOS C function pcpatch_point_n.
func PcpatchPointN(pa *Pcpatch, n int) (_r0 *Pcpoint, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_point_n(pa._inner, C.int(n))
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpoint{_inner: _cret}, nil
}


// PcpatchPoints wraps MEOS C function pcpatch_points.
func PcpatchPoints(pa *Pcpatch, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_points(pa._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// PcpatchHash wraps MEOS C function pcpatch_hash.
func PcpatchHash(pa *Pcpatch) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_hash(pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// PcpatchHashExtended wraps MEOS C function pcpatch_hash_extended.
func PcpatchHashExtended(pa *Pcpatch, seed uint64) (_r0 uint64, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_hash_extended(pa._inner, C.uint64_t(seed))
	if _err = meosError(); _err != nil {
		return
	}
	return uint64(_cret), nil
}


// PcpatchMemSize wraps MEOS C function pcpatch_mem_size.
func PcpatchMemSize(pa *Pcpatch) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_mem_size(pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// PcpatchToGeom wraps MEOS C function pcpatch_to_geom.
func PcpatchToGeom(pa *Pcpatch) (_r0 *Geom, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_to_geom(pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Geom{_inner: _cret}, nil
}


// PcpatchCmp wraps MEOS C function pcpatch_cmp.
func PcpatchCmp(pa1 *Pcpatch, pa2 *Pcpatch) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_cmp(pa1._inner, pa2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// PcpatchEq wraps MEOS C function pcpatch_eq.
func PcpatchEq(pa1 *Pcpatch, pa2 *Pcpatch) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_eq(pa1._inner, pa2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpatchNe wraps MEOS C function pcpatch_ne.
func PcpatchNe(pa1 *Pcpatch, pa2 *Pcpatch) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_ne(pa1._inner, pa2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpatchLt wraps MEOS C function pcpatch_lt.
func PcpatchLt(pa1 *Pcpatch, pa2 *Pcpatch) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_lt(pa1._inner, pa2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpatchLe wraps MEOS C function pcpatch_le.
func PcpatchLe(pa1 *Pcpatch, pa2 *Pcpatch) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_le(pa1._inner, pa2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpatchGt wraps MEOS C function pcpatch_gt.
func PcpatchGt(pa1 *Pcpatch, pa2 *Pcpatch) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_gt(pa1._inner, pa2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpatchGe wraps MEOS C function pcpatch_ge.
func PcpatchGe(pa1 *Pcpatch, pa2 *Pcpatch) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_ge(pa1._inner, pa2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// PcpointsetIn wraps MEOS C function pcpointset_in.
func PcpointsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.pcpointset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PcpointsetOut wraps MEOS C function pcpointset_out.
func PcpointsetOut(s *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpointset_out(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PcpointsetMake wraps MEOS C function pcpointset_make.
func PcpointsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpointset_make((**C.Pcpoint)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PcpointToSet wraps MEOS C function pcpoint_to_set.
func PcpointToSet(pt *Pcpoint) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_to_set(pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PcpointsetStartValue wraps MEOS C function pcpointset_start_value.
func PcpointsetStartValue(s *Set) (_r0 *Pcpoint, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpointset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpoint{_inner: _cret}, nil
}


// PcpointsetEndValue wraps MEOS C function pcpointset_end_value.
func PcpointsetEndValue(s *Set) (_r0 *Pcpoint, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpointset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpoint{_inner: _cret}, nil
}


// PcpointsetValueN wraps MEOS C function pcpointset_value_n.
func PcpointsetValueN(s *Set, n int) (_r0 bool, _r1 *Pcpoint, _err error) {
	var _out_result *C.Pcpoint
	C.meos_errno_reset()
	_cret := C.pcpointset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pcpoint{_inner: _out_result}, nil
}


// PcpointsetValues wraps MEOS C function pcpointset_values.
func PcpointsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpointset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ContainsSetPcpoint wraps MEOS C function contains_set_pcpoint.
func ContainsSetPcpoint(s *Set, pt *Pcpoint) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_pcpoint(s._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedPcpointSet wraps MEOS C function contained_pcpoint_set.
func ContainedPcpointSet(pt *Pcpoint, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_pcpoint_set(pt._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionPcpointSet wraps MEOS C function intersection_pcpoint_set.
func IntersectionPcpointSet(pt *Pcpoint, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_pcpoint_set(pt._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetPcpoint wraps MEOS C function intersection_set_pcpoint.
func IntersectionSetPcpoint(s *Set, pt *Pcpoint) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_pcpoint(s._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusPcpointSet wraps MEOS C function minus_pcpoint_set.
func MinusPcpointSet(pt *Pcpoint, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_pcpoint_set(pt._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetPcpoint wraps MEOS C function minus_set_pcpoint.
func MinusSetPcpoint(s *Set, pt *Pcpoint) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_pcpoint(s._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionPcpointSet wraps MEOS C function union_pcpoint_set.
func UnionPcpointSet(pt *Pcpoint, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_pcpoint_set(pt._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetPcpoint wraps MEOS C function union_set_pcpoint.
func UnionSetPcpoint(s *Set, pt *Pcpoint) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_pcpoint(s._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PcpointUnionTransfn wraps MEOS C function pcpoint_union_transfn.
func PcpointUnionTransfn(state *Set, pt *Pcpoint) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpoint_union_transfn(state._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PcpatchsetIn wraps MEOS C function pcpatchset_in.
func PcpatchsetIn(str string) (_r0 *Set, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.pcpatchset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PcpatchsetOut wraps MEOS C function pcpatchset_out.
func PcpatchsetOut(s *Set, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatchset_out(s._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// PcpatchsetMake wraps MEOS C function pcpatchset_make.
func PcpatchsetMake(values unsafe.Pointer, count int) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatchset_make((**C.Pcpatch)(unsafe.Pointer(values)), C.int(count))
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PcpatchToSet wraps MEOS C function pcpatch_to_set.
func PcpatchToSet(pa *Pcpatch) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_to_set(pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PcpatchsetStartValue wraps MEOS C function pcpatchset_start_value.
func PcpatchsetStartValue(s *Set) (_r0 *Pcpatch, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatchset_start_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpatch{_inner: _cret}, nil
}


// PcpatchsetEndValue wraps MEOS C function pcpatchset_end_value.
func PcpatchsetEndValue(s *Set) (_r0 *Pcpatch, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatchset_end_value(s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpatch{_inner: _cret}, nil
}


// PcpatchsetValueN wraps MEOS C function pcpatchset_value_n.
func PcpatchsetValueN(s *Set, n int) (_r0 bool, _r1 *Pcpatch, _err error) {
	var _out_result *C.Pcpatch
	C.meos_errno_reset()
	_cret := C.pcpatchset_value_n(s._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pcpatch{_inner: _out_result}, nil
}


// PcpatchsetValues wraps MEOS C function pcpatchset_values.
func PcpatchsetValues(s *Set, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatchset_values(s._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// ContainsSetPcpatch wraps MEOS C function contains_set_pcpatch.
func ContainsSetPcpatch(s *Set, pa *Pcpatch) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_set_pcpatch(s._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedPcpatchSet wraps MEOS C function contained_pcpatch_set.
func ContainedPcpatchSet(pa *Pcpatch, s *Set) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_pcpatch_set(pa._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// IntersectionPcpatchSet wraps MEOS C function intersection_pcpatch_set.
func IntersectionPcpatchSet(pa *Pcpatch, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_pcpatch_set(pa._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// IntersectionSetPcpatch wraps MEOS C function intersection_set_pcpatch.
func IntersectionSetPcpatch(s *Set, pa *Pcpatch) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_set_pcpatch(s._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusPcpatchSet wraps MEOS C function minus_pcpatch_set.
func MinusPcpatchSet(pa *Pcpatch, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_pcpatch_set(pa._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// MinusSetPcpatch wraps MEOS C function minus_set_pcpatch.
func MinusSetPcpatch(s *Set, pa *Pcpatch) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.minus_set_pcpatch(s._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionPcpatchSet wraps MEOS C function union_pcpatch_set.
func UnionPcpatchSet(pa *Pcpatch, s *Set) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_pcpatch_set(pa._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// UnionSetPcpatch wraps MEOS C function union_set_pcpatch.
func UnionSetPcpatch(s *Set, pa *Pcpatch) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_set_pcpatch(s._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// PcpatchUnionTransfn wraps MEOS C function pcpatch_union_transfn.
func PcpatchUnionTransfn(state *Set, pa *Pcpatch) (_r0 *Set, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_union_transfn(state._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Set{_inner: _cret}, nil
}


// TpcboxAsEWKB wraps MEOS C function tpcbox_as_ewkb.
func TpcboxAsEWKB(box *TPCBox, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_as_ewkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TpcboxAsHexewkb wraps MEOS C function tpcbox_as_hexewkb.
func TpcboxAsHexewkb(box *TPCBox, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_as_hexewkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TpcboxAsHexwkb wraps MEOS C function tpcbox_as_hexwkb.
func TpcboxAsHexwkb(box *TPCBox, variant uint8, size_out unsafe.Pointer) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_as_hexwkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TpcboxAsWKB wraps MEOS C function tpcbox_as_wkb.
func TpcboxAsWKB(box *TPCBox, variant uint8, size_out unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_as_wkb(box._inner, C.uint8_t(variant), (*C.size_t)(unsafe.Pointer(size_out)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TpcboxFromHexwkb wraps MEOS C function tpcbox_from_hexwkb.
func TpcboxFromHexwkb(hexwkb string) (_r0 *TPCBox, _err error) {
	_c_hexwkb := C.CString(hexwkb)
	defer C.free(unsafe.Pointer(_c_hexwkb))
	C.meos_errno_reset()
	_cret := C.tpcbox_from_hexwkb(_c_hexwkb)
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// TpcboxFromWKB wraps MEOS C function tpcbox_from_wkb.
func TpcboxFromWKB(wkb unsafe.Pointer, size uint) (_r0 *TPCBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_from_wkb((*C.uint8_t)(unsafe.Pointer(wkb)), C.size_t(size))
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// TpcboxIn wraps MEOS C function tpcbox_in.
func TpcboxIn(str string) (_r0 *TPCBox, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpcbox_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// TpcboxOut wraps MEOS C function tpcbox_out.
func TpcboxOut(box *TPCBox, maxdd int) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_out(box._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TpcboxMake wraps MEOS C function tpcbox_make.
func TpcboxMake(hasx bool, hasz bool, hast bool, geodetic bool, pcid uint32, xmin float64, xmax float64, ymin float64, ymax float64, zmin float64, zmax float64, period *Span) (_r0 *TPCBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_make(C.bool(hasx), C.bool(hasz), C.bool(hast), C.bool(geodetic), C.uint32_t(pcid), C.double(xmin), C.double(xmax), C.double(ymin), C.double(ymax), C.double(zmin), C.double(zmax), period._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// TpcboxCopy wraps MEOS C function tpcbox_copy.
func TpcboxCopy(box *TPCBox) (_r0 *TPCBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_copy(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// PcpatchToTpcbox wraps MEOS C function pcpatch_to_tpcbox.
func PcpatchToTpcbox(pa *Pcpatch) (_r0 *TPCBox, _err error) {
	C.meos_errno_reset()
	_cret := C.pcpatch_to_tpcbox(pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// TpcboxHasx wraps MEOS C function tpcbox_hasx.
func TpcboxHasx(box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_hasx(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcboxHasz wraps MEOS C function tpcbox_hasz.
func TpcboxHasz(box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_hasz(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcboxHast wraps MEOS C function tpcbox_hast.
func TpcboxHast(box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_hast(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcboxGeodetic wraps MEOS C function tpcbox_geodetic.
func TpcboxGeodetic(box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_geodetic(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcboxXmin wraps MEOS C function tpcbox_xmin.
func TpcboxXmin(box *TPCBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tpcbox_xmin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TpcboxXmax wraps MEOS C function tpcbox_xmax.
func TpcboxXmax(box *TPCBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tpcbox_xmax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TpcboxYmin wraps MEOS C function tpcbox_ymin.
func TpcboxYmin(box *TPCBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tpcbox_ymin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TpcboxYmax wraps MEOS C function tpcbox_ymax.
func TpcboxYmax(box *TPCBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tpcbox_ymax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TpcboxZmin wraps MEOS C function tpcbox_zmin.
func TpcboxZmin(box *TPCBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tpcbox_zmin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TpcboxZmax wraps MEOS C function tpcbox_zmax.
func TpcboxZmax(box *TPCBox) (_r0 bool, _r1 float64, _err error) {
	var _out_result C.double
	C.meos_errno_reset()
	_cret := C.tpcbox_zmax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), float64(_out_result), nil
}


// TpcboxTmin wraps MEOS C function tpcbox_tmin.
func TpcboxTmin(box *TPCBox) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.TimestampTz
	C.meos_errno_reset()
	_cret := C.tpcbox_tmin(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// TpcboxTminInc wraps MEOS C function tpcbox_tmin_inc.
func TpcboxTminInc(box *TPCBox) (_r0 bool, _r1 bool, _err error) {
	var _out_result C.bool
	C.meos_errno_reset()
	_cret := C.tpcbox_tmin_inc(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), bool(_out_result), nil
}


// TpcboxTmax wraps MEOS C function tpcbox_tmax.
func TpcboxTmax(box *TPCBox) (_r0 bool, _r1 int64, _err error) {
	var _out_result C.TimestampTz
	C.meos_errno_reset()
	_cret := C.tpcbox_tmax(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), int64(_out_result), nil
}


// TpcboxTmaxInc wraps MEOS C function tpcbox_tmax_inc.
func TpcboxTmaxInc(box *TPCBox) (_r0 bool, _r1 bool, _err error) {
	var _out_result C.bool
	C.meos_errno_reset()
	_cret := C.tpcbox_tmax_inc(box._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), bool(_out_result), nil
}


// TpcboxSRID wraps MEOS C function tpcbox_srid.
func TpcboxSRID(box *TPCBox) (_r0 int32, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_srid(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int32(_cret), nil
}


// TpcboxPcid wraps MEOS C function tpcbox_pcid.
func TpcboxPcid(box *TPCBox) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_pcid(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TpcboxToSTBOX wraps MEOS C function tpcbox_to_stbox.
func TpcboxToSTBOX(box *TPCBox) (_r0 *STBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_to_stbox(box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &STBox{_inner: _cret}, nil
}


// TpcboxRound wraps MEOS C function tpcbox_round.
func TpcboxRound(box *TPCBox, maxdd int) (_r0 *TPCBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_round(box._inner, C.int(maxdd))
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// TpcboxSetSRID wraps MEOS C function tpcbox_set_srid.
func TpcboxSetSRID(box *TPCBox, srid int32) (_r0 *TPCBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_set_srid(box._inner, C.int32_t(srid))
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// UnionTpcboxTpcbox wraps MEOS C function union_tpcbox_tpcbox.
func UnionTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox, strict bool) (_r0 *TPCBox, _err error) {
	C.meos_errno_reset()
	_cret := C.gunion_tpcbox_tpcbox(box1._inner, box2._inner, C.bool(strict))
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// InterTpcboxTpcbox wraps MEOS C function inter_tpcbox_tpcbox.
func InterTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _r1 *TPCBox, _err error) {
	var _out_result C.TPCBox
	C.meos_errno_reset()
	_cret := C.inter_tpcbox_tpcbox(box1._inner, box2._inner, &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &TPCBox{_inner: &_out_result}, nil
}


// IntersectionTpcboxTpcbox wraps MEOS C function intersection_tpcbox_tpcbox.
func IntersectionTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 *TPCBox, _err error) {
	C.meos_errno_reset()
	_cret := C.intersection_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// ContainsTpcboxTpcbox wraps MEOS C function contains_tpcbox_tpcbox.
func ContainsTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTpcboxTpcbox wraps MEOS C function contained_tpcbox_tpcbox.
func ContainedTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTpcboxTpcbox wraps MEOS C function overlaps_tpcbox_tpcbox.
func OverlapsTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTpcboxTpcbox wraps MEOS C function same_tpcbox_tpcbox.
func SameTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTpcboxTpcbox wraps MEOS C function adjacent_tpcbox_tpcbox.
func AdjacentTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcboxCmp wraps MEOS C function tpcbox_cmp.
func TpcboxCmp(box1 *TPCBox, box2 *TPCBox) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_cmp(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TpcboxEq wraps MEOS C function tpcbox_eq.
func TpcboxEq(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_eq(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcboxNe wraps MEOS C function tpcbox_ne.
func TpcboxNe(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_ne(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcboxSortkey wraps MEOS C function tpcbox_sortkey.
func TpcboxSortkey(box *TPCBox, size unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_sortkey(box._inner, (*C.size_t)(unsafe.Pointer(size)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TpcboxLt wraps MEOS C function tpcbox_lt.
func TpcboxLt(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_lt(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcboxLe wraps MEOS C function tpcbox_le.
func TpcboxLe(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_le(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcboxGt wraps MEOS C function tpcbox_gt.
func TpcboxGt(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_gt(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcboxGe wraps MEOS C function tpcbox_ge.
func TpcboxGe(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_ge(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTpcboxTpcbox wraps MEOS C function left_tpcbox_tpcbox.
func LeftTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTpcboxTpcbox wraps MEOS C function overleft_tpcbox_tpcbox.
func OverleftTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTpcboxTpcbox wraps MEOS C function right_tpcbox_tpcbox.
func RightTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTpcboxTpcbox wraps MEOS C function overright_tpcbox_tpcbox.
func OverrightTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BelowTpcboxTpcbox wraps MEOS C function below_tpcbox_tpcbox.
func BelowTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.below_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbelowTpcboxTpcbox wraps MEOS C function overbelow_tpcbox_tpcbox.
func OverbelowTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbelow_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AboveTpcboxTpcbox wraps MEOS C function above_tpcbox_tpcbox.
func AboveTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.above_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OveraboveTpcboxTpcbox wraps MEOS C function overabove_tpcbox_tpcbox.
func OveraboveTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overabove_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// FrontTpcboxTpcbox wraps MEOS C function front_tpcbox_tpcbox.
func FrontTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.front_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverfrontTpcboxTpcbox wraps MEOS C function overfront_tpcbox_tpcbox.
func OverfrontTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overfront_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BackTpcboxTpcbox wraps MEOS C function back_tpcbox_tpcbox.
func BackTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.back_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbackTpcboxTpcbox wraps MEOS C function overback_tpcbox_tpcbox.
func OverbackTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overback_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTpcboxTpcbox wraps MEOS C function before_tpcbox_tpcbox.
func BeforeTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTpcboxTpcbox wraps MEOS C function overbefore_tpcbox_tpcbox.
func OverbeforeTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTpcboxTpcbox wraps MEOS C function after_tpcbox_tpcbox.
func AfterTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTpcboxTpcbox wraps MEOS C function overafter_tpcbox_tpcbox.
func OverafterTpcboxTpcbox(box1 *TPCBox, box2 *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_tpcbox_tpcbox(box1._inner, box2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcpointIn wraps MEOS C function tpcpoint_in.
func TpcpointIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpcpoint_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpointOut wraps MEOS C function tpcpoint_out.
func TpcpointOut(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpoint_out(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TpcpointFromMFJSON wraps MEOS C function tpcpoint_from_mfjson.
func TpcpointFromMFJSON(mfjson string) (_r0 *Temporal, _err error) {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	C.meos_errno_reset()
	_cret := C.tpcpoint_from_mfjson(_c_mfjson)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpatchIn wraps MEOS C function tpcpatch_in.
func TpcpatchIn(str string) (_r0 *Temporal, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpcpatch_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpatchOut wraps MEOS C function tpcpatch_out.
func TpcpatchOut(temp *Temporal) (_r0 string, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_out(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	defer C.free(unsafe.Pointer(_cret))
	return C.GoString(_cret), nil
}


// TpcpatchFromMFJSON wraps MEOS C function tpcpatch_from_mfjson.
func TpcpatchFromMFJSON(mfjson string) (_r0 *Temporal, _err error) {
	_c_mfjson := C.CString(mfjson)
	defer C.free(unsafe.Pointer(_c_mfjson))
	C.meos_errno_reset()
	_cret := C.tpcpatch_from_mfjson(_c_mfjson)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpointinstIn wraps MEOS C function tpcpointinst_in.
func TpcpointinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpcpointinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TpcpointseqIn wraps MEOS C function tpcpointseq_in.
func TpcpointseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpcpointseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TpcpointseqsetIn wraps MEOS C function tpcpointseqset_in.
func TpcpointseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpcpointseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TpcpatchinstIn wraps MEOS C function tpcpatchinst_in.
func TpcpatchinstIn(str string) (_r0 *TInstant, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpcpatchinst_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TpcpatchseqIn wraps MEOS C function tpcpatchseq_in.
func TpcpatchseqIn(str string, interp Interpolation) (_r0 *TSequence, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpcpatchseq_in(_c_str, C.interpType(interp))
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TpcpatchseqsetIn wraps MEOS C function tpcpatchseqset_in.
func TpcpatchseqsetIn(str string) (_r0 *TSequenceSet, _err error) {
	_c_str := C.CString(str)
	defer C.free(unsafe.Pointer(_c_str))
	C.meos_errno_reset()
	_cret := C.tpcpatchseqset_in(_c_str)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TpcpointToTgeompoint wraps MEOS C function tpcpoint_to_tgeompoint.
func TpcpointToTgeompoint(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpoint_to_tgeompoint(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpatchToTgeometry wraps MEOS C function tpcpatch_to_tgeometry.
func TpcpatchToTgeometry(temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_to_tgeometry(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpointinstMake wraps MEOS C function tpcpointinst_make.
func TpcpointinstMake(pt *Pcpoint, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpointinst_make(pt._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TpcpointseqFromBaseTstzset wraps MEOS C function tpcpointseq_from_base_tstzset.
func TpcpointseqFromBaseTstzset(pt *Pcpoint, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpointseq_from_base_tstzset(pt._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TpcpointseqFromBaseTstzspan wraps MEOS C function tpcpointseq_from_base_tstzspan.
func TpcpointseqFromBaseTstzspan(pt *Pcpoint, sp *Span) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpointseq_from_base_tstzspan(pt._inner, sp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TpcpointseqsetFromBaseTstzspanset wraps MEOS C function tpcpointseqset_from_base_tstzspanset.
func TpcpointseqsetFromBaseTstzspanset(pt *Pcpoint, ss *SpanSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpointseqset_from_base_tstzspanset(pt._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TpcpointFromBaseTemp wraps MEOS C function tpcpoint_from_base_temp.
func TpcpointFromBaseTemp(pt *Pcpoint, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpoint_from_base_temp(pt._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpointStartValue wraps MEOS C function tpcpoint_start_value.
func TpcpointStartValue(temp *Temporal) (_r0 *Pcpoint, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpoint_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpoint{_inner: _cret}, nil
}


// TpcpointEndValue wraps MEOS C function tpcpoint_end_value.
func TpcpointEndValue(temp *Temporal) (_r0 *Pcpoint, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpoint_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpoint{_inner: _cret}, nil
}


// TpcpointValue wraps MEOS C function tpcpoint_value.
func TpcpointValue(temp *Temporal) (_r0 *Pcpoint, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpoint_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpoint{_inner: _cret}, nil
}


// TpcpointValueN wraps MEOS C function tpcpoint_value_n.
func TpcpointValueN(temp *Temporal, n int) (_r0 bool, _r1 *Pcpoint, _err error) {
	var _out_result *C.Pcpoint
	C.meos_errno_reset()
	_cret := C.tpcpoint_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pcpoint{_inner: _out_result}, nil
}


// TpcpointValues wraps MEOS C function tpcpoint_values.
func TpcpointValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpoint_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TpcpointUnnest wraps MEOS C function tpcpoint_unnest.
func TpcpointUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpoint_unnest(temp._inner, (***C.Pcpoint)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TpcpointValueAtTimestamptz wraps MEOS C function tpcpoint_value_at_timestamptz.
func TpcpointValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 *Pcpoint, _err error) {
	var _out_value *C.Pcpoint
	C.meos_errno_reset()
	_cret := C.tpcpoint_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pcpoint{_inner: _out_value}, nil
}


// TpcpointAtValue wraps MEOS C function tpcpoint_at_value.
func TpcpointAtValue(temp *Temporal, pt *Pcpoint) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpoint_at_value(temp._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpointMinusValue wraps MEOS C function tpcpoint_minus_value.
func TpcpointMinusValue(temp *Temporal, pt *Pcpoint) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpoint_minus_value(temp._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpointPcid wraps MEOS C function tpcpoint_pcid.
func TpcpointPcid(temp *Temporal) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpoint_pcid(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TpcpatchinstMake wraps MEOS C function tpcpatchinst_make.
func TpcpatchinstMake(pa *Pcpatch, t int64) (_r0 *TInstant, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatchinst_make(pa._inner, C.TimestampTz(t))
	if _err = meosError(); _err != nil {
		return
	}
	return &TInstant{_inner: _cret}, nil
}


// TpcpatchseqFromBaseTstzset wraps MEOS C function tpcpatchseq_from_base_tstzset.
func TpcpatchseqFromBaseTstzset(pa *Pcpatch, s *Set) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatchseq_from_base_tstzset(pa._inner, s._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TpcpatchseqFromBaseTstzspan wraps MEOS C function tpcpatchseq_from_base_tstzspan.
func TpcpatchseqFromBaseTstzspan(pa *Pcpatch, sp *Span) (_r0 *TSequence, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatchseq_from_base_tstzspan(pa._inner, sp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequence{_inner: _cret}, nil
}


// TpcpatchseqsetFromBaseTstzspanset wraps MEOS C function tpcpatchseqset_from_base_tstzspanset.
func TpcpatchseqsetFromBaseTstzspanset(pa *Pcpatch, ss *SpanSet) (_r0 *TSequenceSet, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatchseqset_from_base_tstzspanset(pa._inner, ss._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TSequenceSet{_inner: _cret}, nil
}


// TpcpatchFromBaseTemp wraps MEOS C function tpcpatch_from_base_temp.
func TpcpatchFromBaseTemp(pa *Pcpatch, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_from_base_temp(pa._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpatchStartValue wraps MEOS C function tpcpatch_start_value.
func TpcpatchStartValue(temp *Temporal) (_r0 *Pcpatch, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_start_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpatch{_inner: _cret}, nil
}


// TpcpatchEndValue wraps MEOS C function tpcpatch_end_value.
func TpcpatchEndValue(temp *Temporal) (_r0 *Pcpatch, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_end_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpatch{_inner: _cret}, nil
}


// TpcpatchValue wraps MEOS C function tpcpatch_value.
func TpcpatchValue(temp *Temporal) (_r0 *Pcpatch, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_value(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Pcpatch{_inner: _cret}, nil
}


// TpcpatchValueN wraps MEOS C function tpcpatch_value_n.
func TpcpatchValueN(temp *Temporal, n int) (_r0 bool, _r1 *Pcpatch, _err error) {
	var _out_result *C.Pcpatch
	C.meos_errno_reset()
	_cret := C.tpcpatch_value_n(temp._inner, C.int(n), &_out_result)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pcpatch{_inner: _out_result}, nil
}


// TpcpatchValues wraps MEOS C function tpcpatch_values.
func TpcpatchValues(temp *Temporal, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_values(temp._inner, (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TpcpatchUnnest wraps MEOS C function tpcpatch_unnest.
func TpcpatchUnnest(temp *Temporal, values unsafe.Pointer, count unsafe.Pointer) (_r0 unsafe.Pointer, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_unnest(temp._inner, (***C.Pcpatch)(unsafe.Pointer(values)), (*C.int)(unsafe.Pointer(count)))
	if _err = meosError(); _err != nil {
		return
	}
	return unsafe.Pointer(_cret), nil
}


// TpcpatchValueAtTimestamptz wraps MEOS C function tpcpatch_value_at_timestamptz.
func TpcpatchValueAtTimestamptz(temp *Temporal, t int64, strict bool) (_r0 bool, _r1 *Pcpatch, _err error) {
	var _out_value *C.Pcpatch
	C.meos_errno_reset()
	_cret := C.tpcpatch_value_at_timestamptz(temp._inner, C.TimestampTz(t), C.bool(strict), &_out_value)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), &Pcpatch{_inner: _out_value}, nil
}


// TpcpatchAtValue wraps MEOS C function tpcpatch_at_value.
func TpcpatchAtValue(temp *Temporal, pa *Pcpatch) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_at_value(temp._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpatchMinusValue wraps MEOS C function tpcpatch_minus_value.
func TpcpatchMinusValue(temp *Temporal, pa *Pcpatch) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_minus_value(temp._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TpcpatchPcid wraps MEOS C function tpcpatch_pcid.
func TpcpatchPcid(temp *Temporal) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_pcid(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TpcpatchStartNpoints wraps MEOS C function tpcpatch_start_npoints.
func TpcpatchStartNpoints(temp *Temporal) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_start_npoints(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TpcpatchEndNpoints wraps MEOS C function tpcpatch_end_npoints.
func TpcpatchEndNpoints(temp *Temporal) (_r0 uint32, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_end_npoints(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return uint32(_cret), nil
}


// TpcpatchNpoints wraps MEOS C function tpcpatch_npoints.
func TpcpatchNpoints(temp *Temporal) (_r0 int64, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_npoints(temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int64(_cret), nil
}


// EverEqPcpointTpcpoint wraps MEOS C function ever_eq_pcpoint_tpcpoint.
func EverEqPcpointTpcpoint(pt *Pcpoint, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_pcpoint_tpcpoint(pt._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTpcpointPcpoint wraps MEOS C function ever_eq_tpcpoint_pcpoint.
func EverEqTpcpointPcpoint(temp *Temporal, pt *Pcpoint) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tpcpoint_pcpoint(temp._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTpcpointTpcpoint wraps MEOS C function ever_eq_tpcpoint_tpcpoint.
func EverEqTpcpointTpcpoint(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tpcpoint_tpcpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNePcpointTpcpoint wraps MEOS C function ever_ne_pcpoint_tpcpoint.
func EverNePcpointTpcpoint(pt *Pcpoint, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_pcpoint_tpcpoint(pt._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTpcpointPcpoint wraps MEOS C function ever_ne_tpcpoint_pcpoint.
func EverNeTpcpointPcpoint(temp *Temporal, pt *Pcpoint) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tpcpoint_pcpoint(temp._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTpcpointTpcpoint wraps MEOS C function ever_ne_tpcpoint_tpcpoint.
func EverNeTpcpointTpcpoint(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tpcpoint_tpcpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqPcpointTpcpoint wraps MEOS C function always_eq_pcpoint_tpcpoint.
func AlwaysEqPcpointTpcpoint(pt *Pcpoint, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_pcpoint_tpcpoint(pt._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTpcpointPcpoint wraps MEOS C function always_eq_tpcpoint_pcpoint.
func AlwaysEqTpcpointPcpoint(temp *Temporal, pt *Pcpoint) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tpcpoint_pcpoint(temp._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTpcpointTpcpoint wraps MEOS C function always_eq_tpcpoint_tpcpoint.
func AlwaysEqTpcpointTpcpoint(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tpcpoint_tpcpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNePcpointTpcpoint wraps MEOS C function always_ne_pcpoint_tpcpoint.
func AlwaysNePcpointTpcpoint(pt *Pcpoint, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_pcpoint_tpcpoint(pt._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTpcpointPcpoint wraps MEOS C function always_ne_tpcpoint_pcpoint.
func AlwaysNeTpcpointPcpoint(temp *Temporal, pt *Pcpoint) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tpcpoint_pcpoint(temp._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTpcpointTpcpoint wraps MEOS C function always_ne_tpcpoint_tpcpoint.
func AlwaysNeTpcpointTpcpoint(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tpcpoint_tpcpoint(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqPcpatchTpcpatch wraps MEOS C function ever_eq_pcpatch_tpcpatch.
func EverEqPcpatchTpcpatch(pa *Pcpatch, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_pcpatch_tpcpatch(pa._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTpcpatchPcpatch wraps MEOS C function ever_eq_tpcpatch_pcpatch.
func EverEqTpcpatchPcpatch(temp *Temporal, pa *Pcpatch) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tpcpatch_pcpatch(temp._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverEqTpcpatchTpcpatch wraps MEOS C function ever_eq_tpcpatch_tpcpatch.
func EverEqTpcpatchTpcpatch(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_eq_tpcpatch_tpcpatch(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNePcpatchTpcpatch wraps MEOS C function ever_ne_pcpatch_tpcpatch.
func EverNePcpatchTpcpatch(pa *Pcpatch, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_pcpatch_tpcpatch(pa._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTpcpatchPcpatch wraps MEOS C function ever_ne_tpcpatch_pcpatch.
func EverNeTpcpatchPcpatch(temp *Temporal, pa *Pcpatch) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tpcpatch_pcpatch(temp._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// EverNeTpcpatchTpcpatch wraps MEOS C function ever_ne_tpcpatch_tpcpatch.
func EverNeTpcpatchTpcpatch(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.ever_ne_tpcpatch_tpcpatch(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqPcpatchTpcpatch wraps MEOS C function always_eq_pcpatch_tpcpatch.
func AlwaysEqPcpatchTpcpatch(pa *Pcpatch, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_pcpatch_tpcpatch(pa._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTpcpatchPcpatch wraps MEOS C function always_eq_tpcpatch_pcpatch.
func AlwaysEqTpcpatchPcpatch(temp *Temporal, pa *Pcpatch) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tpcpatch_pcpatch(temp._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysEqTpcpatchTpcpatch wraps MEOS C function always_eq_tpcpatch_tpcpatch.
func AlwaysEqTpcpatchTpcpatch(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_eq_tpcpatch_tpcpatch(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNePcpatchTpcpatch wraps MEOS C function always_ne_pcpatch_tpcpatch.
func AlwaysNePcpatchTpcpatch(pa *Pcpatch, temp *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_pcpatch_tpcpatch(pa._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTpcpatchPcpatch wraps MEOS C function always_ne_tpcpatch_pcpatch.
func AlwaysNeTpcpatchPcpatch(temp *Temporal, pa *Pcpatch) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tpcpatch_pcpatch(temp._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// AlwaysNeTpcpatchTpcpatch wraps MEOS C function always_ne_tpcpatch_tpcpatch.
func AlwaysNeTpcpatchTpcpatch(temp1 *Temporal, temp2 *Temporal) (_r0 int, _err error) {
	C.meos_errno_reset()
	_cret := C.always_ne_tpcpatch_tpcpatch(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return int(_cret), nil
}


// TeqPcpointTpcpoint wraps MEOS C function teq_pcpoint_tpcpoint.
func TeqPcpointTpcpoint(pt *Pcpoint, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_pcpoint_tpcpoint(pt._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTpcpointPcpoint wraps MEOS C function teq_tpcpoint_pcpoint.
func TeqTpcpointPcpoint(temp *Temporal, pt *Pcpoint) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tpcpoint_pcpoint(temp._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnePcpointTpcpoint wraps MEOS C function tne_pcpoint_tpcpoint.
func TnePcpointTpcpoint(pt *Pcpoint, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_pcpoint_tpcpoint(pt._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTpcpointPcpoint wraps MEOS C function tne_tpcpoint_pcpoint.
func TneTpcpointPcpoint(temp *Temporal, pt *Pcpoint) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tpcpoint_pcpoint(temp._inner, pt._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqPcpatchTpcpatch wraps MEOS C function teq_pcpatch_tpcpatch.
func TeqPcpatchTpcpatch(pa *Pcpatch, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_pcpatch_tpcpatch(pa._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TeqTpcpatchPcpatch wraps MEOS C function teq_tpcpatch_pcpatch.
func TeqTpcpatchPcpatch(temp *Temporal, pa *Pcpatch) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.teq_tpcpatch_pcpatch(temp._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TnePcpatchTpcpatch wraps MEOS C function tne_pcpatch_tpcpatch.
func TnePcpatchTpcpatch(pa *Pcpatch, temp *Temporal) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_pcpatch_tpcpatch(pa._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// TneTpcpatchPcpatch wraps MEOS C function tne_tpcpatch_pcpatch.
func TneTpcpatchPcpatch(temp *Temporal, pa *Pcpatch) (_r0 *Temporal, _err error) {
	C.meos_errno_reset()
	_cret := C.tne_tpcpatch_pcpatch(temp._inner, pa._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &Temporal{_inner: _cret}, nil
}


// AdjacentTpcboxTpointcloud wraps MEOS C function adjacent_tpcbox_tpointcloud.
func AdjacentTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTpointcloudTpcbox wraps MEOS C function adjacent_tpointcloud_tpcbox.
func AdjacentTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AdjacentTpointcloudTpointcloud wraps MEOS C function adjacent_tpointcloud_tpointcloud.
func AdjacentTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.adjacent_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTpcboxTpointcloud wraps MEOS C function contained_tpcbox_tpointcloud.
func ContainedTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTpointcloudTpcbox wraps MEOS C function contained_tpointcloud_tpcbox.
func ContainedTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainedTpointcloudTpointcloud wraps MEOS C function contained_tpointcloud_tpointcloud.
func ContainedTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contained_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTpcboxTpointcloud wraps MEOS C function contains_tpcbox_tpointcloud.
func ContainsTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTpointcloudTpcbox wraps MEOS C function contains_tpointcloud_tpcbox.
func ContainsTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// ContainsTpointcloudTpointcloud wraps MEOS C function contains_tpointcloud_tpointcloud.
func ContainsTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.contains_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTpcboxTpointcloud wraps MEOS C function overlaps_tpcbox_tpointcloud.
func OverlapsTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTpointcloudTpcbox wraps MEOS C function overlaps_tpointcloud_tpcbox.
func OverlapsTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverlapsTpointcloudTpointcloud wraps MEOS C function overlaps_tpointcloud_tpointcloud.
func OverlapsTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overlaps_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTpcboxTpointcloud wraps MEOS C function same_tpcbox_tpointcloud.
func SameTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTpointcloudTpcbox wraps MEOS C function same_tpointcloud_tpcbox.
func SameTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// SameTpointcloudTpointcloud wraps MEOS C function same_tpointcloud_tpointcloud.
func SameTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.same_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTpcboxTpointcloud wraps MEOS C function left_tpcbox_tpointcloud.
func LeftTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTpointcloudTpcbox wraps MEOS C function left_tpointcloud_tpcbox.
func LeftTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// LeftTpointcloudTpointcloud wraps MEOS C function left_tpointcloud_tpointcloud.
func LeftTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.left_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTpcboxTpointcloud wraps MEOS C function overleft_tpcbox_tpointcloud.
func OverleftTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTpointcloudTpcbox wraps MEOS C function overleft_tpointcloud_tpcbox.
func OverleftTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverleftTpointcloudTpointcloud wraps MEOS C function overleft_tpointcloud_tpointcloud.
func OverleftTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overleft_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTpcboxTpointcloud wraps MEOS C function right_tpcbox_tpointcloud.
func RightTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTpointcloudTpcbox wraps MEOS C function right_tpointcloud_tpcbox.
func RightTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// RightTpointcloudTpointcloud wraps MEOS C function right_tpointcloud_tpointcloud.
func RightTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.right_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTpcboxTpointcloud wraps MEOS C function overright_tpcbox_tpointcloud.
func OverrightTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTpointcloudTpcbox wraps MEOS C function overright_tpointcloud_tpcbox.
func OverrightTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverrightTpointcloudTpointcloud wraps MEOS C function overright_tpointcloud_tpointcloud.
func OverrightTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overright_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BelowTpcboxTpointcloud wraps MEOS C function below_tpcbox_tpointcloud.
func BelowTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.below_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BelowTpointcloudTpcbox wraps MEOS C function below_tpointcloud_tpcbox.
func BelowTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.below_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BelowTpointcloudTpointcloud wraps MEOS C function below_tpointcloud_tpointcloud.
func BelowTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.below_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbelowTpcboxTpointcloud wraps MEOS C function overbelow_tpcbox_tpointcloud.
func OverbelowTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbelow_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbelowTpointcloudTpcbox wraps MEOS C function overbelow_tpointcloud_tpcbox.
func OverbelowTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbelow_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbelowTpointcloudTpointcloud wraps MEOS C function overbelow_tpointcloud_tpointcloud.
func OverbelowTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbelow_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AboveTpcboxTpointcloud wraps MEOS C function above_tpcbox_tpointcloud.
func AboveTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.above_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AboveTpointcloudTpcbox wraps MEOS C function above_tpointcloud_tpcbox.
func AboveTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.above_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AboveTpointcloudTpointcloud wraps MEOS C function above_tpointcloud_tpointcloud.
func AboveTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.above_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OveraboveTpcboxTpointcloud wraps MEOS C function overabove_tpcbox_tpointcloud.
func OveraboveTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overabove_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OveraboveTpointcloudTpcbox wraps MEOS C function overabove_tpointcloud_tpcbox.
func OveraboveTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overabove_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OveraboveTpointcloudTpointcloud wraps MEOS C function overabove_tpointcloud_tpointcloud.
func OveraboveTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overabove_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// FrontTpcboxTpointcloud wraps MEOS C function front_tpcbox_tpointcloud.
func FrontTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.front_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// FrontTpointcloudTpcbox wraps MEOS C function front_tpointcloud_tpcbox.
func FrontTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.front_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// FrontTpointcloudTpointcloud wraps MEOS C function front_tpointcloud_tpointcloud.
func FrontTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.front_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverfrontTpcboxTpointcloud wraps MEOS C function overfront_tpcbox_tpointcloud.
func OverfrontTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overfront_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverfrontTpointcloudTpcbox wraps MEOS C function overfront_tpointcloud_tpcbox.
func OverfrontTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overfront_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverfrontTpointcloudTpointcloud wraps MEOS C function overfront_tpointcloud_tpointcloud.
func OverfrontTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overfront_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BackTpcboxTpointcloud wraps MEOS C function back_tpcbox_tpointcloud.
func BackTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.back_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BackTpointcloudTpcbox wraps MEOS C function back_tpointcloud_tpcbox.
func BackTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.back_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BackTpointcloudTpointcloud wraps MEOS C function back_tpointcloud_tpointcloud.
func BackTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.back_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbackTpcboxTpointcloud wraps MEOS C function overback_tpcbox_tpointcloud.
func OverbackTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overback_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbackTpointcloudTpcbox wraps MEOS C function overback_tpointcloud_tpcbox.
func OverbackTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overback_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbackTpointcloudTpointcloud wraps MEOS C function overback_tpointcloud_tpointcloud.
func OverbackTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overback_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTpcboxTpointcloud wraps MEOS C function before_tpcbox_tpointcloud.
func BeforeTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTpointcloudTpcbox wraps MEOS C function before_tpointcloud_tpcbox.
func BeforeTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// BeforeTpointcloudTpointcloud wraps MEOS C function before_tpointcloud_tpointcloud.
func BeforeTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.before_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTpcboxTpointcloud wraps MEOS C function overbefore_tpcbox_tpointcloud.
func OverbeforeTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTpointcloudTpcbox wraps MEOS C function overbefore_tpointcloud_tpcbox.
func OverbeforeTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverbeforeTpointcloudTpointcloud wraps MEOS C function overbefore_tpointcloud_tpointcloud.
func OverbeforeTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overbefore_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTpcboxTpointcloud wraps MEOS C function after_tpcbox_tpointcloud.
func AfterTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTpointcloudTpcbox wraps MEOS C function after_tpointcloud_tpcbox.
func AfterTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// AfterTpointcloudTpointcloud wraps MEOS C function after_tpointcloud_tpointcloud.
func AfterTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.after_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTpcboxTpointcloud wraps MEOS C function overafter_tpcbox_tpointcloud.
func OverafterTpcboxTpointcloud(box *TPCBox, temp *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_tpcbox_tpointcloud(box._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTpointcloudTpcbox wraps MEOS C function overafter_tpointcloud_tpcbox.
func OverafterTpointcloudTpcbox(temp *Temporal, box *TPCBox) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_tpointcloud_tpcbox(temp._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// OverafterTpointcloudTpointcloud wraps MEOS C function overafter_tpointcloud_tpointcloud.
func OverafterTpointcloudTpointcloud(temp1 *Temporal, temp2 *Temporal) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.overafter_tpointcloud_tpointcloud(temp1._inner, temp2._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// TpcExtentTransfn wraps MEOS C function tpc_extent_transfn.
func TpcExtentTransfn(state *TPCBox, temp *Temporal) (_r0 *TPCBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tpc_extent_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// TpcboxExtentTransfn wraps MEOS C function tpcbox_extent_transfn.
func TpcboxExtentTransfn(state *TPCBox, box *TPCBox) (_r0 *TPCBox, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcbox_extent_transfn(state._inner, box._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &TPCBox{_inner: _cret}, nil
}


// TpcpatchTdensityTransfn wraps MEOS C function tpcpatch_tdensity_transfn.
func TpcpatchTdensityTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_tdensity_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// TpcpatchTnpointsTransfn wraps MEOS C function tpcpatch_tnpoints_transfn.
func TpcpatchTnpointsTransfn(state *SkipList, temp *Temporal) (_r0 *SkipList, _err error) {
	C.meos_errno_reset()
	_cret := C.tpcpatch_tnpoints_transfn(state._inner, temp._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return &SkipList{_inner: _cret}, nil
}


// EintersectsTpcpointGeo wraps MEOS C function eintersects_tpcpoint_geo.
func EintersectsTpcpointGeo(temp *Temporal, gs *Geom) (_r0 bool, _err error) {
	C.meos_errno_reset()
	_cret := C.eintersects_tpcpoint_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return bool(_cret), nil
}


// NadTpcpointGeo wraps MEOS C function nad_tpcpoint_geo.
func NadTpcpointGeo(temp *Temporal, gs *Geom) (_r0 float64, _err error) {
	C.meos_errno_reset()
	_cret := C.nad_tpcpoint_geo(temp._inner, gs._inner)
	if _err = meosError(); _err != nil {
		return
	}
	return float64(_cret), nil
}

