# Hyperledger Go SDK Examples

The aim of this project is to demonstrate Hyperledger fabric features using the HF Client SDK for Golang. This repository contains implementation of [Fabcar](https://hyperledger-fabric.readthedocs.io/en/release-1.4/understand_fabcar_network.html) example in Go SDK using the [first-network](https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html).

## TODO

- Better documentation
- Private data collection example
- Use Raft ordering
- Web server with APIs for operations
- Simple web UI
- Chaincode support in multiple languages
- Fabric 2.0 Compatibility
  
## Prerequisites

Before start, please make sure that you have all the required [prerequisites](https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html) installed.

## How to Run...

- Download/clone this repository inside your $GOPATH/src folder.
- cd $GOPATH/src/gosdk-example && make

## What the make command does?

1. Clean cache of previous run
2. Brings 'first-network' down
3. Build main.go
4. Start the first network
5. Run main.go

## main.go

1. Create admins for Org1 and Org2
2. Create channel mychannel
3. Join peers to mychannel
4. Install and instantiate fabcar chaincode on peers
5. Run sample execute and query commands