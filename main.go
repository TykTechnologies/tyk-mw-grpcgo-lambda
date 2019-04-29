package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	coprocess "github.com/TykTechnologies/tyk-protobuf"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	listenAddress = "/tmp/foo.sock"
	awsRegion     = endpoints.EuWest2RegionID
)

func main() {

	go func() {
		<-handleSIGINTKILL()

		log.Println("received termination signal")

		if err := os.Remove(listenAddress); err != nil {
			log.Println(errors.Wrap(err, "unable to unbind, delete sock file manually"))
			os.Exit(1)
			return
		}

		os.Exit(0)
	}()

	listener, err := net.Listen("unix", listenAddress)
	if err != nil {
		log.Println(errors.Wrap(err, "error opening listener"))
		os.Exit(1)
		return
	}

	log.Printf("gRPC server listening on %s\n", listenAddress)

	server := grpc.NewServer()
	coprocess.RegisterDispatcherServer(server, NewDispatcher())

	log.Println(errors.Wrap(server.Serve(listener), "unable to serve"))
}

func handleSIGINTKILL() chan os.Signal {
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	return sig
}
