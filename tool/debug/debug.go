package debug

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

import "github.com/sirupsen/logrus"

func initSetting() {
	logrus.SetLevel(logrus.DebugLevel)
}
func LogI(args ...interface{}) {
	initSetting()
	logrus.Info(args...)
}
func LogW(args ...interface{}) {
	initSetting()
	logrus.Warn(args...)
}
func LogD(args ...interface{}) {
	initSetting()
	logrus.Debug(args...)
}
func LogE(args ...interface{}) {
	initSetting()
	logrus.Error(args...)
}
//func LogJson(tag,data string)  {
//	if !strings.HasPrefix(data,"{"){
//		return
//	}else if j,err:=json.JsonIndent(data);err!=nil{
//		LogI(tag,":",err)
//	}else {
//		LogI(tag,":",j)
//	}
//}
func LogNotImplement() {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(1)
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo
		filename = filepath.Base(filename)           // /full/path/basename.go => basename.go
	}

	LogE("not implement ", filename, ":", line, ":", funcname, "()")
}
func addCallFunctionName(args ...interface{}) []interface{}{
	_, _, funcname := "???", 0, "???"
	pc, _, _, ok := runtime.Caller(2)
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo
	}
	newArgs:=[]interface{}{funcname+"():"}
	newArgs=append(newArgs,args...)
	return newArgs
}

func getErrorLine() string {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(2)
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo
		filename = filepath.Base(filename)           // /full/path/basename.go => basename.go
	}
	return fmt.Sprint(filename, ":", line, ":", funcname, "()")
}
