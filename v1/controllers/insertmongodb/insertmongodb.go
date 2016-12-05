package insertmongodb

import (
	"vkapi/v1/models/insertmongodb"
	"vkapi/models/result"
)

func InsertMongodb(args ...string) result.Result {
	mchinfo := insertmongodb.Mchinfo{}
	for _, arg := range args {
		if arg == "" {

			result := result.Result{}
			result.Result_code = "FAIL"
			result.Result_des = "请求的参数中存在空值，请检查后再提交"
			return result
		}
	}
	mchinfo.Mchid = args[0]
	mchinfo.Class = "sub"
	mchinfo.Password = "000000"
	mchinfo.Key = "ivknetyyq2015"
	result := insertmongodb.InsertMongodb(mchinfo)
	return result
}
