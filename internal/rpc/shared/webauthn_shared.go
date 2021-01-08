package rpc_shared

const WEBAUTHN_RPC_PORT = 1235

type Webauthn_GenericWebauthnBeginArgs struct {
	UserID int64
}

type Webauthn_DeleteWebauthnFinishArgs struct {
	UserID       int64
	WebauthnData string
}
