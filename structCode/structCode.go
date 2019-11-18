package structCode

import (
	"fmt"
)

func main() {
	// this is how we define struct
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	// lets look at adding the data
	var u user
	u.ID = 1
	u.FirstName = "Sharath"
	u.LastName = "Nasa"
	fmt.Println(u)

	//another way of adding value to struct is using maps
	u2 := user{
		ID:        1,
		FirstName: "Christiano",
		LastName:  "Ronaldo",
	}
	fmt.Println(u2)
}
