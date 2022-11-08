package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		return (operador1 + operador2)
	case "-":
		return (operador1 - operador2)
	case "*":
		return (operador1 * operador2)
	case "/":
		return (operador1 / operador2)
	default:
		fmt.Println(operador, "Operador no esta soportado")
		return 0
	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() string {
	scarnner := bufio.NewScanner(os.Stdin)
	scarnner.Scan()
	return scarnner.Text()
}

func main() {
	entrada := leerEntrada()

	operador := leerEntrada()

	c := calc{}

	fmt.Println(c.operate(entrada, operador))

}
