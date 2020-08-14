package dexfile

//    const DexHeader*    pHeader;
//    const StringId*  pStringIds;
//    const DexTypeId*    pTypeIds;
//    const DexFieldId*   pFieldIds;
//    const DexMethodId*  pMethodIds;
//    const DexProtoId*   pProtoIds;
//    const ClassDef*  pClassDefs;
//    const DexLink*      pLinkData;
//    const DexClassLookup* pClassLookup;
//    const void*         pRegisterMapPool;       // RegisterMapClassPool
//
//    /* points to start of DEX file data */
//    const uint8*           baseAddr;
//
//    /* track memory overhead for auxillary structures */
//    int                 overhead;
type DexFile struct {
	PHeader    *DexHeader
	PStringIds *StringIds
	PTypeIds   *TypeIds
	PFieldIds  *FieldIds
	PMethodIds *MethodIds
	PProtoIds  *ProtoIds
	PClassDefs *ClassDefs
	PMapList   *MapList
}

func NewDexFile(dex []byte,off uint32) *DexFile {
	that := &DexFile{}
	that.PHeader = NewDexHeader(dex, off)
	that.PStringIds=NewStringIds(dex,that.PHeader.StringIdsOff, that.PHeader.StringIdsSize)
	that.PTypeIds=NewTypeIds(dex, that.PHeader.TypeIdsOff,that.PHeader.TypeIdsSize)
	that.PFieldIds=NewFieldIds(dex,that.PHeader.FieldIdsOff,that.PHeader.FieldIdsSize)
	that.PMethodIds=NewMethodIds(dex,that.PHeader.MethodIdsOff,that.PHeader.MethodIdsSize)
	that.PProtoIds=NewProtoIds(dex, that.PHeader.ProtoIdsOff, that.PHeader.ProtoIdsSize)
	that.PClassDefs=NewClassDefs(dex, that.PHeader.ClassDefsOff, that.PHeader.ClassDefsSize)
	that.PMapList=NewMapList(dex,that.PHeader.MapOff)
	return that
}
func (that *DexFile) Write(dex []byte, off uint32) {
	that.PClassDefs.Write(dex,off+that.PHeader.ClassDefsOff)
	//dex修改后需要修复校验
	that.writeFixHeader(dex,off)
}

func (that *DexFile) writeFixHeader(dex []byte, off uint32) {
	//修改dex长度
	that.PHeader.FileSize= uint32(len(dex))
	that.PHeader.Write(dex,off)
	//修复signature校验
	newSignature:=DexSignature(dex,SIGNATURE_OFF+SIGNATURE_LEN)
	that.PHeader.Signature=newSignature
	that.PHeader.Write(dex,off)
	//修复checksum校验
	newChecksum:=DexChecksum(dex,CHECKSUM_OFF+CHECKSUM_LEN)
	that.PHeader.Checksum=newChecksum
	that.PHeader.Write(dex,off)
}

type DexMapItem struct {
	Type_  uint16
	Unused uint16
	Size   uint32
	Offset uint32
}
type DexMapList struct {
	Size uint32
	List [1]*DexMapItem
}

type DexTypeId struct {
	DescriptorIdx uint32
}
type DexFieldId struct {
	ClassIdx uint16
	TypeIdx  uint16
	NameIdx  uint32
}

type DexMethodId struct {
	ClassIdx uint16
	ProtoIdx uint16
	NameIdx  uint32
}

type DexProtoId struct {
	ShortyIdx     uint32
	ReturnTypeIdx uint32
	ParametersOff uint32
	TypeItemList  *DexTypeList
}


type DexTypeItem struct {
	TypeIdx uint16
}
type DexTypeList struct {
	Size uint32
	List []*DexTypeItem
}

type DexCode struct {
	RegistersSize uint16
	InsSize       uint16
	OutsSize      uint16
	TriesSize     uint16
	DebugInfoOff  uint32
	InsnsSize     uint32
	Insns         []uint16
}
type DexTry struct {
	StartAddr  uint32
	InsnCount  uint16
	HandlerOff uint16
}
type DexLink struct {
	Bleargh uint8
}

type DexAnnotationsDirectoryItem struct {
	ClassAnnotationsOff      uint32
	ClassAnnotations         *DexAnnotationSetItem
	FieldsSize               uint32
	MethodsSize              uint32
	ParametersSize           uint32
	FieldAnnotationsItem     []*MethodAnnotation
	MethodAnnotations        []*MethodAnnotation
	ParameterAnnotationsItem []*DexParameterAnnotationsItem
}

//struct method_annotation method_annotations

type MethodAnnotation struct {
	MethodIdx         uint32
	AnnotationsOff    uint32
	MethodAnnotations []*DexAnnotationSetItem
}

//struct annotation_set_item method_annotations
type DexParameterAnnotationsItem struct {
	MethodIdx      uint32
	AnnotationsOff uint32
}
type DexAnnotationSetRefItem struct {
	AnnotationsOff uint32
}
type DexAnnotationSetRefList struct {
	Size uint32
	List *DexAnnotationSetRefItem
}

type DexAnnotationSetItem struct {
	Size    uint32
	Entries []*AnnotationOffItem //offset to DexAnnotationItem struct annotation_off_item entries
}
type AnnotationOffItem struct {
	AnnotationOff uint32
	Item          []*DexAnnotationItem
}

//NOTE: this structure is byte-aligned.
type DexAnnotationItem struct {
	Visibility uint8
	Annotation *EncodedAnnotation // data in encoded_annotation format
}
type EncodedAnnotation struct {
	TypeIdx  uint32
	Size     uint32
	elements []*AnnotationElement
}
type AnnotationElement struct {
	NameIdx *ULeb128
	Values  *EncodedValue
}
type EncodedValue struct {
	ValueType uint8
	ValueArg  uint8
	Array     *StaticValues
}


//NOTE: this structure is byte-aligned.
type DexEncodedArray struct {
	Array []uint8 //data in encoded_array format
}
