package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/appscode/g2/client"
	"github.com/appscode/g2/pkg/runtime"
)

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
func TestBlastp(t *testing.T) {
	var wg sync.WaitGroup
	c, err := client.New("tcp", ":4730")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	c.ErrorHandler = func(e error) {
		t.Fatal(e)
		os.Exit(1)
	}

	jobHandler := func(resp *client.Response) {
		log.Printf("%s", resp.Data)
		wg.Done()

	}

	a := &Arguments{
		Database: "dicty_primary_protein",
		Query:    "test_query.fsa",
		Evalue:   0.1,
		Numalign: 50,
		Wordsize: 3,
		Matrix:   "PAM30",
		Seg:      true,
		//Gapped:   false,
	}
	args, err := json.Marshal(a)
	if err != nil {
		log.Println("error marshaling")
		t.Fatal(err)
	}
	log.Println(args)

	handle, err := c.Do("Blastx", args, runtime.JobNormal, jobHandler)
	if err != nil {
		t.Fatal(err)
	}
	wg.Add(1)

	status, err := c.Status(handle)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%v", *status)

	handle2, err := c.Do("Blastp", args, runtime.JobNormal, jobHandler)
	if err != nil {
		t.Fatal(err)
	}
	wg.Add(1)
	status, err = c.Status(handle2)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%v", *status)
	wg.Wait()
}
