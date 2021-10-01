package main

import (
	"fmt"
)

func diffAndsum(a int, b int) (int, int){
	var sum = (a + b) 
	var diff = (a - b)
	return sum, diff
}
 func main(){

	//fmt.Println()
	a:= 7
	b:= 5

	sum, diff:= diffAndsum(a,b)
	
	fmt.Println(sum)
	fmt.Println(diff)
 }