package api

import (
	"fmt"
	"github.com/1925473403/mihoyou-tools/other"
	"time"
)

const (
	STALT = "xV8v4Qu54lUKrEYFZkJhB8cuOh9Asafs"
)


func (c *Daliy)GetDS(salt string) string {
	timestamp:=time.Now().Unix()
	randomStr:=other.RandChar(6)
	query:=fmt.Sprintf("role_id=%s&server=%s",c.UID,c.Server)
	singsf:=fmt.Sprintf("salt=%s&t=%d&r=%s&b=%s&q=%s",salt,timestamp,randomStr,"",query)
	sing:=other.GetMd5String(singsf)
	DS:=fmt.Sprintf("%d,%s,%s",timestamp,randomStr,sing)
	return DS
}