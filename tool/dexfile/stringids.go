package dexfile

import (
	"dexeditor/tool/bytes"
	"dexeditor/tool/debug"
)

type StringIds struct {
	StringIds []*StringId
}
type StringId struct {
	DataOff uint32
	Size    *ULeb128 //这个是字符长度，并不是字节长度
	Data    []byte
}

func NewStringIds(dex []byte, off, size uint32) *StringIds {
	that := &StringIds{}
	that.StringIds = make([]*StringId, size)
	buffer := bytes.NewBuffer(dex[off:])
	for i := uint32(0); i < size; i++ {
		dataOff := buffer.ReadUInt32(IS_LITTLE_ENDIAN)
		dataSize := ReadUnsignedLeb128(dex[dataOff:])
		that.StringIds[i] = &StringId{DataOff: dataOff, Size: dataSize}
		if dataSize.Val>0{
			//that.StringIds[i].Data= dex[dataOff+dataSize.Len : dataOff+dataSize.Len+dataSize.Val]
			if data,err:= NewReader(dex[ dataOff+dataSize.Len:]).ReadSlices(00);err!=nil{
				debug.LogE(err)
			}else {
				that.StringIds[i].Data=data
			}
		}
	}
	return that
}
