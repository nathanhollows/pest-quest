package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nathanhollows/pest-quest/internal/flash"
)

// Privacy shows the Privacy and Terms required to use the website.
func Privacy(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "blog"

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

	return render(w, data, "privacy/index.html")
}
