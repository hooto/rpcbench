// Copyright 2020 Eryx <evorui аt gmаil dοt cοm>, All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	rpcbench "github.com/hooto/rpcbench/v1"
)

func main() {

	b := rpcbench.NewBench(new(bench), new(benchServiceServer))

	if err := b.Run(); err != nil {
		panic(err)
	}
}

type benchServiceServer struct{}

func (s benchServiceServer) UnaryCall(
	ctx context.Context,
	req *rpcbench.BenchRequest,
) (*rpcbench.BenchReply, error) {

	if req.ReplySpec == nil {
		return nil, errors.New("ReplySpec not found")
	}

	if req.ReplySpec.WorkTime > 0 {
		time.Sleep(time.Microsecond * time.Duration(req.ReplySpec.WorkTime))
	}

	return &rpcbench.BenchReply{
		Id: req.Id,
		Payload: &rpcbench.BenchPayload{
			Body: rpcbench.RandValue(int(req.ReplySpec.Size)),
		},
	}, nil
}

type bench struct {
	server        *grpc.Server
	serviceServer *benchServiceServer
}

type benchClientConn struct {
	conn *grpc.ClientConn
}

func (it *bench) ServerStart(cfg *rpcbench.Config) error {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.ServerPort))
	if err != nil {
		return err
	}

	it.server = grpc.NewServer(
		grpc.MaxMsgSize(rpcbench.NetMsgSizeMax),
		grpc.MaxSendMsgSize(rpcbench.NetMsgSizeMax),
		grpc.MaxRecvMsgSize(rpcbench.NetMsgSizeMax),
	)

	go it.server.Serve(lis)

	rpcbench.RegisterBenchmarkServer(it.server, new(benchServiceServer))

	return nil
}

func (it *bench) ClientConn(opts *rpcbench.ClientConnOptions) (rpcbench.ClientConnInterface, error) {

	c, err := grpc.Dial(opts.Addr, grpc.WithInsecure(),
		// grpc.WithPerRPCCredentials(auth.NewCredentialToken()),
		grpc.WithMaxMsgSize(rpcbench.NetMsgSizeMax),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(rpcbench.NetMsgSizeMax)),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(rpcbench.NetMsgSizeMax)),
	)
	if err != nil {
		return nil, err
	}

	return &benchClientConn{
		conn: c,
	}, nil
}

func (it *benchClientConn) UnaryCall(
	ctx context.Context,
	req *rpcbench.BenchRequest) (*rpcbench.BenchReply, error) {

	return rpcbench.NewBenchmarkClient(it.conn).UnaryCall(ctx, req)
}
