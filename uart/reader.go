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

/**
 * @file reader.go
 * @brief Provides functions for reading data over
 *        UART (Universal Asynchronous Receiver-Transmitter).
 */
package uart

import (
	"machine"

	"github.com/nthnn/n2cu/util"
)

/**
 * @brief Reads a float32 value from UART.
 * @return The read float32 value and nil if successful, or 0 and an error if unsuccessful.
 */
func ReadFloat32() (float32, error) {
	var err error
	buf := []byte{0, 0, 0, 0}

	buf[0], err = machine.Serial.ReadByte()
	for err != nil {
		buf[0], err = machine.Serial.ReadByte()
	}

	buf[1], err = machine.Serial.ReadByte()
	for err != nil {
		buf[1], err = machine.Serial.ReadByte()
	}

	buf[2], err = machine.Serial.ReadByte()
	for err != nil {
		buf[2], err = machine.Serial.ReadByte()
	}

	buf[3], err = machine.Serial.ReadByte()
	for err != nil {
		buf[3], err = machine.Serial.ReadByte()
	}

	return util.BytesToFloat32([4]byte(buf)), nil
}

/**
 * @brief Reads a uint16 value from UART.
 * @return The read uint16 value and nil if successful, or 0 and an error if unsuccessful.
 */
func ReadUint16() uint16 {
	var err error
	buf := []byte{0, 0}

	buf[0], err = machine.Serial.ReadByte()
	for err != nil {
		buf[0], err = machine.Serial.ReadByte()
	}

	buf[1], err = machine.Serial.ReadByte()
	for err != nil {
		buf[1], err = machine.Serial.ReadByte()
	}

	return util.BytesToUint16([2]byte(buf))
}

/**
 * @brief Reads a uint8 value from UART.
 * @return The read uint8 value and nil if successful, or 0 and an error if unsuccessful.
 */
func ReadUint8() uint8 {
	buf := []byte{0}
	_, err := machine.Serial.Read(buf)

	for err != nil {
		_, err = machine.Serial.Read(buf)
	}

	return buf[0]
}
