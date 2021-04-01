package main

import "fmt"
import "strings"

func main(){
	str:= "This is, me"
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.Contains(str,"t"))
	fmt.Println(strings.Contains(str,"T"))
	fmt.Println(strings.Replace(str,"i","I",1))
}

