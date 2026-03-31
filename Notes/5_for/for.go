package main

import "fmt"

func main() {
	// 1. The "While" loop logic in Go
	fmt.Println("--- While Loop ---")
	i := 1
	for i <= 3 {
		fmt.Println("Value of i is:", i)
		i = i + 1 // or i++
	}

	// 2. The Classic For Loop with continue and break
	fmt.Println("\n--- Classic Loop ---")
	for j := 0; j <= 5; j++ {
		if j == 2 {
			continue // Skips the rest of the code below and goes to j=3
		}
		if j == 4 {
			break // Completely breaks out of the loop
		}
		fmt.Println("Value of j is:", j)
	}

	// 3. The New Go 1.22 'range' over integers
	fmt.Println("\n--- Go 1.22 Range ---")
	// Loops from 0 to 4 (runs 5 times)
	for k := range 5 {
		fmt.Println("Range value:", k)
	}

	// 4. Infinite Loop (Usually used for continuously running backend services)
	fmt.Println("\n--- Infinite Loop ---")
	count := 0
	for {
		fmt.Println("Running... (Infinite)")
		count++
		
		// Breaking out so it doesn't crash your terminal
		if count >= 2 {
			fmt.Println("Breaking out of infinite loop!")
			break 
		}
	}
}

/*
========================================
REVISION NOTES FOR LOOPS:
1. Go has NO 'while' or 'do-while' keywords. Only 'for' exists.
2. 'for condition {}' acts exactly like a while loop.
3. DO NOT use parentheses () around the loop condition.
4. 'break' stops the loop entirely.
5. 'continue' skips the current iteration and moves to the next.
6. 'for i := range N' (Go 1.22+) loops from 0 up to N-1.
7. 'for {}' is an infinite loop, heavily used in backend systems to listen for events.
========================================
*/