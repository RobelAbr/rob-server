package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:rabraha@hotmail.io\">rabraha@hotmail.io</a>.")
}

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
func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "<h1>Page not found<h1>", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
