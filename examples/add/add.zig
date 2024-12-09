extern "C" fn _cgo_wait_runtime_init_done() callconv(.C) *const anyopaque;
extern "C" fn _cgo_release_context(v: *const anyopaque) callconv(.C) void;
extern "C" fn crosscall2(ptr: *const anyopaque, data: *const anyopaque, size: c_int, ctx: *const anyopaque) callconv(.C) void;

export fn add(ptr: *const anyopaque) callconv(.c) u32 {
    printHello();

    const ctx = _cgo_wait_runtime_init_done();
    var d:u64 = 321;
    crosscall2(ptr, &d, 0, ctx);
    _cgo_release_context(ctx);

    return 123;
}

const std = @import("std");

pub fn printHello() void {
    const outw = std.io.getStdOut().writer();
    outw.print("Hello world from Zig!\n", .{}) catch return;
}
