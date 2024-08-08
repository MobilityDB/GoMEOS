package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>

TInstant *cast_temporal_to_tinstant(Temporal *temp) {
	TInstant *tinst = (TInstant *) temp;
	return tinst;
}

Temporal *cast_tsequence_to_temporal(TSequence *tseq) {
	Temporal *temp = (Temporal *) tseq;
	return temp;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// ------------------------- TGeogPointInst ---------------------------
type TGeogPointInst struct {
	_inner *C.Temporal
}

func NewTGeogPointInst(tgmpi_in string) TGeogPointInst {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeogpoint_in(c_tgmpi_in)
	g_tgmpi := TGeogPointInst{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeogPointInst) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func TGeogPointSeqMake(instants []TGeogPointInst, count int, lower_inc bool, upper_inc bool, interp string, normalize bool) TGeogPointSeq {
	interpEnum, exists := interpTypeMap[interp]
	if !exists {
		fmt.Printf("Unknown interpolation type: %s\n", interp)
		return TGeogPointSeq{_inner: nil}
	}
	var c_instants []*C.TInstant
	// Fill the C array
	for _, inst := range instants {
		tinst := C.cast_temporal_to_tinstant(inst._inner)
		c_instants = append(c_instants, tinst)
	}
	var input **C.TInstant = &(c_instants[0])
	c_geogseq := C.tsequence_make(input, C.int(count), C.bool(lower_inc), C.bool(upper_inc), interpEnum, C.bool(normalize))
	if c_geogseq == nil {
		fmt.Println("tsequence_make returned NULL")
		return TGeogPointSeq{_inner: nil}
	}
	c_geotemp := C.cast_tsequence_to_temporal(c_geogseq)
	geog_seq := TGeogPointSeq{_inner: c_geotemp}
	return geog_seq
}

func TFloatSeqMake(instants []TFloatInst, count int, lower_inc bool, upper_inc bool, interp string, normalize bool) TFloatSeq {
	interpEnum, exists := interpTypeMap[interp]
	if !exists {
		fmt.Printf("Unknown interpolation type: %s\n", interp)
		return TFloatSeq{_inner: nil}
	}
	var c_instants []*C.TInstant
	// Fill the C array
	for _, inst := range instants {
		tinst := C.cast_temporal_to_tinstant(inst._inner)
		c_instants = append(c_instants, tinst)
	}
	var input **C.TInstant = &(c_instants[0])
	c_geogseq := C.tsequence_make(input, C.int(count), C.bool(lower_inc), C.bool(upper_inc), interpEnum, C.bool(normalize))
	if c_geogseq == nil {
		fmt.Println("tsequence_make returned NULL")
		return TFloatSeq{_inner: nil}
	}
	c_geotemp := C.cast_tsequence_to_temporal(c_geogseq)
	geog_seq := TFloatSeq{_inner: c_geotemp}
	return geog_seq
}

func TPointLength(p TGeogPointSeq) float64 {
	return float64(C.tpoint_length(p._inner))
}

func TemporalNumInstants[T Temporal](temp T) int {
	return int(C.temporal_num_instants(temp.Inner()))
}

// ------------------------- TGeogPointSeq ---------------------------
type TGeogPointSeq struct {
	_inner *C.Temporal
}

func NewTGeogPointSeq(tgmpi_in string) TGeogPointSeq {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeogpoint_in(c_tgmpi_in)
	g_tgmpi := TGeogPointSeq{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeogPointSeq) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func (tgmpi TGeogPointSeq) Inner() *C.Temporal {
	return tgmpi._inner
}

// ------------------------- TGeogPointSeqSet ---------------------------
type TGeogPointSeqSet struct {
	_inner *C.Temporal
}

func NewTGeogPointSeqSet(tgmpi_in string) *TGeogPointSeqSet {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tgeogpoint_in(c_tgmpi_in)
	g_tgmpi := &TGeogPointSeqSet{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TGeogPointSeqSet) TPointOut(maxdd int) string {
	c_tgmpi_out := C.tpoint_as_text(tgmpi._inner, C.int(maxdd))
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}
