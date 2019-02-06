package main

import (
	
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var (
	fileName = "chaincode"
)

type PersonalInfo struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	DOB       string `json:"DOB"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
}


type DocumentInfo struct { 
	aadhar_num int `json:"aadhar"`
}

type KYCapplication struct{
	DocumentInfo DocumentInfo `json:"DocumentInfo"`
	Status string `json:"status"`
}


type SampleChaincode struct {
}


func (t *SampleChaincode) createKYC(stub shim.ChaincodeStubInterface, args []string)pb.Response{
	fmt.Printf("\n Entering method %s", "createKYC")
	//validate input
	if args == nil || len(args) < 2 {
		
		fmt.Errorf("Insufficent input arguments. ")
		return shim.Error("Expected two arguments.")
	}
	var username = args[0]
	var aadhar_num = args[1]

	if len(args) == 3 {
		var passport = argv[2]
	}

	if len(aadhar_num) < 12 {
		fmt.Errorf("Aadhar number too small. ")
		return shim.Error("Aadhar number too small.")
	}

	if len(passport) < 8 {
		fmt.Errorf("Passport invalid. ")
		return shim.Error("Passport invalid.")
	}

	err := stub.PutState(username, []byte(aadhar_num))
	if err != nil {
		fmt.Errorf("Could not save KYC%s ", username)
		return shim.Error("Could not save KYC " + username)
	}
	fmt.Printf("\n Successfully saved KYC %s ", username)

	laBytes, err := stub.GetState(username)
	if err != nil {
		fmt.Errorf("\n File %s Method: %s Msg: %s ", fileName, "createLoanApplication", "Could not fetch loan application "+username+" from ledger")
		return shim.Error("Could not fetch loan application with ID " + username)
	}

	return shim.Success(laBytes)
}



func (t *SampleChaincode) updateKYC(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("Entering updateKYC")

	if len(args) < 2 {
		fmt.Errorf("Invalid number of args")
		return shim.Error("Expected atleast two arguments for KYC application update")
	}

	var username = args[0]
	var aadhar_num = args[1]

	laBytes, err := stub.GetState(username)
	if err != nil {
		fmt.Errorf("Could not fetch aadhar number from ledger %s", err)
		return shim.Error("Could not fetch aadhar number from ledger")
	}
	var kycApplication KYCapplication


	err = json.Unmarshal(laBytes, &kycApplication)
	kycApplication.Status = aadhar_num


	laBytes, err = json.Marshal(&kycApplication)
	if err != nil {
		fmt.Errorf("Could not marshal loan application post update %s", err)
		return shim.Error("Could not marshal loan application post update")
	}

	err = stub.PutState(username, laBytes)
	if err != nil {
		fmt.Errorf("Could not save loan application post update %s", err)
		return shim.Error("Could not save loan application post update")
	}


	var customEvent = "{eventType: 'loanApplicationUpdate', description:" + username + "' Successfully updated status'}"
	err = stub.SetEvent("evtSender", []byte(customEvent))
	if err != nil {
		return shim.Error("Could not set event")
	}
	fmt.Println("Successfully updated loan application")
	return shim.Success(laBytes)

}


func (t *SampleChaincode) getLoanApplicationByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Printf("\n Entering getLoanApplication")

	if args == nil || len(args) < 1 {
		fmt.Errorf("\n File: %s Method: %s Msg: %s ", fileName, "getLoanApplication", "Invalid input.")
		return shim.Error("Invalid loan application ID")
	}
	loanApplicationID := args[0]
	laBytes, err := stub.GetState(loanApplicationID)
	if err != nil {
		fmt.Errorf("\n File %s Method: %s Msg: %s ", fileName, "getLoanApplication", "Could not fetch KYC details"+loanApplicationID+" from ledger")
		return shim.Error("Could not fetch loan application with ID " + loanApplicationID)
	}

	return shim.Success(laBytes)
}

func (t *SampleChaincode) createTestData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Printf("\n Entering method %s", "createTestData")
	
	var data = `{"name":"Varun Ojha","company":"IBM"}`
	var dataID= "varun"
	fmt.Printf(dataID)
	fmt.Printf(data)
	err := stub.PutState(dataID, []byte(data))
	if err != nil {
		fmt.Errorf("Could not save data with id %s ", dataID)
		return shim.Error("Could not save data with ID " + dataID)
	}

	fmt.Printf("\n Successfully saved data with ID %s", dataID)
	laBytes, err := stub.GetState(dataID)
	if err != nil {
		fmt.Errorf("File %s Method: %s Msg: %s ", fileName, "createTestData", "Could not fetch data "+dataID+" from ledger")
		return shim.Error("Could not fetch loan application with ID " + dataID)
	}

	return shim.Success(laBytes)
}

func (t *SampleChaincode) getTestDataByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Printf("\n Entering getTestDataByID")


	if args == nil || len(args) < 1 {
		fmt.Errorf("File: %s Method: %s Msg: %s ", fileName, "getTestDataByID", "Invalid input.")
		return shim.Error("Invalid loan application ID")
	}
	loanApplicationID := args[0]
	laBytes, err := stub.GetState(loanApplicationID)
	if err != nil {
		fmt.Errorf("File %s Method: %s Msg: %s ", fileName, "getTestDataByID", "Could not fetch data "+loanApplicationID+" from ledger")
		return shim.Error("Could not fetch loan application with ID " + loanApplicationID)
	}

	return shim.Success(laBytes)
}


func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("Entering Init")
	fmt.Println("SampleChaincode Init Completed")
	return shim.Success(nil)
}


func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Entering Invoke")
	
	function, args := stub.GetFunctionAndParameters()
	if function == "CreateKYC" {
		return t.createKYC(stub, args)
	}else if function == "UpdateKYC" {
		return t.updateKYC(stub, args)
	}else if function == "GetKYC" {
		return t.getLoanApplicationByID(stub, args)
	} else if function == "CreateTestData" {
		return t.createTestData(stub, args)
	} else if function == "GetTestDataById" {
		return t.getTestDataByID(stub, args)
	}else if function == "init" {
		return t.Init(stub)
	} else {
		return shim.Error("Unsupported function " + function)
	}

}

func main() {
	
	fmt.Println("Entering main() for SampleChaincode")
	

		err := shim.Start(new(SampleChaincode))
		if err != nil {
			fmt.Errorf("Could not start chaincode %s ", "SampleChaincode")
			fmt.Errorf("SampleChaincode error %s",err)
			
		} else {
			fmt.Printf("\n Chaincode started %s", "SampleChaincode")
		}
	}


