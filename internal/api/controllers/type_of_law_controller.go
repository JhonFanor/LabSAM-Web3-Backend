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

type TypeOfLawControllerParams struct {
	fx.In
	TypeOfLawService services.TypeOfLawService
}

type TypeOfLawController struct {
	service services.TypeOfLawService
}

func NewTypeOfLawController(p TypeOfLawControllerParams) *TypeOfLawController {
	return &TypeOfLawController{
		service: p.TypeOfLawService,
	}
}

func (t *TypeOfLawController) GetAllTypeOfLaw(c *gin.Context) {
	typeOfLaw, err := t.service.GetAllTypeOfLaw()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	var typeOfLawResponse []responses.TypeOfLawResponse
	if err := mapstructure.Decode(typeOfLaw, &typeOfLawResponse); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	c.JSON(http.StatusOK, typeOfLawResponse)
}
