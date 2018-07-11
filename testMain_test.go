package main

import (
	"github.com/Blizzardx/GoConfigTool/auto"
	"github.com/Blizzardx/GoConfigTool/decoder"
	"io/ioutil"
	"testing"
)

func Test_GenTestConfig(t1 *testing.T) {
	codeC := new(decoder.MsgPackDecodeC)
	m1 := &auto.WGQueryPlayerInfos{
		PlayerIds: []int64{0, 2, 3, 4, 4, 4, 4, 4, 4, 4},
	}
	content, _ := codeC.Encode(m1)
	ioutil.WriteFile("config/WGQueryPlayerInfos.cfg", content.([]byte), 0666)

	m2 := &auto.WGAddMoney{
		Diamond:   100,
		Gold:      20,
		PlayerIds: []int64{0, 2, 3, 4, 4, 4, 4, 4, 4, 4},
	}
	content, _ = codeC.Encode(m2)
	ioutil.WriteFile("config/WGAddMoney.cfg", content.([]byte), 0666)

	m3 := &auto.WGPaymentSuccess{
		PlayerId: 100000,
		ItemId:   2,
	}
	content, _ = codeC.Encode(m3)
	ioutil.WriteFile("config/WGPaymentSuccess.cfg", content.([]byte), 0666)

	m4 := &auto.GWAddMoney{
		ErrorPlayerIds: []int64{0, 2, 3, 4, 4, 4, 4, 4, 4, 4},
	}
	content, _ = codeC.Encode(m4)
	ioutil.WriteFile("config/GWAddMoney.cfg", content.([]byte), 0666)

	m5 := &auto.WorldPlayerInfo{
		PlayerId:           20,
		PlayerName:         "sbaoxue",
		PlayerState:        1,
		Diamond:            20,
		CreateTime:         10000,
		AccountDisableTime: 0,
		PhoneNumber:        "123121212",
		GuildId:            2,
		GuildName:          "diyigonghui",
		LeagueId:           2,
		LeagueName:         "diyilianmeng",
		IsPayment:          1,
	}
	content, _ = codeC.Encode(m5)
	ioutil.WriteFile("config/WorldPlayerInfo.cfg", content.([]byte), 0666)

	m6 := &auto.WorldApplyCreateLeagueInfo{
		ApplyId:     2,
		PlayerId:    20032233,
		GuildId:     3,
		ApplyStatus: 2,
		ApplyType:   0,
		CreateTime:  2300323,
		Description: "dewwerwer",
		Name:        "apply test name",
		PlayerName:  "test robot 11",
		PlayerIcon:  "shttps://;a;sdfasdfsdf",
		PhoneNumber: "as232323",
	}
	content, _ = codeC.Encode(m6)
	ioutil.WriteFile("config/WorldApplyCreateLeagueInfo.cfg", content.([]byte), 0666)

}
