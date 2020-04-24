package rpc

import (
	"context"
	"goframework/process/controller"
	"goframework/process/rpc/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

//rpc 服务端

type Server struct {
}

//启动 rpc 服务
func StartRpcServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {

	}
	s := grpc.NewServer()
	server.RegisterServerServer(s, &Server{})
	if err := s.Serve(lis); err != nil {

	}
	log.Printf("Listening and serving HTTP on " + addr)
}

func (rp *Server) GetServerTime(ctx context.Context, request *server.ServerTimeRequest) (*server.ServerTimeResponse, error) {
	data, code, msg := controller.GetServerTime()
	resp := &server.ServerTimeResponse{}
	respData := &server.ServerTimeResponseData{}
	resp.Msg = msg
	resp.Code = uint32(code)
	respData.ServerTime = uint64(data.ServerTime)
	resp.Data = respData
	return resp, nil
}
