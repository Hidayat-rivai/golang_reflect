package main

import "fmt"
import "reflect"

func main(){
	var number = 10
	var reflectNumber = reflect.ValueOf(number)
	
	fmt.Println("Tipe :",reflectNumber)
}
