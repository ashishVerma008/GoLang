package main

import "fmt"

func main() {
	var row, col int

	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&row)

	fmt.Print("Enter the number of columns: ")
	fmt.Scan(&col)

	// fmt.Println("Rows:", row)
	// fmt.Println("Columns:", col)
	for i:=1;i<=row;i++{
		for j:=1;j<=col;j++{
			if i==1 || i==row {
				fmt.Printf("*")
			}else {
				if j==1 || j==col {
					fmt.Printf("*")
				}else{
					fmt.Printf(" ")
				}
			}
			
		}
		fmt.Printf("\n")
	}
}