package main

import (
	"os"

	"github.com/appscode/g2/worker"
	"github.com/mikespook/golib/signal"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

//Env holds the logger that we pass to BLAST functions
type Env struct {
	logger *logrus.Logger
}

func RunWorker(c *cli.Context) error {

	env := &Env{logger: getLogger(c)}
	log := env.logger

	log.Info("Starting ...")
	defer log.Info("Shutdown complete!")

	w := worker.New(worker.Unlimited)
	defer w.Close()

	w.ErrorHandler = func(e error) {
		log.Error(e)
	}

	w.JobHandler = func(job worker.Job) error {
		log.WithFields(logrus.Fields{
			"handle:":  job.Handle,
			"uniqueid": job.UniqueId,
			"data":     job.Data,
		}).Info("Job handler")

		return nil
	}

	address := c.String("address") + ":" + c.String("port")
	w.AddServer(c.String("protocol"), address)

	log.WithFields(logrus.Fields{
		"address": address,
		"port":    c.String("port"),
	}).Info("address worker pointed at")

	w.AddFunc("Blastn", env.Blastn, worker.Unlimited)
	w.AddFunc("Blastp", env.Blastp, worker.Unlimited)
	w.AddFunc("Blastx", env.Blastx, worker.Unlimited)
	w.AddFunc("Tblastx", env.Tblastx, worker.Unlimited)
	w.AddFunc("Tblastn", env.Tblastn, worker.Unlimited)

	if err := w.Ready(); err != nil {
		log.Error(err)
		return err
	}

	go w.Work()
	signal.Bind(os.Interrupt, func() uint { return signal.BreakExit })
	signal.Wait()

	return nil
}
