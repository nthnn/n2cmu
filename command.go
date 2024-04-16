package main

const (
	N2CMU_SET_INPUT_COUNT    = 0x00
	N2CMU_SET_HIDDEN_COUNT   = 0x01
	N2CMU_SET_OUTPUT_COUNT   = 0x20
	N2CMU_SET_HIDDEN_NEURON  = 0x03
	N2CMU_SET_OUTPUT_NEURON  = 0x04
	N2CMU_SET_HIDDEN_WEIGHTS = 0x05
	N2CMU_SET_OUTPUT_WEIGHTS = 0x06
	N2CMU_SET_HIDDEN_BIAS    = 0x07
	N2CMU_SET_OUTPUT_BIAS    = 0x08
	N2CMU_SET_HIDDEN_GRAD    = 0x09
	N2CMU_SET_OUTPUT_GRAD    = 0x0a

	N2CMU_GET_INPUT_COUNT    = 0x0b
	N2CMU_GET_HIDDEN_COUNT   = 0x0c
	N2CMU_GET_OUTPUT_COUNT   = 0x0d
	N2CMU_GET_HIDDEN_NEURON  = 0x0e
	N2CMU_GET_OUTPUT_NEURON  = 0x0f
	N2CMU_GET_HIDDEN_WEIGHTS = 0x10
	N2CMU_GET_OUTPUT_WEIGHTS = 0x11
	N2CMU_GET_HIDDEN_BIAS    = 0x12
	N2CMU_GET_OUTPUT_BIAS    = 0x13
	N2CMU_GET_HIDDEN_GRAD    = 0x14
	N2CMU_GET_OUTPUT_GRAD    = 0x15
)