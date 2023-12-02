package web

type ConvertPointRequest struct {
	Point      uint `json:"point" validate:"required"`
	ValuePoint int  `json:"value_point" validate:"required"`
}
