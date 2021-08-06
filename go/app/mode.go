package app

type AUTH struct {
	Uid         int    `json:"uid"`
	Aid         int    `json:"aid"`
	Time        int    `json:"time"`
	Power       string `json:"power"`
	Dev         string `json:"dev"`
	Time_upData int    `json:"time_updata"`
	Time_Login  int    `json:"time_Login"`
	Sign        string `json:"sign"`
}

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type response_data struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type response_list struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int64       `json:"count"`
	Data  interface{} `json:"data"`
}
