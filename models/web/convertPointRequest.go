package web

type ConvertPointRequest struct {
<<<<<<< Updated upstream
	Point      uint `json:"point" validate:"required"`
	ValuePoint int  `json:"value_point" validate:"required"`
=======
<<<<<<< Updated upstream
	Point      uint `json:"point" validate:"required"`
	ValuePoint int  `json:"value_point" validate:"required"`
=======
<<<<<<< Updated upstream
	Point      uint `json:"point" validate:"required"`
	ValuePoint int  `json:"value_point" validate:"required"`
=======
	Point      int `json:"point" validate:"required,numeric"`
	ValuePoint int  `json:"valuePoint" validate:"required,numeric"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}
