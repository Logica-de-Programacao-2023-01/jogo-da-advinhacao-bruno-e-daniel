package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
	var allTries [][]int
	Jogo(&allTries)
	PrintAllTries(allTries)
}

func Jogo(allTries *[][]int) {
	numero := false
	random := rand.Intn(100) + 1
	tries := 0

	var choose int

	fmt.Println("Bem-vindo ao jogo da adivinhação!")

	for !numero {
		fmt.Print("Digite um número entre 1 e 100: ")
		fmt.Scan(&choose)
		tries++

		if choose > 100 || choose <= 0 {
			fmt.Println("Número inválido.")
		} else if choose < random {
			fmt.Println("O número é maior que", choose)
		} else if choose > random {
			fmt.Println("O número é menor que", choose)
		} else {
			fmt.Println("Parabéns, você acertou!")
			fmt.Printf("Você utilizou %d tentativa(s).\n", tries)
			numero = true
		}
	}

	*allTries = append(*allTries, []int{tries})

	fmt.Println("Você deseja jogar novamente? (s/n):")
	var jogarnovamente string
	fmt.Scan(&jogarnovamente)

	if jogarnovamente != "s" && jogarnovamente != "S" {
		fmt.Println("Obrigado por jogar.")
		return
	}

	Jogo(allTries)
}

func PrintAllTries(allTries [][]int) {
	if len(allTries) > 0 {
		for i, tries := range allTries {
			fmt.Printf("Você utilizou %d tentativas na %dª jogada.\n", tries[0], i+1)
		}
	} else {
		fmt.Println("Nenhuma tentativa registrada.")
	}
}
