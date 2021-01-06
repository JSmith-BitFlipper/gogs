package rpc_client

import (
	shared "gogs.io/gogs/internal/rpc/shared"

	"webauthn/protocol"
)

func Webauthn_GenericWebauthnBegin(args *shared.Webauthn_GenericWebauthnBeginArgs, reply *protocol.CredentialAssertion) error {
	return rpcCall("Webauthn.GenericWebauthnBegin", shared.WEBAUTHN_RPC_PORT, args, reply)
}
