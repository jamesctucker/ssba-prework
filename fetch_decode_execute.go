package main

// import a timer
import (
	"fmt"
	"time"
)

// a gigahertz is a unit of a billion
var rounds = 1000000000

func main() {
	result1 := emptyLoops()
	result2 := loopsWithMath()
	// print format result 1 and result 2
	fmt.Printf("Empty Loops: %.2f GHz\n", result1 / 1000000000.0)
	fmt.Println("----------------------")
	fmt.Printf("Loops with Math: %.2f GHz\n", result2 / 1000000000.0)
}

func emptyLoops() float64{
	start := time.Now();

	for i := 0; i < rounds; i++ {
		// do nothing
	}

	end := time.Now();
	elapsed := end.Sub(start).Seconds()
	operationsPerSecond := float64(rounds) / elapsed

	return operationsPerSecond
}

func loopsWithMath() float64 {
	num := 0
	start := time.Now();

	for i := 0; i < rounds; i++ {
		num*=2
	}

	end := time.Now();
	elapsed := end.Sub(start).Seconds()
	operationsPerSecond := float64(rounds) / elapsed

	return operationsPerSecond
}