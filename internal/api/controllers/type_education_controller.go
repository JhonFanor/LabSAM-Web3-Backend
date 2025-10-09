package controllers

import (
	"lamsam-web3-backend/internal/consts"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/fx"
)

type TypeEducationControllerParams struct {
	fx.In
	TypeEducationService services.TypeEducationService
}

type TypeEducationController struct {
	service services.TypeEducationService
}

func NewTypeEducationController(p TypeEducationControllerParams) *TypeEducationController {
	return &TypeEducationController{
		service: p.TypeEducationService,
	}
}

func (t *TypeEducationController) GetAllTypeEducation(c *gin.Context) {
	typeEducation, err := t.service.GetAllTypeEducation()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	var typeEducationResponse []responses.TypeEducationResponse
	if err := mapstructure.Decode(typeEducation, &typeEducationResponse); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	c.JSON(http.StatusOK, typeEducationResponse)
}
