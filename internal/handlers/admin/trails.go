package admin

import (
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/domain"
	"github.com/nathanhollows/pest-quest/internal/flash"
	"github.com/nathanhollows/pest-quest/internal/handlers"
	"gorm.io/gorm/clause"
)

// TrailsIndex allows markers to be added or removed from the map.
func TrailsIndex(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "trails"

	var markers []domain.Location
	env.DB.Preload(clause.Associations).Find(&markers)
	data["markers"] = markers

	return handlers.RenderAdmin(w, data, "admin/trails/index.html")
}
