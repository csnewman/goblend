// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package runtime

import _ "unsafe"

//go:linkname goblend_cgo_setenv goblend_cgo_setenv
//go:linkname _cgo_setenv runtime._cgo_setenv
var goblend_cgo_setenv byte
var _cgo_setenv = &goblend_cgo_setenv

//go:linkname goblend_cgo_unsetenv goblend_cgo_unsetenv
//go:linkname _cgo_unsetenv runtime._cgo_unsetenv
var goblend_cgo_unsetenv byte
var _cgo_unsetenv = &goblend_cgo_unsetenv
