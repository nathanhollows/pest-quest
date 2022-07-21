package handlers

import (
	"net/http"
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
