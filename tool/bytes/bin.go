package bytes

import "encoding/binary"

func Uint32ToBytes(data uint32, isLittleEndian bool) []byte {
	tmp := make([]byte, 4)
	if isLittleEndian{
		binary.LittleEndian.PutUint32(tmp, data)
	}
	return tmp
}
