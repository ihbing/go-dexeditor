package dextask

import (
	"dexeditor/tool/bytes"
	"dexeditor/tool/debug"
	"dexeditor/tool/dexfile"
	"io/ioutil"
	"os"
)

type DexTask struct {
	mDexFile       *dexfile.DexFile
	mDexBuffer     *bytes.BufferV2
	mDexData       []byte
	mDexDataBackup []byte
}

func NewDexTask(dex []byte) *DexTask {
	that := &DexTask{mDexBuffer: bytes.NewBuffer(dex), mDexData: dex}
	copy(that.mDexDataBackup, that.mDexData)
	that.mDexFile = dexfile.NewDexFile(dex, 0)
	return that
}

func (that *DexTask) Reset() {
	that.mDexData = make([]byte, len(that.mDexDataBackup))
	copy(that.mDexData, that.mDexDataBackup)
}
func (that *DexTask) WriteFile(dexPath string) {
	that.mDexFile.Write(that.mDexData, 0)
	if err := ioutil.WriteFile(dexPath, that.mDexData, os.ModeAppend); err != nil {
		debug.LogE("WriteFile ", dexPath, " error:", err)
	}
}
//
//func (that *DexTask) dexBodyParse() (err error) {
//	if err := that.dexIdsParse(); err != nil {
//		return err
//	}
//
//	if err := that.dexMapListParse(); err != nil {
//		return err
//	}
//	return nil
//}
//func (that *DexTask) dexIdsParse() (err error) {
//	if err := that.dexStringIdParse(); err != nil {
//		return err
//	}
//	if err := that.dexTypeIdParse(); err != nil {
//		return err
//	}
//	if err := that.dexProtoIdParse(); err != nil {
//		return err
//	}
//	if err := that.dexFieldIdParse(); err != nil {
//		return err
//	}
//	if err := that.dexMethodIdParse(); err != nil {
//		return err
//	}
//	if err := that.dexClassDefParse(); err != nil {
//		return err
//	}
//
//	return nil
//}
//func (that *DexTask) dexStringIdParse() error {
//	dexHeader := that.mDexFile.PHeader
//	that.mDexFile.PStringIds = make([]*dexfile.StringId, dexHeader.StringIdsSize)
//	dexStringIds := that.mDexFile.PStringIds
//	for i := uint32(0); i < dexHeader.StringIdsSize; i++ {
//		off := that.mDexBuffer.ReadUInt32(true)
//		size, n := dexfile.ReadUnsignedLeb128(that.mDexData[off:])
//		buf := bytes.NewBuffer(that.mDexData[off+n:])
//		if data, err := buf.ReadSlices(00); err != nil {
//			return err
//		} else {
//			dexStringIds[i] = &dexfile.StringId{
//				DataOff: off,
//				StringData: &dexfile.DexStringData{
//					Size: uint32(size),
//					Data: data,
//				},
//			}
//		}
//	}
//	return nil
//}
//
//func (that *DexTask) dexGetStringData(stringId *dexfile.StringId) (string, error) {
//	//if err := that.mDexBuffer.Position(stringId.StringDataOff); err != nil {
//	//	return "", err
//	//}
//	//for ; ; {
//	//	if b, err := that.mDexBuffer.ReadByte(); err != nil {
//	//		debug.LogE(err)
//	//		return "", err
//	//	} else if b > 0x7f {
//	//		if err := that.mDexBuffer.UnreadByte(); err != nil {
//	//			debug.LogE(err)
//	//			return "", err
//	//		} else if line, err := that.mDexBuffer.ReadString(00); err != nil {
//	//			return "", err
//	//		} else {
//	//			return line, nil
//	//		}
//	//	}
//	//}
//	return "", nil
//}
//
//func (that *DexTask) dexTypeIdParse() error {
//	dexHeader := that.mDexFile.PHeader
//	that.mDexFile.PTypeIds = make([]*dexfile.DexTypeId, dexHeader.TypeIdsSize)
//	dexTypeIds := that.mDexFile.PTypeIds
//	for i := uint32(0); i < dexHeader.TypeIdsSize; i++ {
//		idx := that.mDexBuffer.ReadUInt32(true)
//		dexTypeIds[i] = &dexfile.DexTypeId{DescriptorIdx: idx}
//	}
//	return nil
//}
//
//func (that *DexTask) dexProtoIdParse() error {
//	dexHeader := that.mDexFile.PHeader
//	that.mDexFile.PProtoIds = make([]*dexfile.DexProtoId, dexHeader.ProtoIdsSize)
//	dexProtoIds := that.mDexFile.PProtoIds
//	for i := uint32(0); i < dexHeader.ProtoIdsSize; i++ {
//		shortyIdx := that.mDexBuffer.ReadUInt32(true)
//		returnTypeIdx := that.mDexBuffer.ReadUInt32(true)
//		parametersOff := that.mDexBuffer.ReadUInt32(true)
//		dexProtoIds[i] = &dexfile.DexProtoId{
//			ShortyIdx:     shortyIdx,
//			ReturnTypeIdx: returnTypeIdx,
//			ParametersOff: parametersOff,
//		}
//		if parametersOff > 0 {
//			buf := bytes.NewBuffer(that.mDexData)
//			buf.Next(int(parametersOff))
//			size := buf.ReadUInt32(true)
//			typeItemList := &dexfile.DexTypeList{Size: size, List: make([]*dexfile.DexTypeItem, size)}
//			for j := uint32(0); j < size; j++ {
//				if typeIdx, err := buf.ReadUInt16(true); err != nil {
//					return err
//				} else {
//					typeItemList.List[j] = &dexfile.DexTypeItem{TypeIdx: typeIdx}
//				}
//			}
//			dexProtoIds[i].TypeItemList = typeItemList
//		}
//
//	}
//	return nil
//}
//
//func (that *DexTask) dexFieldIdParse() error {
//	dexHeader := that.mDexFile.PHeader
//	that.mDexFile.PFieldIds = make([]*dexfile.DexFieldId, dexHeader.FieldIdsSize)
//	dexFieldIds := that.mDexFile.PFieldIds
//	for i := uint32(0); i < dexHeader.FieldIdsSize; i++ {
//		if classIdx, err := that.mDexBuffer.ReadUInt16(true); err != nil {
//			return err
//		} else if typeIdx, err := that.mDexBuffer.ReadUInt16(true); err != nil {
//			return err
//		} else {
//			nameIdx := that.mDexBuffer.ReadUInt32(true)
//			dexFieldIds[i] = &dexfile.DexFieldId{ClassIdx: classIdx, TypeIdx: typeIdx, NameIdx: nameIdx}
//		}
//	}
//	return nil
//}
//
//func (that *DexTask) dexMethodIdParse() error {
//	dexHeader := that.mDexFile.PHeader
//	that.mDexFile.PMethodIds = make([]*dexfile.DexMethodId, dexHeader.MethodIdsSize)
//	dexMethodIds := that.mDexFile.PMethodIds
//	for i := uint32(0); i < dexHeader.MethodIdsSize; i++ {
//		if classIdx, err := that.mDexBuffer.ReadUInt16(true); err != nil {
//			return err
//		} else if protoIdx, err := that.mDexBuffer.ReadUInt16(true); err != nil {
//			return err
//		} else {
//			nameIdx := that.mDexBuffer.ReadUInt32(true)
//			dexMethodIds[i] = &dexfile.DexMethodId{ClassIdx: classIdx, ProtoIdx: protoIdx, NameIdx: nameIdx}
//		}
//	}
//	return nil
//}
//
//func (that *DexTask) dexClassDefParse() error {
//	dexHeader := that.mDexFile.PHeader
//	that.mDexFile.PClassDefs = make([]*dexfile.ClassDef, dexHeader.ClassDefsSize)
//	dexClassDefs := that.mDexFile.PClassDefs
//	for i := uint32(0); i < dexHeader.ClassDefsSize; i++ {
//		classIdx := that.mDexBuffer.ReadUInt32(true)
//		accessFlags := that.mDexBuffer.ReadUInt32(true)
//		superclassIdx := that.mDexBuffer.ReadUInt32(true)
//		interfacesOff := that.mDexBuffer.ReadUInt32(true)
//		sourceFileIdx := that.mDexBuffer.ReadUInt32(true)
//		annotationsOff := that.mDexBuffer.ReadUInt32(true)
//		classDataOff := that.mDexBuffer.ReadUInt32(true)
//		staticValuesOff := that.mDexBuffer.ReadUInt32(true)
//		dexClassDefs[i] = &dexfile.ClassDef{
//			ClassIdx:        classIdx,
//			AccessFlags:     accessFlags,
//			SuperclassIdx:   superclassIdx,
//			InterfacesOff:   interfacesOff,
//			SourceFileIdx:   sourceFileIdx,
//			AnnotationsOff:  annotationsOff,
//			ClassDataOff:    classDataOff,
//			StaticValuesOff: staticValuesOff,
//		}
//		//classData
//		if classDataOff > 0 {
//
//		}
//		//Interfaces
//		//if interfacesOff > 0 {
//		//	buf := bytes.NewBuffer(that.mDexData)
//		//	buf.Next(int(interfacesOff))
//		//	if size, err := buf.ReadUInt32(true); err != nil {
//		//		return err
//		//	} else {
//		//		typeItemList := &dexfile.DexTypeList{Size: size, List: make([]*dexfile.DexTypeItem, size)}
//		//		for j := uint32(0); j < size; j++ {
//		//			if typeIdx, err := buf.ReadUInt16(true); err != nil {
//		//				return err
//		//			} else {
//		//				typeItemList.List[j] = &dexfile.DexTypeItem{TypeIdx: typeIdx}
//		//			}
//		//		}
//		//	}
//		//}
//		//Annotations
//		//if annotationsOff>0{
//		//	buf := bytes.NewBuffer(that.mOriginalDex)
//		//	buf.Next(int(annotationsOff))
//		//	if classAnnotationsOff, err := buf.ReadUInt32(true); err != nil {
//		//		return err
//		//	} else if fieldsSize, err := buf.ReadUInt32(true); err != nil {
//		//		return err
//		//	} else if methodsSize, err := buf.ReadUInt32(true); err != nil {
//		//		return err
//		//	} else if parametersSize, err := buf.ReadUInt32(true); err != nil {
//		//		return err
//		//	} else{
//		//		annotationsDirectoryItem:=&dexfile.DexAnnotationsDirectoryItem{
//		//			ClassAnnotationsOff: classAnnotationsOff,
//		//			FieldsSize:fieldsSize,
//		//			MethodsSize:methodsSize,
//		//			ParametersSize:parametersSize,
//		//		}
//		//		typeItemList := &dexfile.DexTypeList{Size: classAnnotationsOff, List: make([]*dexfile.DexTypeItem, classAnnotationsOff)}
//		//		for j := uint32(0); j < classAnnotationsOff; j++ {
//		//			if typeIdx, err := buf.ReadUInt16(true); err != nil {
//		//				return err
//		//			} else {
//		//				typeItemList.List[j] = &dexfile.DexTypeItem{TypeIdx: typeIdx}
//		//			}
//		//		}
//		//	}
//		//}
//	}
//	return nil
//}
//
//func (that *DexTask) dexMapListParse() error {
//	//dexHeader := that.mDexFile.PHeader
//	//that.mDexFile.PFieldIds = make([]*dexfile.DexFieldId, dexHeader.FieldIdsSize)
//	//dexFieldIds := that.mDexFile.PFieldIds
//	//for i := uint32(0); i < dexHeader.FieldIdsSize; i++ {
//	//	if classIdx, err := that.mDexBuffer.ReadUInt16(true); err != nil {
//	//		return err
//	//	} else if typeIdx, err := that.mDexBuffer.ReadUInt16(true); err != nil {
//	//		return err
//	//	} else if nameIdx, err := that.mDexBuffer.ReadUInt32(true); err != nil {
//	//		return err
//	//	} else {
//	//		dexFieldIds[i] = &dexfile.DexFieldId{ClassIdx: classIdx, TypeIdx: typeIdx, NameIdx: nameIdx}
//	//	}
//	//}
//	debug.LogNotImplement()
//	return nil
//}

func (that *DexTask) writeDexHeader() {
	err := error(nil)
	dexHeader := that.mDexFile.PHeader
	dexHeader.FileSize = uint32(len(that.mDexData))
	dexHeader.Write(that.mDexData, 0)
	if dexHeader.Signature, err = dexfile.DexSignature(that.mDexData, dexfile.SIGNATURE_OFF+dexfile.SIGNATURE_LEN); err != nil {
		debug.LogE(err)
		return
	}
	dexHeader.Write(that.mDexData, 0)
	if dexHeader.Checksum, err = dexfile.DexChecksum(that.mDexData, dexfile.CHECKSUM_OFF+dexfile.CHECKSUM_LEN); err != nil {
		debug.LogE(err)
		return
	}
	dexHeader.Write(that.mDexData, 0)

}

func (that *DexTask) updateDex() {

	that.writeDexHeader()
}
