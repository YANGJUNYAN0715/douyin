package rpc

import (
	"context"
	// "log"

	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/relation"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/relation/relationservice"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelation() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := relationservice.NewClient(
		consts.RelationServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

// 传递 关注操作 的上下文, 并获取 RPC Server 端的响应.
func RelationAction(ctx context.Context, req *relation.RelationActionRequest) (error) {
	resp, err := relationClient.RelationAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

// 传递 获取正在关注列表操作 的上下文, 并获取 RPC Server 端的响应.
func RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) ([]*relation.User,error) {
	resp, err := relationClient.RelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil,errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	// log.Println("***api-rpc-relation.go***")
	// log.Println(resp.UserList)
	return resp.UserList, nil
}

// 传递 获取粉丝列表操作 的上下文, 并获取 RPC Server 端的响应.
func RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) ([]*relation.User,error) {
	resp, err := relationClient.RelationFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil,errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.UserList, nil
}

// RelationFriendList
func RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) ([]*relation.FriendUser, error) {
	resp, err := relationClient.RelationFriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return 0, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
		return nil,errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.UserList, nil
}



func MessageChat(ctx context.Context, req *relation.MessageChatRequest) ([]*relation.Message,error) {
	resp, err := relationClient.MessageChat(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil,errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.Messages, nil
}


func MessageAction(ctx context.Context, req *relation.MessageActionRequest) error {
	resp, err := relationClient.MessageAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}