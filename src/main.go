package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nats-io/nats.go"
)

func main() {
	natsURI, err := getEnv("NATS_URI")
	if err != nil {
		log.Fatal(err)
	}

	streamNames, err := getEnv("STREAMS")
	if err != nil {
		log.Fatal(err)
	}

	nc, err := nats.Connect(natsURI)
	if err != nil {
		log.Fatal(err)
	}

	js, _ := nc.JetStream()
	streams := strings.Split(streamNames, ",")
	for _, v := range streams {
		stream, err := createStream(js, v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("stream %q has been successfully created \n", stream.Config.Name)
		}
	}
}

func createStream(js nats.JetStreamContext, name string) (*nats.StreamInfo, error) {
	subject := fmt.Sprintf("%s.>", name)
	return js.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: []string{subject},
	})
}

func getEnv(variable string) (string, error) {
	value := os.Getenv(variable)
	if len(value) == 0 {
		return value, fmt.Errorf("env variable %s not found", variable)
	}
	return value, nil
}
