package rpc_server

import (
	"net"
	"net/http"
	"net/rpc"
)

func StartInternalRPC() error {
	rpc_fns := new(Repo)
	rpc.Register(rpc_fns)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		return e
	}

	// TODO: Make sure that this is process isolated!
	go http.Serve(l, nil)
	return nil
}
