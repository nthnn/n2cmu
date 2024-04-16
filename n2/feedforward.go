package n2

func (nt *NeuralNetwork) Infer(input []float32) []float32 {
	for c := uint8(0); c < nt.HiddenCount; c++ {
		sum := nt.HiddenBias[c]
		for r := uint8(0); r < nt.InputCount; r++ {
			sum += input[r] * nt.HiddenWeights[r*nt.HiddenCount+c]
		}

		nt.HiddenNeuron[c] = sigmoidApprox(sum)
	}

	for c := uint8(0); c < nt.OutputCount; c++ {
		sum := nt.OutputBias[c]
		for r := uint8(0); r < nt.HiddenCount; r++ {
			sum += nt.HiddenNeuron[r] * nt.OutputWeights[r*nt.OutputCount+c]
		}

		nt.OutputNeuron[c] = sigmoidApprox(sum)
	}

	return nt.OutputNeuron
}

func (nt *NeuralNetwork) Train(input []float32, y []float32, lr float32) {
	nt.Infer(input)

	for c := uint8(0); c < nt.OutputCount; c++ {
		nt.OutputGrad[c] = (nt.OutputNeuron[c] - y[c]) * 0.25
	}

	for r := uint8(0); r < nt.HiddenCount; r++ {
		nt.HiddenGrad[r] = 0
		for c := uint8(0); c < nt.OutputCount; c++ {
			nt.HiddenGrad[r] += nt.OutputGrad[c] * nt.OutputWeights[r*nt.OutputCount+c]
		}

		nt.HiddenGrad[r] *= 0.25
	}

	for r := uint8(0); r < nt.HiddenCount; r++ {
		for c := uint8(0); c < nt.OutputCount; c++ {
			nt.OutputWeights[r*nt.OutputCount+c] -= lr * nt.OutputGrad[c] * nt.HiddenNeuron[r]
		}
	}

	for r := uint8(0); r < nt.InputCount; r++ {
		for c := uint8(0); c < nt.HiddenCount; c++ {
			nt.HiddenWeights[r*nt.HiddenCount+c] -= lr * nt.HiddenGrad[c] * input[r]
		}
	}

	for c := uint8(0); c < nt.OutputCount; c++ {
		nt.OutputBias[c] -= lr * nt.OutputGrad[c]
	}

	for c := uint8(0); c < nt.HiddenCount; c++ {
		nt.HiddenBias[c] -= lr * nt.HiddenGrad[c]
	}
}
