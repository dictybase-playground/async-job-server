package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/appscode/g2/worker"
	"github.com/mikespook/golib/signal"
)

// A function for handling jobs
func Foobar(job worker.Job) ([]byte, error) {
	log.Printf("ToUpper: Data=[%s]\n", job.Data())
	fmt.Println("in foobar")
	data := []byte(strings.ToUpper(string(job.Data())))
	return data, nil
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
	w.AddServer("tcp", ":1234") //unsure about specs, copied from example
	w.AddFunc("Foobar", Foobar, worker.Unlimited)

	if err := w.Ready(); err != nil {
		log.Fatal(err)
		return
	}
	go w.Work()
	signal.Bind(os.Interrupt, func() uint { return signal.BreakExit })
	signal.Wait()

}
