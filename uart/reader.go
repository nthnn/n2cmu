package uart

import (
	"machine"

	"github.com/nthnn/n2cu/util"
)

func ReadFloat32() (float32, error) {
	var buf []byte
	count, readErr := machine.Serial.Read(buf)

	if count == 4 && readErr != nil {
		return util.BytesToFloat32([4]byte(buf)), nil
	}

	return 0, readErr
}

func ReadUint16() (uint16, error) {
	var buf []byte
	count, readErr := machine.Serial.Read(buf)

	if count == 4 && readErr != nil {
		return util.BytesToUint16([4]byte(buf)), nil
	}

	return 0, readErr
}
