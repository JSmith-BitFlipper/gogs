package rpc_server

import (
	"errors"
	"fmt"

	log "unknwon.dev/clog/v2"

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

func (t *Webauthn) DeleteWebauthnFinish(args *shared.Webauthn_DeleteWebauthnFinishArgs, reply *interface{}) error {
	// Extract the values passed in `args`
	userID := args.UserID
	webauthnData := args.WebauthnData

	// If Webauthn is not enabled, this should error
	if !db.WebauthnEntries.IsUserEnabled(userID) {
		errText := fmt.Sprintf("Webauthn is already disabled, no credentials to delete.")
		log.Error(errText)
		return errors.New(errText)
	}

	// Load the `sessionData`
	sessionData, exists := sessionMap[userID]
	if !exists {
		errText := fmt.Sprintf("Session not found for user ID: %d", userID)
		log.Error(errText)
		return errors.New(errText)
	}

	u, err := db.GetUserByID(userID)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	// Get the webauthn user
	wuser, err := u.ToWebauthnUser()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	var verifyTxAuthSimple protocol.ExtensionsVerifier = func(_, clientDataExtensions protocol.AuthenticationExtensions) error {
		// TODO: Actually check here

		// Success!
		return nil
	}

	// TODO: In an actual implementation, we should perform additional checks on
	// the returned 'credential', i.e. check 'credential.Authenticator.CloneWarning'
	// and then increment the credentials counter
	_, err = db.WebauthnAPI.FinishLogin(wuser, sessionData, verifyTxAuthSimple, webauthnData)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	// Clear the session for this Webauthn authentication event
	delete(sessionMap, userID)

	// Delete the webauthn credentials
	return db.RPCHandler_DeleteWebauthn(userID)
}
