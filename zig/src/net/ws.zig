const websocket = @import("websocket");
const std = @import("std");
const s2s = @import("s2s");
const bls = @import("../bls/bls.zig");
const uniswap = @import("../plugins/uniswap.zig");
const blst = @cImport(@cInclude("blst.h"));

pub const ClientHandler = struct {
    client: websocket.Client,

    pub fn init(allocator: std.mem.Allocator, host: []const u8, port: u16) !ClientHandler {
        return .{
            .client = try websocket.connect(allocator, host, port, .{}),
        };
    }

    pub fn deinit(self: *ClientHandler) void {
        self.client.deinit();
    }

    pub fn connect(self: *ClientHandler, path: []const u8) !void {
        try self.client.handshake(path, .{ .timeout_ms = 5000 });
        const thread = try self.client.readLoopInNewThread(self);
        thread.detach();
    }

    pub fn handle(_: ClientHandler, message: websocket.Message) !void {
        std.debug.print("Feedback received: {s}\n", .{message.data});
    }

    pub fn write(self: *ClientHandler, data: []u8) !void {
        return self.client.write(data);
    }

    pub fn close(_: ClientHandler) void {}
};

// Define a struct for "global" data passed into your websocket handler
// This is whatever you want. You pass it to `listen` and the library will
// pass it back to your handler's `init`. For simple cases, this could be empty
pub const Context = struct {
    allocator: std.mem.Allocator,
};

pub const ServerHandler = struct {
    conn: *websocket.Conn,
    context: *Context,

    pub fn init(
        h: websocket.Handshake,
        conn: *websocket.Conn,
        context: *Context,
    ) !ServerHandler {
        // `h` contains the initial websocket "handshake" request
        // It can be used to apply application-specific logic to verify / allow
        // the connection (e.g. valid url, query string parameters, or headers)

        _ = h; // we're not using this in our simple case

        return ServerHandler{
            .conn = conn,
            .context = context,
        };
    }

    // optional hook that, if present, will be called after initialization is complete
    pub fn afterInit(self: *ServerHandler) !void {
        _ = self;
    }

    pub fn handle(self: *ServerHandler, message: websocket.Message) !void {
        const data = message.data;
        var buffer = std.io.fixedBufferStream(data);
        const parsed: uniswap.PriceData = try s2s.deserialize(buffer.reader(), uniswap.PriceData);

        var msg = std.ArrayList(u8).init(self.context.allocator);
        defer msg.deinit();
        try s2s.serialize(msg.writer(), u256, parsed.price);

        var pk: blst.blst_p1_affine = undefined; // Buffer for serialized data
        _ = blst.blst_p1_deserialize(&pk, &parsed.pk[0]);

        var signature: blst.blst_p2_affine = undefined; // Buffer for serialized data
        _ = blst.blst_p2_deserialize(&signature, &parsed.signature[0]);

        const valid = bls.verify(
            pk,
            signature,
            msg.items,
        );

        if (valid) {
            const price_str = try uniswap.formatPrice(
                parsed.price,
                18,
                self.context.allocator,
            );

            std.debug.print("Received valid price ${s} at block {:}\n", .{
                price_str,
                parsed.block,
            });

            try self.conn.write("accepted");
        } else {
            try self.conn.write("not accepted");
        }
    }

    // called whenever the connection is closed, can do some cleanup in here
    pub fn close(_: *ServerHandler) void {}
};
