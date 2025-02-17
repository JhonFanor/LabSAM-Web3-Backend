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

type EventSubtopicControllerParams struct {
	fx.In
	EventSubtopicService services.EventSubtopicService
	EventService         services.EventService
}

type EventSubtopicController struct {
	service      services.EventSubtopicService
	eventService services.EventService
}

func NewEventSubtopicController(p EventSubtopicControllerParams) *EventSubtopicController {
	return &EventSubtopicController{
		service:      p.EventSubtopicService,
		eventService: p.EventService,
	}
}

func (e *EventSubtopicController) CreateEventSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	eventSubtopicRequest := validatedInput.(*requests.EventSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	event, err := e.eventService.GetEventByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Event not found"})
		return
	}

	if event.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range eventSubtopicRequest.SubtopicIDs {
		eventSubtopic := &models.EventSubtopic{
			EventID:    uint(id),
			SubtopicID: uint(subtopicID),
		}

		_, err := e.service.CreateEventSubtopic(eventSubtopic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{Message: "Successfully assigned subtopics"})
}

func (e *EventSubtopicController) DeleteEventSubtopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	validatedInput, _ := c.Get("input")
	eventSubtopicRequest := validatedInput.(*requests.EventSubtopicRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	event, err := e.eventService.GetEventByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Event not found"})
		return
	}

	if event.UserID != claims.UserID && claims.Role != "admin" {
		c.JSON(http.StatusUnauthorized, customerrors.ErrUnauthorized)
		return
	}

	for _, subtopicID := range eventSubtopicRequest.SubtopicIDs {
		if err := e.service.DeleteEventSubtopic(uint(id), subtopicID); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
