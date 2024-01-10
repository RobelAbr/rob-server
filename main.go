package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// commit something
func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was a error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was a error executing the template.", http.StatusInternalServerError)
		return
	}
}

// commit something
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

// commit something
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

// commit something
func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<b><h1>FAQ Page</h1><p>Welcome for any Questions:<br /></b>
<ul>
	<li>Q: Is there a free version?<br />A: Yes! We offer a free trial for 30 days on any paid plans.</li><br />
	<li>Q: What are your support hour?<br />A: We have support staff answering emails 24/7, though response time may be a bit slower on weekdays.</li><br />
	<li>Q: How do i contact support?<br />A: Email us - <a href="mailto:"support@rabraha.io">support@rabraha.io</a></li><br />
</ul>
`)
}

// starts server and all pages of the webserver that be called
func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Ohhh sorry....Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
