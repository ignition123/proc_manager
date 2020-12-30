package objects

var Adapters = make(map[string]map[string]interface{})

var Response Resp

type Resp struct{
	Msg string `json:"msg"`
	PsName string `json:"psName"`
	Ip string `json:"ip"`
	LUT string `json:"lut"`
	PsPath string `json:"path"`
}