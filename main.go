package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	maximo := flag.Int("m", 10, "Número máximo do sorteio")
	numeroEscolhido := flag.Int("c", 1, "Número escolhido.")
	flag.Parse()

	fmt.Println("Sorteando...")
	time.Sleep(10 * time.Second)

	numeroSorteado := rand.Intn(*maximo) + 1

	fmt.Printf("O valor sorteado foi %d\n", numeroSorteado)

	if *numeroEscolhido == numeroSorteado {
		fmt.Println("O número é igual o que você escolheu, PARABÉNS !!!")
	} else {
		fmt.Println("Não foi dessa vez. Tenta de novo !!!")
		os.Exit(1)
	}
}
