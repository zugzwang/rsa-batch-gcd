package main

import (
	"fmt"
	"github.com/zugzwang/batchgcd"
	"os"
)

func main() {
	moduliFile, err := os.Open("data/moduli/moduli2048.txt")
	if err != nil {
		panic(err)
	}
	defer moduliFile.Close()
	gcds, compromised, err := batchgcd.BatchGcdFromFile(moduliFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d keys analyzed. Compromised keys:\n", len(gcds))
	fmt.Println(compromised)
}
