package main

import (
	"fmt"
	"log"
	"time"
)

// ver 1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

// ver 2.0
// wrappers for logging
/*
func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}
*/

// Applying "commonality - variability" for the above
func logOperation(op func(int, int), x, y int) {
	log.Println("Operation started")
	op(x, y)
	log.Println("Operation completed")
}

func getLogOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started")
		op(x, y)
		log.Println("Operation completed")
	}
}

// ver 3.0
/*
func getProfiledOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		op(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}
*/

// typed version
type OperationFn func(int, int)

func getProfiledOperation(op OperationFn) OperationFn {
	return func(x, y int) {
		start := time.Now()
		op(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
		logOperation(func(i1, i2 int) {
			fmt.Println("Multiply Result :", i1*i2)
		}, 100, 200)
	*/

	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	// composing log & profile
	logAdd := getLogOperation(add)
	profiledLogAdd := getProfiledOperation(logAdd)
	profiledLogAdd(100, 200)
}
