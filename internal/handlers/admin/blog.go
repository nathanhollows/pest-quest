package admin

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nathanhollows/pest-quest/internal/domain"
	"github.com/nathanhollows/pest-quest/internal/flash"
	"github.com/nathanhollows/pest-quest/internal/handlers"
	"github.com/nathanhollows/pest-quest/internal/helpers"
	"gorm.io/gorm/clause"
)

// BlogIndex shows all the blog posts
func BlogIndex(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["title"] = "Blog"
	env.Data["section"] = "blog"

	blogs := []domain.Blog{}
	env.DB.Preload(clause.Associations).Find(&blogs)
	env.Data["blogs"] = blogs

	return handlers.RenderAdmin(w, env.Data, "admin/blog/index.html")
}

// BlogCreate shows the form for creating new posts, and also creates them
func BlogCreate(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["title"] = "New blog post"
	env.Data["section"] = "blog"

	if r.Method == http.MethodPost {
		r.ParseForm()
		blog := domain.Blog{}
		blog.Parse(r.PostForm, &env.DB)
		res := env.DB.Model(&blog).Save(&blog)
		if res.Error != nil {
			flash.Set(w, r, flash.Message{Title: "Could not save blog", Message: res.Error.Error(), Style: "warning"})
		}
	}

	return handlers.RenderAdmin(w, env.Data, "admin/blog/create.html")
}

// BlogEdit shows the form for editing a blog post
func BlogEdit(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["section"] = "blog"

	blog := domain.Blog{}
	res := env.DB.Where("url = ?", chi.URLParam(r, "url")).Find(&blog)
	if res.RowsAffected == 0 || res.Error != nil {
		flash.Set(w, r, flash.Message{Message: "The blog could not be found", Title: "Something went wrong"})
		http.Redirect(w, r, helpers.URL("/admin/blog"), http.StatusTemporaryRedirect)
		return nil
	}
	env.Data["blog"] = blog

	// If the form has been submitted
	if r.Method == http.MethodPost {
		r.ParseForm()
		blog.Parse(r.PostForm, &env.DB)
		res := env.DB.Save(&blog)
		if res.Error != nil || res.RowsAffected == 0 {
			flash.Set(w, r, flash.Message{Title: "Could not save blog", Message: res.Error.Error(), Style: "warning"})
		} else {
			http.Redirect(w, r, helpers.URL(fmt.Sprint("admin/blog/edit/", blog.URL)), http.StatusTemporaryRedirect)
		}
		return nil
	}

	return handlers.RenderAdmin(w, env.Data, "admin/blog/edit.html")
}

// BlogUpdate saves the new values of a blog
func BlogUpdate(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
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

// BlogDelete deletes a blog.
func BlogDelete(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["section"] = "blog"

	return handlers.RenderAdmin(w, env.Data, "admin/markers/index.html")
}
