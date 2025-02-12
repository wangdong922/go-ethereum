package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
)

func createKs() {
	ks := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "wangdong"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(account.Address.Hex()) // 0xd624959551B64A0e1FFe4a1E0393C589dD10e5B4
}

func importKs() {
	file := "./keystore/UTC--2025-02-12T02-06-06.943892000Z--d624959551b64a0e1ffe4a1e0393c589dd10e5b4"
	ks := keystore.NewKeyStore("./keystore/key", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	password := "wangdong"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//createKs()
	importKs()
}
