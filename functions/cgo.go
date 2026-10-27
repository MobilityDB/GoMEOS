package functions

/*
// The installed libmeos describes itself: `pkg-config --cflags meos` reports
// its include dir together with the family macros it is compiled with, and
// `--libs` reports where to link it from. Both matter, because the public
// headers gate declarations on those macros (meos.h holds #if MEOS, #if
// POINTCLOUD and #if JSON blocks), so a hand-kept -D list that disagrees with
// the library either declares a symbol that cannot link or hides a family the
// library provides.
//
// Asking pkg-config also makes the binding relocatable: PKG_CONFIG_PATH selects
// which libmeos to build against, so a build points at a private prefix rather
// than inheriting whatever occupies a machine-wide directory. A prefix outside
// the loader's default search path needs LD_LIBRARY_PATH (or ldconfig).
#cgo pkg-config: meos

// H3 ships no pkg-config file, so the one dependency whose header the public
// MEOS headers expose (meos_h3.h includes <h3api.h>) names its include dir
// here.
#cgo darwin CFLAGS: -I/opt/homebrew/include/h3
#cgo linux CFLAGS: -I/usr/include/h3

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

import "fmt"

// MeosError reports that the MEOS call behind a wrapper failed, carrying the
// error code MEOS recorded for it.
type MeosError struct {
	Code int
}

func (e *MeosError) Error() string {
	return fmt.Sprintf("meos: error %d", e.Code)
}

// meosError reads the out-of-band state MEOS sets when a call fails, and
// clears it so a stale code cannot condemn the next, innocent call.  It
// returns nil when the call succeeded.
//
// The returned value can never answer this on its own: every MEOS sentinel is
// a legitimate value of its own return type -- INT_MAX is a count, DBL_MAX is
// a distance, and a bool has no spare value at all.  So no generated wrapper
// compares against a sentinel; each consults this instead.
//
// Install meos_initialize_noexit_error_handler (wrapped as
// MeosInitializeNoexitErrorHandler) once per process before calling anything
// else, on every thread that calls MEOS: the default handler ends the
// process, and meos_initialize resets the handler per thread.
func meosError() error {
	if code := int(C.meos_errno()); code != 0 {
		C.meos_errno_reset()
		return &MeosError{Code: code}
	}
	return nil
}
