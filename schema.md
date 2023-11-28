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
    "data": { // data to gossip
      "price": "2033.85831990604",
      "block": "18670546"
    },
    "parameters": { // gossip parameters
      "asset": "ethereum",
      "source": "uniswap",
      "chain": "ethereum"
    },
  },
  "signature": "aggregated signature of the request block",
  "signers": [
    // Array of public keys who have signed the packet
  ]
}
```
