package main

import "fmt"

// counter is a function that returns another function (which returns an int)
func counter() func() int {
	// 1. This variable is local to 'counter'.
	// Normally, it would be destroyed when 'counter' finishes.
	count := 0 

	// 2. This inner anonymous function forms a CLOSURE.
	// It "captures" the 'count' variable and keeps it alive.
	return func() int {
		count += 1
		return count
	}
}

func main() {
	// increment is now a function with its own private 'count' backpack
	increment := counter()

	fmt.Println("--- First Counter ---")
	fmt.Println(increment()) // Prints 1
	fmt.Println(increment()) // Prints 2
	fmt.Println(increment()) // Prints 3

	// If we call counter() again, it creates a BRAND NEW isolated state
	increment2 := counter()

	fmt.Println("\n--- Second Counter ---")
	fmt.Println(increment2()) // Prints 1 (Starts from scratch!)
	fmt.Println(increment2()) // Prints 2
	
	// The first counter is still at 3! They don't interfere.
	fmt.Println("\n--- First Counter Again ---")
	fmt.Println(increment()) // Prints 4
}

/*
========================================
REVISION NOTES FOR CLOSURES:
1. Theory: A closure is a function that "remembers" the variables from its outer scope, even after the outer function has finished executing.
2. Significance: It is used for Data Hiding/Encapsulation. It allows a function to have private state without using global variables or structs/classes.
3. Memory: The Go compiler keeps the outer variables alive in memory as long as the inner function (closure) is still being used.
4. Isolation: Calling the outer function multiple times creates independent, isolated states for each closure (e.g., increment and increment2 don't share the same 'count').
========================================
*/