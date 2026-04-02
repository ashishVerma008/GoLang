package main

import "fmt"

func main() {
	// 1. Iterating over Slices
	fmt.Println("--- Range over Slice ---")
	nums := []int{6, 7, 8}
	
	// 'i' is the index, 'num' is the value
	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	// Example: Using Blank Identifier (_) if we don't need the index
	sum := 0
	for _, num := range nums {
		sum = sum + num // Only using the value
	}
	fmt.Println("Sum of elements:", sum)


	// 2. Iterating over Maps
	fmt.Println("\n--- Range over Map ---")
	m := map[string]string{"fname": "john", "lname": "doe"}

	// 'k' is Key, 'v' is Value
	for k, v := range m {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	// If you only need keys, just provide one variable
	fmt.Println("\nPrinting only keys:")
	for k := range m {
		fmt.Println("Key:", k)
	}


	// 3. Iterating over Strings (Unicode Runes)
	fmt.Println("\n--- Range over String ---")
	word := "golang"
	
	// 'i' is the starting byte index, 'c' is the Unicode Rune (int32 representation of the character)
	for i, c := range word {
		// We use string(c) to convert the integer rune back to a readable string character
		fmt.Printf("Starting Byte Index: %d, Rune Value: %v, Character: %s\n", i, c, string(c))
	}
}

/*
========================================
REVISION NOTES FOR RANGE:
1. Theory: 'range' simplifies looping over slices, maps, and strings. Same as C++ range-based loop, but provides both index/key and value.
2. Syntax: for index, value := range collection { ... }
3. Blank Identifier (_): If you don't need the index or value, replace it with an underscore to prevent "unused variable" compile errors (e.g., for _, v := range arr).
4. Map Iteration: Returns (Key, Value). If you only write one variable, it returns only the Key.
5. String Iteration: Strings in Go are UTF-8 bytes. 'range' elegantly jumps character-by-character (runes), returning the starting byte index and the rune (int32) value. Use string(c) to print the actual letter.
========================================
*/	