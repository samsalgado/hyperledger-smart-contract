package main

import (
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
	 type SmartContract struct {
	 contractapi.Contract
	 }

type Token interface{
	addToken() int 
}
type TxSubmission interface {
	passedChecks() bool
}
type Mywallet struct {
	passed bool
	balance int 
}
func (w Mywallet) addToken() int {
	return w.balance + 1
}

func (s *SmartContract) NewBalance(w Token) {


	if transactionSubmitted, ok := w.(TxSubmission); ok {
	fmt.Println(transactionSubmitted.passedChecks())
	fmt.Println(w.addToken())
}
}
func main() {
	assetChaincode, err := contractapi.NewChaincode(&SmartContract{})
    if err != nil {
      log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
    }

    if err := assetChaincode.Start(); err != nil {
      log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
    }


}