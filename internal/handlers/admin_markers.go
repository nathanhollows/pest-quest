package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nathanhollows/pest-quest/internal/domain"
	"github.com/nathanhollows/pest-quest/internal/flash"
	"github.com/nathanhollows/pest-quest/internal/helpers"
	"gorm.io/gorm/clause"
)

// Markers allows markers to be added or removed from the map.
func MarkersIndex(env *Env, w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "markers"

	var markers []domain.Marker
	env.DB.Preload(clause.Associations).Find(&markers)
	data["markers"] = markers

	return renderAdmin(w, data, "admin/markers/index.html")
}

// MarkersAdd shows the form for adding a new marker.
func MarkersCreate(env *Env, w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "markers"

	types := []domain.MarkerType{}
	env.DB.Find(&types)
	data["types"] = types

	if r.Method == http.MethodPost {
		r.ParseForm()
		marker := domain.Marker{}
		marker.Parse(r.PostForm, &env.DB)
		res := env.DB.Model(&marker).Save(&marker)
		if res.Error != nil {
			flash.Set(w, r, flash.Message{Title: "Could not save marker", Message: res.Error.Error(), Style: "warning"})
		}
	}

	return renderAdmin(w, data, "admin/markers/create.html")
}

// Markers allows markers to be added or removed from the map.
func MarkersEdit(env *Env, w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["section"] = "markers"

	marker := domain.Marker{}
	res := env.DB.Where("id = ?", chi.URLParam(r, "id")).Find(&marker)
	if res.RowsAffected == 0 || res.Error != nil {
		flash.Set(w, r, flash.Message{Message: "The marker could not be found", Title: "Something went wrong"})
		http.Redirect(w, r, helpers.URL("/admin/markers"), http.StatusNotFound)
		return nil
	}
	data["marker"] = marker

	types := []domain.MarkerType{}
	env.DB.Find(&types)
	data["types"] = types

	// If the form has been submitted
	if r.Method == http.MethodPost {
		r.ParseForm()
		marker.Parse(r.PostForm, &env.DB)
		res := env.DB.Model(&marker).Save(&marker)
		if res.Error != nil || res.RowsAffected == 0 {
			flash.Set(w, r, flash.Message{Title: "Could not save marker", Message: res.Error.Error(), Style: "warning"})
		}
	}

	return renderAdmin(w, data, "admin/markers/edit.html")
}

// MarkersUpdate saves the new values of a marker
func MarkersUpdate(env *Env, w http.ResponseWriter, r *http.Request) error {
	// If the form has been submitted
	if r.Method == http.MethodPost {
		r.ParseForm()

		marker := domain.Marker{}
		res := env.DB.Where("id = ?", r.Form.Get("id")).Find(&marker)
		if res.RowsAffected == 0 || res.Error != nil {
			flash.Set(w, r, flash.Message{Message: "The marker could not be found", Title: "Something went wrong"})
			fmt.Println("Could not find")
			return Error404(env, w, r)
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
	return Error404(env, w, r)
}

// MarkersDelete removes a marker from the map.
func MarkersDelete(env *Env, w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "markers"

	return renderAdmin(w, data, "admin/markers/index.html")
}
