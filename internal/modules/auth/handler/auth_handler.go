// handler/auth_handler.go
package handler

import (
	"sewapoint/internal/modules/auth/dto"
	"sewapoint/internal/modules/auth/usecase"
	"sewapoint/internal/utils"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Usecase *usecase.Usecase
}

func (h *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if !utils.BindAndValidate(c, &req) {
		return
	}

	token, err := h.Usecase.Register(req)
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	utils.Success(c, dto.AuthResponse{Token: token})
}

func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if !utils.BindAndValidate(c, &req) {
		return
	}

	token, err := h.Usecase.Login(req)
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	utils.Success(c, dto.AuthResponse{Token: token})
}

func (h *Handler) Social(c *gin.Context) {
	var req dto.SocialLoginRequest

	if !utils.BindAndValidate(c, &req) {
		return
	}

	token, err := h.Usecase.SocialLogin(req)
	if err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	utils.Success(c, dto.AuthResponse{Token: token})
}
