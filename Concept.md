
## Project Structure

The Project follows the Clean Architecture pattern. The project is divided into the following layers:

- Services: This layer contains the business logic of the application.
- Repository: This layer contains the data access logic of the application which can be implemented for postgres, mongodb, etc.
- Handler: This layer contains the adapters for websocket packet processing.

and the layout is inspired by this [repository](https://github.com/evrone/go-clean-template), also with some modifications.

## Federated Network

The Unchained network is a federated network that consists of three types of nodes: Broker, Worker, and Consumer. Anyone can run a new node and help the network to process the load and requests.

### Broker

The broker nodes are responsible for managing the network and routing the data to the correct worker and consumer nodes. The clients of brokers consist of following:

- Consumers: They listen to the data from the brokers and save or use them.
- Workers: They process the data and send them the brokers.
- Clients: They request to run a rpc function.

```mermaid
  graph TD;
    U([ٌClient]) <--> |Websocket| W[Server]
    X([ٌWorker]) <--> |Websocket| W[Server]
    V([Consumer]) <--> |Websocket| W[Server]
    W[Server] --> |Sign & Verify| T[Machine Identity]

```

### Worker

The workers are the nodes that process the data and send them to the broker. The workers can be a service that listens to the data from the blockchain, or a service that listens to the data from the broker. These services hold different business logics and will provide different services to the network.
```mermaid
  graph TD;
    Y[Ethereum Network] -->    A[EvmLog Scheduler]
    Y[Ethereum Network] -->    B[Uniswap Scheduler]
    A[EvmLog Scheduler] --> |Every x sec| W[Services]
    B[Uniswap Scheduler] --> |Every x sec| W[Services]

    V[Broker] <--> |Websocket| W[Services]
    W[Services] --> |Sign & Verify| T[Machine Identity]

```

### Consumer

The consumers are the nodes that listen to the data from the broker and save or use them. These data be saved in a database, or be served to the users by an API.

```mermaid
  graph TD;
    V[Broker's Update] --> |Websocket| W[Services]
    W[Services] --> |Sign & Verify| T[Machine Identity]
    W[Services] --> |Save Data| Z[Database]
```

## Identity and Security

In a network like Unchained we send and receive many data from other parties and we need to identify and validate the data sender. in other hand, we communicate with different blockchains and every node needs to keep the keys and addresses of the blockchain.

These keys are hold on a global machine identity and will used once the node wants to sign a message or verify a message.

### BLS

A BLS digital signature, also known as Boneh–Lynn–Shacham (BLS), is a cryptographic signature scheme which allows a user to verify that a signer is authentic. The scheme uses a bilinear pairing for verification, and signatures are elements of an elliptic curve group.

### Ethereum

A Key pair of private and public keys which represent the identity of the node in the Ethereum network. and address of smart contract which is used to sync the nodes together.