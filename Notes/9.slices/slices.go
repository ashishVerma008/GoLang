package main

import (
	"fmt"
	"slices" // Required for comparing slices (Go 1.21+)
)

func main() {
	// 1. Uninitialized Slice (Nil Slice)
	var nilSlice []int
	fmt.Println("Is nilSlice nil?", nilSlice == nil) // Output: true

	// 2. Initialized but Empty Slice
	emptySlice := []int{}
	fmt.Println("Is emptySlice nil?", emptySlice == nil) // Output: false

	// 3. Using make() -> make([]type, length, capacity)
	// Recommended way to create slices if you know the approx size
	nums := make([]int, 0, 5)
	fmt.Printf("\nBefore Append -> Length: %d, Capacity: %d\n", len(nums), cap(nums))

	// 4. Appending data (adds to the end and updates length)
	nums = append(nums, 1)
	nums = append(nums, 2, 3, 4) // Can append multiple at once
	fmt.Println("After Append:", nums)
	fmt.Printf("After Append -> Length: %d, Capacity: %d\n", len(nums), cap(nums))

	// 5. Copying Slices
	nums2 := make([]int, len(nums)) // Create a new slice with the same length
	copy(nums2, nums)               // Deep copy: destination, source
	fmt.Println("\nCopied Slice (nums2):", nums2)

	// 6. Slice Operator [start:end] (start is inclusive, end is exclusive)
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("\nOriginal numbers:", numbers)
	fmt.Println("numbers[0:2] ->", numbers[0:2]) // Gets index 0, 1
	fmt.Println("numbers[:3]  ->", numbers[:3])  // Gets from start up to index 2
	fmt.Println("numbers[2:]  ->", numbers[2:])  // Gets from index 2 up to the end

	// 7. Comparing Slices using slices.Equal
	sliceA := []int{1, 2, 3}
	sliceB := []int{1, 2, 3}
	// sliceA == sliceB -> THIS IS ILLEGAL IN GO
	isEqual := slices.Equal(sliceA, sliceB)
	fmt.Println("\nAre sliceA and sliceB equal?", isEqual)

	// 8. 2D Slices (Dynamic Rows and Columns)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("\n2D Slice:", matrix)
}

/*
========================================
REVISION NOTES FOR SLICES:
1. Slices are dynamic (vectors). They grow and shrink automatically.
2. Uninitialized slice is 'nil'. An empty slice []int{} is NOT 'nil'.
3. make([]int, length, capacity) is the best way to pre-allocate memory for performance.
4. Use append() to add elements. Do not use indexes (nums[5] = 10) if the length is smaller than 5.
5. Slices point to a hidden array. To create an independent copy, use the copy() function.
6. [start:end] slice operator: 'start' is included, 'end' is excluded.
7. You cannot use '==' to compare slices. Use slices.Equal().
========================================
*/