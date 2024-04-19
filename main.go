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
	"machine"

	"github.com/nthnn/n2cu/n2"
	"github.com/nthnn/n2cu/uart"
	"github.com/nthnn/n2cu/util"
)

func main() {
	var network n2.NeuralNetwork
	var epoch uint16 = 0

	machine.UART1.SetBaudRate(115200)
	for {
		command, error := machine.Serial.ReadByte()
		if error != nil {
			continue
		}

		switch command {
		case N2CMU_PROC_HANDSHAKE:
			uart.WriteOk()
			break

		case N2CMU_PROC_CPU_RESET:
			machine.CPUReset()
			break

		case N2CMU_NET_CREATE:
			inputCount := uart.ReadUint8()
			hiddenCount := uart.ReadUint8()
			outputCount := uart.ReadUint8()

			network.InitNetwork(inputCount, hiddenCount, outputCount)
			uart.WriteOk()
			break

		case N2CMU_NET_RESET:
			network.ResetNetwork()
			uart.WriteOk()

			break

		case N2CMU_SET_EPOCH_COUNT:
			epoch = uart.ReadUint16()
			uart.WriteOk()
			break

		case N2CMU_GET_EPOCH_COUNT:
			buf := util.Uint16ToBytes(epoch)
			machine.Serial.Write([]byte{buf[0], buf[1]})
			break
		}
	}
}
