package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("___ CALCULATOR ___")
	for {
		operation := scanOperation()
		if operation == "q" {
			break
		}
		input := scanInput()
		numbers, err := parseNumbers(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(operation, numbers)
	}
}

func outputResult(operation string, numbers []float64) {
	switch operation {
	case "AVG":
		average := calcAvg(numbers)
		result := fmt.Sprintf("The average is: %.2f", average)
		fmt.Println(result)
	case "SUM":
		sum := calcSum(numbers)
		result := fmt.Sprintf("The sum is: %.2f", sum)
		fmt.Println(result)
	case "MED":
		med := calcMedian(numbers)
		result := fmt.Sprintf("The median is: %.2f", med)
		fmt.Println(result)
	default:
		fmt.Println("Unknown operation")
	}
}

func scanOperation() string {
	var operation string
	fmt.Print("Enter operation (AVG,SUM,MED) or q for exit: ")
	fmt.Scan(&operation)
	return operation
}

func scanInput() string {
	var input string
	fmt.Print("Enter numbers separated by commas: ")
	fmt.Scan(&input)
	return input
}

func parseNumbers(value string) ([]float64, error) {
	numbers := []float64{}
	parts := strings.SplitSeq(value, ",")
	for part := range parts {
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
