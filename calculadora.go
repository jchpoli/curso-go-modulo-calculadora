package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct {
}

// Receive funtion
func (Calc) Operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	var resultado int
	switch operador {
	case "+":
		resultado = operador1 + operador2
	case "-":
		resultado = operador1 - operador2
	case "*":
		resultado = operador1 * operador2
	case "/":
		resultado = operador1 / operador2
	default:
		fmt.Println("Operador ", operador, " no soportado")
	}
	return resultado
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada(mensaje string) string {
	scanner := bufio.NewScanner(os.Stdin)
	if mensaje != "" {
		fmt.Println(mensaje)
	}
	scanner.Scan()
	return scanner.Text()
}

/*
func main() {
	entrada := leerEntrada("Ingrese la operacion:")
	operador := leerEntrada("Ingrese el operador:")
	resultado := Calc{}.Operate(entrada, operador)
	fmt.Println("Resultado:", resultado)
}
*/
