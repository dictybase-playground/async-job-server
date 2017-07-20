package main

import (
	"os"

	"github.com/johntdyer/slackrus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const layout = "2006-01-02_150405"

func getLogger(c *cli.Context) *logrus.Logger {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "02/Jan/2006:15:04:05",
	}
	log.Out = os.Stderr
	if c.IsSet("use-logfile") {
		w, err := os.OpenFile("mylogfile2.log", os.O_WRONLY|os.O_CREATE, 0666)
		if err == nil {
			log.Out = w
		} else {
			log.Info("Failed to log to file, using stderr")
		}
	}
	l := c.GlobalString("log-level")
	switch l {
	case "debug":
		log.Level = logrus.DebugLevel
	case "warn":
		log.Level = logrus.WarnLevel
	case "error":
		log.Level = logrus.ErrorLevel
	case "fatal":
		log.Level = logrus.FatalLevel
	case "panic":
		log.Level = logrus.PanicLevel
	}
	// Set up hook
	lh := make(logrus.LevelHooks)
	for _, h := range c.StringSlice("hooks") {
		switch h {
		case "slack":
			lh.Add(&slackrus.SlackrusHook{
				HookURL:        c.String("slack-url"),
				AcceptedLevels: slackrus.LevelThreshold(log.Level),
				IconEmoji:      ":skull:",
			})
		default:
			continue
		}
	}
	log.Hooks = lh
	return log
}
