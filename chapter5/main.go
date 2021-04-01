package main

import "fmt"

func main(){
	x:= [5]int{1,4,3,4,5,}
	for _, value:= range x{
		fmt.Println(value)	
	}
	fmt.Println(x[0])
}