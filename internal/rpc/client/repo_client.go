package rpc_client

import (
	shared "gogs.io/gogs/internal/rpc/shared"
	"net/rpc"

	"webauthn/protocol"
)

func rpcCall(method string, args interface{}, reply interface{}) error {
	// TODO: Move all of the boilerplate code to one place
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
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

func Repo_DeleteRepositoryBegin(args *shared.Repo_DeleteRepositoryBeginArgs, reply *protocol.CredentialAssertion) error {
	return rpcCall("Repo.DeleteRepositoryBegin", args, reply)
}

func Repo_DeleteRepositoryFinish(args *shared.Repo_DeleteRepositoryFinishArgs, reply interface{}) error {
	return rpcCall("Repo.DeleteRepositoryFinish", args, reply)
}

func Repo_DeleteRepository(args *shared.Repo_DeleteRepositoryArgs, reply interface{}) error {
	return rpcCall("Repo.DeleteRepository", args, reply)
}
