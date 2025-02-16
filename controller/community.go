package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mind_share/logic"
	"mind_share/models"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id, community_name） 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 查询该id的社区（community_id, community_name） 以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}

func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数的校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Debug("ShouldBindJSON failed", zap.Any("err", err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 	取得当前发请求的用户ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	p.AuthorID = userID

	// 2. 创建帖子 logic
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
}

func SelectPostByIDHandler(c *gin.Context) {
	// 1. 获取参数
	idStr := c.Param("id")
	postID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 查询该id的帖子
	p, err := logic.SelectPostByID(postID)
	if err != nil {
		zap.L().Error("logic.SelectPostByID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, p)
}
