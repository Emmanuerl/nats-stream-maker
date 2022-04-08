package main

import (
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"strings"
)

func main() {
	natsURI, err := getEnv("NATS_URI")
	if err != nil {
		gracefulExit(err)
	}

	streamNames, err := getEnv("STREAMS")
	if err != nil {
		gracefulExit(err)
	}

	nc, err := nats.Connect(natsURI)
	if err != nil {
		gracefulExit(err)
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
		return value, errors.New(fmt.Sprintf("env variable %s not found\n", variable))
	}
	return value, nil
}

func gracefulExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
