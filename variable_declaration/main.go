package main

import "fmt"


//declaring the v 
var brand string
var price int = 50000
var sold bool 

func main(){
	var student1 string = "suahs"  // this is the normal type of declaration of the variable in go 
	var student2 = "developer"  // without using type of the varible
	x:= 2   //special type of declaration compiler automatically identify the type of varible using assigned value

	var company string   //declaring an variable 
	company = "zerodha" // assigning an value to it

	//assigning the value for declared out of the function variales
	brand = "Nike"   
	sold = false

	//multiple varible declartaion
	var a,b,c,d int = 1,3,5,7

	
    fmt.Printf("brand:%s\n price:%d\n sold:%t\n ", brand, price , sold)

}