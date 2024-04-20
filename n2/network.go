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
 * @file network.go
 * @brief Defines a neural network structure and related functions for inference and training.
 */
package n2

import "math/rand"

/**
 * @struct NeuralNetwork
 * @brief Represents a feedforward neural network.
 */
type NeuralNetwork struct {
	InputCount  uint8 ///< Number of input neurons.
	HiddenCount uint8 ///< Number of hidden neurons.
	OutputCount uint8 ///< Number of output neurons.

	HiddenNeuron  []float32 ///< Neuron values in the hidden layer.
	HiddenBias    []float32 ///< Bias values for the hidden layer neurons.
	HiddenGrad    []float32 ///< Gradient values for the hidden layer neurons.
	HiddenWeights []float32 ///< Weight values for connections between input and hidden layers.

	OutputNeuron  []float32 ///< Neuron values in the output layer.
	OutputBias    []float32 ///< Bias values for the output layer neurons.
	OutputGrad    []float32 ///< Gradient values for the output layer neurons.
	OutputWeights []float32 ///< Weight values for connections between hidden and output layers.
}

/**
 * @brief Initializes the matrix data for neural network with the previous specifications.
 */
func (neuralNet *NeuralNetwork) InitMatrix() {
	neuralNet.HiddenWeights = make(
		[]float32,
		neuralNet.InputCount*neuralNet.HiddenCount,
	)

	neuralNet.OutputWeights = make(
		[]float32,
		neuralNet.HiddenCount*neuralNet.OutputCount,
	)

	neuralNet.HiddenBias = make([]float32, neuralNet.HiddenCount)
	neuralNet.OutputBias = make([]float32, neuralNet.OutputCount)

	neuralNet.HiddenNeuron = make([]float32, neuralNet.HiddenCount)
	neuralNet.OutputNeuron = make([]float32, neuralNet.OutputCount)

	neuralNet.HiddenGrad = make([]float32, neuralNet.HiddenCount)
	neuralNet.OutputGrad = make([]float32, neuralNet.OutputCount)

	for i := range neuralNet.HiddenWeights {
		neuralNet.HiddenWeights[i] = rand.Float32()
	}

	for i := range neuralNet.OutputWeights {
		neuralNet.OutputWeights[i] = rand.Float32()
	}
}

/**
 * @brief Initializes the neural network with specified counts for input, hidden, and output neurons.
 * @param InputCount Number of input neurons.
 * @param HiddenCount Number of hidden neurons.
 * @param OutputCount Number of output neurons.
 */
func (neuralNet *NeuralNetwork) InitNetwork(InputCount, HiddenCount, OutputCount uint8) {
	neuralNet.InputCount = InputCount
	neuralNet.HiddenCount = HiddenCount
	neuralNet.OutputCount = OutputCount

	neuralNet.InitMatrix()
}

/**
 * @brief Resets the neural network, clearing all counts and arrays.
 */
func (neuralNet *NeuralNetwork) ResetNetwork() {
	neuralNet.InputCount = 0
	neuralNet.HiddenCount = 0
	neuralNet.OutputCount = 0

	neuralNet.HiddenNeuron = nil
	neuralNet.HiddenBias = nil
	neuralNet.HiddenGrad = nil
	neuralNet.HiddenWeights = nil

	neuralNet.OutputNeuron = nil
	neuralNet.OutputBias = nil
	neuralNet.OutputGrad = nil
	neuralNet.OutputWeights = nil
}
