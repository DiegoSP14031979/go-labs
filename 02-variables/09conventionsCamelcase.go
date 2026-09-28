package main

import "fmt"

// TODO: renombrar a AppName para que se exporte
var AppName string = "Tracker"

func main() {
    // TODO: renombrar estos cuatro a camelCase, mantén los valores
    var userId int = 42
    var accountBalance float64 = 150.75
    var isActiveNow bool = true
    var fullName string = "Jamie Lee"

    fmt.Println("App:", AppName)
    fmt.Println("User ID:", userId)
    fmt.Println("Balance:", accountBalance)
    fmt.Println("Active:", isActiveNow)
    fmt.Println("Name:", fullName)
}
