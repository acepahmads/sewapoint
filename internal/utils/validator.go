package utils

import (
	"github.com/gin-gonic/gin"
)

// Cara pakai
// var req CreateBookingRequest
// if !utils.BindAndValidate(c, &req) {
//     return
// }

// import "github.com/go-playground/validator/v10"
// type CreateBookingRequest struct {
// 	ProductID int `json:"product_id" validate:"required"`
// }

func BindAndValidate(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		Error(c, 400, err.Error())
		return false
	}
	return true
}
