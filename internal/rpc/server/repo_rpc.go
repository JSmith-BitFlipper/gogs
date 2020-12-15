package rpc_server

import (
	"gogs.io/gogs/internal/db"
	shared "gogs.io/gogs/internal/rpc/shared"
)

type Repo int

func (t *Repo) DeleteRepository(args *shared.Repo_DeleteRepositoryArgs, reply *interface{}) error {
	// Extract the values passed in `args`
	repoID := args.RepoID
	ownerID := args.OwnerID

	return db.RPCHandler_DeleteRepository(ownerID, repoID)
}
