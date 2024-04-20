/*
 * This file is part of the N2CMU (https://github.com/nthnn/n2cmu).
 * Copyright (c) 2024 Nathanne Isip.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"
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
	}

	for _, test := range tests {
		bytes := util.Float32ToBytes(test.input)
		result := util.BytesToFloat32(bytes)

		if result != test.input {
			t.Errorf("Expected %f, but got %f", test.input, result)
		} else {
			fmt.Printf("%v = %.3f\r\n", bytes, result)
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
		} else {
			fmt.Printf("%v = %d\r\n", bytes, result)
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
