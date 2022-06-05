package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"github.com/nathanhollows/hedge-men-ltd/internal/config"
	"github.com/nathanhollows/hedge-men-ltd/internal/helpers"
	"gorm.io/gorm"
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// HandlePublic takes both a game manager and http function.
type HandlePublic struct {
	Env *Env
	H   func(e *Env, w http.ResponseWriter, r *http.Request) error
}

// HandleAdmin takes both a game manager and http function.
type HandleAdmin struct {
	Env *Env
	H   func(e *Env, w http.ResponseWriter, r *http.Request) error
}

// Env is the shared game manager for each request.
type Env struct {
	Session sessions.Store
	DB      gorm.DB
	Data    map[string]interface{}
}

func (h HandlePublic) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

func (h HandleAdmin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, err := h.Env.Session.Get(r, "admin")
	if err != nil || session.Values["id"] == nil {
		http.Redirect(w, r, helpers.URL("login"), http.StatusFound)
	}

	err = h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

var funcs = template.FuncMap{
	"uppercase": func(v string) string {
		return strings.ToUpper(v)
	},
	"divide": func(a, b int) float32 {
		if a == 0 || b == 0 {
			return 0
		}
		return float32(a) / float32(b)
	},
	"progress": func(a, b int) float32 {
		if a == 0 || b == 0 {
			return 0
		}
		return float32(a) / float32(b) * 100
	},
	"add": func(a, b int) int {
		return a + b
	},
	"sub": func(a, b int) int {
		return a - b
	},
	"url": func(s ...string) template.URL {
		return template.URL(helpers.URL(s...))
	},
	"currentYear": func() int {
		return time.Now().UTC().Year()
	},
	"stylesheetversion": func() string {
		file, err := os.Stat("web/static/css/style.css")
		if err != nil {
			fmt.Println(err)
		}
		modifiedtime := file.ModTime().Nanosecond()
		return fmt.Sprint(modifiedtime)
	},
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
}

func parse(patterns ...string) *template.Template {
	patterns = append(patterns, "layout.html", "flash.html")
	for i := 0; i < len(patterns); i++ {
		patterns[i] = "web/templates/" + patterns[i]
	}
	return template.Must(template.New("base").Funcs(funcs).ParseFiles(patterns...))
}

func render(w http.ResponseWriter, data map[string]interface{}, patterns ...string) error {
	w.Header().Set("Content-Type", "text/html")
	if data["siteTitle"] == nil {
		data["siteTitle"] = config.Cfg.Frontend.SiteName
	}
	err := parse(patterns...).ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return err
}
