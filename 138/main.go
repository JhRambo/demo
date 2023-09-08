package main

import (
	"fmt"
	"reflect"
)

type GWResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func main() {
	t := reflect.TypeOf(GWResponse{})
	field, _ := t.FieldByName("Code")
	tag := field.Tag
	newTag := tag.Get("json") + " json:\"Code\""

	u := GWResponse{
		Code: 0,
	}

	structValue := reflect.ValueOf(&u).Elem()
	structFieldValue := structValue.FieldByName("Code")
	structFieldType := structFieldValue.Type()
	newFieldValue := reflect.New(structFieldType).Elem()
	newFieldValue.SetInt(int64(u.Code))
	field.Tag = reflect.StructTag(newTag)
	fmt.Printf("tag：%v\n", field.Tag)
}
