package tools

import "strings"

func FormatTime(t string) string {
	ft := strings.Split(t, "-")
	str := ft[0] + ft[1] + ft[2]
	return str
}
