package rpc_client

import (
	"fmt"
	"net/rpc"
)

func rpcCall(method string, port int32, args interface{}, reply interface{}) error {
	// TODO: Move all of the boilerplate code to one place
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return err
	}

	// Synchronous call
	err = client.Call(method, args, reply)
	if err != nil {
		return err
	}

	// Success!
	return nil
}
