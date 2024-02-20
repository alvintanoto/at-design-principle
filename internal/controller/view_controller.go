package controller

import (
	"net/http"

	vpage "alvintanoto.id/design-principle/internal/view/page"
)

type ViewController interface {
	HomepageHandler() func(http.ResponseWriter, *http.Request)
}

type implViewController struct {
}

func NewViewController() ViewController {
	return &implViewController{}
}

func (i *implViewController) HomepageHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vpage.Homepage().Render(r.Context(), w)
	}
}
