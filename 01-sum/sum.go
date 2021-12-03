package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Método do pacote bufio para ler os inputs do usuário

	for {
		fmt.Print("First number: ")
		scanner.Scan() // Lê o que foi digitado na linha de comando
		n1 := scanner.Text() // Pega o texto que foi digitado e atribui seu valor a variável

		fmt.Print("Second number: ")
		scanner.Scan();
		n2 := scanner.Text()

		A, _ := strconv.Atoi(n1) // Converte o tipo texto para integer
		B, _ := strconv.Atoi(n2) // Converte o tipo texto para integer

		X := A + B

		fmt.Printf("X = %d\n", X) // Exibe a variavel na string
	}
}