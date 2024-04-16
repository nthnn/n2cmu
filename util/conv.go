/**
 * @file util.go
 * @brief Contains utility functions for converting
 *        between float32, uint16, and uint8 and their byte representations.
 */

package util

import "unsafe"

/**
 * @brief Converts a float32 value to its byte representation.
 * @param f The float32 value to convert.
 * @return The byte representation of the float32 value.
 */
func Float32ToBytes(f float32) [4]byte {
	bits := *(*[4]byte)(unsafe.Pointer(&f))
	return bits
}

/**
 * @brief Converts a byte representation to a float32 value.
 * @param buf The byte representation to convert.
 * @return The float32 value converted from the byte representation.
 */
func BytesToFloat32(buf [4]byte) float32 {
	bits := *(*uint32)(unsafe.Pointer(&buf))
	return *(*float32)(unsafe.Pointer(&bits))
}

/**
 * @brief Converts a uint16 value to its byte representation.
 * @param u The uint16 value to convert.
 * @return The byte representation of the uint16 value.
 */
func Uint16ToBytes(u uint16) [2]byte {
	bits := *(*[2]byte)(unsafe.Pointer(&u))
	return bits
}

/**
 * @brief Converts a byte representation to a uint16 value.
 * @param buf The byte representation to convert.
 * @return The uint16 value converted from the byte representation.
 */
func BytesToUint16(buf [2]byte) uint16 {
	bits := *(*uint32)(unsafe.Pointer(&buf))
	return *(*uint16)(unsafe.Pointer(&bits))
}

/**
 * @brief Converts a uint8 value to its byte representation.
 * @param u The uint8 value to convert.
 * @return The byte representation of the uint8 value.
 */
func Uint8ToBytes(u uint8) [1]byte {
	bits := *(*[1]byte)(unsafe.Pointer(&u))
	return bits
}

/**
 * @brief Converts a byte representation to a uint8 value.
 * @param buf The byte representation to convert.
 * @return The uint8 value converted from the byte representation.
 */
func BytesToUint8(buf [1]byte) uint8 {
	bits := *(*uint32)(unsafe.Pointer(&buf))
	return *(*uint8)(unsafe.Pointer(&bits))
}
