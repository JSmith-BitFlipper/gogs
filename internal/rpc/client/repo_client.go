package rpc_client

import (
	shared "gogs.io/gogs/internal/rpc/shared"
	"net/rpc"
)

func GetWebauthnUserOptions() error {
	// TOOD: Get the user options from the rpc_server, such as the challenge, remembering it, etc.
}

func Repo_DeleteRepository(args *shared.Repo_DeleteRepositoryArgs, reply *interface{}) error {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		return err
	}

	// Synchronous call
	err = client.Call("Repo.DeleteRepository", args, reply)
	if err != nil {
		return err
	}

	// Success!
	return nil
}
