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
 * @file command.go
 * @brief Defines constants for communication between
 *        the main microcontroller and the N2CMU.
 */
package main

const (
	N2CMU_PROC_HANDSHAKE = 0x00 ///< Command constant for the N2CMU handshake signal.
	N2CMU_PROC_CPU_RESET = 0x01 ///< Constant for resetting the CPU.

	N2CMU_NET_CREATE = 0x02 //< Command constant for creating a neural network.
	N2CMU_NET_RESET  = 0x03 //< Command constant for resetting a neural network.
	N2CMU_NET_TRAIN  = 0x04 //< Command constant for training a neural network.
	N2CMU_NET_INFER  = 0x05 //< Command constant for making inference with the neural network.

	N2CMU_SET_INPUT_COUNT    = 0x06 //< Command constant for setting the number of input neurons.
	N2CMU_SET_HIDDEN_COUNT   = 0x07 //< Command constant for setting the number of hidden neurons.
	N2CMU_SET_OUTPUT_COUNT   = 0x08 //< Command constant for setting the number of output neurons.
	N2CMU_SET_HIDDEN_NEURON  = 0x09 //< Command constant for setting hidden neuron values.
	N2CMU_SET_OUTPUT_NEURON  = 0x0a //< Command constant for setting output neuron values.
	N2CMU_SET_HIDDEN_WEIGHTS = 0x0b //< Command constant for setting hidden neuron weights
	N2CMU_SET_OUTPUT_WEIGHTS = 0x0c //< Command constant for setting output neuron weights.
	N2CMU_SET_HIDDEN_BIAS    = 0x0d //< Command constant for setting hidden neuron biases.
	N2CMU_SET_OUTPUT_BIAS    = 0x0e //< Command constant for setting output neuron biases.
	N2CMU_SET_HIDDEN_GRAD    = 0x0f //< Command constant for setting hidden neuron gradients.
	N2CMU_SET_OUTPUT_GRAD    = 0x10 //< Command constant for setting output neuron gradients.
	N2CMU_SET_EPOCH_COUNT    = 0x11 //< Command constant for setting the epoch count for training.

	N2CMU_GET_INPUT_COUNT    = 0x12 //< Command constant for getting the number of input neurons.
	N2CMU_GET_HIDDEN_COUNT   = 0x13 //< Command constant for getting the number of hidden neurons.
	N2CMU_GET_OUTPUT_COUNT   = 0x14 //< Command constant for getting the number of output neurons.
	N2CMU_GET_HIDDEN_NEURON  = 0x15 //< Command constant for getting hidden neuron values.
	N2CMU_GET_OUTPUT_NEURON  = 0x16 //< Command constant for getting output neuron values.
	N2CMU_GET_HIDDEN_WEIGHTS = 0x17 //< Command constant for getting hidden neuron weights.
	N2CMU_GET_OUTPUT_WEIGHTS = 0x18 //< Command constant for getting output neuron weights.
	N2CMU_GET_HIDDEN_BIAS    = 0x19 //< Command constant for getting hidden neuron biases.
	N2CMU_GET_OUTPUT_BIAS    = 0x1a //< Command constant for getting output neuron biases.
	N2CMU_GET_HIDDEN_GRAD    = 0x1b //< Command constant for getting hidden neuron gradients.
	N2CMU_GET_OUTPUT_GRAD    = 0x1c //< Command constant for getting output neuron gradients.
	N2CMU_GET_EPOCH_COUNT    = 0x1d //< Command constant for getting the epoch count of training.
)
