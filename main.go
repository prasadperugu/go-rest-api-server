package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name  string
	Email string
	Role  string
}

var users = []User{
	{
		Name:  "Prasad",
		Email: "prasad.perugu@onymos.com",
		Role:  "admin",
	},
	{
		Name:  "Nandu",
		Email: "nandu.ahmed@onymos.com",
		Role:  "clinician",
	},
}

func helloWorld(x http.ResponseWriter, y *http.Request) {
	fmt.Println("prasad")

	fmt.Fprintf(x, "Hello World\n") //server side output

}

func main() {
	http.HandleFunc("/xyz", helloWorld)
	fmt.Println("Server started and listening 90903", users)
	//start the server
	//log.Fatal=> if this doesn't start correctly we need to add log
	log.Fatal(http.ListenAndServe(":9003", nil))
}
