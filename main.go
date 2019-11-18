package main

import (
	"github.com/sharathnasa/webservice/controller"
	"net/http"
)

func main() {

	//this is how we call and assign values to structCode
	//u := models.User{
	//	ID:        2,
	//	FirstName: "Sharath",
	//	LastName:  "Kumar",
	//}
	//fmt.Print(u)

	// to make use of controllers i.e. router we need to register the controller
	controller.RegisterController()
	http.ListenAndServe(":3000", nil)
}
