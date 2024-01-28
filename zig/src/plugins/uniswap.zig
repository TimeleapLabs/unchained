const bls = @import("../bls/bls.zig");
const web3 = @import("web3");
const blst = @cImport(@cInclude("blst.h"));
const std = @import("std");
const clap = @import("clap");
const ws = @import("../net/ws.zig");
const s2s = @import("s2s");

const addr = web3.Address.fromString("0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640") catch unreachable; // Binance wallet

const UniV3 = struct {
    const Self = @This();

    contract: web3.ContractCaller,

    pub fn init(allocator: std.mem.Allocator, address: web3.Address, provider: web3.Provider) UniV3 {
        return UniV3{
            .contract = web3.ContractCaller.init(allocator, address, provider),
        };
    }

    pub fn sqrtPriceX96(self: Self, opts: web3.CallOptions) !u256 {
        const selector = comptime try web3.abi.computeSelectorFromSig("slot0()");
        return self.contract.callSelector(selector, .{}, u256, opts);
    }

    pub fn price(self: Self, opts: web3.CallOptions) !u256 {
        const fetchedSqrtPriceX96 = try self.sqrtPriceX96(opts);
        const raw = (std.math.pow(u256, fetchedSqrtPriceX96, 2)) / (std.math.pow(u256, 2, 192)) * std.math.pow(u256, 10, 6);
        const inverse = true;
        if (inverse) {
            return std.math.pow(u256, 10, 36) / raw;
        }
        return raw / std.math.pow(u256, 10, 18);
    }
};

pub fn formatPrice(price: u256, decimal_places: usize, allocator: std.mem.Allocator) ![]u8 {
    var buffer: [100]u8 = undefined;

    const value_str = try std.fmt.bufPrint(&buffer, "{}", .{price});

    const result_len = value_str.len + 1;
    var result = try allocator.alloc(u8, result_len);

    // Fill with zeros using @memset
    @memset(result[0..result_len], '0');

    const decimal_point_index = if (value_str.len >= decimal_places) value_str.len - decimal_places else 1;
    result[decimal_point_index] = '.';

    if (value_str.len < decimal_places) {
        std.mem.copyForwards(u8, result[1..], value_str);
    } else {
        std.mem.copyForwards(u8, result, value_str[0 .. value_str.len - decimal_places]);
        std.mem.copyForwards(u8, result[decimal_point_index + 1 ..], value_str[value_str.len - decimal_places ..]);
    }

    return result;
}

pub const PriceData = struct {
    price: u256,
    block: u256,
    signature: [192]u8,
    pk: [96]u8,
};

pub fn work(allocator: std.mem.Allocator) !void {
    const rpc_endpoint = "https://eth.llamarpc.com";
    var json_rpc_provider = try web3.JsonRpcProvider.init(allocator, try std.Uri.parse(rpc_endpoint));
    defer json_rpc_provider.deinit();

    const univ3 = UniV3.init(allocator, addr, json_rpc_provider.provider());

    var handler = try ws.ClientHandler.init(allocator, "65.108.48.32", 9223);
    defer handler.deinit();

    // spins up a thread to listen to new messages
    try handler.connect("/");

    const keyPair = bls.keygen();
    var pk_buf: [96]u8 = undefined; // Buffer for serialized data
    blst.blst_p1_serialize(pk_buf[0..], &keyPair.pk);
    var lastBlock: u256 = 0;

    while (true) {
        // Get current block number
        const block_number = try json_rpc_provider.getBlockNumber();

        if (block_number != lastBlock) {
            const price = try univ3.price(.{});

            const float_repr = try formatPrice(price, 18, std.heap.page_allocator);
            defer std.heap.page_allocator.free(float_repr);

            std.debug.print("Eth price: ${s} at block: {:}\n", .{ float_repr, block_number });
            lastBlock = block_number;

            var msg = std.ArrayList(u8).init(allocator);
            defer msg.deinit();
            try s2s.serialize(msg.writer(), u256, price);

            const signature = bls.sign(keyPair, msg.items);

            // Serialize the signature for printing
            var sig_buffer: [192]u8 = undefined; // Buffer for serialized data
            blst.blst_p2_serialize(sig_buffer[0..], &signature);
            // std.debug.print("Signature: {s}\n", .{std.fmt.fmtSliceHexLower(&sig_buffer)});

            const priceData = PriceData{
                .price = price,
                .block = block_number,
                .signature = sig_buffer,
                .pk = pk_buf,
            };

            var payload = std.ArrayList(u8).init(allocator);
            defer payload.deinit();

            try s2s.serialize(payload.writer(), PriceData, priceData);

            const data = try allocator.dupe(u8, payload.items);
            try handler.write(data);
        }

        // without this, we'll exit immediately without having time to receive a
        // message from the server
        std.time.sleep(5 * std.time.ns_per_s);
    }
}
