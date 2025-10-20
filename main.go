package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/karthikbhandary2/to-do/server/routes"
)

func main() {
	r := routes.Router()
	fmt.Println("starting the server on port 9000...")
	log.Fatal(http.ListenAndServe(":9000", r))	
}