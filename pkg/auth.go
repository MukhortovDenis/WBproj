package pkg

import (
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Зарегаться")
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Зайти")
}
