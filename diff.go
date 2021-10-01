<<<<<<< HEAD
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
=======

package main

import "fmt"

func addition(a int, b int) int{
	return a - b
}

func main(){
	fmt.Println(addition(5,2))

}
>>>>>>> ae241788df57e335e4e249f93750bd001501c28c
