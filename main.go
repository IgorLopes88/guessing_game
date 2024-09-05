package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	red := color.New(color.FgRed)
	green := color.New(color.FgHiGreen)
	blue := color.New(color.FgHiBlue)
	yellow := color.New(color.FgYellow)

	red.Println("JOGO DA ADIVINHAÇÃO")
	fmt.Println("Um número aleatório será sorteado. 🎯 Tente acertar!")
	fmt.Println("O número é inteiro e entre 0 e 100.")

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}
	x := rand.Int64N(101)

	for i := range chutes {
		blue.Println("\n🔎 Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um número inteiro...")
			return
		}

		switch {
		case chuteInt < x:
			red.Println("❌ Você errou!")
			yellow.Println("Dica: O número sorteado é maior que", chuteInt)
		case chuteInt > x:
			red.Println("❌ Você errou!")
			yellow.Println("Dica: O número sorteado é menor que", chuteInt)
		case chuteInt == x:
			green.Println("\nParabéns! Você acertou!🥳")
			blue.Printf(
				"✨ O número era %d ✨\n"+
					"Você acertou em %d tentativas.\n\n"+
					"Essas foram suas tentativas: %v\n",
				x, i+1, chutes[:i],
			)
			return
		}

		chutes[i] = chuteInt
	}
	red.Printf(
		"\nInfelizmente, você não acertou o número, que era %d 😅\n",
		x,
	)
	blue.Printf(
		"Você teve 10 tentativas.\n"+
			"Essas foram suas tentativas: %v\n"+
			" ",
		chutes,
	)
}
