package main

import (
	"fmt"
	"time"
)

// 1. Defining a basic Struct
type customer struct {
	name  string
	phone string
}

// 2. Composition: Embedding 'customer' inside 'order'
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // Holds timestamp
	customer            // Embedded struct
}

// 3. Constructor Function Pattern (Returns a pointer to the struct)
// This is the standard Go way to initialize complex structs.
func newOrder(id string, amount float32, cust customer) *order {
	myOrder := order{
		id:        id,
		amount:    amount,
		status:    "Received", // Default status
		createdAt: time.Now(),
		customer:  cust,
	}
	return &myOrder // Returning the memory address!
}

// 4. Pointer Receiver Method (Can MODIFY the struct)
func (o *order) changeStatus(newStatus string) {
	o.status = newStatus
}

// 5. Value Receiver Method (Can only READ, creates a copy)
func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// --- A. Manual Struct Initialization & Composition ---
	fmt.Println("--- Manual Initialization ---")
	manualOrder := order{
		id:     "101",
		amount: 50.0,
		status: "Pending",
		customer: customer{
			name:  "Ashish",
			phone: "9876543210",
		},
	}
	// Modifying embedded struct field directly
	manualOrder.customer.name = "Ashish Verma"
	fmt.Printf("Manual Order ID: %s, Customer: %s\n", manualOrder.id, manualOrder.customer.name)

	// --- B. Using Constructor & Receiver Methods ---
	fmt.Println("\n--- Using Constructors & Methods ---")
	newCust := customer{name: "John Doe", phone: "1122334455"}
	
	// Creating order using our constructor (returns a pointer)
	myOrder := newOrder("102", 99.99, newCust)
	
	fmt.Printf("Order Status Before: %s\n", myOrder.status)
	
	// Calling Pointer Receiver to modify data
	myOrder.changeStatus("Delivered")
	fmt.Printf("Order Status After: %s\n", myOrder.status)
	
	// Calling Value Receiver to read data
	fmt.Printf("Order Amount: $%.2f\n", myOrder.getAmount())

	// --- C. Anonymous Structs ---
	fmt.Println("\n--- Anonymous Struct ---")
	// Useful for quick, one-time data structures (like parsing simple JSON)
	apiResponse := struct {
		statusCode int
		message    string
	}{
		statusCode: 200,
		message:    "Success",
	}
	fmt.Printf("API Response: %d - %s\n", apiResponse.statusCode, apiResponse.message)
}

/*
========================================
REVISION NOTES FOR STRUCTS:
1. Theory: Structs are the Go equivalent of Classes. They group related data together.
2. Composition: Go has NO inheritance (no 'extends'). We embed structs inside other structs to share properties (Composition).
3. Constructors: Go doesn't have native constructors. We write functions returning pointers (func newOrder() *order) to mimic them.
4. Receiver Methods: Functions attached to a struct. 
   - Use POINTER receivers func (s *StructName) to modify the data.
   - Use VALUE receivers func (s StructName) if you only need to read the data.
5. Anonymous Structs: Structs declared and instantiated immediately without a global name. Great for one-off uses.
========================================
*/