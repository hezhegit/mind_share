package logic

import (
	"mind_share/dao/mysql"
	"mind_share/models"
	"mind_share/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断该用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2. 生成ID
	userID := snowflake.GenID()

	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 4. 数据库操作
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return err
	}

	return nil
}
