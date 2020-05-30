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
	"github.com/hooto/hflag4g/hflag"
)

type BenchStatusCallUsageItem struct {
	time int64
	num  int64
}

type BenchStatusOptions struct {
	timeLen       int64 // seconds
	timeStep      int64 // seconds
	latencyMin    int64 // microseconds
	latencyMax    int64 // microseconds
	latencyRanges []int64
	dataFile      string
	dataName      string
}

type BenchStatus struct {
	options     *BenchStatusOptions
	ok          int64
	err         int64
	nps         float64
	npsMap      []*BenchStatusCallUsageItem
	latencyMap  []*BenchStatusCallUsageItem
	latencyTime int64
}

func (it *Bench) statusSetup() {

	if it.status == nil {
		it.status = &BenchStatus{}
	}

	if it.status.options == nil {
		it.status.options = &BenchStatusOptions{
			latencyMin: 10,    // 10 us
			latencyMax: 100e3, // 100 ms
		}
	}

	it.status.options.dataName = it.cfg.Name
	it.status.options.dataFile = "rpcbench.json"

	if v, ok := hflag.ValueOK("latency_min"); ok {
		if it.status.options.latencyMin = v.Int64(); it.status.options.latencyMin < 1 {
			it.status.options.latencyMin = 1 // 1 us
		} else if it.status.options.latencyMin > 1e6 {
			it.status.options.latencyMin = 1e6 // 1 s
		}
	}

	if v, ok := hflag.ValueOK("latency_max"); ok {
		it.status.options.latencyMax = v.Int64()
	}
	if (it.status.options.latencyMin * 10) > it.status.options.latencyMax {
		it.status.options.latencyMax = it.status.options.latencyMin * 10
	}

	// NPS
	it.status.options.timeLen = it.cfg.Time
	it.status.options.timeStep = int64(1)

	// TC
	it.status.options.latencyRanges = []int64{}
	latencyRange := ((it.status.options.latencyMax - it.status.options.latencyMin) >> 20)
	if latencyRange < it.status.options.latencyMin {
		latencyRange = it.status.options.latencyMin
	}
	for i := 0; i < 20; i++ {

		if latencyRange > it.status.options.latencyMax {
			it.status.options.latencyRanges = append(it.status.options.latencyRanges, it.status.options.latencyMax)
			break
		}

		it.status.options.latencyRanges = append(it.status.options.latencyRanges, latencyRange)
		latencyRange = latencyRange << 1
	}

	for _, v := range it.status.options.latencyRanges {
		it.status.latencyMap = append(it.status.latencyMap, &BenchStatusCallUsageItem{
			time: v,
		})
	}
}

func (it *BenchStatus) sync(v ReplyStatus, tc int64) {

	//
	if v == ReplyOK {
		it.ok += 1
	} else {
		it.err += 1
	}

	it.latencyTime += tc

	//
	if tc > it.options.latencyMax {
		tc = it.options.latencyMax
	} else if tc < it.options.latencyMin {
		tc = it.options.latencyMin
	}

	//
	for i := 1; i < len(it.latencyMap); i++ {
		if tc < it.latencyMap[i].time {
			it.latencyMap[i-1].num++
			break
		}
	}
}

func (it *BenchStatus) npsSet(v int64) {
	it.npsMap = append(it.npsMap, &BenchStatusCallUsageItem{
		time: v,
		num:  (it.ok + it.err),
	})
}
