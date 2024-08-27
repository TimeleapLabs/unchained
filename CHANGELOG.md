## v0.13.0 (2024-08-27)

### ✨ Features

- add optional tls support
- **func-config**: add config struct for registering functions >>> ⏰ 3h
- **rpc**: refactor rpc service to accept unix socket >>> ⏰ 6h
- **rpc**: add comments and re arrange codes >>> ⏰ 6h
- **rpc-runtime**: add some runtime and refactor rpc service >>> ⏰ 6h
- **handlers**: refactoring rpc handlers >>> ⏰ 4h
- **bls-test**: add unittest for bls and move bls sign to identity …
- **bls-test**: add unittest for bls and move bls sign to identity method >>> ⏰ 1.5h
- add the new schnorr based pos contract
- **identity**: add flag to get permission to write to secret file,…
- **identity**: add flag to get permission to write to secret file, to prevent losing keys. and some refactor >>> ⏰ 3h
- **identity**: merge and add unit test for identity >>> ⏰ 1h

### 🐛🚑️ Fixes

- **concept**: remove unneccery file >>> ⏰ 5m
- fix rpc issues
- small bug fixes
- small bug fixes with rpc
- filter out unavailable workers
- **lint**: fix linters >>> ⏰ 5m
- **eth-rpc**: fix problem of race condition for client list >>> ⏰ 1h
- **eth-rpc**: rename isExist to isFound >>> ⏰ 2m
- **eth-rpc**: fix problem of race condition for client list >>> ⏰ 1h
- **linters**: do linter fixes >>> ⏰ 2m
- **identity**: rename export signer function name >>> ⏰ 2m
- **identity**: rename export signer function name >>> ⏰ 2m
- **evmlog**: fix mutex problem >>> ⏰ 30m

### fix

- display help correctly when no arguments are entered

### ✅🤡🧪 Tests

- fix mock modules
- fix mock contracts

### 🎨🏗️ Style & Architecture

- **linter**: fix linters problems >>> ⏰ 3h
- fix lint

### 🔐🚧📈✏️ 💩👽️🍻💬🥚🌱🚩🥅🩺 Others

- add unchained AI plugins

### 🚨 Linting

- fix go mod tidy

## v0.12.0 (2024-04-25)

### ✨ Features

- **services**: add some tests to project >>> ⏰ 3h
- **pubsub**: ability to send messages based on channels and sub-channel subscribe >>> ⏰ 6h
- **pubsub**: add internal pubsub and lots of refactor >>> ⏰ 2d
- record consensus info on boolean records
- refactor consensus to save all signatures in db
- add slashing mechanism
- add the proof-of-stake eip712 struct type definitions and signing functions

### 🐛🚑️ Fixes

- fix identity key generation
- **flags**: make config flag unrequired >>> ⏰ 2m
- **flags**: make config flag unrequired >>> ⏰ 2m
- **correctness**: delete unused tests temprorary >>> ⏰ 2m
- **connection**: fix the problem of reconnecting to the broker >>> ⏰ 2h
- **linters**: fix linters >>> ⏰ 2m
- **models**: fix problem of deserializing sia >>> ⏰ 2h
- **badge**: move badger to services and add unit tests >>> ⏰ 30m
- **linters**: solve linters problems >>> ⏰ 1h
- **linters-services**: add a new linter and fix dup in services >>> ⏰ 30m
- **crypto**: fix linters about comments >>> ⏰ 2m
- **crypto**: fix unused assining in evm init >>> ⏰ 10m
- **quickstart**: fix wrong address in the text >>> ⏰ 2m
- **bls**: fix paths on bls >>> ⏰ 2m
- **bls**: fix import paths of bls package >>> ⏰ 2m

### ♻️ Refactorings

- **ctx**: implement ctx passing through project >>> ⏰ 2h
- **services**: capsulate services using interfaces >>> ⏰ 1h
- fix debounce and pre-hook issues
- **crypto**: move etherum to crypto and refactor crypto to make a identity manage >>> ⏰ 2h
- refactor the eip712 module to repository pattern

### feat

- fix problem of config load before logger
- organize cmds / bls / comments
- make complete DI and some fix

### fix

- linter
- remove internal from some urls
- merge
- move Gql from broker to consumer
- remove unneccessry files
- merger
- consumer for worker and some fixes
- ignore plugins if is null
- ignore plugins if is null
- merge
- linter
- remove go.work
- seprate cobra and app init
- linters and update quickstart
- remove scheduler from consumer \ move correctness to services
- change structure of configs
- extract siner from checkPublicKey
- extract siner from checkPublicKey
- merge

### 🎨🏗️ Style & Architecture

- added pre-commit hooks, commitizen, gitmoji, and a bunch of other cool stuff

### 💚👷 CI & Build

- remove commitizen-branch
- **docker**: fix docker build
- migrate to the new org name

### 📌➕⬇️ ➖⬆️ Dependencies

- update the proof-of-stake contract to the latest abi

### 📝💡 Documentation

- **changelog**: add changelog
- **crypto**: add some comments in crypto package >>> ⏰ 10m
- add documentation for installing pre-commit hooks

### 🚨 Linting

- fix unnecessary trailing new line

## v0.11.21 (2024-03-31)

### feat

- collecting all configs to one place and some refactoring
- update config templates
- collecting all configs to one place and some refactoring
- fix all linters and add ci to check linters

### fix

- format of nodes list
- dep cycle problem
- updating quickstart and config templates
- loggers of config file
- add idea directory in gitignore
- add idea directory in gitignore
- problem in wd in CI
- problem in wd in CI
- some fixes according to comments

## v0.11.21-alpha.5 (2024-03-24)

## v0.11.21-alpha.4 (2024-03-24)

## v0.11.21-alpha.3 (2024-03-24)

## v0.11.21-alpha.2 (2024-03-23)

### feat

- add ci for linters
- fix all linters and add ci to check linters
- add golang lint config file and some alise in makefile
- add golang lint config file and some alise in makefile

### fix

- some misspeling
- new line at end of file

## v0.11.21-alpha.1 (2024-03-23)

## v0.11.21-alpha.0 (2024-03-22)

## v0.11.20 (2024-03-22)

## v0.11.20-alpha.0 (2024-03-15)

## v0.11.19 (2024-03-12)

## v0.11.18 (2024-03-10)

## v0.11.17 (2024-03-08)

## v0.11.16 (2024-03-08)

## v0.11.15 (2024-03-04)

## v0.11.14 (2024-03-02)

## v0.11.13 (2024-03-01)

## v0.11.12 (2024-02-26)

## v0.11.11 (2024-02-26)

## v0.11.10 (2024-02-26)

## v0.11.9 (2024-02-24)

## v0.11.9-rc.6 (2024-02-24)

## v0.11.9-rc.5 (2024-02-24)

## v0.11.9-rc.4 (2024-02-24)

## v0.11.9-rc.3 (2024-02-24)

## v0.11.9-rc.2 (2024-02-24)

## v0.11.9-rc.1 (2024-02-24)

## v0.11.9-rc.0 (2024-02-24)

## v0.11.8 (2024-02-21)

## v0.11.7 (2024-02-20)

## v0.11.6 (2024-02-18)

## v0.11.5 (2024-02-17)

## v0.11.4 (2024-02-16)

## v0.11.3 (2024-02-07)

## v0.11.3-alpha.2 (2024-02-07)

## v0.11.3-alpha.1 (2024-02-07)

## v0.11.3-alpha.0 (2024-02-07)

## v0.11.2 (2024-02-07)

## v0.11.1 (2024-02-02)

## v0.11.0 (2024-02-01)

## v0.11.0-alpha.5 (2024-02-01)

## v0.11.0-alpha.4 (2024-02-01)

## v0.11.0-alpha.3 (2024-02-01)

## v0.11.0-alpha.2 (2024-02-01)

## v0.10.28 (2024-01-25)

## v0.10.27 (2024-01-23)

## v0.10.26 (2024-01-23)

## v0.10.25 (2024-01-23)

## v0.10.24 (2024-01-23)

## v0.10.23 (2024-01-23)

## v0.10.22 (2024-01-22)

## v0.10.21 (2024-01-21)

## v0.10.20 (2024-01-21)

## v0.10.19 (2024-01-20)

## v0.10.18 (2024-01-18)

## v0.10.17 (2024-01-18)

## v0.10.16 (2024-01-18)

## v0.10.15 (2024-01-18)

## v0.10.14 (2024-01-17)

## v0.10.13 (2024-01-17)

## v0.10.12 (2024-01-17)

## v0.10.11 (2024-01-17)

## v0.10.10 (2024-01-16)

## v0.10.10-rc.2 (2024-01-16)

## v0.10.10-rc.1 (2024-01-16)

## v0.10.10-rc.0 (2024-01-16)

## v0.10.9 (2024-01-15)

## v0.10.9-rc.14 (2024-01-15)

## v0.10.9-rc.13 (2024-01-15)

## v0.10.9-rc.12 (2024-01-15)

## v0.10.9-rc.11 (2024-01-15)

## v0.10.9-rc.10 (2024-01-15)

## v0.10.9-rc.9 (2024-01-15)

## v0.10.9-rc.8 (2024-01-15)

## v0.10.9-rc.7 (2024-01-15)

## v0.10.9-rc.6 (2024-01-15)

## v0.10.9-rc.5 (2024-01-15)

## v0.10.9-rc.4 (2024-01-15)

## v0.10.9-rc.3 (2024-01-15)

## v0.10.9-rc.2 (2024-01-15)

## v0.10.9-rc.1 (2024-01-14)

## v0.10.9-rc.0 (2024-01-14)

## v0.10.8 (2024-01-14)

## v0.10.8-rc.1 (2024-01-14)

## v0.10.8-rc.0 (2024-01-14)

## v0.10.7 (2024-01-13)

## v0.10.6 (2024-01-13)

## v0.10.5 (2024-01-13)

## v0.10.5-rc.1 (2024-01-13)

## v0.10.5-rc.0 (2024-01-13)

## v0.10.4 (2024-01-12)

## v0.10.3 (2024-01-11)

## v0.10.3-rc.5 (2024-01-11)

## v0.10.3-rc.4 (2024-01-11)

## v0.10.3-rc.3 (2024-01-11)

## v0.10.3-rc.2 (2024-01-11)

## v0.10.3-rc.1 (2024-01-11)

## v0.10.3-rc.0 (2024-01-11)

## v0.10.2 (2024-01-10)

## v0.10.1 (2024-01-10)

## v0.10.0 (2024-01-10)

## v0.9.13 (2024-01-08)

## v0.9.12 (2024-01-08)

## v0.9.11 (2024-01-07)

## v0.9.10 (2024-01-06)

## v0.9.9 (2024-01-06)

## v0.9.8 (2024-01-06)

## v0.9.7 (2024-01-05)

## v0.9.6 (2024-01-05)

## v0.9.5 (2024-01-05)

## v0.9.4 (2024-01-04)

## v0.9.3 (2024-01-04)

## v0.9.2 (2024-01-04)

## v0.9.1 (2024-01-04)

## v0.8.13 (2024-01-02)

## v0.8.12 (2024-01-01)

## v0.8.11 (2023-12-28)

## v0.8.10 (2023-12-25)

## v0.8.9 (2023-12-25)

## v0.8.8 (2023-12-25)

## v0.8.7 (2023-12-23)

## v0.8.6 (2023-12-22)

## v0.8.5 (2023-12-21)

## v0.8.4 (2023-12-20)

## v0.8.3 (2023-12-20)

## v0.8.2 (2023-12-18)

## v0.8.1 (2023-12-18)

## v0.8.0 (2023-12-18)

## v0.7.1 (2023-12-09)

## v0.7.0 (2023-12-09)

## v0.6.0 (2023-12-06)

## v0.5.8 (2023-12-04)

## v0.5.7 (2023-12-03)

## v0.5.6 (2023-12-03)

## v0.5.5 (2023-12-03)

## v0.5.4 (2023-12-02)

## v0.5.3 (2023-12-02)

## v0.5.2 (2023-12-01)

## v0.5.1 (2023-12-01)

## v0.5.0 (2023-12-01)

## v0.4.2 (2023-11-30)

## v0.4.1 (2023-11-29)

## v0.4.0 (2023-11-29)

### feat

- collecting all configs to one place and some refactoring
- update config templates
- collecting all configs to one place and some refactoring
- fix all linters and add ci to check linters
- add ci for linters
- fix all linters and add ci to check linters
- add golang lint config file and some alise in makefile
- add golang lint config file and some alise in makefile

### fix

- rename pos and add default path for secrets
- rename o to option
- remove bin file
- linters
- linters
- linters
- merge
- merge
- format of nodes list
- dep cycle problem
- updating quickstart and config templates
- loggers of config file
- add idea directory in gitignore
- add idea directory in gitignore
- deleted some unused
- move Gql from broker to consumer
- problem in wd in CI
- problem in wd in CI
- some fixes according to comments
- some misspeling
- new line at end of file

## v0.3.3 (2023-11-28)

## v0.3.2 (2023-11-28)

## v0.3.1 (2023-11-28)

## v0.3.0 (2023-11-28)

## v0.2.7 (2023-11-27)

## v0.2.6 (2023-11-27)

## v0.2.5 (2023-11-27)

## v0.2.4 (2023-11-27)

## v0.2.3 (2023-11-27)

## v0.2.2 (2023-11-27)

## v0.2.1 (2023-11-27)

## v0.2.0 (2023-11-27)

## v0.1.3 (2023-11-25)

## v0.1.2 (2023-11-25)

## v0.1.1 (2023-11-25)
