package main

import (
	"github.com/joho/godotenv"
	"golang-btc-wallet/pkg/hdwallet"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load environment variable %v", err)
	}

	log.Println("BTC WALLET SERVER START")

	seed, err := hdwallet.MnemonicToSeed(os.Getenv("MNEMONIC"), "")
	if err != nil {
		log.Fatalf("failed to get seed: %v", err)
	}
	params := hdwallet.GetTestNetworkParams()
	wallet, err := hdwallet.CreateHDWallet(seed, params)
	if err != nil {
		log.Fatalf("failed to create hd wallet : %v", err)
	}

	log.Println(wallet.String())
	log.Println(wallet.Address(params))

	privateKey, err := hdwallet.GetPrivateKey(wallet)
	if err != nil {
		log.Fatalf("failed to get private key : %v", err)
	}

	log.Println(privateKey)

	p2shAddress, err := hdwallet.CreateP2SHAddress(wallet, params)
	if err != nil {
		log.Fatalf("failed to create p2sh address : %v", err)
	}

	log.Println(p2shAddress)

	p2wpkhAddress, err := hdwallet.CreateP2WPKHAddress(wallet, params)
	if err != nil {
		log.Fatalf("failed to create p2wpkh address : %v", err)
	}

	log.Println(p2wpkhAddress)

	child := hdwallet.GetChildFromRoot(wallet, 1)
	log.Println(child.String())
}
