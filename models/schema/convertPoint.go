package schema

import (
	"time"

	"gorm.io/gorm"
)

type ConvertPoint struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

<<<<<<< Updated upstream
	Point uint `json:"point"`
<<<<<<< Updated upstream
	ValuePoint int `json:"value_point"`
=======
<<<<<<< Updated upstream
	PointValue int `json:"point_value"`
=======
<<<<<<< Updated upstream
	PointValue int `json:"point_value"`
=======
<<<<<<< Updated upstream
	ValuePoint int `json:"value_point"`
=======
	PointValue int `json:"point_value"`
=======
<<<<<<< Updated upstream
	Point uint `json:"point"`
<<<<<<< Updated upstream
	ValuePoint int `json:"value_point"`
=======
	PointValue int `json:"point_value"`
=======
	Point int `json:"point"`
	ValuePoint int `json:"valuePoint"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}