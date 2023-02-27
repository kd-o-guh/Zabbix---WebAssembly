package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//fsf := http.FileServer(http.Dir("favicons"))
	//http.Handle("/favicons/", http.StripPrefix("/favicons", fsf))

	//fsi := http.FileServer(http.Dir("images"))
	//http.Handle("/images/", http.StripPrefix("/images", fsi))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server is up and listening on port 8080.")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	index := template.Must(template.ParseFiles("tmpl/index.html"))

	if err := index.Execute(w, nil); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execucao do modelo: ", err.Error())
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	login := template.Must(template.ParseFiles("tmpl/login.html"))

	if err := login.Execute(w, nil); err != nil {
		http.Error(w, "Houve um erro na renderização da página.", http.StatusInternalServerError)
		fmt.Println("[Ola] Erro na execucao do modelo: ", err.Error())
	}

}
