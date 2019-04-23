package interfaces

import (
	"github.com/krzysztofSkolimowski/url-shrinker/pkg/app"
	"net/http"
)

func SaveHandler(svc app.Shrinker) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//	todo
		w.Write([]byte("todo save"))
	})
}

func KeyHandler(svc app.Retriever) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//	todo
		w.Write([]byte("todo key"))
	})
}

func DecodeHandler(svc app.Retriever) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//	todo
		w.Write([]byte("todo - decode"))
	})
}

func RedirectHandler(svc app.Retriever) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// todo
		w.Write([]byte("todo 0 redirect"))
		http.Redirect(w, r, string("https://krzysztof-skolimowski.netlify.com/"), 301)
	})
}
