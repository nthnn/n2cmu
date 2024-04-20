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
)

func main() {
	var network n2.NeuralNetwork
	var epoch uint16 = 0

	machine.UART1.SetBaudRate(31250)
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
			break

		case N2CMU_NET_RESET:
			network.ResetNetwork()
			break

		case N2CMU_NET_TRAIN:
			dataSize := int(uart.ReadUint16())
			dataSet := make([][]float32, dataSize)

			for j := 0; j < dataSize; j++ {
				data := make([]float32, network.InputCount)
				for k := 0; k < int(network.InputCount); k++ {
					data[k] = uart.ReadFloat32()
				}

				dataSet[j] = data
			}

			output := make([][]float32, dataSize)
			for j := 0; j < dataSize; j++ {
				data := make([]float32, network.OutputCount)
				for k := 0; k < int(network.OutputCount); k++ {
					data[k] = uart.ReadFloat32()
				}

				output[j] = data
			}

			learningRate := uart.ReadFloat32()
			for i := 0; i < int(epoch); i++ {
				dataIndex := i % dataSize
				network.Train(dataSet[dataIndex], output[dataIndex], learningRate)
			}

			uart.WriteOk()
			break

		case N2CMU_NET_INFER:
			input := make([]float32, network.InputCount)
			for j := 0; j < int(network.InputCount); j++ {
				input[j] = uart.ReadFloat32()
			}

			output := network.Infer(input)
			for i := 0; i < int(network.OutputCount); i++ {
				uart.WriteFloat32(output[i])
			}

			uart.WriteOk()
			break

		case N2CMU_SET_INPUT_COUNT:
			network.InputCount = uart.ReadUint8()
			network.InitMatrix()
			break

		case N2CMU_GET_INPUT_COUNT:
			uart.WriteUint8(network.InputCount)
			break

		case N2CMU_SET_HIDDEN_COUNT:
			network.HiddenCount = uart.ReadUint8()
			network.InitMatrix()
			break

		case N2CMU_GET_HIDDEN_COUNT:
			uart.WriteUint8(network.HiddenCount)
			break

		case N2CMU_SET_OUTPUT_COUNT:
			network.OutputCount = uart.ReadUint8()
			network.InitMatrix()
			break

		case N2CMU_GET_OUTPUT_COUNT:
			uart.WriteUint8(network.OutputCount)
			break

		case N2CMU_SET_HIDDEN_NEURON:
			for i := 0; i < int(network.HiddenCount); i++ {
				network.HiddenNeuron[i] = uart.ReadFloat32()
			}

			uart.WriteOk()
			break

		case N2CMU_GET_HIDDEN_NEURON:
			for i := 0; i < int(network.HiddenCount); i++ {
				uart.WriteFloat32(network.HiddenNeuron[i])
			}
			break

		case N2CMU_SET_OUTPUT_NEURON:
			for i := 0; i < int(network.OutputCount); i++ {
				network.OutputNeuron[i] = uart.ReadFloat32()
			}

			uart.WriteOk()
			break

		case N2CMU_GET_OUTPUT_NEURON:
			for i := 0; i < int(network.OutputCount); i++ {
				uart.WriteFloat32(network.OutputNeuron[i])
			}
			break

		case N2CMU_SET_HIDDEN_WEIGHTS:
			for i := 0; i < len(network.HiddenWeights); i++ {
				network.HiddenWeights[i] = uart.ReadFloat32()
			}

			uart.WriteOk()
			break

		case N2CMU_GET_HIDDEN_WEIGHTS:
			for i := 0; i < len(network.HiddenWeights); i++ {
				uart.WriteFloat32(network.HiddenWeights[i])
			}
			break

		case N2CMU_SET_OUTPUT_WEIGHTS:
			for i := 0; i < len(network.OutputWeights); i++ {
				network.OutputWeights[i] = uart.ReadFloat32()
			}

			uart.WriteOk()
			break

		case N2CMU_GET_OUTPUT_WEIGHTS:
			for i := 0; i < len(network.OutputWeights); i++ {
				uart.WriteFloat32(network.OutputWeights[i])
			}
			break

		case N2CMU_SET_HIDDEN_BIAS:
			for i := 0; i < int(network.HiddenCount); i++ {
				network.HiddenBias[i] = uart.ReadFloat32()
			}

			uart.WriteOk()
			break

		case N2CMU_GET_HIDDEN_BIAS:
			for i := 0; i < int(network.HiddenCount); i++ {
				uart.WriteFloat32(network.HiddenBias[i])
			}
			break

		case N2CMU_SET_OUTPUT_BIAS:
			for i := 0; i < int(network.OutputCount); i++ {
				network.OutputBias[i] = uart.ReadFloat32()
			}

			uart.WriteOk()
			break

		case N2CMU_GET_OUTPUT_BIAS:
			for i := 0; i < int(network.OutputCount); i++ {
				uart.WriteFloat32(network.OutputBias[i])
			}
			break

		case N2CMU_SET_HIDDEN_GRAD:
			for i := 0; i < int(network.HiddenCount); i++ {
				network.HiddenGrad[i] = uart.ReadFloat32()
			}

			uart.WriteOk()
			break

		case N2CMU_GET_HIDDEN_GRAD:
			for i := 0; i < int(network.HiddenCount); i++ {
				uart.WriteFloat32(network.HiddenGrad[i])
			}
			break

		case N2CMU_SET_EPOCH_COUNT:
			epoch = uart.ReadUint16()
			break

		case N2CMU_GET_EPOCH_COUNT:
			uart.WriteUint16(epoch)
			break
		}
	}
}
