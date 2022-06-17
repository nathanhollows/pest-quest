package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/nathanhollows/pest-quest/internal/domain"
	"github.com/nathanhollows/pest-quest/internal/flash"
	"github.com/nathanhollows/pest-quest/internal/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login handles user logins
func Login(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")
	data := make(map[string]interface{})
	data["section"] = "session"

	r.ParseForm()

	if r.Method == http.MethodPost {
		username := r.Form.Get("username")

		var user domain.Admin
		env.DB.Model(&user).Where("username = ?", username).Find(&user)

		if checkHashPassword(user.Password, r.Form.Get("password")) {
			session, err := env.Session.Get(r, "admin")
			if err != nil || session.Values["id"] == nil {
				session, _ = env.Session.New(r, "admin")
				session.Options.HttpOnly = true
				session.Options.SameSite = http.SameSiteStrictMode
				session.Options.Secure = true
				id := uuid.New()
				session.Values["id"] = id.String()
				session.Save(r, w)
				http.Redirect(w, r, helpers.URL("admin"), http.StatusFound)
				return nil
			}
		} else {
			flash.Set(w, r, flash.Message{Message: "Incorrect username or password"})
		}
	}

	data["messages"] = flash.Get(w, r)
	return render(w, data, "session/login.html")
}

// Logout destroys the user session
func Logout(env *Env, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/html")

	session, err := env.Session.Get(r, "admin")
	if err == nil {
		if session.Values["id"] != nil {
			session.Options.MaxAge = -1
			session.Save(r, w)
			flash.Set(w, r, flash.Message{Message: "You have been logged out"})
		}
	}
	http.Redirect(w, r, fmt.Sprint(helpers.URL("login")), http.StatusFound)
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkHashPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createUser(username string, password string, db *gorm.DB) error {
	user := domain.Admin{}
	user.Username = username
	pw, err := hashPassword(password)
	if err != nil {
		return err
	}
	user.Password = pw
	result := db.Model(&domain.Admin{}).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
