package rpc_shared

const REPO_RPC_PORT = 1234

type Repo_GenericWebauthnBeginArgs struct {
	UserID int64
}

type Repo_DeleteRepositoryFinishArgs struct {
	UserID       int64
	OwnerID      int64
	RepoID       int64
	WebauthnData string
}

type Repo_DeleteRepositoryArgs struct {
	OwnerID, RepoID int64
}
