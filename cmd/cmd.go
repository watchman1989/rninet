package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"projects/rninet/generator"
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
				cli.StringFlag{
					Name: "proto, p",
					Usage: "porot3 file",
					Destination: &protoFile,
				},
			},
			Action: NewAction,
		},
	}

	srvName string
	cliName string
	protoFile string

)


func NewAction (c *cli.Context) error {

	fmt.Printf("new server: %s, client: %s\n", srvName, cliName)

	if len(srvName) != 0 {
		G := generator.NewGenerator(
			generator.WithSrvFlag(),
			generator.WithOutput(srvName),
			generator.WithProtoFile(protoFile),
		)

		G.Gen()
	}

	if len(cliName) != 0 {

	}


	return nil
}