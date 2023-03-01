/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-samples/auction/chaincode-go/smart-contract"
)

func main() {
	cartrack, err := contractapi.NewChaincode(&cartrack.SmartContract{})
	if err != nil {
		log.Panicf("Error creating cartrack chaincode: %v", err)
	}

	if err := cartrack.Start(); err != nil {
		log.Panicf("Error starting cartrack chaincode: %v", err)
	}
}
