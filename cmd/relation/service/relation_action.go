
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"

	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/relation/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/relation"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
)

type RelationActionService struct {
	ctx context.Context
}

// NewRelationActionService new RelationActionService
func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// Register create user info.
func (s *RelationActionService) RelationAction(req *relation.RelationActionRequest) error {
	//新建关注
	if req.ActionType==1{
		return db.NewAction(s.ctx,req.UserId,req.ToUserId);
	}
	//删除关注
	if req.ActionType==2{
		return db.DelAction(s.ctx,req.UserId,req.ToUserId);
	}

	return errno.ActionTypeErr
}
