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