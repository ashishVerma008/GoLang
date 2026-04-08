package main

import "fmt"

// Global variables (Outside Function). Here shorthand (:=) not works.
var globalScore int = 100

func main() {
	// 1. Explicit Declaration (Full syntax)
	var language string = "Golang"
	
	// 2. Type Inference (Compiler understand by itself what is the type)
	var age = 21 
	
	// 3. Shorthand Syntax (Most used method)
	framework := "Fiber"
	isEasy := true
	version := 1.22 // Automatically becomes float64 

	// 4. Multiple Variable Declaration 
	var x, y, z int = 1, 2, 3
	name, role := "Ashish", "Backend Developer" 

	// 5. Zero Values (Declare but not assign value)
	var emptyString string // default is ""
	var defaultInt int     // default is 0
	var defaultBool bool   // default is false

	// Printing all variables , otherwise go will not compile
	fmt.Println("Language:", language)
	fmt.Println("Age:", age)
	fmt.Println("Framework:", framework, "Is Easy?", isEasy, "Version:", version)
	fmt.Println("Coordinates:", x, y, z)
	fmt.Println("Profile:", name, "-", role)
	
	fmt.Println("--- Zero Values ---")
	fmt.Printf("String: '%s', Int: %d, Bool: %v\n", emptyString, defaultInt, defaultBool)
	fmt.Println("Global Score:", globalScore)
}

/*
========================================
REVISION NOTES FOR VARIABLES:
1. 4 main ways to declare:
   - var name string = "Go" (Explicit)
   - var name = "Go"        (Inferred)
   - name := "Go"           (Shorthand - Most used, local scope only)
   - var name string        (Zero value concept)
   
2. IMPORTANT GOTCHA: You MUST use every declared variable, otherwise Go will throw a compile-time error.
3. Zero Values: If not initialized, int = 0, float = 0.0, string = "", bool = false.
4. Shorthand (:=) operator cannot be used outside functions (e.g., global variables).
========================================
*/