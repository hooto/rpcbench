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
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/hooto/hchart/v2/hcapi"
	"github.com/lessos/lessgo/encoding/json"
	psLoad "github.com/shirou/gopsutil/load"
	psNet "github.com/shirou/gopsutil/net"
)

func (it *Bench) runClient() error {

	var (
		avg *psLoad.AvgStat
		err error
	)

	for {
		avg, err = psLoad.Avg()

		if err == nil && avg.Load1 < it.cfg.loadAvg {
			break
		}

		fmt.Printf("Waiting %.2f\r", avg.Load1)
		time.Sleep(3e9)
	}

	fmt.Printf("Bench %s ...\n", it.cfg.Name)

	infos := []string{}

	infos = append(infos, []string{
		"Start at", time.Now().Format("2006-01-02 15:04:05")}...)

	/**
	infos = append(infos, []string{
		"Avg Load", fmt.Sprintf("%.2f", avg.Load1)}...)
	*/

	infos = append(infos, []string{
		"Bench Time, Client Num", fmt.Sprintf("%d s, %d", it.cfg.Time, it.cfg.ClientNum)}...)

	infos = append(infos, []string{
		"Call Write/Read Size", fmt.Sprintf("%d/%d bytes", it.cfg.WriteSize, it.cfg.ReadSize)}...)

	if it.cfg.WorkTime > 0 {
		infos = append(infos, []string{
			"Call Work Time", fmt.Sprintf("%d us", it.cfg.WorkTime)}...)
	}

	for i := 0; i < len(infos); i += 2 {
		fmt.Printf("%30s   %s\n", infos[i], infos[i+1])
	}

	var pSets hcapi.DataList
	if err := json.DecodeFile(it.status.options.dataFile, &pSets); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	var (
		clients            = []ClientConnInterface{}
		ticker             = time.NewTicker(time.Second * time.Duration(it.status.options.timeStep))
		gts                = time.Now().UnixNano() / 1e3
		timeUsed           = int64(0)
		statsNetRecv int64 = 0
		statsNetSent int64 = 0
	)
	defer ticker.Stop()

	nio, _ := psNet.IOCounters(false)
	statsNetRecv = int64(nio[0].BytesRecv)
	statsNetSent = int64(nio[0].BytesSent)

	for n := 0; n < it.cfg.ClientNum; n++ {

		c, err := it.proj.ClientConn(&ClientConnOptions{
			Addr:    fmt.Sprintf("127.0.0.1:%d", it.cfg.ServerPort),
			ConnNum: it.cfg.ClientNum,
		})

		if err != nil {
			return err
		}

		clients = append(clients, c)
	}

	for _, c := range clients {

		go func(c ClientConnInterface) {

			for !it.quit {

				ts := timeus()
				st, workTime := it.clientCall(c)
				tc := timeus() - ts - workTime

				if tc <= 0 {
					tc = 0
				}

				it.status.sync(st, tc)
			}
		}(c)
	}

	it.status.npsSet(0)
	for !it.quit {
		_ = <-ticker.C
		timeUsed += it.status.options.timeStep
		it.status.npsSet(timeUsed)
		if timeUsed >= it.status.options.timeLen {
			it.quit = true
		}
	}

	gtc := (time.Now().UnixNano() / 1e3) - gts
	if gtc < 1 {
		gtc = 1
	}

	it.status.nps = (float64(it.status.ok+it.status.err) / float64(gtc)) * 1e6

	if it.status.ok > 0 && len(it.status.npsMap) > 0 {

		ds := hcapi.NewDataItem(it.status.options.dataName)
		ds.AttrSet("throughput")
		ds.AttrSet(fmt.Sprintf("client-num:%d", it.cfg.ClientNum))
		ds.AttrSet(fmt.Sprintf("write-read-size:%d-%d",
			it.cfg.WriteSize, it.cfg.ReadSize))

		/**
		for _, av := range fn.Attrs() {
			ds.AttrSet(av)
		}
		*/

		for _, v := range it.status.npsMap {

			ds.Points = append(ds.Points, &hcapi.DataPoint{
				X: float64(v.time),
				Y: float64(v.num),
			})
		}

		pSets.Set(ds)
	}

	if it.status.ok > 0 && len(it.status.latencyMap) > 0 {

		ds := hcapi.NewDataItem(it.status.options.dataName)
		ds.AttrSet("latency-avg")
		ds.AttrSet(fmt.Sprintf("client-num:%d", it.cfg.ClientNum))
		ds.AttrSet(fmt.Sprintf("write-read-size:%d-%d",
			it.cfg.WriteSize, it.cfg.ReadSize))
		/**
		for _, av := range fn.Attrs() {
			ds.AttrSet(av)
		}
		*/
		ds.Points = append(ds.Points, &hcapi.DataPoint{
			Y: float64Round(float64(it.status.latencyTime)/float64(it.status.ok), 4),
		})
		pSets.Set(ds)

		ds = hcapi.NewDataItem(it.status.options.dataName)
		ds.AttrSet("latency")
		ds.AttrSet(fmt.Sprintf("client-num:%d", it.cfg.ClientNum))
		ds.AttrSet(fmt.Sprintf("write-read-size:%d-%d",
			it.cfg.WriteSize, it.cfg.ReadSize))
		/**
		for _, av := range fn.Attrs() {
			ds.AttrSet(av)
		}
		*/

		for _, v := range it.status.latencyMap {

			ds.Points = append(ds.Points, &hcapi.DataPoint{
				X: float64(v.time),
				Y: float64(v.num),
			})
		}

		pSets.Set(ds)
	}

	if err := json.EncodeToFile(pSets, it.status.options.dataFile, "  "); err != nil {
		return err
	}

	nio, _ = psNet.IOCounters(false)

	ctAvg := int64(-1)
	if n := int64(it.status.nps); n > 0 {
		ctAvg = (1e6 * int64(it.cfg.ClientNum)) / n
	}

	sizeAvgSent, sizeAvgRecv := int64(0), int64(0)
	if n := it.status.ok + it.status.err; n > 0 {
		sizeAvgSent = (int64(nio[0].BytesSent) - statsNetSent) / n
		sizeAvgRecv = (int64(nio[0].BytesRecv) - statsNetRecv) / n
	}

	infos = []string{
		"Avg Throughput QPS", fmt.Sprintf("%d", int64(it.status.nps)),
		"Avg Sent/Recv Size", fmt.Sprintf("%d/%d bytes", sizeAvgSent, sizeAvgRecv),
	}

	infos = append(infos, "Avg Call Latency Time")
	if ctAvg >= 10000 {
		infos = append(infos, fmt.Sprintf("%d ms", ctAvg/1e3))
	} else {
		infos = append(infos, fmt.Sprintf("%d us", ctAvg))
	}

	if it.status.err > 0 {
		infos = append(infos, []string{
			"Number of Success/Fail", fmt.Sprintf("%d, %d", it.status.ok, it.status.err),
		}...)
	}

	for i := 0; i < len(infos); i += 2 {
		fmt.Printf("%30s > %s\n", infos[i], infos[i+1])
	}

	time.Sleep(1e9)

	return nil
}

func (it *Bench) clientCall(c ClientConnInterface) (ReplyStatus, int64) {

	ctx, fc := context.WithTimeout(context.Background(), CallTimeout)
	defer fc()

	req := &BenchRequest{
		Id: rand.Uint64(),
		Payload: &BenchPayload{
			Body: RandValue(int(it.cfg.WriteSize)),
		},
		ReplySpec: &BenchPayloadSpec{
			Size:     int32(it.cfg.ReadSize),
			WorkTime: 0,
		},
	}

	if it.cfg.WorkTime > 0 {

		// if WorkTime == 1000 us
		//  3% 10 ~ 100 ms
		//  7% 1 ~ 10 ms
		// 90% 0 ~ 1000 us
		if p := rand.Intn(100); p < 3 {
			req.ReplySpec.WorkTime = int32((it.cfg.WorkTime * 10) + rand.Int63n(it.cfg.WorkTime*90))
		} else if p < 10 {
			req.ReplySpec.WorkTime = int32((it.cfg.WorkTime) + rand.Int63n(it.cfg.WorkTime*9))
		} else {
			req.ReplySpec.WorkTime = int32(rand.Int63n(it.cfg.WorkTime))
		}
	}

	if rep, err := c.UnaryCall(ctx, req); err == nil && rep.Id == req.Id {
		return ReplyOK, int64(req.ReplySpec.WorkTime)
	}

	return ReplyER, int64(req.ReplySpec.WorkTime)
}
