package web

import (
	"fmt"
	"net/http"
	"teleport/internal/database"

	"github.com/a-h/templ"
)

func ShortUrlHandler(db database.Service, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	url := r.FormValue("url")
	alias := r.FormValue("alias")
	hash := db.SetLongUrl(url, alias)

	var component templ.Component
	if len(hash) > 0 {
		component = ShortUrl(fmt.Sprintf("http://localhost:3033/v1/%v", hash))
	} else {
		component = Error("This alias is already in use!")
	}
	component.Render(r.Context(), w)
}
