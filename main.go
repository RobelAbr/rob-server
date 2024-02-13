package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// executeTemplate is a utility function to render an HTML template.
// It takes an http.ResponseWriter and the filepath of the template as input.
func executeTemplate(w http.ResponseWriter, filepath string) {
	// Set the Content-Type header to indicate that the response is HTML.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Parse the HTML template located at the specified filepath.
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		// If there is an error parsing the template, log the error,
		// send an internal server error response, and return from the function.
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was a error parsing the template.", http.StatusInternalServerError)
		return
	}

	// Execute the parsed template, writing the output to the http.ResponseWriter.
	// Pass nil as the data since we are not providing any specific data to the template.
	err = tpl.Execute(w, nil)
	if err != nil {
		// If there is an error executing the template, log the error,
		// send an internal server error response, and return from the function.
		log.Printf("executing template: %v", err)
		http.Error(w, "There was a error executing the template.", http.StatusInternalServerError)
		return
	}
}

// here we call the template home
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

// here we call the template contact
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

// here we call the template career
func CareerHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "career.gohtml")
	executeTemplate(w, tplPath)
}

// commit something
func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

// here we call the template career
func loginHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "login.gohtml")
	executeTemplate(w, tplPath)
}

// here we call the template career
func registerHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "register.gohtml")
	executeTemplate(w, tplPath)
}

// starts server and all pages of the webserver that be called
func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/career", CareerHandler)
	r.Get("/faq", faqHandler)
	r.Get("/login", loginHandler)
	r.Get("/register", registerHandler)
	r.Handle("/rob-server/web/img/{name}", http.StripPrefix("/rob-server/web/img/", http.FileServer(http.Dir("./web/img/"))))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Ohhh sorry....Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
