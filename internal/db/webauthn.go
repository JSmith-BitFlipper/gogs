package db

import (
	"encoding/gob"
	"fmt"
	"gorm.io/gorm"
	"time"

	"gogs.io/gogs/internal/conf"
	rpc_client "gogs.io/gogs/internal/rpc/client"
	rpc_shared "gogs.io/gogs/internal/rpc/shared"
	log "unknwon.dev/clog/v2"

	"webauthn/protocol"
	"webauthn/webauthn"
)

var WebauthnAPI *webauthn.WebAuthn

func InitWebauthnAPI() error {
	var err error
	WebauthnAPI, err = webauthn.New(&webauthn.Config{
		RPDisplayName: "Gogs",                     // Display Name for your site
		RPID:          conf.Server.URL.Hostname(), // Generally the domain name for your site
		RPOrigin:      conf.Server.ExternalURL,    // The origin URL for WebAuthn requests
	})

	// Register the `webauthn.SessionData` struct so that it can be stored in a web session
	gob.Register(webauthn.SessionData{})

	return err
}

// Webauthn is public key of a user.
type WebauthnEntry struct {
	UserID      int64     `xorm:"UNIQUE" gorm:"UNIQUE"`
	Created     time.Time `xorm:"-" gorm:"-" json:"-"`
	CreatedUnix int64

	PubKey    []byte `xorm:"VARCHAR(65) UNIQUE" gorm:"TYPE:VARCHAR(65);UNIQUE"`
	CredID    []byte `xorm:"VARCHAR(250) UNIQUE" gorm:"TYPE:VARCHAR(250);UNIQUE"`
	SignCount uint32 `xorm:"DEFAULT 0" gorm:"DEFAULT:0"`
	RPID      string `xorm:"rp_id VARCHAR(253)" gorm:"COLUMN:rp_id;TYPE:VARCHAR(253)"`
}

// WebauthnStore is the persistent interface for WebAuthn.
//
// NOTE: All methods are sorted in alphabetical order.
type WebauthnStore interface {
	// Create creates a new Webauthn credential entry for given user.
	Create(userID int64, credential webauthn.Credential) error

	// IsUserEnabled returns true if the user has enabled Webauthn 2FA.
	IsUserEnabled(userID int64) bool

	//
	// Private functions
	//

	// Query the database for the `WebauthnEntry`s of `userID`
	getCredentials(userID int64) ([]*WebauthnEntry, error)
}

var WebauthnEntries WebauthnStore

// NOTE: This is a GORM create hook.
func (t *WebauthnEntry) BeforeCreate(tx *gorm.DB) error {
	if t.CreatedUnix == 0 {
		t.CreatedUnix = tx.NowFunc().Unix()
	}
	return nil
}

// NOTE: This is a GORM query hook.
func (t *WebauthnEntry) AfterFind(tx *gorm.DB) error {
	t.Created = time.Unix(t.CreatedUnix, 0).Local()
	return nil
}

// Make sure `*webauthnEntries` implements `WebauthnStore`
var _ WebauthnStore = (*webauthnEntries)(nil)

type webauthnEntries struct {
	*gorm.DB
}

func (db *webauthnEntries) Create(userID int64, credential webauthn.Credential) error {
	wentry := &WebauthnEntry{
		UserID:    userID,
		PubKey:    credential.PublicKey,
		CredID:    credential.ID,
		SignCount: credential.Authenticator.SignCount,
		RPID:      "TODO",
	}

	return db.DB.Create(&wentry).Error
}

func (db *webauthnEntries) numCredentials(userID int64) (count int64) {
	err := db.Model(new(WebauthnEntry)).Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		log.Error("Failed to count webauthn entries [user_id: %d]: %v", userID, err)
	}
	return count
}

func (db *webauthnEntries) getCredentials(userID int64) ([]*WebauthnEntry, error) {
	ncreds := db.numCredentials(userID)
	entries := make([]*WebauthnEntry, 0, ncreds)

	err := db.Model(new(WebauthnEntry)).Where("user_id = ?", userID).Find(&entries).Error
	if err != nil {
		log.Error("Failed to get webauthn entries [user_id: %d]: %v", userID, err)
		return []*WebauthnEntry{}, err
	}

	return entries, nil
}

func (db *webauthnEntries) IsUserEnabled(userID int64) bool {
	return db.numCredentials(userID) > 0
}

// DeleteWebauthn removes Webauthn two-factor authentication entry from the database
func DeleteWebauthn(userID int64) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Where("user_id = ?", userID).Delete(new(WebauthnEntry)); err != nil {
		return fmt.Errorf("delete webauthn two-factor: %v", err)
	}

	return sess.Commit()
}

// The `GenericWebauthnBegin` initiates a transaction authentication assertion request
// without the extensions field filled, hence 'generic'
func Webauthn_GenericWebauthnBegin(userID int64) (reply *protocol.CredentialAssertion, err error) {
	// Call the RPC procedure for `GenericWebauthnBegin`
	args := &rpc_shared.Webauthn_GenericWebauthnBeginArgs{UserID: userID}
	reply = new(protocol.CredentialAssertion)

	err = rpc_client.Webauthn_GenericWebauthnBegin(args, reply)

	return
}

// TODO: Maybe move this to a specific file in db package with only RPCHandlers
//
// Get the transaction authentication details without any extensions
func RPCHandler_GenericWebauthnBegin(userID int64) (*protocol.CredentialAssertion, *webauthn.SessionData, error) {
	// Get the `user`
	user, err := GetUserByID(userID)

	// User doesn't exist
	if err != nil {
		log.Error(err.Error())
		return nil, nil, err
	}

	wuser, err := user.ToWebauthnUser()
	if err != nil {
		log.Error(err.Error())
		return nil, nil, err
	}

	// Generate PublicKeyCredentialRequestOptions, session data
	options, sessionData, err := WebauthnAPI.BeginLogin(wuser, nil)
	if err != nil {
		log.Error(err.Error())
		return nil, nil, err
	}

	return options, sessionData, nil
}
