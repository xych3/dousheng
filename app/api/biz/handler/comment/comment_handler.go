// Code generated by hertz generator.

package comment

import (
	"context"
	"fmt"
	"log"

	"gitee.com/derrickball/douyin/app/api/biz/handler/pack"
	"gitee.com/derrickball/douyin/app/comment/kitex_gen/comment_rpc"

	"gitee.com/derrickball/douyin/app/api/biz/model/comment"
	"gitee.com/derrickball/douyin/app/api/biz/mw"
	"gitee.com/derrickball/douyin/app/api/biz/rpc"
	"gitee.com/derrickball/douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Action .
// @router /douyin/comment/action/ [POST]
func Action(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.CommentActionReq
	var response *comment_rpc.RPCActionResp
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//resp := new(comment.CommentActionResp)

	if len(req.GetToken()) == 0 {
		c.JSON(consts.StatusOK, errno.MConvertErr(errno.AuthorizationFailedErr))
		return
	}
	if len(fmt.Sprint(req.GetVideoId())) == 0 || len(fmt.Sprint(req.GetActionType())) == 0 {
		c.JSON(consts.StatusOK, errno.MConvertErr(errno.ParamErr))
		return
	}
	uid := c.GetInt64(mw.AuthMiddleware.IdentityKey)
	switch req.GetActionType() {
	case comment.CommentActionType_PUBLISH:
		response, err = rpc.CreateComment(ctx, &comment_rpc.RPCCommentCreateReq{
			UserId:      uid,
			VideoId:     req.GetVideoId(),
			CommentText: req.GetCommentText(),
		})
		if err != nil {
			c.JSON(consts.StatusBadRequest, err.Error())
			return
		}

	case comment.CommentActionType_DELETE:
		response, err = rpc.DelComment(ctx, &comment_rpc.RPCCommentDelReq{
			CommentId: req.GetCommentId(),
			UserId:    uid,
			VideoId:   req.GetVideoId(),
		})
		if err != nil {
			c.JSON(consts.StatusBadRequest, err.Error())
			return
		}

	default:
		c.JSON(consts.StatusOK, errno.MConvertErr(errno.ParamErr))
		return
	}
	resp := &comment.CommentActionResp{
		StatusCode: &errno.Success.ErrCode,
		StatusMsg:  &errno.Success.ErrMsg,
		Comment:    pack.Comment(response.GetComment()),
	}
	c.JSON(consts.StatusOK, resp)
}

// List .
// @router /douyin/comment/list/ [GET]
func List(ctx context.Context, c *app.RequestContext) {
	var err error
	var req comment.CommentListReq
	var userId int64
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userId = -1
	if len(req.GetToken()) != 0 {
		userId = c.GetInt64(mw.AuthMiddleware.IdentityKey)
	}
	response, err := rpc.CommentList(ctx, &comment_rpc.RPCCommentListReq{
		VideoId: req.GetVideoId(),
		UserId:  userId,
	})
	if err != nil {
		c.JSON(consts.StatusOK, errno.MConvertErr(err))
		return
	}
	log.Println(response.CommentList)

	resp := &comment.CommentListResp{
		StatusCode:  &errno.Success.ErrCode,
		StatusMsg:   &errno.Success.ErrMsg,
		CommentList: pack.Comments(response.CommentList),
	}
	c.JSON(consts.StatusOK, resp)
}