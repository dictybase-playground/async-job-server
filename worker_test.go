package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"testing"

	"github.com/appscode/g2/client"
	"github.com/appscode/g2/pkg/runtime"

	dockertest "gopkg.in/ory-am/dockertest.v3"
)

var port string

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{Repository: "appscode/gearmand", Tag: "0.5.1", Cmd: []string{"run", "--v=5", "--storage-dir=/var/db"}, ExposedPorts: []string{"4730"}})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	//port is dynamically assigned based on availability
	port = resource.GetPort("4730/tcp")

	//make database
	makedb := exec.Command("makeblastdb", "-in", "dicty_primary_protein", "-dbtype", "prot")
	err = makedb.Run()
	if err != nil {
		fmt.Println("error makin db")
		log.Fatal(err)
	}

	//runs all the other tests
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	//delete database
	deletedb1 := exec.Command("rm", "dicty_primary_protein.phr")
	deletedb2 := exec.Command("rm", "dicty_primary_protein.pin")
	deletedb3 := exec.Command("rm", "dicty_primary_protein.psq")

	err = deletedb1.Run()
	if err != nil {
		log.Fatal(err)
	}
	err = deletedb2.Run()
	if err != nil {
		log.Fatal(err)
	}
	err = deletedb3.Run()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

func TestBlastp(t *testing.T) {

	//start worker
	cmd := exec.Command("async-job-server", "run", "-p", port)
	err := cmd.Start()
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	c, err := client.New("tcp", ":"+port)
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
