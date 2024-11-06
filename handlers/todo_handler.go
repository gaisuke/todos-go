package handlers

import (
	"net/http"
	"strconv"
	"todos-go/controllers"
	"todos-go/models"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	listController    controllers.ListController
	sublistController controllers.SublistController
}

func NewTodoHandler(listController controllers.ListController, sublistController controllers.SublistController) *TodoHandler {
	return &TodoHandler{
		listController:    listController,
		sublistController: sublistController,
	}
}

func (h *TodoHandler) GetLists(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	filters := map[string]interface{}{
		"title":       ctx.QueryParam("title"),
		"description": ctx.QueryParam("description"),
	}

	lists, total, err := h.listController.GetAll(page, limit, filters)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":  lists,
		"total": total,
	})
}

func (h *TodoHandler) GetListByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	list, err := h.listController.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, list)
}

func (h *TodoHandler) GetSublistsByListID(ctx echo.Context) error {
	listID, _ := strconv.Atoi(ctx.Param("id"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	filters := map[string]interface{}{
		"title":       ctx.QueryParam("title"),
		"description": ctx.QueryParam("description"),
	}

	sublists, total, err := h.sublistController.GetAllByListID(listID, page, limit, filters)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":  sublists,
		"total": total,
	})
}

func (h *TodoHandler) GetSublistByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	sublist, err := h.sublistController.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, sublist)
}

func (h *TodoHandler) CreateList(ctx echo.Context) error {
	var list models.List
	if err := ctx.Bind(&list); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := h.listController.Create(&list); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, list)
}

func (h *TodoHandler) CreateSublist(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var sublist models.Sublist
	if err := ctx.Bind(&sublist); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	sublist.ListID = id

	if err := h.sublistController.Create(&sublist); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, sublist)
}

func (h *TodoHandler) UpdateList(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var list models.List
	if err := ctx.Bind(&list); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	list.ID = id

	if err := h.listController.Update(&list); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, list)
}

func (h *TodoHandler) UpdateSublist(ctx echo.Context) error {
	listId, _ := strconv.Atoi(ctx.Param("list_id"))
	id, _ := strconv.Atoi(ctx.Param("id"))
	var sublist models.Sublist
	if err := ctx.Bind(&sublist); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	sublist.ListID = listId
	sublist.ID = id

	if err := h.sublistController.Update(&sublist); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, sublist)
}

func (h *TodoHandler) DeleteList(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.listController.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "List deleted successfully",
	})
}

func (h *TodoHandler) DeleteSublist(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := h.sublistController.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Sublist deleted successfully",
	})
}
