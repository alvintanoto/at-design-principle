package controller

import (
	"net/http"

	"alvintanoto.id/design-principle/internal/dto"
	vpage "alvintanoto.id/design-principle/internal/view/page"
)

type ViewController interface {
	HomepageHandler() func(http.ResponseWriter, *http.Request)
	ColorpageHandler() func(http.ResponseWriter, *http.Request)
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

func (i *implViewController) ColorpageHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vpage.Colorpage(dto.ViewBaseDTO{
			Name: "Color",
		}).Render(r.Context(), w)
	}
}
