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

	"github.com/nthnn/n2cu/n2"
)

func TestNeuralNetwork(t *testing.T) {
	trainingData := [][]float32{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	expectedOutput := []float32{0, 0, 0, 1}

	var network n2.NeuralNetwork
	network.InitNetwork(2, 2, 1)

	for i := 0; i < 4000; i++ {
		dataIndex := i % 4
		network.Train(trainingData[dataIndex], []float32{expectedOutput[dataIndex]}, 1.0)
	}

	for _, data := range trainingData {
		fmt.Printf("%.0f, %.0f = %.0f\r\n",
			data[0],
			data[1],
			network.Infer(data)[0],
		)
	}

	network.ResetNetwork()
}
