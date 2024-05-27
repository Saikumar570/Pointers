// array of pointers
/*
package main

import "fmt"

func main() {
	arr := []int{10, 20, 30}
	var i int
	var ptrArr [3]*int

	for i = 0; i < 3; i++ {
		ptrArr[i] = &arr[i]
	}

	for i = 0; i < 3; i++ {
		fmt.Printf("*ptrArr[%d] = %d\n", i, *ptrArr[i])
	}
}


// double pointer

package main

import "fmt"

func main() {
	var X int = 100
    var ptrA *int = &X

    var ptrB **int = &ptrA

	fmt.Printf("Value of variable X: %d\n", X)
    fmt.Printf("Address of variable X: %d\n", &X)
    fmt.Printf("Value of pointer ptrA: %d\n", ptrA)
	fmt.Printf("Address of pointer ptrA: %d\n", &ptrA)
    fmt.Printf("Value of pointer ptrB: %d\n", ptrB)

	fmt.Printf("Value of X from ptrA : %d\n", *ptrA)
	fmt.Printf("Value of X from ptrB : %d\n", **ptrB)
}


// pointer to function

package main
import "fmt"

func myfunction(pvar *int){
     *pvar = *pvar + 10
}

func main() {
	var x int = 25
	fmt.Println("Before function call, x =  ", x)

    var intPtr *int = &x

    myfunction(intPtr)
	fmt.Println("After function call, x =  ", x)

    myfunction(&x)
    fmt.Println("After passing Address, x = ", x)
}




// Pointer to struct

package main

import "fmt"

type Employee struct {
	name string
	empid int
}

func main() {

	emp := Employee{"Raju", 19078}

	pts := &emp

	fmt.Println(pts)

	fmt.Println(pts.name)

	//dereferencing
	fmt.Println((*pts).name)

}
*/

// pointers usage in the interfaces

package main

import (
	"fmt"
)

type age interface {
	show(int)
}

func show(a age, newage int) {
	a.show(newage)
}

type ageofone struct {
	oldage int
}

func (aoo ageofone) show(newage int) {
	aoo.oldage = newage
}

type ageofsecond struct {
	oldage int
}

func (aos *ageofsecond) show(newage int) {
	aos.oldage = newage
}
func main() {
	first := ageofone{oldage: 21}
	second := ageofsecond{oldage: 30}
	//after four years updating the ages of two persons
	//using value
	show(first, 25)
	//using reference
	show(&second, 34)
	fmt.Println("The age of first one is ", first.oldage)//The age wont be updated as we are passing by the value
	fmt.Println("The age of second one is ", second.oldage)//The age will be updated to 34 as we are using reference
}
