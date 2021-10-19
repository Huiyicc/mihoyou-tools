package api

import (
	"github.com/1925473403/mihoyou-tools/other"
	"strings"
)

const (
	APPVERSION = "2.11.1"
)

func (c *Daliy)GetHeader (cookies,DS string) (ret map[string]string) {
	ret = make(map[string]string)
	ret["Host"]="api-takumi.mihoyo.com"
	ret["Cookie"]=cookies
	ret["Accept"]="application/json, text/plain, */*"
	ret["Content-Type"]="application/json"
	ret["User-Agent"]="Mozilla/5.0 (Linux; Android 6.0.1; MuMu Build/V417IR; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.70 Mobile Safari/537.36 miHoYoBBS/"+ APPVERSION
	ret["Referer"]="https://webstatic.mihoyo.com/app/community-game-records/index.html?v=6"
	ret["x-rpc-channel"]="appstore"
	ret["x-rpc-device_id"]=strings.ToUpper(other.RandChar(32))
	ret["x-rpc-app_version"]=APPVERSION
	ret["x-rpc-device_model"]="iPhone11,8"
	//ret["x-rpc-device_name"]=
	ret["x-rpc-client_type"]="5"
	ret["Origin"]= "https://webstatic.mihoyo.com"
	ret["DS"] = DS
	return
}
