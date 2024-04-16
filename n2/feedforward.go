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
 * @file feedforward.go
 * @brief Defines a neural network structure and related functions for inference and training.
 */
package n2

/**
 * @brief Performs inference on the neural network with the given input.
 * @param input Input values for inference.
 * @return Output values after inference.
 */
func (neuralNet *NeuralNetwork) Infer(input []float32) []float32 {
	for c := uint8(0); c < neuralNet.HiddenCount; c++ {
		sum := neuralNet.HiddenBias[c]
		for r := uint8(0); r < neuralNet.InputCount; r++ {
			sum += input[r] * neuralNet.HiddenWeights[r*neuralNet.HiddenCount+c]
		}

		neuralNet.HiddenNeuron[c] = sigmoidApprox(sum)
	}

	for c := uint8(0); c < neuralNet.OutputCount; c++ {
		sum := neuralNet.OutputBias[c]
		for r := uint8(0); r < neuralNet.HiddenCount; r++ {
			sum += neuralNet.HiddenNeuron[r] * neuralNet.OutputWeights[r*neuralNet.OutputCount+c]
		}

		neuralNet.OutputNeuron[c] = sigmoidApprox(sum)
	}

	return neuralNet.OutputNeuron
}

/**
 * @brief Trains the neural network using backpropagation.
 * @param input Input values for training.
 * @param output Target output values.
 * @param rate Learning rate for training.
 */
func (neuralNet *NeuralNetwork) Train(input []float32, output []float32, rate float32) {
	neuralNet.Infer(input)

	for c := uint8(0); c < neuralNet.OutputCount; c++ {
		neuralNet.OutputGrad[c] = (neuralNet.OutputNeuron[c] - output[c]) * 0.25
	}

	for r := uint8(0); r < neuralNet.HiddenCount; r++ {
		neuralNet.HiddenGrad[r] = 0
		for c := uint8(0); c < neuralNet.OutputCount; c++ {
			neuralNet.HiddenGrad[r] += neuralNet.OutputGrad[c] * neuralNet.OutputWeights[r*neuralNet.OutputCount+c]
		}

		neuralNet.HiddenGrad[r] *= 0.25
	}

	for r := uint8(0); r < neuralNet.HiddenCount; r++ {
		for c := uint8(0); c < neuralNet.OutputCount; c++ {
			neuralNet.OutputWeights[r*neuralNet.OutputCount+c] -= rate * neuralNet.OutputGrad[c] * neuralNet.HiddenNeuron[r]
		}
	}

	for r := uint8(0); r < neuralNet.InputCount; r++ {
		for c := uint8(0); c < neuralNet.HiddenCount; c++ {
			neuralNet.HiddenWeights[r*neuralNet.HiddenCount+c] -= rate * neuralNet.HiddenGrad[c] * input[r]
		}
	}

	for c := uint8(0); c < neuralNet.OutputCount; c++ {
		neuralNet.OutputBias[c] -= rate * neuralNet.OutputGrad[c]
	}

	for c := uint8(0); c < neuralNet.HiddenCount; c++ {
		neuralNet.HiddenBias[c] -= rate * neuralNet.HiddenGrad[c]
	}
}
