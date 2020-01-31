package main

import (
	"errors"
	"fmt"
	"strconv"
	 "encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// These are function names from Invoke first parameter
const (
	Init    string = "Init"
	InitData    string = "InitData"
	RetrieveUser    string = "RetrieveUser"
	ListUser    string = "ListUser"
	UpdateUser    string = "UpdateUser"
	CreateUser    string = "CreateUser"
)

// Invoke - Implements shim.Chaincode interface Invoke() method
func (t *CVSChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	// Retrieve function and input arguments. This is not a recommended approach.
	// Instead, GetArgs() is a more suitable method and works perfectly with protocol buffers
	function, args := stub.GetFunctionAndParameters()
	logger.Info(fmt.Sprintf("Starting Phantom chaincode Invoke for %s and no of argument passed are %d", function, len(args)))

	if function == Init {
		return t.Init(stub)
	} else if function == InitData {
     	return t.initData(stub)
    } else if function == RetrieveUser {
		return t.retrieveUser(stub, args)
	} else if function == ListUser {
		return t.listUser(stub, args)
	} else if function == CreateUser {
		return t.createUser(stub, args)
	} else if function == UpdateUser {
		return t.updateUser(stub, args)
	}

	e := chainError{"Invoke", "", CODE_UNKNOWN_INVOKE, errors.New("Unknown function invoke")}
	return shim.Error(e.Error())
}

func (t *CVSChaincode) initData(APIstub shim.ChaincodeStubInterface) pb.Response {
    organisation := Organisation{Id: "1", Name: "HKU", Region: "hk"}
	users := []User{
		User{Id: "30354", Name: "Nelson", Gender: "M", NationalId: "R123", PhoneNumber:"66666666", Address: "mk", Organisation: organisation, UserType: "Student"},
		User{Id: "30355", Name: "Ray", Gender: "M", NationalId: "R124", PhoneNumber:"99999999", Address: "cwb", Organisation: organisation, UserType: "Student"},
	}

	i := 0
	for i < len(users) {
		fmt.Println("i is ", i)
		usersAsBytes, _ := json.Marshal(users[i])
		APIstub.PutState(strconv.Itoa(i+1), usersAsBytes)
		fmt.Println("Added", users[i])
		i = i + 1
	}

	return shim.Success(nil)
}


func (t *CVSChaincode) retrieveUser(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userAsBytes, _ := APIstub.GetState(args[0])
	if userAsBytes == nil {
		return shim.Error("Could not locate tuna")
	}
	return shim.Success(userAsBytes)
}

func (t *CVSChaincode) listUser(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userAsBytes, _ := APIstub.GetState(args[0])
	if userAsBytes == nil {
		return shim.Error("Could not locate tuna")
	}
	return shim.Success(userAsBytes)
}

func (t *CVSChaincode) createUser(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userAsBytes, _ := APIstub.GetState(args[0])
	if userAsBytes == nil {
		return shim.Error("Could not locate tuna")
	}
	return shim.Success(userAsBytes)
}

func (t *CVSChaincode) updateUser(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userAsBytes, _ := APIstub.GetState(args[0])
	if userAsBytes == nil {
		return shim.Error("Could not locate tuna")
	}
	return shim.Success(userAsBytes)
}


export CORE_PEER_ID=peer0.org1.example.com
export CORE_LOGGING_PEER=debug
export CORE_CHAINCODE_LOGGING_LEVEL=DEBUG
export CORE_PEER_LOCALMSPID=Org1MSP
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/msp
export CORE_PEER_ADDRESS=peer0.org1.example.com:7051


CONFIG_ROOT=/opt/gopath/src/github.com/hyperledger/fabric/peer
ORG1_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
ORG1_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
ORG2_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
ORG2_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
ORDERER_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
