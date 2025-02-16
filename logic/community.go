package logic

import (
	"mind_share/dao/mysql"
	"mind_share/models"
	"mind_share/pkg/snowflake"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetail(id)
}

func CreatePost(p *models.Post) error {
	// 1. 参数处理 id生成等
	p.ID = snowflake.GenID()
	// 2. 保存到数据库
	return mysql.CreatePost(p)

}

func SelectPostByID(id int64) (data *models.PostDetail, err error) {
	post, err := mysql.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	communityDetail, err := mysql.GetCommunityDetail(post.CommunityID)
	if err != nil {
		return nil, err
	}
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		return nil, err
	}
	data = new(models.PostDetail)
	data.AuthorName = user.Username
	data.Post = post
	data.CommunityDetail = communityDetail
	return

}
