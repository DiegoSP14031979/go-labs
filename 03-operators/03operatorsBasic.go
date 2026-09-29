package main

import "fmt"

func main() {
    totalPoints := 23
    rounds := 4

    // TODO: declara average como totalPoints / rounds (división entera)
    var average int = totalPoints / rounds

    totalPointsF := 23.0
    roundsF := 4.0

    // TODO: declara preciseAverage como totalPointsF / roundsF
    var preciseAverage float64 = totalPointsF / roundsF

    fmt.Println("Average (int):", average)
    fmt.Println("Average (precise):", preciseAverage)
}
