package handlers

import (
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/flash"
)

// About gives some background information about the game (in fiction).
func About(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "about"

	_, err := env.Session.Get(r, "uid")
	if err != nil {
		_, err = startSession(env.Session, r, w)
	}

	return render(w, data, "about/index.html")
}
