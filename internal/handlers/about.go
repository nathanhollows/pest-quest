package handlers

import (
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/flash"
)

// About gives some background information about the game (in fiction).
func About(env *Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["section"] = "about"

	return render(w, env.Data, "about/index.html")
}
