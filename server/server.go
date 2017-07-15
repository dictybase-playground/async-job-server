package main

import "github.com/appscode/g2/pkg/server"

//test
func main() {
	cfg := server.Config{
		ListenAddr: ":1234",
		Storage:    "234",
		WebAddress: "localhost:1234",
	}
	s := server.NewServer(cfg)
	s.Start()
}
