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

package rpcbench

//go:generate protoc --go_out=plugins=grpc:. types.proto
//go:generate hrpc-gen types.proto

import (
	"context"
	"time"
)

type ReplyStatus int

const (
	ReplyOK ReplyStatus = 1
	ReplyER ReplyStatus = 2

	NetMsgSizeMax = 8 * 1024 * 1024

	CallWriteSizeMin int64 = 10
	CallWriteSizeMax int64 = 4 * 1024 * 1024
	CallReadSizeMin  int64 = 10
	CallReadSizeMax  int64 = 4 * 1024 * 1024
	CallWorkTimeMax  int64 = 3 * 1e6 // us

	CallTimeout = time.Second * 10
)

type Config struct {
	Name       string
	ServerPort int
	Time       int64 // bench time in seconds
	ClientNum  int
	WriteSize  int64
	ReadSize   int64
	WorkTime   int64 // a random time to simulate "work" in µs
	loadAvg    float64
}

type BenchProjectInterface interface {
	ServerStart(cfg *Config) error
	ClientConn(opts *ClientConnOptions) (ClientConnInterface, error)
}

type Bench struct {
	cfg    *Config
	proj   BenchProjectInterface
	srv    BenchmarkServer
	quit   bool
	status *BenchStatus
}

type ClientConnOptions struct {
	Addr    string
	ConnNum int
}

type ClientConnInterface interface {
	UnaryCall(ctx context.Context, req *BenchRequest) (*BenchReply, error)
}
