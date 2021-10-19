package api

import "github.com/1925473403/mihoyou-tools/other"

type Daliy struct {
	UID string
	Cookies string
	Server string
	rmap	map[string]string
}

const (
	TIANKONGDAO = "cn_gf01"
	SHIJIESHU = "cn_qd01"

	EXPLORATION_STATUS_ONGOING = "Ongoing"	//派遣中
	EXPLORATION_STATUS_FINISHED = "Finished"	//派遣完成
)


func NewCli(UID string,Cookies string,Server string) *Daliy {
	c:=&Daliy{}
	c.UID = UID
	c.Cookies = Cookies
	c.Server = Server
	c.rmap = make(map[string]string)
	c.rmap["Albedo"]="阿贝多"
	c.rmap["Ambor"]="安柏"
	c.rmap["Barbara"]="芭芭拉"
	c.rmap["Beidou"]="北斗"
	c.rmap["Bennett"]="班尼特"
	c.rmap["Chongyun"]="重云"
	c.rmap["Diluc"]="迪卢克"
	c.rmap["Diona"]="迪奥娜"
	c.rmap["Eula"]="优菈"
	c.rmap["Fischl"]="菲谢尔"
	c.rmap["Ganyu"]="甘雨"
	c.rmap["Hutao"]="胡桃"
	c.rmap["Jean"]="琴"
	c.rmap["Kazuha"]="枫原万叶"
	c.rmap["Kaeya"]="凯亚"
	c.rmap["Ayaka"]="神里绫华"
	c.rmap["Keqing"]="刻晴"
	c.rmap["Klee"]="可莉"
	c.rmap["Lisa"]="丽莎"
	c.rmap["Mona"]="莫娜"
	c.rmap["Ningguang"]="凝光"
	c.rmap["Noelle"]="诺艾尔"
	c.rmap["Qiqi"]="七七"
	c.rmap["Razor"]="雷泽"
	c.rmap["Rosaria"]="罗莎莉亚"
	c.rmap["Sucrose"]="砂糖"
	c.rmap["Tartaglia"]="达达利亚"
	c.rmap["Venti"]="温迪"
	c.rmap["Xiangling"]="香菱"
	c.rmap["Xiao"]="魈"
	c.rmap["Xingqiu"]="行秋"
	c.rmap["Xinyan"]="辛焱"
	c.rmap["Yanfei"]="烟绯"
	c.rmap["Zhongli"]="钟离"
	c.rmap["Yoimiya"]="宵宫"
	c.rmap["Sayu"]="早柚"
	c.rmap["Shogun"]="雷电将军"
	c.rmap["Aloy"]="埃洛伊"
	c.rmap["Sara"]="九条裟罗"
	c.rmap["Kokomi"]="珊瑚宫心海"
	return c
}

type RetJson struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		CurrentResin              int    `json:"current_resin"`
		MaxResin                  int    `json:"max_resin"`
		ResinRecoveryTime         string `json:"resin_recovery_time"`
		FinishedTaskNum           int    `json:"finished_task_num"`
		TotalTaskNum              int    `json:"total_task_num"`
		IsExtraTaskRewardReceived bool   `json:"is_extra_task_reward_received"`
		RemainResinDiscountNum    int    `json:"remain_resin_discount_num"`
		ResinDiscountNumLimit     int    `json:"resin_discount_num_limit"`
		CurrentExpeditionNum      int    `json:"current_expedition_num"`
		MaxExpeditionNum          int    `json:"max_expedition_num"`
		Expeditions               []struct {
			AvatarSideIcon string `json:"avatar_side_icon"`
			Status         string `json:"status"`
			RemainedTime   string `json:"remained_time"`
		} `json:"expeditions"`
	} `json:"data"`
}

func (c *Daliy)SwitchName (name string) (ret string) {
	u:=other.GetBetweenStr (name,"UI_AvatarIcon_Side_",".png")
	ret=c.rmap[u]
	if ret == "" {
		ret = name
	}
	return
}