package main

import "fmt"

func main() {
	// Estas variables están declaradas pero no inicializadas
	// Go les asignará automáticamente sus valores cero
	var number int
	var decimal float64
	var isTrue bool
	var text string
	
	// TODO: Imprime cada variable para ver su valor cero
	// Hint: Usa fmt.Println() para cada variable
    fmt.Println(number)
    fmt.Println(decimal)
    fmt.Println(isTrue)
    fmt.Println(text)
}
