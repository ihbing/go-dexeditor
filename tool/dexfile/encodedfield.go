package dexfile

type EncodedField struct {
	FieldIdxDiff *ULeb128
	AccessFlags  *ULeb128
}

func NewEncodedField(reader *Reader) *EncodedField {
	that:= &EncodedField{}
	that.FieldIdxDiff =reader.ReadUnsignedLeb128()
	that.AccessFlags =reader.ReadUnsignedLeb128()
	return that
}

func (that *EncodedField)Write(writer *Writer)  {
	writer.WriteUnsignedLeb128(that.FieldIdxDiff)
	writer.WriteUnsignedLeb128(that.AccessFlags)
}