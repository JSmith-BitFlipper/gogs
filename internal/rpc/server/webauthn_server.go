package rpc_server

import (
	"gogs.io/gogs/internal/db"
	shared "gogs.io/gogs/internal/rpc/shared"

	"webauthn/protocol"
)

type Webauthn int

func (t *Webauthn) GenericWebauthnBegin(args *shared.Webauthn_GenericWebauthnBeginArgs, reply *protocol.CredentialAssertion) error {
	// Extract the values passed in `args`
	userID := args.UserID

	options, sessionData, err := db.RPCHandler_GenericWebauthnBegin(userID)

	if err != nil {
		return err
	}

	// Save the `sessionData` for the current `userID`
	sessionMap[userID] = *sessionData

	// Return the `options` as the `reply`
	*reply = *options

	// Success!
	return nil
}
