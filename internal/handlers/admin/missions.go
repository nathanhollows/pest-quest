package admin

import (
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/domain"
	"github.com/nathanhollows/pest-quest/internal/flash"
	"github.com/nathanhollows/pest-quest/internal/handlers"
	"gorm.io/gorm/clause"
)

// MissionsIndex shows the quests available, past, present, and future
func MissionsIndex(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "missions"

	var missions []domain.Mission
	env.DB.Preload(clause.Associations).Find(&missions)
	data["missions"] = missions

	return handlers.RenderAdmin(w, data, "admin/missions/index.html")
}
