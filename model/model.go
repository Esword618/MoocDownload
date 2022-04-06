package model

type CheckStruct struct {
	Data struct {
		Version string `json:"version"`
		Info    string `json:"info"`
	} `json:"data"`
	Msg interface{} `json:"msg"`
}
