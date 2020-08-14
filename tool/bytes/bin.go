package bytes

import "encoding/binary"

func Uint32ToBytes(data uint32, isLittleEndian bool) []byte {
	tmp := make([]byte, 4)
	if isLittleEndian{
		binary.LittleEndian.PutUint32(tmp, data)
	}
	return tmp
}
func Uint16ToBytes(data uint16, isLittleEndian bool) []byte {
	tmp := make([]byte, 2)
	if isLittleEndian{
		binary.LittleEndian.PutUint16(tmp, data)
	}
	return tmp
}
