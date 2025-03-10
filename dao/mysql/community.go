package mysql

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"mind_share/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err = db.Select(&communityList, sqlStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 查询到没有结果时，给调用者一个明确的返回
			zap.L().Warn("no community list", zap.String("sql", sqlStr))
			return nil, sql.ErrNoRows // 明确返回错误
		}
		// 如果是其他错误，直接返回错误
		return nil, err
	}

	// 查询成功后，正常返回数据
	return communityList, nil
}

func GetCommunityDetail(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := "select community_id, community_name, introduction, create_time from community where community_id = ?"
	if err = db.Get(communityDetail, sqlStr, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrInvalidID
		}
	}
	return communityDetail, err
}
