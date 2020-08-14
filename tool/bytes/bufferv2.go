package bytes

import (
	"bytes"
	"dexeditor/tool/debug"
	"encoding/binary"
)

type BufferV2 struct {
	*bytes.Buffer

}

func (that *BufferV2) Read(data []byte)error{
	if _,err:= that.Buffer.Read(data);err!=nil{
		return err
	}
	return nil
}

func (that *BufferV2) ReadBytes(len uint32)([]byte,error){
	result:=make([]byte,0,len)
	for ; len>0;  {
		if len<1024{
			buf:=make([]byte,len)
			if n,err:= that.Buffer.Read(buf);err!=nil{
				return nil,err
			}else {
				len=len-uint32(n)
				result = append(result, buf[:n]...)
			}
		}else {
			buf:=make([]byte,1024)
			if n,err:= that.Buffer.Read(buf);err!=nil{
				return nil,err
			}else {
				len=len-uint32(n)
				result = append(result, buf[:n]...)
			}
		}
	}
	return result,nil
}
func (that *BufferV2) ReadUInt32(isLittleEndian bool)uint32{
	buf:=make([]byte,4)
	if _,err:= that.Buffer.Read(buf);err!=nil{
		debug.LogE(err)
		return 0
	}else if isLittleEndian {
		return binary.LittleEndian.Uint32(buf)
	}else {
		return binary.BigEndian.Uint32(buf)
	}
}
func (that *BufferV2) WriteUInt32(data []byte,isLittleEndian bool)(uint32,error){
	buf:=make([]byte,4)
	if _,err:= that.Buffer.Read(buf);err!=nil{
		return 0,err
	}else if isLittleEndian {
		return binary.LittleEndian.Uint32(buf),nil
	}else {
		return binary.BigEndian.Uint32(buf),nil
	}
}


func (that *BufferV2) ReadUInt16(isLittleEndian bool) uint16 {
	buf:=make([]byte,2)
	if _,err:= that.Buffer.Read(buf);err!=nil{
		debug.LogE(err)
		return 0
	}else if isLittleEndian {
		return binary.LittleEndian.Uint16(buf)
	}else {
		return binary.BigEndian.Uint16(buf)
	}
}

func (that *BufferV2) ReadSlices(delim byte) ([]byte, error) {
	buf:=make([]byte,0)
	for ; ;  {
		if b,err:=that.ReadByte();err!=nil{
			return nil,err
		}else if b!=delim{
			buf=append(buf,b)
		}else {
			break
		}
	}
	return buf,nil
}


func NewBuffer(data []byte) *BufferV2 {
	return &BufferV2{Buffer: bytes.NewBuffer(data)}
}
