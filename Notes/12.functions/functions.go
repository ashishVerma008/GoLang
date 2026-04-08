package main

import "fmt"

// 1. Basic Function (Shorthand parameters: a and b are both int)
// The last 'int' is the return type.
func add(a, b int) int {
	return a + b
}

// 2. Multiple Return Values
// This function returns two strings and one boolean.
func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

// 3. Higher-Order Function: Takes a function as an argument (Callback)
func processItArg(fn func(a int) int) {
	fmt.Println("Inside processItArg, calling the passed function...")
	result := fn(10) // Calling the passed function with value 10
	fmt.Println("Result from passed function:", result)
}

// 4. Higher-Order Function: Returns a function
func processItReturn() func(a int) int {
	// Returning an anonymous function
	return func(a int) int {
		return a * 4 // Multiplies the input by 4
	}
}

func main() {
	fmt.Println("--- 1. Basic Function ---")
	result := add(3, 5)
	fmt.Println("3 + 5 =", result)

	fmt.Println("\n--- 2. Multiple Returns ---")
	// Using blank identifier '_' to ignore the boolean value
	lang1, lang2, _ := getLanguages()
	fmt.Println("Languages:", lang1, "and", lang2)

	fmt.Println("\n--- 3. Anonymous Function ---")
	// Creating a function without a name and assigning it to a variable
	myFunc := func(a int) int {
		return a + 100
	}
	fmt.Println("myFunc(5) =", myFunc(5))

	fmt.Println("\n--- 4. Passing Function as Arg ---")
	// Passing our anonymous 'myFunc' variable as an argument
	processItArg(myFunc)

	fmt.Println("\n--- 5. Returning a Function ---")
	// multiplierFunc now holds the function returned by processItReturn
	multiplierFunc := processItReturn()
	fmt.Println("multiplierFunc(6) =", multiplierFunc(6)) // 6 * 4 = 24
}

/*
========================================
REVISION NOTES FOR FUNCTIONS:
1. Theory: Functions are first-class citizens in Go. They can be stored in variables, passed as args, and returned.
2. Shorthand Params: func add(a, b int) is valid.
3. Multiple Returns: Go natively supports returning multiple values: func get() (int, string) {}.
4. NO Function Overloading: You cannot have two functions with the same name in the same package.
5. Anonymous Functions: You can create nameless functions and assign them to variables (like JS arrow functions).
========================================
*/