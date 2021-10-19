package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Daliy)GetDaily() (ret RetJson,err error){
	cookies:="mi18nLang=zh-cn; _MHYUUID=26e0599a-8bcf-4f20-8930-227633a5456d; _ga_5DECE6NN1T=GS1.1.1611394350.1.1.1611394352.0; _ga_E36KSL9TFE=GS1.1.1611825767.6.1.1611825852.0; ltoken=WP6Xmz0rYQpirIB8WXIpf7hWrzCLx7WQ1BUfF0a7; ltuid=186057372; _ga_FLGX56JEBS=GS1.1.1613882720.2.1.1613882903.0; _ga_0S6JZDKDXS=GS1.1.1619159309.1.1.1619159345.0; _ga_QH3ZTJN1H1=GS1.1.1620815870.8.0.1620815874.0; _ga_ELSF9S5LND=GS1.1.1621138803.3.1.1621138803.0; _ga_H3M17VQB59=GS1.1.1622796593.2.0.1622796597.0; UM_distinctid=17b619e506c2ef-03e5975fac39ba-7868786b-1fa400-17b619e506dc9e; _ga_HKTGWLY8PN=GS1.1.1630118034.3.1.1630118298.0; _ga_831VBKXN1V=GS1.1.1632822704.1.1.1632822731.0; _ga=GA1.2.282685535.1605164115; _ga_9TTX3TE5YL=GS1.1.1632823350.20.1.1632824240.0; cookie_token=ObPFOelUQwFEniHEPx9YhsLm6tlKyKYU7Nh9Fcb7; account_id=186057372; mihoyo-survey-3627=1; login_uid=186057372; login_ticket=yux1LJeUvELZvnqmmgAzMztVdWto4GTEb7HBIm9A; _gid=GA1.2.377472684.1634524329; CNZZDATA1275023096=1596781702-1605159618-%7C1634530122"
	//url:="https://api-takumi.mihoyo.com/game_record/app/genshin/api/dailyNote"
	url:=fmt.Sprintf("https://api-takumi.mihoyo.com/game_record/app/genshin/api/dailyNote?role_id=%s&server=%s",c.UID,c.Server)
	//提交请求
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	head:=c.GetHeader(cookies,c.GetDS(STALT))
	eoc:=""
	for k,v:=range head{
		reqest.Header.Add(k,v)
		eoc+=k+": "+v+"\n"
	}
	fmt.Println(eoc)
	//处理返回结果
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	bod,err:=io.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(bod))
	err=json.Unmarshal(bod,&ret)
	return
}