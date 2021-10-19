package other

import "strings"

//DisableEncodeJSON
//转义json字符串去掉一层转义
func DisableEncodeJSON(str string) string {
	str = strings.ReplaceAll(str, "\\\"", "\"")
	str = strings.ReplaceAll(str, "\\\\", "\\")
	return str
}

//HttpsUrl2HttpUrl
//https链接转http链接
func HttpsUrl2HttpUrl (url string) (ret string){
	if len(url) < 5 {
		return url
	}
	ret = url
	if ret [0:5] == "https" {
		ret = "http"+ret [5:]

	}
	return
}

func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + len(start)  // 增加了else，不加的会把start带上
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}
