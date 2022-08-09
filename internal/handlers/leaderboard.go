package handlers

import (
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/flash"
)

// Leaderboard shows players who is currently in the lead.
// The leaderboard is with respect to the first game and the ARG.
// Alternatively, this could be a list of completed quests.
func Leaderboard(env *Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["section"] = "leaderboard"

	return render(w, env.Data, "leaderboard/index.html")
}
