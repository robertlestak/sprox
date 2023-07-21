package main

import (
	"flag"
	"os"

	"git.shdw.tech/shdw.tech/sproxy/pkg/sproxy"
	log "github.com/sirupsen/logrus"
)

func init() {
	ll, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		ll = log.InfoLevel
	}
	log.SetLevel(ll)
}

func cmdClient() {
	l := log.WithFields(log.Fields{
		"fn":  "cmdClient",
		"app": "sprox",
	})
	l.Debug("Starting client")
	// usage:
	// sprox connect [options] <address>
	var address *string
	clientFlags := flag.NewFlagSet("client", flag.ExitOnError)
	authToken := clientFlags.String("token", "", "Auth token")
	authTokenCmd := clientFlags.String("token-cmd", "", "Auth token command")
	tcpAddress := clientFlags.String("listen", "127.0.0.1:7654", "Local TCP listener address")
	clientFlags.Parse(os.Args[2:])
	args := clientFlags.Args()
	if len(args) > 0 {
		address = &args[0]
	}
	if *address == "" {
		l.Fatal("Missing server address")
	}
	c := &sproxy.Client{
		Listen:  *tcpAddress,
		Connect: *address,
	}
	if *authToken != "" || *authTokenCmd != "" {
		c.Auth = &sproxy.ClientAuth{
			Token:    *authToken,
			TokenCmd: *authTokenCmd,
		}
		l.Debug("Using auth token")
	}
	if err := c.Run(); err != nil {
		l.WithError(err).Fatal("Error while running client")
	}
}

func cmdServer() {
	l := log.WithFields(log.Fields{
		"fn":  "cmdServer",
		"app": "sprox",
	})
	l.Debug("Starting server")
	// usage:
	// sprox proxy [options] <address>
	serverFlags := flag.NewFlagSet("server", flag.ExitOnError)
	address := serverFlags.String("listen", "0.0.0.0:6543", "Address of the server")
	metricsAddress := serverFlags.String("metrics-address", "0.0.0.0:9090", "Address of the metrics server")
	serverFlags.Parse(os.Args[2:])
	var tcpAddress *string
	args := serverFlags.Args()
	if len(args) > 0 {
		tcpAddress = &args[0]
	}
	s := &sproxy.Server{
		MetricsAddress: *metricsAddress,
		Listen:         *address,
		Connect:        *tcpAddress,
	}
	if err := s.Run(); err != nil {
		l.Errorf("Failed to start sproxy server: %v", err)
	}
}

func usage() {
	log.Printf("Usage: %s <command> [options]", os.Args[0])
	log.Printf("Commands:")
	log.Printf("  connect [options] <address>")
	log.Printf("  proxy [options] <address>")
}

func main() {
	l := log.WithFields(log.Fields{
		"fn":  "main",
		"app": "sprox",
	})
	l.Debug("Starting sprox server")
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		usage()
		os.Exit(0)
	}
	switch os.Args[1] {
	case "connect":
		cmdClient()
	case "proxy":
		cmdServer()
	default:
		l.Fatalf("Unknown command: %s", os.Args[1])
	}
}
