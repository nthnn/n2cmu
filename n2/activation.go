package n2

func sigmoidApprox(d float32) float32 {
	result := 0.5 + (0.5 * d)
	if result < 0 {
		return 0.0
	}

	return result
}
