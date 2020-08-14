package dexfile

type CodeData struct {
	RegistersSize uint16
	InsSize       uint16
	OutsSize      uint16
	TriesSize     uint16
	DebugInfoOff  uint32
	InsnsSize     uint32
	Insns         []uint16
}


func NewCodeData(dex []byte, off uint32) *CodeData {
	that := new(CodeData)
	reader:=NewReader(dex[off:])
	that.RegistersSize=reader.ReadUInt16(IS_LITTLE_ENDIAN)
	that.InsSize=reader.ReadUInt16(IS_LITTLE_ENDIAN)
	that.OutsSize=reader.ReadUInt16(IS_LITTLE_ENDIAN)
	that.TriesSize=reader.ReadUInt16(IS_LITTLE_ENDIAN)
	that.DebugInfoOff=reader.ReadUInt32(IS_LITTLE_ENDIAN)
	that.InsnsSize=reader.ReadUInt32(IS_LITTLE_ENDIAN)
	that.Insns=make([]uint16,that.InsnsSize)
	for i := uint32(0); i < that.InsnsSize; i++ {
		that.Insns[i]=reader.ReadUInt16(IS_LITTLE_ENDIAN)
	}
	return that
}
func (that *CodeData)Write(dex []byte,off uint32)  {
	writer:=NewWriter(dex[off:])
	writer.WriteUint16(that.RegistersSize,IS_LITTLE_ENDIAN)
	writer.WriteUint16(that.InsSize,IS_LITTLE_ENDIAN)
	writer.WriteUint16(that.OutsSize,IS_LITTLE_ENDIAN)
	writer.WriteUint16(that.TriesSize,IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.DebugInfoOff,IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.InsnsSize,IS_LITTLE_ENDIAN)
	for i := uint32(0); i < that.InsnsSize; i++ {
		writer.WriteUint16(that.Insns[i],IS_LITTLE_ENDIAN)
	}
}
