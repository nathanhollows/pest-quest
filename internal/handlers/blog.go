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
	env.Data["messages"] = flash.Get(w, r)
	env.Data["section"] = "blog"

	env.Data["posts"] = blog.GetLatest(env.DB, 10)

	return render(w, env.Data, "blog/index.html")
}

// BlogPost display a single post.
func BlogPost(env *Env, w http.ResponseWriter, r *http.Request) error {
	env.Data["messages"] = flash.Get(w, r)
	env.Data["section"] = "blog"

	post := domain.Blog{}
	res := env.DB.Find(&post, "url = ?", chi.URLParam(r, "title"))
	env.Data["post"] = post

	if res.RowsAffected == 0 {
		return StatusError{http.StatusNotFound, errors.New("could not find the blog post")}
	}

	return render(w, env.Data, "blog/post.html")
}
