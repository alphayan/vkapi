package tinyurl

import (
	"vkapi/models/result"
	"vkapi/v1/models/tinyurl"
)

func Tinyurl(args ...string) result.Result {
	ui := tinyurl.UrlInfo{}
	ui.Device_info = args[0]
	ui.Body = args[1]
	ui.Mchid = args[2]
	ui.Tips = args[3]
	result := tinyurl.TinyUrl(ui)
	return result
}
