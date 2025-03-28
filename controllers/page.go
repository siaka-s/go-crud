package controllers

import (
	"html/template"
	"net/http"
)

func Accueil(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/accueil.html"))
	tmpl.Execute(w, nil)
}

func AddUser(w http.ResponseWriter, r *http.Request) {

}

func EditUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
