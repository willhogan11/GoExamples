package main

// Good style to use the factored import statement.
import (
	"fmt"
)

func addNumbers(num1, num2 int)(int) {
	var result int = num1 + num2
	return result
}

type person struct {
	name string
	age int
}


func main() {
	fmt.Println("This is a test")
	fmt.Println("The result is: ", addNumbers(12, 15))

	sum := 0
	for i := 0; i < 10 ; i++{
		sum += i
	}
	fmt.Println(sum)


	// Struct Exmaple
	var P person
	P.name, P.age = "Will", 39
	fmt.Println(P)


	// Pointers Example
	var num int = 5 //  Iniatialise num to 5
	var ptr *int // Declare an int pointer
	ptr = &num // ptr assigned the value of num

	fmt.Println("Value of memory address that pointer is pointing to : ", &ptr) // Prints the Address that's been stored, ie the address of num
	fmt.Println("Actual value at location of what pointer is pointing to (Dereferencing) : ", *ptr) // Prints the actual value stored by inital variable num, (AKA Dereferencing)
}