package rpc_server

import (
	"net"
	"net/http"
	"net/rpc"
	"sync"
)

func StartInternalRPC() (*sync.WaitGroup, error) {
	rpc_fns := new(Repo)
	rpc.Register(rpc_fns)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		return nil, e
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Start the http RPC server in a goroutine
	go func() {
		// Decrement the counter when the goroutine completes
		defer wg.Done()
		http.Serve(l, nil)
	}()

	return &wg, nil
}
