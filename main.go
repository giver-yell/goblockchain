package main

import (
	"fmt"
	"log"

	"github.com/giver-yell/goblockchain/block"
	"github.com/giver-yell/goblockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0)

	// blockchain
	blockchain := block.NewBlockchain(walletM.BlockChainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0, walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added? ", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockChainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockChainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockChainAddress()))
}
