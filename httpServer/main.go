package main

//! METHOD
// GET
// POST
// PUT
// DELETE
//! STATUS
// 200
// 201
// 204
// 400
// 401
// 404
// 405
// 500
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
//!----------------------------------------------------------

func information(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Azat"
	}
	fmt.Fprintf(w, "Method:    %s\n", r.Method)
	fmt.Fprintf(w, "Path:      %s\n", r.URL.Path)
	fmt.Fprintf(w, "Name:      %s\n", name)
	fmt.Fprintf(w, "Web-site:  %s\n", r.Header.Get("User-agent"))
	fmt.Fprintf(w, "Language:  %s\n", r.Header.Get("Accept-Language"))
	fmt.Fprintf(w, "Host: 	   %s\n", r.Host)
}

//! ---------------------------------------------------------------------------


func helloUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello user")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "The ID must be", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "user's #%d\n", id)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	// WriteHeader status kodini belgilaydi
	// w serverning javobini yozish uchun ishlatiladi
	fmt.Fprintf(w, "new user is created")
}

var users = map[int]string{
	1: "Ali",
	2: "Vali",
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, " this ID doesn't exist", http.StatusBadRequest)
		return
	}
	if _, ok := users[id]; !ok {
		http.Error(w, "this ID doesn't exist", http.StatusNotFound)
		return
	}

	delete(users, id)
	w.WriteHeader(http.StatusNoContent) // 204
}

//! -------------------------------------------------------------------------------------



func main() {
	// http.HandleFunc("/", hello)
	// http.HandleFunc("/dashboard", information)


	mux := http.NewServeMux()
	mux.HandleFunc("GET /helloUser", helloUser)
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("DELETE /delete/user/{id}", deleteUser)	

	log.Fatal(http.ListenAndServe(":8080", mux))
}