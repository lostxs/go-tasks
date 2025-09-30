package main

import (
	"fmt"
)

type exchangesMap = map[string]map[string]float64

func main() {
	const (
		UsdToEur = 0.95
		UsdToRub = 94.20
		EurToUsd = 1 / UsdToEur
		EurToRub = UsdToRub / UsdToEur
	)
	exchanges := exchangesMap{
		"USD": {
			"EUR": UsdToEur,
			"RUB": UsdToRub,
		},
		"EUR": {
			"USD": EurToUsd,
			"RUB": EurToRub,
		},
		"RUB": {
			"USD": 1 / UsdToRub,
			"EUR": UsdToEur / UsdToRub,
		},
	}

	fmt.Println("📊 Exchange rates")
	fmt.Printf("USD = %.2f RUB\n", UsdToRub)
	fmt.Printf("EUR = %.2f RUB\n", EurToRub)
	for {
		source := getSourceCurrency()
		if _, ok := exchanges[source]; !ok {
			fmt.Println("Invalid currency")
			continue
		}
		target := getTargetCurrency()
		if _, ok := exchanges[target]; !ok {
			fmt.Println("Invalid currency")
			continue
		}
		amount := getAmount()
		if amount < 0 {
			fmt.Println("Invalid amount")
			continue
		}
		exchange := calculateExchange(source, target, amount, exchanges)
		fmt.Println(exchange)
	}
}

func getSourceCurrency() string {
	var source string
	fmt.Printf("Enter source currency (USD,EUR,RUB): ")
	fmt.Scan(&source)
	return source
}

func getTargetCurrency() string {
	var target string
	fmt.Printf("Enter target currency (USD,EUR,RUB): ")
	fmt.Scan(&target)
	return target
}

func getAmount() float64 {
	var amount float64
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)
	return amount
}

func calculateExchange(source, target string, amount float64, exchanges exchangesMap) string {
	if source == target {
		return fmt.Sprintf("%.2f %s = %.2f %s", amount, source, amount, target)
	}
	rate := exchanges[source][target]
	exchange := amount * rate
	return fmt.Sprintf("%.2f %s = %.2f %s", amount, source, exchange, target)
}
