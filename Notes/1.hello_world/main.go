package main

// Importing the 'fmt' package for formatted input/output operations
import "fmt"

// The main function is the starting point of the application
func main() {
	// 1. Println: Prints text and automatically moves to a new line
	fmt.Println("hello world")

	// 2. Print: Prints text on the same line (no automatic new line)
	fmt.Print("This is on the ")
	fmt.Print("same line.\n") // Added \n manually to break the line

	// 3. Printf: Prints formatted text using verbs like %s (string), %d (digit), %v (value)
	channelName := "Coders "
	fmt.Printf("Learning Go from %s!\n", channelName)
}

/*
========================================
REVISION NOTES:
- 'package main' is strictly required for the file to be executable.
- 'fmt' works like <iostream> for input/output.
- The opening brace '{' MUST be on the same line as the function declaration.
- Println = Adds new line automatically.
- Print = No new line.
- Printf = Formatted string (use %v for general values, \n for a manual new line).
========================================
*/