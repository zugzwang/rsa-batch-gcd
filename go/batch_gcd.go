package main

import (
	"fmt"
	"github.com/zugzwang/batchgcd"
	"time"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Give a target moduli file (ex: data/moduli/mod.txt)")
		return
	}
	moduliFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer moduliFile.Close()
	start := time.Now()
	gcds, compromised, err := batchgcd.BatchGcdFromFile(moduliFile, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d keys analyzed. Compromised keys:\n", len(gcds))
	fmt.Println(len(compromised))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
