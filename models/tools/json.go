package tools

import "encoding/json"

//把struct数据解压出来转化为json
func JsonEncodestruct(v interface{}) (jsonresult string, err error) {
	output, err := json.Marshal(v)
	jsonresult = string(output)
	return jsonresult, err
}

//把byte数据格式化为v
func JsonDecodebytes(b []byte, v interface{}) error {
	err := json.Unmarshal(b, v)
	return err
}