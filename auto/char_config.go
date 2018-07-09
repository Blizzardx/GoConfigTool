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

