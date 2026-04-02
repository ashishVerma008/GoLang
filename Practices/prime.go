package main

import "fmt"

func checkPrime(a int)bool{
	for i:=2;i<a;i++{
		if a%i==0 {
			return false
		}
	}
	return true
}
func Prime(nums []int)[]int{
	res:=[]int{}
	for _,val:=range nums {
		if checkPrime(val) {
			res=append(res, val)
		}
	}
	return res
}

func main() {
	// nums:=make([]int,2,5)
	
//    for i:=0;i<len(nums);i++ {
// 	nums[i]=i+1
//    }
//    for i:=0;i<5;i++{
// 	nums=append(nums, i+1)
//    }
// 	fmt.Println(nums)
// 	fmt.Printf("The new capacity of the nums aray is %d",cap(nums))
	// fmt.Printf("Hi this costs around %d ",20)  


	nums:=[]int{2,3,4,5,100,300,400,9000}
	fmt.Println(Prime(nums))
}