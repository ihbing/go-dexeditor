package dextask

import (
	"dexeditor/tool/dexfile"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestNewDexTask(t *testing.T) {

}
func TestDexTask_Parse(t *testing.T) {
	if bin,err:=ioutil.ReadFile("test.dex");err!=nil{
		t.Error(err)
	}else{
		dexTask:=NewDexTask(bin)
		classDefs := dexTask.mDexFile.PClassDefs.ClassDefs
		stringIds := dexTask.mDexFile.PStringIds.StringIds
		typeIds := dexTask.mDexFile.PTypeIds.TypeIds
		for _, v := range classDefs {
			className:=string(stringIds[typeIds[v.ClassIdx].DescriptorIdx].Data)
			fmt.Println("className:",className)
			//将类名包含Main的类所有函数隐藏
			if strings.Contains(className,"Main"){
				v.ClassData.Header.VirtualMethodsSize=dexfile.NewULeb128(0)
				v.ClassData.Header.DirectMethodsSize=dexfile.NewULeb128(0)
			}
		}
		dexTask.WriteFile("new.dex")
	}
}