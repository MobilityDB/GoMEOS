package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
TInstant *cast_temporal_to_tinstant(Temporal *temp) {
	TInstant *tinst = (TInstant *) temp;
	return tinst;
}

Temporal *cast_tinstant_to_temporal(TInstant *tinst) {
	Temporal *temp = (Temporal *) tinst;
	return temp;
}

Temporal *cast_tsequence_to_temporal(TSequence *tseq) {
	Temporal *temp = (Temporal *) tseq;
	return temp;
}

GSERIALIZED *cast_pointer_to_geo(Datum *p) {
	GSERIALIZED *gs = (GSERIALIZED *) p;
	return gs;
}
*/
import "C"
import "time"

type Temporal interface {
	Inner() *C.Temporal
	Init(*C.Temporal)
}

func (tgmpi *TGeomPointSeq) Inner() *C.Temporal {
	return tgmpi._inner
}

func (tgmpi *TGeomPointSeq) Init(c_temp *C.Temporal) {
	tgmpi._inner = c_temp
}

type TInstant interface {
	Inner() *C.Temporal
	TemporalToInst() *C.TInstant
	InstToTemporal(*C.TInstant) *C.Temporal
	Init(*C.Temporal)
}

func (tgmpi *TGeomPointInst) Inner() *C.Temporal {
	return tgmpi._inner
}

func (tgmpi *TGeomPointInst) TemporalToInst() *C.TInstant {
	return C.cast_temporal_to_tinstant(tgmpi._inner)
}

func (tgmpi *TGeomPointInst) InstToTemporal(c_inst *C.TInstant) *C.Temporal {
	return C.cast_tinstant_to_temporal(c_inst)
}

func (tgmpi *TGeomPointInst) Init(c_temp *C.Temporal) {
	tgmpi._inner = c_temp
}

func (tgmpi *TGeomPointInst) Timestamptz() time.Time {
	c_inst := C.cast_temporal_to_tinstant(tgmpi._inner)
	return TimestamptzToDatetime(c_inst.t)
}

func (tgmpi *TGeomPointInst) TimestampOut() string {
	c_inst := C.cast_temporal_to_tinstant(tgmpi._inner)
	return C.GoString(C.pg_timestamptz_out(c_inst.t))
}

type TGeo interface {
	Inner() *C.Temporal
}

func GeoAsEWKT[I TGeo](inst I, precision int) string {
	c_temp := inst.Inner()
	c_inst := C.cast_temporal_to_tinstant(c_temp)
	c_geo := C.cast_pointer_to_geo(&c_inst.value)
	return C.GoString(C.geo_as_ewkt(c_geo, C.int(precision)))
}

type TPoint interface {
	TPointOut() string
}
