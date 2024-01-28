const std = @import("std");
const clap = @import("clap");
const websocket = @import("websocket");
const ws = @import("net/ws.zig");
const uniswap = @import("plugins/uniswap.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.detectLeaks();
    const allocator = gpa.allocator();

    // First we specify what parameters our program can take.
    // We can use `parseParamsComptime` to parse a string into an array of `Param(Help)`
    const params = comptime clap.parseParamsComptime(
        \\-h, --help             Display this help and exit.
        \\-a, --actor            Run in Actor mode.
        \\-r, --reactor          Run in Reactor mode.
        \\
    );

    // Initialize our diagnostics, which can be used for reporting useful errors.
    // This is optional. You can also pass `.{}` to `clap.parse` if you don't
    // care about the extra information `Diagnostics` provides.
    var diag = clap.Diagnostic{};

    var res = clap.parse(clap.Help, &params, clap.parsers.default, .{
        .diagnostic = &diag,
        .allocator = gpa.allocator(),
    }) catch |err| {
        // Report useful error and exit
        diag.report(std.io.getStdErr().writer(), err) catch {};
        return err;
    };

    defer res.deinit();

    if (res.args.help != 0) {
        return clap.help(std.io.getStdErr().writer(), clap.Help, &params, .{});
    } else if (res.args.reactor != 0) {
        var context = ws.Context{ .allocator = allocator };
        try websocket.listen(ws.ServerHandler, allocator, &context, .{
            .port = 9223,
            .max_headers = 10,
            .address = "0.0.0.0",
        });
    } else {
        try uniswap.work(allocator);
    }
}
