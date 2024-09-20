package form

import (
	"fmt"
	"net/http"

	"github.com/aygumov-g/shortlinks-go/application/src/internal/app/web/templates"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, templates.FORM_HTML_PAGE)
	}
}
