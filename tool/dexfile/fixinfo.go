package dexfile

import "dexeditor/tool/debug"

type FixInfos struct {
	FixInfos []*FixInfo
}

const FixInfoMagic uint32 = 0x09050207
type FixInfo struct {
	Off  uint32
	Data []byte
}
func NewFixInfos(dex []byte,off uint32) *FixInfos {
	that:= &FixInfos{}
	reader:=NewReader(dex[off:])
	magic:=reader.ReadUInt32(IS_LITTLE_ENDIAN)
	if magic!=FixInfoMagic {
		debug.LogE("this dex has not fix info.")
		return that
	}
	size:=reader.ReadUInt32(IS_LITTLE_ENDIAN)
	that.FixInfos=make([]*FixInfo,size)
	for i := uint32(0); i < size; i++ {
		that.FixInfos[i]=new(FixInfo)
		that.FixInfos[i].Off =reader.ReadUInt32(IS_LITTLE_ENDIAN)
		if data,err:=reader.ReadBytes(that.FixInfos[i].Off);err!=nil{
			debug.LogE(err)
		}else {
			that.FixInfos[i].Data=data
		}
	}
	return that
}