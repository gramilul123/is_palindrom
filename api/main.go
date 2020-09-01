package main

import (
	"isPalindrom/template"
	"isPalindrom/word"
	"log"
	"net/http"
	"os"
)

func main() {

	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
		log.Printf("Default port: %s \n", port)
	}
	log.Printf("Server is listening on %s port \n", port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var err error
	var words []word.Word
	var templateData template.TemplateData

	if r.Method == "POST" {

		err = r.ParseForm()
		if err == nil {
			wordExemplar := word.Word{
				Word:        r.FormValue("word"),
				IsPalindrom: word.IsPalindrom(r.FormValue("word"))}
			wordExemplar.Save()
		}
	}

	words, err = word.GetWordsList()

	if err == nil {
		templateData, err = template.GetTemplate(words)

		if err == nil {
			templateData.Template.Execute(w, templateData)
		}
	}

	if err != nil {
		log.Fatalln(err)
	}
}
