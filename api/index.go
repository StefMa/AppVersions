package api

import (
	"fmt"
	"net/http"

	"github.com/StefMa/AppVersions/presentation"
)

func Index(w http.ResponseWriter, r *http.Request) {
	androidQuery := r.URL.Query().Get("android")
	iosQuery := r.URL.Query().Get("ios")
	format := r.URL.Query().Get("format")

	if androidQuery != "" || iosQuery != "" {
		url := "/lookup?android=" + androidQuery + "&ios=" + iosQuery + "&format=" + format
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	} else {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprint(w, presentation.Index())
	}
}
