package web

import (
	"fmt"
	"net/http"
	"teleport/internal/database"
)

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	url := r.FormValue("url")
	hash := database.New().SetLongUrl(url)

	component := ShortUrl(fmt.Sprintf("http://localhost:3032/v1/%v", hash))
	component.Render(r.Context(), w)
}
