package main

import "fmt"

func main() {
	// 1. Basic If-Else
	age := 16
	fmt.Println("--- Basic If-Else ---")
	if age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	// 2. If, Else-If, Else Ladder
	age = 14
	fmt.Println("\n--- Else-If Ladder ---")
	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is a teenager")
	} else {
		fmt.Println("Person is a kid")
	}

	// 3. Using Logical Operators (&& = AND, || = OR)
	role := "admin"
	hasPermissions := true
	fmt.Println("\n--- Logical Operators ---")
	// Both conditions must be true
	if role == "admin" && hasPermissions {
		fmt.Println("Access Granted: Welcome Admin!")
	} else {
		fmt.Println("Access Denied")
	}

	// 4. Inline Variable Declaration (Super important in Go!)
	fmt.Println("\n--- Inline Declaration ---")
	// 'userAge' is created here and will be destroyed after this if-else block ends
	if userAge := 20; userAge >= 18 {
		fmt.Println("User is an adult, age is:", userAge)
	} else {
		fmt.Println("User is a minor, age is:", userAge)
	}
	// Note: You CANNOT use 'userAge' here outside the block. It will give an error.

	// 5. NO TERNARY OPERATOR in Go.
	// You cannot do: result := (age > 18) ? "yes" : "no"
	// You MUST use the standard if-else structure.
}

/*
========================================
REVISION NOTES FOR IF-ELSE:
1. No parentheses () required around conditions.
2. The '{' must be on the same line as 'if'.
3. The '} else {' must be strictly on the same line.
4. Go does NOT have the ternary operator (? :).
5. Inline declaration (if x := 10; x > 5) limits the variable's scope to that block only, keeping memory clean.
========================================
*/