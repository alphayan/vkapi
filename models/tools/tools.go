package tools

import (
	"bytes"
	"reflect"
	"strings"
	"fmt"
)

func ToKeyValueStr(v interface{}) string {
	var signstr bytes.Buffer
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		name := field.Name
		keytemp := field.Tag.Get("json")
		if keytemp == "" {
			break
		}
		keymap := strings.Split(keytemp, ",")
		key := keymap[0]
		value := (reflect.Indirect(vv).FieldByName(name)).String()
		if value != "" && key != "json" {
			signstr.WriteString(key + "=" + value + "&")
		}
	}
	stra := strings.Split(signstr.String(), "&")//去掉最后一个&符，强迫症
	fmt.Println(stra)

	strb := strings.Join(stra[:len(stra) - 1], "&")
	fmt.Println(strb)

	return strb
}