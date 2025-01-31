package userscontroller

import (
	"go-rekrut-rizani/models/usersmodel"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	pengguna := usersmodel.GetAll()
	data := map[string]any{
		"users": pengguna,
	}

	temp, err := template.ParseFiles("views/users/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
