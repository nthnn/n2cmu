package n2

/**
 * @brief Applies the sigmoid activation function.
 * @param d Input value.
 * @return Output value after applying the sigmoid function.
 */
func sigmoidApprox(d float32) float32 {
	result := 0.5 + (0.5 * d)
	if result < 0.0 {
		return 0.0
	}

	return result
}
