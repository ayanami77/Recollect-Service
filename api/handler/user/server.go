package user

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/myerror"
	"github.com/Seiya-Tagami/Recollect-Service/api/response"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	CheckEmailDuplication(c *gin.Context)
	CheckUserIDDuplication(c *gin.Context)
}

type handler struct {
	userInteractor user.Interactor
}

func New(userInteractor user.Interactor) Handler {
	return &handler{userInteractor}
}

func (h *handler) CreateUser(c *gin.Context) {
	userReq := entity.User{}
	if err := c.BindJSON(&userReq); err != nil {
		panic(err)
	}

	user, err := h.userInteractor.CreateUser(userReq)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}
	userResponse := response.ToUserResponse(&user)

	c.JSON(http.StatusOK, userResponse)
}

func (h *handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	userReq := entity.User{}
	if err := c.BindJSON(&userReq); err != nil {
		myerror.HandleError(c, err)
		return
	}

	user, err := h.userInteractor.UpdateUser(userReq, id)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	userResponse := response.ToUserResponse(&user)

	c.JSON(http.StatusOK, userResponse)
}

func (h *handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.userInteractor.DeleteUser(id)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *handler) CheckEmailDuplication(c *gin.Context) {
	emailReq := EmailRequest{}
	if err := c.BindJSON(&emailReq); err != nil {
		myerror.HandleError(c, err)
		return
	}

	isDuplicated, err := h.userInteractor.CheckEmailDuplication(emailReq.Email)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, isDuplicated)
}

func (h *handler) CheckUserIDDuplication(c *gin.Context) {
	userIDReq := UserIDRequest{}
	if err := c.BindJSON(&userIDReq); err != nil {
		myerror.HandleError(c, err)
		return
	}

	isDuplicated, err := h.userInteractor.CheckUserIDDuplication(userIDReq.UserID)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, isDuplicated)
}

type EmailRequest struct {
	Email string `json:"email"`
}

type UserIDRequest struct {
	UserID string `json:"user_id"`
}
