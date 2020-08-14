package dexfile

import "dexeditor/tool/bytes"

type Writer struct {
	mData []byte
	mOff int
}

func NewWriter(data []byte) *Writer {
	return &Writer{mData: data}
}
func (that *Writer)WriteUnsignedLeb128(leb128 *ULeb128)  {
	WriteUnsignedLeb128(that.mData[that.mOff:],leb128.Val)
	that.mOff+=int(leb128.Len)
}
func (that *Writer)Write(data []byte)  {
	copy(that.mData[that.mOff:],data)
	that.mOff+= len(data)
}
func (that *Writer)WriteUint32(data uint32,isLittleEndian bool)  {
	copy(that.mData[that.mOff:], bytes.Uint32ToBytes(data,isLittleEndian))
	that.mOff+=4
}

func (that *Writer) WriteUint16(data uint16, isLittleEndian bool) {
	copy(that.mData[that.mOff:], bytes.Uint16ToBytes(data,isLittleEndian))
	that.mOff+=2
}

