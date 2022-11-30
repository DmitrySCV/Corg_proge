package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello world"))
}
func addCorg(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Создать новую картинку"))
}

func showCorgs(w http.ResponseWriter, r *http.Request) {
	// MAIN SECTION HTML CODE
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
	fmt.Fprintf(w, "<title>Go</title>")
	fmt.Fprintf(w, "<img src='assets/corg.jpg' alt='corg' style='width:235px;height:320px;'>")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	// ABOUT SECTION HTML CODE
	fmt.Fprintf(w, "<title>Go/about/</title>")
	fmt.Fprintf(w, "Expert web design by JT Skrivanek")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/corg", showCorgs)
	mux.HandleFunc("/corg/create", addCorg)
	http.HandleFunc("/about/", about_handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
