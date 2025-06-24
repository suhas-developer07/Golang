package main

import (
	"fmt"
	"myapp/calculator"
)

func main(){
	var a int  //This variable will not access to other package until passthrough function 
	var b int 
	var operator string

	//taking input from console
	fmt.Printf("Enter two number to calculate\n");
	fmt.Scanf("%d\n %d\n",&a,&b)
    
	//here i utilize printf and scanf to write on console

	fmt.Println("Enter operartor to perform")
	fmt.Scan(&operator)

  
    result :=	calculator.Calculate(a,b,operator)

	fmt.Printf("Answer: %d\n", result)
}
