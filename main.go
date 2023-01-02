package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("jogos.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo " + err.Error())
	}

	fileContent := string(file)

	numeros := os.Args[1:]

	fmt.Println(numeros)

	var valor string
	for index, number := range numeros {
		if index == len(numeros)-1 {
			valor += number
			break
		}
		valor += fmt.Sprintf("%s\r\n", number)
	}

	confereJogo(fileContent, valor)

}

func confereJogo(fileContent string, valor string) {
	venceu := strings.Contains(fileContent, valor)

	if venceu {
		fmt.Println("Voce ganhou milhoes")
	} else {
		fmt.Println("Voce perdeu, tenta ano que vem")
	}
}
