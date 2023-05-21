// `main.go`
package main

import (
	"github.com/galymzhantolepbergen/Booking-Store-Golang/app/controllers"
	
)

func main() {

	println("~ Welcome!")
	println("~ http://localhost:8000/welcome/")
	controllers.HandlerRequest()

}
