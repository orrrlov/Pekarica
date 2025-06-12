package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

var (
	domain = flag.String("domain", "localhost", "domain")
	port   = flag.String("port", "8080", "port")
	done   = make(chan os.Signal, 1)
)

func main() {
	flag.Parse()

	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	srv := initialize(*domain, *port)

	go srv.run()

	openBrowser(url(*domain, *port))

	<-done

	srv.shutdown()
}
