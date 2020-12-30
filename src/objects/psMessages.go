package objects

type Msg struct{
	Msg string `json:"msg"`
	PsName string `json:"psName"`
	Ip string `json:"ip"`
	LUT string `json:"lut"`
	PsPath string `json:"path"`
}