
# C-CHAIN
Use this repo to create your first cosmos chain 

## Prerequisites
- [Ignite CLI](https://docs.ignite.com/network/validator)
- [Docker](https://www.docker.com/)
- [Cosmos SDK](https://tutorials.cosmos.network/)
- [Golang](https://go.dev/)

### Install Ignite CLI - docker

Build the following docker file. This creates an image that contains ignite and golang which we will be using for the rest of this setup

`docker build -f c-chain/prod/dockerfile-ignite . -t ignite`

### Create Cosmos Chain
Create tendermint chain by running the following.

`docker run --rm -it -v $(pwd):/ignite -w /ignite ignite ignite scaffold chain github.com/conare/c-chain`

This creates a Tendermint chain with a basic node out of the box. You can optimize `c-chain` application and run it in dev mode by running `docker run --rm -it -v $(pwd):/ignite -w /ignite ignite ignite chain serve`.

When you are ready to productionalize your application, proceed to the next section.

## Build Production Image
We now need to package the application into a docker image to make it ready for production use. Do this by running the following step

`docker build -f c-chain/dockerfile-prod . --build-arg BUILDARCH=amd64 -t c-chaind`

### Test Image
Test that the image can run
`docker run --rm -it c-chaindhelp`

You should see an output similar to the following:

```
 => docker run --rm -it c-chaind help
Stargate CosmosHub App

Usage:
  c-chaind [command]

Available Commands:
  add-genesis-account Add a genesis account to genesis.json
  collect-gentxs      Collect genesis txs and output a genesis.json file
  config              Create or query an application CLI configuration file
  debug               Tool for helping with debugging your application
```

## Validator Keys

Let's run KMS for our validators. we do this using docker by running the following:

`docker build -f prod-sim/dockerfile-tmkms . -t tmkms_i:v0.12.2`

### Initialize KMS

Let's initialize the KMS folder

```
docker run --rm -it \
    -v $(pwd)/prod/kms-chris:/root/tmkms \
    tmkms_i:v0.12.2 \
    init /root/tmkms
```

## Create Validators

Initialize validators by running:

```
echo -e val-alice'\n'val-bob'\n'desk-chris'\n'val-chris'\n'node-jane'\n'sentry-alice'\n'sentry-chris'\n'sentry-bob\
    | xargs -I {} \
    docker run --rm -i \
    -v $(pwd)/prod/{}:/root/.c-chain \
    c-chaind \
    init c-chaind
```

### Validators Keys

let's create keys for validators

#### Alice
```
docker run --rm -it \
    -v $(pwd)/prod/val-alice:/root/.c-chain \
    c-chaind \
    keys \
    --keyring-backend file --keyring-dir /root/.c-chain/keys \
    add alice
```
#### Chris
```
docker run --rm -it \
    -v $(pwd)/prod/desk-chris:/root/.c-chain \
    c-chaind \
    keys \
    --keyring-backend file --keyring-dir /root/.c-chain/keys \
    add chris
```

#### Import Consensus Key 

Import import val-chris's consensus key in secrets/val-chris-consensus.key and remove it from the validator

```
docker run --rm -t \
    -v $(pwd)/prod/val-chris:/root/.c-chain \
    c-chaind \
    tendermint show-validator \
    | tr -d '\n' | tr -d '\r' \
    > prod/desk-chris/config/pub_validator_key-val-chris.json
```

### Set up Validators

Ensure Validator Chris listens on port 26659

```
# Notify KMS of address chris will run on
docker run --rm -i \
    -v $(pwd)/prod/kms-chris:/root/tmkms \
    --entrypoint sed \
    tmkms_i:v0.12.2 \
    -Ei 's/^addr = "tcp:.*$/addr = "tcp:\/\/val-chris:26659"/g' /root/tmkms/tmkms.toml

# Make sure validator chris runs on the specified port
docker run --rm -i \
  -v $(pwd)/prod/val-chris:/root/.c-chain \
  --entrypoint sed \
  c-chaind \
  -Ei 's/priv_validator_laddr = ""/priv_validator_laddr = "tcp:\/\/0.0.0.0:26659"/g' \
  /root/.c-chain/config/config.toml
```

## Coming Together

We are now ready to run our chain with validators. To simply running of this application, we are going to use docker-compose to run our set up as docker images locally. We do this by running

` docker-compose -f c-chain/docker-compose.yaml up -d`

## Hosting Options

To make this application available to the world, we would need a robust mechanism of hosting it beyond docker-compose. Some of this option could be using:
- Kubernetes to run our docker containers
- Bare metal VM if we need speed
