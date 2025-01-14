package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mind_share/logic"
	"mind_share/models"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误
		zap.L().Error("SignUpHandler ShouldBindJSON error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数有误！"})
		return
	}
	// 2. 业务处理
	logic.SignUp(p)
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
