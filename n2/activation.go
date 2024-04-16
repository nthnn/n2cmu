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

/**
 * @file activations.go
 * @brief Defines a neural network structure and related functions for inference and training.
 */
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
