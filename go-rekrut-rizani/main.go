package main

import (
	"go-rekrut-rizani/config"
	"go-rekrut-rizani/controllers/berandacontroller"
	"go-rekrut-rizani/controllers/userscontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	//1.beranda
	http.HandleFunc("/", berandacontroller.Welcome)
	//2. user
	http.HandleFunc("/users", userscontroller.Index)
	http.HandleFunc("/users/add", userscontroller.Add)
	http.HandleFunc("/users/edit", userscontroller.Edit)
	http.HandleFunc("/users/delete", userscontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
