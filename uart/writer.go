/**
 * @file writer.go
 * @brief Provides functions for writing data over
 *        UART (Universal Asynchronous Receiver-Transmitter).
 */
package uart

import (
	"machine"

	"github.com/nthnn/n2cu/util"
)

/**
 * @brief Signals an "OK" message over UART.
 */
func WriteOk() {
	machine.Serial.WriteByte(0x00)
}

/**
 * @brief Signals an "Error" message over UART.
 */
func WriteError() {
	machine.Serial.WriteByte(0x01)
}

/**
 * @brief Writes a float32 value over UART.
 * @param f The float32 value to write.
 */
func WriteFloat32(f float32) {
	buf := util.Float32ToBytes(f)
	machine.Serial.Write(buf[:])
}

/**
 * @brief Writes a uint16 value over UART.
 * @param u The uint16 value to write.
 */
func WriteUint16(u uint16) {
	buf := util.Uint16ToBytes(u)
	machine.Serial.Write(buf[:])
}

/**
 * @brief Writes a uint8 value over UART.
 * @param u The uint8 value to write.
 */
func WriteUint8(u uint8) {
	buf := util.Uint8ToBytes(u)
	machine.Serial.Write(buf[:])
}
