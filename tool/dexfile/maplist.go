package dexfile

import "dexeditor/tool/bytes"

type MapList struct {
	MapItems []*MapItem
}
type MapItem struct {
	Type uint16 //类型
	Unused uint16 //对齐方式
	Size uint32 //大小
	Offset uint32 //偏移
}

func NewMapList(dex []byte, off uint32) *MapList {
	that := &MapList{}
	buffer := bytes.NewBuffer(dex[off:])
	size:=buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.MapItems = make([]*MapItem, size)
	for i := uint32(0); i < size; i++ {
		that.MapItems[i]=new(MapItem)
		that.MapItems[i].Type=buffer.ReadUInt16(IS_LITTLE_ENDIAN)
		that.MapItems[i].Unused=buffer.ReadUInt16(IS_LITTLE_ENDIAN)
		that.MapItems[i].Size=buffer.ReadUInt32(IS_LITTLE_ENDIAN)
		that.MapItems[i].Offset=buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	}
	return that
}
