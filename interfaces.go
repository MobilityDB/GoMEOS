package gomeos

/*
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"
import "fmt"

type Temporal interface {
	Inner() *C.Temporal
	Init(*C.Temporal)
	String() string
	Type() string
}

type TInstant interface {
	Temporal
	IsTInstant() bool
}

type TSequence interface {
	Temporal
	IsTSequence() bool
}

type TSequenceSet interface {
	Temporal
	IsTSequenceSet() bool
}

type TBool interface {
	Temporal
	IsTBool() bool
}

type TNumber interface {
	Temporal
	IsTNumber() bool
}

type TInt interface {
	TNumber
	IsTInt() bool
}

type TFloat interface {
	TNumber
	IsTFloat() bool
}

type TText interface {
	Temporal
	IsTText() bool
}

type TPoint interface {
	Temporal
	IsTPoint() bool
}

type TGeomPoint interface {
	TPoint
	IsTGeomPoint() bool
}

type TGeogPoint interface {
	TPoint
	IsTGeogPoint() bool
}

type Geo interface {
	Inner() *C.GSERIALIZED
	IsGeo() bool
}

type Set interface {
	Inner() *C.Set
	Init(*C.Set)
}

type Span interface {
	Inner() *C.Span
	Init(*C.Span)
}

type SpanSet interface {
	Inner() *C.SpanSet
	Init(*C.SpanSet)
}

// ------------------------- Factory ----------------------------------
func CreateTemporal(inner *C.Temporal) Temporal {
	meosType := inner.temptype
	subtype := inner.subtype
	// meosType MeosType, subtype MeosTemporalSubtype
	switch meosType {
	case C.T_TBOOL:
		switch subtype {
		case C.TINSTANT:
			return &TBoolInst{_inner: inner}
		case C.TSEQUENCE:
			return &TBoolSeq{_inner: inner}
		case C.TSEQUENCESET:
			return &TBoolSeqSet{_inner: inner}
		}
	case C.T_TINT:
		switch subtype {
		case C.TINSTANT:
			return &TIntInst{_inner: inner}
		case C.TSEQUENCE:
			return &TIntSeq{_inner: inner}
		case C.TSEQUENCESET:
			return &TIntSeqSet{_inner: inner}
		}
	case C.T_TFLOAT:
		switch subtype {
		case C.TINSTANT:
			return &TFloatInst{_inner: inner}
		case C.TSEQUENCE:
			return &TFloatSeq{_inner: inner}
		case C.TSEQUENCESET:
			return &TFloatSeqSet{_inner: inner}
		}
	case C.T_TTEXT:
		switch subtype {
		case C.TINSTANT:
			return &TTextInst{_inner: inner}
		case C.TSEQUENCE:
			return &TTextSeq{_inner: inner}
		case C.TSEQUENCESET:
			return &TTextSeqSet{_inner: inner}
		}
	case C.T_TGEOMPOINT:
		switch subtype {
		case C.TINSTANT:
			return &TGeomPointInst{_inner: inner}
		case C.TSEQUENCE:
			return &TGeomPointSeq{_inner: inner}
		case C.TSEQUENCESET:
			return &TGeomPointSeqSet{_inner: inner}
		}
	case C.T_TGEOGPOINT:
		switch subtype {
		case C.TINSTANT:
			return &TGeogPointInst{_inner: inner}
		case C.TSEQUENCE:
			return &TGeogPointSeq{_inner: inner}
		case C.TSEQUENCESET:
			return &TGeogPointSeqSet{_inner: inner}
		}
	}
	return nil // or handle error
}

func TemporalToTBoolInst[T Temporal](temp T) (*TBoolInst, error) {
	t_type := temp.Type()
	if t_type == "TBoolInst" {
		return &TBoolInst{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TBoolInst")
	}
}

func TemporalToTBoolSeq[T Temporal](temp T) (*TBoolSeq, error) {
	t_type := temp.Type()
	if t_type == "TBoolSeq" {
		return &TBoolSeq{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TBoolSeq")
	}
}

func TemporalToTBoolSeqSet[T Temporal](temp T) (*TBoolSeqSet, error) {
	t_type := temp.Type()
	if t_type == "TBoolSeqSet" {
		return &TBoolSeqSet{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TBoolSeqSet")
	}
}

func TemporalToTIntInst[T Temporal](temp T) (*TIntInst, error) {
	t_type := temp.Type()
	if t_type == "TIntInst" {
		return &TIntInst{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TIntInst")
	}
}

func TemporalToTIntSeq[T Temporal](temp T) (*TIntSeq, error) {
	t_type := temp.Type()
	if t_type == "TIntSeq" {
		return &TIntSeq{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TIntSeq")
	}
}

func TemporalToTIntSeqSet[T Temporal](temp T) (*TIntSeqSet, error) {
	t_type := temp.Type()
	if t_type == "TIntSeqSet" {
		return &TIntSeqSet{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TIntSeqSet")
	}
}

func TemporalToTFloatInst[T Temporal](temp T) (*TFloatInst, error) {
	t_type := temp.Type()
	if t_type == "TFloatInst" {
		return &TFloatInst{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TFloatInst")
	}
}

func TemporalToTFloatSeq[T Temporal](temp T) (*TFloatSeq, error) {
	t_type := temp.Type()
	if t_type == "TFloatSeq" {
		return &TFloatSeq{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TFloatSeq")
	}
}

func TemporalToTFloatSeqSet[T Temporal](temp T) (*TFloatSeqSet, error) {
	t_type := temp.Type()
	if t_type == "TFloatSeqSet" {
		return &TFloatSeqSet{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TFloatSeqSet")
	}
}

func TemporalToTTextInst[T Temporal](temp T) (*TTextInst, error) {
	t_type := temp.Type()
	if t_type == "TTextInst" {
		return &TTextInst{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TTextInst")
	}
}

func TemporalToTTextSeq[T Temporal](temp T) (*TTextSeq, error) {
	t_type := temp.Type()
	if t_type == "TTextSeq" {
		return &TTextSeq{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TTextSeq")
	}
}

func TemporalToTTextSeqSet[T Temporal](temp T) (*TTextSeqSet, error) {
	t_type := temp.Type()
	if t_type == "TTextSeqSet" {
		return &TTextSeqSet{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TTextSeqSet")
	}
}

func TemporalToGeomPointInst[T Temporal](temp T) (*TGeomPointInst, error) {
	t_type := temp.Type()
	if t_type == "TGeomPointInst" {
		return &TGeomPointInst{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TGeomPointInst")
	}
}

func TemporalToTGeomPointSeq[T Temporal](temp T) (*TGeomPointSeq, error) {
	t_type := temp.Type()
	if t_type == "TGeomPointSeq" {
		return &TGeomPointSeq{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TGeomPointSeq")
	}
}

func TemporalToTGeomPointSeqSet[T Temporal](temp T) (*TGeomPointSeqSet, error) {
	t_type := temp.Type()
	if t_type == "TGeomPointSeqSet" {
		return &TGeomPointSeqSet{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TGeomPointSeqSet")
	}
}

func TemporalToTGeogPointInst[T Temporal](temp T) (*TGeogPointInst, error) {
	t_type := temp.Type()
	if t_type == "TGeogPointInst" {
		return &TGeogPointInst{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TGeogPointInst")
	}
}

func TemporalToTGeogPointSeq[T Temporal](temp T) (*TGeogPointSeq, error) {
	t_type := temp.Type()
	if t_type == "TGeogPointSeq" {
		return &TGeogPointSeq{_inner: temp.Inner()}, nil
	} else {
		return nil, fmt.Errorf("Temporal is not TGeogPointSeq")
	}
}
