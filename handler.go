package urlshort

import (
	"net/http"
)

func MapHandler(inToRedirected map[string]string, fallback http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		val, ok := inToRedirected[r.URL.String()]
		if ok {
			http.Redirect(w, r, val, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	})
}
