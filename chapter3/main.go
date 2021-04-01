package main

import "fmt"


func main(){
fmt.Println("Enter temperature in Fahrenhite: ")
var temp float64
fmt.Scanf("%f", &temp)	
temp = ((temp-32)*5/9)
fmt.Println("Temperature in Celsius is: ",temp)
}


