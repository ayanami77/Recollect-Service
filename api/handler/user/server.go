package user

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/jwtutil"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/myerror"
	"github.com/Seiya-Tagami/Recollect-Service/api/response"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:generate mockgen -source=$GOFILE -destination=$GOPATH/Recollect-Service/api/mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE
type Handler interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	CheckEmailDuplication(c *gin.Context)
	CheckUserIDDuplication(c *gin.Context)
	AnalyzeUserHistory(c *gin.Context)
}

type handler struct {
	userInteractor user.Interactor
}

func New(userInteractor user.Interactor) Handler {
	return &handler{userInteractor}
}

func (h *handler) GetUser(c *gin.Context) {
	sub, err := jwtutil.SubFromBearerToken(c)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	user, err := h.userInteractor.GetUser(sub)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	userResponse := response.ToUserResponse(&user)

	c.JSON(http.StatusOK, userResponse)
}

func (h *handler) CreateUser(c *gin.Context) {
	sub, err := jwtutil.SubFromBearerToken(c)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	userReq := entity.User{}
	if err := c.BindJSON(&userReq); err != nil {
		panic(err)
	}

	userReq.Sub = sub

	user, err := h.userInteractor.CreateUser(userReq)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}
	userResponse := response.ToUserResponse(&user)

	c.JSON(http.StatusOK, userResponse)
}

func (h *handler) UpdateUser(c *gin.Context) {
	sub, err := jwtutil.SubFromBearerToken(c)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	userReq := entity.User{}
	if err := c.BindJSON(&userReq); err != nil {
		myerror.HandleError(c, err)
		return
	}

	userReq.Sub = sub

	user, err := h.userInteractor.UpdateUser(userReq, sub)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	userResponse := response.ToUserResponse(&user)

	c.JSON(http.StatusOK, userResponse)
}

func (h *handler) DeleteUser(c *gin.Context) {
	sub, err := jwtutil.SubFromBearerToken(c)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	err = h.userInteractor.DeleteUser(sub)
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

func (h *handler) AnalyzeUserHistory(c *gin.Context) {
	sub, err := jwtutil.SubFromBearerToken(c)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	user, err := h.userInteractor.AnalyzeUserHistory(sub)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	userResponse := response.ToUserResponse(&user)

	c.JSON(http.StatusOK, userResponse)
}

type EmailRequest struct {
	Email string `json:"email"`
}

type UserIDRequest struct {
	UserID string `json:"user_id"`
}
