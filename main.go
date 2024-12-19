package main

import (
	"fmt"
	"forum/app/controllers"
	"log"
	"net/http"
)

func main() {
    // Register routes for your application
    http.HandleFunc("/", controllers.HomePage)

    // Start the server
	fmt.Println("run: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
