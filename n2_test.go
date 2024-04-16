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
