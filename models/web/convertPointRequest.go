package web

type ConvertPointRequest struct {
	Point      int `json:"point" validate:"required,numeric"`
	ValuePoint int  `json:"valuePoint" validate:"required,numeric"`
}
