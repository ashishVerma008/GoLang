package main

import (
	"errors"
	"fmt"
)
type Output struct{
	res int
	err string
}

func divide(a int, b int ) (float64,error){
	if b==0 {
		return 0,errors.New("Cannot divide by 0")
	}
	return float64(a)/float64(b),nil
}


func main(){
data,err:=divide(2,0)

if err!=nil{
	fmt.Println(err)
}else{
	fmt.Println(data)
}
}