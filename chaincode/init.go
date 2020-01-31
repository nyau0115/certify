package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// AidChaincode - define a type for chaincode.
// AidChaincode type must implements shim.Chaincode interface
type CVSChaincode struct {
}

// Constants
const (
	// Error codes
	CODE_ALL_A_OK              string = "P2001" // Success
	CODE_NOT_FOUND            string = "P4001" // resource not found
	CODE_UNKNOWN_INVOKE       string = "P4002" // Unknown invoke
	CODE_UNPROCESSABLE_ENTITY string = "P4003" // Invalid input
	CODE_GEN_EXCEPTION        string = "P5001" // Unknown exception
	CODE_AlRD_EXIST           string = "P5002" // Not unique
	CODE_NOT_ALLWD            string = "P4004" // Operation not allowed

	// Couch DB Doc types for asset
	GPRJCT string = "GPRJCT"
	GITEM  string = "GITEM"
	DONIN  string = "DONIN"
	DONOUT string = "DONOUT"

	// Range index name - to perform range queries
	INDXNM string = "bitmask~txnID~amount" //bitmask is "0" for donation (spending) & "1" donation(incoming)

	FIXEDPT int32 = 4 // All currency values rounded off to 4 decimals i.e. 0.0000
)

// Init - Implements shim.Chaincode interface Init() method
func (t *CVSChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	r := response{(CODE_ALL_A_OK), "CVScc started", nil}
	return shim.Success((r.formatResponse()))
}


