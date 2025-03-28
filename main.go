package main

import (
	"fmt"
	"net/http"
	"siaka-s/go-crud/routes"
)

func main() {
	routes.Web()

	fmt.Println("demarre le serveur sur http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Printf("Une erreur a ete detecté : %s", err)
	}

}
