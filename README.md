# Hyperledger Go SDK Example

The aim of this project is to demonstrate Hyperledger fabric features using the HF Client SDK for Golang. This repository contains implementation of [Fabcar Network](https://hyperledger-fabric.readthedocs.io/en/release-1.4/understand_fabcar_network.html) in Go SDK using the [first-network](https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html), although it is possible to use other chaincode/network setup. Implementation of Fabcar network in NodeJS can be found [here](https://hyperledger-fabric.readthedocs.io/en/release-1.4/write_first_app.html). Go SDK v1.0.0-beta1 version is used to run the application.

## TODO

- [x] Better documentation (done)
- [ ] Private data collection example
- [ ] Add query ledger apis
- [x] Use Raft ordering (done)
- [x] Web server with APIs for operations (done)
- [ ] Simple web UI
- [ ] ~~Chaincode support in multiple languages~~ (SDK v1.0.0-beta1 release supports Golang chaincode only)
- [ ] ~~Fabric 2.0 Compatibility~~ (SDK v1.0.0-beta1 supports Fabric 1.4 only)
  
Pulls are welcome!!!
  
## Prerequisites

Before start, please make sure that you have all the required prerequisites ([link1](https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html),[link2](https://hyperledger-fabric.readthedocs.io/en/release-1.4/install.html)) installed. You also need the following additional packages installed.

1. make (sudo apt install make)

## How to Run...

- Clone this repository inside your $GOPATH/src folder.
- cd $GOPATH/src/gosdk-example && make
- Now, a web server will be up and listening on localhost:3000 port.
- Open a new terminal and run testAPIs.sh
- This application uses CouchDB database and Raft ordering. This can be changed in the startFabric.sh file.

## What the make command does?

1. Clean cache of previous run
2. Brings 'first-network' down
3. Build main.go
4. Start the first network
5. Run main.go (Initiate setups of two organizations and starts the web server.)

## API requests

Web server supports following requests. testAPIs.sh file contains sample curl requests.
1. User Enrollment
2. Channel Creation
3. Join Channel
4. Install Chaincode
5. Instantiate Chaincode
6. Execute chaincode
7. Query chaincode