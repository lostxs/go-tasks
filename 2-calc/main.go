package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var menu = map[string]func([]float64) float64{
	"AVG": calcAvg,
	"SUM": calcSum,
	"MED": calcMedian,
}

func main() {
	fmt.Println("___ CALCULATOR ___")
MENU:
	for {
		operation := scanOperation()
		if operation == "q" {
			break MENU
		}
		menuFunc, ok := menu[operation]
		if !ok {
			fmt.Println("Unknown operation")
			continue MENU
		}
	NUMBERS:
		for {
			numbers, err := scanNumbers()
			if err != nil {
				fmt.Println(err)
				continue NUMBERS
			}
			result := menuFunc(numbers)
			fmt.Printf("Result: %.2f\n", result)
			break NUMBERS
		}
	}
}

func scanOperation() string {
	fmt.Print("Enter operation (AVG,SUM,MED) or q for exit: ")
	var operation string
	fmt.Scan(&operation)
	return operation
}

func scanNumbers() ([]float64, error) {
	fmt.Print("Enter numbers separated by commas: ")
	var input string
	fmt.Scan(&input)
	numbers := []float64{}
	for part := range strings.SplitSeq(input, ",") {
		num, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func calcAvg(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	var sum float64
	for _, d := range data {
		sum += d
	}
	return sum / float64(len(data))
}

func calcSum(data []float64) float64 {
	var sum float64
	for _, num := range data {
		sum += num
	}
	return sum
}

func calcMedian(data []float64) float64 {
	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)
	sort.Float64s(dataCopy)
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		return (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		return dataCopy[l/2]
	}
}
