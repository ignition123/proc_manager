package objects

var Conf *Config

type Server struct{
	Host *string `json:"host"`
	Port int `json:"port"`
}

type Applications struct{
	Name *string `json:"name"`
	Restart *string `json:"restart"`
	Path *string `json:"path"`
	Cmd *string `json:"cmd"`
	Sleep *int `json:"sleep"`
	Status bool `json:"status"`
}

type Config struct{
	Server *Server
	App []*Applications `json:"applications"`
}