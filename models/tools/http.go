package tools

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)

func Get(geturl string) []byte {

	res, err := http.Get(geturl)
	if err != nil {
		fmt.Println(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	return robots
}
func Post(posturl, data string) []byte {
	res, err := http.Post(posturl, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	return robots
}
