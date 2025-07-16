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

func (n *NewsController) CreateNews(c *gin.Context) {
	validatedInput, _ := c.Get("input")
	newsRequest := validatedInput.(*requests.NewsRequest)

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	var news models.News
	if err := mapstructure.Decode(newsRequest, &news); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	news.Date = &newsRequest.Date

	createdNews, err := n.service.CreateNews(&news, claims.UserID, newsRequest.SubtopicIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdNews)
}

func (n *NewsController) GetAllNews(c *gin.Context) {
	pagination, err := n.service.GetAllNews(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.NewsGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (n *NewsController) GetAllNewsByUserID(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	pagination, err := n.service.GetAllNewsByUserID(c, claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.NewsGetAllByUserIDResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (n *NewsController) GetAllNewsNotApproved(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	pagination, err := n.service.GetAllNewsNotApproved(c, claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.NewsGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (n *NewsController) CountNewsNotApproved(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	count, err := n.service.CountNewsNotApproved(claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.CountResponse{Count: count})
}

func (n *NewsController) GetNewsByID(c *gin.Context) {
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

	news, err := n.service.GetNewsByID(uint(id), userID, role)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "News not found"})
		return
	}

	var newsResponse responses.NewsGetResponse
	if err := mapstructure.Decode(news, &newsResponse); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	c.JSON(http.StatusOK, newsResponse)
}

func (n *NewsController) UpdateNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	validatedInput, _ := c.Get("input")
	newsRequest := validatedInput.(*requests.NewsUpdateRequest)
	var news models.News
	if err := mapstructure.Decode(newsRequest, &news); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: consts.ErrorMapConst})
		return
	}

	news.ID = uint(id)
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := n.service.UpdateNews(&news, claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, news)
}

func (b *NewsController) DeleteNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "Invalid ID"})
		return
	}

	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	if err := b.service.DeleteNews(uint(id), claims.UserID, claims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully deleted"})
}
