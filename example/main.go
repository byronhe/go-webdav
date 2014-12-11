// Copyright 2014 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/google/go-webdav"
	"github.com/google/go-webdav/memfs"
)

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(30 * time.Second)
	return tc, nil
}

func main() {
	srv := webdav.NewWebDAV(memfs.NewMemFS())
	srv.Debug = true
	var addr = "0.0.0.0:8080"
	log.Printf("Listening on http://" + addr + "/...")
	var server = &http.Server{Addr: addr, Handler: srv}
	ln, err := net.Listen("tcp4", addr)
	if err != nil {
		panic(err)
	}
	server.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})

}
