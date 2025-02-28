package handlers

import (
	"net/http"

	"rbp/internal/models"
	"rbp/internal/service"
	"rbp/pkg/validator"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": ErrorBadRequest})

		return
	}

	if err := validator.Validate(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": ErrorBadRequest})

		return
	}

	if _, err := h.userService.Register(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": ErrorInternalServer})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": SuccessUserRegister})
}

func (h *UserHandler) Login(c *gin.Context) {
	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	token, refreshToken, err := h.userService.Login(c, login.Email, login.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

		return
	}

	c.SetCookie("refresh_token", refreshToken, 60*60*24*7, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"token": token, "code": SuccessUserLogin})
}
