package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"strconv"

	"github.com/appscode/g2/worker"
	"github.com/mikespook/golib/signal"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func RunWorker(c *cli.Context) error {

	log := getLogger(c)

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
	//Blastn runs the blastx program and returns result in job.Data
	Blastn := func(job worker.Job) ([]byte, error) {
		//unmarshal the Arguments
		args := Arguments{}
		err := json.Unmarshal(job.Data(), &args)
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}
		log.WithFields(logrus.Fields{
			"program":  "Blastn",
			"database": args.Database,
			"query":    args.Query,
			"evalue":   args.Evalue,
			"numalign": args.Numalign,
			"wordsize": args.Wordsize,
			"matrix":   args.Matrix,
			"seg":      args.Seg,
			"gapped":   args.Gapped,
		}).Info("Parameters")
		evalue := strconv.FormatFloat(args.Evalue, 'f', -1, 64)

		cmd := exec.Command("blastn", "-db", args.Database, "-query", args.Query, "-evalue", evalue, "-num_alignments", string(args.Numalign), "-matrix", args.Matrix)
		if args.Seg {
			cmd.Args = append(cmd.Args, "-seg")
			cmd.Args = append(cmd.Args, "yes")
		}
		// if !args.Gapped {
		// 	cmd.Args = append(cmd.Args, "--ungapped")
		// }
		out, err := cmd.Output()
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}
		return out, nil
	}

	//Blastp runs the blastp program and returns result in job.Data
	Blastp := func(job worker.Job) ([]byte, error) {
		//unmarshal the Arguments
		args := Arguments{}
		err := json.Unmarshal(job.Data(), &args)
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}
		log.WithFields(logrus.Fields{
			"program":  "Blastp",
			"database": args.Database,
			"query":    args.Query,
			"evalue":   args.Evalue,
			"numalign": args.Numalign,
			"wordsize": args.Wordsize,
			"matrix":   args.Matrix,
			"seg":      args.Seg,
			"gapped":   args.Gapped,
		}).Info("Parameters")
		evalue := strconv.FormatFloat(args.Evalue, 'f', -1, 64)

		cmd := exec.Command("blastp", "-db", args.Database, "-query", args.Query, "-evalue", evalue, "-num_alignments", string(args.Numalign), "-matrix", args.Matrix)
		if args.Seg {
			cmd.Args = append(cmd.Args, "-seg")
			cmd.Args = append(cmd.Args, "yes")
		}
		// if !args.Gapped {
		// 	cmd.Args = append(cmd.Args, "--ungapped")
		// }
		out, err := cmd.Output()
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}
		return out, nil
	}

	//Blastx runs the blastx program and returns result in job.Data
	Blastx := func(job worker.Job) ([]byte, error) {
		//unmarshal the Arguments
		args := Arguments{}
		err := json.Unmarshal(job.Data(), &args)
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}
		log.WithFields(logrus.Fields{
			"program":  "Blastx",
			"database": args.Database,
			"query":    args.Query,
			"evalue":   args.Evalue,
			"numalign": args.Numalign,
			"wordsize": args.Wordsize,
			"matrix":   args.Matrix,
			"seg":      args.Seg,
			"gapped":   args.Gapped,
		}).Info("Parameters")
		evalue := strconv.FormatFloat(args.Evalue, 'f', -1, 64)

		cmd := exec.Command("blastx", "-db", args.Database, "-query", args.Query, "-evalue", evalue, "-num_alignments", string(args.Numalign), "-matrix", args.Matrix)
		if args.Seg {
			cmd.Args = append(cmd.Args, "-seg")
			cmd.Args = append(cmd.Args, "yes")
		}
		// if !args.Gapped {
		// 	cmd.Args = append(cmd.Args, "--ungapped")
		// }
		out, err := cmd.Output()
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}
		return out, nil
	}
	//Tblastx runs the blastx program and returns result in job.Data
	Tblastx := func(job worker.Job) ([]byte, error) {
		//unmarshal the Arguments
		args := Arguments{}
		err := json.Unmarshal(job.Data(), &args)
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}
		log.WithFields(logrus.Fields{
			"program":  "Tblastx",
			"database": args.Database,
			"query":    args.Query,
			"evalue":   args.Evalue,
			"numalign": args.Numalign,
			"wordsize": args.Wordsize,
			"matrix":   args.Matrix,
			"seg":      args.Seg,
			"gapped":   args.Gapped,
		}).Info("Parameters")
		evalue := strconv.FormatFloat(args.Evalue, 'f', -1, 64)

		cmd := exec.Command("tblastx", "-db", args.Database, "-query", args.Query, "-evalue", evalue, "-num_alignments", string(args.Numalign), "-matrix", args.Matrix)
		if args.Seg {
			cmd.Args = append(cmd.Args, "-seg")
			cmd.Args = append(cmd.Args, "yes")
		}
		// if !args.Gapped {
		// 	cmd.Args = append(cmd.Args, "--ungapped")
		// }
		out, err := cmd.Output()
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}
		return out, nil
	}
	//Tblastn runs the blastx program and returns result in job.Data
	Tblastn := func(job worker.Job) ([]byte, error) {
		//unmarshal the Arguments
		args := Arguments{}
		err := json.Unmarshal(job.Data(), &args)
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}
		log.WithFields(logrus.Fields{
			"program":  "Tblastn",
			"database": args.Database,
			"query":    args.Query,
			"evalue":   args.Evalue,
			"numalign": args.Numalign,
			"wordsize": args.Wordsize,
			"matrix":   args.Matrix,
			"seg":      args.Seg,
			"gapped":   args.Gapped,
		}).Info("Parameters")
		evalue := strconv.FormatFloat(args.Evalue, 'f', -1, 64)

		cmd := exec.Command("tblastn", "-db", args.Database, "-query", args.Query, "-evalue", evalue, "-num_alignments", string(args.Numalign), "-matrix", args.Matrix)
		if args.Seg {
			cmd.Args = append(cmd.Args, "-seg")
			cmd.Args = append(cmd.Args, "yes")
		}
		// if !args.Gapped {
		// 	cmd.Args = append(cmd.Args, "--ungapped")
		// }
		out, err := cmd.Output()
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}
		return out, nil
	}

	address := c.String("address") + ":" + c.String("port")
	w.AddServer(c.String("protocol"), address)

	log.WithFields(logrus.Fields{
		"address": address,
		"port":    c.String("port"),
	}).Info("address worker pointed at")

	w.AddFunc("Blastn", Blastn, worker.Unlimited)
	w.AddFunc("Blastp", Blastp, worker.Unlimited)
	w.AddFunc("Blastx", Blastx, worker.Unlimited)
	w.AddFunc("Tblastx", Tblastx, worker.Unlimited)
	w.AddFunc("Tblastn", Tblastn, worker.Unlimited)

	if err := w.Ready(); err != nil {
		log.Error(err)
		return err
	}

	go w.Work()
	signal.Bind(os.Interrupt, func() uint { return signal.BreakExit })
	signal.Wait()

	return nil
}
