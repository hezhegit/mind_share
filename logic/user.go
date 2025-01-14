package logic

import (
	"mind_share/dao/mysql"
	"mind_share/models"
	"mind_share/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// 1. 判断该用户存不存在
	mysql.QueryUserByUsername()
	// 2. 生成ID
	snowflake.GenID()
	// 3. 密码加密

	// 4. 数据库操作
	mysql.InsertUser()
}
