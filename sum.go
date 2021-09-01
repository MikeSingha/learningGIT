
package main

import "fmt"

func main(){
	a:=2
	b:=3
 addition(a,b)
}

func addition(a int, b int){
	
	c := a + b
	fmt.Println("The sum is:", c)
}

/* or

func addition(a int, b int) int{
	return a + b
}

func main(){
	fmt.Println(addition(4,6))

}
*/