package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Hexlet"))
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hexlet is the leading educational platform for Software Engineers"))
	})

	http.HandleFunc("/courses", func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")
		var pageCourses string

		switch page {
		case "":
			pageCourses = "Введите номер курса!"
		case "1":
			pageCourses = "Как написать свой первый \"Hello world\" на go..."
		case "2":
			pageCourses = "Как работает сборщик мусора в go..."
		default:
			pageCourses = "Курс в разработке..."
		}

		// возвращаем страницу курса
		w.Write([]byte(pageCourses))
	})

	http.ListenAndServe(":80", nil)
}
