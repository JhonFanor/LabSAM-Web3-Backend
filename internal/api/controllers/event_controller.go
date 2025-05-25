package controllers

import (
	"encoding/json"
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

type EventControllerParams struct {
	fx.In
	EventService        services.EventService
	LocalitationService services.LocalitationService
}

type EventController struct {
	service             services.EventService
	localitationService services.LocalitationService
}

func NewEventController(p EventControllerParams) *EventController {
	return &EventController{
		service:             p.EventService,
		localitationService: p.LocalitationService,
	}
}

func (e *EventController) CreateEvent(ctx *gin.Context) {

	validatedInput, _ := ctx.Get("input")

	eventRequest := validatedInput.(*requests.EventRequest)

	claimsValue, _ := ctx.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	var event models.Event
	if err := mapstructure.Decode(eventRequest, &event); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	event.Date = eventRequest.Date

	if eventRequest.Localiatation != nil {
		var localitationRequest models.Localitation
		if err := mapstructure.Decode(eventRequest.Localiatation, &localitationRequest); err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
			return
		}

		localitation, err := e.localitationService.AssignLocalitation(&localitationRequest)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}

		event.LocalitationID = localitation.ID
	}

	createdEvent, err := e.service.CreateEvent(&event, claims.UserID, eventRequest.SubtopicIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdEvent)
}

func (e *EventController) GetAllEvents(c *gin.Context) {
	pagination, err := e.service.GetAllEvents(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.EventGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (e *EventController) GetEventByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	event, err := e.service.GetEventByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Event not found"})
		return
	}

	var eventResponse responses.EventGetResponse
	if err := mapstructure.Decode(event, &eventResponse); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	c.JSON(http.StatusOK, eventResponse)
}

func (e *EventController) UpdateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	eventRequest := validatedInput.(*requests.EventUpdateRequest)
	var event models.Event
	if err := mapstructure.Decode(eventRequest, &event); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	event.ID = uint(id)
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if eventRequest.Localiatation != nil {
		var localitationRequest models.Localitation
		if err := mapstructure.Decode(eventRequest.Localiatation, &localitationRequest); err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
			return
		}

		localitation, err := e.localitationService.AssignLocalitation(&localitationRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
			return
		}

		event.LocalitationID = localitation.ID
	}

	if err := e.service.UpdateEvent(&event, claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (e *EventController) DeleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := e.service.DeleteEvent(uint(id), claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
