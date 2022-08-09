package handlers

import (
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/core/agentname"
	"github.com/nathanhollows/pest-quest/internal/core/missions"
	"github.com/nathanhollows/pest-quest/internal/flash"
)

// Index is the homepage of the game.
// Prints a very simple page asking only for a team code.
func Index(env *Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["section"] = "index"

	env.Data["agent"] = agentname.Generate()
	env.Data["missions"] = missions.Get(&env.DB)

	return render(w, env.Data, "index/index.html")
}
