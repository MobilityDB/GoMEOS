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
*/
import "C"
import (
	"unsafe"
)

var _ = unsafe.Pointer(nil)

// -------------------- opaque handle types --------------------
type Temporal struct { _inner *C.Temporal }
func (x *Temporal) Inner() *C.Temporal { if x == nil { return nil }; return x._inner }
func (x *Temporal) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func TemporalFromPointer(p unsafe.Pointer) *Temporal { if p == nil { return nil }; return &Temporal{_inner: (*C.Temporal)(p)} }
type TInstant struct { _inner *C.TInstant }
func (x *TInstant) Inner() *C.TInstant { if x == nil { return nil }; return x._inner }
func (x *TInstant) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func TInstantFromPointer(p unsafe.Pointer) *TInstant { if p == nil { return nil }; return &TInstant{_inner: (*C.TInstant)(p)} }
type TSequence struct { _inner *C.TSequence }
func (x *TSequence) Inner() *C.TSequence { if x == nil { return nil }; return x._inner }
func (x *TSequence) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func TSequenceFromPointer(p unsafe.Pointer) *TSequence { if p == nil { return nil }; return &TSequence{_inner: (*C.TSequence)(p)} }
type TSequenceSet struct { _inner *C.TSequenceSet }
func (x *TSequenceSet) Inner() *C.TSequenceSet { if x == nil { return nil }; return x._inner }
func (x *TSequenceSet) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func TSequenceSetFromPointer(p unsafe.Pointer) *TSequenceSet { if p == nil { return nil }; return &TSequenceSet{_inner: (*C.TSequenceSet)(p)} }
type STBox struct { _inner *C.STBox }
func (x *STBox) Inner() *C.STBox { if x == nil { return nil }; return x._inner }
func (x *STBox) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func STBoxFromPointer(p unsafe.Pointer) *STBox { if p == nil { return nil }; return &STBox{_inner: (*C.STBox)(p)} }
type TBox struct { _inner *C.TBox }
func (x *TBox) Inner() *C.TBox { if x == nil { return nil }; return x._inner }
func (x *TBox) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func TBoxFromPointer(p unsafe.Pointer) *TBox { if p == nil { return nil }; return &TBox{_inner: (*C.TBox)(p)} }
type Span struct { _inner *C.Span }
func (x *Span) Inner() *C.Span { if x == nil { return nil }; return x._inner }
func (x *Span) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func SpanFromPointer(p unsafe.Pointer) *Span { if p == nil { return nil }; return &Span{_inner: (*C.Span)(p)} }
type SpanSet struct { _inner *C.SpanSet }
func (x *SpanSet) Inner() *C.SpanSet { if x == nil { return nil }; return x._inner }
func (x *SpanSet) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func SpanSetFromPointer(p unsafe.Pointer) *SpanSet { if p == nil { return nil }; return &SpanSet{_inner: (*C.SpanSet)(p)} }
type Set struct { _inner *C.Set }
func (x *Set) Inner() *C.Set { if x == nil { return nil }; return x._inner }
func (x *Set) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func SetFromPointer(p unsafe.Pointer) *Set { if p == nil { return nil }; return &Set{_inner: (*C.Set)(p)} }
type Geom struct { _inner *C.GSERIALIZED }
func (x *Geom) Inner() *C.GSERIALIZED { if x == nil { return nil }; return x._inner }
func (x *Geom) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func GeomFromPointer(p unsafe.Pointer) *Geom { if p == nil { return nil }; return &Geom{_inner: (*C.GSERIALIZED)(p)} }
type Interval struct { _inner *C.Interval }
func (x *Interval) Inner() *C.Interval { if x == nil { return nil }; return x._inner }
func (x *Interval) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func IntervalFromPointer(p unsafe.Pointer) *Interval { if p == nil { return nil }; return &Interval{_inner: (*C.Interval)(p)} }
type Npoint struct { _inner *C.Npoint }
func (x *Npoint) Inner() *C.Npoint { if x == nil { return nil }; return x._inner }
func (x *Npoint) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func NpointFromPointer(p unsafe.Pointer) *Npoint { if p == nil { return nil }; return &Npoint{_inner: (*C.Npoint)(p)} }
type Nsegment struct { _inner *C.Nsegment }
func (x *Nsegment) Inner() *C.Nsegment { if x == nil { return nil }; return x._inner }
func (x *Nsegment) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func NsegmentFromPointer(p unsafe.Pointer) *Nsegment { if p == nil { return nil }; return &Nsegment{_inner: (*C.Nsegment)(p)} }
type Cbuffer struct { _inner *C.Cbuffer }
func (x *Cbuffer) Inner() *C.Cbuffer { if x == nil { return nil }; return x._inner }
func (x *Cbuffer) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func CbufferFromPointer(p unsafe.Pointer) *Cbuffer { if p == nil { return nil }; return &Cbuffer{_inner: (*C.Cbuffer)(p)} }
type Pose struct { _inner *C.Pose }
func (x *Pose) Inner() *C.Pose { if x == nil { return nil }; return x._inner }
func (x *Pose) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func PoseFromPointer(p unsafe.Pointer) *Pose { if p == nil { return nil }; return &Pose{_inner: (*C.Pose)(p)} }
type Jsonb struct { _inner *C.Jsonb }
func (x *Jsonb) Inner() *C.Jsonb { if x == nil { return nil }; return x._inner }
func (x *Jsonb) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func JsonbFromPointer(p unsafe.Pointer) *Jsonb { if p == nil { return nil }; return &Jsonb{_inner: (*C.Jsonb)(p)} }
type JsonPath struct { _inner *C.JsonPath }
func (x *JsonPath) Inner() *C.JsonPath { if x == nil { return nil }; return x._inner }
func (x *JsonPath) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func JsonPathFromPointer(p unsafe.Pointer) *JsonPath { if p == nil { return nil }; return &JsonPath{_inner: (*C.JsonPath)(p)} }
type Pcpoint struct { _inner *C.Pcpoint }
func (x *Pcpoint) Inner() *C.Pcpoint { if x == nil { return nil }; return x._inner }
func (x *Pcpoint) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func PcpointFromPointer(p unsafe.Pointer) *Pcpoint { if p == nil { return nil }; return &Pcpoint{_inner: (*C.Pcpoint)(p)} }
type Pcpatch struct { _inner *C.Pcpatch }
func (x *Pcpatch) Inner() *C.Pcpatch { if x == nil { return nil }; return x._inner }
func (x *Pcpatch) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func PcpatchFromPointer(p unsafe.Pointer) *Pcpatch { if p == nil { return nil }; return &Pcpatch{_inner: (*C.Pcpatch)(p)} }
type TPCBox struct { _inner *C.TPCBox }
func (x *TPCBox) Inner() *C.TPCBox { if x == nil { return nil }; return x._inner }
func (x *TPCBox) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func TPCBoxFromPointer(p unsafe.Pointer) *TPCBox { if p == nil { return nil }; return &TPCBox{_inner: (*C.TPCBox)(p)} }
type MeosArray struct { _inner *C.MeosArray }
func (x *MeosArray) Inner() *C.MeosArray { if x == nil { return nil }; return x._inner }
func (x *MeosArray) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func MeosArrayFromPointer(p unsafe.Pointer) *MeosArray { if p == nil { return nil }; return &MeosArray{_inner: (*C.MeosArray)(p)} }
type PCSchema struct { _inner *C.PCSCHEMA }
func (x *PCSchema) Inner() *C.PCSCHEMA { if x == nil { return nil }; return x._inner }
func (x *PCSchema) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func PCSchemaFromPointer(p unsafe.Pointer) *PCSchema { if p == nil { return nil }; return &PCSchema{_inner: (*C.PCSCHEMA)(p)} }
type SkipList struct { _inner *C.SkipList }
func (x *SkipList) Inner() *C.SkipList { if x == nil { return nil }; return x._inner }
func (x *SkipList) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func SkipListFromPointer(p unsafe.Pointer) *SkipList { if p == nil { return nil }; return &SkipList{_inner: (*C.SkipList)(p)} }
type RTree struct { _inner *C.RTree }
func (x *RTree) Inner() *C.RTree { if x == nil { return nil }; return x._inner }
func (x *RTree) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func RTreeFromPointer(p unsafe.Pointer) *RTree { if p == nil { return nil }; return &RTree{_inner: (*C.RTree)(p)} }
type Match struct { _inner *C.Match }
func (x *Match) Inner() *C.Match { if x == nil { return nil }; return x._inner }
func (x *Match) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func MatchFromPointer(p unsafe.Pointer) *Match { if p == nil { return nil }; return &Match{_inner: (*C.Match)(p)} }
type Box3D struct { _inner *C.BOX3D }
func (x *Box3D) Inner() *C.BOX3D { if x == nil { return nil }; return x._inner }
func (x *Box3D) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func Box3DFromPointer(p unsafe.Pointer) *Box3D { if p == nil { return nil }; return &Box3D{_inner: (*C.BOX3D)(p)} }
type GBox struct { _inner *C.GBOX }
func (x *GBox) Inner() *C.GBOX { if x == nil { return nil }; return x._inner }
func (x *GBox) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func GBoxFromPointer(p unsafe.Pointer) *GBox { if p == nil { return nil }; return &GBox{_inner: (*C.GBOX)(p)} }
type PJContext struct { _inner *C.PJ_CONTEXT }
func (x *PJContext) Inner() *C.PJ_CONTEXT { if x == nil { return nil }; return x._inner }
func (x *PJContext) Pointer() unsafe.Pointer { if x == nil { return nil }; return unsafe.Pointer(x._inner) }
func PJContextFromPointer(p unsafe.Pointer) *PJContext { if p == nil { return nil }; return &PJContext{_inner: (*C.PJ_CONTEXT)(p)} }

// -------------------- enums --------------------
type TempSubtype C.tempSubtype
const (
	TempSubtype_ANYTEMPSUBTYPE TempSubtype = 0
	TempSubtype_TINSTANT TempSubtype = 1
	TempSubtype_TSEQUENCE TempSubtype = 2
	TempSubtype_TSEQUENCESET TempSubtype = 3
)
type Interpolation C.interpType
const (
	Interpolation_INTERP_NONE Interpolation = 0
	Interpolation_DISCRETE Interpolation = 1
	Interpolation_STEP Interpolation = 2
	Interpolation_LINEAR Interpolation = 3
)
type IndexSearchOp C.IndexSearchOp
const (
	IndexSearchOp_INDEX_OVERLAPS IndexSearchOp = 0
	IndexSearchOp_INDEX_CONTAINS IndexSearchOp = 1
	IndexSearchOp_INDEX_CONTAINED_BY IndexSearchOp = 2
	IndexSearchOp_INDEX_LEFT IndexSearchOp = 3
	IndexSearchOp_INDEX_OVERLEFT IndexSearchOp = 4
	IndexSearchOp_INDEX_RIGHT IndexSearchOp = 5
	IndexSearchOp_INDEX_OVERRIGHT IndexSearchOp = 6
	IndexSearchOp_INDEX_BELOW IndexSearchOp = 7
	IndexSearchOp_INDEX_OVERBELOW IndexSearchOp = 8
	IndexSearchOp_INDEX_ABOVE IndexSearchOp = 9
	IndexSearchOp_INDEX_OVERABOVE IndexSearchOp = 10
	IndexSearchOp_INDEX_FRONT IndexSearchOp = 11
	IndexSearchOp_INDEX_OVERFRONT IndexSearchOp = 12
	IndexSearchOp_INDEX_BACK IndexSearchOp = 13
	IndexSearchOp_INDEX_OVERBACK IndexSearchOp = 14
	IndexSearchOp_INDEX_BEFORE IndexSearchOp = 15
	IndexSearchOp_INDEX_OVERBEFORE IndexSearchOp = 16
	IndexSearchOp_INDEX_AFTER IndexSearchOp = 17
	IndexSearchOp_INDEX_OVERAFTER IndexSearchOp = 18
	IndexSearchOp_INDEX_SAME IndexSearchOp = 19
	IndexSearchOp_INDEX_ADJACENT IndexSearchOp = 20
)
type SPTreeKind C.SPTreeKind
const (
	SPTreeKind_SPTREE_QUADTREE SPTreeKind = 0
	SPTreeKind_SPTREE_KDTREE SPTreeKind = 1
)
type MeosType C.MeosType
const (
	MeosType_T_UNKNOWN MeosType = 0
	MeosType_T_BOOL MeosType = 1
	MeosType_T_DATE MeosType = 2
	MeosType_T_DATEMULTIRANGE MeosType = 3
	MeosType_T_DATERANGE MeosType = 4
	MeosType_T_DATESET MeosType = 5
	MeosType_T_DATESPAN MeosType = 6
	MeosType_T_DATESPANSET MeosType = 7
	MeosType_T_DOUBLE2 MeosType = 8
	MeosType_T_DOUBLE3 MeosType = 9
	MeosType_T_DOUBLE4 MeosType = 10
	MeosType_T_FLOAT8 MeosType = 11
	MeosType_T_FLOATSET MeosType = 12
	MeosType_T_FLOATSPAN MeosType = 13
	MeosType_T_FLOATSPANSET MeosType = 14
	MeosType_T_INT4 MeosType = 15
	MeosType_T_INT4MULTIRANGE MeosType = 16
	MeosType_T_INT4RANGE MeosType = 17
	MeosType_T_INTSET MeosType = 18
	MeosType_T_INTSPAN MeosType = 19
	MeosType_T_INTSPANSET MeosType = 20
	MeosType_T_INT8 MeosType = 21
	MeosType_T_INT8MULTIRANGE MeosType = 52
	MeosType_T_INT8RANGE MeosType = 53
	MeosType_T_BIGINTSET MeosType = 22
	MeosType_T_BIGINTSPAN MeosType = 23
	MeosType_T_BIGINTSPANSET MeosType = 24
	MeosType_T_STBOX MeosType = 25
	MeosType_T_TBOOL MeosType = 26
	MeosType_T_TBOX MeosType = 27
	MeosType_T_TDOUBLE2 MeosType = 28
	MeosType_T_TDOUBLE3 MeosType = 29
	MeosType_T_TDOUBLE4 MeosType = 30
	MeosType_T_TEXT MeosType = 31
	MeosType_T_TEXTSET MeosType = 32
	MeosType_T_TFLOAT MeosType = 33
	MeosType_T_TIMESTAMPTZ MeosType = 34
	MeosType_T_TINT MeosType = 35
	MeosType_T_TSTZMULTIRANGE MeosType = 36
	MeosType_T_TSTZRANGE MeosType = 37
	MeosType_T_TSTZSET MeosType = 38
	MeosType_T_TSTZSPAN MeosType = 39
	MeosType_T_TSTZSPANSET MeosType = 40
	MeosType_T_TTEXT MeosType = 41
	MeosType_T_GEOMETRY MeosType = 42
	MeosType_T_GEOMSET MeosType = 43
	MeosType_T_GEOGRAPHY MeosType = 44
	MeosType_T_GEOGSET MeosType = 45
	MeosType_T_TGEOMPOINT MeosType = 46
	MeosType_T_TGEOGPOINT MeosType = 47
	MeosType_T_NPOINT MeosType = 48
	MeosType_T_NPOINTSET MeosType = 49
	MeosType_T_NSEGMENT MeosType = 50
	MeosType_T_TNPOINT MeosType = 51
	MeosType_T_POSE MeosType = 54
	MeosType_T_POSESET MeosType = 55
	MeosType_T_TPOSE MeosType = 56
	MeosType_T_CBUFFER MeosType = 57
	MeosType_T_CBUFFERSET MeosType = 58
	MeosType_T_TCBUFFER MeosType = 59
	MeosType_T_TGEOMETRY MeosType = 60
	MeosType_T_TGEOGRAPHY MeosType = 61
	MeosType_T_TRGEOMETRY MeosType = 62
	MeosType_T_JSONB MeosType = 63
	MeosType_T_JSONPATH MeosType = 64
	MeosType_T_JSONBSET MeosType = 65
	MeosType_T_TJSONB MeosType = 66
	MeosType_T_TBIGINT MeosType = 67
	MeosType_T_H3INDEX MeosType = 68
	MeosType_T_H3INDEXSET MeosType = 69
	MeosType_T_TH3INDEX MeosType = 70
	MeosType_T_QUADBIN MeosType = 71
	MeosType_T_QUADBINSET MeosType = 72
	MeosType_T_TQUADBIN MeosType = 73
	MeosType_T_PCPOINT MeosType = 74
	MeosType_T_PCPOINTSET MeosType = 75
	MeosType_T_TPCPOINT MeosType = 76
	MeosType_T_PCPATCH MeosType = 77
	MeosType_T_PCPATCHSET MeosType = 78
	MeosType_T_TPCPATCH MeosType = 79
	MeosType_T_TPCBOX MeosType = 80
	MeosType_T_RAQUET MeosType = 81
	MeosType_T_POSECHAIN MeosType = 82
	MeosType_T_POSECHAINSET MeosType = 83
	MeosType_T_TPOSECHAIN MeosType = 84
	MeosType_T_S2CELL MeosType = 85
	MeosType_T_S2CELLSET MeosType = 86
	MeosType_T_TS2CELL MeosType = 87
	MeosType_NUM_MEOS_TYPES MeosType = 88
)
type MeosOper C.MeosOper
const (
	MeosOper_UNKNOWN_OP MeosOper = 0
	MeosOper_EQ_OP MeosOper = 1
	MeosOper_NE_OP MeosOper = 2
	MeosOper_LT_OP MeosOper = 3
	MeosOper_LE_OP MeosOper = 4
	MeosOper_GT_OP MeosOper = 5
	MeosOper_GE_OP MeosOper = 6
	MeosOper_ADJACENT_OP MeosOper = 7
	MeosOper_UNION_OP MeosOper = 8
	MeosOper_MINUS_OP MeosOper = 9
	MeosOper_INTERSECT_OP MeosOper = 10
	MeosOper_OVERLAPS_OP MeosOper = 11
	MeosOper_CONTAINS_OP MeosOper = 12
	MeosOper_CONTAINED_OP MeosOper = 13
	MeosOper_SAME_OP MeosOper = 14
	MeosOper_LEFT_OP MeosOper = 15
	MeosOper_OVERLEFT_OP MeosOper = 16
	MeosOper_RIGHT_OP MeosOper = 17
	MeosOper_OVERRIGHT_OP MeosOper = 18
	MeosOper_BELOW_OP MeosOper = 19
	MeosOper_OVERBELOW_OP MeosOper = 20
	MeosOper_ABOVE_OP MeosOper = 21
	MeosOper_OVERABOVE_OP MeosOper = 22
	MeosOper_FRONT_OP MeosOper = 23
	MeosOper_OVERFRONT_OP MeosOper = 24
	MeosOper_BACK_OP MeosOper = 25
	MeosOper_OVERBACK_OP MeosOper = 26
	MeosOper_BEFORE_OP MeosOper = 27
	MeosOper_OVERBEFORE_OP MeosOper = 28
	MeosOper_AFTER_OP MeosOper = 29
	MeosOper_OVERAFTER_OP MeosOper = 30
	MeosOper_EVEREQ_OP MeosOper = 31
	MeosOper_EVERNE_OP MeosOper = 32
	MeosOper_EVERLT_OP MeosOper = 33
	MeosOper_EVERLE_OP MeosOper = 34
	MeosOper_EVERGT_OP MeosOper = 35
	MeosOper_EVERGE_OP MeosOper = 36
	MeosOper_ALWAYSEQ_OP MeosOper = 37
	MeosOper_ALWAYSNE_OP MeosOper = 38
	MeosOper_ALWAYSLT_OP MeosOper = 39
	MeosOper_ALWAYSLE_OP MeosOper = 40
	MeosOper_ALWAYSGT_OP MeosOper = 41
	MeosOper_ALWAYSGE_OP MeosOper = 42
	MeosOper_TEMPCONTAINS_OP MeosOper = 43
	MeosOper_TEMPCONTAINED_OP MeosOper = 44
)
type SkipListType C.SkipListType
const (
	SkipListType_SKIPLIST_TEMPORAL SkipListType = 0
	SkipListType_SKIPLIST_KEYVALUE SkipListType = 1
)
type NullHandleType C.nullHandleType
const (
	NullHandleType_NULL_INVALID NullHandleType = 0
	NullHandleType_NULL_ERROR NullHandleType = 1
	NullHandleType_NULL_JSON_NULL NullHandleType = 2
	NullHandleType_NULL_DELETE NullHandleType = 3
	NullHandleType_NULL_RETURN NullHandleType = 4
)
