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

type EventControllerParams struct {
	fx.In
	EventService services.EventService
}

type EventController struct {
	service services.EventService
}

func NewEventController(p EventControllerParams) *EventController {
	return &EventController{
		service: p.EventService,
	}
}

// CreateEvent godoc
// @Summary Create a new Event entry
// @Description This endpoint creates a new Event record. Requires authentication.
// @Tags Event
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param input body requests.EventRequest true "Event Information"
// @Success 201 {object} models.Event "Successfully created"
// @Failure 400 {object} responses.ErrorResponse "Bad request"
// @Failure 401 {object} responses.ErrorResponse "Unauthorized"
// @Failure 500 {object} responses.ErrorResponse "Internal server error"
// @Router /event/create [post]
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

	createdEvent, err := e.service.CreateEvent(&event, claims.UserID, eventRequest.SubtopicIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdEvent)
}
