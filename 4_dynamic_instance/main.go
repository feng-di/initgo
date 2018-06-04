package main

import (
	"fmt"
	"reflect"
	"initGo/4_dynamic_instance/special"
	"strings"
	"github.com/pkg/errors"
)

var typeRegistry = make(map[string]reflect.Type)

func registerType(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	fmt.Printf("---------- : `%+#v`, \n\n\n", t.Name())
	typeRegistry[t.Name()] = t

}

func newStruct(name string) (interface{}, bool) {


	elem, ok := typeRegistry[name]
	if !ok {
		return nil, false
	}
	return reflect.New(elem).Elem().Interface(), true
}

func init() {
	registerType((*test)(nil))
	registerType((*special.Special)(nil))
}


type test struct {
	Name string
	Sex  int
}

func main() {
	fmt.Printf("---------- : `%+#v`, \n\n\n", typeRegistry)

	structName := "test"

	s, ok := newStruct(structName)
	if !ok {
		return
	}

	fmt.Println(s, reflect.TypeOf(s))

	t, ok := s.(test)
	if !ok {
		return
	}
	t.Name = "i am test"
	fmt.Println(t, reflect.TypeOf(t))

	// Test struct in package
	sp, ok := newStruct("special.Special")
	fmt.Println(sp, reflect.TypeOf(sp))
	tsp, ok := sp.(special.Special)
	tsp.Name = "i am sptest"
	fmt.Println(t, reflect.TypeOf(tsp))

	fmt.Printf("---------- : `%+#v`, \n\n\n", strings.Title("createCard"))
}


func InvokeObjectMethod(object interface{}, methodName string, args ...interface{}) ([]reflect.Value, error) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	// If method not exist return error
	if _, exist := reflect.ValueOf(object).Type().MethodByName(methodName); exist == false {
		return nil, errors.New("error message")
	}

	return reflect.ValueOf(object).MethodByName(methodName).Call(inputs), nil
}
