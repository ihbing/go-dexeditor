package dexfile

import "dexeditor/tool/bytes"

type MethodIds struct {
	MethodIds []*MethodId
}
type MethodId struct {
	ClassIdx uint16 //typeIds里索引
	ProtoIdx uint16 //protoIds里索引
	NameIdx  uint32 //stringIds里索引
}

func NewMethodIds(dex []byte, off, size uint32) *MethodIds {
	that := &MethodIds{}
	that.MethodIds = make([]*MethodId, size)
	buffer := bytes.NewBuffer(dex[off:])
	for i := uint32(0); i < size; i++ {
		classIdx := buffer.ReadUInt16(IS_LITTLE_ENDIAN)
		protoIdx := buffer.ReadUInt16(IS_LITTLE_ENDIAN)
		nameIdx := buffer.ReadUInt32(IS_LITTLE_ENDIAN)
		that.MethodIds[i] = &MethodId{ClassIdx: classIdx, ProtoIdx: protoIdx, NameIdx: nameIdx}
	}
	return that
}
