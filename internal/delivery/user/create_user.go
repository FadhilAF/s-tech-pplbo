package delivery

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/utils"

	"github.com/FadhilAF/s-tech-pplbo/internal/model"
	"github.com/gin-gonic/gin"
)

func (handler *userHandler) CreateUser(ctx *gin.Context) {
	var req model.CreateUserRequest

	ok := utils.BindJSONAndValidate(ctx, &req)
	if !ok {
		return
	}

	res := handler.usecase.CreateUser(req)

	ctx.JSON(res.Status, res)
}
