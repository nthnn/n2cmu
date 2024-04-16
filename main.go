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
	machine.UART1.SetBaudRate(2000000)

	var network n2.NeuralNetwork
	var epoch uint16 = 0

	for machine.Serial.Buffered() > 0 {
		command, _ := machine.Serial.ReadByte()

		switch command {
		case N2CMU_PROC_HANDSHAKE:
			uart.WriteOk()
			break

		case N2CMU_PROC_CPU_RESET:
			machine.CPUReset()
			uart.WriteOk()

			break

		case N2CMU_NET_CREATE:
			inputCount, err := uart.ReadUint8()
			if err != nil {
				uart.WriteError()
				continue
			}

			hiddenCount, err := uart.ReadUint8()
			if err != nil {
				uart.WriteError()
				continue
			}

			outputCount, err := uart.ReadUint8()
			if err != nil {
				uart.WriteError()
				continue
			}

			network.InitNetwork(inputCount, hiddenCount, outputCount)
			uart.WriteOk()
			break

		case N2CMU_NET_RESET:
			network.ResetNetwork()
			uart.WriteOk()

			break

		case N2CMU_SET_EPOCH_COUNT:
			if num, err := uart.ReadUint16(); err != nil {
				epoch = num
				uart.WriteOk()
			} else {
				uart.WriteError()
			}

			break

		case N2CMU_GET_EPOCH_COUNT:
			buf := util.Uint16ToBytes(epoch)
			machine.Serial.Write(buf[:])

			break
		}
	}
}
