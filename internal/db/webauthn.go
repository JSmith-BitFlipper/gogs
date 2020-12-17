package db

import (
	"encoding/gob"
	"gorm.io/gorm"
	"time"
	log "unknwon.dev/clog/v2"

	"gogs.io/gogs/internal/conf"

	"webauthn/webauthn"
)

var WebauthnAPI *webauthn.WebAuthn

func InitWebauthnAPI() error {
	var err error
	WebauthnAPI, err = webauthn.New(&webauthn.Config{
		RPDisplayName: "Gogs",                     // Display Name for your site
		RPID:          conf.Server.URL.Hostname(), // Generally the domain name for your site
		RPOrigin:      conf.Server.HTTPAddr,       // The origin URL for WebAuthn requests
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

	UKey      string `xorm:"VARCHAR(20) UNIQUE" gorm:"TYPE:VARCHAR(20);UNIQUE"`
	PubKey    string `xorm:"VARCHAR(65) UNIQUE" gorm:"TYPE:VARCHAR(65);UNIQUE"`
	CredID    string `xorm:"VARCHAR(250) UNIQUE" gorm:"TYPE:VARCHAR(250);UNIQUE"`
	SignCount int64  `xorm:"DEFAULT 0" gorm:"DEFAULT:0"`
	RPID      string `xorm:"VARCHAR(253)" gorm:"TYPE:VARCHAR(253)"`
}

// WebauthnStore is the persistent interface for WebAuthn.
//
// NOTE: All methods are sorted in alphabetical order.
type WebauthnStore interface {
	// Create creates a new Webauthn credential entry for given user.
	Create() error
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

var _ WebauthnStore = (*webauthnEntries)(nil)

type webauthnEntries struct {
	*gorm.DB
}

func (db *webauthnEntries) Create() error {
	log.Info("ADDED I'll take some from you tonight!")

	wentry := &WebauthnEntry{
		UserID:    69,
		UKey:      "damian",
		PubKey:    "is really cool!",
		CredID:    "but makes mistakes",
		SignCount: 420,
		RPID:      "but learns from them.",
	}

	return db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&wentry).Error
	})
}
