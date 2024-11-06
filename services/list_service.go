package services

import (
	"todos-go/models"

	"gorm.io/gorm"
)

type ListService interface {
	GetAll(page, limit int, filters map[string]interface{}) ([]models.List, int, error)
	GetByID(id int) (*models.List, error)
	Create(list *models.List) error
	Update(list *models.List) error
	Delete(id int) error
}

type listService struct {
	db *gorm.DB
}

// Create implements ListService.
func (s *listService) Create(list *models.List) error {
	return s.db.Create(list).Error
}

// Delete implements ListService.
func (s *listService) Delete(id int) error {
	return s.db.Delete(&models.List{ID: id}).Error
}

// GetAll implements ListService.
func (s *listService) GetAll(page int, limit int, filters map[string]interface{}) ([]models.List, int, error) {
	var lists []models.List
	var totalCount int64

	query := s.db.Model(&models.List{})

	for key, value := range filters {
		query = query.Where(key+" LIKE ?", "%"+value.(string)+"%")
	}

	query.Count(&totalCount)
	query.Offset((page - 1) * limit).Limit(limit).Find(&lists)

	return lists, int(totalCount), nil
}

// GetByID implements ListService.
func (s *listService) GetByID(id int) (*models.List, error) {
	var list models.List
	if err := s.db.Preload("Sublists").First(&list, id).Error; err != nil {
		return nil, err
	}
	return &list, nil
}

// Update implements ListService.
func (s *listService) Update(list *models.List) error {
	return s.db.Save(list).Error
}

func NewListService(db *gorm.DB) ListService {
	return &listService{db: db}
}
