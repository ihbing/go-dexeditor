package dexfile

import (
	"crypto/sha1"
	"dexeditor/tool/bytes"
	"dexeditor/tool/debug"
	"hash/adler32"
)

//struct DexHeader {
//    U1  magic[8];           /* includes version number */
//    U4  checksum;           /* adler32 checksum */
//    U1  signature[kSHA1DigestLen]; /* SHA-1 hash */
//    U4  fileSize;           /* length of entire file */
//    U4  headerSize;         /* offset to start of next section */
//    U4  endianTag;
//    U4  linkSize;
//    U4  linkOff;
//    U4  mapOff;
//    U4  stringIdsSize;
//    U4  stringIdsOff;
//    U4  typeIdsSize;
//    U4  typeIdsOff;
//    U4  protoIdsSize;
//    U4  protoIdsOff;
//    U4  fieldIdsSize;
//    U4  fieldIdsOff;
//    U4  methodIdsSize;
//    U4  methodIdsOff;
//    U4  classDefsSize;
//    U4  classDefsOff;
//    U4  dataSize;
//    U4  dataOff;
//};
type DexHeader struct {
	Magic         []uint8 //8 len
	Checksum      uint32
	Signature     []uint8 //KSHA1DigestLen 20
	FileSize      uint32
	HeaderSize    uint32
	EndianTag     uint32
	LinkSize      uint32
	LinkOff       uint32
	MapOff        uint32
	StringIdsSize uint32
	StringIdsOff  uint32
	TypeIdsSize   uint32
	TypeIdsOff    uint32
	ProtoIdsSize  uint32
	ProtoIdsOff   uint32
	FieldIdsSize  uint32
	FieldIdsOff   uint32
	MethodIdsSize uint32
	MethodIdsOff  uint32
	ClassDefsSize uint32
	ClassDefsOff  uint32
	DataSize      uint32
	DataOff       uint32
}

func NewDexHeader(dex []byte, off uint32) *DexHeader {
	that := &DexHeader{}
	if err := that.parse(dex, off); err != nil {
		debug.LogE(err)
	}
	return that
}
func (that *DexHeader) parse(dex []byte, off uint32) (err error) {
	buffer := bytes.NewBuffer(dex[off:])
	if that.Magic, err = buffer.ReadBytes(8); err != nil {
		return err
	}
	that.Checksum = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	if that.Signature, err = buffer.ReadBytes(KSHA1DigestLen); err != nil {
		return err
	}
	that.FileSize = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.HeaderSize = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.EndianTag = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.LinkSize = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.LinkOff = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.MapOff = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.StringIdsSize = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.StringIdsOff = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.TypeIdsSize = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.TypeIdsOff = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.ProtoIdsSize = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.ProtoIdsOff = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.FieldIdsSize = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.FieldIdsOff = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.MethodIdsSize = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.MethodIdsOff = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.ClassDefsSize = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.ClassDefsOff = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.DataSize = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	that.DataOff = buffer.ReadUInt32(IS_LITTLE_ENDIAN)
	return
}
func (that *DexHeader) Write(dex []byte, off uint32) {
	writer := NewWriter(dex[off:])
	writer.Write(that.Magic)
	writer.WriteUint32(that.Checksum, IS_LITTLE_ENDIAN)
	writer.Write(that.Signature)
	writer.WriteUint32(that.FileSize, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.HeaderSize, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.EndianTag, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.LinkSize, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.LinkOff, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.MapOff, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.StringIdsSize, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.StringIdsOff, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.TypeIdsSize, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.TypeIdsOff, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.ProtoIdsSize, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.ProtoIdsOff, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.FieldIdsSize, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.FieldIdsOff, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.MethodIdsSize, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.MethodIdsOff, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.ClassDefsSize, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.ClassDefsOff, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.DataSize, IS_LITTLE_ENDIAN)
	writer.WriteUint32(that.DataOff, IS_LITTLE_ENDIAN)

	return
}

func DexSignature(dex []byte, off int)(signature []byte){
	s1 := sha1.New()
	if _, err := s1.Write(dex[off:]); err != nil {
		debug.LogE(err)
		return
	}
	signature = s1.Sum(nil)
	return
}
func DexChecksum(dex []byte, off int) (checksum uint32){
	return adler32.Checksum(dex[off:])
}
