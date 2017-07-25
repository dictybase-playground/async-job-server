package main

import (
	"os"

	"github.com/urfave/cli"
)

//Arguments struct is the parameters that blastp takes
type Arguments struct {
	Database string  `json:"database"`
	Query    string  `json:"query"`
	Evalue   float64 `json:"evalue"`
	Numalign int     `json:"numalign"`
	Wordsize int     `json:"wordsize"`
	Matrix   string  `json:"matrix"`
	Seg      bool    `json:"seg"`
	Gapped   bool    `json:"gapped"`
}

func main() {

	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "BLAST worker"
	app.Usage = "Run BLAST queries"
	app.Commands = []cli.Command{

		{
			Name:   "run",
			Usage:  "Starts the worker",
			Action: RunWorker,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port, p",
					Usage: "port the server is on",
					Value: 4730,
				},
				cli.StringFlag{
					Name:  "address, a",
					Usage: "address the server is on, default localhost",
					Value: "",
				},
				cli.StringFlag{
					Name:  "protocol, pr",
					Usage: "web protocol to use, default tcp",
					Value: "tcp",
				},
				cli.StringSliceFlag{
					Name:  "hooks",
					Usage: "hook names for sending log in addition to stderr",
					Value: &cli.StringSlice{},
				},
				cli.StringFlag{
					Name:  "log-level",
					Usage: "log level for the application",
					Value: "error",
				},
				cli.StringFlag{
					Name:   "slack-channel",
					EnvVar: "SLACK_CHANNEL",
					Usage:  "Slack channel where the log will be posted",
				},
				cli.StringFlag{
					Name:   "slack-url",
					EnvVar: "SLACK_URL",
					Usage:  "Slack webhook url[required if slack channel is provided]",
				},
				cli.BoolFlag{
					Name:   "use-logfile",
					EnvVar: "USE_LOG_FILE",
					Usage:  "Instead of stderr, write the script(s) log to a file",
				},
			},
		},
	}
	app.Run(os.Args)

}
