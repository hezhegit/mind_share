package mysql

import (
	"database/sql"
	"errors"
	"mind_share/models"
)

func CreatePost(p *models.Post) error {
	sqlStr := "insert into post (post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)"

	_, err := db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return err
}

func GetPostByID(id int64) (p *models.Post, err error) {
	p = new(models.Post)
	sqlStr := "select post_id, title, content, author_id, community_id, create_time from post where post_id = ?"
	if err = db.Get(p, sqlStr, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrInvalidID
		}
	}
	return
}

func GetPostList(pageNum, pageSize int64) (p []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post limit ?,?`
	p = make([]*models.Post, 0, 2) // 长度-容量
	err = db.Select(&p, sqlStr, (pageNum-1)*pageSize, pageSize)
	return

}
