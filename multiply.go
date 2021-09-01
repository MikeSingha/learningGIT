package main

import "fmt"

func main(){
	a:= 3
	b:= 5
	fmt.Println(multiply(a,b))

}

func multiply(a int, b int) int{
	return a * b
}