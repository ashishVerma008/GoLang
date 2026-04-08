package main

import "fmt"

// 1. Pass by Value (Creates a copy)
func changeNumByValue(num int) {
	num = 5
	fmt.Println("Inside changeNumByValue:", num) // This prints 5
}

// 2. Pass by Reference (Uses memory address)
// The parameter type is '*int' (Pointer to an integer)
func changeNumByReference(num *int) {
	// 3. Dereferencing: Use '*' to get/set the actual value at the memory address
	*num = 5
	fmt.Println("Inside changeNumByReference:", *num)
}

func main() {
	// Original variable
	score := 1
	fmt.Println("Initial Score:", score)

	// --- PASS BY VALUE ---
	// changeNumByValue(score)
	// fmt.Println("After Pass by Value, Score is still:", score) // Output: 1

	// --- PASS BY REFERENCE ---
	// 4. Address-of Operator: Use '&' to pass the memory address of the variable
	fmt.Println("\nMemory address of score is:", &score)
	
	changeNumByReference(&score)

	// Since the memory address was passed, the original variable is modified
	fmt.Println("After Pass by Reference, final Score is:", score) // Output: 5
}

/*
========================================
REVISION NOTES FOR POINTERS:
1. Theory: Pointers hold the memory address of another variable. They avoid copying large amounts of data and allow direct modification of original variables.
2. Operators: 
   - '&' (Address-of): Gets the memory location (e.g., &num).
   - '*' (Dereference): Accesses or modifies the value at that memory location (e.g., *num = 10).
3. Type: '*int' means a pointer that points to an integer.
4. Go vs C++: Go has NO pointer arithmetic. You cannot do ptr++ to jump to the next memory block. This makes Go highly memory-safe.
5. Danger: Never dereference a 'nil' pointer (*ptr = 5 when ptr is nil), it will cause a runtime Panic!
========================================
*/