package main

import "fmt"

func main() {
    var day, year int
    var month string
    fmt.Scanf("%d %s %d", &day, &month, &year)
    fmt.Printf("captured: %d %s %d\n", day, month, year)
}
