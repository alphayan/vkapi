package tinyurl

import (
	"vkapi/models/result"
	"vkapi/models/tools"
	"strings"
)

type UrlInfo struct {
	Mchid string `json:"mch_id"`
	Tips string `json:"tips"`
	Body string `json:"body"`
	Device_info string `json:"device_info"`
}
type ResponeInfo struct {
	Tinyurl string `json:"tinyurl"`
	Status  int    `json:"status"`
	Longurl string `json:"longurl"`
	Err_msg string `json:"err_msg"`
}
func UrlEncode(url string) string {

	replacer := strings.NewReplacer("&", "%26")
	return replacer.Replace(url)
}
func TinyUrl(ui UrlInfo)result.Result{
	result:=result.Result{}
	str:=tools.ToKeyValueStr(ui)
	str="http://www.vkaifu.com/vkpay/pay?"+str
	result.Result_code="SUCCESS"
	result.Result_des="获取短链接成功"
	result.Data=str
	return result
}
