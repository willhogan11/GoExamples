package main

// Good style to use the factored import statement.
import (
	"fmt"
)

func addNumbers(num1, num2 int)(int) {
	var result int = num1 + num2
	return result
}

func main() {
	fmt.Println("This is a test")
	fmt.Println("The result is: ", addNumbers(12, 15))

	sum := 0
	for i := 0; i < 10 ; i++{
		sum += i
	}
	fmt.Println(sum)
}

