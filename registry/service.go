package registry



type Service struct {
	Name string `json: "name"`
	Addr string `json: "addr"`
	Metadata map[string]string `json: "metadata"`
}
