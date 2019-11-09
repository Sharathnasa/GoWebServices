package pointer

import (
	"fmt"
)

func yoy() {

	// This is one way of using pointer
	// this is how we declare the pointers in GO!
	var firstName *string = new(string)
	// here i'm de-referencing the pointer to store the value
	*firstName = "Arthur Curry"

	// While printing the result again we need to do de-referencing
	fmt.Println(*firstName)

	// second way
	lastName := "Meera"
	fmt.Println(lastName)

	ptr := &lastName
	fmt.Println(ptr, *ptr)
}
