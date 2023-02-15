package middleware

import (
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(200)
		w.Write([]byte("Таблица успешного обновлена."))
		return
	}
}
