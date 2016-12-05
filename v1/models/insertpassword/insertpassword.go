package insertpassword

import (
	"vkapi/models/result"

	"vkapi/models/tools"
	"fmt"
)

type PasswordInfo struct {
	Mch_id      string `json:"mch_id"`
	Name        string  `json:"name"`
	Device_info string `json:"device_info"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	Prin        string `json:"prin"`
}

func InsertPassword(pi PasswordInfo) result.Result {
	result := result.Result{}
	str := tools.ToKeyValueStr(pi)
	str = "http://api.vkaifu.com:8000/insertpassword?pwd=911202&" + str
	re := tools.Get(str)
	fmt.Println(re)
	tools.JsonDecodebytes(re, &result)
	return result
}
