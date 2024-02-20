package dto

type ViewBaseDTO struct {
	Name string
}

type ChartViewDTO struct {
	ViewBaseDTO
	Dataset []ChartDataset
}

type ChartDataset struct {
	AcquisitionCount string `json:"count"`
	Year             string `json:"year"`
}
