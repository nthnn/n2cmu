package n2

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
 * @brief Initializes the neural network with specified counts for input, hidden, and output neurons.
 * @param InputCount Number of input neurons.
 * @param HiddenCount Number of hidden neurons.
 * @param OutputCount Number of output neurons.
 */
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
