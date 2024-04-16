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

func Uint16ToBytes(u uint16) [2]byte {
	bits := *(*[2]byte)(unsafe.Pointer(&u))
	return bits
}

func BytesToUint16(buf [2]byte) uint16 {
	bits := *(*uint32)(unsafe.Pointer(&buf))
	return *(*uint16)(unsafe.Pointer(&bits))
}

func Uint8ToBytes(u uint8) [1]byte {
	bits := *(*[1]byte)(unsafe.Pointer(&u))
	return bits
}

func BytesToUint8(buf [1]byte) uint8 {
	bits := *(*uint32)(unsafe.Pointer(&buf))
	return *(*uint8)(unsafe.Pointer(&bits))
}
