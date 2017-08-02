package main

import (
	"encoding/json"
	"os/exec"
	"strconv"

	"github.com/appscode/g2/worker"
	"github.com/sirupsen/logrus"
)

func (env *Env) Blastn(job worker.Job) ([]byte, error) {
	//unmarshal the Arguments
	log := env.logger
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

	cmd := exec.Command("blastn", "-db", args.Database, "-query", args.Query, "-evalue", evalue, "-num_alignments", string(args.Numalign), "-matrix", args.Matrix, "-outfmt", "15")
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
func (env *Env) Blastp(job worker.Job) ([]byte, error) {
	log := env.logger
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

	cmd := exec.Command("blastp", "-db", args.Database, "-query", args.Query, "-evalue", evalue, "-num_alignments", string(args.Numalign), "-matrix", args.Matrix, "-outfmt", "15")
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
func (env *Env) Blastx(job worker.Job) ([]byte, error) {
	log := env.logger
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

	cmd := exec.Command("blastx", "-db", args.Database, "-query", args.Query, "-evalue", evalue, "-num_alignments", string(args.Numalign), "-matrix", args.Matrix, "-outfmt", "15")
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
func (env *Env) Tblastx(job worker.Job) ([]byte, error) {
	log := env.logger
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

	cmd := exec.Command("tblastx", "-db", args.Database, "-query", args.Query, "-evalue", evalue, "-num_alignments", string(args.Numalign), "-matrix", args.Matrix, "-outfmt", "15")
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
func (env *Env) Tblastn(job worker.Job) ([]byte, error) {
	log := env.logger
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

	cmd := exec.Command("tblastn", "-db", args.Database, "-query", args.Query, "-evalue", evalue, "-num_alignments", string(args.Numalign), "-matrix", args.Matrix, "-outfmt", "15")
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
