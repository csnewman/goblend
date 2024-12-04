// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (linux && (amd64 || arm64 || loong64)) || (freebsd && amd64)

package runtime

// Import "unsafe" because we use go:linkname.
import _ "unsafe"

// When using cgo, call the C library for mmap, so that we call into
// any sanitizer interceptors. This supports using the memory
// sanitizer with Go programs. The memory sanitizer only applies to
// C/C++ code; this permits that code to see the Go code as normal
// program addresses that have been initialized.

// To support interceptors that look for both mmap and munmap,
// also call the C library for munmap.

//go:linkname goblend_cgo_mmap goblend_cgo_mmap
//go:linkname _cgo_mmap _cgo_mmap
var goblend_cgo_mmap byte
var _cgo_mmap = &goblend_cgo_mmap

//go:linkname goblend_cgo_munmap goblend_cgo_munmap
//go:linkname _cgo_munmap _cgo_munmap
var goblend_cgo_munmap byte
var _cgo_munmap = &goblend_cgo_munmap
