package dexfile

import "dexeditor/tool/bytes"

type FieldIds struct {
	FieldIds []*FieldId
}
type FieldId struct {
	ClassIdx uint16 //typeIds里索引
	TypeIdx  uint16//typeIds里索引
	NameIdx  uint32//stringIds里索引
}

func NewFieldIds(dex []byte, off, size uint32) *FieldIds {
	that := &FieldIds{}
	that.FieldIds = make([]*FieldId, size)
	buffer := bytes.NewBuffer(dex[off:])
	for i := uint32(0); i < size; i++ {
		classIdx := buffer.ReadUInt16(IS_LITTLE_ENDIAN)
		typeIdx := buffer.ReadUInt16(IS_LITTLE_ENDIAN)
		nameIdx := buffer.ReadUInt32(IS_LITTLE_ENDIAN)
		that.FieldIds[i] = &FieldId{ClassIdx: classIdx, TypeIdx: typeIdx, NameIdx: nameIdx}
	}
	return that
}
