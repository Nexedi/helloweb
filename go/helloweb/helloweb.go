// Copyright (C) 2015-2021  Nexedi SA and Contributors.
//                          Kirill Smelkov <kirr@nexedi.com>
//                          Gabriel Monnerat <gabriel@nexedi.com>
//
// This program is free software: you can Use, Study, Modify and Redistribute
// it under the terms of the GNU General Public License version 3, or (at your
// option) any later version, as published by the Free Software Foundation.
//
// You can also Link and Combine this program with other software covered by
// the terms of any of the Free Software licenses or any of the Open Source
// Initiative approved licenses and Convey the resulting work. Corresponding
// source of such a combination shall include the source code for all other
// software used.
//
// This program is distributed WITHOUT ANY WARRANTY; without even the implied
// warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
//
// See COPYING file for full licensing terms.
// See https://www.nexedi.com/licensing for rationale and options.

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
	"runtime"
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
	fmt.Fprintf(w, "Hello %s at `%s`  ; %s  (go %s)\n", name,
		r.URL.Path, asctime(), runtime.Version())
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
