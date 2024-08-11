

## Project layout

This project follow the clean layout structure inspired from https://github.com/evrone/go-clean-template with some modifications. 

## Available Services

- **Pos:** This service will handle the logics related to the POS smart contract. 
- **Correctness:** This service will handle the logics of incoming events from correctness events.
- **Evm Log:** This service will handle the logics of incoming events from EVM logs.
- **Uniswap:** This service will handle the logics of incoming price events from Uniswap.
- **Rpc:** This service consist of two coordinator and worker which make ability to call RPC functions from brokers.
    - The coordinator will handle the selection and handover requests and responses between workers.
    - The worker will register and run the functions which are requested by broker's clients.
    - Rpc service contains multiple runtimes which can help the unchained to run external applications.

## Workflow

```mermaid
  flowchart TD;
      A[-->B;
      A-->C;
      B-->D;
      C-->D;
```