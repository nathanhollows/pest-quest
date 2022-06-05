package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nathanhollows/hedge-men-ltd/internal/flash"
)

// Index is the homepage of the game.
// Prints a very simple page asking only for a team code.
func Index(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-store")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "index"

	session, err := env.Session.Get(r, "uid")
	if err != nil || session.Values["id"] == nil {
		session, _ = env.Session.New(r, "uid")
		session.Options.HttpOnly = true
		session.Options.SameSite = http.SameSiteStrictMode
		session.Options.Secure = true
		id := uuid.New()
		session.Values["id"] = id.String()
		session.Save(r, w)
	}

	return render(w, data, "index/index.html")
}
