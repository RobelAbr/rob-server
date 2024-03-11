package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// executeTemplate ist eine Hilfsfunktion zum Rendern eines HTML-Templates.
// Sie nimmt einen http.ResponseWriter und den Dateipfad des Templates als Eingabe.
func executeTemplate(w http.ResponseWriter, filepath string) {
	// Setze den Content-Type-Header, um anzuzeigen, dass die Antwort HTML ist.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Parsen des HTML-Templates am angegebenen Dateipfad.
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		// Bei einem Fehler beim Parsen des Templates: Fehler protokollieren,
		// eine interne Serverfehlerantwort senden und die Funktion beenden.
		log.Printf("Template-Parsing-Fehler: %v", err)
		http.Error(w, "Beim Parsen des Templates ist ein Fehler aufgetreten.", http.StatusInternalServerError)
		return
	}

	// Ausführen des geparsen Templates und Schreiben des Ergebnisses in den http.ResponseWriter.
	// Nil wird als Daten übergeben, da wir keine spezifischen Daten an das Template liefern.
	err = tpl.Execute(w, nil)
	if err != nil {
		// Bei einem Fehler beim Ausführen des Templates: Fehler protokollieren,
		// eine interne Serverfehlerantwort senden und die Funktion beenden.
		log.Printf("Template-Ausführungsfehler: %v", err)
		http.Error(w, "Beim Ausführen des Templates ist ein Fehler aufgetreten.", http.StatusInternalServerError)
		return
	}
}

// homeHandler ruft das Template "home" auf
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

// contactHandler ruft das Template "contact" auf
func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

// careerHandler ruft das Template "career" auf
func careerHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "career.gohtml")
	executeTemplate(w, tplPath)
}

// faqHandler ruft das Template "faq" auf
func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

// loginHandler ruft das Template "login" auf
func loginHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "login.gohtml")
	executeTemplate(w, tplPath)
}

// registerHandler ruft das Template "register" auf
func registerHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "register.gohtml")
	executeTemplate(w, tplPath)
}

// tutorialsHandler ruft das Template "tutorials" auf
func tutorialsHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "tutorials.gohtml")
	executeTemplate(w, tplPath)
}

// developHandler ruft das Template "develop" auf
func developHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "develop.gohtml")
	executeTemplate(w, tplPath)
}

// learnHandler ruft das Template "learn" auf
func learnHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "learn.gohtml")
	executeTemplate(w, tplPath)
}

// materialHandler ruft das Template "material" auf
func materialHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "material.gohtml")
	executeTemplate(w, tplPath)
}

// main startet den Server und ruft alle Seiten des Webservers auf, die aufgerufen werden können
func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/career", careerHandler)
	r.Get("/faq", faqHandler)
	r.Get("/login", loginHandler)
	r.Get("/register", registerHandler)
	r.Get("/tutorials", tutorialsHandler)
	r.Get("/develop", developHandler)
	r.Get("/learn", learnHandler)
	r.Get("/material", materialHandler)
	r.Handle("/rob-server/web/img/{name}", http.StripPrefix("/rob-server/web/img/", http.FileServer(http.Dir("./web/img/"))))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Ohhh sorry....Page not found", http.StatusNotFound)
	})
	fmt.Println("Server wird auf :3000 gestartet...")
	http.ListenAndServe(":3000", r)
}
