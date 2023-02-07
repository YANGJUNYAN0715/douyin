package main

import (
	"log"
	"net"

	user "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/userservice"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	// "github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/rpcinfo"
	server "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}), // server name
		server.WithServiceAddr(addr), server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}