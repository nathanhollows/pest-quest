package admin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nathanhollows/pest-quest/internal/domain"
	"github.com/nathanhollows/pest-quest/internal/flash"
	"github.com/nathanhollows/pest-quest/internal/handlers"
	"github.com/nathanhollows/pest-quest/internal/helpers"
	"gorm.io/gorm/clause"
)

// AdviceIndex lists the advice available
func AdviceIndex(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["title"] = "Sage Advice"
	env.Data["section"] = "advice"

	advice := []domain.Advice{}
	env.DB.Preload(clause.Associations).Find(&advice)
	env.Data["advice"] = advice

	return handlers.RenderAdmin(w, env.Data, "admin/advice/index.html")
}

// AdviceCreates shows the form for adding new advice, and also handles saving
func AdviceCreate(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["title"] = "New Sage Advice"
	env.Data["section"] = "advice"

	log.Println("Hit AdviceCreate")
	if r.Method == http.MethodPost {
		log.Println("Hit AdviceCreatePost")

		r.ParseForm()
		advice := domain.Advice{}
		advice.Parse(r.PostForm, &env.DB)
		res := env.DB.Model(&advice).Save(&advice)
		if res.Error != nil {
			log.Println("Could not save")

			flash.Set(w, r, flash.Message{Title: "Could not save advice", Message: res.Error.Error(), Style: "warning"})
		}
		log.Println("Saved")

	}

	return handlers.RenderAdmin(w, env.Data, "admin/advice/create.html")
}

// AdviceEdit shows the form for editing advice
func AdviceEdit(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["section"] = "advice"

	advice := domain.Advice{}
	res := env.DB.Where("id = ?", chi.URLParam(r, "id")).Find(&advice)
	if res.RowsAffected == 0 || res.Error != nil {
		flash.Set(w, r, flash.Message{Message: "The Sage Advice could not be found", Title: "Something went wrong"})
		http.Redirect(w, r, helpers.URL("/admin/advice"), http.StatusTemporaryRedirect)
		return nil
	}
	env.Data["advice"] = advice

	// If the form has been submitted
	if r.Method == http.MethodPost {
		r.ParseForm()
		advice.Parse(r.PostForm, &env.DB)
		res := env.DB.Save(&advice)
		if res.Error != nil || res.RowsAffected == 0 {
			flash.Set(w, r, flash.Message{Title: "Could not save blog", Message: res.Error.Error(), Style: "warning"})
		} else {
			http.Redirect(w, r, helpers.URL(fmt.Sprint("admin/advice/edit/", advice.ID)), http.StatusTemporaryRedirect)
		}
		return nil
	}

	return handlers.RenderAdmin(w, env.Data, "admin/advice/edit.html")
}

// AdviceUpdate saves the new values of Advice
func AdviceUpdate(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	// If the form has been submitted
	if r.Method == http.MethodPost {
		r.ParseForm()

		marker := domain.Location{}
		res := env.DB.Where("id = ?", r.Form.Get("id")).Find(&marker)
		if res.RowsAffected == 0 || res.Error != nil {
			flash.Set(w, r, flash.Message{Message: "The marker could not be found", Title: "Something went wrong"})
			fmt.Println("Could not find")
			return handlers.Error404(env, w, r)
		}

		marker.Parse(r.PostForm, &env.DB)
		res = env.DB.Model(&marker).Save(&marker)
		if res.Error != nil || res.RowsAffected == 0 {
			flash.Set(w, r, flash.Message{Title: "Could not save marker", Message: res.Error.Error(), Style: "warning"})
			http.Redirect(w, r, helpers.URL("/admin/markers"), http.StatusSeeOther)
			return nil
		}
		url := fmt.Sprint("/admin/markers/edit/", marker.ID)
		http.Redirect(w, r, helpers.URL(url), http.StatusSeeOther)
		return nil
	}
	return handlers.Error404(env, w, r)
}

// AdviceDelete deletes a piece of Advice
func AdviceDelete(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["section"] = "advice"

	return handlers.RenderAdmin(w, env.Data, "admin/markers/index.html")
}
