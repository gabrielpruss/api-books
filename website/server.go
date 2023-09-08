package website

import (
	"html/template"
	"log"
	"net/http"
)

func Run() {
	tmpl, err := template.ParseFiles("website/templates/index.html")
	if err != nil {
		panic(err)
	}

	print(tmpl)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//tmpl.Execute(w, nil)
		http.ServeFile(w, r, "website/templates/index.html")
	})

	http.HandleFunc("/css/style.css", func(w http.ResponseWriter, r *http.Request) {
		// Seta o cabeçalho `Content-Type` para `text/css`.
		w.Header().Set("Content-Type", "text/css")

		// Serve o arquivo `style.css`.
		http.ServeFile(w, r, "css/style.css")
	})

	log.Fatal(http.ListenAndServe(":5050", nil))
}
