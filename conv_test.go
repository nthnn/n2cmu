package main

import (
	"testing"

	"github.com/nthnn/n2cu/util"
)

func TestFloat32ToBytesAndBack(t *testing.T) {
	tests := []struct {
		input float32
	}{
		{input: 3.14},
		{input: -123.456},
		{input: 0.0},
		{input: -0.0},
	}

	for _, test := range tests {
		bytes := util.Float32ToBytes(test.input)
		result := util.BytesToFloat32(bytes)

		if result != test.input {
			t.Errorf("Expected %f, but got %f", test.input, result)
		}
	}
}

func TestUint16ToBytesAndBack(t *testing.T) {
	tests := []struct {
		input uint16
	}{
		{input: 100},
		{input: 65535},
		{input: 0},
	}

	for _, test := range tests {
		bytes := util.Uint16ToBytes(test.input)
		result := util.BytesToUint16(bytes)

		if result != test.input {
			t.Errorf("Expected %d, but got %d", test.input, result)
		}
	}
}

func TestUint8ToBytesAndBack(t *testing.T) {
	tests := []struct {
		input uint8
	}{
		{input: 255},
		{input: 0},
	}

	for _, test := range tests {
		bytes := util.Uint8ToBytes(test.input)
		result := util.BytesToUint8(bytes)

		if result != test.input {
			t.Errorf("Expected %d, but got %d", test.input, result)
		}
	}
}

func BenchmarkFloat32ToBytes(b *testing.B) {
	input := float32(3.14)
	for i := 0; i < b.N; i++ {
		_ = util.Float32ToBytes(input)
	}
}

func BenchmarkBytesToFloat32(b *testing.B) {
	input := [4]byte{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		_ = util.BytesToFloat32(input)
	}
}

func BenchmarkUint16ToBytes(b *testing.B) {
	input := uint16(100)
	for i := 0; i < b.N; i++ {
		_ = util.Uint16ToBytes(input)
	}
}

func BenchmarkBytesToUint16(b *testing.B) {
	input := [2]byte{1, 2}
	for i := 0; i < b.N; i++ {
		_ = util.BytesToUint16(input)
	}
}

func BenchmarkUint8ToBytes(b *testing.B) {
	input := uint8(100)
	for i := 0; i < b.N; i++ {
		_ = util.Uint8ToBytes(input)
	}
}

func BenchmarkBytesToUint8(b *testing.B) {
	input := [1]byte{1}
	for i := 0; i < b.N; i++ {
		_ = util.BytesToUint8(input)
	}
}
