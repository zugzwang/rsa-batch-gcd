package main

import (
	"fmt"
	"github.com/zugzwang/batchgcd"
	"os"
)

func main() {
	collectModuli()
}

func collectModuli() {
	fmt.Println("Collecting moduli")
	// Files to collect moduli from openpgp armored keys
	output2048, err := os.Create("data/moduli/moduli2048.txt")
	if err != nil {
		panic(err)
	}
	output3072, err := os.Create("data/moduli/moduli3072.txt")
	if err != nil {
		panic(err)
	}
	output4096, err := os.Create("data/moduli/moduli4096.txt")
	if err != nil {
		panic(err)
	}
	// Collect moduli
	defer func() {
		output2048.Close()
		output3072.Close()
		output4096.Close()
	}()
	err = batchgcd.ModuliFromDir("./data/keys/", output2048, output3072, output4096)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done.")
}
