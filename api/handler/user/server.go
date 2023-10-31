package user

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/myerror"
	"net/http"
	"os"

	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/response"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/user"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	LoginUser(c *gin.Context)
	LogoutUser(c *gin.Context)
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

func (h *handler) LoginUser(c *gin.Context) {
	userReq := entity.User{}
	if err := c.BindJSON(&userReq); err != nil {
		myerror.HandleError(c, err)
		return
	}

	tokenString, err := h.userInteractor.LoginUser(userReq.UserID, userReq.Password)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)

	// sameSite = Noneの時は、secure属性をつけてあげるようにする。
	c.SetCookie("user_token", tokenString, 3600, "/", os.Getenv("API_DOMAIN"), true, true)

	c.Status(http.StatusNoContent)
}

func (h *handler) LogoutUser(c *gin.Context) {
	c.SetSameSite(http.SameSiteNoneMode)

	c.SetCookie("user_token", "", 0, "/", os.Getenv("API_DOMAIN"), true, true)

	c.Status(http.StatusNoContent)
}
