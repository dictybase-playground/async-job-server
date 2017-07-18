package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/appscode/g2/client"
	"github.com/appscode/g2/pkg/runtime"
)

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
	var wg sync.WaitGroup
	c, err := client.New("tcp", ":4730")
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	c.ErrorHandler = func(e error) {
		log.Println(e)
		os.Exit(1)
	}
	echo := []byte("Hello\x00 world")
	echomsg, err := c.Echo(echo)
	if err != nil {
		log.Fatalln(err)

	}
	log.Println(string(echomsg))

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
		log.Fatal(err)
	}
	log.Println(args)

	handle, err := c.Do("Blastp", args, runtime.JobNormal, jobHandler)
	wg.Add(1)

	log.Println(string(handle))
	status, err := c.Status(handle)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%v", *status)
	wg.Wait()
}
