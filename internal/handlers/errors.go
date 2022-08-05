package handlers

import (
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/flash"
)

// NotFound is the 404 handlers.
func NotFound(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = "Error 404"
	data["section"] = "error"
	render(w, data, "errors/404.html")
}

// Error404 is a directly accessible 404 page (via /404)
func Error404(env *Env, w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["title"] = "Error 404"
	data["section"] = "error"
	return render(w, data, "errors/404.html")
}

// AccessDenied is the page visitors see when they try to access they shouldn't
func AccessDenied(env *Env, w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["title"] = "Access Denied"
	data["section"] = "error"
	data["messages"] = flash.Get(w, r)
	return render(w, data, "errors/accessdenied.html")
}
