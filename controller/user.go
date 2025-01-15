package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": errs.Translate(trans), // 翻译错误
		})
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
