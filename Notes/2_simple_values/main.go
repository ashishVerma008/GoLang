package main

import "fmt"

func main() {
	// 1. Integers (int): Whole numbers without decimals.
	// We can perform math operations directly.
	fmt.Println("1 + 1 =", 1+1)

	// 2. Strings (string): Text wrapped in double quotes.
	fmt.Println("hello golang")

	// 3. Booleans (bool): Represents true or false logic.
	// Note: They must be lowercase.
	fmt.Println(true)
	fmt.Println(false)

	// 4. Floats (float64): Numbers with decimal points.
	fmt.Println(10.5)
	
	// Float division vs Integer division example
	fmt.Println("Float Division (7.0/3.0) :", 7.0/3.0) 
	fmt.Println("Integer Division (7/3) :", 7/3) // Drops the decimal part
}

/*
========================================
REVISION NOTES:
1. 'int' = whole numbers (1, -5, 100).
2. 'string' = text in double quotes ("hello"). Single quotes ('') are for characters.
3. 'bool' = true or false (always lowercase).
4. 'float64' = decimal numbers (10.5).
5. Math Gotcha: Integer division (7/3) gives an integer (2). You need floats (7.0/3.0) to get exact decimal answers.
6. Use '//' for single-line comments.
========================================
*/