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
	var buf [4]byte
	for i := uint8(0); i < uint8(4); i++ {
		value, err := machine.Serial.ReadByte()

		if err != nil {
			return float32(0.0), err
		}

		buf[i] = value
	}

	return util.BytesToFloat32(buf), nil
}

/**
 * @brief Reads a uint16 value from UART.
 * @return The read uint16 value and nil if successful, or 0 and an error if unsuccessful.
 */
func ReadUint16() (uint16, error) {
	var buf [2]byte
	for i := uint8(0); i < uint8(2); i++ {
		value, err := machine.Serial.ReadByte()

		if err != nil {
			return uint16(0), err
		}

		buf[i] = value
	}

	return util.BytesToUint16(buf), nil
}

/**
 * @brief Reads a uint8 value from UART.
 * @return The read uint8 value and nil if successful, or 0 and an error if unsuccessful.
 */
func ReadUint8() (uint8, error) {
	value, err := machine.Serial.ReadByte()
	if err != nil {
		return uint8(0), err
	}

	return value, nil
}
