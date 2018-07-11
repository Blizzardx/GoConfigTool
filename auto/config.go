package auto

type WGQueryPlayerInfos struct {
	PlayerIds []int64
}
type WGAddMoney struct {
	PlayerIds []int64 `json:"playerIds"`
	Diamond   int32   `json:"diamond"`
	Gold      int32   `json:"gold"`
}

//支付成功，通知加币和客户端
type WGPaymentSuccess struct {
	PlayerId int64
	ItemId   int32
}

type GWAddMoney struct {
	ErrorPlayerIds []int64
}

type WorldPlayerInfo struct {
	PlayerId           int64  `json:"id"`
	PlayerName         string `json:"name"`
	PlayerState        int32  `json:"state"`
	Diamond            int32  `json:"diamond"`
	CreateTime         int64  `json:"createTime"`
	AccountDisableTime int64  `json:"disabletime"`
	PhoneNumber        string `json:"phoneNumber"`
	GuildId            int64  `json:"guildId"`
	GuildName          string `json:"guildName"`
	LeagueId           int64  `json:"leagueId"`
	LeagueName         string `json:"leagueName"`
	IsPayment          int    `json:"isPayment"`
}
type WorldApplyCreateLeagueInfo struct {
	ApplyId     int64  `json:"id"`
	PlayerId    int64  `json:"playerId"`
	GuildId     int64  `json:"guildId"`
	ApplyStatus int32  `json:"applyState"`
	ApplyType   int32  `json:"applyType"`
	CreateTime  int64  `json:"createTime"`
	Description string `json:"desc"`
	Name        string `json:"leagueName"`
	PlayerName  string `json:"playerName"`
	PlayerIcon  string `json:"playerIcon"`
	PhoneNumber string `json:"phoneNumber"`
}
