package main

import (
	"fmt"
	mdns "mdns/pkg"
	"net"
	"time"
)

func main() {
	s, err := mdns.NewMDNSService(
		"xxxx",
		"xxxx",
		"micro"+".",
		"",
		9999,
		[]net.IP{net.ParseIP("127.0.0.1")},
		[]string{"x"},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	server, err := mdns.NewServer(&mdns.Config{Zone: s})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(server)

	time.Sleep(100 * time.Second)
}
