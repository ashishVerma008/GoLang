package main

import "fmt"

func main() {
	// 1. Array Declaration (Fills with Zero Values automatically)
	// Integer array defaults to 0
	var nums [4]int
	nums[0] = 10
	nums[1] = 20
	fmt.Println("Int Array:", nums)
	fmt.Println("Length of nums:", len(nums))

	// Boolean array defaults to false
	var vals [4]bool
	vals[2] = true
	fmt.Println("\nBool Array:", vals)

	// String array defaults to "" (empty string)
	var names [3]string
	names[0] = "Golang"
	names[1] = "Backend"
	fmt.Println("\nString Array:", names)

	// 2. Single Line Declaration & Initialization (Shorthand)
	scores := [3]int{95, 88, 72}
	fmt.Println("\nScores Array:", scores)

	// 3. 2D Arrays (Matrix)
	// A 2x2 grid initialized directly
	matrix := [2][2]int{
		{3, 4},
		{5, 6},
	}
	fmt.Println("\n2D Array (Matrix):", matrix)
	fmt.Println("Element at row 0, col 1:", matrix[0][1]) // Prints 4
}

/*
========================================
REVISION NOTES FOR ARRAYS:
1. Arrays have a FIXED size. Once declared as [4]int, it cannot grow or shrink.
2. Uninitialized elements are filled with 'Zero Values' (0, false, ""). NO garbage values!
3. len(arrayName) gives the size of the array.
4. IMPORTANT: In Go, Arrays are VALUE types. Assigning an array to another copies all its elements.
5. [3]int and [4]int are completely different types in Go.
6. Pros: Predictable memory size, constant time access O(1).
========================================
*/