package handlers

import (
	"net/http"

	"github.com/nt2311-vn/goth/views/foo"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, foo.Index())
}
