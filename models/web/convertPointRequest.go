package web

type ConvertPointRequest struct {
	Point      uint `json:"point" validate:"required,numeric"`
	ValuePoint int  `json:"valuePoint" validate:"required,numeric"`
}
