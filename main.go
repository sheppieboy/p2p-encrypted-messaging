package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grandcat/zeroconf"
)

var (
    name     = flag.String("name", "MyChatApp", "The name for the chat service.")
    service  = flag.String("service", "_chat._tcp", "Set the service type of the chat service.")
    domain   = flag.String("domain", "local.", "Set the network domain. Default should be fine.")
    port     = flag.Int("port", 5454, "Set the port the chat service is listening to.")
    waitTime = flag.Int("wait", 10, "Duration in [s] to publish service for.")
)

func main() {
    flag.Parse()

    server, err := zeroconf.Register(*name, *service, *domain, *port, []string{"txtv=0", "lo=1", "la=2"}, nil)
    if err != nil {
        log.Fatalf("Failed to register chat service: %v", err)
    }
    defer server.Shutdown()
    log.Println("Published chat service.")

    // Discover other chat services
    resolver, err := zeroconf.NewResolver(nil)
    if err != nil {
        log.Fatalf("Failed to create resolver: %v", err)
    }
    entries := make(chan *zeroconf.ServiceEntry)
    go func(results <-chan *zeroconf.ServiceEntry) {
        for entry := range results {
            log.Printf("Found chat service: %s:%d,%s ", entry.AddrIPv4, entry.Port, entry.HostName)
            // Connect to the chat service
            // Establish chat connection with entry.AddrIPv4 and entry.Port
        }
    }(entries)
    ctx := context.Background()
    err = resolver.Browse(ctx, *service, *domain, entries)
    if err != nil {
        log.Fatalf("Failed to browse chat services: %v", err)
    }

    // Clean exit.
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
    // Timeout timer.
    var tc <-chan time.Time
    if *waitTime > 0 {
        tc = time.After(time.Second * time.Duration(*waitTime))
    }

    select {
    case <-sig:
        // Exit by user
    case <-tc:
        // Exit by timeout
    }

    log.Println("Shutting down.")
}
