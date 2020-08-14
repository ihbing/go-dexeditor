package dexfile

type StaticValues struct {
	size        *ULeb128
	StaticValue []*StaticValue
}
type StaticValue struct {

}
func NewStaticValues(dex []byte,off uint32) *StaticValues {
	buffer:= NewReader(dex[off:])
	that:= &StaticValues{size: buffer.ReadUnsignedLeb128()}
	that.StaticValue =make([]*StaticValue,that.size.Val)
	return that
}
