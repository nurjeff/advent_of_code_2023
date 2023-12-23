package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Input struct {
	Seeds []uint64
	Maps  [7]map[uint64][2]uint64
}

func solve(input *Input) uint64 {
	var answer uint64 = ^uint64(0)
	for i := 0; i < len(input.Seeds); i += 2 {
		if i+1 < len(input.Seeds) {
			key := [2]uint64{input.Seeds[i], input.Seeds[i] + input.Seeds[i+1]}
			remapRange(key, input.Maps[0], func(key [2]uint64) {
				remapRange(key, input.Maps[1], func(key [2]uint64) {
					remapRange(key, input.Maps[2], func(key [2]uint64) {
						remapRange(key, input.Maps[3], func(key [2]uint64) {
							remapRange(key, input.Maps[4], func(key [2]uint64) {
								remapRange(key, input.Maps[5], func(key [2]uint64) {
									remapRange(key, input.Maps[6], func(key [2]uint64) {
										if key[0] < answer {
											answer = key[0]
										}
									})
								})
							})
						})
					})
				})
			})
		}
	}

	return answer
}

func remapRange(key [2]uint64, src map[uint64][2]uint64, consume func([2]uint64)) {
	start, end := key[0], key[1]

	for end > start {
		var found bool
		var srcRange, dstRange, rangeLen uint64

		for k, v := range src {
			if k < end {
				if !found || k > srcRange {
					found = true
					srcRange, dstRange, rangeLen = k, v[0], v[1]
				}
			}
		}

		if !found || (srcRange+rangeLen <= start) {
			consume([2]uint64{start, end})
			break
		}

		if srcRange+rangeLen < end {
			consume([2]uint64{srcRange + rangeLen, end})
			end = srcRange + rangeLen
		}

		begin := max(start, srcRange)
		consume([2]uint64{
			remap(begin, srcRange, dstRange),
			remap(end, srcRange, dstRange),
		})
		end = begin
	}
}

func remap(key, src, dst uint64) uint64 {
	return dst + (key - src)
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	inp, _ := io.ReadAll(file)
	input := string(inp)

	ip := &Input{
		Maps: [7]map[uint64][2]uint64{
			make(map[uint64][2]uint64),
			make(map[uint64][2]uint64),
			make(map[uint64][2]uint64),
			make(map[uint64][2]uint64),
			make(map[uint64][2]uint64),
			make(map[uint64][2]uint64),
			make(map[uint64][2]uint64),
		},
	}

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	seedsLines := strings.Split(lines[0][7:], " ")

	for i := 0; i < len(seedsLines); i++ {
		s, _ := strconv.ParseInt(seedsLines[i], 10, 64)
		ip.Seeds = append(ip.Seeds, uint64(s))
	}

	mapIndex := 0
	for _, e := range lines[3:] {
		if len(e) == 0 {
			continue
		}
		vals := strings.Split(e, " ")
		v1, err := strconv.ParseInt(vals[0], 10, 64)
		if err != nil {
			mapIndex++
			continue
		}
		v2, _ := strconv.ParseInt(vals[1], 10, 64)
		v3, _ := strconv.ParseInt(vals[2], 10, 64)
		ip.Maps[mapIndex][uint64(v2)] = [2]uint64{uint64(v1), uint64(v3)}
	}

	t := time.Now()
	result := solve(ip)
	fmt.Println(result, time.Since(t))
}
