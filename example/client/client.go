package main

import (
	"fmt"
	mdns "mdns/pkg"
	"time"
)

func main() {
	entries := make(chan *mdns.ServiceEntry, 10)
	exit := make(chan struct{})
	p := &mdns.QueryParam{
		Service:             "provider",
		Domain:              "micro.",
		Timeout:             time.Second,
		Entries:             entries,
		WantUnicastResponse: false,
	}

	go func() {
		for {
			select {
			case e := <-entries:
				fmt.Println(e)
			}
		}
	}()

	// execute query
	if err := mdns.Query(p); err != nil {
		fmt.Println(err)
		return
	}

	mdns.Listen(entries, exit)
	time.Sleep(1000 * time.Second)
}
