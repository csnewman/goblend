export fn add(ptr: *const anyopaque) callconv(.c) u32 {
    _ = ptr;

    printHello();

    return 123;
}

const std = @import("std");

pub fn printHello() void {
    const outw = std.io.getStdOut().writer();
    outw.print("Hello world from Zig!\n", .{}) catch return;
}
