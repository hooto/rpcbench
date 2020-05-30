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

import (
	"errors"

	"github.com/hooto/hflag4g/hflag"
)

func NewBench(p BenchProjectInterface, s BenchmarkServer) *Bench {

	c := &Config{
		Name:       hflag.Value("name").String(),
		ServerPort: hflag.Value("server_port").Int(),
		ClientNum:  hflag.Value("client_num").Int(),
		Time:       hflag.Value("time").Int64(),
		WriteSize:  hflag.Value("write_size").Int64(),
		ReadSize:   hflag.Value("read_size").Int64(),
		WorkTime:   hflag.Value("work_time").Int64(),
		loadAvg:    1.0,
	}

	if c.Name == "" {
		c.Name = "rpc"
	}

	if c.ServerPort < 1 || c.ServerPort > 65535 {
		c.ServerPort = 3301
	}

	if c.ClientNum < 1 {
		c.ClientNum = 1
	} else if c.ClientNum > 2048 {
		c.ClientNum = 2048
	}

	if c.Time < 10 {
		c.Time = 10
	} else if c.Time > 3600 {
		c.Time = 3600
	}

	if c.WriteSize < CallWriteSizeMin {
		c.WriteSize = CallWriteSizeMin
	} else if c.WriteSize > CallWriteSizeMax {
		c.WriteSize = CallWriteSizeMax
	}

	if c.ReadSize < CallReadSizeMin {
		c.ReadSize = CallReadSizeMin
	} else if c.ReadSize > CallReadSizeMax {
		c.ReadSize = CallReadSizeMax
	}

	if c.WorkTime > CallWorkTimeMax {
		c.WorkTime = CallWorkTimeMax
	}

	b := &Bench{
		cfg:  c,
		proj: p,
		srv:  s,
		quit: false,
	}

	b.statusSetup()

	return b
}

func (it *Bench) Run() error {

	if _, ok := hflag.ValueOK("chart-export"); ok {
		return chartOutput()
	}

	if it.proj == nil {
		return errors.New("project not found")
	}

	if it.srv == nil {
		return errors.New("service-server not found")
	}

	// run in server-only mode
	if _, ok := hflag.ValueOK("server-only"); ok {
		if err := it.proj.ServerStart(it.cfg); err != nil {
			return err
		}
		select {}
	}

	// run client-only mode
	if _, ok := hflag.ValueOK("client-only"); ok {
		return it.runClient()
	}

	// run server + client in simultaneous mode
	if err := it.proj.ServerStart(it.cfg); err != nil {
		return err
	}
	return it.runClient()
}
