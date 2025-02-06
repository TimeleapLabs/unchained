# Timeleap Data Schema

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
