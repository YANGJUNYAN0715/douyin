// Code generated by hertz generator.

package api

import (
	"context"
	"log"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/model/api"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/mw"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/rpc"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/comment"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/relation"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)
// RelationAction .
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//从token中获取id
	u, _ := c.Get(consts.IdentityKey)
	if u == nil {
		SendResponse(c, errno.Token2UserIdErr, nil)
		return
	}
	log.Println(u)
	// log.Println("userid:%d",u.(*api.User).ID)
	err = rpc.RelationAction(context.Background(), &relation.RelationActionRequest{
		UserId:     u.(*api.User).ID,
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, nil)
}
// RelationFollowList .
// @router /douyin/relation/follow/list/ [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	//从token中获取id
	u, _ := c.Get(consts.IdentityKey)
	if u == nil {
		SendResponse(c, errno.Token2UserIdErr, nil)
		return
	}
	//越权错误
	if req.UserID != 0 && req.UserID != u.(*api.User).ID {
		SendResponse(c, errno.BrokenAccessControlErr, nil)
		return
	}
	if req.UserID == 0 {
		req.UserID = u.(*api.User).ID
	}

	// resp := new(api.RelationFollowListResponse)
	users, err := rpc.RelationFollowList(context.Background(), &relation.RelationFollowListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// log.Println("***api.go***")
	// log.Println(users)
	// SendResponse(c, errno.Success, utils.H{
	// 	"user_list": users,
	// })
	// 客户端对JSON的名称有要求，不适用SendResponse
	Err := errno.ConvertErr(errno.Success)
	c.JSON(200, utils.H{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"user_list":   users,
	})
}

// RelationFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//从token中获取id
	u, _ := c.Get(consts.IdentityKey)
	if u == nil {
		SendResponse(c, errno.Token2UserIdErr, nil)
		return
	}
	//越权错误
	if req.UserID != 0 && req.UserID != u.(*api.User).ID {
		SendResponse(c, errno.BrokenAccessControlErr, nil)
		return
	}
	if req.UserID == 0 {
		req.UserID = u.(*api.User).ID
	}
	// resp := new(api.RelationFollowerListResponse)
	users, err := rpc.RelationFollowerList(context.Background(), &relation.RelationFollowerListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	// SendResponse(c, errno.Success, utils.H{
	// 	"user_list": users,
	// })
	// 客户端对JSON的名称有要求，不适用SendResponse
	Err := errno.ConvertErr(errno.Success)
	c.JSON(200, utils.H{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"user_list":   users,
	})
}

// RelationFriendList .
// @router /douyin/relation/friend/list/ [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//从token中获取id
	u, _ := c.Get(consts.IdentityKey)
	if u == nil {
		SendResponse(c, errno.Token2UserIdErr, nil)
		return
	}
	//越权错误
	if req.UserID != 0 && req.UserID != u.(*api.User).ID {
		SendResponse(c, errno.BrokenAccessControlErr, nil)
		return
	}
	if req.UserID == 0 {
		req.UserID = u.(*api.User).ID
	}
	friends, err := rpc.RelationFriendList(context.Background(), &relation.RelationFriendListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})

	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	// SendResponse(c, errno.Success, utils.H{
	// 	"user_list": friends,
	// })
	// 客户端对JSON的名称有要求，不适用SendResponse
	Err := errno.ConvertErr(errno.Success)
	c.JSON(200, utils.H{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"user_list":   friends,
	})
}


// // CommentAction .
// // @router /douyin/comment/action/ [POST]
// func CommentAction(ctx context.Context, c *app.RequestContext) {
// 	var err error
// 	var req api.CommentActionRequest
// 	err = c.BindAndValidate(&req)
// 	if err != nil {
// 		SendResponse(c, errno.ConvertErr(err), nil)
// 		return
// 	}
// 	//从token中获取id TODO middleware中处理
// 	u, _ := c.Get(consts.IdentityKey)
// 	if u == nil {
// 		SendResponse(c, errno.Token2UserIdErr, nil)
// 		return
// 	}
// 	//越权错误
// 	if req.UserID != 0 && req.UserID != u.(*api.User).ID {
// 		SendResponse(c, errno.BrokenAccessControlErr, nil)
// 		return
// 	}
// 	resp, err := rpc.CommentAction(ctx, &comment.CommentActionRequest{
// 		UserId:      u.(*api.User).ID,
// 		Token:       req.Token,
// 		VideoId:     req.VideoID,
// 		ActionType:  req.ActionType,
// 		CommentText: req.CommentText,
// 		CommentId:   req.CommentID,
// 	})
// 	Err := errno.ConvertErr(errno.Success)
// 	c.JSON(200, utils.H{
// 		"status_code": Err.ErrCode,
// 		"status_msg":  Err.ErrMsg,
// 		"comment":     resp,
// 	})
// }

// // CommentList .
// // @router /douyin/comment/list/ [GET]
// func CommentList(ctx context.Context, c *app.RequestContext) {
// 	var err error
// 	var req api.CommentListRequest
// 	err = c.BindAndValidate(&req)
// 	if err != nil {
// 		SendResponse(c, errno.ConvertErr(err), nil)
// 		return
// 	}
// 	//从token中获取id TODO middleware中处理
// 	u, _ := c.Get(consts.IdentityKey)
// 	if u == nil {
// 		SendResponse(c, errno.Token2UserIdErr, nil)
// 		return
// 	}
// 	resp, err := rpc.CommentList(ctx, &comment.CommentListRequest{
// 		Token:   req.Token,
// 		VideoId: req.VideoID,
// 	})
// 	Err := errno.ConvertErr(errno.Success)
// 	c.JSON(200, utils.H{
// 		"status_code":  Err.ErrCode,
// 		"status_msg":   Err.ErrMsg,
// 		"comment_list": resp,
// 	})
// }
// MessageChat .
// @router /douyin/message/chat/ [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	var err error
	var req api.MessageChatRequest
	err = c.BindAndValidate(&req)
	
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	messages, err := rpc.MessageChat(context.Background(), &relation.MessageChatRequest{
		FromUserId: v.(*api.User).ID,
		Token: req.Token,
		ToUserId: req.ToUserID,
<<<<<<< HEAD

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
		"message_list":  messages,
	})
=======
	// var err error
	// var req api.MessageChatRequest
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	c.String(consts.StatusBadRequest, err.Error())
	// 	return
	// }

	// resp := new(api.MessageChatResponse)

	// c.JSON(200, resp)
>>>>>>> origin/guo
=======

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
		"message_list":  messages,
	})
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
	var err error
	var req api.MessageActionRequest
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
	err = rpc.MessageAction(context.Background(), &relation.MessageActionRequest{
		FromUserId:  v.(*api.User).ID,
		Token: req.Token,
		ToUserId: req.ToUserID,
		ActionType: req.ActionType,
		Content: req.Content,
<<<<<<< HEAD

	})
	log.Println("hz----------------------------", req.ToUserID)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// log.Println("hz----------------------------", req.ToUserID)
	SendResponse(c, errno.Success, nil)
=======
	// var err error
	// var req api.MessageActionRequest
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	c.String(consts.StatusBadRequest, err.Error())
	// 	return
	// }

	// resp := new(api.MessageActionResponse)

	// c.JSON(200, resp)
>>>>>>> origin/guo
=======

	})
	log.Println("hz----------------------------", req.ToUserID)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// log.Println("hz----------------------------", req.ToUserID)
	SendResponse(c, errno.Success, nil)
>>>>>>> 2f592bb30236c8349ec8e629984207ec905ef48a
}
