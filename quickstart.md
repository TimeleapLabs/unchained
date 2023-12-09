# Quick Start

## Prerequisites

To run a Kenshi Unchained validator, you need to install Node.js. Follow the
installation instructions for your platform on the Node.js
[official installation](https://nodejs.org/en/download/package-manager) page.

## Install the Unchained Client

The easiest way to get started is to install the Unchained client globally.
On windows, Linux, MacOS, and BSD, run this command in CMD or terminal:

```bash
npm i -g @kenshi.io/unchained
```

Note: On UNIX-like operating systems, you might need to run this command with
`sudo`:

```bash
sudo npm i -g @kenshi.io/unchained
```

### Updates

To update the Unchained client, you can re-run the installation command above.
Adding `@latest` to the end would result in installing the latest version.

```bash
sudo npm i -g @kenshi.io/unchained@latest
```

## MongoDB

Note: Skip this step if you're planning to run a lite node.

To run a full node, you need to have an instance of MongoDB. You can either run
your own, or make a subscription to a cloud service. You can get a free
subscription for the Unchained testnet on
[MongoDB Atlas](https://www.mongodb.com/pricing) or
[OVH](https://www.ovhcloud.com/en/public-cloud/mongodb/). You can also get $200
free credits on [Digital Ocean](https://try.digitalocean.com/freetrialoffer/).

Contact us on [Telegram](https://t.me/kenshi) if you need help with this step.

## Configuration

You need a configuration file to get started. You can start with the following
config:

```yaml
log: info
name: Change me
lite: true
rpc:
  ethereum: https://ethereum.publicnode.com
database:
  url: mongodb+srv://<user>:<password>@<url>/?retryWrites=true&w=majority
  name: unchained
```

Save the above configuration in a file named `conf.yaml` on your system and make
the following modifications if required:

- `log`: Defines the validator log level. Change it to `silly` or `debug` to see
  all messages. Leaving this at `info` level gives you all the important
  messages.
- `name`: This name will be associated with your validator node, and is published to
  all peers.
- `lite`: To run a lite node, set this to `true`, otherwise set it to `false`.
- `rpc.ethereum`: You need to modify the `ethereum` RPC address to the one of your
  choice. You can find a list of Ethereum RPC nodes on
  [Chainlist](https://chainlist.org/chain/1).
- `database.url`: Your
  [MongoDB connection string](https://www.mongodb.com/docs/manual/reference/connection-string/)
  goes here. Ignore this if you're running a lite node.
- `database.name`: Name of the database to choose. If unsure, leave the default
  value. Ignore this if you're running a lite node.

You can also use RPC nodes that start with `wss://` instead of `https://`.

## Starting the Unchained Validator

To start the validator and join the Unchained network, you need to run the
following command (in CMD or terminal, depending on your OS) in the directory
where you saved the above configuration file:

```bash
unchained start conf.yaml
```

If you are running the `start` command for the first time, you also need to pass `--generate` to generate a random secret key. This key will be saved to the configuraion file and you won't have to generate a new key every time.

```bash
unchained start conf.yaml --generate
```

## Help

Running the following command in CMD or terminal should give you a list of
available options for the validator CLI tool:

```bash
unchained help
```

If you need more help, reach out on [Telegram](https://t.me/kenshi).
