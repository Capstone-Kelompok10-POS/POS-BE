package repository

import (
	"qbills/models/domain"
	"qbills/models/schema"
	req "qbills/utils/request"
	res "qbills/utils/response"

	"gorm.io/gorm"
)

type ConvertPointRepository interface {
	Create(convertPoint *domain.ConvertPoint) (*domain.ConvertPoint, error)
	Update(convertPoint *domain.ConvertPoint, id int) (*domain.ConvertPoint, error)
	FindById(id int) (*domain.ConvertPoint, error)
	FindAll() ([]domain.ConvertPoint, error)
	Delete(id int) error
}

type ConvertPointRepositoryImpl struct {
	DB *gorm.DB
}

func NewConvertPointRepository(db *gorm.DB) ConvertPointRepository {
	return &ConvertPointRepositoryImpl{DB: db}
}

func (repository *ConvertPointRepositoryImpl) Create(convertPoint *domain.ConvertPoint) (*domain.ConvertPoint, error) {
	convertPointDB := req.ConvertPointDomainToConvertPointSchema(*convertPoint)
	result := repository.DB.Create(&convertPointDB)
	if result.Error != nil {
		return nil,result.Error
	}

	results := res.ConvertPointSchemaToConvertPointDomain(convertPointDB)
	return results, nil
}

func (repository *ConvertPointRepositoryImpl) Update(convertPoint *domain.ConvertPoint, id int) (*domain.ConvertPoint, error) {
	result := repository.DB.Table("convert_points").Where("id = ?", id).Updates(domain.ConvertPoint{Point: convertPoint.Point, ValuePoint: convertPoint.ValuePoint})
	if result.Error != nil {
		return nil, result.Error
	}

	return convertPoint, nil

}

func (repository *ConvertPointRepositoryImpl) FindById(id int) (*domain.ConvertPoint, error) {
	var convertPoint domain.ConvertPoint

	if err := repository.DB.First(&convertPoint, id).Error; err != nil {
		return nil, err
	}
	query := "SELECT * FROM convert_points WHERE convert_points.id = ? AND convert_points.deleted_at IS NULL"
	result := repository.DB.Raw(query, id).Scan(&convertPoint)
	if result.Error != nil {
		return nil, result.Error
	}

	return &convertPoint, nil
}

func (repository *ConvertPointRepositoryImpl) FindAll() ([]domain.ConvertPoint, error) {
	convertPoint := []domain.ConvertPoint{}
	query := "SELECT * FROM convert_points WHERE deleted_at IS NULL"
	result := repository.DB.Raw(query).Scan(&convertPoint)
	if result.Error != nil {
		return nil, result.Error
	}
	return convertPoint, nil
}


func (repository *ConvertPointRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.ConvertPoint{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}