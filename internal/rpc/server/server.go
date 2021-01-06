package rpc_server

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"sync"

	"gogs.io/gogs/internal/route"

	"webauthn/webauthn"
)

// TODO: Make this more legit
//
// A hacky way to handle user sessions. Maps UserID -> SessionData
var sessionMap = map[int64]webauthn.SessionData{}

func InitServer(customConf string) error {
	// Initilize the database
	if err := route.InitOnlyDB(customConf); err != nil {
		return err
	}

	// Success!
	return nil
}

func StartInternalRPC(rpc_fns interface{}, port int32) (*sync.WaitGroup, error) {
	rpc.Register(rpc_fns)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", fmt.Sprintf(":%d", port))
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
