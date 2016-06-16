package insertpassword

import (
	"vkapi/models/result"
	"vkapi/v1/models/insertpassword"
	"fmt"
)

func InsertPassword(args ...string)result.Result{
	pi:=insertpassword.PasswordInfo{}
	for _,arg:=range args{
		fmt.Println(arg)
		if arg==""{
			result:=result.Result{}
			result.Result_code="FAIL"
			result.Result_des="请求的参数中存在空值，请检查后再提交"
			return result
		}
	}
	pi.Password=args[0]
	pi.Address=args[1]
	pi.Device_info=args[2]
	pi.Name=args[3]
	pi.Prin=args[4]
	pi.Mch_id=args[5]
	result:=insertpassword.InsertPassword(pi)
	return  result
}