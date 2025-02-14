package mysql

import "errors"

var (
	ErrUserExists      = errors.New("用户已存在")
	ErrUserNotExists   = errors.New("用户不存在")
	ErrInvalidPassword = errors.New("用户名或密码错误")
	ErrInvalidID       = errors.New("无效的ID")
)
