package rpc_client

import (
	shared "gogs.io/gogs/internal/rpc/shared"

	"webauthn/protocol"
)

func Repo_GenericWebauthnBegin(args *shared.Repo_GenericWebauthnBeginArgs, reply *protocol.CredentialAssertion) error {
	return rpcCall("Repo.GenericWebauthnBegin", shared.REPO_RPC_PORT, args, reply)
}

func Repo_DeleteRepositoryFinish(args *shared.Repo_DeleteRepositoryFinishArgs, reply *interface{}) error {
	return rpcCall("Repo.DeleteRepositoryFinish", shared.REPO_RPC_PORT, args, reply)
}

func Repo_DeleteRepository(args *shared.Repo_DeleteRepositoryArgs, reply *interface{}) error {
	return rpcCall("Repo.DeleteRepository", shared.REPO_RPC_PORT, args, reply)
}

func Repo_DeleteMissingRepositories(args *interface{}, reply *interface{}) error {
	return rpcCall("Repo.DeleteMissingRepositories", shared.REPO_RPC_PORT, args, reply)
}
