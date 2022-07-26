package handlers

import (
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/core/markers"
	"github.com/nathanhollows/pest-quest/internal/flash"
)

// Map shows locations of interest for the public and for players
func Map(env *Env, w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "map"

	data["markers"] = markers.GetPublished(env.DB)

	return render(w, data, "map/index.html")
}
