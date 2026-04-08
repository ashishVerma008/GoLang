package main

import "fmt"

// 1. Global Constant: Can be accessed anywhere in the package
const AppName = "MyBackendServer"

func main() {
	// 2. Local Constant: Explicit type vs Inferred type
	const language string = "Golang" // Explicit
	const version = 1.22             // Inferred (Compiler knows it's a float)

	// 3. Grouped Constants: Best way to write multiple config values
	const (
		port = 5000
		host = "localhost"
		db   = "mongodb"
	)

	// Error Example (Uncommenting the line below will cause a compile error)
	// port = 8080 // "cannot assign to port"

	// 4. No Error for Unused Constants!
	const unusedConst = "I am not used, but Go won't complain!"

	fmt.Println("Starting", AppName)
	fmt.Printf("Language: %s, Version: %v\n", language, version)
	fmt.Printf("Server running on %s:%d\n", host, port)
	fmt.Println("Database:", db)
}

/*
========================================
REVISION NOTES FOR CONSTANTS:
1. Constants ('const') cannot be changed once assigned.
2. You CANNOT use the shorthand operator (:=) with constants.
3. Group constants using const ( ... ) for clean configuration blocks.
4. BIG DIFFERENCE: Unused variables cause errors, but unused constants DO NOT.
5. Constant values must be known at compile time (no dynamic function results).
========================================
*/