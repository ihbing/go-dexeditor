package dexfile

const IS_LITTLE_ENDIAN bool = true
const KSHA1DigestLen uint32 = 20
const KSHA1DigestOutputLen = KSHA1DigestLen*2 + 1
const KDexEndianConstant = 0x12345678
const KDexNoIndex = 0xffffffff

const MAGIC_LEN = 0x0008

const CHECKSUM_OFF = 0x0008
const CHECKSUM_LEN = 0x0004

const SIGNATURE_OFF = 0x000C
const SIGNATURE_LEN = 0x0014

const UINT_LEN = 0x0004
const USHORT_LEN = 0x0002

const MAP_OFF_OFF = MAGIC_LEN + UINT_LEN + SIGNATURE_LEN + UINT_LEN*5
const FILE_SIZE_OFF = MAGIC_LEN + UINT_LEN + SIGNATURE_LEN

const MAP_ITEM_LEN = 0x000C
