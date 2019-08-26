package main

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/watchman1989/rninet/cmd"
	"os"
	"runtime"
)


func InitEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}


func InitArg() {

}



func init () {
	InitEnv()
	InitArg()

}




func main () {

	fmt.Println("MAIN_START")

	app := cli.NewApp()
	app.Name = "rninet"
	app.Usage = "A micro service frame"

	app.Commands = cmd.Commands


	if err := app.Run(os.Args); err != nil {
		fmt.Printf("APP_RUN_ERROR: %v\n", err)
	}


	//time.Sleep(1000 * time.Second)

}
