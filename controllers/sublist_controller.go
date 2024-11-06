package controllers

import (
	"todos-go/models"
	"todos-go/services"
)

type SublistController interface {
	GetAllByListID(listID, page, limit int, filters map[string]interface{}) ([]models.Sublist, int, error)
	GetByID(id int) (*models.Sublist, error)
	Create(sublist *models.Sublist) error
	Update(sublist *models.Sublist) error
	Delete(id int) error
}

type sublistController struct {
	service services.SublistService
}

// Create implements SublistController.
func (c *sublistController) Create(sublist *models.Sublist) error {
	return c.service.Create(sublist)
}

// Delete implements SublistController.
func (c *sublistController) Delete(id int) error {
	return c.service.Delete(id)
}

// GetAllByListID implements SublistController.
func (c *sublistController) GetAllByListID(listID int, page int, limit int, filters map[string]interface{}) ([]models.Sublist, int, error) {
	return c.service.GetAllByListID(listID, page, limit, filters)
}

// GetByID implements SublistController.
func (c *sublistController) GetByID(id int) (*models.Sublist, error) {
	return c.service.GetByID(id)
}

// Update implements SublistController.
func (c *sublistController) Update(sublist *models.Sublist) error {
	return c.service.Update(sublist)
}

func NewSublistController(service services.SublistService) SublistController {
	return &sublistController{
		service: service,
	}
}
