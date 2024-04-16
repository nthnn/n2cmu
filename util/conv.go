package util

import (
	"math"
)

func Float32ToBytes(f float32) [4]byte {
	var buf [4]byte
	bits := math.Float32bits(f)

	for i := 0; i < 4; i++ {
		buf[i] = byte(bits >> uint(24-8*i))
	}

	return buf
}

func BytesToFloat32(buf [4]byte) float32 {
	var bits uint32
	for i := 0; i < 4; i++ {
		bits |= uint32(buf[i]) << uint(24-8*i)
	}

	return math.Float32frombits(bits)
}
