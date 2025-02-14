package logic

import (
	"mind_share/dao/mysql"
	"mind_share/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}
