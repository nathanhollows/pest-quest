package handlers

import (
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/flash"
)

// Privacy shows the Privacy and Terms required to use the website.
func Privacy(env *Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["section"] = "blog"

	return render(w, env.Data, "privacy/index.html")
}
