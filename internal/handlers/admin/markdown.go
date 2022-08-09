package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nathanhollows/pest-quest/internal/handlers"
)

// PreviewMD accepts MD and returns HTML
func PreviewMD(env *handlers.Env, w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {

		type markdown struct {
			Md string `json:"md"`
		}

		var response markdown
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&response)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return nil
		}

		// t := template.Must(template.New("md").Parse("{{.}}"))
		// t.Execute(w, handlers.ParseMD(response.Md, &env.DB))

		return nil
	}
	return handlers.StatusError{Code: http.StatusBadRequest, Err: errors.New("must be POST")}
}
