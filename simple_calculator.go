package main

import (
	"errors"
	"fmt"
	"os"
)

var x, y float64
var operator string

func add(x, y float64) float64 {
	return x + y
}

func subtraction(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) (float64, error) {
	if y != 0 {
		return x / y, nil
	}
	return 0, errors.New("Dividion by 0")
}

func exponentiation(x, y float64) float64 {
	var res float64 = 1
	for i := 1; i <= int(y); i++ {
		res *= x
	}
	return res
}

func doCalc(x, y float64, op string) (float64, error) {
	switch op {
	case "+":
		return add(x, y), nil
	case "-":
		return subtraction(x, y), nil
	case "*":
		return multiply(x, y), nil
	case "/":
		return divide(x, y)
	case "^":
		return exponentiation(x, y), nil
	}
	return 0, errors.New("Not identified operator")
}

func handleRequest() {
	for {
		fmt.Print("Enter request:\n")
		n, err := fmt.Scanf("%f%s%f\n", &x, &operator, &y)

		if n != 3 && err != nil {
			fmt.Println("exit")
			break
		} else {
			result, err := doCalc(x, y, operator)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
			fmt.Print("= ", result)
			fmt.Println()
		}
	}
}
