package web

import (
	"fmt"
	"net/http"
	"teleport/internal/database"
)

func ShortUrlHandler(db database.Service, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	url := r.FormValue("url")
	hash := db.SetLongUrl(url)

	component := ShortUrl(fmt.Sprintf("https://tp.airtxt.me/v1/%v", hash))
	component.Render(r.Context(), w)
}
