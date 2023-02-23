// Code generated by hertz generator.

package api

import (

	"context"
	// "fmt"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/model/api"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/mw"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/rpc"
	// "log"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/interact"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/message"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	// "path/filepath"
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	if req.ActionType != 1{
		SendResponse(c, errno.Success, nil)
		return
	}
	err = rpc.FavoriteAction(context.Background(), &interact.FavoriteActionRequest{
		UserId:  v.(*api.User).ID,
		Token: req.Token,
		VideoId: req.VideoID,
		ActionType: req.ActionType,
		

	})
	
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	
	SendResponse(c, errno.Success, nil)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteListRequest
	err = c.BindAndValidate(&req)
	
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	video_list, err := rpc.FavoriteList(context.Background(), &interact.FavoriteListRequest{
		UserId: v.(*api.User).ID,
		Token: req.Token,
		

	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// SendResponse(c, errno.Success, nil)
	Err := errno.ConvertErr(errno.Success)
	c.JSON(200, utils.H{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"video_list":  video_list,
	})
}


// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	// var err error
	// var req api.CommentActionRequest
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	SendResponse(c, errno.ConvertErr(err), nil)
	// 	return
	// }
	// //从token中获取id TODO middleware中处理
	// u, _ := c.Get(consts.IdentityKey)
	// if u == nil {
	// 	SendResponse(c, errno.Token2UserIdErr, nil)
	// 	return
	// }
	// //越权错误
	// if req.UserID != 0 && req.UserID != u.(*api.User).ID {
	// 	SendResponse(c, errno.BrokenAccessControlErr, nil)
	// 	return
	// }
	// resp, err := rpc.CommentAction(ctx, &interact.CommentActionRequest{
	// 	UserId:      u.(*api.User).ID,
	// 	Token:       req.Token,
	// 	VideoId:     req.VideoID,
	// 	ActionType:  req.ActionType,
	// 	CommentText: req.CommentText,
	// 	CommentId:   req.CommentID,
	// })
	// Err := errno.ConvertErr(errno.Success)
	// c.JSON(consts.StatusOK, utils.H{
	// 	"status_code": Err.ErrCode,
	// 	"status_msg":  Err.ErrMsg,
	// 	"comment":     resp,
	// })
}

// CommentList .
// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	// var err error
	// var req api.CommentListRequest
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	SendResponse(c, errno.ConvertErr(err), nil)
	// 	return
	// }
	// //从token中获取id TODO middleware中处理
	// u, _ := c.Get(consts.IdentityKey)
	// if u == nil {
	// 	SendResponse(c, errno.Token2UserIdErr, nil)
	// 	return
	// }
	// resp, err := rpc.CommentList(ctx, &interact.CommentListRequest{
	// 	Token:   req.Token,
	// 	VideoId: req.VideoID,
	// })
	// Err := errno.ConvertErr(errno.Success)
	// c.JSON(consts.StatusOK, utils.H{
	// 	"status_code":  Err.ErrCode,
	// 	"status_msg":   Err.ErrMsg,
	// 	"comment_list": resp,
	// })
}

