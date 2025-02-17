package controllers

import (
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto/requests"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type NewsSubtopicControllerParams struct {
	fx.In
	NewsSubtopicService services.NewsSubtopicService
	NewsService         services.NewsService
}

type NewsSubtopicController struct {
	service     services.NewsSubtopicService
	newsService services.NewsService
}

func NewNewsSubtopicController(p NewsSubtopicControllerParams) *NewsSubtopicController {
	return &NewsSubtopicController{
		service:     p.NewsSubtopicService,
		newsService: p.NewsService,
	}
}

func (n *NewsSubtopicController) CreateNewsSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	newsSubtopicRequest := validatedInput.(*requests.NewsSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	news, err := n.newsService.GetNewsByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "News not found"})
		return
	}

	if news.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range newsSubtopicRequest.SubtopicIDs {
		newsSubtopic := &models.NewsSubtopic{
			NewsID:     uint(id),
			SubtopicID: uint(subtopicID),
		}

		_, err := n.service.CreateNewsSubtopic(newsSubtopic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{Message: "Successfully assigned subtopics"})
}

func (n *NewsSubtopicController) DeleteNewsSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	newsSubtopicRequest := validatedInput.(*requests.NewsSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	news, err := n.newsService.GetNewsByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "News not found"})
		return
	}

	if news.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range newsSubtopicRequest.SubtopicIDs {
		if err := n.service.DeleteNewsSubtopic(uint(id), subtopicID); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
