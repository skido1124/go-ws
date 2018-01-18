// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"text/template"
)

var addr = flag.String("addr", ":" + os.Getenv("PORT"), "http service address")
var homeTempl = template.Must(template.ParseFiles("home.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is nothing")
	}

	log.SetFlags(log.Lshortfile)
	flag.Parse()
	go h.run()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/ws/{room}", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})
	//err := http.ListenAndServe(*addr, r)
	err := http.ListenAndServe(":" + port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("ListenAndServe: success port" + port)
	}
}
