package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divisor cannot be 0")

func main() {
	defer func() {
		fmt.Println("	[main] - deferred")
		if err := recover(); err != nil {
			fmt.Println("Shutdown due to panic, err :", err)
			return
		}
		fmt.Println("	[main] - Thank You!")
	}()
	var divisor int
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		if q, r, err := divideWrapper(100, divisor); err != nil {
			fmt.Println("Error :", err)
			continue
		} else {
			fmt.Println(q, r)
			break
		}
	}
}

// wrapper function to conver the panic into an error inorder to take a different course of action
func divideWrapper(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api (that panics)
func divide(x, y int) (quotient, remainder int) {
	defer func() {
		fmt.Println("	[divide] - deferred")
	}()
	fmt.Println("[divide] - Calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y

	fmt.Println("[divide] - Calculating remainder")
	remainder = x % y

	return
}
