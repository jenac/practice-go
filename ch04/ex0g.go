package main

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

type Example struct {
	Field  int    `json:"one"`
	Field2 string `json:"two"`
}

func (e Example) String() string {
	return fmt.Sprintf("%d, %s", e.Field, e.Field2)
}

type ReflectUsage struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (ref ReflectUsage) String() string {
	return fmt.Sprintf("Name: %s, Age %d", ref.Name, ref.Age)
}

func (ref ReflectUsage) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Name: %s", ref.Name))
	return buf.Bytes(), nil
}

func Usage() {
	var example ReflectUsage
	example.Name = "James"
	example.Age = 20
	typ := reflect.TypeOf(example)
	fmt.Printf("%T\n", example)
	fmt.Println(typ)

	val := reflect.ValueOf(example)
	fmt.Printf("%#v\n", example)
	fmt.Printf("%v\n", example)
	fmt.Println(val)

	fmt.Println(typ.NumField(), typ.NumMethod())
	fmt.Println(val.NumField(), val.NumMethod())

	fmt.Println(typ.FieldByName("Name"))
	fmt.Println(typ.FieldByName("Age"))
	fmt.Println(val.FieldByName("Name"))
	fmt.Println(val.FieldByName("Age"))

	fmt.Println(typ.NumMethod(), typ.Method(0))
	fmt.Println(typ.NumMethod(), typ.Method(1))
	method1 := val.Method(1)
	args := make([]reflect.Value, 0)
	result := method1.Call(args)
	fmt.Println(result)

	method2 := val.MethodByName("MarshalJSON")
	args2 := make([]reflect.Value, 0)
	result2 := method2.Call(args2)
	fmt.Println(string(result2[0].Bytes()))

	valCanSet := reflect.ValueOf(&example)
	ptr := valCanSet.Elem()
	ptr.FieldByName("Age").SetInt(1000)
	fmt.Println(example)
}

func UnsafeUsage() {
	var example ReflectUsage
	example.Name = "goLang"
	example.Age = 10
	typ:=reflect.TypeOf(unsafe.Sizeof(example))
	fmt.Println(typ)
	fmt.Println(unsafe.Sizeof(example))

	ptr:=unsafe.Pointer(&example)
	fmt.Println(ptr)
	fmt.Println(*(*string)(ptr))

	ptrOfSecondField := unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(example.Age))
	fmt.Println(ptrOfSecondField)
	fmt.Println(*(*int)(ptrOfSecondField))

	*(*int)(ptrOfSecondField) = 32
	fmt.Println(example)

}

func main() {
	Age := 20
	fmt.Printf("%T\n", Age)
	typ := reflect.TypeOf(Age)
	fmt.Println(typ)
	fmt.Println(reflect.ValueOf(Age))
	fmt.Println("----------")
	Usage()
	fmt.Println("----------")
	UnsafeUsage()

}
