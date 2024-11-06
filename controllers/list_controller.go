package controllers

import (
	"todos-go/models"
	"todos-go/services"
)

type ListController interface {
	GetAll(page, limit int, filters map[string]interface{}) ([]models.List, int, error)
	GetByID(id int) (*models.List, error)
	Create(list *models.List) error
	Update(list *models.List) error
	Delete(id int) error
}

type listController struct {
	service services.ListService
}

// Create implements ListController.
func (c *listController) Create(list *models.List) error {
	return c.service.Create(list)
}

// Delete implements ListController.
func (c *listController) Delete(id int) error {
	return c.service.Delete(id)
}

// GetAll implements ListController.
func (c *listController) GetAll(page int, limit int, filters map[string]interface{}) ([]models.List, int, error) {
	return c.service.GetAll(page, limit, filters)
}

// GetByID implements ListController.
func (c *listController) GetByID(id int) (*models.List, error) {
	return c.service.GetByID(id)
}

// Update implements ListController.
func (c *listController) Update(list *models.List) error {
	return c.service.Update(list)
}

func NewListController(service services.ListService) ListController {
	return &listController{
		service: service,
	}
}
