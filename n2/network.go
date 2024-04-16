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

func CreateNetwork(InputCount, HiddenCount, OutputCount uint8) *NeuralNetwork {
	network := &NeuralNetwork{
		InputCount:    InputCount,
		HiddenCount:   HiddenCount,
		OutputCount:   OutputCount,
		HiddenWeights: make([]float32, InputCount*HiddenCount),
		OutputWeights: make([]float32, HiddenCount*OutputCount),
		HiddenBias:    make([]float32, HiddenCount),
		OutputBias:    make([]float32, OutputCount),
		HiddenNeuron:  make([]float32, HiddenCount),
		OutputNeuron:  make([]float32, OutputCount),
		HiddenGrad:    make([]float32, HiddenCount),
		OutputGrad:    make([]float32, OutputCount),
	}

	for i := range network.HiddenWeights {
		network.HiddenWeights[i] = 0.5
	}

	for i := range network.OutputWeights {
		network.OutputWeights[i] = 0.5
	}

	return network
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
