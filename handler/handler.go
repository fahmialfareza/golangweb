package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Wellcome to Home"))

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)

		http.Error(w, "Error is happened, keep calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title": "I'm learning Golang Web",
	// 	"content": "I'm learning Golang Web with Fahmi Alfareza",
	// }

	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 3}

	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 11},
		{ID: 2, Name: "Xpander", Price: 270000000, Stock: 8},
		{ID: 3, Name: "Pajero Sport", Price: 500000000, Stock: 1},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)

		http.Error(w, "Error is happened, keep calm", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World, saya sedang belajar Golang web"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from Nintendo"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprintf(w, "Product page: %d", idNumb)

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)

		http.Error(w, "Error is happened, keep calm", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content": idNumb,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)

		http.Error(w, "Error is happened, keep calm", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method // GET POST

	switch method {
		case "GET": 
			w.Write([]byte("Ini adalah GET"))
		case "POST":
			w.Write([]byte("Ini adalah POST"))
		default:
			http.Error(w, "Error is happened, keep calm", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)

			http.Error(w, "Error is happened, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)

			http.Error(w, "Error is happened, keep calm", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error is happened, keep calm", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)

			http.Error(w, "Error is happened, keep calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{} {
			"name": name,
			"message": message,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)

			http.Error(w, "Error is happened, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)

			http.Error(w, "Error is happened, keep calm", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error is happened, keep calm", http.StatusBadRequest)
}