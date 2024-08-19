package gomeos

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lmeos -Wl,-rpath,/opt/homebrew/lib
#include "meos.h"
#include "meos_catalog.h"
#include <stdio.h>
#include <stdlib.h>
#include "cast.h"
*/
import "C"

//	var interpTypeMap = map[string]C.interpType{
//		"INTERP_NONE": C.INTERP_NONE,
//		"DISCRETE":    C.DISCRETE,
//		"STEP":        C.STEP,
//		"LINEAR":      C.LINEAR,
//	}

// try godoc
type Interpolation C.int

const (
	INTERP_NONE    Interpolation = C.INTERP_NONE
	DISCRETE       Interpolation = C.DISCRETE
	STEP           Interpolation = C.STEP
	LINEAR                       = C.LINEAR
	ANYTEMPSUBTYPE               = C.ANYTEMPSUBTYPE /**< Any temporal subtype */
	TINSTANT                     = C.TINSTANT       /**< Temporal instant subtype */
	TSEQUENCE                    = C.TSEQUENCE      /**< Temporal sequence subtype */
	TSEQUENCESET                 = C.TSEQUENCESET   /**< Temporal sequence set subtype */
	T_TBOOL                      = C.T_TBOOL        /**< temporal boolean type */
	T_TFLOAT                     = C.T_TFLOAT       /**< temporal float type */
	T_TINT                       = C.T_TINT         /**< temporal integer type */
	T_TTEXT                      = C.T_TTEXT        /**< temporal text type */
	T_TGEOMPOINT                 = C.T_TGEOMPOINT   /**< temporal geometry point type */
	T_TGEOGPOINT                 = C.T_TGEOGPOINT   /**< temporal geography point type */
)

// func CreateTemporal(meosType MeosType, subtype MeosTemporalSubtype) func(string) *TBoolInst {
// 	switch meosType {
// 	case T_TBOOL:
// 		switch subtype {
// 		case TINSTANT:
// 			return NewTBoolInst
// 		case TSEQUENCE:
// 			return NewTBoolSeq
// 		case TSEQUENCESET:
// 			return NewTBoolSeqSet
// 		}
// 		return nil // or handle error
// 	}
// 	return nil
// }
