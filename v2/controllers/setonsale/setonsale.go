package setonsale

import (
	"github.com/astaxie/beego"
	"vkapi/models/result"
	"vkapi/v2/models/setonsale"
	"vkapi/models/tools"
	"fmt"
)

type VtwoSetonsaleController  struct {
	beego.Controller
}
func Setonsale(args ...string)[]result.Result{
	onsale:=setonsale.Onsaledata{}
	for _,arg:=range args{
		if arg==""{
			resulta:=[]result.Result{}
			resultb:=result.Result{}
			resultb.Result_code="FAIL"
			resultb.Result_des="请求的参数中存在空值，请检查后再提交"
			resulta=append(resulta,resultb)
			return resulta
		}
	}
	onsale.Device_info=args[0]
	onsale.Max=args[1]
	onsale.Mch_id=args[2]
	onsale.Min=args[3]
	onsale.Money=args[4]
	onsale.Number=args[5]
	fmt.Println(args[6])
	onsale.Timestart=tools.FormatTime(args[6])

	onsale.Timeend=tools.FormatTime(args[7])
	onsale.From="weikai"
	onsale.Goods_tag="xian"
	result:=setonsale.Setonsale(onsale)
  	return  result
}