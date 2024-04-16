package main

import (
	"machine"

	"github.com/nthnn/n2cu/n2"
	"github.com/nthnn/n2cu/util"
)

func test() {
	machine.UART1.SetBaudRate(2000000)

	trainingData := [][]float32{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	expectedOutput := []float32{0, 0, 0, 1}

	network := n2.CreateNetwork(2, 2, 1)
	for i := 0; i < 4000; i++ {
		dataIndex := i % 4
		network.Train(trainingData[dataIndex], []float32{expectedOutput[dataIndex]}, 1.0)
	}

	for _, data := range trainingData {
		infered := util.Float32ToBytes(network.Infer(data)[0])
		machine.Serial.Write(infered[:])
	}

	network.ResetNetwork()
}

func main() {
	var network n2.NeuralNetwork

	for machine.Serial.Buffered() > 0 {
		command, _ := machine.Serial.ReadByte()
		switch command {
		case N2CMU_PROC_HANDSHAKE:
			machine.Serial.WriteByte(0x01)
			break

		case N2CMU_NET_RESET:
			network.ResetNetwork()
			machine.Serial.WriteByte(0x01)
			break
		}
	}
}
