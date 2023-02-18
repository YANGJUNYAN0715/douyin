package service

import (
	"context"
	"errors"
	"github.com/YANGJUNYAN0715/douyin/tree/li/cmd/favorite/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/li/kitex_gen/favorite"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/jwt"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction implement the like and unlike operations
func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionRequest) error {
	Jwt := jwt.NewJWT([]byte(consts.SecretKey))
	claim, err := Jwt.ParseToken(req.Token)
	if err != nil {
		return err
	}
	currentId := claim.Id

	videos, err := db.QueryVideoByVideoIds(s.ctx, []int64{req.VideoId})
	if err != nil {
		return err
	}
	if len(videos) == 0 {
		return errors.New("video not exist")
	}

	//若ActionType（操作类型）等于1，则向favorite表创建一条记录，同时向video表的目标video增加点赞数
	//若ActionType等于2，则向favorite表删除一条记录，同时向video表的目标video减少点赞数
	//若ActionType不等于1和2，则返回错误
	if req.ActionType == consts.Like {
		favorite := &db.Favorite{
			UserId:  currentId,
			VideoId: req.VideoId,
		}

		err := db.CreateFavorite(s.ctx, favorite, req.VideoId)
		if err != nil {
			return err
		}
	}
	if req.ActionType == consts.Unlike {
		err := db.DeleteFavorite(s.ctx, currentId, req.VideoId)
		if err != nil {
			return err
		}

	}
	if req.ActionType != consts.Like && req.ActionType != consts.Unlike {
		return errors.New("action type no equal 1 and 2")
	}
	return nil
}
