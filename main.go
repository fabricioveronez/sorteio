package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	maximo := flag.Int("m", 1, "Número escolhido")
	flag.Parse()

	fmt.Println("Sorteando...")
	time.Sleep(10 * time.Second)

	numeroSorteado := rand.Intn(*maximo-1) + 1
	fmt.Printf("O valor sorteado foi %d", numeroSorteado)
}
