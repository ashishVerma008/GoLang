package main

import (
	"fmt"
	"maps" // For maps.Equal (Go 1.21+)
)

func main() {
	// 1. Creating a Map using make()
	// map[KeyType]ValueType
	m := make(map[string]string)

	// 2. Setting elements
	m["name"] = "Golang"
	m["area"] = "Backend"

	// 3. Getting elements & The "Zero Value" rule
	fmt.Println("Name:", m["name"])
	fmt.Println("Area:", m["area"])
	// "version" key will return the empty string , the compiler will not crash 
	fmt.Println("Version (missing key):", m["version"]) 

	// 4. Map with integer values
	cart := make(map[string]int)
	cart["laptop"] = 1
	cart["mouse"] = 2
	
	fmt.Println("\nCart:", cart)
	fmt.Println("Total unique items in cart (length):", len(cart))

	// 5. Deleting a key
	delete(cart, "mouse")
	fmt.Println("Cart after deleting mouse:", cart)

	// 6. Clearing the whole map (Go 1.21+)
	clear(cart)
	fmt.Println("Cart after clearing:", cart)

	// 7. Literal initialization (Declaring and adding values at once)
	inventory := map[string]int{"phones": 3, "tablets": 5}

	// 8. The "Comma OK" Idiom (Very Important in Go)
	// 'v' gets the value, 'ok' gets a boolean (true if key exists, false if not)
	v, ok := inventory["phones"]
	if ok {
		fmt.Println("\nPhones found in inventory. Count:", v)
	} else {
		fmt.Println("\nPhones key does NOT exist in inventory.")
	}

	// 9. Comparing Maps
	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 3}
	// m1 == m2 -> ILLEGAL IN GO
	isEqual := maps.Equal(m1, m2)
	fmt.Println("\nAre m1 and m2 equal?", isEqual)
}

/*
========================================
REVISION NOTES FOR MAPS (Hash/Dict):
1. Theory: Maps store Key-Value pairs. Same as std::unordered_map in C++. Time complexity for operations is $O(1)$.
2. Use make() to create a map: make(map[string]int).
3. NEVER write to an uninitialized map (var m map[string]int). It will cause a runtime panic.
4. If you access a missing key, it returns the 'Zero Value' (0, "", false). It does NOT crash.
5. "Comma OK" Idiom: val, ok := map["key"]. Use 'ok' to check if the key actually exists.
6. delete(map, key) removes a key. clear(map) empties the whole map.
7. Use maps.Equal() to compare two maps.
========================================
*/