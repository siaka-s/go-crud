package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Définir les routes
	http.HandleFunc("/", accueilHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/contact", contactHandler)

	// Démarrer le serveur
	fmt.Println("Serveur démarré sur http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erreur du serveur: %s\n", err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenue sur la page de contact!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bonjour le monde!")
}

func accueilHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenue sur la page d'accueil!")
}
