package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZuhybDev/waqt/internal/routes"
)

func main() {
	routers := routes.Register()
	fmt.Println("Server is running at PORT: 8080")
	log.Fatal(http.ListenAndServe(":8080", routers))

}
