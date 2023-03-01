package cartrack 

import (
	"bytes"
	"crypto/sha256"
	
	"encoding/json"
	"encoding/base64"
	
	"fmt"
	
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

type SmartContract struct {
	contractapi.Contract
}

type Car struct {
	
	Name         string				`json:"name"`
	Manufactor   string             `json:"manufactor"`
	Model        string				`json:"model"`
	Price        int                `json:"price"`
	Dealer       string             `json:"dealer"`
	Customer     string				`json:"customer"`
	Status       string             `json:"status"`
}


//car is manufactured by the manufacturer it is CREATED state
func (s *SmartContract) ManufactorCar(ctx contractapi.TransactionContextInterface, carId string, name string, model string) error {


	role,error = getRole(ctx)
	
	if err != nil {
		return "", fmt.Errorf("Failed to read clientID: %v", err)
	}
	
	if role != "Maker"{
		return "", fmt.Errorf("Permission is not to this clientID for creating new car: %v", err)
	}
	
	// get org of submitting client
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	
	if err != nil {
		return fmt.Errorf("failed to get client identity %v", err)
	}
	
	car := Car{
	
	Name         name,
	Manufactor   clientOrgID,
	Model        model,
	Price        10000,
	Dealer       "",
	Customer     "",
	Status       "CREATED",
		
	}
	
	carJSON, err := json.Marshal(car)
	if err != nil {
		return err
	}

	// put auction into state
	err = ctx.GetStub().PutState(carId, carJSON)
	if err != nil {
		return fmt.Errorf("failed to put auction in public data: %v", err)
	}

	return nil
}

// After it is delivered to a dealer it will be in READY_FOR_SALE state
func (s *SmartContract) deliverToDealer(ctx contractapi.TransactionContextInterface, carId string) (*Car, error) {

	
	role,error = getRole(ctx)
	
	if err != nil {
		return "", fmt.Errorf("Failed to read clientID: %v", err)
	}
	
	if role != "Maker"{
		return "", fmt.Errorf("Permission is not to this clientID for deliver new car: %v", err)
	}
	
	// get org of submitting client
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	
	if err != nil {
		return fmt.Errorf("failed to get client identity %v", err)
	}
	
	
	carAsBytes, err := ctx.GetStub().GetState(carId)
	
	car := new(Car)
	_ = json.Unmarshal(carAsBytes, car)
	
	if car.Status != "CREATED"{
		return fmt.Errorf("Car is not available for sale.", carId)
	}
	
	
	car.Dealer = clientOrgID
    car.Status = " READY_FOR_SALE"

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carId, carAsBytes)
}

// Once it is sold to a customer it will be in SOLD state
func (s *SmartContract) saleToCustomer(ctx contractapi.TransactionContextInterface, carId string) (*Car, error) {

	
	role,error = getRole(ctx)
	
	if err != nil {
		return "", fmt.Errorf("Failed to read clientID: %v", err)
	}
	
	if role != "Dealer"{
		return "", fmt.Errorf("Permission is not to this clientID for sale new car: %v", err)
	}
	
	// get org of submitting client
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	
	if err != nil {
		return fmt.Errorf("failed to get client identity %v", err)
	}
	
	carAsBytes, err := ctx.GetStub().GetState(carId)
	
	car := new(Car)
	_ = json.Unmarshal(carAsBytes, car)
	
	if car.Status != "CREATED"{
		return fmt.Errorf("Car is not available for sale.", carId)
	}
	
	
	car.Customer = clientOrgID
    car.Status = "SOLD"

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carId, carAsBytes)
}


//Validation of permisions.
func (s *SmartContract) getRole(ctx contractapi.TransactionContextInterface) (string role, error) {
	
	b64ID, found, err := ctx.GetClientIdentity().GetAttributeValue()
	
	if err != nil {
		return "", fmt.Errorf("Failed to read clientID: %v", err)
	}
	decodeID, err := base64.StdEncoding.DecodeString(b64ID)
	
	if err != nil {
		return "", fmt.Errorf("failed to base64 decode clientID: %v", err)
	}
	
	role =  string(decodeID)
	
	return role;
}



