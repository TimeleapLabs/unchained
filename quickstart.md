# Quick Start

## Prerequisites

To run a Kenshi Unchained validator, you need to install Node.js. Follow the
installation instructions for your platform on the Node.js
[official installation](https://nodejs.org/en/download/package-manager) page.

## Install the Unchained Software

The easiest way to get started is to install the Unchained software globally.
On windows, run this command in CMD:

```bash
npm i -g @kenshi.io/unchained
```

On Linux, MacOS, and BSD, run this command in terminal:

```bash
sudo npm i -g @kenshi.io/unchained
```

## Configuration

You need a configuration file to get started. You can start with the following
config:

```yaml
log: info
store: ~/.unchained
privateKey: dummy
rpc:
  eth: https://eth.llamarpc.com
```

Save the above configuration in a file named `conf.yaml` on your system and make
the following modifications if required:

- `log` defines the validator log level. Change it to `silly` or `debug` to see
  all messages. Leaving this at `info` level gives you all the important
  messages.
- `store` is where the validator stores the data it needs to operate. It needs
  to be writable. It's safe to leave it as is.
- `privateKey` is not used in this release, it can be any string value. Leave it
  as is.
- `rpc.eth`: You need to modify the `eth` RPC address from
  `https://eth.llamarpc.com` to the one of your choice. You can find a list of
  Ethereum RPC nodes on [Chainlist](https://chainlist.org/chain/1)

## Starting the Unchained Validator

To start the validator and join the Unchained network, you need to run the
following command (in CMD or terminal, depending on your OS) in the directory
where you saved the above configuration file:

```bash
unchained start conf.yaml
```

## Help

Running the following command in CMD or terminal should give you a list of
available options for the validator CLI tool:

```bash
unchained help
```

If you need more help, reach out on [Telegram](https://t.me/kenshi).
