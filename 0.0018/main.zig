const std = @import("std");

pub fn main() void {
    std.debug.print("Hello, world!\n", .{});
 const a: i32 = 5;
    const b: i32 = 7;
    const sum = a + b;

    std.debug.print("Sum = {}\n", .{sum});
}
