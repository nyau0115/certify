package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("CVSChaincode")

// Chaincode entry point
func main() {
	logger.SetLevel(shim.LogInfo)
	err := shim.Start(new(CVSChaincode))
	if err != nil {
		logger.Error("Error starting CVSChaincode - ", err)
	}

}