package goblend

import _ "unsafe"

//go:linkname runtime_cgocall runtime.cgocall
func runtime_cgocall(fn uintptr, arg uintptr) int32

func Call(fn uintptr, arg uintptr) int32 {
	return runtime_cgocall(fn, arg)
}
