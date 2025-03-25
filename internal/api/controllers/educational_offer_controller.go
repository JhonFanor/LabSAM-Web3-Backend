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

type EducationalOfferControllerParams struct {
	fx.In
	EducationalOfferService services.EducationalOfferService
}

type EducationalOfferController struct {
	service services.EducationalOfferService
}

func NewEducationalOfferController(p EducationalOfferControllerParams) *EducationalOfferController {
	return &EducationalOfferController{
		service: p.EducationalOfferService,
	}
}

func (e *EducationalOfferController) CreateEducationalOffer(ctx *gin.Context) {

	validatedInput, _ := ctx.Get("input")

	educationalOfferRequest := validatedInput.(*requests.EducationalOfferRequest)

	claimsValue, _ := ctx.Get("claims")

	claims, _ := claimsValue.(*security.Claims)

	var educationalOffer models.EducationalOffer
	if err := mapstructure.Decode(educationalOfferRequest, &educationalOffer); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	createdEducationalOffer, err := e.service.CreateEducationalOffer(&educationalOffer, claims.UserID, educationalOfferRequest.SubtopicIDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdEducationalOffer)
}

func (e *EducationalOfferController) GetAllEducationalOffers(c *gin.Context) {
	pagination, err := e.service.GetAllEducationalOffers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, pagination)
}

func (e *EducationalOfferController) GetEducationalOfferByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	educationalOffer, err := e.service.GetEducationalOfferByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Educational Offer not found"})
		return
	}

	c.JSON(http.StatusOK, educationalOffer)
}

func (e *EducationalOfferController) UpdateEducationalOffer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	educationalOfferRequest := validatedInput.(*requests.EducationalUpdateOfferRequest)
	var educationalOffer models.EducationalOffer
	if err := mapstructure.Decode(educationalOfferRequest, &educationalOffer); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	educationalOffer.ID = uint(id)
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := e.service.UpdateEducationalOffer(&educationalOffer, claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, educationalOffer)
}

func (e *EducationalOfferController) DeleteEducationalOffer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := e.service.DeleteEducationalOffer(uint(id), claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
