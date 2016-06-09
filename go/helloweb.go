// helloweb is a simple web-server that says "Hello World" for every path
//
// helloweb [--logfile <logfile>] <bind-ip> <bind-port> ...
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func asctime() string {
	return time.Now().Format(time.ANSIC)
}

// wrapper for http.Handler to log all requests
func logit(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

var name string

func webhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s at `%s`  ; %s  (go)\n", name,
		r.URL.Path, asctime())
}

func main() {
	logfile := flag.String("logfile", "", "log output to file instead of stderr")
	flag.Usage = func() {
		fmt.Println("Usage: helloweb.go [options] bindIP bindPort ...")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	bindIP := flag.Arg(0)
	bindPort := flag.Arg(1)
	bindAddr := net.JoinHostPort(bindIP, bindPort)

	name = strings.Join(flag.Args()[2:], " ")
	if name == "" {
		name = "world"
	}

	// redirect log to file, if requested
	if *logfile != "" {
		f, err := os.OpenFile(*logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	log.Printf("* %s helloweb.go starting at %s", asctime(), bindAddr)

	http.HandleFunc("/", webhello)
	log.Fatal(
		http.ListenAndServe(bindAddr, logit(http.DefaultServeMux)))
}
