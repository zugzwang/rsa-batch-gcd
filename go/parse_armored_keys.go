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
	output, err := os.Create("data/moduli/collected_moduli.txt")
	if err != nil {
		panic(err)
	}
	// Collect moduli
	defer func() {
		output.Close()
	}()
	err = batchgcd.ModuliFromDir("./data/keys/", output)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done.")
}
