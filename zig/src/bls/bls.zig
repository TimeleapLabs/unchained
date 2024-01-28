const blst = @cImport(@cInclude("blst.h"));
const std = @import("std");

const crypto = std.crypto;

pub const KeyPair = struct {
    sk: blst.blst_scalar,
    pk: blst.blst_p1,
};

pub fn keygen() KeyPair {
    // Generate secret key (SK)
    var ikm: [32]u8 = undefined; // Create a buffer for the IKM
    crypto.random.bytes(&ikm); // Fill the buffer with random bytes

    var keyPair: KeyPair = undefined;

    blst.blst_keygen(&keyPair.sk, &ikm, ikm.len, null, 0);
    blst.blst_sk_to_pk_in_g1(&keyPair.pk, &keyPair.sk);

    var buffer: [96]u8 = undefined; // Buffer for serialized data
    blst.blst_p1_serialize(buffer[0..], &keyPair.pk);
    std.debug.print("Public Key (G1): {s}\n", .{std.fmt.fmtSliceHexLower(&buffer)});

    return keyPair;
}

pub fn sign(keyPair: KeyPair, msg: []u8) blst.blst_p2 {
    // Hash the message to a point on G1
    const dst: []const u8 = "BLS_SIG_BLS12381G1_XMD:SHA-256_SSWU_RO_NUL_"; // Domain separation tag
    var hashed_msg: blst.blst_p2 = undefined;
    blst.blst_hash_to_g2(&hashed_msg, msg.ptr, msg.len, dst.ptr, dst.len, null, 0);

    // Sign the message
    var signature: blst.blst_p2 = undefined;
    blst.blst_sign_pk_in_g1(&signature, &hashed_msg, &keyPair.sk);

    return signature;
}

pub fn verify(pk_g1_affine: blst.blst_p1_affine, signature_affine: blst.blst_p2_affine, msg: []u8) bool {
    const dst: []const u8 = "BLS_SIG_BLS12381G1_XMD:SHA-256_SSWU_RO_NUL_"; // Domain separation tag

    // Verify the signature
    // Initialize context for pairing-based verification
    var allocator = std.heap.page_allocator;

    // Get the size of blst_pairing structure
    const blst_pairing_size = blst.blst_pairing_sizeof();

    // Allocate memory for the blst_pairing context
    const ctx_ptr = allocator.alloc(u8, blst_pairing_size) catch {
        std.debug.panic("Failed to allocate memory for blst_pairing context\n", .{});
    };
    defer allocator.free(ctx_ptr);

    // Cast the allocated memory to a pointer to blst_pairing
    const ctx = @as(*blst.blst_pairing, @ptrCast(ctx_ptr.ptr));

    blst.blst_pairing_init(ctx, true, dst.ptr, dst.len); // true for hash, false for encode

    // Aggregate the public key and message
    _ = blst.blst_pairing_aggregate_pk_in_g1(ctx, &pk_g1_affine, null, msg.ptr, msg.len, null, 0);

    // Finalize the setup
    blst.blst_pairing_commit(ctx);

    // Prepare the signature for verification
    var gtsig: blst.blst_fp12 = undefined;
    blst.blst_aggregated_in_g2(&gtsig, &signature_affine);

    // Perform the final verification
    const result = blst.blst_pairing_finalverify(ctx, &gtsig);

    return result;
}
