package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nathanhollows/pest-quest/internal/core/blog"
	"github.com/nathanhollows/pest-quest/internal/domain"
	"github.com/nathanhollows/pest-quest/internal/flash"
)

// Blog displays the most recent posts.
func Blog(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "blog"

	data["posts"] = blog.GetLatest(env.DB, 10)

	_, err := env.Session.Get(r, "uid")
	if err != nil {
		_, err = startSession(env.Session, r, w)
	}

	return render(w, data, "blog/index.html")
}

// BlogPost display a single post.
func BlogPost(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["messages"] = flash.Get(w, r)
	data["section"] = "blog"

	post := domain.Blog{}
	res := env.DB.Find(&post, "url = ?", chi.URLParam(r, "title"))
	data["post"] = post

	if res.RowsAffected == 0 {
		return StatusError{http.StatusNotFound, errors.New("could not find the blog post")}
	}

	_, err := env.Session.Get(r, "uid")
	if err != nil {
		_, err = startSession(env.Session, r, w)
	}

	return render(w, data, "blog/post.html")
}
