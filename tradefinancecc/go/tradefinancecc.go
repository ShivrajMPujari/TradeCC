package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	peer "github.com/hyperledger/fabric/protos/peer"
)

type Account struct {
	AccountNumber     string
	AccountHolderName string
	Balance           int
}

type Contract struct {
}

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()
	var A string
	var B string
	var Avalue string
	var Bvalue string

	A = args[0]
	Avalue = args[1]
	B = args[2]
	Bvalue = args[3]

	// AvalueInt,_ := strconv.Atoi(Avalue)
	// AvalueInt,_:= strconv.Atoi(Bvalue)
	err1 := stub.PutState(A, []byte(Avalue))

	err2 := stub.PutState(B, []byte(Bvalue))

	if err1 != nil || err2 != nil {
		return shim.Error("error while updating the ledger")
	}

	return shim.Success(nil)
	fmt.Println("init called")
	return shim.Success(nil)
}

// Invoke
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("init called")
	function, args := stub.GetFunctionAndParameters()

	if len(args) > 2 {
		shim.Error("input arguments are less ....")
	}
	if function == "invoke" {
		return t.invoke(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	} else if function == "createAccount" {
		return t.createAccount(stub, args)
	} else if function == "createContract" {
		return t.createContract(stub, args)
	} else if function == "getBalance" {
		return t.getBalance(stub, args)
	} else if function == "getAccount" {
		return t.getAccount(stub, args)
	} else if function == "getContract" {
		return t.getContract(stub, args)
	} else if function == "deleteAccount" {
		return t.deleteAccount(stub, args)
	}

	return shim.Error("Invalid invoke function must pass invoke/query as arguments ")
}

func (t *SimpleAsset) invoke(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var A string
	var B string
	var Avalue string
	var Bvalue string

	A = args[0]
	Avalue = args[1]
	B = args[2]
	Bvalue = args[3]

	// AvalueInt,_ := strconv.Atoi(Avalue)
	// AvalueInt,_:= strconv.Atoi(Bvalue)
	err1 := stub.PutState(A, []byte(Avalue))

	err2 := stub.PutState(B, []byte(Bvalue))

	stub.PutState("C", []byte("100"))
	if err1 != nil || err2 != nil {
		return shim.Error("error while updating the ledger")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleAsset) query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func (t *SimpleAsset) createAccount(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	accountBal, errBal := strconv.Atoi(args[2])

	if errBal != nil {
		return shim.Error("invalid balance.....")
	}
	var account = Account{AccountNumber: args[0], AccountHolderName: args[1], Balance: accountBal}
	accountByte, err := json.Marshal(account)

	if err != nil {
		return shim.Error("account is not been created")
	}
	stub.PutState(args[0], accountByte)

	return shim.Success(nil)
}

func (t *SimpleAsset) createContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success(nil)
}

func (t *SimpleAsset) getBalance(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	accountNumber := args[0]

	accountByte, accountFetchErr := stub.GetState(accountNumber)
	if accountFetchErr != nil {
		return shim.Error("error while gettting account info in getstate method ")
	}
	var tempAccountSruct Account
	errAccountStruct := json.Unmarshal(accountByte, &tempAccountSruct)
	if errAccountStruct != nil {
		shim.Error("error while converting to Struct Account")
	}

	tempBalance := strconv.Itoa(tempAccountSruct.Balance)
	return shim.Success([]byte(tempBalance))
}

func (t *SimpleAsset) getAccount(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	accountNumber := args[0]

	accountByte, accountFetchErr := stub.GetState(accountNumber)
	if accountFetchErr != nil {
		return shim.Error("error while gettting account info in getstate method ")
	}
	var tempAccountSruct Account
	errAccountStruct := json.Unmarshal(accountByte, &tempAccountSruct)
	if errAccountStruct != nil {
		shim.Error("error while converting to Struct Account")
	}

	return shim.Success(accountByte)
}

func (t *SimpleAsset) deleteAccount(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	account := args[0]
	errorInDelAcc := stub.DelState(account)

	if errorInDelAcc != nil {
		return shim.Error("Error while deleting the account")
	}

	return shim.Success(nil)

}

func (t *SimpleAsset) getContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success(nil)
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
