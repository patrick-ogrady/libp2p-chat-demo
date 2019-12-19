package main

import (
	"context"
	"flag"
	"log"

	golog "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/peer"
	gologging "github.com/whyrusleeping/go-logging"
)

func main() {
	// Parse options from the command line
	listenF := flag.Int("p", 0, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
	verbose := flag.Bool("v", false, "print debug logs")
	flag.Parse()

	if *listenF == 0 {
		log.Fatal("Please provide a port to bind on with -p")
	}

	if *verbose == true {
		// LibP2P code uses golog to log messages. They log with different
		// string IDs (i.e. "swarm"). We can control the verbosity level for
		// all loggers with:
		golog.SetAllLoggers(gologging.INFO)
	} else {
		golog.SetAllLoggers(gologging.ERROR)
	}

	ha, err := makeRoutedHost(*listenF)
	if err != nil {
		log.Fatal(err)
	}

	// Set a stream handler on host A. /chat/1.0.0 is
	// a user-defined protocol name. The stream handler
	// is called when another peer connects to us.
	ha.SetStreamHandler("/chat/1.0.0", handleStream)

	if *target == "" {
		log.Printf("Run \"./chat -p %d -d %s\" on a different terminal or computer to start a chat\n", *listenF+1, ha.ID().Pretty())
		log.Println("listening for streams....")
		select {} // hang forever
	}
	/**** This is where the listener code ends ****/

	peerid, err := peer.IDB58Decode(*target)
	if err != nil {
		log.Fatalln(err)
	}

	// make a new stream from host B to host A
	// it should be handled on host A by the handler
	// we set above because we use the same /chat/1.0.0 protocol
	s, err := ha.NewStream(context.Background(), peerid, "/chat/1.0.0")
	if err != nil {
		log.Fatalln(err)
	} else {
		handleStream(s)
	}

	select {}
}
