package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"projects/rninet/code"
	"projects/rninet/code/generator"
)

var (
	Commands []cli.Command = []cli.Command {
		{
			Name: "new",
			Aliases: []string{},
			Usage: "new a server/cli code",
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "server, s",
					Usage: "server name",
					Destination: &srvName,
				},
				cli.StringFlag{
					Name: "client, c",
					Usage: "client name",
					Destination: &cliName,
				},
			},
			Action: NewAction,
		},
	}

	srvName string
	cliName string

)


func NewAction (c *cli.Context) error {

	fmt.Printf("new server: %s, client: %s\n", srvName, cliName)

	if len(srvName) != 0 {
		code.AddGenerator("dir", &generator.DirGenerator{})
	}

	if len(cliName) != 0 {

	}


	return nil
}