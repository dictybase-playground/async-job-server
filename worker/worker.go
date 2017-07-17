package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/appscode/g2/worker"
	"github.com/mikespook/golib/signal"
)

type Arguments struct {
	Database string `json:"database"`
	Query    string `json:"query"`
}

// A function for handling jobs
func Foobar(job worker.Job) ([]byte, error) {
	log.Printf("ToUpper: Data=[%s]\n", job.Data())
	data := []byte(strings.ToUpper(string(job.Data())))
	log.Println(string(data))
	log.Println(data)
	return data, nil
}

func Blastp(job worker.Job) ([]byte, error) {
	//unmarshal the Arguments
	log.Println(job.Data())

	args := Arguments{}
	err := json.Unmarshal(job.Data(), &args)
	if err != nil {
		log.Println(err.Error())
		return []byte{}, err
	}
	log.Println(args)

	cmd := exec.Command("blastp", "-db", args.Database, "-query", args.Query)
	out, err := cmd.Output()
	if err != nil {
		log.Print("output err")
		log.Fatal(err)
	}
	return out, nil
}

//blastp -db dicty_primary_protein -query test_query.fsa
// -evalue 0.1 -num_alignments 50 -word_size 3 -seg 'yes'

//
// cmd := exec.Command("blastp", "-db", "dicty_primary_protein", "-query", "test_query.fsa")
// var out bytes.Buffer
// cmd.Stdout = &out
// err := cmd.Run()
//
// if err != nil {
//     fmt.Println("error!!!!")
//     fmt.Printf(err.Error())
//     log.Fatal(err)
// }
// fmt.Println("done")
// fmt.Print(out.String())
func main() {

	log.Println("Starting ...")
	defer log.Println("Shutdown complete!")
	w := worker.New(worker.Unlimited)
	defer w.Close()
	w.ErrorHandler = func(e error) {
		log.Println(e)
	}
	w.JobHandler = func(job worker.Job) error {
		log.Printf("H=%s, UID=%s, Data=%s", job.Handle,
			job.UniqueId, job.Data)
		return nil
	}
	w.AddServer("tcp", ":4730")
	w.AddFunc("Foobar", Foobar, worker.Unlimited)
	w.AddFunc("Blastp", Blastp, worker.Unlimited)

	if err := w.Ready(); err != nil {
		log.Fatal(err)
		return
	}
	go w.Work()
	signal.Bind(os.Interrupt, func() uint { return signal.BreakExit })
	signal.Wait()

}
