package uart

import (
	"machine"

	"github.com/nthnn/n2cu/util"
)

func WriteOk() {
	machine.Serial.WriteByte(0x00)
}

func WriteError() {
	machine.Serial.WriteByte(0x01)
}

func WriteFloat32(f float32) {
	buf := util.Float32ToBytes(f)
	machine.Serial.Write(buf[:])
}

func WriteUint16(u uint16) {
	buf := util.Uint16ToBytes(u)
	machine.Serial.Write(buf[:])
}

func WriteUint8(u uint8) {
	buf := util.Uint8ToBytes(u)
	machine.Serial.Write(buf[:])
}
