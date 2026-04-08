package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Simple Switch
	fmt.Println("--- Simple Switch ---")
	i := 3
	switch i {
	case 1:
		fmt.Println("Value is One")
	case 2:
		fmt.Println("Value is Two")
	case 3:
		fmt.Println("Value is Three") // Ye print hoga aur automatically break ho jayega
	default:
		fmt.Println("Other Value")
	}

	// 2. Multiple Condition Switch (Grouping cases)
	fmt.Println("\n--- Multiple Condition Switch ---")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend! Time to relax.")
	default:
		fmt.Println("It's a workday. Back to code.")
	}

	// 3. Type Switch (Finding out the data type at runtime)
	fmt.Println("\n--- Type Switch ---")
	// interface{} means this function can accept ANY data type
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("It's an integer!")
		case string:
			fmt.Println("It's a string!")
		case bool:
			fmt.Println("It's a boolean!")
		default:
			fmt.Println("Unknown type")
		}
	}
	whoAmI(55)        // Output: integer
	whoAmI("Golang")  // Output: string
	whoAmI(true)      // Output: boolean

	// 4. Switch without an expression (Acts like an if-else ladder)
	fmt.Println("\n--- Switch Without Expression ---")
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}

/*
========================================
REVISION NOTES FOR SWITCH:
1. In Go, 'break' is AUTOMATIC at the end of each case. No need to write it.
2. If you want it to jump to the next case (like C++), use the 'fallthrough' keyword.
3. You can group multiple values in one case using commas (case 1, 2, 3:).
4. switch i.(type) is used to find the data type of an interface{} at runtime.
5. 'switch {}' without a condition acts exactly like an if-else if-else chain.
========================================
*/