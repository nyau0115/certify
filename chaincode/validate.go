package main
//
//import (
//    "fmt"
//    "strconv"
//    "encoding/json"
//	"github.com/hyperledger/fabric/core/chaincode/shim"
//	pb "github.com/hyperledger/fabric/protos/peer"
//)
//
//		return shim.Error(cErr.Error())
//	}
//	if len(args[1]) == 0 {
//		cErr := &chainError{"validateProjectW", "", CODEUNPROCESSABLEENTITY, errors.New("Project name can not be empty")}
//		//logger.Error(cErr.Error())
//		return shim.Error(cErr.Error())
//	}
//
//	epochTime, _ := stub.GetTxTimestamp()
//	startDt := time.Unix(epochTime.GetSeconds(), 0)
//	// Bypass whilst running unit test
//	var callerID string
//	if os.Getenv("MODE") != "TEST" {
//		var err *chainError
//		callerID, err = getCallerID(stub)
//		if err != nil {
//			return shim.Error(err.Error())
//		}
//	} else {
//		callerID = "Test Caller"
//	}
//
//	prjBase := projectBase{ProjectName: args[1], StartDt: startDt, RunBy: callerID}
//	p := &project{ObjectType: GPRJCT, ProjectID: args[0], Data: prjBase}
//	return saveAsset(stub, p)
//}
//
//func validateItemW(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//
//	if len(args) != 3 {
//		cErr := &chainError{"validateItemW", "", CODEUNPROCESSABLEENTITY, errors.New("Incorrect no of input args, excepting 3")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[0]) == 0 {
//		cErr := &chainError{"validateItemW", "", CODEUNPROCESSABLEENTITY, errors.New("Item ID can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[1]) == 0 {
//		cErr := &chainError{"validateItemW", "", CODEUNPROCESSABLEENTITY, errors.New("Item type can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[2]) == 0 {
//		cErr := &chainError{"validateItemW", "", CODEUNPROCESSABLEENTITY, errors.New("Item narrative can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//
//	itmBase := itemBase{ItemType: args[1], Narrative: args[2]}
//	it := &item{ObjectType: GITEM, ItemID: args[0], Data: itmBase}
//
//	return saveAsset(stub, it)
//}
//
//func validateDonationW(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//
//	if len(args) != 5 {
//		cErr := &chainError{"validateDonationW", "", CODEUNPROCESSABLEENTITY, errors.New("Incorrect no of input args, excepting 5")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[0]) == 0 {
//		cErr := &chainError{"validateDonationW", "", CODEUNPROCESSABLEENTITY, errors.New("Txn ID can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[1]) == 0 {
//		cErr := &chainError{"validateDonationW", "", CODEUNPROCESSABLEENTITY, errors.New("Donor name can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[2]) == 0 {
//		cErr := &chainError{"validateDonationW", "", CODEUNPROCESSABLEENTITY, errors.New("Affiliated project ID can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[3]) == 0 {
//		cErr := &chainError{"validateDonationW", "", CODEUNPROCESSABLEENTITY, errors.New("Item ID can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[4]) == 0 {
//		cErr := &chainError{"validateDonationW", "", CODEUNPROCESSABLEENTITY, errors.New("Donation amount can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	zeroDecimal, _ := decimal.NewFromString("0")
//	amount, _ := decimal.NewFromString(args[4])
//	if amount.LessThanOrEqual(zeroDecimal) {
//		cErr := &chainError{"validateDonationW", "", CODEUNPROCESSABLEENTITY, errors.New("Donation amount can not less than or equal to zero")}
//		return shim.Error(cErr.Error())
//	}
//
//	epochTime, _ := stub.GetTxTimestamp()
//	timeStamp := time.Unix(epochTime.GetSeconds(), 0)
//
//	donBase := donationBase{ProjectID: args[2], ItemID: args[3], Amount: amount.RoundBank(FIXEDPT), TimeStamp: timeStamp}
//	d := &donation{ObjectType: DONIN, TxnID: args[0], Donor: args[1], Data: donBase}
//
//	return saveAsset(stub, d)
//}
//
//func validateSpendW(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//
//	if len(args) != 5 {
//		cErr := &chainError{"validateSpendW", "", CODEUNPROCESSABLEENTITY, errors.New("Incorrect no of input args, excepting 5")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[0]) == 0 {
//		cErr := &chainError{"validateSpendW", "", CODEUNPROCESSABLEENTITY, errors.New("Txn ID can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[1]) == 0 {
//		cErr := &chainError{"validateSpendW", "", CODEUNPROCESSABLEENTITY, errors.New("Beneficiary can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[2]) == 0 {
//		cErr := &chainError{"validateSpendW", "", CODEUNPROCESSABLEENTITY, errors.New("Affiliated project ID can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[3]) == 0 {
//		cErr := &chainError{"validateSpendW", "", CODEUNPROCESSABLEENTITY, errors.New("Item ID can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[4]) == 0 {
//		cErr := &chainError{"validateSpendW", "", CODEUNPROCESSABLEENTITY, errors.New("Donation amount can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//	zeroDecimal, _ := decimal.NewFromString("0")
//	amount, _ := decimal.NewFromString(args[4])
//	if amount.LessThanOrEqual(zeroDecimal) {
//		cErr := &chainError{"validateSpendW", "", CODEUNPROCESSABLEENTITY, errors.New("Donation amount can not less than or equal to zero")}
//		return shim.Error(cErr.Error())
//	}
//
//	epochTime, _ := stub.GetTxTimestamp()
//	timeStamp := time.Unix(epochTime.GetSeconds(), 0)
//
//	spendData := donationBase{ProjectID: args[2], ItemID: args[3], Amount: amount.RoundBank(FIXEDPT), TimeStamp: timeStamp}
//	s := &spend{ObjectType: DONOUT, TxnID: args[0], Benficiary: args[1], Data: spendData}
//
//	return saveAsset(stub, s)
//}
//
//func validateProjectR(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//
//	if len(args) != 1 {
//		cErr := &chainError{"validateProjectR", "", CODEUNPROCESSABLEENTITY, errors.New("Incorrect no of input args, excepting project ID only")}
//		return shim.Error(cErr.Error())
//	}
//
//	if len(args[0]) == 0 {
//		cErr := &chainError{"validateProjectR", "", CODEUNPROCESSABLEENTITY, errors.New("Project ID can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//
//	p := &project{ProjectID: args[0]}
//	return readAsset(stub, p)
//}
//
//func validateItemR(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//
//	if len(args) != 1 {
//		cErr := &chainError{"validateItemR", "", CODEUNPROCESSABLEENTITY, errors.New("Incorrect no of input args, excepting Item ID only")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[0]) == 0 {
//		cErr := &chainError{"validateItemR", "", CODEUNPROCESSABLEENTITY, errors.New("Item ID can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//
//	it := &item{ItemID: args[0]}
//	return readAsset(stub, it)
//}
//
//func validateDonationR(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//
//	if len(args) != 1 {
//		cErr := &chainError{"validateDonationR", "", CODEUNPROCESSABLEENTITY, errors.New("Incorrect no of input args, excepting Transaction ID only")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[0]) == 0 {
//		cErr := &chainError{"validateDonationR", "", CODEUNPROCESSABLEENTITY, errors.New("Txn ID can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//
//	d := &donation{TxnID: args[0]}
//	return readAsset(stub, d)
//}
//
//func validateSpendR(stub shim.ChaincodeStubInterface, args []string) pb.Response {
//
//	if len(args) != 1 {
//		cErr := &chainError{"validateSpendR", "", CODEUNPROCESSABLEENTITY, errors.New("Incorrect no of input args, excepting Transaction ID only")}
//		return shim.Error(cErr.Error())
//	}
//	if len(args[0]) == 0 {
//		cErr := &chainError{"validateSpendR", "", CODEUNPROCESSABLEENTITY, errors.New("Txn ID can not be empty")}
//		return shim.Error(cErr.Error())
//	}
//
//	s := &spend{TxnID: args[0]}
//	return readAsset(stub, s)
//}

//func (t *CVSChaincode) initData(APIstub shim.ChaincodeStubInterface) pb.Response {
//    organisation := Organisation{Id: "1", Name: "HKU", Region: "hk"}
//	users := []User{
//		User{Id: "30354", Name: "Nelson", Gender: "M", NationalId: "R123", PhoneNumber:"66666666", Address: "mk", Organisation: organisation, UserType: "Student"},
//		User{Id: "30355", Name: "Ray", Gender: "M", NationalId: "R124", PhoneNumber:"99999999", Address: "cwb", Organisation: organisation, UserType: "Student"},
//	}
//
//	i := 0
//	for i < len(users) {
//		fmt.Println("i is ", i)
//		usersAsBytes, _ := json.Marshal(users[i])
//		APIstub.PutState(strconv.Itoa(i+1), usersAsBytes)
//		fmt.Println("Added", users[i])
//		i = i + 1
//	}
//
//	return shim.Success(nil)
//}
//
//
//func (t *CVSChaincode) retrieveUser(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
//
//	if len(args) != 1 {
//		return shim.Error("Incorrect number of arguments. Expecting 1")
//	}
//
//	userAsBytes, _ := APIstub.GetState(args[0])
//	if userAsBytes == nil {
//		return shim.Error("Could not locate tuna")
//	}
//	return shim.Success(userAsBytes)
//}
//
//func (t *CVSChaincode) listUser(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
//
//	if len(args) != 1 {
//		return shim.Error("Incorrect number of arguments. Expecting 1")
//	}
//
//	userAsBytes, _ := APIstub.GetState(args[0])
//	if userAsBytes == nil {
//		return shim.Error("Could not locate tuna")
//	}
//	return shim.Success(userAsBytes)
//}
//
//func (t *CVSChaincode) createUser(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
//
//	if len(args) != 1 {
//		return shim.Error("Incorrect number of arguments. Expecting 1")
//	}
//
//	userAsBytes, _ := APIstub.GetState(args[0])
//	if userAsBytes == nil {
//		return shim.Error("Could not locate tuna")
//	}
//	return shim.Success(userAsBytes)
//}
//
//func (t *CVSChaincode) updateUser(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
//
//	if len(args) != 1 {
//		return shim.Error("Incorrect number of arguments. Expecting 1")
//	}
//
//	userAsBytes, _ := APIstub.GetState(args[0])
//	if userAsBytes == nil {
//		return shim.Error("Could not locate tuna")
//	}
//	return shim.Success(userAsBytes)
//}