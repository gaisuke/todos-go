package services

import (
	"todos-go/models"

	"gorm.io/gorm"
)

type SublistService interface {
	GetAllByListID(listID, page, limit int, filters map[string]interface{}) ([]models.Sublist, int, error)
	GetByID(id int) (*models.Sublist, error)
	Create(sublist *models.Sublist) error
	Update(sublist *models.Sublist) error
	Delete(id int) error
}

type sublistService struct {
	db *gorm.DB
}

// Create implements SublistService.
func (s *sublistService) Create(sublist *models.Sublist) error {
	return s.db.Create(sublist).Error
}

// Delete implements SublistService.
func (s *sublistService) Delete(id int) error {
	return s.db.Delete(&models.Sublist{ID: id}).Error
}

// GetAllByListID implements SublistService.
func (s *sublistService) GetAllByListID(listID int, page int, limit int, filters map[string]interface{}) ([]models.Sublist, int, error) {
	var sublists []models.Sublist
	var totalCount int64

	query := s.db.Model(&models.Sublist{}).Where("list_id = ?", listID)

	for key, value := range filters {
		query = query.Where(key+" LIKE ?", "%"+value.(string)+"%")
	}

	query.Count(&totalCount)
	query.Offset((page - 1) * limit).Limit(limit).Find(&sublists)

	return sublists, int(totalCount), nil
}

// GetByID implements SublistService.
func (s *sublistService) GetByID(id int) (*models.Sublist, error) {
	var sublist models.Sublist
	if err := s.db.Preload("List").First(&sublist, id).Error; err != nil {
		return nil, err
	}
	return &sublist, nil
}

// Update implements SublistService.
func (s *sublistService) Update(sublist *models.Sublist) error {
	return s.db.Save(sublist).Error
}

func NewSublistService(db *gorm.DB) SublistService {
	return &sublistService{db}
}
