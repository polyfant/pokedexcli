package main

import (
	"fmt"
	"log"
	"net/http"
	"pokedex/web"
)
	func main() {
	handler := web.NewHandler()

	// Setup routes
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/search", handler.Search)
	http.HandleFunc("/pokemon/details", handler.PokemonDetails)

	fmt.Println("Server starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
