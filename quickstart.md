# Quick Start

Follow the guide on this page to run either a broker or a worker node:

- **Worker nodes** are nodes that only validate data, but they don't store any of
  the validated datapoints. It's easier to set up a worker node than a broker node.
- **Broker nodes** are nodes that validate data, and store the validated data for
  future queries. Broker nodes require extra steps to set up, and need more
  resources. You can either store the validated data on a local Postgres
  instance, or use DBaaS such as [Neon](https://neon.tech),
  [AWS RDS](https://aws.amazon.com/rds/), [ElephantSQL](https://www.elephantsql.com/)
  or any other.

For now, only worker nodes are implemented.

You can setup a node using a docker deployment or local deployment:

- **Docker deployment** requires Docker and makes the installation and
  management of the node straighforward.
- **Local deployment** runs Unchained locally on your machine and requires you
  to manually setup Unchained.

## Using Docker

Running Unchained with Docker is straightforward. Make sure you
[have Docker installed](https://docs.docker.com/engine/install/).
Then, head over to the
[Unchained release page](https://github.com/KenshiTech/unchained/releases)
on GitHub, find the latest Docker release file (file name ends with
`-docker.zip`), download it, and uncompress it.

Once done, head to the uncompressed release directory.

### Node types

The docker deployment is compatible with the 3 different node types:

- `broker`: Unchained and Postgres run in Docker.
- `worker`: Unchained runs in Docker, no need for Postgres.
- `remote`: Unchained runs in Docker, Postgres runs elsewhere.

Choose which node type you'd like to run on your machine.

### Configuring your node

#### Broker node

Make a copy of the environment template:

```bash
cp .env.template .env
```

Edit the newly created file with a username and password of your choice for Postgres
and Grafana.

Make a copy of the broker configuration template:

```bash
cp conf.broker.yaml.template conf.broker.yaml
```

Edit the newly created file. Set a name for your node. Set the Postgres
username, password, and database name to the ones you defined in the previous
step. Set the host of the Postgres local instance to `postgres`, and the port to 5432.

#### Worker node

Make a copy of the worker configuration template:

```bash
cp conf.worker.yaml.template conf.worker.yaml
```

Edit the newly created file and set a name for your node.

#### Remote node

Make a copy of the remote configuration template:

```bash
cp conf.remote.yaml.template conf.remote.yaml
```

Edit the newly created file. Set a name for your node. Set the Postgres username,
password, database name, host, and port to the ones of your DBaaS instance.

### Managing your node

To manage your node, use the following commands in your favorite terminal emulator.

#### Start Node

To start the node, run this command while in the release directory:

```bash
./unchained.sh [node] up -d
```

#### Stop Node

To stop the node, run this command while in the release directory:

```bash
./unchained.sh [node] stop
```

#### View Node

To view the node, run this command while in the release directory:

```bash
./unchained.sh [node] logs -f
```

#### Update Node

To update the node to the latest docker image version, run this command while in
the release directory:

```bash
./unchained.sh [node] pull
./unchained.sh [node] up -d --force-recreate
```

### Grafana Dashboard

Running a local node starts a Grafana dashboard instance on port 3000. You can
access it in your browser by visiting http://localhost:3000. You must set username
and password of your Grafana instance in the .env file mentioned above.

## Installing Locally

Follow these instructions if you want to install Unchained and its dependencies
locally, on a RaspberryPi, on a server, or on your computer/laptop.

### Prerequisites

To run a Kenshi Unchained node, you need to download the latest release. Head over to the
[Unchained release page](https://github.com/KenshiTech/unchained/releases)
on GitHub, find the latest release corresponding to your architecture and OS and download it.

On Unix-like operating systems, you'll first need to make the file executable:

```bash
chmod +x unchained.OS.ARCH 
```

#### Updates

To get the latest update of Unchained, you can download the latest release on the [Unchained release page](https://github.com/KenshiTech/unchained/releases).

### Postgres

Note: Skip this step if you're planning to run a worker node.

To run a broker node, you need to have an instance of Postgres. You can either run
your own, or make a subscription to a cloud service.

Contact us on [Telegram](https://t.me/kenshi) if you need help with this step.

#### Installing Postgres Locally

If you want to install Postgres locally, first follow the official Postgres
installation [instructions](https://www.postgresql.org/download/), then use the
following url in your Unchained config file:

```
postgres://<user>:<pass>@localhost:5432/<database>
```

Replace `<user>`, `<pass>`, and `<database>` with the user, password, and name
of the database you have created.

### Configuration

You need a configuration file to get started. You can start with the following
config:

```yaml
log: info
name: <name>
rpc:
  ethereum:
    - https://ethereum.publicnode.com
    - https://eth.llamarpc.com
    - wss://ethereum.publicnode.com
    - https://eth.rpc.blxrbdn.com
```

Save the above configuration in a file named `conf.yaml` on your system and make
the following modifications if required:

- `log`: Defines the validator log level. Change it to `silly` or `debug` to see
  all messages. Leaving this at `info` level gives you all the important
  messages.
- `name`: This name will be associated with your validator node, and is published to
  all peers.
- `rpc.ethereum`: Unchained testnet has automatic RPC rotation and renewal when
  issues are detected with the RPC connection. You can find a list of Ethereum
  RPC nodes on [Chainlist](https://chainlist.org/chain/1).

You can also use RPC nodes that start with `wss://` instead of `https://`.

## Migrations / Database Initialization

Note: Skip this step if you're running a worker node.

Before running the Unchained client, you need to get your database schema ready
for storing Unchained data. To do so, you should run:

```bash
unchained postgres migrate conf.yaml
```

You'll need to run this command again if you're installing a new version of the
client that makes changes to the Unchained data structure.

## Starting an Unchained worker node

To start a worker node and join the Unchained network, you need to run the
following command (in CMD or terminal, depending on your OS) in the directory
where you saved the above configuration file:

```bash
unchained.OS.ARCH worker
```

If you are running the node for the first time, Unchained will generate a random secret key. This key will be saved to the
secrets.yaml file. It is your responsibility to keep this file safe.

## Max Open Sockets

Depending on your OS and OS configuration, you might run into issues if you have
too many peers connected. Follow the guides below to increase the maximum open
connections limit on your OS.

### MacOS

To increase the limit on MacOS, run these commands:

```bash
sudo sysctl kern.maxfiles=2000000 kern.maxfilesperproc=2000000
echo "ulimit -Hn 2000000" >> ~/.zshrc
echo "ulimit -Sn 2000000" >> ~/.zshrc
source ~/.zshrc
```

### Linux

To increase the limit on Linux, run these commands:

```bash
sudo sysctl -w fs.nr_open=33554432
echo "ulimit -Hn 33554432" >> ~/.bashrc
echo "ulimit -Sn 33554432" >> ~/.bashrc
source ~/.bashrc
```

## Help

Running the following command in CMD or terminal should give you a list of
available options for the validator CLI tool:

```bash
unchained help
```

If you need more help, reach out on [Telegram](https://t.me/kenshi).
