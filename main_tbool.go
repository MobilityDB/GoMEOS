package gomeos

/*
#include "meos.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

type TBoolInst struct {
	_inner *C.Temporal
}

func NewTBoolInst(tgmpi_in string) *TBoolInst {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tbool_in(c_tgmpi_in)
	g_tgmpi := &TBoolInst{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TBoolInst) TBoolOut() string {
	c_tgmpi_out := C.tbool_out(tgmpi._inner)
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func (tb *TBoolInst) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TBoolInst) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TBoolInst) IsTBool() bool {
	return true
}

func (tb *TBoolInst) String() string {
	return tb.TBoolOut()
}

func (tb *TBoolInst) Type() string {
	return "TBoolInst"
}

type TBoolSeq struct {
	_inner *C.Temporal
}

func NewTBoolSeq(tgmpi_in string) *TBoolSeq {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tbool_in(c_tgmpi_in)
	g_tgmpi := &TBoolSeq{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TBoolSeq) TBoolOut() string {
	c_tgmpi_out := C.tbool_out(tgmpi._inner)
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func (tb *TBoolSeq) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TBoolSeq) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TBoolSeq) IsTBool() bool {
	return true
}

func (tb *TBoolSeq) String() string {
	return tb.TBoolOut()
}

func (tb *TBoolSeq) Type() string {
	return "TBoolSeq"
}

func (tb *TBoolSeq) IsTSequence() bool {
	return true
}

type TBoolSeqSet struct {
	_inner *C.Temporal
}

func NewTBoolSeqSet(tgmpi_in string) *TBoolSeqSet {
	c_tgmpi_in := C.CString(tgmpi_in)
	defer C.free(unsafe.Pointer(c_tgmpi_in))
	c_tgmpi := C.tbool_in(c_tgmpi_in)
	g_tgmpi := &TBoolSeqSet{_inner: c_tgmpi}
	return g_tgmpi
}

func (tgmpi *TBoolSeqSet) TBoolOut() string {
	c_tgmpi_out := C.tbool_out(tgmpi._inner)
	defer C.free(unsafe.Pointer(c_tgmpi_out))
	tgmpi_out := C.GoString(c_tgmpi_out)
	return tgmpi_out
}

func (tb *TBoolSeqSet) Init(c_temp *C.Temporal) {
	tb._inner = c_temp
}

func (tb *TBoolSeqSet) Inner() *C.Temporal {
	return tb._inner
}

func (tb *TBoolSeqSet) IsTBool() bool {
	return true
}

func (tb *TBoolSeqSet) String() string {
	return tb.TBoolOut()
}

func (tb *TBoolSeqSet) Type() string {
	return "TBoolSeqSet"
}

// TBoolFromBaseTemporal Create a temporal Boolean from a Boolean value and the time frame of another temporal object.
func TBoolFromBaseTemporal[T1 Temporal, T2 Temporal](value bool, base T1, output T2) T2 {
	c_temp := C.tbool_from_base_temp(C.bool(value), base.Inner())
	output.Init(c_temp)
	return output
}

func TBoolIn[TB TBool](input string, output TB) TB {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_bool := C.tbool_in(c_str)
	output.Init(c_bool)
	return output
}

// TBoolFromMFJSON Return a temporal boolean from its MF-JSON representation
func TBoolFromMFJSON[TB TBool](input string, output TB) TB {
	c_str := C.CString(input)
	defer C.free(unsafe.Pointer(c_str))
	c_bool := C.tbool_from_mfjson(c_str)
	output.Init(c_bool)
	return output
}

func TBoolOut[TB TBool](tb TB) string {
	c_bool := C.tbool_out(tb.Inner())
	defer C.free(unsafe.Pointer(c_bool))
	bool_out := C.GoString(c_bool)
	return bool_out
}

// TBoolValueSet Return the array of base values of a temporal boolean
func TBoolValueSet[TB TBool](tb TB) ([]bool, error) {
	var count C.int

	// Call the C function
	cValues := C.tbool_values(tb.Inner(), &count)
	if cValues == nil {
		return nil, fmt.Errorf("failed to retrieve bool values")
	}

	// Convert the C bool array to a Go slice
	length := int(count)
	values := unsafe.Slice((*C.bool)(cValues), length)
	// Convert the C bool values to Go bool values
	goValues := make([]bool, length)
	for i := 0; i < length; i++ {
		goValues[i] = bool(values[i])
	}
	return goValues, nil
}

// TBoolStartValue Return the start value of a temporal boolean
func TBoolStartValue[TB TBool](tb TB) bool {
	cValue := C.tbool_start_value(tb.Inner())
	return bool(cValue)
}

// TBoolEndValue Return the end value of a temporal boolean
func TBoolEndValue[TB TBool](tb TB) bool {
	cValue := C.tbool_end_value(tb.Inner())
	return bool(cValue)
}

// TBoolValueAtTimestamp Return the value of a temporal boolean at a timestamptz
func TBoolValueAtTimestamp[TB TBool](tb TB, ts time.Time) bool {
	tboolinst, _ := TemporalToTBoolInst(TemporalAtTimestamptz(tb, ts))
	return TBoolStartValue(tboolinst)
}

// AlwaysEqTBoolBool Return true if a temporal boolean is always equal to a boolean
func AlwaysEqTBoolBool[TB TBool](tb TB, value bool) bool {
	return int(C.always_eq_tbool_bool(tb.Inner(), C.bool(value))) > 0
}

// EverEqTBoolBool Return true if a temporal boolean is always equal to a boolean
func EverEqTBoolBool[TB TBool](tb TB, value bool) bool {
	return int(C.ever_eq_tbool_bool(tb.Inner(), C.bool(value))) > 0
}

// TEqTBoolBool Return the temporal equality of a temporal boolean and a boolean
func TEqTBoolBool[TB TBool](tb TB, value bool) Temporal {
	c_temp := C.teq_tbool_bool(tb.Inner(), C.bool(value))
	return CreateTemporal(c_temp)
}

// TNEqTBoolBool Return the temporal equality of a temporal boolean and a boolean
func TNEqTBoolBool[TB TBool](tb TB, value bool) Temporal {
	c_temp := C.tne_tbool_bool(tb.Inner(), C.bool(value))
	return CreateTemporal(c_temp)
}

// TBoolAtValue Return a temporal boolean restricted to a boolean
func TBoolAtValue[TB TBool](tb TB, value bool) Temporal {
	c_tbools := C.tbool_at_value(tb.Inner(), C.bool(value))
	return CreateTemporal(c_tbools)
}

// TBoolMinusValue Return a temporal boolean restricted to the complement of a boolean
func TBoolMinusValue[TB TBool](tb TB, value bool) Temporal {
	c_tbools := C.tbool_minus_value(tb.Inner(), C.bool(value))
	return CreateTemporal(c_tbools)
}

// TAndTBoolBool Return the boolean and of a temporal boolean and a boolean
func TAndTBoolBool[TB TBool](tb TB, value bool) Temporal {
	c_tbools := C.tand_tbool_bool(tb.Inner(), C.bool(value))
	return CreateTemporal(c_tbools)
}

// TAndTBoolTBool Return the boolean and of the temporal booleans
func TAndTBoolTBool[TB TBool](tb TB, other TB) Temporal {
	c_tbools := C.tand_tbool_tbool(tb.Inner(), other.Inner())
	return CreateTemporal(c_tbools)
}

// TOrTBoolBool Return the boolean or of a temporal boolean and a boolean
func TOrTBoolBool[TB TBool](tb TB, value bool) Temporal {
	c_tbools := C.tor_tbool_bool(tb.Inner(), C.bool(value))
	return CreateTemporal(c_tbools)
}

// TOrTBoolTBool Return the boolean or of the temporal booleans
func TOrTBoolTBool[TB TBool](tb TB, other TB) Temporal {
	c_tbools := C.tor_tbool_tbool(tb.Inner(), other.Inner())
	return CreateTemporal(c_tbools)
}

// TNotTBool Return the boolean not of a temporal boolean
func TNotTBool[TB TBool](tb TB, output TB) TB {
	c_temp := C.tnot_tbool(tb.Inner())
	output.Init(c_temp)
	return output
}

// TBoolWhenTrue Return the time when the temporal boolean has value true
func TBoolWhenTrue[TB TBool](tb TB) TsTzSpanSet {
	tss := C.tbool_when_true(tb.Inner())
	return TsTzSpanSet{tss}
}

// TBoolInstMake Return a temporal boolean instant from a boolean and a timestamptz
func TBoolInstMake(value bool, base time.Time) *TBoolInst {
	res := C.tboolinst_make(C.bool(value), DatetimeToTimestamptz(base))
	return &TBoolInst{_inner: C.cast_tinstant_to_temporal(res)}
}

// TBoolSeqFromBaseTstzset Return a temporal boolean discrete sequence from a boolean and a timestamptz set
func TBoolSeqFromBaseTstzset(value bool, base TsTzSet) *TBoolSeq {
	res := C.tboolseq_from_base_tstzset(C.bool(value), base._inner)
	return &TBoolSeq{_inner: C.cast_tsequence_to_temporal(res)}
}

// TBoolSeqFromBaseTstzspan Return a temporal boolean sequence from a boolean and a timestamptz span
func TBoolSeqFromBaseTstzspan(value bool, base TsTzSpan) *TBoolSeq {
	res := C.tboolseq_from_base_tstzspan(C.bool(value), base._inner)
	return &TBoolSeq{_inner: C.cast_tsequence_to_temporal(res)}
}

// TBoolSeqSetFromBaseTstzspanset Return a temporal boolean sequence set from a boolean and a timestamptz span set
func TBoolSeqSetFromBaseTstzspanset(value bool, base *TsTzSpanSet) *TBoolSeqSet {
	res := C.tboolseqset_from_base_tstzspanset(C.bool(value), base._inner)
	return &TBoolSeqSet{_inner: C.cast_tsequenceset_to_temporal(res)}
}
