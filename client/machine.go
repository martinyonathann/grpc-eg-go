package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/martinyonathann/grpc-eg-go/machine"
)

var (
	serverAddr = flag.String("server_addr", "localhost:9111", "The server address in the format of host:port")
)

func runExecute(client machine.MachineClient, instructions *machine.InstructionSet) {
	log.Printf("Executing  %v", instructions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := client.Execute(ctx, instructions)
	if err != nil {
		log.Fatalf("%v.Execute(_) = _, %v: ", client, err)
	}
	log.Println(result)
}
