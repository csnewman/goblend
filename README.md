# goblend
CGO-less FFI for Go.

goblend uses the `.syso` trick from [GcToolchainTricks](https://go.dev/wiki/GcToolchainTricks), to enable the
embedding of object files using only the inbuilt Go linker. When combined with a small amount of trampoline assembly,
this allows programs to use pre-compiled code from another language without requiring a cross-compiler. 

The main motivation behind this library is to enable library authors to reuse native code, without forcing downstream
users to modify their build steps. Provided a user is on a supported platform, a goblend library can be used like any
other Go library, with a simple `go get $pkg`.

The main downside is that all non-Go code must be pre-compiled outside of `go build`, and for library authors, must be
commited to the repository. The friction of using native code is shifted from the end user to the library maintainers.

## Features
- CGO-less FFI calls and callbacks (`CGO_ENABLED=0`)
- Cross-compilation
  - e.g. Build for Windows from a Linux machine, without configuring a cross-compiler.
- CGO compatible (`CGO_ENABLED=1`)
- `runtime.iscgo == true` on all platforms
  - The runtime is built from the original CGO source (for `.S` and `.c` files) for all supported
    platforms.
  - Should be (virtually) indistinguishable from CGO.

## Supported platforms

- Linux: arm64
- Windows: amd64
- Darwin: arm64

## Example

See `examples/add` for a complete example.

To demonstrate the library is compatible:
```bash
$ go mod init example
$ go get github.com/csnewman/goblend/examples/add
$ go run github.com/csnewman/goblend/examples/add
2024/12/09 11:23:39 Add example
Hello world from Zig!
2024/12/09 11:23:39 123
2024/12/09 11:23:39 End of example
```

### Native code

The exported functions must have the following signature:

```zig
export fn add(ptr: *const anyopaque) callconv(.c) u32 {
    // code here
}
```

and should be compiled to a native object file (archives/`.a` and libraries/`.dll` are unsupported): 

```bash
zig build-obj -O ReleaseFast -target aarch64-linux-gnu --library c -femit-bin=core_linux_arm64.syso add.zig add2.zig etc
```

The native object file should have the `.syso` extension, so the Go linker detects them. On Windows and Linux you can
take advantage of partial linking to reduce the number of object files down to one per platform.

### Native imports

When building, you will likely experience errors such as:

```
github.com/csnewman/goblend/runtime(.text): unknown symbol memset in callarm64
github.com/csnewman/goblend/runtime(.text): unknown symbol memset in callarm64
github.com/csnewman/goblend/runtime(.text): relocation target memset not defined
```

This is caused by your native code importing other code, such as `memset` from `libc` in the above example. As such, you
will need to create a file per platform, such as `core_linux.go` that imports the necessary dependencies:

```Go
//go:build linux

package runtime

import "unsafe"

//go:cgo_import_dynamic memset memset "libc.so.6"

// More dependencies...
```

Note: On Windows, the situation is slightly more complex due to `__imp_{some name}` symbols being a pointer to a
function, while the `{some name}` symbol expects to be the function itself. Please see `WriteFile` in the `examples/add`
example.

### Calling from Go

First you will need to create a `trampoline` for each function you intend to call: (`core.s`)

```asm
#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT goblend_add(SB),NOSPLIT|NOFRAME,$0-0
  JMP add(SB)
  RET
```

Note, no parameter information is necessary here, as this is a simple `jmp`.

Then in a Go file, such as `core.go`, you will need to get a ptr to the trampoline:

```go
package main

import "unsafe"

//go:linkname add goblend_add
var add uintptr
var addPtr = uintptr(unsafe.Pointer(&add))
```

Finally, the native function can be called:

```go
ret := goblend.Call(addPtr, 0)
```

The `0` argument is the `ptr` to pass to the native function. You can pass and return any arbitrary data this way.

### Callbacks

The pointer to a function can be resolved via `goblend.FuncPtr`:

```go
func ExampleCallback(ptr uintptr) {
	// callback logic here
}

ptr := goblend.FuncPtr(ExampleCallback)
```

Callbacks must have the above signature and can be invoked via:

```zig
extern "C" fn _cgo_wait_runtime_init_done() callconv(.C) *const anyopaque;
extern "C" fn _cgo_release_context(v: *const anyopaque) callconv(.C) void;
extern "C" fn crosscall2(ptr: *const anyopaque, data: *const anyopaque, size: c_int, ctx: *const anyopaque) callconv(.C) void;

fn someExample(ptr: *const anyopaque) void {
    var value :u64 = 321;

    const ctx = _cgo_wait_runtime_init_done();
    crosscall2(ptr, &value, 0, ctx);
    _cgo_release_context(ctx);
}
```

The `ctx` setup and releasing is optional in most circumstances and instead `0`/`null` can be passed, however
[runtime.SetCgoTraceback](https://pkg.go.dev/runtime#SetCgoTraceback) will not work. `size` appears to be unused in the
cgo runtime and can be left as 0.
