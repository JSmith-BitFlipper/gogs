package rpc_shared

type Repo_DeleteRepositoryBeginArgs struct {
	UserID int64
}

type Repo_DeleteRepositoryFinishArgs struct {
	UserID      int64
	OwnerID     int64
	RepoID      int64
	RequestData []byte
}

type Repo_DeleteRepositoryArgs struct {
	OwnerID, RepoID int64
}
