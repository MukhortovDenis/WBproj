package pkg

import (
	"fmt"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "не обоссан")
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Обоссан")
}
