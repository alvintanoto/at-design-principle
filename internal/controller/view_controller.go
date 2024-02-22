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
	ButtonpageHandler() func(http.ResponseWriter, *http.Request)
	ChartpageHandler() func(http.ResponseWriter, *http.Request)
	AlertPageHandler() func(http.ResponseWriter, *http.Request)
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

func (i *implViewController) ButtonpageHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		buttonType := r.PostFormValue("button_type")
		if buttonType == "" {
			buttonType = "Primary"
		}

		vpage.Buttonpage(dto.ButtonViewDTO{
			ViewBaseDTO: dto.ViewBaseDTO{
				Name: "Button",
			},
			Type: buttonType,
		}).Render(r.Context(), w)
	}
}

func (i *implViewController) ChartpageHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vpage.Chartpage(dto.ChartViewDTO{
			ViewBaseDTO: dto.ViewBaseDTO{
				Name: "Chart",
			},
			Dataset: []dto.ChartDataset{
				{Year: "2010", AcquisitionCount: "10"},
				{Year: "2011", AcquisitionCount: "22"},
				{Year: "2012", AcquisitionCount: "15"},
				{Year: "2013", AcquisitionCount: "25"},
				{Year: "2014", AcquisitionCount: "22"},
				{Year: "2015", AcquisitionCount: "30"},
				{Year: "2016", AcquisitionCount: "28"},
			},
		}).Render(r.Context(), w)
	}
}

func (i *implViewController) AlertPageHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vpage.Alertpage(dto.ViewBaseDTO{
			Name: "Alert",
		}).Render(r.Context(), w)
	}
}
