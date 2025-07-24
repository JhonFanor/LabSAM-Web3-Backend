package controllers

import (
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type PublicationsControllerParams struct {
	fx.In
	PublicationsService services.PublicationsService
}

type PublicationsController struct {
	service services.PublicationsService
}

func NewPublicationsController(p PublicationsControllerParams) *PublicationsController {
	return &PublicationsController{
		service: p.PublicationsService,
	}
}

func (pc *PublicationsController) GetAllPublications(c *gin.Context) {
	publications, err := pc.service.GetAllPublications(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: "Publications not found",
		})
		return
	}

	c.JSON(http.StatusOK, publications)
}
