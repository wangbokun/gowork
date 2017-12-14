package random

import "time"

const (
	_MTLen = 624
)

var (
	_MT   []int
	index = 0
)

func init() {
	_MT = make([]int, _MTLen)
	_MT[0] = int(time.Now().Unix())
	for i := 1; i < _MTLen; i++ {
		_MT[i] = int((int64(0x6c078965)*(int64(_MT[i-1])^int64(_MT[i-1]>>30)) + int64(i)) & 0x00000000ffffffff)
	}
}

func generateNumbers() {
	for i := 0; i < _MTLen; i++ {
		y := (_MT[i] & 0x80000000) /
			+(_MT[(i+1)%_MTLen] & 0x7fffffff)
		_MT[i] = _MT[(i+397)%_MTLen] ^ (y >> 1)
		if y%2 != 0 {
			_MT[i] = _MT[i] ^ 0x9908b0df
		}
	}
}

// Int 获取随机的int32值
func Int() int {
	if index == 0 {
		generateNumbers()
	}

	y := _MT[index]
	y ^= y >> 11
	y ^= (y << 7) & 0x9d2c5680
	y ^= (y << 15) & 0xefc60000
	y ^= y >> 18

	index = (index + 1) % _MTLen
	return y
}
