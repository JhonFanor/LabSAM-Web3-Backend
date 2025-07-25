package controllers

import (
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type UniversityTypeControllerParams struct {
	fx.In
	UniversityTypeService services.UniversityTypeService
}

type UniversityTypeController struct {
	service services.UniversityTypeService
}

func NewUniversityTypeController(p UniversityTypeControllerParams) *UniversityTypeController {
	return &UniversityTypeController{
		service: p.UniversityTypeService,
	}
}

func (ctrl *UniversityTypeController) GetAllUniversityTypes(c *gin.Context) {
	types, err := ctrl.service.GetAllUniversityTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: "Failed to fetch university types",
		})
		return
	}

	c.JSON(http.StatusOK, types)
}
