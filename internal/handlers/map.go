package handlers

import (
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/core/markers"
	"github.com/nathanhollows/pest-quest/internal/flash"
)

// Map shows locations of interest for the public and for players
func Map(env *Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["section"] = "map"

	env.Data["markers"] = markers.GetPublished(env.DB)

	return render(w, env.Data, "map/index.html")
}
