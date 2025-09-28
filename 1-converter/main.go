package main

import "fmt"

const (
	UsdToEur = 0.95
	UsdToRub = 94.20
)

func main() {
	eurToRub := UsdToRub / UsdToEur

	fmt.Println("Exchange rates")
	fmt.Printf("1 USD = %.2f EUR\n", UsdToEur)
	fmt.Printf("2 USD = %.2f RUB\n", UsdToRub)
	fmt.Printf("3 EUR = %.2f RUB\n", eurToRub)
}

func getUserInput() {
	var val float64
	fmt.Print("Enter value: ")
	fmt.Scan(&val)
}

func calculateExchange(val int, original, target string) {}
