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
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	mrand "math/rand"
	"time"
)

var (
	randBytesMax = 1024 * 1024
	randWordCap  = 1000
	randWordLen  = 8
	randWords    = [][]byte{}
)

func init() {
	mrand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < randWordCap; i++ {
		randWords = append(randWords, randBytes(randWordLen))
	}
}

func RandKey(size int, seq uint64) []byte {
	if seq == 0 {
		return []byte(randHexString(size))
	}
	return []byte(randHexString(size-16) + uint64ToHexString(seq))
}

func RandValue(size int) []byte {
	text := []byte{}
	for i := 0; i < size; i += (randWordLen + 1) {
		text = append(text, bytesClone(randWords[mrand.Intn(randWordCap)])...)
		text = append(text, ' ')
	}
	return text
}

func randBytes(size int) []byte {

	if size < 1 {
		size = 1
	} else if size > randBytesMax {
		size = randBytesMax
	}

	bs := make([]byte, size)

	if _, err := rand.Read(bs); err != nil {
		for i := range bs {
			bs[i] = uint8(mrand.Intn(256))
		}
	}

	return bs
}

func randHexString(length int) string {
	return hex.EncodeToString(randBytes(length / 2))
}

func uint64ToBytes(v uint64) []byte {

	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, v)

	return bs
}

func uint64ToHexString(v uint64) string {
	return hex.EncodeToString(uint64ToBytes(v))
}

func float64Round(f float64, pa_num int64) float64 {
	pa_fix := float64(1e4)
	switch pa_num {
	case 2:
		pa_fix = 1e2
	case 4:
		pa_fix = 1e4
	case 6:
		pa_fix = 1e6
	default:
		pa_fix = 1e4
	}
	return float64(int64(f*pa_fix+0.5)) / pa_fix
}

func bytesClone(src []byte) []byte {

	dst := make([]byte, len(src))
	copy(dst, src)

	return dst
}

func timeus() int64 {
	return (time.Now().UnixNano() / 1e3)
}
