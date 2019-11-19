// Copyright (C) 2019 ProtonTech AG

package main

import (
	"encoding/base64"
	"math/big"
	"fmt"
	"crypto/rand"
	"os"
	"strconv"
)

func main() {
	// Create directory if it doesn't exist
	dir := "data/moduli/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}
	}
	var n, bitLen int
	fmt.Println("Number of keys?")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bit length of modulus? (Ex. 4096)")
	_, err = fmt.Scanf("%d", &bitLen)
	if err != nil {
		panic(err)
	}
	fileName := dir + strconv.Itoa(n) + "moduli" + strconv.Itoa(bitLen) + ".txt"
	modFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer modFile.Close()
	// Generate keys and write to file
	for i := 0; i < n; i++ {
		p, err := rand.Prime(rand.Reader, bitLen/2)
		if err != nil {
			panic(err)
		}
		q, err := rand.Prime(rand.Reader, bitLen/2)
		if err != nil {
			panic(err)
		}
		mod := new(big.Int)
		mod.Mul(p, q)
		mod64 := base64.StdEncoding.EncodeToString(mod.Bytes())
		bitLen := 2048
		fmt.Println(i, bitLen, mod64[:60] + "...")
		modFile.WriteString(strconv.Itoa(i))
		modFile.WriteString(",")
		modFile.WriteString(strconv.Itoa(int(bitLen)))
		modFile.WriteString(",")
		modFile.WriteString(mod64)
		modFile.WriteString("\n")
	}
	fmt.Println("NOTE: Written to " + fileName)
}
