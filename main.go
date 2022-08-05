// Argon is a QR code based education platform
package main

import (
	"fmt"
	"log"
	"net"
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
	"github.com/nathanhollows/pest-quest/internal/handlers/admin"
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
		// 	&models.Gallery{},
		&domain.Advice{},
		&domain.User{},
		&domain.Trail{},
		&domain.Page{},
		&domain.Blog{},
		&domain.Location{},
		&domain.MarkerType{},
		&domain.User{},
	)
	routes()
	fmt.Printf("%s%s%s", "http://", GetOutboundIP(), ":8803")
	fmt.Println(http.ListenAndServe(":"+config.Cfg.Server.Port, router))
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// Set up the routes needed for the game.
func routes() {
	router.Handle("/", handlers.HandlePublic{Env: &env, H: handlers.Index})
	router.Handle("/map", handlers.HandlePublic{Env: &env, H: handlers.Map})
	router.Handle("/blog", handlers.HandlePublic{Env: &env, H: handlers.Blog})
	router.Handle("/blog/{title}", handlers.HandlePublic{Env: &env, H: handlers.BlogPost})
	router.Handle("/about", handlers.HandlePublic{Env: &env, H: handlers.About})
	router.Handle("/leaderboard", handlers.HandlePublic{Env: &env, H: handlers.Leaderboard})
	router.Handle("/privacy-and-terms", handlers.HandlePublic{Env: &env, H: handlers.Privacy})
	router.Handle("/name/suggest", handlers.HandlePublic{Env: &env, H: handlers.SuggestName})

	router.Handle("/admin/markers", handlers.HandleAdmin{Env: &env, H: admin.MarkersIndex})
	router.Handle("/admin/markers/add", handlers.HandleAdmin{Env: &env, H: admin.MarkersCreate})
	router.Handle("/admin/markers/edit/{id}", handlers.HandleAdmin{Env: &env, H: admin.MarkersEdit})
	router.Handle("/admin/markers/update", handlers.HandleAdmin{Env: &env, H: admin.MarkersUpdate})
	router.Handle("/admin/markers/delete", handlers.HandleAdmin{Env: &env, H: admin.MarkersDelete})

	router.Handle("/admin/trails", handlers.HandleAdmin{Env: &env, H: admin.TrailsIndex})

	router.Handle("/admin/blog", handlers.HandleAdmin{Env: &env, H: admin.BlogIndex})
	router.Handle("/admin/blog/create", handlers.HandleAdmin{Env: &env, H: admin.BlogCreate})
	router.Handle("/admin/blog/edit/{url}", handlers.HandleAdmin{Env: &env, H: admin.BlogEdit})
	router.Handle("/admin/blog/delete", handlers.HandleAdmin{Env: &env, H: admin.BlogDelete})

	router.Handle("/admin/advice", handlers.HandleAdmin{Env: &env, H: admin.AdviceIndex})
	router.Handle("/admin/advice/create", handlers.HandleAdmin{Env: &env, H: admin.AdviceCreate})
	router.Handle("/admin/advice/edit/{id}", handlers.HandleAdmin{Env: &env, H: admin.AdviceEdit})
	router.Handle("/admin/advice/delete", handlers.HandleAdmin{Env: &env, H: admin.AdviceDelete})

	router.Handle("/admin/md/preview", handlers.HandleAdmin{Env: &env, H: admin.PreviewMD})

	router.Handle("/login", handlers.HandlePublic{Env: &env, H: handlers.Login})
	router.Handle("/logout", handlers.HandlePublic{Env: &env, H: handlers.Logout})
	router.Handle("/register", handlers.HandlePublic{Env: &env, H: handlers.Register})

	router.Handle("/404", handlers.HandlePublic{Env: &env, H: handlers.Error404})
	router.NotFound(handlers.NotFound)

	workDir, _ := os.Getwd()
	filesDir := filesystem.Myfs{Dir: http.Dir(filepath.Join(workDir, "web/static"))}
	filesystem.FileServer(router, "/public", filesDir)
}
