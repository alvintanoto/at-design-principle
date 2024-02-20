package controller

import (
	"net/http"

	"alvintanoto.id/design-principle/internal/dto"
	vpage "alvintanoto.id/design-principle/internal/view/page"
)

type ViewController interface {
	HomepageHandler() func(http.ResponseWriter, *http.Request)
	ColorpageHandler() func(http.ResponseWriter, *http.Request)
	FontpageHandler() func(http.ResponseWriter, *http.Request)
	ChartpageHandler() func(http.ResponseWriter, *http.Request)
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

func (i *implViewController) FontpageHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vpage.Fontpage(dto.ViewBaseDTO{
			Name: "Font",
		}).Render(r.Context(), w)
	}
}

func (i *implViewController) ChartpageHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vpage.Chartpage(dto.ViewBaseDTO{
			Name: "Chart",
		}).Render(r.Context(), w)
	}
}
