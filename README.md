# Hyperledger Go SDK Example

The aim of this project is to demonstrate Hyperledger fabric features using the HF Client SDK for Golang. This repository contains implementation of [Fabcar](https://hyperledger-fabric.readthedocs.io/en/release-1.4/understand_fabcar_network.html) example in Go SDK using the [first-network](https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html). SDK v1.0.0-beta1 version is used to run the application.

## TODO

- [x] Better documentation (done)
- [ ] Private data collection example
- [ ] Query ledger 
- [ ] Use Raft ordering
- [x] Web server with APIs for operations (done)
- [ ] Simple web UI
- [ ] ~~Chaincode support in multiple languages~~ (SDK v1.0.0-beta1 release supports Golang chaincode only)
- [ ] ~~Fabric 2.0 Compatibility~~ (SDK v1.0.0-beta1 supports Fabric 1.4 only)
  
Pulls are welcome!!!
  
## Prerequisites

Before start, please make sure that you have all the required [prerequisites](https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html) installed.

## How to Run...

- Clone this repository inside your $GOPATH/src folder.
- cd $GOPATH/src/gosdk-example && make
- Now, a web server will be up and listening on localhost:3000 port.
- Open a new terminal and run testAPIs.sh

## What the make command does?

1. Clean cache of previous run
2. Brings 'first-network' down
3. Build main.go
4. Start the first network
5. Run main.go (Initiate setups of two organizations and starts the web server.)

## API requests

Web server supports following requests.testAPIs.sh file contains sample curl requests.
1. User Enrollment
2. Channel Creation
3. Join Channel
4. Install Chaincode
5. Instantiate Chaincode
6. Execute chaincode
7. Query chaincode