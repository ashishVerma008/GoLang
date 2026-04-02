package main

import "fmt"

// 1. Variadic Function: Accepts any number of integers
func sum(nums ...int) int {
	// Inside the function, 'nums' is just a slice of ints: []int
	total := 0
	
	for _, num := range nums {
		total = total + num
	}
	
	return total
}

// 2. Mixed Parameters: Variadic MUST be the last parameter
func printReceipt(customerName string, prices ...int) {
	total := sum(prices...) // Passing variadic arguments to another variadic function
	fmt.Printf("Customer: %s | Total Bill: $%d\n", customerName, total)
}

func main() {
	// 1. Passing comma-separated values directly (Packing)
	fmt.Println("--- Passing individual values ---")
	res1 := sum(10, 20)
	res2 := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of 10, 20:", res1)
	fmt.Println("Sum of 1,2,3,4,5:", res2)

	// 2. Passing an existing slice (Unpacking)
	fmt.Println("\n--- Passing a slice ---")
	mySlice := []int{3, 4, 5, 6}
	// We use '...' after the slice to unpack it into individual arguments
	res3 := sum(mySlice...) 
	fmt.Println("Sum of slice:", res3)

	// 3. Using the mixed parameter function
	fmt.Println("\n--- Mixed Parameters ---")
	printReceipt("Ashish", 100, 50, 200) 
}

/*
========================================
REVISION NOTES FOR VARIADIC FUNCTIONS:
1. Theory: A function that can take zero or multiple arguments of a specific type.
2. Syntax: Use '...' before the type in the function definition (e.g., nums ...int).
3. Under the hood: The variadic parameter behaves exactly like a Slice.
4. Unpacking: If you already have a slice and want to pass it to a variadic function, use '...' after the variable name (e.g., sum(mySlice...)).
5. STRICT RULE: A variadic parameter must ALWAYS be the last parameter in the function signature, and you can only have one per function.
========================================
*/