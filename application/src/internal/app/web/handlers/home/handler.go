package home

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aygumov-g/shortlinks-go/application/src/internal/app/web/templates"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/form?type=CR", http.StatusMovedPermanently)
		} else {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, strings.ReplaceAll(
				templates.OTHER_HTML_PAGE, "{BACKEND_SERVER_URL}", os.Getenv("BACKEND_SERVER_URL"),
			))
		}
	}
}
