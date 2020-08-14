package dexfile

import (
	"dexeditor/tool/bytes"
	"dexeditor/tool/debug"
)

type Reader struct {
	*bytes.BufferV2
}
func NewReader(data []byte) *Reader {
	return &Reader{BufferV2: bytes.NewBuffer(data)}
}

func (that *Reader) ReadUnsignedLeb128() *ULeb128 {
	result:=ReadUnsignedLeb128(that.Bytes())
	if _, err:= that.ReadBytes(result.Len);err!=nil{
		debug.LogE(err)
	}
	return result
}
