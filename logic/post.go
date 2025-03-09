package logic

import (
	"mind_share/dao/mysql"
	"mind_share/models"
	"mind_share/pkg/snowflake"
)

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

func GetPostList() (data []*models.PostDetail, err error) {
	posts, err := mysql.GetPostList()
	if err != nil {
		return nil, err
	}
	data = make([]*models.PostDetail, 0, len(posts))

	for _, post := range posts {
		communityDetail, err := mysql.GetCommunityDetail(post.CommunityID)
		if err != nil {
			continue
		}
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			continue
		}
		postDetail := new(models.PostDetail)
		postDetail.AuthorName = user.Username
		postDetail.Post = post
		postDetail.CommunityDetail = communityDetail
		data = append(data, postDetail)
	}

	return

}
