package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Jogo da Advinhação")	
	fmt.Println("Um número aleatório será sorteado. Tente acertar. O número é um inteiro entre 0 e 100")	

	x:= rand.Int64N(101)
	fmt.Println("o numero é ", x)

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10] int64{}

	for i := range chutes{
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)
		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil{
			fmt.Println("O seu número precisa ser um inteiro")
			return
		}

		switch{
		case chuteInt < x:
			fmt.Println("Você errou. O número sorteado é menor que !", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou. O número sorteado é maior que !", chuteInt)
		case chuteInt == x:
				fmt.Printf("Acertou, parabéns ! O número é %d\n" +
				"Você acertou em %d tentativas.\n" +
				"Essas tentativas foram: %v\n", 
				x, i+1, chutes[:i])
			return 
		}

		fmt.Println("Suas chances", 10-1-i)
		fmt.Println("-----------------")
		chutes[i] = chuteInt
	}

	fmt.Printf("Infelizmente você não acertou o número que era: %d\n" + 
		"Você teve 10 tentativas.\n" +
		"Essas tentativas foram: %v", 
	x, chutes)
}