package session

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/nathanhollows/pest-quest/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func GetUser(request *http.Request) (user domain.User) {
	if cookie, err := request.Cookie("user"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("user", cookie.Value, &cookieValue); err == nil {
			user.DisplayName = cookieValue["name"]
			id, _ := strconv.ParseUint(cookieValue["id"], 10, 0)
			user.ID = uint(id)
			user.Admin, _ = strconv.ParseBool(cookieValue["admin"])
		}
	}
	return user
}

func SetSession(user domain.User, response http.ResponseWriter) {
	value := map[string]string{
		"name":  user.DisplayName,
		"admin": fmt.Sprint(user.Admin),
		"id":    fmt.Sprint(user.ID),
	}
	if encoded, err := cookieHandler.Encode("user", value); err == nil {
		cookie := &http.Cookie{
			Name:    "user",
			Value:   encoded,
			Path:    "/",
			Expires: time.Now().Add(time.Hour * 24 * 365),
		}
		http.SetCookie(response, cookie)
	}
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "user",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHashPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateUser(user *domain.User, db *gorm.DB) error {
	var err error
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		return err
	}
	result := db.Model(&domain.User{}).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
