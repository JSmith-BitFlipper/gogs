package rpc_server

import (
	"errors"
	"fmt"
	"reflect"

	log "unknwon.dev/clog/v2"

	"gogs.io/gogs/internal/db"
	shared "gogs.io/gogs/internal/rpc/shared"

	"webauthn/protocol"
)

type Repo int

func (t *Repo) GenericWebauthnBegin(args *shared.Repo_GenericWebauthnBeginArgs, reply *protocol.CredentialAssertion) error {
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

func (t *Repo) DeleteRepositoryFinish(args *shared.Repo_DeleteRepositoryFinishArgs, reply *interface{}) error {
	// Extract the values passed in `args`
	userID := args.UserID
	ownerID := args.OwnerID
	repoID := args.RepoID
	webauthnData := args.WebauthnData

	// If Webauthn is not enabled, delete without any authentication check
	if !db.WebauthnEntries.IsUserEnabled(userID) {
		return db.RPCHandler_DeleteRepository(ownerID, repoID)
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
		repo, err := db.GetRepositoryByID(repoID)
		if err != nil {
			return err
		}

		expectedExtensions := protocol.AuthenticationExtensions{
			"txAuthSimple": fmt.Sprintf("Confirm deletion of repository: %s", repo.FullName()),
		}

		if !reflect.DeepEqual(expectedExtensions, clientDataExtensions) {
			return fmt.Errorf("Extensions verification failed: Expected %v, Received %v",
				expectedExtensions,
				clientDataExtensions)
		}

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

	// Delete the repository
	return db.RPCHandler_DeleteRepository(ownerID, repoID)
}

func (t *Repo) DeleteRepository(args *shared.Repo_DeleteRepositoryArgs, reply *interface{}) error {
	// Extract the values passed in `args`
	repoID := args.RepoID
	ownerID := args.OwnerID

	return db.RPCHandler_DeleteRepository(ownerID, repoID)
}
