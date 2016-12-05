package insertmongodb

import (
	"vkapi/models/result"
	"vkapi/models/tools"
	"fmt"
)

type Mchinfo struct {
	Class    string    `json:"class"`
	Mchid    string    `json:"mch_id"`
	Password string    `json:"password"`
	Key      string    `json:"key"`
}

func InsertMongodb(mchinfo Mchinfo) result.Result {
	result := result.Result{}
	fmt.Println(mchinfo)
	str := tools.ToKeyValueStr(mchinfo)
	fmt.Println(str)
	str = "http://api.vkaifu.com:8000/setmchid?pwd=911202&" + str
	re := tools.Get(str)
	tools.JsonDecodebytes(re, &result)
	return result
}