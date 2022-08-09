package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/nathanhollows/pest-quest/internal/config"
	"github.com/nathanhollows/pest-quest/internal/core/session"
	"github.com/nathanhollows/pest-quest/internal/flash"
	"github.com/nathanhollows/pest-quest/internal/helpers"
	"gitlab.com/golang-commonmark/markdown"
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
	h.Env.Data["session"] = session.GetUser(r)
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
	user := session.GetUser(r)
	if !user.Admin {
		flash.Set(w, r, flash.Message{Message: "You don't have access to that"})
		http.Redirect(w, r, helpers.URL("denied"), http.StatusFound)
	} else {
		h.Env.Data["session"] = session.GetUser(r)
	}

	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
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
	// Compares the activeSection to the checking section.
	// Prints "active" if the sections match
	"active": func(activeSection string, checking string) string {
		if activeSection == checking {
			return "active"
		}
		return ""
	},
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
	"svg": func(icon string, inline bool) template.HTML {
		if inline {
			path := fmt.Sprint("web/static/img/icons/", icon, ".svg")
			file, err := os.Open(path)
			if err != nil {
				return "ERROR: Icon does not exist"
			}
			b, _ := ioutil.ReadAll(file)
			return template.HTML(b)
		} else {
			url := helpers.URL(fmt.Sprint("/public/img/icons/", icon, ".svg"))
			return template.HTML(url)
		}
	},
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
}

func render(w http.ResponseWriter, data map[string]interface{}, patterns ...string) error {
	w.Header().Set("Content-Type", "text/html")
	if data["siteTitle"] == nil {
		data["siteTitle"] = config.Cfg.Frontend.SiteName
	}
	if data["section"] == nil {
		data["section"] = "error"
		log.Println("Section is not set. \t Patterns: ", patterns)
	}
	err := parse(patterns...).ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return err
}

func parse(patterns ...string) *template.Template {
	for i := 0; i < len(patterns); i++ {
		patterns[i] = "web/templates/public/" + patterns[i]
	}
	patterns = append(patterns, "web/templates/public/layout.html", "web/templates/partials/flash.html")
	return template.Must(template.New("base").Funcs(funcs).ParseFiles(patterns...))
}

func parseMD(page string, db *gorm.DB) template.HTML {
	md := markdown.New(
		markdown.XHTMLOutput(true),
		markdown.HTML(true),
		markdown.Breaks(true))

	page = regexp.MustCompile("==(.*)==").ReplaceAllString(page, "<mark>$1</mark>")
	page = regexp.MustCompile(":::([^:::]*):::").ReplaceAllString(page, `<article>
$1</article>`)

	return template.HTML(md.RenderToString([]byte(page)))
}

// TODO: Abstract this out so it is automatic
func RenderAdmin(w http.ResponseWriter, data map[string]interface{}, patterns ...string) error {
	w.Header().Set("Content-Type", "text/html")
	if data["siteTitle"] == nil {
		data["siteTitle"] = config.Cfg.Frontend.SiteName
	}
	err := parseAdmin(patterns...).ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return err
}

func parseAdmin(patterns ...string) *template.Template {
	for i := 0; i < len(patterns); i++ {
		patterns[i] = "web/templates/" + patterns[i]
	}
	patterns = append(patterns, "web/templates/admin/layout.html", "web/templates/partials/flash.html")
	return template.Must(template.New("base").Funcs(funcs).ParseFiles(patterns...))
}

func startSession(s sessions.Store, r *http.Request, w http.ResponseWriter) (*sessions.Session, error) {
	session, err := s.Get(r, "uid")
	if err != nil || session.Values["id"] == nil {
		session, err = s.New(r, "uid")
		session.Options.HttpOnly = true
		session.Options.SameSite = http.SameSiteStrictMode
		session.Options.Secure = true
		id := uuid.New()
		session.Values["id"] = id.String()
		session.Save(r, w)
	}
	return session, err
}

func renderBlank(w http.ResponseWriter, data map[string]interface{}, patterns ...string) error {
	w.Header().Set("Content-Type", "text/html")
	err := parseBlank(patterns...).ExecuteTemplate(w, "body", data)
	if err != nil {
		http.Error(w, err.Error(), 0)
		log.Print("Template executing error: ", err)
	}
	return err
}

func parseBlank(patterns ...string) *template.Template {
	for i := 0; i < len(patterns); i++ {
		patterns[i] = "web/templates/public/" + patterns[i]
	}
	return template.Must(template.New("body").Funcs(funcs).ParseFiles(patterns...))
}
