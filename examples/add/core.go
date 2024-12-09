package main

import "unsafe"

//go:linkname add goblend_add
var add uintptr
var addPtr = uintptr(unsafe.Pointer(&add))
