package main

import (
	"bytes"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// respnse struct to have consistent response structure for all chaincode invokes
type response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Payload []byte `json:"payload"`
}

func (r *response) formatResponse() []byte {
	var buffer bytes.Buffer
	buffer.WriteString("{\"code\":")
	buffer.WriteString("\"")
	buffer.WriteString(r.Code)
	buffer.WriteString("\",")
	buffer.WriteString("\"message\":")
	buffer.WriteString("\"")
	buffer.WriteString(r.Message)
	if r.Payload == nil {
		buffer.WriteString("\"}")
	} else {
		buffer.WriteString("\",")
		buffer.WriteString("\"payload\":")
		buffer.Write(r.Payload)
		buffer.WriteString("}")
	}
	return buffer.Bytes()
}

// chainError - custom error
type chainError struct {
	fcn  string // function/method
	key  string // associated key, if any
	code string // Error code
	err  error  // Error.
}

func (e *chainError) Error() string {
	return e.fcn + " " + e.key + " " + e.code + ": " + e.err.Error()
}

// checkAsset - check if an asset ( with a key) is already available on the ledger
func checkAsset(stub shim.ChaincodeStubInterface, assetID string) (bool, *chainError) {

	assetBytes, err := stub.GetState(assetID)

	if err != nil {
		e := &chainError{"checkAsset", assetID, CODE_GEN_EXCEPTION, err}
		return false, e
	} else if assetBytes != nil {
		//e := &chainError{"checkAsset", assetID, CODEAlRDEXIST, errors.New("Asset with key already exists")}
		return true, nil
	}
	return false, nil
}

// queryAsset - return query state from the ledger
func queryAsset(stub shim.ChaincodeStubInterface, assetID string) ([]byte, *chainError) {

	assetBytes, err := stub.GetState(assetID)

	if err != nil {
		e := &chainError{"queryAsset", assetID, CODE_GEN_EXCEPTION, err}
		return nil, e
	} else if assetBytes == nil {
		e := &chainError{"queryAsset", assetID, CODE_NOT_FOUND, errors.New("Asset ID not found")}
		return nil, e
	}
	return assetBytes, nil
}
