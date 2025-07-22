package controllers

import (
	"encoding/json"
	"lamsam-web3-backend/internal/consts"
	"lamsam-web3-backend/internal/customerrors"
	"lamsam-web3-backend/internal/dto/responses"
	"lamsam-web3-backend/internal/services"
	"lamsam-web3-backend/pkg/security"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/fx"
)

type NotificationControllerParams struct {
	fx.In
	Service services.NotificationService
}

type NotificationController struct {
	service services.NotificationService
}

func NewNotificationController(p NotificationControllerParams) *NotificationController {
	return &NotificationController{
		service: p.Service,
	}
}

func (n *NotificationController) GetAll(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims := claimsValue.(*security.Claims)

	pagination, err := n.service.GetAllNotifications(c, claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, _ := json.Marshal(pagination.Data)
	var list []responses.NotificationGetAllResponse
	_ = json.Unmarshal(jsonData, &list)
	pagination.Data = list

	c.JSON(http.StatusOK, pagination)
}

func (n *NotificationController) CountNotificationNotRead(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims, _ := claimsValue.(*security.Claims)

	count, err := n.service.CountNotificationNotRead(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.CountResponse{Count: count})
}

func (n *NotificationController) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	claimsValue, _ := c.Get("claims")
	claims := claimsValue.(*security.Claims)

	noti, err := n.service.GetNotificationByID(uint(id), claims.UserID)
	if err != nil {
		status := http.StatusInternalServerError
		if err == customerrors.ErrUnauthorized {
			status = http.StatusForbidden
		}
		c.JSON(status, responses.ErrorResponse{Error: err.Error()})
		return
	}

	var notificationResponse responses.NotificationGetResponse
	if err := mapstructure.Decode(noti, &notificationResponse); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Error: consts.ErrorMapConst,
		})
		return
	}

	c.JSON(http.StatusOK, noti)
}

func (n *NotificationController) MarkAsRead(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	claimsValue, _ := c.Get("claims")
	claims := claimsValue.(*security.Claims)

	err = n.service.UpdateIsRead(uint(id), claims.UserID)
	if err != nil {
		status := http.StatusInternalServerError
		if err == customerrors.ErrUnauthorized {
			status = http.StatusForbidden
		}
		c.JSON(status, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Notification marked as read"})
}

func (n *NotificationController) MarkAllAsRead(c *gin.Context) {
	claimsValue, _ := c.Get("claims")
	claims := claimsValue.(*security.Claims)

	err := n.service.UpdateAllIsReadByUserID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "All notifications marked as read"})
}

func (n *NotificationController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: customerrors.ErrInvalidID.Error()})
		return
	}

	claimsValue, _ := c.Get("claims")
	claims := claimsValue.(*security.Claims)

	err = n.service.DeleteNotification(uint(id), claims.UserID)
	if err != nil {
		status := http.StatusInternalServerError
		if err == customerrors.ErrUnauthorized {
			status = http.StatusForbidden
		}
		c.JSON(status, responses.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{Message: "Notification deleted"})
}
