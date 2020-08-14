package dexfile

type EncodedMethod  struct {
	MethodIdxDiff *ULeb128
	AccessFlags   *ULeb128
	CodeOff       *ULeb128
}

func NewEncodedMethod(reader *Reader) *EncodedMethod  {
	that:= &EncodedMethod{}
	that.MethodIdxDiff =reader.ReadUnsignedLeb128()
	that.AccessFlags =reader.ReadUnsignedLeb128()
	that.CodeOff =reader.ReadUnsignedLeb128()
	return that
}

func (that *EncodedMethod )Write(writer *Writer)  {
	writer.WriteUnsignedLeb128(that.MethodIdxDiff)
	writer.WriteUnsignedLeb128(that.AccessFlags)
	writer.WriteUnsignedLeb128(that.CodeOff)
}