package main

import "fmt"
import "reflect"

func main(){
	var number = 10
	var reflectNumber = reflect.ValueOf(number)
	
	if reflectNumber.Kind() == reflect.Int{
		fmt.Println("Sama :",reflectNumber.Int())
	} else {
		fmt.Println("Beda")
	}
}
