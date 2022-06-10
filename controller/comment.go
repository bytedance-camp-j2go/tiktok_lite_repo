package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"tiktok-lite/dao"
	"tiktok-lite/model"
	"tiktok-lite/response"
	"tiktok-lite/util"
)

// CommentAction 评论操作
// action_type 	1-新增 2-删除
// comment_text 评论内容
// comment_id 	评论的标识
// video_id		视频 id
func CommentAction(ctx *gin.Context) {
	user := CtxUser(ctx)

	queryStr := ctx.Query("action_type")
	if len(queryStr) != 1 {
		CtxInputError(ctx, "error action_type:"+queryStr)
		return
	}

	switch []byte(queryStr)[0] - '0' {
	case ActionAppend:
		commentAppend(ctx, user)
	case ActionDel:
		commentDel(ctx, user)
	}
}

func commentAppend(ctx *gin.Context, user *model.User) {
	var (
		videoId  int64
		err      error
		queryStr string
	)
	queryStr = ctx.Query("video_id")
	if videoId, err = util.String10Bit2Int64(queryStr); err != nil {
		CtxInputError(ctx, "error video_id! "+err.Error())
		return
	}

	queryStr = ctx.Query("comment_text")
	if err = checkCommentText(queryStr); err != nil {
		CtxInputError(ctx, "error comment_text:"+err.Error())
		return
	}

	comment := model.NewComment(util.UniqueID())
	{
		comment.UsrID = user.Id
		comment.Content = queryStr
		comment.VideoID = videoId
	}

	err = dao.CommentSave(comment)
	if err != nil {
		CtxServerError(ctx, "error server process!!")
		zap.L().Error("save 2 db err!", zap.Error(err))
		return
	}
	// CtxBaseSuccess(ctx, "opt success")
	ctx.JSON(http.StatusOK,
		response.NewCommentResp(*comment, user.Id),
	)
}

func commentDel(ctx *gin.Context, user *model.User) {
	var (
		err      error
		queryStr string
	)

	queryStr = ctx.Query("comment_id")
	commentId, err := util.String10Bit2Int64(queryStr)
	if err != nil {
		CtxInputError(ctx, "error comment_id:"+err.Error())
		return
	}

	cUid, err := dao.CommentQueryUserId(commentId)
	checkDel := checkCommentDel(user.Id, cUid)
	if err != nil || !checkDel {
		CtxInputError(ctx, "error comment del illegal !")
		zap.L().Debug("comment del error!!",
			zap.Bool("can_do_del", checkDel),
			zap.Error(err),
			zap.Int64("owner_uid", cUid),
			zap.Int64("opter_uid", user.Id),
		)
		return
	}

	if !checkDel {
		CtxInputError(ctx, "You can't del this comment! No permission")
		return
	}

	err = dao.CommentDel(&model.Comment{ID: commentId})
	if err != nil {
		CtxServerError(ctx, "")
		return
	}

}

// TODO 权限分级
// 由于目前 User 对象中不需要实现权限分级，所以使用 uid 简单判断
func checkCommentDel(uid, cuid int64) bool {
	// admin 账号保留前 10000
	return uid == cuid || uid < 10000
}

// TODO 完善 SQL 注入防护
// 执行一些防止 SQL 注入的查询
func checkCommentText(text string) error {
	if text != "" {
		return nil
	}

	// ... 定义一些文本判断的报错信息
	return nil
}

// CommentList 视频下的评论列表
// video_id
func CommentList(ctx *gin.Context) {
	user := *CtxUser(ctx)
	queryStr := ctx.Query("video_id")
	videoId, err := util.String10Bit2Int64(queryStr)
	if err != nil {
		CtxInputError(ctx, "error video_id:"+queryStr)
		return
	}

	commentList, err := dao.CommentList(videoId)
	if err != nil {
		CtxServerError(ctx, "")
		zap.L().Debug("query comment list err!",
			zap.Int64("vid", videoId),
			zap.Error(err),
		)
		return
	}

	ctx.JSON(http.StatusOK, response.NewCommentListResp(commentList, user.Id))
}
