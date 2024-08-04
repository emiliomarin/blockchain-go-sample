## Sample Blockchain + Go app

### Install
- [Ganache](https://archive.trufflesuite.com/ganache/) 
- [Truffle](https://archive.trufflesuite.com/docs/truffle/how-to/install/)

### Usage
- Start Ganache to start a new ethereum network in localhost:7545
- Run smart contract migrations to deploy the contract: `npx truffle migrate`
- Run app `go run main.go`

This will run the application which creates some sensors, reads data from the chain and updates the sensor value.