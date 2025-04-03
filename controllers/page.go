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
	tmpl := template.Must(template.ParseFiles("views/addUser.html"))
	tmpl.Execute(w, nil)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/edit.html"))
	tmpl.Execute(w, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/delete.html"))
	tmpl.Execute(w, nil)
}
