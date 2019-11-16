// Copyright (C) 2019 ProtonTech AG

package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
	"golang.org/x/crypto/rsa"
	"os"
	"strconv"
)

func main() {
	// Create directory if it doesn't exist
	dir := "data/moduli/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, 0755); err != nil {
			panic(err)
		}
	}
	var n int
	fmt.Println("Number of keys?")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}
	fileName := dir + strconv.Itoa(n) + "-moduli.txt"
	modFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	config := &packet.Config{
		RSABits:   2048,
		Algorithm: packet.PubKeyAlgoRSA,
	}
	name := "tester"
	comment := "batch-gcd"
	email := "zug@zwang.com"

	defer modFile.Close()
	// Generate keys and write to file
	for i := 0; i < n; i++ {
		key, errKG := openpgp.NewEntity(name, comment, email, config)
		if errKG != nil {
			panic(errKG)
		}
		mod := key.PrimaryKey.PublicKey.(*rsa.PublicKey).N.Bytes()
		mod64 := base64.StdEncoding.EncodeToString(mod)
		fmt.Println(mod64)
		modFile.WriteString(mod64)
		modFile.WriteString("\n")
		for _, subkey := range key.Subkeys {
			subMod := subkey.PublicKey.PublicKey.(*rsa.PublicKey).N.Bytes()
			subMod64 := base64.StdEncoding.EncodeToString(subMod)
			modFile.WriteString(subMod64)
			fmt.Println(subMod64)
			modFile.WriteString("\n")
		}
	}
}
