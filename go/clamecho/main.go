package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// curl -X POST -d "{\"test\": \"that\"}" http://localhost:8080
// wget --post-data "username=test" http://localhost:8080
// {"apiVersion":"v1alpha","imageName":"","imageID":"sha256:348549493ddfd2eb0827c90738f2add96","containerID":"07caee20a5dcf04bf78070a7f100770344f2d1e3fc2e0542bf5a9c8bc96d4067","results":[{"name":"clamav","scannerVersion":"0.50.1","timestamp":"2017-11-10T17:08:38.871064799Z","reference":"file://root/.viminfo","description":"Spammer.PHP.Private.B(Test).UNOFFICIAL FOUND"},{"name":"clamav","scannerVersion":"0.50.1","timestamp":"2017-11-10T17:08:38.871064799Z","reference":"file://root/test.py","description":"Spammer.PHP.Private.B(Test).UNOFFICIAL FOUND"}]}

func showLocalAddrs() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		fmt.Println(addr.String())
	}
}

// Listen - receive function
func Listen(port int) error {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("Accept error:", err)
		}
		log.Println("accept:", conn.RemoteAddr())

		go func(c net.Conn) {
			io.Copy(os.Stdout, c)
			log.Println("closed:", conn.RemoteAddr())
			defer c.Close()
		}(conn)
	}
}

// Dial - send function
func Dial(host string, port int) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	defer conn.Close()

	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		// To retrieve text through | (pipe)
		buffer, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalln(err)
		}

		_, err = io.Copy(conn, bytes.NewReader(buffer))
	} else {
		_, err = io.Copy(conn, os.Stdin)
	}

	return err
}

func main() {
	port := flag.Int("p", 0, "local port number")
	flag.Usage = func() {
		fmt.Println(strings.Replace(
			`options:
connect to somewhere:	$name [-options] hostname port
listen:			$name -p port
	-p		listen port number`,
			"$name", filepath.Base(os.Args[0]), -1))
	}
	flag.Parse()
	if *port > 0 {
		log.Fatal(Listen(*port))
	}

	if flag.NArg() != 2 {
		flag.Usage()
		return
	}
	dialPort := 0
	fmt.Sscanf(flag.Arg(1), "%d", &dialPort)
	Dial(flag.Arg(0), dialPort)
}
