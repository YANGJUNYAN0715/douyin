// Code generated by hertz generator.

package api

import (
	"context"
	"fmt"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/api/biz/model/api"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/api/biz/mw"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/api/biz/rpc"
	"log"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/user"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/message"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"path/filepath"
	"time"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	// "mime/multipart"
)

// LoginUser .
// @router /douyin/user/login/ [POST]
func LoginUser(ctx context.Context, c *app.RequestContext) {
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// RegisterUser .
// @router /douyin/user/register/ [POST]
func RegisterUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegisterUserRequest
	err = c.BindAndValidate(&req)
	
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	err = rpc.RegisterUser(context.Background(), &user.RegisterUserRequest{
		Username: req.Username,
		Password: req.Password,

	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// UserInfo .
// @router /douyin/user/ [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserInfoRequest
	err = c.BindAndValidate(&req)
	
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	user_info, err := rpc.UserInfo(context.Background(), &user.UserInfoRequest{
		UserId: v.(*api.User).ID,
		Token: req.Token,

	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	log.Println("***------------------------------------UserInfo-service---------------------------------------***")
	log.Println(user_info)
	Err := errno.ConvertErr(errno.Success)
	c.JSON(200, utils.H{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"user":   user_info,
	})
}

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	log.Println("***------------------------------------PublishAction-service---------------------------------------***")
	var err error
	// var req multipart.FileHeader

	// _ = c.ParseMultipartForm(1024)
	// var req api.PublishActionRequest
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	SendResponse(c, errno.ConvertErr(err), nil)
	// 	return
	// }
	if(c.PostForm("token")==""){
		SendResponse(c, errno.ConvertErr(err), nil)
		return 
	}
	video_data, err := c.FormFile("data")
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	title := c.PostForm("title")
	if title == "" {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	u, _ := c.Get(consts.IdentityKey)
	
	
	filename := filepath.Base(video_data.Filename)
	finalName := fmt.Sprintf("%s", filename)
	video_path := fmt.Sprintf("%s", filepath.Join(consts.VideoSavePath, finalName))
	// video_path := filepath.Join(consts.VideoSavePath, finalName)
	err = c.SaveUploadedFile(video_data, video_path)
	log.Println("2///////////////////",err,"//////////////////////////")
	if err != nil {
		
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	log.Println("1///////////////////",video_path,"//////////////////////////")
	
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	endpoint := "https://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	accessKeyId := "LTAI5tRapYeKEkQYL2QD1xnc"
	accessKeySecret := "t1EiOyPLbi67pCANcRbIy3r8zyCzMP"
	// yourBucketName填写存储空间名称。
	bucketName := "douyin-test-guo"
	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	VideoName := finalName
	ImgName := "img"
	// yourLocalFileName填写本地文件的完整路径。
	localFileName_video := video_path
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	// 上传文件。
	
	err = bucket.PutObjectFromFile("video/" + VideoName, localFileName_video)
	if err != nil {
		log.Println("上传失败",err)
		SendResponse(c, errno.ConvertErr(err), nil)
	}

	
	log.Println("3/////////////////////////////////////////////")
	
	// 获取视频截图
	snapshotName, err := GetSnapshot(video_path, consts.CoverPath, 1)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	cover_path := fmt.Sprintf("%s.jpg", filepath.Join(consts.CoverPath, snapshotName))
	
	
	// 上传文件。
	err = bucket.PutObjectFromFile(ImgName, cover_path)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}

	
	log.Println("4///////////////////",snapshotName,"//////////////////////////")
	
	err = rpc.PublishAction(context.Background(), &user.PublishActionRequest{
		UserId:  u.(*api.User).ID,
		// Token: c.PostForm("token"),
		Title: title,
		
		FileUrl: video_path,
		CoverUrl: cover_path,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	log.Println(fmt.Sprintf("%s.jpg", filepath.Join(consts.CoverPath, snapshotName)),video_path)
	SendResponse(c, errno.Success, nil)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	videos, err := rpc.PublishList(context.Background(), &user.PublishListRequest{
		UserId: v.(*api.User).ID,
		Token: req.Token,

	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	Err := errno.ConvertErr(errno.Success)
	c.JSON(200, utils.H{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"video_list":   videos,
	})
}

// GetUserFeed .
// @router /douyin/feed/ [GET]
func GetUserFeed(ctx context.Context, c *app.RequestContext) {
	log.Println("-------------------------------------------------hertz feed----------------------------------------------------")
	var err error
	var req api.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//从token中获取id
	u, _ := c.Get(consts.IdentityKey)
	
	// if u == nil {
	// 	SendResponse(c, errno.Token2UserIdErr, nil)
	// 	return
	// }
	//越权错误
	// if req.UserID != 0 && req.UserID != u.(*api.User).ID {
	// 	SendResponse(c, errno.BrokenAccessControlErr, nil)
	// 	return
	// }
	// if req.UserID == 0 {
	// 	req.UserID = u.(*api.User).ID
	// }
	var user_id int64
	var token string
	var latest_time int64
	user_id = 0
	log.Println("-------token: ",req.Token,"-----")
	log.Println("---------------u: ",u,"--------------------------")
	if req.Token !=""{
		user_id = u.(*api.User).ID
		token = req.Token
	}
	if req.LatestTime == 0{
		latest_time = time.Now().UnixMilli()

	}else {
		latest_time = req.LatestTime
		
	}
	log.Println("-------req.UserID: -----")
	log.Println(user_id, "--------------------------", token, "--------------------------", latest_time)
	
	feed, err := rpc.GetUserFeed(ctx,&user.FeedRequest{
		UserId: user_id,
		LatestTime: latest_time,
		Token:  req.Token,
	})
	log.Println(feed, "---+++++++++++++++++++++++++++++++++++++++----feed--+++++++++++++++++++++++++---")
	//SendResponse2(c, feedresponse)
	Err := errno.ConvertErr(errno.Success)
	c.JSON(200, utils.H{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"video_list": feed.VideoList,
		"next_time": feed.NextTime,
    	
	})
	
}
