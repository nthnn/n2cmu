package n2

type NeuralNetwork struct {
	InputCount  uint8
	HiddenCount uint8
	OutputCount uint8

	HiddenNeuron  []float32
	HiddenBias    []float32
	HiddenGrad    []float32
	HiddenWeights []float32

	OutputNeuron  []float32
	OutputBias    []float32
	OutputGrad    []float32
	OutputWeights []float32
}

func (neuralNet *NeuralNetwork) InitNetwork(InputCount, HiddenCount, OutputCount uint8) {
	neuralNet.InputCount = InputCount
	neuralNet.HiddenCount = HiddenCount
	neuralNet.OutputCount = OutputCount

	neuralNet.HiddenWeights = make([]float32, InputCount*HiddenCount)
	neuralNet.OutputWeights = make([]float32, HiddenCount*OutputCount)

	neuralNet.HiddenBias = make([]float32, HiddenCount)
	neuralNet.OutputBias = make([]float32, OutputCount)

	neuralNet.HiddenNeuron = make([]float32, HiddenCount)
	neuralNet.OutputNeuron = make([]float32, OutputCount)

	neuralNet.HiddenGrad = make([]float32, HiddenCount)
	neuralNet.OutputGrad = make([]float32, OutputCount)

	for i := range neuralNet.HiddenWeights {
		neuralNet.HiddenWeights[i] = 0.5
	}

	for i := range neuralNet.OutputWeights {
		neuralNet.OutputWeights[i] = 0.5
	}
}

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
