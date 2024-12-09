package goblend

import (
	"reflect"
	_ "unsafe"
)

//go:linkname runtime_cgocall runtime.cgocall
func runtime_cgocall(fn uintptr, arg uintptr) int32

// Call invokes the native function, passing a single argument.
//
// It is important that any pointers referenced from arg are alive and are valid throughout the duration of their use.
//
// [KeepAliveSpill] is the preferred choice for keeping pointers alive. It will cause the value to be stored on the
// meaning that the native code is allowed to perform Go callbacks or store the pointer (assuming the memory has been
// pinned with [runtime.Pinner]).
//
// [KeepAliveNoSpill] can be used in limited circumstance, where the function does not call Go code or attempt to store
// the value.
func Call(fn uintptr, arg uintptr) int32 {
	return runtime_cgocall(fn, arg)
}

// FuncPtr returns the address of the provided function, suitable for passing to native code.
func FuncPtr(f any) uintptr {
	return reflect.ValueOf(f).Pointer()
}

// AlwaysFalse is always false.
//
// This value is hidden from the compiler, so code using it will not be eliminated.
//
//go:linkname AlwaysFalse runtime.cgoAlwaysFalse
var AlwaysFalse bool

//go:linkname runtime_cgo_use runtime.cgoUse
func runtime_cgo_use(interface{})

// TODO: replace runtime_cgo_use_noescape with runtime.cgoKeepAlive in Go 1.24.

// TODO: expose no callback helpers in Go 1.24.

//go:linkname runtime_cgo_use_noescape runtime.cgoUse
//go:noescape
func runtime_cgo_use_noescape(interface{})

// KeepAliveSpill will keep the value alive until this point and will cause the value to be stored on the heap.
func KeepAliveSpill(v any) {
	if AlwaysFalse {
		runtime_cgo_use(v)
	}
}

// KeepAliveNoSpill will keep the value alive until this point without causing the value to be stored on the heap.
func KeepAliveNoSpill(v any) {
	if AlwaysFalse {
		runtime_cgo_use_noescape(v)
	}
}

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
//go:noescape
func _cgoCheckPointer(any, any)

// CheckPointer checks if the argument contains a Go pointer that points to an unpinned Go pointer, and panics if it
// does. See [runtime.cgoCheckPointer] for more details.
//
// NOTE: This may not yet be working.
func CheckPointer(ptr any, arg any) {
	_cgoCheckPointer(ptr, arg)
}

//go:linkname _cgoCheckResult runtime.cgoCheckResult
//go:noescape
func _cgoCheckResult(any)

// CheckResult checks the result parameter of an exported Go function. It panics if the result is or contains any other
// pointer into unpinned Go memory. See [runtime.cgoCheckResult] for more details.
//
// NOTE: This may not yet be working.
func CheckResult(ptr any) {
	_cgoCheckResult(ptr)
}
