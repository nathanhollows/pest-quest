// Argon is a QR code based education platform
package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/sessions"
	"github.com/nathanhollows/pest-quest/internal/config"
	"github.com/nathanhollows/pest-quest/internal/domain"
	"github.com/nathanhollows/pest-quest/internal/filesystem"
	"github.com/nathanhollows/pest-quest/internal/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var router *chi.Mux
var env handlers.Env

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Compress(5))

	var store sessions.Store = sessions.NewCookieStore([]byte(config.Cfg.Server.SessionKey))

	db, err := gorm.Open(sqlite.Open(config.Cfg.Database.File), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	env = handlers.Env{
		Session: store,
		DB:      *db,
		Data:    make(map[string]interface{}),
	}
}

func main() {
	env.DB.AutoMigrate(
	// 	&models.Poll{},
	// 	&models.Admin{},
	// 	&models.ScanEvent{},
	// 	&models.User{},
	// 	&models.Page{},
	// 	&models.Trail{},
	// 	&models.Gallery{},
		&domain.Blog{},
		&domain.Marker{},
		&domain.MarkerType{},
	)
	routes()
	fmt.Println(http.ListenAndServe(":"+config.Cfg.Server.Port, router))
}

// Set up the routes needed for the game.
func routes() {
	router.Handle("/", handlers.HandlePublic{Env: &env, H: handlers.Index})
	router.Handle("/blog", handlers.HandlePublic{Env: &env, H: handlers.Blog})
	router.Handle("/about", handlers.HandlePublic{Env: &env, H: handlers.About})
	router.Handle("/leaderboard", handlers.HandlePublic{Env: &env, H: handlers.Leaderboard})
	router.Handle("/privacy-and-terms", handlers.HandlePublic{Env: &env, H: handlers.Privacy})

	router.Handle("/login", handlers.HandlePublic{Env: &env, H: handlers.Privacy})
	router.Handle("/logout", handlers.HandlePublic{Env: &env, H: handlers.Privacy})

	router.Handle("/404", handlers.HandlePublic{Env: &env, H: handlers.Error404})
	router.NotFound(handlers.NotFound)

	workDir, _ := os.Getwd()
	filesDir := filesystem.Myfs{Dir: http.Dir(filepath.Join(workDir, "web/static"))}
	filesystem.FileServer(router, "/public", filesDir)
}
