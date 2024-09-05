package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func FlakyDivide(a, b int) (int, error) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Println(r.Float32())
	if r.Float32() < 0.2 { // Increased failure probability to 50%
		return 0, fmt.Errorf("random error occurred")
	}
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	
	// Simulate a delay of 0 to 1 second
	time.Sleep(time.Duration(r.Intn(500)) * time.Millisecond)

	return a / b, nil
}
