package main 

import "github.com/go-ping/ping"

pinger, err := ping.NewPinger("www.google.com")
if err != nil {
	panic(err)
}
pinger.Count = 3
pinger.Run()                 // blocks until finished
stats := pinger.Statistics() // get send/receive/rtt stats