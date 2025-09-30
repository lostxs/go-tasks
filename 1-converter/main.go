package main

import (
	"errors"
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
	currencies := joinCurrencies(&exchanges)
	for {
		source, err := getCurrency("Enter source currency", currencies, &exchanges)
		if err != nil {
			fmt.Println(err)
			continue
		}
		amount, err := getAmount()
		if err != nil {
			fmt.Println(err)
			continue
		}
		target, err := getCurrency("Enter target currency", currencies, &exchanges)
		if err != nil {
			fmt.Println(err)
			continue
		}
		exchange := calculateExchange(source, target, amount, &exchanges)
		fmt.Printf("Exchange result: %.2f %s = %.2f %s\n", amount, source, exchange, target)
	}
}

func joinCurrencies(exchanges *exchangesMap) string {
	var keys string
	first := true
	for key := range *exchanges {
		if !first {
			keys += ","
		}
		keys += key
		first = false
	}
	return keys
}

func getCurrency(prompt, currencies string, exchanges *exchangesMap) (string, error) {
	var curr string
	fmt.Printf("%s (%v): ", prompt, currencies)
	fmt.Scan(&curr)
	if _, ok := (*exchanges)[curr]; !ok {
		return "", errors.New("Invalid currency")
	}
	return curr, nil
}

func getAmount() (float64, error) {
	var amount float64
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)
	if amount >= 0 {
		return amount, nil
	}
	return 0, errors.New("Invalid amount")
}

func calculateExchange(source, target string, amount float64, exchanges *exchangesMap) float64 {
	if source == target {
		return amount
	}
	rate := (*exchanges)[source][target]
	return amount * rate
}
