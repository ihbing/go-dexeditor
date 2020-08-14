package dexfile

import "dexeditor/tool/bytes"

type ProtoIds struct {
	ProtoIds []*ProtoId
}
type Parameters struct {
	Size     uint32
	TypeIdxs []uint16
}
type ProtoId struct {
	ShortyIdx     uint32      //stringIds编号
	ReturnTypeIdx uint32      //typeIds编号
	ParametersOff uint32      //参数偏移位置
	Parameters    *Parameters //参数偏移位置
}

func NewProtoIds(dex []byte, off, size uint32) *ProtoIds {
	that := &ProtoIds{}
	that.ProtoIds = make([]*ProtoId, size)
	buffer := bytes.NewBuffer(dex[off:])
	for i := uint32(0); i < size; i++ {
		shortyIdx := buffer.ReadUInt32(IS_LITTLE_ENDIAN)
		rtIdx := buffer.ReadUInt32(IS_LITTLE_ENDIAN)
		pOff := buffer.ReadUInt32(IS_LITTLE_ENDIAN)
		that.ProtoIds[i] = &ProtoId{ShortyIdx: shortyIdx, ReturnTypeIdx: rtIdx, ParametersOff: pOff}
		if pOff > 0 {
			buffer1 := bytes.NewBuffer(dex[pOff:])
			pSize := buffer1.ReadUInt32(IS_LITTLE_ENDIAN)
			pts := make([]uint16, pSize)
			for j := uint32(0); j < pSize; j++ {
				pts[j] = buffer1.ReadUInt16(IS_LITTLE_ENDIAN)
			}
			that.ProtoIds[i].Parameters = &Parameters{Size: pSize, TypeIdxs: pts,}
		}
	}
	return that
}
