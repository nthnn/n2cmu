package util

import "unsafe"

func Float32ToBytes(f float32) [4]byte {
	bits := *(*[4]byte)(unsafe.Pointer(&f))
	return bits
}

func BytesToFloat32(buf [4]byte) float32 {
	bits := *(*uint32)(unsafe.Pointer(&buf))
	return *(*float32)(unsafe.Pointer(&bits))
}
