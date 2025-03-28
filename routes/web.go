package routes

import (
	"net/http"
	"siaka-s/go-crud/controllers"
)

func Web() {

	http.HandleFunc("/", controllers.Accueil)
	http.HandleFunc("/add", controllers.AddUser)
	http.HandleFunc("/edit", controllers.EditUser)
	http.HandleFunc("/delete", controllers.DeleteUser)

}
