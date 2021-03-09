package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	v:= struct {
		FString string `json:"f_string"`
		FInt int
		Slice []int
		Obj struct{
			NF int
		}
	}{
		FString: "str",
		FInt: 188,
		Slice: []int{1,2,3},
		Obj: struct{ NF int }{NF: 33333},
	}

	PrintStruct(&v)
}

func PrintStruct(in interface{}) {
	if in == nil {
		return
	}

	val := reflect.ValueOf(in)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	for i:=0; i<val.NumField(); i++ { // NumField only for structs
		typeF:=val.Type().Field(i)
		if typeF.Type.Kind() == reflect.Struct {
			fmt.Printf("nested field: %v\n", typeF.Name)
			PrintStruct(val.Field(i).Interface())
			continue
		}
		log.Printf("name = %s, type=%s,val=%v,tag=`%s`\n", typeF.Name, typeF.Type, val.Field(i), typeF.Tag)
	}
}
