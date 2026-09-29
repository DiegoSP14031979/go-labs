package main

import "fmt"

func main() {
    pricePerItem := 8
    tax := 2
    quantity := 3
    discount := 5
    result := 0

    // TODO: calcula ((pricePerItem + tax) * quantity) - discount y guárdalo en result
    result = (pricePerItem + tax)*quantity - discount
    fmt.Println("Final bill:", result)
}
