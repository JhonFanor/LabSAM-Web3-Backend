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

	educationalOffer.StartDate = educationalOfferRequest.StartDate
	educationalOffer.EndDate = educationalOfferRequest.EndDate

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

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.EducationalOfferGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (e *EducationalOfferController) GetAllEducationalOffersByUserID(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	pagination, err := e.service.GetAllEducationalOffersByUserID(c, claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.EducationalOfferGetAllByUserIDResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (e *EducationalOfferController) GetAllEducationalOffersNotApproved(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	pagination, err := e.service.GetAllEducationalOffersNotApproved(c, claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.EducationalOfferGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (e *EducationalOfferController) CountEducationalOffersNotApproved(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	count, err := e.service.CountEducationalOffersNotApproved(claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.CountResponse{Count: count})
}

func (e *EducationalOfferController) GetEducationalOfferByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	var userID uint
	var role string

	if claimsValue, exists := c.Get("claims"); exists {
		if claims, ok := claimsValue.(*security.Claims); ok {
			userID = claims.UserID
			role = claims.Role
		}
	}

	educationalOffer, err := e.service.GetEducationalOfferByID(uint(id), userID, role)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Educational Offer not found"})
		return
	}

	var educationalOfferResponse responses.EducationalOfferGetResponse
	if err := mapstructure.Decode(educationalOffer, &educationalOfferResponse); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	c.JSON(http.StatusOK, educationalOfferResponse)
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

func (e *EducationalOfferController) SetEducationalOfferApproval(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	req := validatedInput.(*requests.ApprovalRequest)

	claimsValue, _ := c.Get("claims")
	claims := claimsValue.(*security.Claims)

	err = e.service.SetEducationalOfferApproval(uint(id), req.Approved, claims.Role)
	if err != nil {
		if err == customerrors.ErrUnauthorized {
			c.JSON(http.StatusForbidden, responses.ErrorResponse{Error: err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Approval status updated"})
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
