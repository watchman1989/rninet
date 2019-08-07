package registry



type Node struct {
	Id string
	Addr string
	Port int
	Metadata map[string]string
}


type Service struct {
	Name string
	Metadata map[string]string
	Nodes []*Node
}
