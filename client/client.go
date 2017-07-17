package main

import (
	"log"
	"os"
	"sync"

	"github.com/appscode/g2/client"
	"github.com/appscode/g2/pkg/runtime"
)

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
	handle, err := c.Do("Foobar", echo, runtime.JobNormal, jobHandler)
	if err != nil {
		log.Fatalln(err)
	}
	wg.Add(1)

	log.Println(string(handle))
	status, err := c.Status(handle)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%v", *status)
	wg.Wait()
}
