package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/DaveAppleton/ether_go/ethKeys"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	currentTime := time.Now()
	dateStr := currentTime.Format("01-02-2006")
	granny := ethKeys.NewKey("keys/granny")
	if err := granny.RestoreOrCreate(); err != nil {
		log.Fatal("Cannot revive granny", err)
	}
	tom := ethKeys.NewKey("keys/tom")
	if err := tom.RestoreOrCreate(); err != nil {
		log.Fatal("Cannot revive tom", err)
	}
	fmt.Println("Tom ", tom.PublicKeyAsHexString())
	ownerTx := bind.NewKeyedTransactor(tom.GetKey())
	dateBytes := []byte(dateStr)
	dateHash := crypto.Keccak256(dateBytes)
	fullToSign := append([]byte("\x19Ethereum Signed Transaction\n32"), dateHash...)
	fullHash := crypto.Keccak256(fullToSign)
	sig, err := crypto.Sign(fullHash, granny.GetKey())
	if err != nil {
		log.Fatal("cannot sign : ", err)
	}
	client, err := ethclient.Dial("https://kovan.infura.io/DaveAppleton")
	if err != nil {
		log.Fatal("Cannot connect", err)
	}
	contractAddress := common.HexToAddress("0x5D58398589de1a119117a5271EA8f75Db192fcD3")
	helper, err := NewMain(contractAddress, client)
	if err != nil {
		log.Fatal("linking  contract ", err)
	}
	tx, err := helper.HelperVisit(ownerTx, dateStr, sig)
	if err != nil {
		log.Fatal("HelperVisit error ", err)
	}
	log.Println(tx.Hash().Hex())
}
