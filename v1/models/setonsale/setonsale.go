package setonsale

import (
	"vkapi/models/result"
	"vkapi/models/tools"
	"fmt"
	"time"
)

type  Onsaledata struct {
	Time      	string            `json:"time"`
	Goods_tag 	string            `json:"goods_tag"`
	From      	string            `json:"from"`
	Number    	string            `json:"number"`
	Max       	string            `json:"max"`
	Min       	string            `json:"min"`
	Mch_id    	string     	  `json:"mch_id"`
	Money     	string   	  `json:"money"`
	Device_info     string            `json:"device_info"`
	Timestart       string
	Timeend         string
}

func Setonsale(o Onsaledata)[]result.Result  {
	resulta:=[]result.Result{}
	resultb:=result.Result{}
	if o.Timestart==o.Timeend{
		o.Time=o.Timestart
		str:=tools.ToKeyValueStr(o)
		fmt.Println(str)
		str="http://localhost:8000/setgoodstag?pwd=911202&"+str
		re:=tools.Get(str)
		fmt.Println(string(re))
		tools.JsonDecodebytes(re,&resultb)
		resulta=append(resulta,resultb)
		return resulta
	}else{
		tm1,_:=time.Parse("20060102",o.Timestart)//将时间转化为时间戳
		tm2,_:=time.Parse("20060102",o.Timeend)
		fmt.Println(tm1.Unix())

		fmt.Println(tm2.Unix())
		num:=(tm2.Unix()-tm1.Unix())/86400
		k:=int(num)
		fmt.Println(k)
		for i:=0;i<=k;i++{
			o.Time=o.Timestart

			tm, _ := time.Parse("20060102", o.Timestart)
			temp := tm.Add(time.Hour * 24)
			fmt.Println(temp)
			o.Timestart = temp.Format("20060102")

			str:=tools.ToKeyValueStr(o)
			str="http://localhost:8000/setgoodstag?pwd=911202&"+str
			fmt.Println(str)
			re:=tools.Get(str)
			//fmt.Println(string(re))
			tools.JsonDecodebytes(re,&resultb)
			resulta=append(resulta,resultb)
		}

		return resulta
	}
}