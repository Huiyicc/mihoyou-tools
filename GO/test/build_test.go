package main

import (
	"fmt"
	"github.com/1925473403/mihoyou-tools/api"
	"github.com/1925473403/mihoyou-tools/other"
	"testing"
	"time"
)

var (
	YouUUID = ""	//你的游戏uid
	Cookies = ""
	server = api.TIANKONGDAO	//账号类型 天空岛|世界树
)

func build(t *testing.T)  {
	cli:=api.NewCli(YouUUID,Cookies,server)
	res,err:=cli.GetDaily()
	if err != nil {
		panic(err)
	}
	if res.Retcode != 0 {
		fmt.Printf("错误代码:%d\n错误信息:%s\n",res.Retcode,res.Message)
		return
	}
	fmt.Println(res)
	log:=fmt.Sprintf("原粹树脂: %d/%d\n",res.Data.CurrentResin,res.Data.MaxResin)
	log+=fmt.Sprintf("原粹树脂恢复时间: %s\n",time.Unix (time.Now().Unix() + other.AtoInt64(res.Data.ResinRecoveryTime),0).Format("2006-01-02 15:04"))
	log+=fmt.Sprintf("每日委托完成度: %d/%d\n",res.Data.FinishedTaskNum,res.Data.TotalTaskNum)
	log+=fmt.Sprintf("周本减半次数: %d/%d\n",res.Data.RemainResinDiscountNum,res.Data.ResinDiscountNumLimit)
	log+=fmt.Sprintf("探索派遣次数: %d/%d\n",res.Data.CurrentExpeditionNum,res.Data.MaxExpeditionNum)
	for i:=0;i<len(res.Data.Expeditions);i++{
		tim:=time.Now().Unix()
		tim += other.AtoInt64(res.Data.Expeditions[i].RemainedTime)
		log+=fmt.Sprintf("\t[%s]派遣人员: %s 截止至:%s\n",
			other.Ifs(res.Data.Expeditions[i].Status==api.EXPLORATION_STATUS_ONGOING,"派遣中","派遣完成"),	//api.EXPLORATION_STATUS_ONGOING | api.EXPLORATION_STATUS_ONGOING
			cli.SwitchName(res.Data.Expeditions[i].AvatarSideIcon),
			time.Unix(tim, 0).Format("2006-01-02 15:04"))
	}
	fmt.Println(log)
}
