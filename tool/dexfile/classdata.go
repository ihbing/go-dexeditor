package dexfile

type ClassDataHeader struct {
	StaticFieldsSize   *ULeb128
	InstanceFieldsSize *ULeb128
	DirectMethodsSize  *ULeb128
	VirtualMethodsSize *ULeb128
}

type ClassData struct {
	Header         *ClassDataHeader
	StaticFields   []*EncodedField
	InstanceFields []*EncodedField
	DirectMethods  []*EncodedMethod
	VirtualMethods []*EncodedMethod
}

func NewClassData(dex []byte, off uint32) *ClassData {
	that := new(ClassData)
	that.parse(dex, off)
	return that
}

func (that *ClassData) parse(dex []byte, off uint32) {
	buffer := NewReader(dex[off:])
	that.Header = new(ClassDataHeader)
	that.Header.StaticFieldsSize = buffer.ReadUnsignedLeb128()
	that.Header.InstanceFieldsSize = buffer.ReadUnsignedLeb128()
	that.Header.DirectMethodsSize = buffer.ReadUnsignedLeb128()
	that.Header.VirtualMethodsSize = buffer.ReadUnsignedLeb128()
	that.StaticFields = make([]*EncodedField, that.Header.StaticFieldsSize.Val)
	for i := 0; i < len(that.StaticFields); i++ {
		that.StaticFields[i] = NewEncodedField(buffer)
	}
	that.InstanceFields = make([]*EncodedField, that.Header.InstanceFieldsSize.Val)
	for i := 0; i < len(that.InstanceFields); i++ {
		that.InstanceFields[i] = NewEncodedField(buffer)
	}
	that.DirectMethods = make([]*EncodedMethod, that.Header.DirectMethodsSize.Val)
	for i := 0; i < len(that.DirectMethods); i++ {
		that.DirectMethods[i] = NewEncodedMethod(buffer,dex)
	}
	that.VirtualMethods = make([]*EncodedMethod, that.Header.VirtualMethodsSize.Val)
	for i := 0; i < len(that.VirtualMethods); i++ {
		that.VirtualMethods[i] = NewEncodedMethod(buffer,dex)
	}
}
func (that *ClassData) Write(dex []byte, off uint32) {
	writer := NewWriter(dex[off:])
	writer.WriteUnsignedLeb128(that.Header.StaticFieldsSize)
	writer.WriteUnsignedLeb128(that.Header.InstanceFieldsSize)
	writer.WriteUnsignedLeb128(that.Header.DirectMethodsSize)
	writer.WriteUnsignedLeb128(that.Header.VirtualMethodsSize)
	for i := uint32(0); i < that.Header.StaticFieldsSize.Val; i++ {
		that.StaticFields[i].Write(writer)
	}
	for i := uint32(0); i < that.Header.InstanceFieldsSize.Val; i++ {
		that.InstanceFields[i].Write(writer)
	}
	for i := uint32(0); i < that.Header.DirectMethodsSize.Val; i++ {
		that.DirectMethods[i].Write(writer,dex)
	}
	for i := uint32(0); i < that.Header.VirtualMethodsSize.Val; i++ {
		that.VirtualMethods[i].Write(writer,dex)
	}

}
