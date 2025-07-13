package controllers

import (
	"lamsam-web3-backend/internal/consts"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto/requests"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/models"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/fx"
)

type TopicControllerParams struct {
	fx.In
	TopicService services.TopicService
}

type TopicController struct {
	service services.TopicService
}

func NewTopicController(p TopicControllerParams) *TopicController {
	return &TopicController{
		service: p.TopicService,
	}
}

func (t *TopicController) CreateTopic(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if claims.Role == "admin" {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	topicRequest := validatedInput.(*requests.TopicRequest)

	var topic models.Topic
	if err := mapstructure.Decode(topicRequest, &topic); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	createdTopic, err := t.service.CreateTopic(&topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTopic)
}

func (t *TopicController) GetAllTopics(c *gin.Context) {
	topics, err := t.service.GetAllTopic()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	var topicsResponse []responses.TopicGetAllResponse
	if err := mapstructure.Decode(topics, &topicsResponse); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	c.JSON(http.StatusOK, topicsResponse)
}

func (b *TopicController) GetTopicByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	topic, err := b.service.GetTopicByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Topic not found"})
		return
	}

	c.JSON(http.StatusOK, topic)
}

func (b *TopicController) UpdateTopic(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if claims.Role == "admin" {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	topicRequest := validatedInput.(*requests.TopicRequest)
	var topic models.Topic
	if err := mapstructure.Decode(topicRequest, &topic); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	topic.ID = uint(id)
	if err := b.service.UpdateTopic(&topic); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, topic)
}

func (b *TopicController) DeleteTopic(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if claims.Role == "admin" {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: customerrors.ErrUnauthorized.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	if err := b.service.DeleteTopic(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
