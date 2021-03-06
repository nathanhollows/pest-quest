package handlers

import (
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/core/agentname"
	"github.com/nathanhollows/pest-quest/internal/core/markers"
	"github.com/nathanhollows/pest-quest/internal/flash"
)

// Index is the homepage of the game.
// Prints a very simple page asking only for a team code.
func Index(env *Env, w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "index"

	data["agent"] = agentname.Generate()
	data["markers"] = markers.GetPublished(env.DB)

	return render(w, data, "index/index.html")
}
