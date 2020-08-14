package dexfile

import "dexeditor/tool/bytes"

type TypeIds struct {
	TypeIds []*TypeId
}
type TypeId struct {
	DescriptorIdx uint32 //对应StringIds编号
}

func NewTypeIds(dex []byte, off, size uint32) *TypeIds {
	that := &TypeIds{}
	that.TypeIds = make([]*TypeId, size)
	buffer := bytes.NewBuffer(dex[off:])
	for i := uint32(0); i < size; i++ {
		idx := buffer.ReadUInt32(IS_LITTLE_ENDIAN)
		that.TypeIds[i] = &TypeId{DescriptorIdx:idx}
	}
	return that
}
