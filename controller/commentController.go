package controller

import (
	"github.com/gin-gonic/gin"
)

// CommentAction 评论操作
// action_type 1-新增 2-删除
// comment_text 评论内容
// comment_id 评论的标识
//
func CommentAction(ctx *gin.Context) {
	// user := *CtxUser(ctx)

	queryStr := ctx.Query("action_type")
	if len(queryStr) != 1 {
		CtxInputError(ctx, "error action_type:"+queryStr)
		return
	}

	switch []byte(queryStr)[0] - '0' {
	case ActionAppend:
		queryStr = ctx.Query("comment_text")
		if err := checkCommentText(queryStr); err != nil {
			CtxInputError(ctx, "error comment_text:"+err.Error())
			return
		}

		// TODO 写入数据库
	case ActionDel:
		// TODO 执行删除
		queryStr = ctx.Query("comment_id")
		// commentId, err := util.String10Bit2Int64(queryStr)
		// if err != nil {
		// 	CtxInputError(ctx, "error comment_id:"+err.Error())
		// 	return
		// }
		// comment := dao.
	}
}

// 执行一些防止 SQL 注入的查询
func checkCommentText(text string) error {
	if text != "" {
		return nil
	}

	// ... 定义一些文本判断的报错信息
	return nil
}

func CommentList(ctx *gin.Context) {

}
