package dextask

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNewDexTask(t *testing.T) {

}
func TestDexTask_Parse(t *testing.T) {
	if bin, err := ioutil.ReadFile("test.dex"); err != nil {
		t.Error(err)
	} else {
		dexTask := NewDexTask(bin)
		classDefs := dexTask.mDexFile.PClassDefs.ClassDefs
		stringIds := dexTask.mDexFile.PStringIds.StringIds
		methodIds := dexTask.mDexFile.PMethodIds.MethodIds
		typeIds := dexTask.mDexFile.PTypeIds.TypeIds
		for _, v := range classDefs {
			className:=string(stringIds[typeIds[v.ClassIdx].DescriptorIdx].Data)
			if className=="Lformatfa/xposed/Fdex2/MainActivity;"{
				fmt.Println("className:",className)
				lastIndex:=uint32(0)
				for _, vm := range v.ClassData.VirtualMethods {
					lastIndex+=vm.MethodIdx.Val
					methodName := string(stringIds[methodIds[lastIndex].NameIdx].Data)
					if methodName=="onCreate" {
						vm.AccessFlags.Val = uint32(2)
						vm.CodeData.Insns=make([]uint16,vm.CodeData.InsnsSize)
						fmt.Println("native method:",methodName)
					}
				}
				//for _, vm := range v.ClassData.DirectMethods {
				//	methodName := string(stringIds[methodIds[vm.MethodIdxDiff.Val].NameIdx].Data)
				//	if strings.Contains(methodName, "onCreate") {
				//		vm.AccessFlags.Val |= 0x00000100
				//		fmt.Println("native method:",methodName)
				//	}
				//}
			}


			//className := string(stringIds[typeIds[v.ClassIdx].DescriptorIdx].Data)
			//fmt.Println("className:", className)
			////将类名包含Main的类所有函数隐藏
			//if strings.Contains(className, "Main") {
			//	v.ClassData.Header.VirtualMethodsSize = dexfile.NewULeb128(0)
			//	v.ClassData.Header.DirectMethodsSize = dexfile.NewULeb128(0)
			//}
		}
		dexTask.WriteFile("new.dex")
	}
}
