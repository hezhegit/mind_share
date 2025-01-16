package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"mind_share/models"
)

const secret = "mind-share"

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(1) from user where username=?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}

	if count > 0 {
		return errors.New("用户已存在！")
	}

	return nil
}

func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user (user_id, username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
