package dexfile

type U1 uint8
type U2 uint16
type U4 uint32
type U8 uint64
type ULeb128 struct {
	Val uint32
	Len uint32
}

func NewULeb128(num uint32) *ULeb128 {
	data:=ULeb128encode(uint64(num))
	return &ULeb128{num, uint32(len(data))}
}
//参考https://berryjam.github.io/2019/09/LEB128(Little-Endian-Base-128)%E6%A0%BC%E5%BC%8F%E4%BB%8B%E7%BB%8D/
//参考https://www.androidos.net.cn/android/4.0.2_r1/xref/dalvik/libdex/Leb128.h
//func ReadUnsignedLeb128(data []byte)(uleb128 *ULeb128) {
//	uleb128=new(ULeb128)
//	index :=uint32(0)
//	uleb128.Len=0
//	uleb128.Val=uint32(data[index])
//	index++
//	uleb128.Len++
//	if uleb128.Val>0x7f{
//		cur:=uint32(data[index])
//		index++
//		uleb128.Val=(uleb128.Val&0x7f)|((cur&0x7f)<<7)
//		uleb128.Len++
//		if cur>0x7f{
//			cur=uint32(data[index])
//			index++
//			uleb128.Val|=(cur&0x7f)<<24
//			uleb128.Len++
//			if cur>0x7f{
//				cur=uint32(data[index])
//				index++
//				uleb128.Val|=(cur&0x7f)<<21
//				uleb128.Len++
//				if cur>0x7f{
//					cur=uint32(data[index])
//					index++
//					uleb128.Val|=cur<<28
//					uleb128.Len++
//				}
//			}
//		}
//	}
//	return
//}
func WriteUnsignedLeb128(data []byte,num uint32){
	if num == 0 {
		data[0]=0
	} else {
		for i:=0;num != 0;i++ {
			b := (byte)(num & 0x7F)
			num >>= 7
			if num != 0 { /* more bytes to come */
				b |= 0x80
			}
			data[i]=b
		}
	}
}
func ULeb128encode(num uint64) []byte {
	res := []byte{}

	if num == 0 {
		res = append(res, 0)
	} else {
		for num != 0 {
			b := (byte)(num & 0x7F)
			num >>= 7
			if num != 0 { /* more bytes to come */
				b |= 0x80
			}
			res = append(res, b)
		}
	}

	return res
}

func ReadUnsignedLeb128(data []byte)(uleb128 *ULeb128){
	if len(data) == 0 {
		panic("illegal input")
	}
	uleb128=new(ULeb128)
	uleb128.Val = 0
	uleb128.Len = 1
	for {
		flag := data[uleb128.Len-1] & 0x80
		low7bit := data[uleb128.Len-1] & 0x7F
		uleb128.Val |= uint32(low7bit) << (7 * (uleb128.Len-1))
		if flag != 0 {
			uleb128.Len++
		} else {
			break
		}
	}

	return
}
func ULeb128decode(bytes []byte) uint64 {
	if len(bytes) == 0 {
		panic("illegal input")
	}
	var res uint64 = 0
	var i uint8 = 0
	for {
		flag := bytes[i] & 0x80
		low7bit := bytes[i] & 0x7F
		res |= uint64(low7bit) << (7 * i)
		if flag != 0 {
			i++
		} else {
			break
		}
	}

	return res
}

func SLeb128encode(value int64) []byte {
	res := []byte{}

	more := 1

	for more != 0 {
		b := (byte)(value & 0x7F)
		signFlag := (byte)(value & 0x40)
		value >>= 7
		if (value == 0 && signFlag == 0) ||  // 正数
			(value == -1 && signFlag != 0) { // 负数
			more = 0
		} else {
			b |= 0x80
		}
		res = append(res, b)
	}

	return res
}

func SLeb128decode(bytes []byte) int64 {
	if len(bytes) == 0 {
		panic("illegal input")
	}
	var res uint64 = 0
	var i uint8 = 0
	isNegative := false
	var shift uint64 = 0
	for {
		flag := bytes[i] & 0x80
		low7bit := bytes[i] & 0x7F
		res |= uint64(low7bit) << (shift)
		shift+=7
		if flag != 0 {
			i++
		} else {
			signFlag := bytes[i] & 0x40
			if signFlag != 0 {
				isNegative = true
			}
			break
		}
	}
	if !isNegative {
		return int64(res)
	} else {
		tmp := int64(res)
		tmp |= -(1 << shift)
		return tmp
	}
}