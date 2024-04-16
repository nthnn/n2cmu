/**
 * @file uart.go
 * @brief Provides functions for reading data over
 *        UART (Universal Asynchronous Receiver-Transmitter).
 */
package uart

import (
	"machine"

	"github.com/nthnn/n2cu/util"
)

/**
 * @brief Reads a float32 value from UART.
 * @return The read float32 value and nil if successful, or 0 and an error if unsuccessful.
 */
func ReadFloat32() (float32, error) {
	var buf []byte
	count, readErr := machine.Serial.Read(buf)

	if count == 4 && readErr != nil {
		return util.BytesToFloat32([4]byte(buf)), nil
	}

	return 0, readErr
}

/**
 * @brief Reads a uint16 value from UART.
 * @return The read uint16 value and nil if successful, or 0 and an error if unsuccessful.
 */
func ReadUint16() (uint16, error) {
	var buf []byte
	count, readErr := machine.Serial.Read(buf)

	if count == 2 && readErr != nil {
		return util.BytesToUint16([2]byte(buf)), nil
	}

	return 0, readErr
}

/**
 * @brief Reads a uint8 value from UART.
 * @return The read uint8 value and nil if successful, or 0 and an error if unsuccessful.
 */
func ReadUint8() (uint8, error) {
	var buf []byte
	count, readErr := machine.Serial.Read(buf)

	if count == 1 && readErr != nil {
		return util.BytesToUint8([1]byte(buf)), nil
	}

	return 0, readErr
}
