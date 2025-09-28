package main

import (
	"errors"
	"fmt"
)

const (
	UsdToEur = 0.95
	UsdToRub = 94.20
)

func main() {
	eurToRub := UsdToRub / UsdToEur

	fmt.Println("📊 Exchange rates")
	fmt.Printf("1 USD = %.2f EUR\n", UsdToEur)
	fmt.Printf("2 USD = %.2f RUB\n", UsdToRub)
	fmt.Printf("3 EUR = %.2f RUB\n", eurToRub)

	for {
		source, err := getCurrency("Enter source currency")
		if err != nil {
			fmt.Println(err)
			continue
		}
		amount, err := getAmount("Enter amount")
		if err != nil {
			fmt.Println(err)
			continue
		}
		target, err := getCurrency("Enter target currency")
		if err != nil {
			fmt.Println(err)
			continue
		}
		exchange := calculateExchange(amount, source, target)
		fmt.Printf("💱 Result: %.2f %s = %.2f %s\n\n", amount, source, exchange, target)
	}
}

func getCurrency(prompt string) (string, error) {
	var input string
	fmt.Printf("%s (USD,EUR,RUB): ", prompt)
	fmt.Scan(&input)
	if input == "USD" || input == "EUR" || input == "RUB" {
		return input, nil
	}
	return "", errors.New("❌ Invalid currency, try again.")
}

func getAmount(prompt string) (float64, error) {
	var input float64
	fmt.Printf("%s: ", prompt)
	fmt.Scan(&input)
	if input >= 0 {
		return input, nil
	}
	return 0, errors.New("❌ Invalid amount, try again.")
}

func calculateExchange(amount float64, source, target string) float64 {
	if source == target {
		return amount
	}
	switch source {
	case "USD":
		switch target {
		case "EUR":
			return amount * UsdToEur
		case "RUB":
			return amount * UsdToRub
		}
	case "EUR":
		switch target {
		case "USD":
			return amount / UsdToEur
		case "RUB":
			return amount * (UsdToRub / UsdToEur)
		}
	case "RUB":
		switch target {
		case "USD":
			return amount / UsdToRub
		case "EUR":
			return amount / (UsdToRub / UsdToEur)
		}
	}
	return 0
}
