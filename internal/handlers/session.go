package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/nathanhollows/pest-quest/internal/core/agentname"
	"github.com/nathanhollows/pest-quest/internal/core/session"
	"github.com/nathanhollows/pest-quest/internal/domain"
	"github.com/nathanhollows/pest-quest/internal/flash"
	"github.com/nathanhollows/pest-quest/internal/helpers"
)

// Login handles user logins
func Login(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["section"] = "session"

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		pass := r.FormValue("password")

		user := domain.User{}
		user.FindByEmail(&env.DB, email)

		if user.ID > 0 && session.CheckHashPassword(user.Password, pass) {
			session.SetSession(user, w)
			flash.Set(w, r, flash.Message{Message: "Welcome back, " + user.DisplayName})
			http.Redirect(w, r, helpers.URL(), http.StatusFound)
		} else {
			flash.Set(w, r, flash.Message{Message: "Please double check your email and password"})
		}
	}

	data["messages"] = flash.Get(w, r)
	return render(w, data, "session/login.html")
}

// Logout destroys the user session
func Logout(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	session.ClearSession(w)
	flash.Set(w, r, flash.Message{Message: "You have been logged out"})
	http.Redirect(w, r, fmt.Sprint(helpers.URL("login")), http.StatusFound)
	return nil
}

// Register creates a new user
func Register(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["section"] = "session"
	data["username"] = agentname.Generate()

	r.ParseForm()

	if r.Method == http.MethodPost && r.Form.Get("terms") != "" {
		user := domain.User{}
		user.DisplayName = r.Form.Get("agent")
		user.Email = r.Form.Get("email")
		user.Password = r.Form.Get("password") // This gets hashed before saving!

		var err = session.CreateUser(&user, &env.DB)
		if err == nil {
			session, err := env.Session.Get(r, "user")
			if err != nil || session.Values["id"] == nil {
				session, _ = env.Session.New(r, "user")
				session.Options.HttpOnly = true
				session.Options.SameSite = http.SameSiteStrictMode
				// session.Options.Secure = true
				id := uuid.New()
				session.Values["id"] = id.String()
				session.Save(r, w)
				http.Redirect(w, r, helpers.URL(""), http.StatusFound)
				flash.Set(w, r, flash.Message{Message: fmt.Sprint("Welcome Agent ", user.DisplayName)})
				return nil
			}
		} else {
			flash.Set(w, r, flash.Message{Message: err.Error()})
		}
	} else if r.Method == http.MethodPost && r.Form.Get("terms") == "" {
		flash.Set(w, r, flash.Message{Message: "Please accept the terms and conditions."})
	}

	data["messages"] = flash.Get(w, r)
	return render(w, data, "session/register.html")
}

// Suggestname suggests a new name
func SuggestName(env *Env, w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["name"] = agentname.Generate()
	return renderBlank(w, data, "session/name.html")
}
