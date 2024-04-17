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
 * @file writer.go
 * @brief Provides functions for writing data over
 *        UART (Universal Asynchronous Receiver-Transmitter).
 */
package uart

import (
	"machine"

	"github.com/nthnn/n2cu/util"
)

/**
 * @brief Signals an "OK" message over UART.
 */
func WriteOk() {
	machine.Serial.WriteByte(0x01)
}

/**
 * @brief Signals an "Error" message over UART.
 */
func WriteError() {
	machine.Serial.WriteByte(0x00)
}

/**
 * @brief Writes a float32 value over UART.
 * @param f The float32 value to write.
 */
func WriteFloat32(f float32) {
	buf := util.Float32ToBytes(f)
	machine.Serial.Write(buf[:])
}

/**
 * @brief Writes a uint16 value over UART.
 * @param u The uint16 value to write.
 */
func WriteUint16(u uint16) {
	buf := util.Uint16ToBytes(u)
	machine.Serial.Write(buf[:])
}

/**
 * @brief Writes a uint8 value over UART.
 * @param u The uint8 value to write.
 */
func WriteUint8(u uint8) {
	buf := util.Uint8ToBytes(u)
	machine.Serial.Write(buf[:])
}
