package main

//
// //Blastp runs the blastp program and returns result in job.Data
// import (
// 	"encoding/json"
// 	"log"
// 	"os/exec"
// 	"strconv"
// 	"testing"
//
// 	"github.com/appscode/g2/worker"
// )
//
// func Blastp(job worker.Job) ([]byte, error) {
// 	//unmarshal the Arguments
// 	args := Arguments{}
// 	err := json.Unmarshal(job.Data(), &args)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	evalue := strconv.FormatFloat(args.Evalue, 'f', -1, 64)
//
// 	cmd := exec.Command("blastp", "-db", args.Database, "-query", args.Query, "-evalue", evalue, "-num_alignments", string(args.Numalign), "-matrix", args.Matrix)
// 	if args.Seg {
// 		cmd.Args = append(cmd.Args, "-seg")
// 		cmd.Args = append(cmd.Args, "yes")
// 	}
// 	// if !args.Gapped {
// 	// 	cmd.Args = append(cmd.Args, "--ungapped")
// 	// }
// 	out, err := cmd.Output()
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	return out, nil
// }
//
// func TestBlastp(t *testing.T) {
// 	a := &Arguments{
// 		Database: "dicty_primary_protein",
// 		Query:    "test_query.fsa",
// 		Evalue:   0.1,
// 		Numalign: 50,
// 		Wordsize: 3,
// 		Matrix:   "PAM30",
// 		Seg:      true,
// 		//Gapped:   false,
// 	}
// 	args, err := json.Marshal(a)
// 	if err != nil {
// 		log.Println("error marshaling")
// 		log.Fatal(err)
// 	}
//
// 	var j *worker.Job
// 	&j.Handle() = "bla"
//
// }
