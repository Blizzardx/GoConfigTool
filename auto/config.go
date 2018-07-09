package auto



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
