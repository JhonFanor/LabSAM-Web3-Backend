package controllers

import (
	"lamsam-web3-backend/internal/consts"
	"lamsam-web3-backend/internal/dto/requests"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/fx"
)

type NewsControllerParams struct {
	fx.In
	NewsService services.NewsService
}

type NewsController struct {
	service services.NewsService
}

func NewNewsController(p NewsControllerParams) *NewsController {
	return &NewsController{
		service: p.NewsService,
	}
}

// CreateNews godoc
// @Summary Create a new News entry
// @Description This endpoint creates a new News record. Requires authentication.
// @Tags News
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param input body requests.NewsRequest true "News Information"
// @Success 201 {object} models.News "Successfully created"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /news/create [post]
func (n *NewsController) CreateNews(ctx *gin.Context) {

	validatedInput, _ := ctx.Get("input")

	newsRequest := validatedInput.(*requests.NewsRequest)

	claimsValue, _ := ctx.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	var news models.News
	if err := mapstructure.Decode(newsRequest, &news); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	createdNews, err := n.service.CreateNews(&news, claims.UserID, newsRequest.SubtopicIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdNews)
}
