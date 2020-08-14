package dexfile

import "dexeditor/tool/bytes"

type ClassDefs struct {
	ClassDefs []*ClassDef
	size uint32
}

type ClassDef struct {
	ClassIdx        uint32 //typeIds中索引
	AccessFlags     uint32
	SuperclassIdx   uint32
	InterfacesOff   uint32
	Interfaces      *DexTypeList
	SourceFileIdx   uint32
	AnnotationsOff  uint32
	Annotations     *DexAnnotationsDirectoryItem
	ClassDataOff    uint32 //file offset to class_data_item
	ClassData       *ClassData
	StaticValuesOff uint32 //file offset to DexEncodedArray
	StaticValues    *StaticValues
}
func NewClassDefs(dex []byte,off,size uint32) *ClassDefs {
	that := &ClassDefs{ClassDefs:make([]*ClassDef, size),size:size}
	buffer := bytes.NewBuffer(dex[off:])
	for i := uint32(0); i < size; i++ {
		that.ClassDefs[i]=new(ClassDef)
		that.ClassDefs[i].ClassIdx = buffer.ReadUInt32(true)
		that.ClassDefs[i].AccessFlags = buffer.ReadUInt32(true)
		that.ClassDefs[i].SuperclassIdx = buffer.ReadUInt32(true)
		that.ClassDefs[i].InterfacesOff = buffer.ReadUInt32(true)
		that.ClassDefs[i].SourceFileIdx = buffer.ReadUInt32(true)
		that.ClassDefs[i].AnnotationsOff = buffer.ReadUInt32(true)
		that.ClassDefs[i].ClassDataOff = buffer.ReadUInt32(true)
		that.ClassDefs[i].StaticValuesOff = buffer.ReadUInt32(true)
		if that.ClassDefs[i].ClassDataOff > 0 {
			that.ClassDefs[i].ClassData = NewClassData(dex, that.ClassDefs[i].ClassDataOff)
		}
		if that.ClassDefs[i].StaticValuesOff > 0 {
			that.ClassDefs[i].StaticValues = NewStaticValues(dex, that.ClassDefs[i].StaticValuesOff)
		}

	}
	return that
}
func (that *ClassDefs) Write(dex []byte, off uint32) {
	writer :=NewWriter(dex[off:])
	for i:=uint32(0);i<that.size ;i++  {
		classDef:=that.ClassDefs[i]
		writer.WriteUint32(classDef.ClassIdx,IS_LITTLE_ENDIAN)
		writer.WriteUint32(classDef.AccessFlags,IS_LITTLE_ENDIAN)
		writer.WriteUint32(classDef.SuperclassIdx,IS_LITTLE_ENDIAN)
		writer.WriteUint32(classDef.InterfacesOff,IS_LITTLE_ENDIAN)
		writer.WriteUint32(classDef.SourceFileIdx,IS_LITTLE_ENDIAN)
		writer.WriteUint32(classDef.AnnotationsOff,IS_LITTLE_ENDIAN)
		writer.WriteUint32(classDef.ClassDataOff,IS_LITTLE_ENDIAN)
		if classDef.ClassDataOff>0{
			classDef.ClassData.Write(dex,classDef.ClassDataOff)
		}
		writer.WriteUint32(classDef.StaticValuesOff,IS_LITTLE_ENDIAN)
		if classDef.StaticValuesOff>0{
			//暂时不做处理
		}
	}
}
