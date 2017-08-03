package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"sync"
	"testing"

	"github.com/appscode/g2/client"
	"github.com/appscode/g2/pkg/runtime"
	"github.com/appscode/g2/worker"
	"github.com/sirupsen/logrus"

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
		log.Fatalf("Error creating database: %s", err)
	}

	//runs all the other tests
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	//delete database

	os.Remove("dicty_primary_protein.phr")
	os.Remove("dicty_primary_protein.pin")
	os.Remove("dicty_primary_protein.psq")

	os.Exit(code)
}

func TestBlastp(t *testing.T) {
	//start worker
	log := logrus.New()
	env := &Env{logger: log}
	w := worker.New(1)
	defer w.Close()
	w.AddServer("tcp", ":"+port)
	w.AddFunc("Blastp", env.Blastp, worker.Unlimited)
	if err := w.Ready(); err != nil {
		t.Fatalf("Error starting worker:%s", err)
	}
	go w.Work()

	//start client
	var wg sync.WaitGroup
	c, err := client.New("tcp", ":"+port)
	if err != nil {
		t.Fatalf("Error starting client: %s", err)
	}
	defer c.Close()

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
		t.Fatalf("Error marshaling data: %s", err)
	}
	c.ErrorHandler = func(e error) {
		t.Fatal(e)
		os.Exit(1)
	}

	//jobHandler gets the result from the worker and checks if its what is expected
	jobHandler := func(resp *client.Response) {
		expectcmd := exec.Command("blastp", "-db", a.Database, "-query", a.Query, "-evalue", strconv.FormatFloat(a.Evalue, 'f', -1, 64), "-num_alignments", string(a.Numalign), "-matrix", a.Matrix, "-outfmt", "15")
		if a.Seg {
			expectcmd.Args = append(expectcmd.Args, "-seg")
			expectcmd.Args = append(expectcmd.Args, "yes")
		}
		expected, err := expectcmd.Output()
		if err != nil {
			t.Fatalf("Error getting expected output: %s", err)
		}
		actual := resp.Data
		var a interface{}
		var e interface{}
		err = json.Unmarshal(actual, &a)
		if err != nil {
			t.Fatalf("Error unmrashaling: %s", err)
		}
		err = json.Unmarshal(expected, &e)
		if err != nil {
			t.Fatalf("Error unmrashaling: %s", err)

		}
		eq := reflect.DeepEqual(a, e)
		if !eq {
			t.Fatalf("Actual result from Blastp does not equal expect result")
		}

		wg.Done()

	}
	_, err = c.Do("Blastp", args, runtime.JobNormal, jobHandler)
	if err != nil {
		t.Fatalf("Error sending job to worker: %s", err)
	}
	wg.Add(1)
	wg.Wait()
}
