package n2

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

func (neuralNet *NeuralNetwork) Train(input []float32, y []float32, lr float32) {
	neuralNet.Infer(input)

	for c := uint8(0); c < neuralNet.OutputCount; c++ {
		neuralNet.OutputGrad[c] = (neuralNet.OutputNeuron[c] - y[c]) * 0.25
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
			neuralNet.OutputWeights[r*neuralNet.OutputCount+c] -= lr * neuralNet.OutputGrad[c] * neuralNet.HiddenNeuron[r]
		}
	}

	for r := uint8(0); r < neuralNet.InputCount; r++ {
		for c := uint8(0); c < neuralNet.HiddenCount; c++ {
			neuralNet.HiddenWeights[r*neuralNet.HiddenCount+c] -= lr * neuralNet.HiddenGrad[c] * input[r]
		}
	}

	for c := uint8(0); c < neuralNet.OutputCount; c++ {
		neuralNet.OutputBias[c] -= lr * neuralNet.OutputGrad[c]
	}

	for c := uint8(0); c < neuralNet.HiddenCount; c++ {
		neuralNet.HiddenBias[c] -= lr * neuralNet.HiddenGrad[c]
	}
}
