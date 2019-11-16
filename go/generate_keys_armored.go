// Copyright (C) 2019 ProtonTech AG

package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
	"golang.org/x/crypto/openpgp/armor"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create directory if it doesn't exist
	dir = "data/1Mkeys/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, 0755); err != nil {
			panic(err)
		}
	}
	config := &packet.Config{
		RSABits:   2048,
		Algorithm: packet.PubKeyAlgoRSA,
	}
	name := "tester1M"
	comment := "batch-gcd"
	email := "test@test.com"

	// Generate keys
	for i := 0; i < 1000000; i++ {
		fmt.Println(i)
		entity, errKG := openpgp.NewEntity(name, comment, email, config)
		if errKG != nil {
			panic(errKG)
		}
		w := bytes.NewBuffer(nil)
		if err := entity.SerializePrivateNoSign(w, nil); err != nil {
			panic(err)
		}

		serialized := w.Bytes()

		privateKey, _ := armorWithType(serialized, "PGP PRIVATE KEY BLOCK")
		pk, _ := publicKey(privateKey)
		filename := dir + "rsa2048-1M-" + strconv.Itoa(i) + ".asc"
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		file.WriteString(pk)
		file.Close()
	}
}

// armorWithType make bytes input to armor format
func armorWithType(input []byte, armorType string) (string, error) {
	var b bytes.Buffer
	w, err := armor.Encode(&b, armorType, nil)
	if err != nil {
		return "", err
	}
	if _, err = w.Write(input); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	return b.String(), nil
}

func publicKey(privateKey string) (string, error) {
	privKeyReader := strings.NewReader(privateKey)
	entries, err := openpgp.ReadArmoredKeyRing(privKeyReader)
	if err != nil {
		return "", err
	}

	var outBuf bytes.Buffer
	for _, e := range entries {
		err := e.Serialize(&outBuf)
		if err != nil {
			return "", err
		}
	}

	outString, err := armorWithType(outBuf.Bytes(), "PGP PUBLIC KEY BLOCK")
	if err != nil {
		return "", err
	}

	return outString, nil
}
