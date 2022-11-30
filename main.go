package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
func showCorgs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображать картинку"))
}
func addCorg(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Создать новую картинку"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/corg", showCorgs)
	mux.HandleFunc("/corg/create", addCorg)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
