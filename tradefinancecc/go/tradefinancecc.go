package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	peer "github.com/hyperledger/fabric/protos/peer"
)

//Account structure
type Account struct {
	AccountNumber     string `json:"accountNumber"`
	AccountHolderName string `json:"accountHolderName"`
	Balance           int    `json:"balance"`
	Bank              string `json:"bank"`
}

//Contract structure
type Contract struct {
	ContractID         string `json:"contractId"`
	ContentDescription string `json:"contentDescription"`
	Value              int    `json:"value"`
	ImporterID         string `json:"importerId"`
	ExporterID         string `json:"exporterId"`
	CustomID           string `json:"customId"`
	ImporterBankID     string `json:"importerBankId"`
	InsuranceID        string `json:"insuranceId"`
	PortOfLoading      string `json:"portOfLoading"`
	PortOfEntery       string `json:"portOfEntry"`
}

// SimpleAsset implements a chaincode to manage an asset
type SimpleAsset struct {
}

//Documents structure
type Documents struct {
	BillOfLading   string
	LetterOfCredit string
}

//Logger creation ......
var Logger = shim.NewLogger("tradefinancecc")

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

	fmt.Println("init called")
	return shim.Success(nil)
}

// Invoke ....
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("init called")
	function, args := stub.GetFunctionAndParameters()
	Logger.Info("success init")
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
	} else if function == "importerAssurity" {
		return t.importerAssurity(stub, args)
	} else if function == "exporterAssurity" {
		return t.exporterAssurity(stub, args)
	} else if function == "customAssurity" {
		return t.customAssurity(stub, args)
	} else if function == "importerBankAssurity" {
		return t.importerBankAssurity(stub, args)
	} else if function == "insuranceAssurity" {
		return t.insuranceAssurity(stub, args)
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
	var account = Account{AccountNumber: args[0], AccountHolderName: args[1], Balance: accountBal, Bank: args[2]}
	accountByte, err := json.Marshal(account)

	if err != nil {
		return shim.Error("account is not been created")
	}
	stub.PutState(args[0], accountByte)

	return shim.Success(nil)
}

func (t *SimpleAsset) createContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 8 {
		shim.Error("input arguments are less")
	}

	contractValue, errContractValue := strconv.Atoi(args[2])
	if errContractValue != nil {
		return shim.Error("error while converting string to int ")
	}
	var contract = Contract{ContractID: args[0], ContentDescription: args[1], Value: contractValue, ImporterID: args[3], ExporterID: args[4], CustomID: args[5], ImporterBankID: args[6], InsuranceID: args[7], PortOfLoading: args[8], PortOfEntery: args[9]}
	contractByte, errContractByte := json.Marshal(contract)
	if errContractByte != nil {
		return shim.Error("error while converting to json byte stream")
	}
	stub.PutState(contract.ContractID, contractByte)

	return shim.Success(contractByte)
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
	tempAccountSruct := Account{}
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

	contractInByte, errInContractByte := stub.GetState(args[0])
	if errInContractByte != nil {
		return shim.Error("error while getting state")
	}
	if contractInByte == nil {
		return shim.Error("error while getting state")
	}

	return shim.Success(contractInByte)
}

func (t *SimpleAsset) importerAssurity(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	importerID := args[0]
	contractID := args[1]

	importerByte, errImporter := stub.GetState(importerID)
	if errImporter != nil {
		return shim.Error("error while getting importer")
	}
	Logger.Info(importerByte)
	contractByte, contractErr := stub.GetState(contractID)
	if contractErr != nil {
		return shim.Error("error while getting contract...")
	}

	tempImporterSruct := Account{}
	errAccountStruct := json.Unmarshal(importerByte, &tempImporterSruct)

	tempContractSruct := Contract{}
	errContractStruct := json.Unmarshal(contractByte, &tempContractSruct)

	if errContractStruct != nil {

		return shim.Error("cant unmarshall the byte contractByte ")
	}

	if errAccountStruct != nil {
		return shim.Error("error while converting to Struct Account")
	}

	if tempContractSruct.PortOfLoading != "denmark" || tempContractSruct.PortOfEntery != "berlin" {

		return shim.Error("port of loading or port of entry doesn't matches")
	}

	if tempImporterSruct.AccountNumber != tempContractSruct.ImporterID {

		return shim.Error("importer is not valid,according to contract")
	} else if tempImporterSruct.Balance < tempContractSruct.Value {
		return shim.Error("insufficient balance in importers account  ")
	}
	respond := "importer is validated"
	return shim.Success([]byte(respond))
}

func (t *SimpleAsset) exporterAssurity(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	exporterID := args[0]
	contractID := args[1]

	exporterByte, errExporter := stub.GetState(exporterID)
	if errExporter != nil {
		return shim.Error("error while getting exporter")
	}

	tempExporterSruct := Account{}
	errAccountStruct := json.Unmarshal(exporterByte, &tempExporterSruct)

	if errAccountStruct != nil {
		return shim.Error("error while converting to Struct Account")
	}

	contractByte, contractErr := stub.GetState(contractID)
	if contractErr != nil {
		return shim.Error("error while getting contract...")
	}
	tempContractSruct := Contract{}
	errContractStruct := json.Unmarshal(contractByte, &tempContractSruct)

	if errContractStruct != nil {

		return shim.Error("cant unmarshall the byte contractByte ")
	}

	if tempExporterSruct.AccountNumber != tempContractSruct.ExporterID {

		return shim.Error("Exporter doesn't exist on contract")
	}

	if tempContractSruct.PortOfLoading != "denmark" || tempContractSruct.PortOfEntery != "berlin" {

		return shim.Error("port of loading or port of entry doesn't matches")
	}

	importerID := tempContractSruct.ImporterID
	importerByte, importerErr := stub.GetState(importerID)
	if importerErr != nil {
		return shim.Error("error while getting state of importer")
	}

	tempImporterSruct := Account{}
	errImpAccountStruct := json.Unmarshal(importerByte, &tempImporterSruct)
	if errImpAccountStruct != nil {
		return shim.Error("error while converting to Struct Account importer")
	}

	if tempImporterSruct.Balance < tempContractSruct.Value {
		return shim.Error("insufficient balance in importers account  ")
	}

	respond := "exporter is validated"
	return shim.Success([]byte(respond))
}

func (t *SimpleAsset) customAssurity(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	customID := args[0]
	contractID := args[1]

	customByte, errCustom := stub.GetState(customID)
	if errCustom != nil {
		return shim.Error("error while getting custom")
	}

	tempCustomSruct := Account{}
	errAccountStruct := json.Unmarshal(customByte, &tempCustomSruct)

	if errAccountStruct != nil {
		return shim.Error("cant unmarshall the byte customByte ")
	}

	contractByte, contractErr := stub.GetState(contractID)
	if contractErr != nil {
		return shim.Error("error while getting contract...")
	}
	tempContractSruct := Contract{}
	errContractStruct := json.Unmarshal(contractByte, &tempContractSruct)

	if errContractStruct != nil {

		return shim.Error("cant unmarshall the byte contractByte ")
	}

	if tempCustomSruct.AccountNumber != tempContractSruct.CustomID {
		return shim.Error("invalid custom account")
	}
	if tempContractSruct.PortOfLoading != "denmark" || tempContractSruct.PortOfEntery != "berlin" {

		return shim.Error("port of loading or port of entry doesn't matches")
	}

	importerID := tempContractSruct.ImporterID
	importerByte, importerErr := stub.GetState(importerID)
	if importerErr != nil {
		return shim.Error("error while getting state of importer")
	}

	tempImporterSruct := Account{}
	errImpAccountStruct := json.Unmarshal(importerByte, &tempImporterSruct)
	if errImpAccountStruct != nil {
		return shim.Error("error while converting to Struct Account importer")
	}

	if tempImporterSruct.Balance < tempContractSruct.Value {
		return shim.Error("insufficient balance in importers account  ")
	}
	respond := "custom is validated"
	return shim.Success([]byte(respond))
}

func (t *SimpleAsset) importerBankAssurity(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	importerBankID := args[0]
	contractID := args[1]

	importerBankByte, errimporterBank := stub.GetState(importerBankID)

	if errimporterBank != nil {
		return shim.Error("error while getting errimporterBank")
	}

	importerBankStruct := Account{}

	errImporterBankStruct := json.Unmarshal(importerBankByte, &importerBankStruct)

	if errImporterBankStruct != nil {

		return shim.Error("error while unmarshall importerBank structure")
	}

	contractByte, contractErr := stub.GetState(contractID)
	if contractErr != nil {
		return shim.Error("error while getting contract...")
	}
	tempContractSruct := Contract{}
	errContractStruct := json.Unmarshal(contractByte, &tempContractSruct)

	if errContractStruct != nil {

		return shim.Error("cant unmarshall the byte contractByte ")
	}

	if tempContractSruct.ImporterBankID != importerBankStruct.AccountNumber {

		return shim.Error("invalid importer Id ")
	}

	if tempContractSruct.PortOfLoading != "denmark" || tempContractSruct.PortOfEntery != "berlin" {

		return shim.Error("port of loading or port of entry doesn't matches")
	}

	importerID := tempContractSruct.ImporterID
	importerByte, importerErr := stub.GetState(importerID)
	if importerErr != nil {
		return shim.Error("error while getting state of importer")
	}

	tempImporterSruct := Account{}
	errImpAccountStruct := json.Unmarshal(importerByte, &tempImporterSruct)
	if errImpAccountStruct != nil {
		return shim.Error("error while converting to Struct Account importer")
	}

	if tempImporterSruct.Balance < tempContractSruct.Value {
		return shim.Error("insufficient balance in importers account  ")
	}
	respond := "importerBank is validated"
	return shim.Success([]byte(respond))
}

func (t *SimpleAsset) insuranceAssurity(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	accountID := args[0]
	contractID := args[1]

	accountStruct, tempContractSruct, anyError := t.createStruct(stub, accountID, contractID)

	if anyError != "" {
		return shim.Error(anyError)
	}

	// insuranceByte, errInsurance := stub.GetState(insuranceID)

	// if errInsurance != nil {
	// 	return shim.Error("error while getting errimporterBank")
	// }

	// insuranceStruct := Account{}

	// errInsuranceStruct := json.Unmarshal(insuranceByte, &insuranceStruct)

	// if errInsuranceStruct != nil {

	// 	return shim.Error("error while unmarshall insurance structure")
	// }

	// contractByte, contractErr := stub.GetState(contractID)
	// if contractErr != nil {
	// 	return shim.Error("error while getting contract...")
	// }
	// tempContractSruct := Contract{}
	// errContractStruct := json.Unmarshal(contractByte, &tempContractSruct)

	// if errContractStruct != nil {

	// 	return shim.Error("can't unmarshall the byte contractByte ")
	// }

	if accountStruct.AccountNumber != tempContractSruct.InsuranceID {

		return shim.Error("insurance account is not valid...")
	}

	if tempContractSruct.PortOfLoading != "denmark" || tempContractSruct.PortOfEntery != "berlin" {

		return shim.Error("port of loading or port of entry doesn't matches")
	}

	importerID := tempContractSruct.ImporterID
	importerByte, importerErr := stub.GetState(importerID)
	if importerErr != nil {
		return shim.Error("error while getting state of importer")
	}

	tempImporterSruct := Account{}
	errImpAccountStruct := json.Unmarshal(importerByte, &tempImporterSruct)
	if errImpAccountStruct != nil {
		return shim.Error("error while converting to Struct Account importer")
	}

	if tempImporterSruct.Balance < tempContractSruct.Value {
		return shim.Error("insufficient balance in importers account  ")
	}

	//respond := "insurance is validated"
	return shim.Success(importerByte)
}

func (t *SimpleAsset) createStruct(stub shim.ChaincodeStubInterface, accountID string, contractID string) (Account, Contract, string) {

	accountByte, errAccount := stub.GetState(accountID)
	var recordError string
	if errAccount != nil {
		recordError = "error while getting state Account"
		return Account{}, Contract{}, recordError
	}

	accountStruct := Account{}
	error1 := json.Unmarshal(accountByte, &accountStruct)
	if error1 != nil {
		recordError = "error while unmarshalling accountbyte"
		return Account{}, Contract{}, recordError
	}

	contractByte, errContract := stub.GetState(contractID)
	if errContract != nil {
		recordError = "error while getting state contract"
		return Account{}, Contract{}, recordError
	}

	contractStruct := Contract{}
	error2 := json.Unmarshal(contractByte, &contractStruct)

	if error2 != nil {

		recordError = "error while unmarshalling contract"
		return Account{}, Contract{}, recordError
	}

	return accountStruct, contractStruct, ""
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
