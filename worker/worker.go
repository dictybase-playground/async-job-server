package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"

	"github.com/appscode/g2/worker"
	"github.com/mikespook/golib/signal"
)

type Arguments struct {
	Database string  `json:"database"`
	Query    string  `json:"query"`
	Evalue   float64 `json:"evalue"`
	Numalign int     `json:"numalign"`
	Wordsize int     `json:"wordsize"`
}

func Blastp(job worker.Job) ([]byte, error) {
	//unmarshal the Arguments
	args := Arguments{}
	err := json.Unmarshal(job.Data(), &args)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("blastp", "-db", args.Database, "-query", args.Query)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return out, nil
}

func main() {

	log.Println("Starting ...")
	defer log.Println("Shutdown complete!")
	w := worker.New(worker.Unlimited)
	defer w.Close()
	w.ErrorHandler = func(e error) { //unsure if i need this function
		log.Println(e)
	}
	w.JobHandler = func(job worker.Job) error { //or this function
		log.Printf("H=%s, UID=%s, Data=%s", job.Handle,
			job.UniqueId, job.Data)
		return nil
	}

	w.AddServer("tcp", ":4730")
	w.AddFunc("Blastp", Blastp, worker.Unlimited)

	if err := w.Ready(); err != nil {
		log.Fatal(err)
		return
	}
	go w.Work()
	signal.Bind(os.Interrupt, func() uint { return signal.BreakExit })
	signal.Wait()

}
