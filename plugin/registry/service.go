package registry



type Service struct {
	Name string `json: "name"`
	Addr string `json: "addr"`
	Ip string `json: "ip"`
	Port int `json: "port"`
	Metadata map[string]string `json: "metadata"`
}
