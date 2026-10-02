package main

import "fmt"

func main() {
    totalCandies := 29
    kids := 4

    // TODO: declara leftover como totalCandies % kids
    var leftover int = totalCandies % kids

    // TODO: declara isLeftoverEven como (leftover % 2 == 0)
    var isLeftoverEven bool = leftover % 2 == 0

    fmt.Println("Leftover:", leftover)
    fmt.Println("Leftover is even:", isLeftoverEven)
}
