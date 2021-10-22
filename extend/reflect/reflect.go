package cherryReflect

import (
	"fmt"
	"github.com/cherry-game/cherry/error"
	"github.com/cherry-game/cherry/extend/string"
	"github.com/cherry-game/cherry/facade"
	"reflect"
	"runtime"
)

func ReflectTry(f reflect.Value, args []reflect.Value, handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("-------------panic recover---------------")
			if handler != nil {
				handler(err)
			}
		}
	}()
	f.Call(args)
}

func GetStructName(v interface{}) string {
	return reflect.Indirect(reflect.ValueOf(v)).Type().Name()
}

func GetFuncName(fn interface{}) string {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic(fmt.Sprintf("[fn = %v] is not func type.", fn))
	}

	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	return cherryString.CutLastString(fullName, ".", "-")
}

//GetInvokeFunc reflect function convert to HandlerFn
func GetInvokeFunc(name string, fn interface{}) (*cherryFacade.HandlerFn, error) {
	if name == "" {
		return nil, cherryError.Error("func name is nil")
	}

	if fn == nil {
		return nil, cherryError.Errorf("func is nil. name = %s", name)
	}

	typ := reflect.TypeOf(fn)
	val := reflect.ValueOf(fn)

	if typ.Kind() != reflect.Func {
		return nil, cherryError.Errorf("name = %s is not func type.", name)
	}

	var inArgs []reflect.Type
	for i := 0; i < typ.NumIn(); i++ {
		t := typ.In(i)
		inArgs = append(inArgs, t)
	}

	var outArgs []reflect.Type
	for i := 0; i < typ.NumOut(); i++ {
		t := typ.Out(i)
		outArgs = append(outArgs, t)
	}

	invoke := &cherryFacade.HandlerFn{
		Type:    typ,
		Value:   val,
		InArgs:  inArgs,
		OutArgs: outArgs,
	}

	return invoke, nil
}
