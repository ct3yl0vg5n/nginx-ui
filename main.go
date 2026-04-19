package main

import (
	"fmt"
	"os"

	"github.com/0xJacky/nginx-ui/app"
	"github.com/0xJacky/nginx-ui/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	cliApp := &cli.App{
		Name:    "nginx-ui",
		Usage:   "A web UI for managing Nginx configurations",
		Version: app.Version,
		Commands: []*cli.Command{
			cmd.ServeCommand,
		},
		Action: func(c *cli.Context) error {
			// Default action: start the server
			return cmd.ServeCommand.Action(c)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Path to the configuration file",
				Value:   "app.ini",
				EnvVars: []string{"NGINX_UI_CONFIG"},
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
