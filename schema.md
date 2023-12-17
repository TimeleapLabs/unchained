# Unchained Data Schema

## RPC

To call a method on a node, you can use the following payload schema:

```JSON
{
  "type": "call",
  "request":  {
     // Unique ID of this request (NanoID or UUIDv4 are recommended)
    "id": "qKcZZtdvgAo14PZP2BjRc9KrjxR3z4T2e8zs",
    "method": "getPriceOf", // Method name to call
    "args": { // Arguments to pass to the method
      "asset": "ethereum"
    }
  },
  "signature": "Signature of the request block",
  "signer": "Public key of the signer"
}
```

## Gossip

Gossip is a piece of data that is broadcast to all nodes. To gossip, use
the following payload schema:

```JSON
{
  "type": "gossip",
  "request": {
    "method": "assetPrice", // gossip method to call
    "dataset": "ethereum::uniswap::ethereum", // chain::source::asset
    "metric": { "block": "18670546" }, // M -> V metric
    "signature": "Singature" // Hash { metric, value }, then sign
  },
  "origin": "Public key of the signer",
  "seen": [] // Array of public keys who have seen the packet with this signature
}
```

Signature data format:

```JSON
{
  "metric": { "block": "18670546" },
  "value": { "price": "$2229.8986944638564" }
}
```
