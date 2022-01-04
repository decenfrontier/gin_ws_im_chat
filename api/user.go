package api

import (
	"chat/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserRegisterService
	err := c.ShouldBind(&userRegisterService)
	if err != nil {
		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
	ret := userRegisterService.Register()
	c.JSON(200, ret)
}
