package webauthn

import (
	"gogs.io/gogs/internal/context"
	"gogs.io/gogs/internal/db"
)

func IsUserEnabled(c *context.APIContext) {
	enabled := db.WebauthnEntries.IsUserEnabled(c.User.ID)

	c.JSONSuccess(enabled)
	return
}
