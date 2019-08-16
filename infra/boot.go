package infra


import "fmt"

var (
	boot *Boot = &Boot{}
)

func NewBoot() *Boot {
	return boot
}

type Boot struct {}



func (b *Boot) Load () {
	for name, starter := range AllStarters() {
		fmt.Println("Load ", name)
		starter.Init()
	}
}


func (b *Boot) Stop () {
	for _, starter := range AllStarters() {
		starter.Stop()
	}
}
