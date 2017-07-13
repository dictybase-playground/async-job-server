package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

//blastp -db dicty_primary_protein -query test_query.fsa
// -evalue 0.1 -num_alignments 50 -word_size 3 -seg 'yes'
func main() {
	cmd := exec.Command("tr", "-s", "s")
	cmd.Stdin = strings.NewReader("sssss input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		fmt.Println("error!!!!")
		fmt.Printf(err.Error())
		log.Fatal(err)
	}
	fmt.Println("done")
	fmt.Print(out.String())
}
